package sdl

/*
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// SpinLock is an atomic spinlock type.
type SpinLock int

// AtomicInt is a type representing an atomic integer value.
type AtomicInt struct {
	c C.SDL_AtomicInt
}

// AtomicU32 is a type representing an atomic unsigned 32-bit value.
type AtomicU32 struct {
	c C.SDL_AtomicU32
}

// TryLockSpinlock tries to lock a spin lock by setting it to a non-zero value.
// Returns true if the lock succeeded, false if the lock is already held.
func TryLockSpinlock(lock *SpinLock) bool {
	return bool(C.SDL_TryLockSpinlock((*C.SDL_SpinLock)(unsafe.Pointer(lock))))
}

// LockSpinlock locks a spin lock by setting it to a non-zero value.
func LockSpinlock(lock *SpinLock) {
	C.SDL_LockSpinlock((*C.SDL_SpinLock)(unsafe.Pointer(lock)))
}

// UnlockSpinlock unlocks a spin lock by setting it to 0.
func UnlockSpinlock(lock *SpinLock) {
	C.SDL_UnlockSpinlock((*C.SDL_SpinLock)(unsafe.Pointer(lock)))
}

// CompareAndSwap sets the atomic variable to a new value if it is currently the old value.
// Returns true if the atomic variable was set, false otherwise.
func (a *AtomicInt) CompareAndSwap(oldval, newval int) bool {
	return bool(C.SDL_CompareAndSwapAtomicInt(&a.c, C.int(oldval), C.int(newval)))
}

// Set sets the atomic variable to a value.
// Returns the previous value.
func (a *AtomicInt) Set(v int) int {
	return int(C.SDL_SetAtomicInt(&a.c, C.int(v)))
}

// Get returns the current value of the atomic variable.
func (a *AtomicInt) Get() int {
	return int(C.SDL_GetAtomicInt(&a.c))
}

// Add adds to the atomic variable.
// Returns the previous value.
func (a *AtomicInt) Add(v int) int {
	return int(C.SDL_AddAtomicInt(&a.c, C.int(v)))
}

// CompareAndSwap sets the atomic variable to a new value if it is currently the old value.
// Returns true if the atomic variable was set, false otherwise.
func (a *AtomicU32) CompareAndSwap(oldval, newval uint32) bool {
	return bool(C.SDL_CompareAndSwapAtomicU32(&a.c, C.Uint32(oldval), C.Uint32(newval)))
}

// Set sets the atomic variable to a value.
// Returns the previous value.
func (a *AtomicU32) Set(v uint32) uint32 {
	return uint32(C.SDL_SetAtomicU32(&a.c, C.Uint32(v)))
}

// Get returns the current value of the atomic variable.
func (a *AtomicU32) Get() uint32 {
	return uint32(C.SDL_GetAtomicU32(&a.c))
}

// Add adds to the atomic variable.
// Returns the previous value.
func (a *AtomicU32) Add(v int) uint32 {
	return uint32(C.SDL_AddAtomicU32(&a.c, C.int(v)))
}

// CompareAndSwapAtomicPointer sets a pointer to a new value if it is currently an old value.
// Returns true if the pointer was set, false otherwise.
func CompareAndSwapAtomicPointer(a *unsafe.Pointer, oldval, newval unsafe.Pointer) bool {
	return bool(C.SDL_CompareAndSwapAtomicPointer((*unsafe.Pointer)(unsafe.Pointer(a)), oldval, newval))
}

// SetAtomicPointer sets a pointer to a value atomically.
// Returns the previous value.
func SetAtomicPointer(a *unsafe.Pointer, v unsafe.Pointer) unsafe.Pointer {
	return C.SDL_SetAtomicPointer((*unsafe.Pointer)(unsafe.Pointer(a)), v)
}

// GetAtomicPointer gets the value of a pointer atomically.
func GetAtomicPointer(a *unsafe.Pointer) unsafe.Pointer {
	return C.SDL_GetAtomicPointer((*unsafe.Pointer)(unsafe.Pointer(a)))
}

// MemoryBarrierReleaseFunction issues a memory barrier release fence (function call version).
func MemoryBarrierReleaseFunction() {
	C.SDL_MemoryBarrierReleaseFunction()
}

// MemoryBarrierAcquireFunction issues a memory barrier acquire fence (function call version).
func MemoryBarrierAcquireFunction() {
	C.SDL_MemoryBarrierAcquireFunction()
}
