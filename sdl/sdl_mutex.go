package sdl

/*
#include <SDL3/SDL.h>
*/
import "C"

// Mutex is a means to serialize access to a resource between threads.
type Mutex struct {
	c *C.SDL_Mutex
}

// CreateMutex creates a new mutex.
func CreateMutex() (*Mutex, error) {
	c := C.SDL_CreateMutex()
	if c == nil {
		return nil, getError()
	}
	return &Mutex{c: c}, nil
}

// Lock locks the mutex.
func (m *Mutex) Lock() {
	C.SDL_LockMutex(m.c)
}

// TryLock tries to lock the mutex without blocking.
// Returns true on success, false if the mutex would block.
func (m *Mutex) TryLock() bool {
	return bool(C.SDL_TryLockMutex(m.c))
}

// Unlock unlocks the mutex.
func (m *Mutex) Unlock() {
	C.SDL_UnlockMutex(m.c)
}

// Destroy destroys the mutex.
func (m *Mutex) Destroy() {
	if m.c != nil {
		C.SDL_DestroyMutex(m.c)
		m.c = nil
	}
}

// RWLock is a mutex that allows read-only threads to run in parallel.
type RWLock struct {
	c *C.SDL_RWLock
}

// CreateRWLock creates a new read/write lock.
func CreateRWLock() (*RWLock, error) {
	c := C.SDL_CreateRWLock()
	if c == nil {
		return nil, getError()
	}
	return &RWLock{c: c}, nil
}

// LockForReading locks the read/write lock for read-only operations.
func (rw *RWLock) LockForReading() {
	C.SDL_LockRWLockForReading(rw.c)
}

// LockForWriting locks the read/write lock for write operations.
func (rw *RWLock) LockForWriting() {
	C.SDL_LockRWLockForWriting(rw.c)
}

// TryLockForReading tries to lock the read/write lock for reading without blocking.
// Returns true on success, false if the lock would block.
func (rw *RWLock) TryLockForReading() bool {
	return bool(C.SDL_TryLockRWLockForReading(rw.c))
}

// TryLockForWriting tries to lock the read/write lock for writing without blocking.
// Returns true on success, false if the lock would block.
func (rw *RWLock) TryLockForWriting() bool {
	return bool(C.SDL_TryLockRWLockForWriting(rw.c))
}

// Unlock unlocks the read/write lock.
func (rw *RWLock) Unlock() {
	C.SDL_UnlockRWLock(rw.c)
}

// Destroy destroys the read/write lock.
func (rw *RWLock) Destroy() {
	if rw.c != nil {
		C.SDL_DestroyRWLock(rw.c)
		rw.c = nil
	}
}

// Semaphore is a means to manage access to a resource, by count, between threads.
type Semaphore struct {
	c *C.SDL_Semaphore
}

// CreateSemaphore creates a new semaphore with the given initial value.
func CreateSemaphore(initialValue uint32) (*Semaphore, error) {
	c := C.SDL_CreateSemaphore(C.Uint32(initialValue))
	if c == nil {
		return nil, getError()
	}
	return &Semaphore{c: c}, nil
}

// Wait waits until the semaphore has a positive value and then decrements it.
func (s *Semaphore) Wait() {
	C.SDL_WaitSemaphore(s.c)
}

// TryWait checks if the semaphore has a positive value and decrements it if so.
// Returns true if the wait succeeds, false if the wait would block.
func (s *Semaphore) TryWait() bool {
	return bool(C.SDL_TryWaitSemaphore(s.c))
}

// WaitTimeout waits until the semaphore has a positive value or the timeout elapses.
// Returns true if the wait succeeds, false if the wait times out.
func (s *Semaphore) WaitTimeout(timeoutMS int32) bool {
	return bool(C.SDL_WaitSemaphoreTimeout(s.c, C.Sint32(timeoutMS)))
}

// Signal atomically increments the semaphore's value and wakes waiting threads.
func (s *Semaphore) Signal() {
	C.SDL_SignalSemaphore(s.c)
}

// Value returns the current value of the semaphore.
func (s *Semaphore) Value() uint32 {
	return uint32(C.SDL_GetSemaphoreValue(s.c))
}

// Destroy destroys the semaphore.
func (s *Semaphore) Destroy() {
	if s.c != nil {
		C.SDL_DestroySemaphore(s.c)
		s.c = nil
	}
}

// Condition is a means to block multiple threads until a condition is satisfied.
type Condition struct {
	c *C.SDL_Condition
}

// CreateCondition creates a new condition variable.
func CreateCondition() (*Condition, error) {
	c := C.SDL_CreateCondition()
	if c == nil {
		return nil, getError()
	}
	return &Condition{c: c}, nil
}

// Signal restarts one of the threads waiting on the condition variable.
func (cond *Condition) Signal() {
	C.SDL_SignalCondition(cond.c)
}

// Broadcast restarts all threads waiting on the condition variable.
func (cond *Condition) Broadcast() {
	C.SDL_BroadcastCondition(cond.c)
}

// Wait waits until the condition variable is signaled.
// The mutex must be locked before calling this function.
func (cond *Condition) Wait(mutex *Mutex) {
	C.SDL_WaitCondition(cond.c, mutex.c)
}

// WaitTimeout waits until the condition variable is signaled or the timeout elapses.
// The mutex must be locked before calling this function.
// Returns true if the condition is signaled, false if the wait times out.
func (cond *Condition) WaitTimeout(mutex *Mutex, timeoutMS int32) bool {
	return bool(C.SDL_WaitConditionTimeout(cond.c, mutex.c, C.Sint32(timeoutMS)))
}

// Destroy destroys the condition variable.
func (cond *Condition) Destroy() {
	if cond.c != nil {
		C.SDL_DestroyCondition(cond.c)
		cond.c = nil
	}
}

// InitState is used to manage one-time initialization and cleanup.
type InitState struct {
	c C.SDL_InitState
}

// ShouldInit returns true if initialization should be performed.
func (s *InitState) ShouldInit() bool {
	return bool(C.SDL_ShouldInit(&s.c))
}

// ShouldQuit returns true if cleanup should be performed.
func (s *InitState) ShouldQuit() bool {
	return bool(C.SDL_ShouldQuit(&s.c))
}

// SetInitialized marks the initialization state as completed.
func (s *InitState) SetInitialized(initialized bool) {
	C.SDL_SetInitialized(&s.c, C.bool(initialized))
}

// InitStatus represents the status of an SDL_InitState.
type InitStatus int

const (
	INIT_STATUS_UNINITIALIZED  InitStatus = C.SDL_INIT_STATUS_UNINITIALIZED
	INIT_STATUS_INITIALIZING   InitStatus = C.SDL_INIT_STATUS_INITIALIZING
	INIT_STATUS_INITIALIZED    InitStatus = C.SDL_INIT_STATUS_INITIALIZED
	INIT_STATUS_UNINITIALIZING InitStatus = C.SDL_INIT_STATUS_UNINITIALIZING
)

// Note: SDL_mutex.h thread safety annotation macros (SDL_ACQUIRE, SDL_GUARDED_BY,
// SDL_REQUIRES, etc.) are C compiler annotations with no Go equivalent.
