package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>

static bool _sdl_set_error(const char *msg) {
	return SDL_SetError("%s", msg);
}
*/
import "C"

import (
	"errors"
	"unsafe"
)

// GetError returns the last SDL error message.
func GetError() string {
	return C.GoString(C.SDL_GetError())
}

// ClearError clears the current error message.
func ClearError() {
	C.SDL_ClearError()
}

func getError() error {
	msg := C.SDL_GetError()
	if msg == nil || *msg == 0 {
		return errors.New("sdl: unknown error")
	}
	return errors.New("sdl: " + C.GoString(msg))
}

// SetError sets the SDL error message for the current thread.
func SetError(msg string) {
	cm := C.CString(msg)
	defer C.free(unsafe.Pointer(cm))
	C._sdl_set_error(cm)
}

// OutOfMemory sets an error indicating that memory allocation failed.
func OutOfMemory() {
	C.SDL_OutOfMemory()
}

// Note: SDL_SetErrorV is not wrapped because it takes a C va_list argument
// which has no equivalent in Go. Use SetError instead.
