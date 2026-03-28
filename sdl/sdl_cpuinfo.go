package sdl

/*
#include <SDL3/SDL.h>
*/
import "C"

// GetNumLogicalCPUCores returns the number of logical CPU cores.
func GetNumLogicalCPUCores() int {
	return int(C.SDL_GetNumLogicalCPUCores())
}

// GetCPUCacheLineSize returns the L1 cache line size in bytes.
func GetCPUCacheLineSize() int {
	return int(C.SDL_GetCPUCacheLineSize())
}

// HasAltiVec returns true if the CPU has AltiVec features.
func HasAltiVec() bool { return bool(C.SDL_HasAltiVec()) }

// HasMMX returns true if the CPU has MMX features.
func HasMMX() bool { return bool(C.SDL_HasMMX()) }

// HasSSE returns true if the CPU has SSE features.
func HasSSE() bool { return bool(C.SDL_HasSSE()) }

// HasSSE2 returns true if the CPU has SSE2 features.
func HasSSE2() bool { return bool(C.SDL_HasSSE2()) }

// HasSSE3 returns true if the CPU has SSE3 features.
func HasSSE3() bool { return bool(C.SDL_HasSSE3()) }

// HasSSE41 returns true if the CPU has SSE4.1 features.
func HasSSE41() bool { return bool(C.SDL_HasSSE41()) }

// HasSSE42 returns true if the CPU has SSE4.2 features.
func HasSSE42() bool { return bool(C.SDL_HasSSE42()) }

// HasAVX returns true if the CPU has AVX features.
func HasAVX() bool { return bool(C.SDL_HasAVX()) }

// HasAVX2 returns true if the CPU has AVX2 features.
func HasAVX2() bool { return bool(C.SDL_HasAVX2()) }

// HasAVX512F returns true if the CPU has AVX-512F features.
func HasAVX512F() bool { return bool(C.SDL_HasAVX512F()) }

// HasARMSIMD returns true if the CPU has ARM SIMD features.
func HasARMSIMD() bool { return bool(C.SDL_HasARMSIMD()) }

// HasNEON returns true if the CPU has ARM NEON features.
func HasNEON() bool { return bool(C.SDL_HasNEON()) }

// HasLSX returns true if the CPU has LSX features.
func HasLSX() bool { return bool(C.SDL_HasLSX()) }

// HasLASX returns true if the CPU has LASX features.
func HasLASX() bool { return bool(C.SDL_HasLASX()) }

// GetSystemRAM returns the amount of RAM in MB.
func GetSystemRAM() int {
	return int(C.SDL_GetSystemRAM())
}

// GetSIMDAlignment returns the alignment needed for SIMD operations.
func GetSIMDAlignment() int {
	return int(C.SDL_GetSIMDAlignment())
}

// GetSystemPageSize returns the system's page size.
func GetSystemPageSize() int {
	return int(C.SDL_GetSystemPageSize())
}

// CACHELINE_SIZE is the assumed L1 cache line size for alignment purposes.
const CACHELINE_SIZE = C.SDL_CACHELINE_SIZE
