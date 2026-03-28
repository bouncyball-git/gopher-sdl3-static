package sdl

/*
#include <SDL3/SDL.h>
*/
import "C"

// PowerState represents the battery/power state.
// Note: PowerState is also defined in sdl_joystick.go since both joystick
// and power APIs share it. This file just provides the GetPowerInfo function.

// GetPowerInfo returns the current power state and battery information.
func GetPowerInfo() (state PowerState, seconds, percent int) {
	var cs, cp C.int
	ps := C.SDL_GetPowerInfo(&cs, &cp)
	return PowerState(ps), int(cs), int(cp)
}
