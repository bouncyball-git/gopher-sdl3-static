package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>

extern void goDialogFileCallback(void *userdata, char **filelist, int filter);

static void cgoDialogFileCallbackTrampoline(void *userdata, const char * const *filelist, int filter) {
	goDialogFileCallback(userdata, (char **)filelist, filter);
}

static inline SDL_DialogFileCallback _get_dialog_trampoline(void) {
	return cgoDialogFileCallbackTrampoline;
}
*/
import "C"

import (
	"unsafe"
)

//go:nosplit
func ptrFromID(id uintptr) unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&id))
}

// FileDialogType represents the type of file dialog.
type FileDialogType int

// File dialog type constants.
const (
	FILEDIALOG_OPENFILE   FileDialogType = C.SDL_FILEDIALOG_OPENFILE
	FILEDIALOG_SAVEFILE   FileDialogType = C.SDL_FILEDIALOG_SAVEFILE
	FILEDIALOG_OPENFOLDER FileDialogType = C.SDL_FILEDIALOG_OPENFOLDER
)

// DialogFileFilter describes a file filter for file dialogs.
type DialogFileFilter struct {
	Name    string
	Pattern string
}

// DialogFileCallback is the type for file dialog result callbacks.
type DialogFileCallback func(files []string, filter int)

//export goDialogFileCallback
func goDialogFileCallback(userdata unsafe.Pointer, filelist **C.char, filter C.int) {
	id := uintptr(userdata)
	fn := getCallback(id).(DialogFileCallback)
	defer unregisterCallback(id)

	var files []string
	if filelist != nil {
		for ptr := filelist; *ptr != nil; ptr = (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(filelist))) {
			files = append(files, C.GoString(*ptr))
		}
	}
	fn(files, int(filter))
}

// ShowOpenFileDialog displays an open file dialog.
func ShowOpenFileDialog(callback DialogFileCallback, window *Window, filters []DialogFileFilter, defaultLocation string, allowMany bool) {
	id := registerCallback(callback)
	var cw *C.SDL_Window
	if window != nil {
		cw = window.c
	}
	var cfilters *C.SDL_DialogFileFilter
	var cnames, cpatterns []*C.char
	if len(filters) > 0 {
		cf := make([]C.SDL_DialogFileFilter, len(filters))
		cnames = make([]*C.char, len(filters))
		cpatterns = make([]*C.char, len(filters))
		for i, f := range filters {
			cnames[i] = C.CString(f.Name)
			cpatterns[i] = C.CString(f.Pattern)
			cf[i].name = cnames[i]
			cf[i].pattern = cpatterns[i]
		}
		cfilters = &cf[0]
		defer func() {
			for i := range filters {
				C.free(unsafe.Pointer(cnames[i]))
				C.free(unsafe.Pointer(cpatterns[i]))
			}
		}()
	}
	var cdl *C.char
	if defaultLocation != "" {
		cdl = C.CString(defaultLocation)
		defer C.free(unsafe.Pointer(cdl))
	}
	C.SDL_ShowOpenFileDialog(C._get_dialog_trampoline(), ptrFromID(id), cw, cfilters, C.int(len(filters)), cdl, C.bool(allowMany))
}

// ShowSaveFileDialog displays a save file dialog.
func ShowSaveFileDialog(callback DialogFileCallback, window *Window, filters []DialogFileFilter, defaultLocation string) {
	id := registerCallback(callback)
	var cw *C.SDL_Window
	if window != nil {
		cw = window.c
	}
	var cfilters *C.SDL_DialogFileFilter
	var cnames, cpatterns []*C.char
	if len(filters) > 0 {
		cf := make([]C.SDL_DialogFileFilter, len(filters))
		cnames = make([]*C.char, len(filters))
		cpatterns = make([]*C.char, len(filters))
		for i, f := range filters {
			cnames[i] = C.CString(f.Name)
			cpatterns[i] = C.CString(f.Pattern)
			cf[i].name = cnames[i]
			cf[i].pattern = cpatterns[i]
		}
		cfilters = &cf[0]
		defer func() {
			for i := range filters {
				C.free(unsafe.Pointer(cnames[i]))
				C.free(unsafe.Pointer(cpatterns[i]))
			}
		}()
	}
	var cdl *C.char
	if defaultLocation != "" {
		cdl = C.CString(defaultLocation)
		defer C.free(unsafe.Pointer(cdl))
	}
	C.SDL_ShowSaveFileDialog(C._get_dialog_trampoline(), ptrFromID(id), cw, cfilters, C.int(len(filters)), cdl)
}

// ShowOpenFolderDialog displays an open folder dialog.
func ShowOpenFolderDialog(callback DialogFileCallback, window *Window, defaultLocation string, allowMany bool) {
	id := registerCallback(callback)
	var cw *C.SDL_Window
	if window != nil {
		cw = window.c
	}
	var cdl *C.char
	if defaultLocation != "" {
		cdl = C.CString(defaultLocation)
		defer C.free(unsafe.Pointer(cdl))
	}
	C.SDL_ShowOpenFolderDialog(C._get_dialog_trampoline(), ptrFromID(id), cw, cdl, C.bool(allowMany))
}

// ShowFileDialogWithProperties displays a file dialog with the specified properties.
func ShowFileDialogWithProperties(dialogType FileDialogType, callback DialogFileCallback, props PropertiesID) {
	id := registerCallback(callback)
	C.SDL_ShowFileDialogWithProperties(C.SDL_FileDialogType(dialogType), C._get_dialog_trampoline(), ptrFromID(id), C.SDL_PropertiesID(props))
}
