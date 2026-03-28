package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// OpenURL opens a URL in the default browser.
func OpenURL(url string) error {
	curl := C.CString(url)
	defer C.free(unsafe.Pointer(curl))
	if !C.SDL_OpenURL(curl) {
		return getError()
	}
	return nil
}
