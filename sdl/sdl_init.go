package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>

extern void goMainThreadCallback(void *userdata);

static void cgoMainThreadCallbackTrampoline(void *userdata) {
	goMainThreadCallback(userdata);
}
*/
import "C"

import "unsafe"

// InitFlags specifies which subsystems to initialize.
type InitFlags uint32

const (
	INIT_AUDIO    InitFlags = C.SDL_INIT_AUDIO
	INIT_VIDEO    InitFlags = C.SDL_INIT_VIDEO
	INIT_JOYSTICK InitFlags = C.SDL_INIT_JOYSTICK
	INIT_HAPTIC   InitFlags = C.SDL_INIT_HAPTIC
	INIT_GAMEPAD  InitFlags = C.SDL_INIT_GAMEPAD
	INIT_EVENTS   InitFlags = C.SDL_INIT_EVENTS
	INIT_SENSOR   InitFlags = C.SDL_INIT_SENSOR
	INIT_CAMERA   InitFlags = C.SDL_INIT_CAMERA
)

// AppResult represents the return value from SDL app callbacks.
type AppResult int

const (
	APP_CONTINUE AppResult = C.SDL_APP_CONTINUE
	APP_SUCCESS  AppResult = C.SDL_APP_SUCCESS
	APP_FAILURE  AppResult = C.SDL_APP_FAILURE
)

// Init initializes the SDL library.
func Init(flags InitFlags) error {
	if !C.SDL_Init(C.SDL_InitFlags(flags)) {
		return getError()
	}
	return nil
}

// InitSubSystem initializes specific SDL subsystems.
func InitSubSystem(flags InitFlags) error {
	if !C.SDL_InitSubSystem(C.SDL_InitFlags(flags)) {
		return getError()
	}
	return nil
}

// QuitSubSystem shuts down specific SDL subsystems.
func QuitSubSystem(flags InitFlags) {
	C.SDL_QuitSubSystem(C.SDL_InitFlags(flags))
}

// WasInit returns the subsystems which have previously been initialized.
func WasInit(flags InitFlags) InitFlags {
	return InitFlags(C.SDL_WasInit(C.SDL_InitFlags(flags)))
}

// Quit cleans up all initialized subsystems.
func Quit() {
	C.SDL_Quit()
}

// IsMainThread returns true if the current thread is the main thread.
func IsMainThread() bool {
	return bool(C.SDL_IsMainThread())
}

// SetAppMetadata sets the application metadata.
func SetAppMetadata(appname, appversion, appidentifier string) error {
	cn := C.CString(appname)
	cv := C.CString(appversion)
	ci := C.CString(appidentifier)
	defer C.free(unsafe.Pointer(cn))
	defer C.free(unsafe.Pointer(cv))
	defer C.free(unsafe.Pointer(ci))
	if !C.SDL_SetAppMetadata(cn, cv, ci) {
		return getError()
	}
	return nil
}

// SetAppMetadataProperty sets a metadata property for the application.
func SetAppMetadataProperty(name, value string) error {
	cn := C.CString(name)
	cv := C.CString(value)
	defer C.free(unsafe.Pointer(cn))
	defer C.free(unsafe.Pointer(cv))
	if !C.SDL_SetAppMetadataProperty(cn, cv) {
		return getError()
	}
	return nil
}

// GetAppMetadataProperty returns a metadata property for the application.
func GetAppMetadataProperty(name string) string {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	return C.GoString(C.SDL_GetAppMetadataProperty(cn))
}

// MainThreadCallback is a function to run on the main thread.
type MainThreadCallback func()

//export goMainThreadCallback
func goMainThreadCallback(userdata unsafe.Pointer) {
	id := uintptr(userdata)
	fn := getCallback(id).(MainThreadCallback)
	unregisterCallback(id)
	fn()
}

// RunOnMainThread runs a callback on the main thread.
// If waitComplete is true, the function blocks until the callback finishes.
func RunOnMainThread(callback MainThreadCallback, waitComplete bool) error {
	id := registerCallback(callback)
	if !C.SDL_RunOnMainThread(C.SDL_MainThreadCallback(C.cgoMainThreadCallbackTrampoline), ptrFromID(id), C.bool(waitComplete)) {
		unregisterCallback(id)
		return getError()
	}
	return nil
}

// App metadata PROP_* constants are in sdl_props_constants.go.
