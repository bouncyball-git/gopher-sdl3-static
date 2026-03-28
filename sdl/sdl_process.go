package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// Process represents a running process.
type Process struct {
	c *C.SDL_Process
}

// CreateProcess creates a new process with the specified arguments.
func CreateProcess(args []string, pipeStdio bool) (*Process, error) {
	cargs := make([]*C.char, len(args)+1)
	for i, a := range args {
		cargs[i] = C.CString(a)
	}
	cargs[len(args)] = nil
	defer func() {
		for _, ca := range cargs {
			if ca != nil {
				C.free(unsafe.Pointer(ca))
			}
		}
	}()
	cp := C.SDL_CreateProcess(&cargs[0], C.bool(pipeStdio))
	if cp == nil {
		return nil, getError()
	}
	return &Process{c: cp}, nil
}

// Properties returns the properties of the process.
func (p *Process) Properties() PropertiesID {
	return PropertiesID(C.SDL_GetProcessProperties(p.c))
}

// GetInput returns the I/O stream for the process's stdin.
func (p *Process) GetInput() *IOStream {
	cs := C.SDL_GetProcessInput(p.c)
	if cs == nil {
		return nil
	}
	return &IOStream{c: cs}
}

// GetOutput returns the I/O stream for the process's stdout.
func (p *Process) GetOutput() *IOStream {
	cs := C.SDL_GetProcessOutput(p.c)
	if cs == nil {
		return nil
	}
	return &IOStream{c: cs}
}

// Kill sends a signal to terminate the process.
func (p *Process) Kill(force bool) error {
	if !C.SDL_KillProcess(p.c, C.bool(force)) {
		return getError()
	}
	return nil
}

// Wait waits for the process to finish.
func (p *Process) Wait(block bool) (exitcode int, finished bool, err error) {
	var ec C.int
	if !C.SDL_WaitProcess(p.c, C.bool(block), &ec) {
		return 0, false, getError()
	}
	return int(ec), true, nil
}

// ReadAll reads all output from the process and waits for it to finish.
func (p *Process) ReadAll() ([]byte, int, error) {
	var size C.size_t
	var exitcode C.int
	data := C.SDL_ReadProcess(p.c, &size, &exitcode)
	if data == nil {
		return nil, int(exitcode), getError()
	}
	defer C.SDL_free(data)
	result := make([]byte, int(size))
	copy(result, unsafe.Slice((*byte)(data), int(size)))
	return result, int(exitcode), nil
}

// Destroy destroys the process object.
func (p *Process) Destroy() {
	if p.c != nil {
		C.SDL_DestroyProcess(p.c)
		p.c = nil
	}
}

// CreateProcessWithProperties creates a new process with properties.
func CreateProcessWithProperties(props PropertiesID) (*Process, error) {
	cp := C.SDL_CreateProcessWithProperties(C.SDL_PropertiesID(props))
	if cp == nil {
		return nil, getError()
	}
	return &Process{c: cp}, nil
}

// ProcessIO represents the I/O mode for process stdio.
type ProcessIO int

const (
	PROCESS_STDIO_INHERITED ProcessIO = C.SDL_PROCESS_STDIO_INHERITED
	PROCESS_STDIO_NULL      ProcessIO = C.SDL_PROCESS_STDIO_NULL
	PROCESS_STDIO_APP       ProcessIO = C.SDL_PROCESS_STDIO_APP
	PROCESS_STDIO_REDIRECT  ProcessIO = C.SDL_PROCESS_STDIO_REDIRECT
)
