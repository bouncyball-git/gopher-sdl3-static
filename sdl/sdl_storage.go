package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>

extern SDL_EnumerationResult goEnumerateDirectoryCallback(void *userdata, char *dirname, char *fname);

static SDL_EnumerationResult cgoEnumerateDirectoryTrampoline_storage(void *userdata, const char *dirname, const char *fname) {
	return goEnumerateDirectoryCallback(userdata, (char *)dirname, (char *)fname);
}
*/
import "C"

import "unsafe"

// Storage represents an abstract filesystem storage container.
type Storage struct {
	c *C.SDL_Storage
}

// OpenTitleStorage opens a read-only container for the application's filesystem.
// override is an optional path to override the backend's default title root (may be empty).
func OpenTitleStorage(override string, props PropertiesID) (*Storage, error) {
	var coverride *C.char
	if override != "" {
		coverride = C.CString(override)
		defer C.free(unsafe.Pointer(coverride))
	}
	s := C.SDL_OpenTitleStorage(coverride, C.SDL_PropertiesID(props))
	if s == nil {
		return nil, getError()
	}
	return &Storage{c: s}, nil
}

// OpenUserStorage opens a container for a user's unique read/write filesystem.
func OpenUserStorage(org, app string, props PropertiesID) (*Storage, error) {
	corg := C.CString(org)
	capp := C.CString(app)
	defer C.free(unsafe.Pointer(corg))
	defer C.free(unsafe.Pointer(capp))
	s := C.SDL_OpenUserStorage(corg, capp, C.SDL_PropertiesID(props))
	if s == nil {
		return nil, getError()
	}
	return &Storage{c: s}, nil
}

// OpenFileStorage opens a container for local filesystem storage.
// path is the base path prepended to all storage paths (may be empty for no base path).
func OpenFileStorage(path string) (*Storage, error) {
	var cpath *C.char
	if path != "" {
		cpath = C.CString(path)
		defer C.free(unsafe.Pointer(cpath))
	}
	s := C.SDL_OpenFileStorage(cpath)
	if s == nil {
		return nil, getError()
	}
	return &Storage{c: s}, nil
}

// Close closes and frees the storage container.
func (s *Storage) Close() error {
	if !C.SDL_CloseStorage(s.c) {
		return getError()
	}
	s.c = nil
	return nil
}

// Ready checks if the storage container is ready to use.
func (s *Storage) Ready() bool {
	return bool(C.SDL_StorageReady(s.c))
}

// GetFileSize queries the size of a file within the storage container.
func (s *Storage) GetFileSize(path string) (uint64, error) {
	cp := C.CString(path)
	defer C.free(unsafe.Pointer(cp))
	var length C.Uint64
	if !C.SDL_GetStorageFileSize(s.c, cp, &length) {
		return 0, getError()
	}
	return uint64(length), nil
}

// ReadFile reads a file from the storage container into the provided buffer.
// The length of data must match the file size exactly.
func (s *Storage) ReadFile(path string, data []byte) error {
	cp := C.CString(path)
	defer C.free(unsafe.Pointer(cp))
	if !C.SDL_ReadStorageFile(s.c, cp, unsafe.Pointer(&data[0]), C.Uint64(len(data))) {
		return getError()
	}
	return nil
}

// WriteFile writes data from the buffer into a file in the storage container.
func (s *Storage) WriteFile(path string, data []byte) error {
	cp := C.CString(path)
	defer C.free(unsafe.Pointer(cp))
	if !C.SDL_WriteStorageFile(s.c, cp, unsafe.Pointer(&data[0]), C.Uint64(len(data))) {
		return getError()
	}
	return nil
}

// CreateDirectory creates a directory in the writable storage container.
func (s *Storage) CreateDirectory(path string) error {
	cp := C.CString(path)
	defer C.free(unsafe.Pointer(cp))
	if !C.SDL_CreateStorageDirectory(s.c, cp) {
		return getError()
	}
	return nil
}

