package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// TouchID is a unique identifier for a touch device.
type TouchID uint64

// FingerID is a unique identifier for a single finger on a touch device.
type FingerID uint64

// TouchDeviceType describes the type of a touch device.
type TouchDeviceType int32

// Touch device type constants.
const (
	TOUCH_DEVICE_INVALID           TouchDeviceType = C.SDL_TOUCH_DEVICE_INVALID
	TOUCH_DEVICE_DIRECT            TouchDeviceType = C.SDL_TOUCH_DEVICE_DIRECT
	TOUCH_DEVICE_INDIRECT_ABSOLUTE TouchDeviceType = C.SDL_TOUCH_DEVICE_INDIRECT_ABSOLUTE
	TOUCH_DEVICE_INDIRECT_RELATIVE TouchDeviceType = C.SDL_TOUCH_DEVICE_INDIRECT_RELATIVE
)

// Finger contains data about a single finger in a multitouch event.
type Finger struct {
	ID       FingerID
	X        float32
	Y        float32
	Pressure float32
}

// GetTouchDevices returns a list of registered touch device IDs.
func GetTouchDevices() []TouchID {
	var count C.int
	cids := C.SDL_GetTouchDevices(&count)
	if cids == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cids))
	n := int(count)
	result := make([]TouchID, n)
	slice := unsafe.Slice((*C.SDL_TouchID)(cids), n)
	for i, id := range slice {
		result[i] = TouchID(id)
	}
	return result
}

// GetTouchDeviceName returns the name of a touch device.
func GetTouchDeviceName(touchID TouchID) string {
	return C.GoString(C.SDL_GetTouchDeviceName(C.SDL_TouchID(touchID)))
}

// GetTouchDeviceType returns the type of the given touch device.
func GetTouchDeviceType(touchID TouchID) TouchDeviceType {
	return TouchDeviceType(C.SDL_GetTouchDeviceType(C.SDL_TouchID(touchID)))
}

// GetTouchFingers returns a list of active fingers for a given touch device.
func GetTouchFingers(touchID TouchID) []Finger {
	var count C.int
	cfingers := C.SDL_GetTouchFingers(C.SDL_TouchID(touchID), &count)
	if cfingers == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cfingers))
	n := int(count)
	result := make([]Finger, n)
	slice := unsafe.Slice((**C.SDL_Finger)(cfingers), n)
	for i, cf := range slice {
		result[i] = Finger{
			ID:       FingerID(cf.id),
			X:        float32(cf.x),
			Y:        float32(cf.y),
			Pressure: float32(cf.pressure),
		}
	}
	return result
}

// Special IDs for touch-simulated mouse events and vice versa.
const (
	TOUCH_MOUSEID = C.SDL_TOUCH_MOUSEID
	MOUSE_TOUCHID = C.SDL_MOUSE_TOUCHID
)
