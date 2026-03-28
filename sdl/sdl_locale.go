package sdl

/*
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// Locale describes a locale with language and optional country.
type Locale struct {
	Language string
	Country  string
}

// GetPreferredLocales returns the user's preferred locales.
func GetPreferredLocales() []Locale {
	var count C.int
	cls := C.SDL_GetPreferredLocales(&count)
	if cls == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cls))
	n := int(count)
	result := make([]Locale, n)
	slice := unsafe.Slice((**C.SDL_Locale)(unsafe.Pointer(cls)), n)
	for i, cl := range slice {
		result[i].Language = C.GoString(cl.language)
		if cl.country != nil {
			result[i].Country = C.GoString(cl.country)
		}
	}
	return result
}
