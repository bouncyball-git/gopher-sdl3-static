// Package sdl provides Go bindings for SDL3 (Simple DirectMedia Layer).
package sdl

/*
#cgo CFLAGS: -I${SRCDIR}/../sdl-3.5.0-static/include
#cgo LDFLAGS: -L${SRCDIR}/../sdl-3.5.0-static/lib -lSDL3 -lm -pthread
#cgo linux LDFLAGS: -ldl
#define SDL_THREAD_SAFETY_ANALYSIS 0
#include <SDL3/SDL.h>
*/
import "C"

import "runtime"

func init() {
	runtime.LockOSThread()
}

const (
	WINDOWPOS_UNDEFINED = int(0x1FFF0000)
	WINDOWPOS_CENTERED  = int(0x2FFF0000)
)
