package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"

// PenID is a unique identifier for a pen device.
type PenID uint32

// PenInputFlags represents pen input state flags.
type PenInputFlags uint32

// Pen input flag constants.
const (
	PEN_INPUT_DOWN         PenInputFlags = C.SDL_PEN_INPUT_DOWN
	PEN_INPUT_BUTTON_1     PenInputFlags = C.SDL_PEN_INPUT_BUTTON_1
	PEN_INPUT_BUTTON_2     PenInputFlags = C.SDL_PEN_INPUT_BUTTON_2
	PEN_INPUT_BUTTON_3     PenInputFlags = C.SDL_PEN_INPUT_BUTTON_3
	PEN_INPUT_BUTTON_4     PenInputFlags = C.SDL_PEN_INPUT_BUTTON_4
	PEN_INPUT_BUTTON_5     PenInputFlags = C.SDL_PEN_INPUT_BUTTON_5
	PEN_INPUT_ERASER_TIP   PenInputFlags = C.SDL_PEN_INPUT_ERASER_TIP
	PEN_INPUT_IN_PROXIMITY PenInputFlags = C.SDL_PEN_INPUT_IN_PROXIMITY
)

// PenAxis represents a pen axis index.
type PenAxis int32

// Pen axis constants.
const (
	PEN_AXIS_PRESSURE              PenAxis = C.SDL_PEN_AXIS_PRESSURE
	PEN_AXIS_XTILT                 PenAxis = C.SDL_PEN_AXIS_XTILT
	PEN_AXIS_YTILT                 PenAxis = C.SDL_PEN_AXIS_YTILT
	PEN_AXIS_DISTANCE              PenAxis = C.SDL_PEN_AXIS_DISTANCE
	PEN_AXIS_ROTATION              PenAxis = C.SDL_PEN_AXIS_ROTATION
	PEN_AXIS_SLIDER                PenAxis = C.SDL_PEN_AXIS_SLIDER
	PEN_AXIS_TANGENTIAL_PRESSURE   PenAxis = C.SDL_PEN_AXIS_TANGENTIAL_PRESSURE
	PEN_AXIS_COUNT                 PenAxis = C.SDL_PEN_AXIS_COUNT
)

// PenDeviceType describes the type of a pen device.
type PenDeviceType int32

// Pen device type constants.
const (
	PEN_DEVICE_TYPE_INVALID  PenDeviceType = C.SDL_PEN_DEVICE_TYPE_INVALID
	PEN_DEVICE_TYPE_UNKNOWN  PenDeviceType = C.SDL_PEN_DEVICE_TYPE_UNKNOWN
	PEN_DEVICE_TYPE_DIRECT   PenDeviceType = C.SDL_PEN_DEVICE_TYPE_DIRECT
	PEN_DEVICE_TYPE_INDIRECT PenDeviceType = C.SDL_PEN_DEVICE_TYPE_INDIRECT
)

// GetPenDeviceType returns the device type of the given pen.
func GetPenDeviceType(instanceID PenID) PenDeviceType {
	return PenDeviceType(C.SDL_GetPenDeviceType(C.SDL_PenID(instanceID)))
}

// Special IDs for pen-simulated mouse/touch events.
const (
	PEN_MOUSEID = C.SDL_PEN_MOUSEID
	PEN_TOUCHID = C.SDL_PEN_TOUCHID
)
