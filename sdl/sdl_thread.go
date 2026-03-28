package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>

extern int goThreadFunction(void *data);

static SDL_Thread *_sdl_create_thread(const char *name, void *data) {
	return SDL_CreateThread(goThreadFunction, name, data);
}

static SDL_Thread *_sdl_create_thread_with_properties(SDL_PropertiesID props) {
	return SDL_CreateThreadWithProperties(props);
}
*/
import "C"

import "unsafe"

// Thread represents an SDL thread.
type Thread struct {
	c *C.SDL_Thread
}

// ThreadID is a unique numeric ID that identifies a thread.
type ThreadID uint64

// ThreadPriority represents the SDL thread priority.
type ThreadPriority int

// Thread priority levels.
const (
	THREAD_PRIORITY_LOW           ThreadPriority = C.SDL_THREAD_PRIORITY_LOW
	THREAD_PRIORITY_NORMAL        ThreadPriority = C.SDL_THREAD_PRIORITY_NORMAL
	THREAD_PRIORITY_HIGH          ThreadPriority = C.SDL_THREAD_PRIORITY_HIGH
	THREAD_PRIORITY_TIME_CRITICAL ThreadPriority = C.SDL_THREAD_PRIORITY_TIME_CRITICAL
)

// ThreadState represents the current state of a thread.
type ThreadState int

// Thread states.
const (
	THREAD_STATE_UNKNOWN  ThreadState = C.SDL_THREAD_UNKNOWN
	THREAD_STATE_ALIVE    ThreadState = C.SDL_THREAD_ALIVE
	THREAD_STATE_DETACHED ThreadState = C.SDL_THREAD_DETACHED
	THREAD_STATE_COMPLETE ThreadState = C.SDL_THREAD_COMPLETE
)

// GetCurrentThreadID returns the thread identifier for the current thread.
func GetCurrentThreadID() ThreadID {
	return ThreadID(C.SDL_GetCurrentThreadID())
}

// SetCurrentThreadPriority sets the priority for the current thread.
func SetCurrentThreadPriority(priority ThreadPriority) error {
	if !C.SDL_SetCurrentThreadPriority(C.SDL_ThreadPriority(priority)) {
		return getError()
	}
	return nil
}

// Name returns the name of the thread as specified in SDL_CreateThread.
func (t *Thread) Name() string {
	return C.GoString(C.SDL_GetThreadName(t.c))
}

// ID returns the thread identifier for this thread.
func (t *Thread) ID() ThreadID {
	return ThreadID(C.SDL_GetThreadID(t.c))
}

// State returns the current state of the thread.
func (t *Thread) State() ThreadState {
	return ThreadState(C.SDL_GetThreadState(t.c))
}

// Wait waits for the thread to finish and returns the thread's return code.
func (t *Thread) Wait() int {
	var status C.int
	C.SDL_WaitThread(t.c, &status)
	t.c = nil
	return int(status)
}

// Detach lets the thread clean up on exit without intervention.
func (t *Thread) Detach() {
	C.SDL_DetachThread(t.c)
	t.c = nil
}

// TLSID represents a thread local storage identifier.
type TLSID struct {
	c C.SDL_TLSID
}

// GetTLS returns the value associated with a TLS ID for the current thread.
func GetTLS(id *TLSID) unsafe.Pointer {
	return C.SDL_GetTLS(&id.c)
}

// SetTLS sets the value associated with a TLS ID for the current thread.
func SetTLS(id *TLSID, value unsafe.Pointer) error {
	if !C.SDL_SetTLS(&id.c, value, nil) {
		return getError()
	}
	return nil
}

// CleanupTLS cleans up all TLS data for the current thread.
func CleanupTLS() {
	C.SDL_CleanupTLS()
}

// ThreadFunction is a Go function to run in an SDL thread.
type ThreadFunction func() int

//export goThreadFunction
func goThreadFunction(data unsafe.Pointer) C.int {
	id := uintptr(data)
	fn := getCallback(id).(ThreadFunction)
	// Don't unregister here - the callback stays alive for the thread's lifetime.
	// It gets cleaned up when the Thread is waited on or detached.
	return C.int(fn())
}

// CreateThread creates a new SDL thread that runs the given function.
// The thread must eventually be waited on with Wait() or detached with Detach().
func CreateThread(fn ThreadFunction, name string) (*Thread, error) {
	id := registerCallback(fn)
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	ct := C._sdl_create_thread(cn, ptrFromID(id))
	if ct == nil {
		unregisterCallback(id)
		return nil, getError()
	}
	return &Thread{c: ct}, nil
}

// CreateThreadWithProperties creates a new SDL thread with properties.
func CreateThreadWithProperties(props PropertiesID) (*Thread, error) {
	ct := C._sdl_create_thread_with_properties(C.SDL_PropertiesID(props))
	if ct == nil {
		return nil, getError()
	}
	return &Thread{c: ct}, nil
}

// Note: SDL_CreateThreadRuntime and SDL_CreateThreadWithPropertiesRuntime are
// internal implementation functions. They are wrapped via the SDL_CreateThread
// and SDL_CreateThreadWithProperties macros used in the C helper functions above.

// Thread creation PROP_* constants are in sdl_props_constants.go.
