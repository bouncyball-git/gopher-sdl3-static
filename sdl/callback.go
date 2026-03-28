package sdl

import "sync"

var (
	callbackMu     sync.Mutex
	callbackNextID uintptr = 1
	callbackMap            = make(map[uintptr]any)
)

func registerCallback(fn any) uintptr {
	callbackMu.Lock()
	defer callbackMu.Unlock()
	id := callbackNextID
	callbackNextID++
	callbackMap[id] = fn
	return id
}

func unregisterCallback(id uintptr) {
	callbackMu.Lock()
	defer callbackMu.Unlock()
	delete(callbackMap, id)
}

func getCallback(id uintptr) any {
	callbackMu.Lock()
	defer callbackMu.Unlock()
	return callbackMap[id]
}
