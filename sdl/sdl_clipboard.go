package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>

extern void *goClipboardDataCallback(void *userdata, char *mime_type, size_t *size);
extern void goClipboardCleanupCallback(void *userdata);

static const void *cgoClipboardDataTrampoline(void *userdata, const char *mime_type, size_t *size) {
	return goClipboardDataCallback(userdata, (char *)mime_type, size);
}

static void cgoClipboardCleanupTrampoline(void *userdata) {
	goClipboardCleanupCallback(userdata);
}
*/
import "C"

import "unsafe"

// SetClipboardText sets the clipboard to the specified text.
func SetClipboardText(text string) error {
	ct := C.CString(text)
	defer C.free(unsafe.Pointer(ct))
	if !C.SDL_SetClipboardText(ct) {
		return getError()
	}
	return nil
}

// GetClipboardText returns the clipboard text.
func GetClipboardText() string {
	ct := C.SDL_GetClipboardText()
	if ct == nil {
		return ""
	}
	s := C.GoString(ct)
	C.SDL_free(unsafe.Pointer(ct))
	return s
}

// HasClipboardText returns true if the clipboard has text.
func HasClipboardText() bool {
	return bool(C.SDL_HasClipboardText())
}

// SetPrimarySelectionText sets the primary selection to the specified text.
func SetPrimarySelectionText(text string) error {
	ct := C.CString(text)
	defer C.free(unsafe.Pointer(ct))
	if !C.SDL_SetPrimarySelectionText(ct) {
		return getError()
	}
	return nil
}

// GetPrimarySelectionText returns the primary selection text.
func GetPrimarySelectionText() string {
	ct := C.SDL_GetPrimarySelectionText()
	if ct == nil {
		return ""
	}
	s := C.GoString(ct)
	C.SDL_free(unsafe.Pointer(ct))
	return s
}

// HasPrimarySelectionText returns true if the primary selection has text.
func HasPrimarySelectionText() bool {
	return bool(C.SDL_HasPrimarySelectionText())
}

// ClearClipboardData clears the clipboard data.
func ClearClipboardData() error {
	if !C.SDL_ClearClipboardData() {
		return getError()
	}
	return nil
}

// GetClipboardData returns the clipboard data for a given MIME type.
// Caller receives a Go byte slice copy; the C allocation is freed.
func GetClipboardData(mimeType string) ([]byte, error) {
	cm := C.CString(mimeType)
	defer C.free(unsafe.Pointer(cm))
	var size C.size_t
	data := C.SDL_GetClipboardData(cm, &size)
	if data == nil {
		return nil, getError()
	}
	defer C.SDL_free(data)
	result := make([]byte, int(size))
	copy(result, unsafe.Slice((*byte)(data), int(size)))
	return result, nil
}

// HasClipboardData returns whether the clipboard has data for a given MIME type.
func HasClipboardData(mimeType string) bool {
	cm := C.CString(mimeType)
	defer C.free(unsafe.Pointer(cm))
	return bool(C.SDL_HasClipboardData(cm))
}

// GetClipboardMimeTypes returns the MIME types available on the clipboard.
func GetClipboardMimeTypes() []string {
	var count C.size_t
	ctypes := C.SDL_GetClipboardMimeTypes(&count)
	if ctypes == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(ctypes))
	n := int(count)
	result := make([]string, n)
	slice := unsafe.Slice((**C.char)(unsafe.Pointer(ctypes)), n)
	for i, s := range slice {
		result[i] = C.GoString(s)
	}
	return result
}

// ClipboardDataFunc is called when clipboard data is requested.
// It must return the data for the given MIME type.
type ClipboardDataFunc func(mimeType string) []byte

// clipboardDataHolder stores data returned by clipboard callbacks so the C pointer remains valid.
var clipboardDataHolder []byte

//export goClipboardDataCallback
func goClipboardDataCallback(userdata unsafe.Pointer, mimeType *C.char, size *C.size_t) unsafe.Pointer {
	id := uintptr(userdata)
	fn := getCallback(id).(ClipboardDataFunc)
	data := fn(C.GoString(mimeType))
	if data == nil {
		*size = 0
		return nil
	}
	clipboardDataHolder = data
	*size = C.size_t(len(data))
	return unsafe.Pointer(&clipboardDataHolder[0])
}

//export goClipboardCleanupCallback
func goClipboardCleanupCallback(userdata unsafe.Pointer) {
	id := uintptr(userdata)
	unregisterCallback(id)
	clipboardDataHolder = nil
}

// SetClipboardData offers clipboard data to other applications.
// The callback is called when another application requests the data.
func SetClipboardData(callback ClipboardDataFunc, mimeTypes []string) error {
	id := registerCallback(callback)
	cmimes := make([]*C.char, len(mimeTypes))
	for i, m := range mimeTypes {
		cmimes[i] = C.CString(m)
	}
	defer func() {
		for _, cm := range cmimes {
			C.free(unsafe.Pointer(cm))
		}
	}()
	var cmPtr **C.char
	if len(cmimes) > 0 {
		cmPtr = &cmimes[0]
	}
	if !C.SDL_SetClipboardData(
		C.SDL_ClipboardDataCallback(C.cgoClipboardDataTrampoline),
		C.SDL_ClipboardCleanupCallback(C.cgoClipboardCleanupTrampoline),
		ptrFromID(id), cmPtr, C.size_t(len(mimeTypes))) {
		unregisterCallback(id)
		return getError()
	}
	return nil
}
