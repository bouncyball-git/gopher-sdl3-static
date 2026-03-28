package sdl

/*
#include <SDL3/SDL.h>

extern Uint32 goTimerCallback(void *userdata, SDL_TimerID timerID, Uint32 interval);
extern Uint64 goTimerNSCallback(void *userdata, SDL_TimerID timerID, Uint64 interval);

static Uint32 cgoTimerCallbackTrampoline(void *userdata, SDL_TimerID timerID, Uint32 interval) {
    return goTimerCallback(userdata, timerID, interval);
}

static Uint64 cgoTimerNSCallbackTrampoline(void *userdata, SDL_TimerID timerID, Uint64 interval) {
    return goTimerNSCallback(userdata, timerID, interval);
}
*/
import "C"

import "unsafe"

// GetTicks returns the number of milliseconds since SDL library initialization.
func GetTicks() uint64 {
	return uint64(C.SDL_GetTicks())
}

// GetTicksNS returns the number of nanoseconds since SDL library initialization.
func GetTicksNS() uint64 {
	return uint64(C.SDL_GetTicksNS())
}

// GetPerformanceCounter returns the current value of the high-resolution counter.
func GetPerformanceCounter() uint64 {
	return uint64(C.SDL_GetPerformanceCounter())
}

// GetPerformanceFrequency returns the frequency of the high-resolution counter.
func GetPerformanceFrequency() uint64 {
	return uint64(C.SDL_GetPerformanceFrequency())
}

// Delay waits the specified number of milliseconds.
func Delay(ms uint32) {
	C.SDL_Delay(C.Uint32(ms))
}

// DelayNS waits the specified number of nanoseconds.
func DelayNS(ns uint64) {
	C.SDL_DelayNS(C.Uint64(ns))
}

// DelayPrecise waits precisely the specified number of nanoseconds.
func DelayPrecise(ns uint64) {
	C.SDL_DelayPrecise(C.Uint64(ns))
}

// TimerID is the unique identifier for a timer.
type TimerID uint32

// TimerCallback is the callback function for millisecond timers.
// Return the next interval in milliseconds, or 0 to cancel.
type TimerCallback func(timerID TimerID, interval uint32) uint32

// TimerNSCallback is the callback function for nanosecond timers.
// Return the next interval in nanoseconds, or 0 to cancel.
type TimerNSCallback func(timerID TimerID, interval uint64) uint64

//export goTimerCallback
func goTimerCallback(userdata unsafe.Pointer, timerID C.SDL_TimerID, interval C.Uint32) C.Uint32 {
	id := uintptr(userdata)
	fn := getCallback(id).(TimerCallback)
	ret := fn(TimerID(timerID), uint32(interval))
	if ret == 0 {
		unregisterCallback(id)
	}
	return C.Uint32(ret)
}

//export goTimerNSCallback
func goTimerNSCallback(userdata unsafe.Pointer, timerID C.SDL_TimerID, interval C.Uint64) C.Uint64 {
	id := uintptr(userdata)
	fn := getCallback(id).(TimerNSCallback)
	ret := fn(TimerID(timerID), uint64(interval))
	if ret == 0 {
		unregisterCallback(id)
	}
	return C.Uint64(ret)
}

// AddTimer sets a timer that fires a callback after the given interval in milliseconds.
func AddTimer(interval uint32, callback TimerCallback) (TimerID, error) {
	id := registerCallback(callback)
	tid := C.SDL_AddTimer(C.Uint32(interval), C.SDL_TimerCallback(C.cgoTimerCallbackTrampoline), ptrFromID(id))
	if tid == 0 {
		unregisterCallback(id)
		return 0, getError()
	}
	return TimerID(tid), nil
}

// AddTimerNS sets a timer that fires a callback after the given interval in nanoseconds.
func AddTimerNS(interval uint64, callback TimerNSCallback) (TimerID, error) {
	id := registerCallback(callback)
	tid := C.SDL_AddTimerNS(C.Uint64(interval), C.SDL_NSTimerCallback(C.cgoTimerNSCallbackTrampoline), ptrFromID(id))
	if tid == 0 {
		unregisterCallback(id)
		return 0, getError()
	}
	return TimerID(tid), nil
}

// RemoveTimer removes a timer created with AddTimer or AddTimerNS.
func RemoveTimer(id TimerID) error {
	if !C.SDL_RemoveTimer(C.SDL_TimerID(id)) {
		return getError()
	}
	return nil
}

// Time conversion constants.
const (
	MS_PER_SECOND  = 1000
	US_PER_SECOND  = 1000000
	NS_PER_SECOND  = 1000000000
	NS_PER_MS      = 1000000
	NS_PER_US      = 1000
	SECONDS_TO_NS  = NS_PER_SECOND
	NS_TO_SECONDS  = 1.0 / NS_PER_SECOND
	MS_TO_NS       = NS_PER_MS
	NS_TO_MS       = 1.0 / NS_PER_MS
	US_TO_NS       = NS_PER_US
	NS_TO_US       = 1.0 / NS_PER_US
)
