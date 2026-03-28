package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>

extern SDL_EnumerationResult goEnumerateDirectoryCallback(void *userdata, char *dirname, char *fname);

static SDL_EnumerationResult cgoEnumerateDirectoryTrampoline(void *userdata, const char *dirname, const char *fname) {
	return goEnumerateDirectoryCallback(userdata, (char *)dirname, (char *)fname);
}
*/
import "C"

import "unsafe"

// Folder represents a well-known folder type.
type Folder int

// Well-known user folder types.
const (
	FOLDER_HOME        Folder = C.SDL_FOLDER_HOME
	FOLDER_DESKTOP     Folder = C.SDL_FOLDER_DESKTOP
	FOLDER_DOCUMENTS   Folder = C.SDL_FOLDER_DOCUMENTS
	FOLDER_DOWNLOADS   Folder = C.SDL_FOLDER_DOWNLOADS
	FOLDER_MUSIC       Folder = C.SDL_FOLDER_MUSIC
	FOLDER_PICTURES    Folder = C.SDL_FOLDER_PICTURES
	FOLDER_PUBLICSHARE Folder = C.SDL_FOLDER_PUBLICSHARE
	FOLDER_SAVEDGAMES  Folder = C.SDL_FOLDER_SAVEDGAMES
	FOLDER_SCREENSHOTS Folder = C.SDL_FOLDER_SCREENSHOTS
	FOLDER_TEMPLATES   Folder = C.SDL_FOLDER_TEMPLATES
	FOLDER_VIDEOS      Folder = C.SDL_FOLDER_VIDEOS
)

// PathType represents the type of a filesystem path.
type PathType int

// Filesystem path types.
const (
	PATHTYPE_NONE      PathType = C.SDL_PATHTYPE_NONE
	PATHTYPE_FILE      PathType = C.SDL_PATHTYPE_FILE
	PATHTYPE_DIRECTORY PathType = C.SDL_PATHTYPE_DIRECTORY
	PATHTYPE_OTHER     PathType = C.SDL_PATHTYPE_OTHER
)

// PathInfo describes filesystem path information.
type PathInfo struct {
	Type       PathType
	Size       int64
	CreateTime int64
	ModifyTime int64
	AccessTime int64
}

// GetBasePath returns the directory where the application was run from.
func GetBasePath() string {
	return C.GoString(C.SDL_GetBasePath())
}

// GetPrefPath returns the user's preferred path for application data.
func GetPrefPath(org, app string) string {
	corg := C.CString(org)
	capp := C.CString(app)
	defer C.free(unsafe.Pointer(corg))
	defer C.free(unsafe.Pointer(capp))
	cs := C.SDL_GetPrefPath(corg, capp)
	if cs == nil {
		return ""
	}
	s := C.GoString(cs)
	C.SDL_free(unsafe.Pointer(cs))
	return s
}

// GetUserFolder returns the path to a well-known user folder.
func GetUserFolder(folder Folder) string {
	return C.GoString(C.SDL_GetUserFolder(C.SDL_Folder(folder)))
}

// CreateDirectory creates a directory and any missing parent directories.
func CreateDirectory(path string) error {
	cp := C.CString(path)
	defer C.free(unsafe.Pointer(cp))
	if !C.SDL_CreateDirectory(cp) {
		return getError()
	}
	return nil
}

// RemovePath removes a file or empty directory.
func RemovePath(path string) error {
	cp := C.CString(path)
	defer C.free(unsafe.Pointer(cp))
	if !C.SDL_RemovePath(cp) {
		return getError()
	}
	return nil
}

// RenamePath renames a file or directory.
func RenamePath(oldpath, newpath string) error {
	co := C.CString(oldpath)
	cn := C.CString(newpath)
	defer C.free(unsafe.Pointer(co))
	defer C.free(unsafe.Pointer(cn))
	if !C.SDL_RenamePath(co, cn) {
		return getError()
	}
	return nil
}

// CopyFile copies a file.
func CopyFile(oldpath, newpath string) error {
	co := C.CString(oldpath)
	cn := C.CString(newpath)
	defer C.free(unsafe.Pointer(co))
	defer C.free(unsafe.Pointer(cn))
	if !C.SDL_CopyFile(co, cn) {
		return getError()
	}
	return nil
}

// GetPathInfo returns information about a filesystem path.
func GetPathInfo(path string) (PathInfo, error) {
	cp := C.CString(path)
	defer C.free(unsafe.Pointer(cp))
	var ci C.SDL_PathInfo
	if !C.SDL_GetPathInfo(cp, &ci) {
		return PathInfo{}, getError()
	}
	return PathInfo{
		Type:       PathType(ci._type),
		Size:       int64(ci.size),
		CreateTime: int64(ci.create_time),
		ModifyTime: int64(ci.modify_time),
		AccessTime: int64(ci.access_time),
	}, nil
}

// GetCurrentDirectory returns the current working directory.
func GetCurrentDirectory() string {
	cs := C.SDL_GetCurrentDirectory()
	if cs == nil {
		return ""
	}
	s := C.GoString(cs)
	C.SDL_free(unsafe.Pointer(cs))
	return s
}

// GlobFlags represents flags for directory globbing.
type GlobFlags uint32

const (
	GLOB_CASEINSENSITIVE GlobFlags = 1
)

// GlobDirectory returns file paths matching a pattern in a directory.
func GlobDirectory(path, pattern string, flags GlobFlags) []string {
	cp := C.CString(path)
	cpt := C.CString(pattern)
	defer C.free(unsafe.Pointer(cp))
	defer C.free(unsafe.Pointer(cpt))
	var count C.int
	cresult := C.SDL_GlobDirectory(cp, cpt, C.SDL_GlobFlags(flags), &count)
	if cresult == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cresult))
	n := int(count)
	result := make([]string, n)
	slice := unsafe.Slice((**C.char)(unsafe.Pointer(cresult)), n)
	for i, s := range slice {
		result[i] = C.GoString(s)
	}
	return result
}

// EnumerationResult controls directory enumeration behavior.
type EnumerationResult int

const (
	ENUMERATION_CONTINUE EnumerationResult = C.SDL_ENUM_CONTINUE
	ENUMERATION_SUCCESS  EnumerationResult = C.SDL_ENUM_SUCCESS
	ENUMERATION_FAILURE  EnumerationResult = C.SDL_ENUM_FAILURE
)

// EnumerateDirectoryFunc is called for each entry during directory enumeration.
// Return ENUMERATION_CONTINUE to keep going, ENUMERATION_SUCCESS to stop successfully,
// or ENUMERATION_FAILURE to stop with an error.
type EnumerateDirectoryFunc func(dirname, fname string) EnumerationResult

//export goEnumerateDirectoryCallback
func goEnumerateDirectoryCallback(userdata unsafe.Pointer, dirname, fname *C.char) C.SDL_EnumerationResult {
	id := uintptr(userdata)
	fn := getCallback(id).(EnumerateDirectoryFunc)
	return C.SDL_EnumerationResult(fn(C.GoString(dirname), C.GoString(fname)))
}

// EnumerateDirectory enumerates a directory, calling the callback for each entry.
func EnumerateDirectory(path string, callback EnumerateDirectoryFunc) error {
	cp := C.CString(path)
	defer C.free(unsafe.Pointer(cp))
	id := registerCallback(callback)
	defer unregisterCallback(id)
	if !C.SDL_EnumerateDirectory(cp, C.SDL_EnumerateDirectoryCallback(C.cgoEnumerateDirectoryTrampoline), ptrFromID(id)) {
		return getError()
	}
	return nil
}