// RemovePath removes a file or empty directory in the storage container.
func (s *Storage) RemovePath(path string) error {
	cp := C.CString(path)
	defer C.free(unsafe.Pointer(cp))
	if !C.SDL_RemoveStoragePath(s.c, cp) {
		return getError()
	}
	return nil
}

// RenamePath renames a file or directory in the storage container.
func (s *Storage) RenamePath(oldpath, newpath string) error {
	co := C.CString(oldpath)
	cn := C.CString(newpath)
	defer C.free(unsafe.Pointer(co))
	defer C.free(unsafe.Pointer(cn))
	if !C.SDL_RenameStoragePath(s.c, co, cn) {
		return getError()
	}
	return nil
}

// CopyFile copies a file in the storage container.
func (s *Storage) CopyFile(oldpath, newpath string) error {
	co := C.CString(oldpath)
	cn := C.CString(newpath)
	defer C.free(unsafe.Pointer(co))
	defer C.free(unsafe.Pointer(cn))
	if !C.SDL_CopyStorageFile(s.c, co, cn) {
		return getError()
	}
	return nil
}

// GetPathInfo gets information about a filesystem path in the storage container.
func (s *Storage) GetPathInfo(path string) (PathInfo, error) {
	cp := C.CString(path)
	defer C.free(unsafe.Pointer(cp))
	var ci C.SDL_PathInfo
	if !C.SDL_GetStoragePathInfo(s.c, cp, &ci) {
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

// GetSpaceRemaining queries the remaining space in the storage container.
func (s *Storage) GetSpaceRemaining() uint64 {
	return uint64(C.SDL_GetStorageSpaceRemaining(s.c))
}

// GlobFlags and GLOB_CASEINSENSITIVE are defined in sdl_filesystem.go.

// GlobStorageDirectory enumerates a directory tree filtered by pattern.
func (s *Storage) GlobStorageDirectory(path, pattern string, flags GlobFlags) ([]string, error) {
	var cp *C.char
	if path != "" {
		cp = C.CString(path)
		defer C.free(unsafe.Pointer(cp))
	}
	var cpat *C.char
	if pattern != "" {
		cpat = C.CString(pattern)
		defer C.free(unsafe.Pointer(cpat))
	}
	var count C.int
	cresult := C.SDL_GlobStorageDirectory(s.c, cp, cpat, C.SDL_GlobFlags(flags), &count)
	if cresult == nil {
		return nil, getError()
	}
	defer C.SDL_free(unsafe.Pointer(cresult))
	n := int(count)
	result := make([]string, n)
	slice := unsafe.Slice((**C.char)(unsafe.Pointer(cresult)), n)
	for i, cs := range slice {
		result[i] = C.GoString(cs)
	}
	return result, nil
}

// EnumerateStorageDirectory enumerates a directory in storage, calling the callback for each entry.
// Uses the same EnumerateDirectoryFunc and trampoline from sdl_filesystem.go.
func (s *Storage) EnumerateDirectory(path string, callback EnumerateDirectoryFunc) error {
	cp := C.CString(path)
	defer C.free(unsafe.Pointer(cp))
	id := registerCallback(callback)
	defer unregisterCallback(id)
	if !C.SDL_EnumerateStorageDirectory(s.c, cp, C.SDL_EnumerateDirectoryCallback(C.cgoEnumerateDirectoryTrampoline_storage), ptrFromID(id)) {
		return getError()
	}
	return nil
}

// OpenStorage creates a custom storage container from a C SDL_StorageInterface.
// The iface parameter should point to an initialized SDL_StorageInterface struct.
func OpenStorage(iface unsafe.Pointer, userdata unsafe.Pointer) (*Storage, error) {
	cs := C.SDL_OpenStorage((*C.SDL_StorageInterface)(iface), userdata)
	if cs == nil {
		return nil, getError()
	}
	return &Storage{c: cs}, nil
}
