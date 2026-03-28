package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
#include <SDL3/SDL_vulkan.h>
*/
import "C"

import "unsafe"

// Vulkan_LoadLibrary dynamically loads the Vulkan loader library.
// Pass an empty string to use the default library.
func Vulkan_LoadLibrary(path string) error {
	var cp *C.char
	if path != "" {
		cp = C.CString(path)
		defer C.free(unsafe.Pointer(cp))
	}
	if !C.SDL_Vulkan_LoadLibrary(cp) {
		return getError()
	}
	return nil
}

// Vulkan_UnloadLibrary unloads the Vulkan library previously loaded by Vulkan_LoadLibrary.
func Vulkan_UnloadLibrary() {
	C.SDL_Vulkan_UnloadLibrary()
}

// Vulkan_GetInstanceExtensions returns the Vulkan instance extensions needed for vkCreateInstance.
func Vulkan_GetInstanceExtensions() []string {
	var count C.Uint32
	cexts := C.SDL_Vulkan_GetInstanceExtensions(&count)
	if cexts == nil {
		return nil
	}
	n := int(count)
	result := make([]string, n)
	slice := unsafe.Slice(cexts, n)
	for i, s := range slice {
		result[i] = C.GoString(s)
	}
	return result
}

// Vulkan_GetVkGetInstanceProcAddr returns the address of the vkGetInstanceProcAddr function.
func Vulkan_GetVkGetInstanceProcAddr() unsafe.Pointer {
	return unsafe.Pointer(C.SDL_Vulkan_GetVkGetInstanceProcAddr())
}

// VkInstance represents a Vulkan instance handle.
type VkInstance = unsafe.Pointer

// VkPhysicalDevice represents a Vulkan physical device handle.
type VkPhysicalDevice = unsafe.Pointer

// VkSurfaceKHR represents a Vulkan surface handle.
type VkSurfaceKHR = C.VkSurfaceKHR

// Vulkan_CreateSurface creates a Vulkan rendering surface for a window.
// Pass nil for allocator to use the default allocator.
func Vulkan_CreateSurface(window *Window, instance VkInstance, allocator unsafe.Pointer) (VkSurfaceKHR, error) {
	var surface C.VkSurfaceKHR
	if !C.SDL_Vulkan_CreateSurface(window.c, C.VkInstance(instance), (*C.struct_VkAllocationCallbacks)(allocator), &surface) {
		return VkSurfaceKHR(nil), getError()
	}
	return VkSurfaceKHR(surface), nil
}

// Vulkan_DestroySurface destroys a Vulkan surface.
func Vulkan_DestroySurface(instance VkInstance, surface VkSurfaceKHR, allocator unsafe.Pointer) {
	C.SDL_Vulkan_DestroySurface(C.VkInstance(instance), C.VkSurfaceKHR(surface), (*C.struct_VkAllocationCallbacks)(allocator))
}

// Vulkan_GetPresentationSupport queries whether a queue family of a physical device supports presentation.
func Vulkan_GetPresentationSupport(instance VkInstance, physicalDevice VkPhysicalDevice, queueFamilyIndex uint32) bool {
	return bool(C.SDL_Vulkan_GetPresentationSupport(C.VkInstance(instance), C.VkPhysicalDevice(physicalDevice), C.Uint32(queueFamilyIndex)))
}
