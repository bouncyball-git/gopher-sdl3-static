package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// GUID represents a 128-bit identifier.
type GUID struct {
	Data [16]byte
}

func (g GUID) cval() C.SDL_GUID {
	var cg C.SDL_GUID
	for i := 0; i < 16; i++ {
		cg.data[i] = C.Uint8(g.Data[i])
	}
	return cg
}

func guidFromC(cg C.SDL_GUID) GUID {
	var g GUID
	for i := 0; i < 16; i++ {
		g.Data[i] = byte(cg.data[i])
	}
	return g
}

// String returns the GUID as a string.
func (g GUID) String() string {
	cg := g.cval()
	buf := make([]C.char, 33)
	C.SDL_GUIDToString(cg, &buf[0], 33)
	return C.GoString(&buf[0])
}

// GUIDFromString converts a string to a GUID.
func GUIDFromString(s string) GUID {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	return guidFromC(C.SDL_StringToGUID(cs))
}
