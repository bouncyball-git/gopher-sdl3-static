package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// AsyncIO represents an asynchronous I/O operation handle.
type AsyncIO struct {
	c *C.SDL_AsyncIO
}

// AsyncIOQueue represents a queue of completed asynchronous I/O tasks.
type AsyncIOQueue struct {
	c *C.SDL_AsyncIOQueue
}

// AsyncIOTaskType represents the type of an asynchronous I/O task.
type AsyncIOTaskType int

// Async I/O task types.
const (
	ASYNCIO_TASK_READ  AsyncIOTaskType = C.SDL_ASYNCIO_TASK_READ
	ASYNCIO_TASK_WRITE AsyncIOTaskType = C.SDL_ASYNCIO_TASK_WRITE
	ASYNCIO_TASK_CLOSE AsyncIOTaskType = C.SDL_ASYNCIO_TASK_CLOSE
)

// AsyncIOResult represents the outcome of an asynchronous I/O task.
type AsyncIOResult int

// Async I/O result values.
const (
	ASYNCIO_COMPLETE AsyncIOResult = C.SDL_ASYNCIO_COMPLETE
	ASYNCIO_FAILURE  AsyncIOResult = C.SDL_ASYNCIO_FAILURE
	ASYNCIO_CANCELED AsyncIOResult = C.SDL_ASYNCIO_CANCELED
)

// AsyncIOOutcome contains information about a completed asynchronous I/O request.
type AsyncIOOutcome struct {
	Type             AsyncIOTaskType
	Result           AsyncIOResult
	Buffer           unsafe.Pointer
	Offset           uint64
	BytesRequested   uint64
	BytesTransferred uint64
	Userdata         unsafe.Pointer
}

func outcomeFromC(co *C.SDL_AsyncIOOutcome) AsyncIOOutcome {
	return AsyncIOOutcome{
		Type:             AsyncIOTaskType(co._type),
		Result:           AsyncIOResult(co.result),
		Buffer:           co.buffer,
		Offset:           uint64(co.offset),
		BytesRequested:   uint64(co.bytes_requested),
		BytesTransferred: uint64(co.bytes_transferred),
		Userdata:         co.userdata,
	}
}

// AsyncIOFromFile creates a new AsyncIO object for reading from and/or writing to a named file.
func AsyncIOFromFile(file, mode string) (*AsyncIO, error) {
	cf := C.CString(file)
	cm := C.CString(mode)
	defer C.free(unsafe.Pointer(cf))
	defer C.free(unsafe.Pointer(cm))
	aio := C.SDL_AsyncIOFromFile(cf, cm)
	if aio == nil {
		return nil, getError()
	}
	return &AsyncIO{c: aio}, nil
}

// GetSize returns the size of the data stream in the AsyncIO.
func (aio *AsyncIO) GetSize() (int64, error) {
	sz := C.SDL_GetAsyncIOSize(aio.c)
	if sz < 0 {
		return 0, getError()
	}
	return int64(sz), nil
}

// Read starts an async read operation.
// ptr must remain valid until the work is done.
func (aio *AsyncIO) Read(ptr unsafe.Pointer, offset, size uint64, queue *AsyncIOQueue, userdata unsafe.Pointer) error {
	if !C.SDL_ReadAsyncIO(aio.c, ptr, C.Uint64(offset), C.Uint64(size), queue.c, userdata) {
		return getError()
	}
	return nil
}

// Write starts an async write operation.
// ptr must remain valid until the work is done.
func (aio *AsyncIO) Write(ptr unsafe.Pointer, offset, size uint64, queue *AsyncIOQueue, userdata unsafe.Pointer) error {
	if !C.SDL_WriteAsyncIO(aio.c, ptr, C.Uint64(offset), C.Uint64(size), queue.c, userdata) {
		return getError()
	}
	return nil
}

// Close closes and frees any allocated resources for the async I/O object.
// If flush is true, data is synced to disk before the task completes.
func (aio *AsyncIO) Close(flush bool, queue *AsyncIOQueue, userdata unsafe.Pointer) error {
	if !C.SDL_CloseAsyncIO(aio.c, C.bool(flush), queue.c, userdata) {
		return getError()
	}
	aio.c = nil
	return nil
}

// CreateAsyncIOQueue creates a new task queue for tracking multiple I/O operations.
func CreateAsyncIOQueue() (*AsyncIOQueue, error) {
	q := C.SDL_CreateAsyncIOQueue()
	if q == nil {
		return nil, getError()
	}
	return &AsyncIOQueue{c: q}, nil
}

// Destroy destroys the async I/O task queue.
func (q *AsyncIOQueue) Destroy() {
	if q.c != nil {
		C.SDL_DestroyAsyncIOQueue(q.c)
		q.c = nil
	}
}

// GetResult queries the queue for a completed task without blocking.
// Returns the outcome and true if a task has completed, or false otherwise.
func (q *AsyncIOQueue) GetResult() (AsyncIOOutcome, bool) {
	var co C.SDL_AsyncIOOutcome
	if !C.SDL_GetAsyncIOResult(q.c, &co) {
		return AsyncIOOutcome{}, false
	}
	return outcomeFromC(&co), true
}

// WaitResult blocks until the queue has a completed task or the timeout expires.
// timeoutMS is the maximum time to wait in milliseconds, or -1 to wait indefinitely.
// Returns the outcome and true if a task completed, or false otherwise.
func (q *AsyncIOQueue) WaitResult(timeoutMS int32) (AsyncIOOutcome, bool) {
	var co C.SDL_AsyncIOOutcome
	if !C.SDL_WaitAsyncIOResult(q.c, &co, C.Sint32(timeoutMS)) {
		return AsyncIOOutcome{}, false
	}
	return outcomeFromC(&co), true
}

// Signal wakes up any threads that are blocking in WaitResult.
func (q *AsyncIOQueue) Signal() {
	C.SDL_SignalAsyncIOQueue(q.c)
}

// LoadFileAsync loads all the data from a file path asynchronously.
// The data will be allocated by SDL and must be freed with SDL_free via the outcome buffer.
func LoadFileAsync(file string, queue *AsyncIOQueue, userdata unsafe.Pointer) error {
	cf := C.CString(file)
	defer C.free(unsafe.Pointer(cf))
	if !C.SDL_LoadFileAsync(cf, queue.c, userdata) {
		return getError()
	}
	return nil
}
