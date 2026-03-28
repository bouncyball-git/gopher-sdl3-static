package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// SharedObject represents a loaded shared library.
type SharedObject struct {
	c *C.SDL_SharedObject
}

// LoadObject loads a shared object.
func LoadObject(sofile string) (*SharedObject, error) {
	cs := C.CString(sofile)
	defer C.free(unsafe.Pointer(cs))
	co := C.SDL_LoadObject(cs)
	if co == nil {
		return nil, getError()
	}
	return &SharedObject{c: co}, nil
}

// LoadFunction looks up a function in the shared object.
func (so *SharedObject) LoadFunction(name string) (unsafe.Pointer, error) {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	fp := C.SDL_LoadFunction(so.c, cn)
	if fp == nil {
		return nil, getError()
	}
	return unsafe.Pointer(fp), nil
}

// Unload unloads the shared object.
func (so *SharedObject) Unload() {
	if so.c != nil {
		C.SDL_UnloadObject(so.c)
		so.c = nil
	}
}
