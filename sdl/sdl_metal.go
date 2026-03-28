package sdl

/*
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// MetalView is a handle to a CAMetalLayer-backed NSView (macOS) or UIView (iOS/tvOS).
type MetalView unsafe.Pointer

// Metal_CreateView creates a CAMetalLayer-backed NSView/UIView and attaches it to the specified window.
func Metal_CreateView(window *Window) MetalView {
	return MetalView(C.SDL_Metal_CreateView(window.c))
}

// Metal_DestroyView destroys an existing MetalView object.
func Metal_DestroyView(view MetalView) {
	C.SDL_Metal_DestroyView(C.SDL_MetalView(view))
}

// Metal_GetLayer returns a pointer to the backing CAMetalLayer for the given view.
func Metal_GetLayer(view MetalView) unsafe.Pointer {
	return C.SDL_Metal_GetLayer(C.SDL_MetalView(view))
}
