package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>

static void _sdl_log(const char *msg) {
	SDL_Log("%s", msg);
}

static void _sdl_log_category(int category, SDL_LogPriority priority, const char *msg) {
	SDL_LogMessage(category, priority, "%s", msg);
}

extern void goLogOutputFunction(void *userdata, int category, SDL_LogPriority priority, char *message);

static void cgoLogOutputTrampoline(void *userdata, int category, SDL_LogPriority priority, const char *message) {
	goLogOutputFunction(userdata, category, priority, (char *)message);
}

static SDL_LogOutputFunction _get_log_output_trampoline(void) {
	return cgoLogOutputTrampoline;
}
*/
import "C"

import "unsafe"

// LogPriority specifies the priority of a log message.
type LogPriority int

const (
	LOG_PRIORITY_INVALID  LogPriority = C.SDL_LOG_PRIORITY_INVALID
	LOG_PRIORITY_TRACE    LogPriority = C.SDL_LOG_PRIORITY_TRACE
	LOG_PRIORITY_VERBOSE  LogPriority = C.SDL_LOG_PRIORITY_VERBOSE
	LOG_PRIORITY_DEBUG    LogPriority = C.SDL_LOG_PRIORITY_DEBUG
	LOG_PRIORITY_INFO     LogPriority = C.SDL_LOG_PRIORITY_INFO
	LOG_PRIORITY_WARN     LogPriority = C.SDL_LOG_PRIORITY_WARN
	LOG_PRIORITY_ERROR    LogPriority = C.SDL_LOG_PRIORITY_ERROR
	LOG_PRIORITY_CRITICAL LogPriority = C.SDL_LOG_PRIORITY_CRITICAL
	LOG_PRIORITY_COUNT    LogPriority = C.SDL_LOG_PRIORITY_COUNT
)

// LogCategory represents a log category type.
type LogCategory int

const (
	LOG_CATEGORY_APPLICATION = int(C.SDL_LOG_CATEGORY_APPLICATION)
	LOG_CATEGORY_ERROR       = int(C.SDL_LOG_CATEGORY_ERROR)
	LOG_CATEGORY_ASSERT      = int(C.SDL_LOG_CATEGORY_ASSERT)
	LOG_CATEGORY_SYSTEM      = int(C.SDL_LOG_CATEGORY_SYSTEM)
	LOG_CATEGORY_AUDIO       = int(C.SDL_LOG_CATEGORY_AUDIO)
	LOG_CATEGORY_VIDEO       = int(C.SDL_LOG_CATEGORY_VIDEO)
	LOG_CATEGORY_RENDER      = int(C.SDL_LOG_CATEGORY_RENDER)
	LOG_CATEGORY_INPUT       = int(C.SDL_LOG_CATEGORY_INPUT)
	LOG_CATEGORY_TEST        = int(C.SDL_LOG_CATEGORY_TEST)
	LOG_CATEGORY_GPU         = int(C.SDL_LOG_CATEGORY_GPU)
	LOG_CATEGORY_RESERVED2  = int(C.SDL_LOG_CATEGORY_RESERVED2)
	LOG_CATEGORY_RESERVED3  = int(C.SDL_LOG_CATEGORY_RESERVED3)
	LOG_CATEGORY_RESERVED4  = int(C.SDL_LOG_CATEGORY_RESERVED4)
	LOG_CATEGORY_RESERVED5  = int(C.SDL_LOG_CATEGORY_RESERVED5)
	LOG_CATEGORY_RESERVED6  = int(C.SDL_LOG_CATEGORY_RESERVED6)
	LOG_CATEGORY_RESERVED7  = int(C.SDL_LOG_CATEGORY_RESERVED7)
	LOG_CATEGORY_RESERVED8  = int(C.SDL_LOG_CATEGORY_RESERVED8)
	LOG_CATEGORY_RESERVED9  = int(C.SDL_LOG_CATEGORY_RESERVED9)
	LOG_CATEGORY_RESERVED10 = int(C.SDL_LOG_CATEGORY_RESERVED10)
	LOG_CATEGORY_CUSTOM     = int(C.SDL_LOG_CATEGORY_CUSTOM)
)

// Log logs a message with SDL_LOG_CATEGORY_APPLICATION and SDL_LOG_PRIORITY_INFO.
func Log(msg string) {
	cm := C.CString(msg)
	defer C.free(unsafe.Pointer(cm))
	C._sdl_log(cm)
}

