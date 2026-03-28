package sdl

/*
#include <SDL3/SDL.h>
*/
import "C"

// GetVersion returns the version of SDL that is linked against.
func GetVersion() int {
	return int(C.SDL_GetVersion())
}

// GetRevision returns the revision string of the SDL library.
func GetRevision() string {
	return C.GoString(C.SDL_GetRevision())
}

const (
	MAJOR_VERSION = C.SDL_MAJOR_VERSION
	MINOR_VERSION = C.SDL_MINOR_VERSION
	MICRO_VERSION = C.SDL_MICRO_VERSION
)

// VersionNum encodes a version number from major, minor, micro components.
func VersionNum(major, minor, micro int) int {
	return major*1000000 + minor*1000 + micro
}

// VersionMajor extracts the major version from an encoded version.
func VersionMajor(version int) int {
	return version / 1000000
}

// VersionMinor extracts the minor version from an encoded version.
func VersionMinor(version int) int {
	return (version / 1000) % 1000
}

// VersionMicro extracts the micro version from an encoded version.
func VersionMicro(version int) int {
	return version % 1000
}

// VersionAtLeast returns true if the SDL version is at least the given version.
func VersionAtLeast(x, y, z int) bool {
	return GetVersion() >= VersionNum(x, y, z)
}