// LogMessage logs a message with the specified category and priority.
func LogMessage(category int, priority LogPriority, msg string) {
	cm := C.CString(msg)
	defer C.free(unsafe.Pointer(cm))
	C._sdl_log_category(C.int(category), C.SDL_LogPriority(priority), cm)
}

// Note: SDL_LogMessageV is not wrapped because it takes a C va_list argument
// which has no equivalent in Go. Use LogMessage instead.

// LogVerbose logs a message with SDL_LOG_PRIORITY_VERBOSE.
func LogVerbose(category int, msg string) {
	LogMessage(category, LOG_PRIORITY_VERBOSE, msg)
}

// LogDebug logs a message with SDL_LOG_PRIORITY_DEBUG.
func LogDebug(category int, msg string) {
	LogMessage(category, LOG_PRIORITY_DEBUG, msg)
}

// LogInfo logs a message with SDL_LOG_PRIORITY_INFO.
func LogInfo(category int, msg string) {
	LogMessage(category, LOG_PRIORITY_INFO, msg)
}

// LogWarn logs a message with SDL_LOG_PRIORITY_WARN.
func LogWarn(category int, msg string) {
	LogMessage(category, LOG_PRIORITY_WARN, msg)
}

// LogError logs a message with SDL_LOG_PRIORITY_ERROR.
func LogError(category int, msg string) {
	LogMessage(category, LOG_PRIORITY_ERROR, msg)
}

// LogCritical logs a message with SDL_LOG_PRIORITY_CRITICAL.
func LogCritical(category int, msg string) {
	LogMessage(category, LOG_PRIORITY_CRITICAL, msg)
}

// SetLogPriority sets the priority of a log category.
func SetLogPriority(category int, priority LogPriority) {
	C.SDL_SetLogPriority(C.int(category), C.SDL_LogPriority(priority))
}

// GetLogPriority returns the priority of a log category.
func GetLogPriority(category int) LogPriority {
	return LogPriority(C.SDL_GetLogPriority(C.int(category)))
}

// SetLogPriorities sets the priority of all log categories.
func SetLogPriorities(priority LogPriority) {
	C.SDL_SetLogPriorities(C.SDL_LogPriority(priority))
}

// ResetLogPriorities resets all log priorities to their default values.
func ResetLogPriorities() {
	C.SDL_ResetLogPriorities()
}

// SetLogPriorityPrefix sets the text prepended to log messages of a given priority.
func SetLogPriorityPrefix(priority LogPriority, prefix string) error {
	var cp *C.char
	if prefix != "" {
		cp = C.CString(prefix)
		defer C.free(unsafe.Pointer(cp))
	}
	if !C.SDL_SetLogPriorityPrefix(C.SDL_LogPriority(priority), cp) {
		return getError()
	}
	return nil
}

// LogTrace logs a message with SDL_LOG_PRIORITY_TRACE.
func LogTrace(category int, msg string) {
	LogMessage(category, LOG_PRIORITY_TRACE, msg)
}

// LogOutputFunc is a function that handles log output.
type LogOutputFunc func(category int, priority LogPriority, message string)

//export goLogOutputFunction
func goLogOutputFunction(userdata unsafe.Pointer, category C.int, priority C.SDL_LogPriority, message *C.char) {
	id := uintptr(userdata)
	fn := getCallback(id).(LogOutputFunc)
	fn(int(category), LogPriority(priority), C.GoString(message))
}

// SetLogOutputFunction replaces the default log output function.
func SetLogOutputFunction(callback LogOutputFunc) {
	id := registerCallback(callback)
	C.SDL_SetLogOutputFunction(C._get_log_output_trampoline(), ptrFromID(id))
}

// GetLogOutputFunction returns whether a custom log output function is currently set.
// The raw C function pointers are not directly accessible in Go.
func GetLogOutputFunction() bool {
	var cb C.SDL_LogOutputFunction
	var ud unsafe.Pointer
	C.SDL_GetLogOutputFunction(&cb, &ud)
	return cb != nil
}

// GetDefaultLogOutputFunction returns the default log output function as an opaque pointer.
func GetDefaultLogOutputFunction() unsafe.Pointer {
	return unsafe.Pointer(C.SDL_GetDefaultLogOutputFunction())
}
