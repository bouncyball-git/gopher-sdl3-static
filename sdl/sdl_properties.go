package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>

extern void goEnumeratePropertiesCallback(void *userdata, SDL_PropertiesID props, char *name);
extern void goCleanupPropertyCallback(void *userdata, void *value);

static void cgoEnumeratePropertiesTrampoline(void *userdata, SDL_PropertiesID props, const char *name) {
	goEnumeratePropertiesCallback(userdata, props, (char *)name);
}

static void cgoCleanupPropertyTrampoline(void *userdata, void *value) {
	goCleanupPropertyCallback(userdata, value);
}
*/
import "C"

import "unsafe"

// PropertiesID is a unique identifier for a set of properties.
type PropertiesID uint32

// CreateProperties creates a new group of properties.
func CreateProperties() (PropertiesID, error) {
	id := C.SDL_CreateProperties()
	if id == 0 {
		return 0, getError()
	}
	return PropertiesID(id), nil
}

// DestroyProperties destroys a group of properties.
func DestroyProperties(props PropertiesID) {
	C.SDL_DestroyProperties(C.SDL_PropertiesID(props))
}

// SetStringProperty sets a string property.
func SetStringProperty(props PropertiesID, name, value string) error {
	cn := C.CString(name)
	cv := C.CString(value)
	defer C.free(unsafe.Pointer(cn))
	defer C.free(unsafe.Pointer(cv))
	if !C.SDL_SetStringProperty(C.SDL_PropertiesID(props), cn, cv) {
		return getError()
	}
	return nil
}

// GetStringProperty returns a string property.
func GetStringProperty(props PropertiesID, name, defaultValue string) string {
	cn := C.CString(name)
	cd := C.CString(defaultValue)
	defer C.free(unsafe.Pointer(cn))
	defer C.free(unsafe.Pointer(cd))
	return C.GoString(C.SDL_GetStringProperty(C.SDL_PropertiesID(props), cn, cd))
}

// SetNumberProperty sets a number property.
func SetNumberProperty(props PropertiesID, name string, value int64) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	if !C.SDL_SetNumberProperty(C.SDL_PropertiesID(props), cn, C.Sint64(value)) {
		return getError()
	}
	return nil
}

// GetNumberProperty returns a number property.
func GetNumberProperty(props PropertiesID, name string, defaultValue int64) int64 {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	return int64(C.SDL_GetNumberProperty(C.SDL_PropertiesID(props), cn, C.Sint64(defaultValue)))
}

// SetFloatProperty sets a float property.
func SetFloatProperty(props PropertiesID, name string, value float32) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	if !C.SDL_SetFloatProperty(C.SDL_PropertiesID(props), cn, C.float(value)) {
		return getError()
	}
	return nil
}

// GetFloatProperty returns a float property.
func GetFloatProperty(props PropertiesID, name string, defaultValue float32) float32 {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	return float32(C.SDL_GetFloatProperty(C.SDL_PropertiesID(props), cn, C.float(defaultValue)))
}

// SetBooleanProperty sets a boolean property.
func SetBooleanProperty(props PropertiesID, name string, value bool) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	if !C.SDL_SetBooleanProperty(C.SDL_PropertiesID(props), cn, C.bool(value)) {
		return getError()
	}
	return nil
}

// GetBooleanProperty returns a boolean property.
func GetBooleanProperty(props PropertiesID, name string, defaultValue bool) bool {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	return bool(C.SDL_GetBooleanProperty(C.SDL_PropertiesID(props), cn, C.bool(defaultValue)))
}

// HasProperty returns true if a property exists.
func HasProperty(props PropertiesID, name string) bool {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	return bool(C.SDL_HasProperty(C.SDL_PropertiesID(props), cn))
}

// ClearProperty clears a property.
func ClearProperty(props PropertiesID, name string) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	if !C.SDL_ClearProperty(C.SDL_PropertiesID(props), cn) {
		return getError()
	}
	return nil
}

// PropertyType represents the type of a property.
type PropertyType int

const (
	PROPERTY_TYPE_INVALID PropertyType = C.SDL_PROPERTY_TYPE_INVALID
	PROPERTY_TYPE_POINTER PropertyType = C.SDL_PROPERTY_TYPE_POINTER
	PROPERTY_TYPE_STRING  PropertyType = C.SDL_PROPERTY_TYPE_STRING
	PROPERTY_TYPE_NUMBER  PropertyType = C.SDL_PROPERTY_TYPE_NUMBER
	PROPERTY_TYPE_FLOAT   PropertyType = C.SDL_PROPERTY_TYPE_FLOAT
	PROPERTY_TYPE_BOOLEAN PropertyType = C.SDL_PROPERTY_TYPE_BOOLEAN
)

// GetGlobalProperties returns the global SDL properties.
func GetGlobalProperties() (PropertiesID, error) {
	id := C.SDL_GetGlobalProperties()
	if id == 0 {
		return 0, getError()
	}
	return PropertiesID(id), nil
}

// CopyProperties copies all properties from src to dst.
func CopyProperties(src, dst PropertiesID) error {
	if !C.SDL_CopyProperties(C.SDL_PropertiesID(src), C.SDL_PropertiesID(dst)) {
		return getError()
	}
	return nil
}

// LockProperties locks a group of properties for thread-safe access.
func LockProperties(props PropertiesID) error {
	if !C.SDL_LockProperties(C.SDL_PropertiesID(props)) {
		return getError()
	}
	return nil
}

// UnlockProperties unlocks a group of properties.
func UnlockProperties(props PropertiesID) {
	C.SDL_UnlockProperties(C.SDL_PropertiesID(props))
}

// SetPointerProperty sets a pointer property.
func SetPointerProperty(props PropertiesID, name string, value unsafe.Pointer) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	if !C.SDL_SetPointerProperty(C.SDL_PropertiesID(props), cn, value) {
		return getError()
	}
	return nil
}

// GetPointerProperty returns a pointer property.
func GetPointerProperty(props PropertiesID, name string, defaultValue unsafe.Pointer) unsafe.Pointer {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	return C.SDL_GetPointerProperty(C.SDL_PropertiesID(props), cn, defaultValue)
}

// GetPropertyType returns the type of a property.
func GetPropertyType(props PropertiesID, name string) PropertyType {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	return PropertyType(C.SDL_GetPropertyType(C.SDL_PropertiesID(props), cn))
}

// EnumeratePropertiesFunc is called for each property during enumeration.
type EnumeratePropertiesFunc func(props PropertiesID, name string)

// CleanupPropertyFunc is called when a pointer property is cleaned up.
type CleanupPropertyFunc func(value unsafe.Pointer)

//export goEnumeratePropertiesCallback
func goEnumeratePropertiesCallback(userdata unsafe.Pointer, props C.SDL_PropertiesID, name *C.char) {
	id := uintptr(userdata)
	fn := getCallback(id).(EnumeratePropertiesFunc)
	fn(PropertiesID(props), C.GoString(name))
}

//export goCleanupPropertyCallback
func goCleanupPropertyCallback(userdata unsafe.Pointer, value unsafe.Pointer) {
	id := uintptr(userdata)
	fn := getCallback(id).(CleanupPropertyFunc)
	fn(value)
	unregisterCallback(id)
}

// EnumerateProperties calls a function for each property in a set.
func EnumerateProperties(props PropertiesID, callback EnumeratePropertiesFunc) error {
	id := registerCallback(callback)
	defer unregisterCallback(id)
	if !C.SDL_EnumerateProperties(C.SDL_PropertiesID(props), C.SDL_EnumeratePropertiesCallback(C.cgoEnumeratePropertiesTrampoline), ptrFromID(id)) {
		return getError()
	}
	return nil
}

// SetPointerPropertyWithCleanup sets a pointer property with a cleanup callback.
func SetPointerPropertyWithCleanup(props PropertiesID, name string, value unsafe.Pointer, cleanup CleanupPropertyFunc) error {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	var ccb C.SDL_CleanupPropertyCallback
	var cud unsafe.Pointer
	if cleanup != nil {
		id := registerCallback(cleanup)
		ccb = C.SDL_CleanupPropertyCallback(C.cgoCleanupPropertyTrampoline)
		cud = ptrFromID(id)
	}
	if !C.SDL_SetPointerPropertyWithCleanup(C.SDL_PropertiesID(props), cn, value, ccb, cud) {
		return getError()
	}
	return nil
}
// SDL_PROP_* property name string constants.

// Property constants from SDL_init.h.
const (
	PROP_APP_METADATA_NAME_STRING = "SDL.app.metadata.name"
	PROP_APP_METADATA_VERSION_STRING = "SDL.app.metadata.version"
	PROP_APP_METADATA_IDENTIFIER_STRING = "SDL.app.metadata.identifier"
	PROP_APP_METADATA_CREATOR_STRING = "SDL.app.metadata.creator"
	PROP_APP_METADATA_COPYRIGHT_STRING = "SDL.app.metadata.copyright"
	PROP_APP_METADATA_URL_STRING = "SDL.app.metadata.url"
	PROP_APP_METADATA_TYPE_STRING = "SDL.app.metadata.type"
)

// Property constants from SDL_video.h.
const (
	PROP_GLOBAL_VIDEO_WAYLAND_WL_DISPLAY_POINTER = "SDL.video.wayland.wl_display"
	PROP_DISPLAY_HDR_ENABLED_BOOLEAN = "SDL.display.HDR_enabled"
	PROP_DISPLAY_KMSDRM_PANEL_ORIENTATION_NUMBER = "SDL.display.KMSDRM.panel_orientation"
	PROP_DISPLAY_WAYLAND_WL_OUTPUT_POINTER = "SDL.display.wayland.wl_output"
	PROP_DISPLAY_WINDOWS_HMONITOR_POINTER = "SDL.display.windows.hmonitor"
	PROP_WINDOW_CREATE_ALWAYS_ON_TOP_BOOLEAN = "SDL.window.create.always_on_top"
	PROP_WINDOW_CREATE_BORDERLESS_BOOLEAN = "SDL.window.create.borderless"
	PROP_WINDOW_CREATE_CONSTRAIN_POPUP_BOOLEAN = "SDL.window.create.constrain_popup"
	PROP_WINDOW_CREATE_FOCUSABLE_BOOLEAN = "SDL.window.create.focusable"
	PROP_WINDOW_CREATE_EXTERNAL_GRAPHICS_CONTEXT_BOOLEAN = "SDL.window.create.external_graphics_context"
	PROP_WINDOW_CREATE_FLAGS_NUMBER = "SDL.window.create.flags"
	PROP_WINDOW_CREATE_FULLSCREEN_BOOLEAN = "SDL.window.create.fullscreen"
	PROP_WINDOW_CREATE_HEIGHT_NUMBER = "SDL.window.create.height"
	PROP_WINDOW_CREATE_HIDDEN_BOOLEAN = "SDL.window.create.hidden"
	PROP_WINDOW_CREATE_HIGH_PIXEL_DENSITY_BOOLEAN = "SDL.window.create.high_pixel_density"
	PROP_WINDOW_CREATE_MAXIMIZED_BOOLEAN = "SDL.window.create.maximized"
	PROP_WINDOW_CREATE_MENU_BOOLEAN = "SDL.window.create.menu"
	PROP_WINDOW_CREATE_METAL_BOOLEAN = "SDL.window.create.metal"
	PROP_WINDOW_CREATE_MINIMIZED_BOOLEAN = "SDL.window.create.minimized"
	PROP_WINDOW_CREATE_MODAL_BOOLEAN = "SDL.window.create.modal"
	PROP_WINDOW_CREATE_MOUSE_GRABBED_BOOLEAN = "SDL.window.create.mouse_grabbed"
	PROP_WINDOW_CREATE_OPENGL_BOOLEAN = "SDL.window.create.opengl"
	PROP_WINDOW_CREATE_PARENT_POINTER = "SDL.window.create.parent"
	PROP_WINDOW_CREATE_RESIZABLE_BOOLEAN = "SDL.window.create.resizable"
	PROP_WINDOW_CREATE_TITLE_STRING = "SDL.window.create.title"
	PROP_WINDOW_CREATE_TRANSPARENT_BOOLEAN = "SDL.window.create.transparent"
	PROP_WINDOW_CREATE_TOOLTIP_BOOLEAN = "SDL.window.create.tooltip"
	PROP_WINDOW_CREATE_UTILITY_BOOLEAN = "SDL.window.create.utility"
	PROP_WINDOW_CREATE_VULKAN_BOOLEAN = "SDL.window.create.vulkan"
	PROP_WINDOW_CREATE_WIDTH_NUMBER = "SDL.window.create.width"
	PROP_WINDOW_CREATE_X_NUMBER = "SDL.window.create.x"
	PROP_WINDOW_CREATE_Y_NUMBER = "SDL.window.create.y"
	PROP_WINDOW_CREATE_COCOA_WINDOW_POINTER = "SDL.window.create.cocoa.window"
	PROP_WINDOW_CREATE_COCOA_VIEW_POINTER = "SDL.window.create.cocoa.view"
	PROP_WINDOW_CREATE_WINDOWSCENE_POINTER = "SDL.window.create.uikit.windowscene"
	PROP_WINDOW_CREATE_WAYLAND_SURFACE_ROLE_CUSTOM_BOOLEAN = "SDL.window.create.wayland.surface_role_custom"
	PROP_WINDOW_CREATE_WAYLAND_CREATE_EGL_WINDOW_BOOLEAN = "SDL.window.create.wayland.create_egl_window"
	PROP_WINDOW_CREATE_WAYLAND_WL_SURFACE_POINTER = "SDL.window.create.wayland.wl_surface"
	PROP_WINDOW_CREATE_WIN32_HWND_POINTER = "SDL.window.create.win32.hwnd"
	PROP_WINDOW_CREATE_WIN32_PIXEL_FORMAT_HWND_POINTER = "SDL.window.create.win32.pixel_format_hwnd"
	PROP_WINDOW_CREATE_X11_WINDOW_NUMBER = "SDL.window.create.x11.window"
	PROP_WINDOW_CREATE_EMSCRIPTEN_CANVAS_ID_STRING = "SDL.window.create.emscripten.canvas_id"
	PROP_WINDOW_CREATE_EMSCRIPTEN_KEYBOARD_ELEMENT_STRING = "SDL.window.create.emscripten.keyboard_element"
	PROP_WINDOW_SHAPE_POINTER = "SDL.window.shape"
	PROP_WINDOW_HDR_ENABLED_BOOLEAN = "SDL.window.HDR_enabled"
	PROP_WINDOW_SDR_WHITE_LEVEL_FLOAT = "SDL.window.SDR_white_level"
	PROP_WINDOW_HDR_HEADROOM_FLOAT = "SDL.window.HDR_headroom"
	PROP_WINDOW_ANDROID_WINDOW_POINTER = "SDL.window.android.window"
	PROP_WINDOW_ANDROID_SURFACE_POINTER = "SDL.window.android.surface"
	PROP_WINDOW_UIKIT_WINDOW_POINTER = "SDL.window.uikit.window"
	PROP_WINDOW_UIKIT_METAL_VIEW_TAG_NUMBER = "SDL.window.uikit.metal_view_tag"
	PROP_WINDOW_UIKIT_OPENGL_FRAMEBUFFER_NUMBER = "SDL.window.uikit.opengl.framebuffer"
	PROP_WINDOW_UIKIT_OPENGL_RENDERBUFFER_NUMBER = "SDL.window.uikit.opengl.renderbuffer"
	PROP_WINDOW_UIKIT_OPENGL_RESOLVE_FRAMEBUFFER_NUMBER = "SDL.window.uikit.opengl.resolve_framebuffer"
	PROP_WINDOW_KMSDRM_DEVICE_INDEX_NUMBER = "SDL.window.kmsdrm.dev_index"
	PROP_WINDOW_KMSDRM_DRM_FD_NUMBER = "SDL.window.kmsdrm.drm_fd"
	PROP_WINDOW_KMSDRM_GBM_DEVICE_POINTER = "SDL.window.kmsdrm.gbm_dev"
	PROP_WINDOW_COCOA_WINDOW_POINTER = "SDL.window.cocoa.window"
	PROP_WINDOW_COCOA_METAL_VIEW_TAG_NUMBER = "SDL.window.cocoa.metal_view_tag"
	PROP_WINDOW_OPENVR_OVERLAY_ID_NUMBER = "SDL.window.openvr.overlay_id"
	PROP_WINDOW_QNX_WINDOW_POINTER = "SDL.window.qnx.window"
	PROP_WINDOW_QNX_SURFACE_POINTER = "SDL.window.qnx.surface"
	PROP_WINDOW_VIVANTE_DISPLAY_POINTER = "SDL.window.vivante.display"
	PROP_WINDOW_VIVANTE_WINDOW_POINTER = "SDL.window.vivante.window"
	PROP_WINDOW_VIVANTE_SURFACE_POINTER = "SDL.window.vivante.surface"
	PROP_WINDOW_WIN32_HWND_POINTER = "SDL.window.win32.hwnd"
	PROP_WINDOW_WIN32_HDC_POINTER = "SDL.window.win32.hdc"
	PROP_WINDOW_WIN32_INSTANCE_POINTER = "SDL.window.win32.instance"
	PROP_WINDOW_WAYLAND_DISPLAY_POINTER = "SDL.window.wayland.display"
	PROP_WINDOW_WAYLAND_SURFACE_POINTER = "SDL.window.wayland.surface"
	PROP_WINDOW_WAYLAND_VIEWPORT_POINTER = "SDL.window.wayland.viewport"
	PROP_WINDOW_WAYLAND_EGL_WINDOW_POINTER = "SDL.window.wayland.egl_window"
	PROP_WINDOW_WAYLAND_XDG_SURFACE_POINTER = "SDL.window.wayland.xdg_surface"
	PROP_WINDOW_WAYLAND_XDG_TOPLEVEL_POINTER = "SDL.window.wayland.xdg_toplevel"
	PROP_WINDOW_WAYLAND_XDG_TOPLEVEL_EXPORT_HANDLE_STRING = "SDL.window.wayland.xdg_toplevel_export_handle"
	PROP_WINDOW_WAYLAND_XDG_POPUP_POINTER = "SDL.window.wayland.xdg_popup"
	PROP_WINDOW_WAYLAND_XDG_POSITIONER_POINTER = "SDL.window.wayland.xdg_positioner"
	PROP_WINDOW_X11_DISPLAY_POINTER = "SDL.window.x11.display"
	PROP_WINDOW_X11_SCREEN_NUMBER = "SDL.window.x11.screen"
	PROP_WINDOW_X11_WINDOW_NUMBER = "SDL.window.x11.window"
	PROP_WINDOW_EMSCRIPTEN_CANVAS_ID_STRING = "SDL.window.emscripten.canvas_id"
	PROP_WINDOW_EMSCRIPTEN_KEYBOARD_ELEMENT_STRING = "SDL.window.emscripten.keyboard_element"
)

// Property constants from SDL_render.h.
const (
	PROP_RENDERER_CREATE_NAME_STRING = "SDL.renderer.create.name"
	PROP_RENDERER_CREATE_WINDOW_POINTER = "SDL.renderer.create.window"
	PROP_RENDERER_CREATE_SURFACE_POINTER = "SDL.renderer.create.surface"
	PROP_RENDERER_CREATE_OUTPUT_COLORSPACE_NUMBER = "SDL.renderer.create.output_colorspace"
	PROP_RENDERER_CREATE_PRESENT_VSYNC_NUMBER = "SDL.renderer.create.present_vsync"
	PROP_RENDERER_CREATE_GPU_DEVICE_POINTER = "SDL.renderer.create.gpu.device"
	PROP_RENDERER_CREATE_GPU_SHADERS_SPIRV_BOOLEAN = "SDL.renderer.create.gpu.shaders_spirv"
	PROP_RENDERER_CREATE_GPU_SHADERS_DXIL_BOOLEAN = "SDL.renderer.create.gpu.shaders_dxil"
	PROP_RENDERER_CREATE_GPU_SHADERS_MSL_BOOLEAN = "SDL.renderer.create.gpu.shaders_msl"
	PROP_RENDERER_CREATE_VULKAN_INSTANCE_POINTER = "SDL.renderer.create.vulkan.instance"
	PROP_RENDERER_CREATE_VULKAN_SURFACE_NUMBER = "SDL.renderer.create.vulkan.surface"
	PROP_RENDERER_CREATE_VULKAN_PHYSICAL_DEVICE_POINTER = "SDL.renderer.create.vulkan.physical_device"
	PROP_RENDERER_CREATE_VULKAN_DEVICE_POINTER = "SDL.renderer.create.vulkan.device"
	PROP_RENDERER_CREATE_VULKAN_GRAPHICS_QUEUE_FAMILY_INDEX_NUMBER = "SDL.renderer.create.vulkan.graphics_queue_family_index"
	PROP_RENDERER_CREATE_VULKAN_PRESENT_QUEUE_FAMILY_INDEX_NUMBER = "SDL.renderer.create.vulkan.present_queue_family_index"
	PROP_RENDERER_NAME_STRING = "SDL.renderer.name"
	PROP_RENDERER_WINDOW_POINTER = "SDL.renderer.window"
	PROP_RENDERER_SURFACE_POINTER = "SDL.renderer.surface"
	PROP_RENDERER_VSYNC_NUMBER = "SDL.renderer.vsync"
	PROP_RENDERER_MAX_TEXTURE_SIZE_NUMBER = "SDL.renderer.max_texture_size"
	PROP_RENDERER_TEXTURE_FORMATS_POINTER = "SDL.renderer.texture_formats"
	PROP_RENDERER_TEXTURE_WRAPPING_BOOLEAN = "SDL.renderer.texture_wrapping"
	PROP_RENDERER_OUTPUT_COLORSPACE_NUMBER = "SDL.renderer.output_colorspace"
	PROP_RENDERER_HDR_ENABLED_BOOLEAN = "SDL.renderer.HDR_enabled"
	PROP_RENDERER_SDR_WHITE_POINT_FLOAT = "SDL.renderer.SDR_white_point"
	PROP_RENDERER_HDR_HEADROOM_FLOAT = "SDL.renderer.HDR_headroom"
	PROP_RENDERER_D3D9_DEVICE_POINTER = "SDL.renderer.d3d9.device"
	PROP_RENDERER_D3D11_DEVICE_POINTER = "SDL.renderer.d3d11.device"
	PROP_RENDERER_D3D11_SWAPCHAIN_POINTER = "SDL.renderer.d3d11.swap_chain"
	PROP_RENDERER_D3D12_DEVICE_POINTER = "SDL.renderer.d3d12.device"
	PROP_RENDERER_D3D12_SWAPCHAIN_POINTER = "SDL.renderer.d3d12.swap_chain"
	PROP_RENDERER_D3D12_COMMAND_QUEUE_POINTER = "SDL.renderer.d3d12.command_queue"
	PROP_RENDERER_VULKAN_INSTANCE_POINTER = "SDL.renderer.vulkan.instance"
	PROP_RENDERER_VULKAN_SURFACE_NUMBER = "SDL.renderer.vulkan.surface"
	PROP_RENDERER_VULKAN_PHYSICAL_DEVICE_POINTER = "SDL.renderer.vulkan.physical_device"
	PROP_RENDERER_VULKAN_DEVICE_POINTER = "SDL.renderer.vulkan.device"
	PROP_RENDERER_VULKAN_GRAPHICS_QUEUE_FAMILY_INDEX_NUMBER = "SDL.renderer.vulkan.graphics_queue_family_index"
	PROP_RENDERER_VULKAN_PRESENT_QUEUE_FAMILY_INDEX_NUMBER = "SDL.renderer.vulkan.present_queue_family_index"
	PROP_RENDERER_VULKAN_SWAPCHAIN_IMAGE_COUNT_NUMBER = "SDL.renderer.vulkan.swapchain_image_count"
	PROP_RENDERER_GPU_DEVICE_POINTER = "SDL.renderer.gpu.device"
	PROP_TEXTURE_CREATE_COLORSPACE_NUMBER = "SDL.texture.create.colorspace"
	PROP_TEXTURE_CREATE_FORMAT_NUMBER = "SDL.texture.create.format"
	PROP_TEXTURE_CREATE_ACCESS_NUMBER = "SDL.texture.create.access"
	PROP_TEXTURE_CREATE_WIDTH_NUMBER = "SDL.texture.create.width"
	PROP_TEXTURE_CREATE_HEIGHT_NUMBER = "SDL.texture.create.height"
	PROP_TEXTURE_CREATE_PALETTE_POINTER = "SDL.texture.create.palette"
	PROP_TEXTURE_CREATE_SDR_WHITE_POINT_FLOAT = "SDL.texture.create.SDR_white_point"
	PROP_TEXTURE_CREATE_HDR_HEADROOM_FLOAT = "SDL.texture.create.HDR_headroom"
	PROP_TEXTURE_CREATE_D3D11_TEXTURE_POINTER = "SDL.texture.create.d3d11.texture"
	PROP_TEXTURE_CREATE_D3D11_TEXTURE_U_POINTER = "SDL.texture.create.d3d11.texture_u"
	PROP_TEXTURE_CREATE_D3D11_TEXTURE_V_POINTER = "SDL.texture.create.d3d11.texture_v"
	PROP_TEXTURE_CREATE_D3D12_TEXTURE_POINTER = "SDL.texture.create.d3d12.texture"
	PROP_TEXTURE_CREATE_D3D12_TEXTURE_U_POINTER = "SDL.texture.create.d3d12.texture_u"
	PROP_TEXTURE_CREATE_D3D12_TEXTURE_V_POINTER = "SDL.texture.create.d3d12.texture_v"
	PROP_TEXTURE_CREATE_METAL_PIXELBUFFER_POINTER = "SDL.texture.create.metal.pixelbuffer"
	PROP_TEXTURE_CREATE_OPENGL_TEXTURE_NUMBER = "SDL.texture.create.opengl.texture"
	PROP_TEXTURE_CREATE_OPENGL_TEXTURE_UV_NUMBER = "SDL.texture.create.opengl.texture_uv"
	PROP_TEXTURE_CREATE_OPENGL_TEXTURE_U_NUMBER = "SDL.texture.create.opengl.texture_u"
	PROP_TEXTURE_CREATE_OPENGL_TEXTURE_V_NUMBER = "SDL.texture.create.opengl.texture_v"
	PROP_TEXTURE_CREATE_OPENGLES2_TEXTURE_NUMBER = "SDL.texture.create.opengles2.texture"
	PROP_TEXTURE_CREATE_OPENGLES2_TEXTURE_UV_NUMBER = "SDL.texture.create.opengles2.texture_uv"
	PROP_TEXTURE_CREATE_OPENGLES2_TEXTURE_U_NUMBER = "SDL.texture.create.opengles2.texture_u"
	PROP_TEXTURE_CREATE_OPENGLES2_TEXTURE_V_NUMBER = "SDL.texture.create.opengles2.texture_v"
	PROP_TEXTURE_CREATE_VULKAN_TEXTURE_NUMBER = "SDL.texture.create.vulkan.texture"
	PROP_TEXTURE_CREATE_VULKAN_LAYOUT_NUMBER = "SDL.texture.create.vulkan.layout"
	PROP_TEXTURE_CREATE_GPU_TEXTURE_POINTER = "SDL.texture.create.gpu.texture"
	PROP_TEXTURE_CREATE_GPU_TEXTURE_UV_POINTER = "SDL.texture.create.gpu.texture_uv"
	PROP_TEXTURE_CREATE_GPU_TEXTURE_U_POINTER = "SDL.texture.create.gpu.texture_u"
	PROP_TEXTURE_CREATE_GPU_TEXTURE_V_POINTER = "SDL.texture.create.gpu.texture_v"
	PROP_TEXTURE_COLORSPACE_NUMBER = "SDL.texture.colorspace"
	PROP_TEXTURE_FORMAT_NUMBER = "SDL.texture.format"
	PROP_TEXTURE_ACCESS_NUMBER = "SDL.texture.access"
	PROP_TEXTURE_WIDTH_NUMBER = "SDL.texture.width"
	PROP_TEXTURE_HEIGHT_NUMBER = "SDL.texture.height"
	PROP_TEXTURE_SDR_WHITE_POINT_FLOAT = "SDL.texture.SDR_white_point"
	PROP_TEXTURE_HDR_HEADROOM_FLOAT = "SDL.texture.HDR_headroom"
	PROP_TEXTURE_D3D11_TEXTURE_POINTER = "SDL.texture.d3d11.texture"
	PROP_TEXTURE_D3D11_TEXTURE_U_POINTER = "SDL.texture.d3d11.texture_u"
	PROP_TEXTURE_D3D11_TEXTURE_V_POINTER = "SDL.texture.d3d11.texture_v"
	PROP_TEXTURE_D3D12_TEXTURE_POINTER = "SDL.texture.d3d12.texture"
	PROP_TEXTURE_D3D12_TEXTURE_U_POINTER = "SDL.texture.d3d12.texture_u"
	PROP_TEXTURE_D3D12_TEXTURE_V_POINTER = "SDL.texture.d3d12.texture_v"
	PROP_TEXTURE_OPENGL_TEXTURE_NUMBER = "SDL.texture.opengl.texture"
	PROP_TEXTURE_OPENGL_TEXTURE_UV_NUMBER = "SDL.texture.opengl.texture_uv"
	PROP_TEXTURE_OPENGL_TEXTURE_U_NUMBER = "SDL.texture.opengl.texture_u"
	PROP_TEXTURE_OPENGL_TEXTURE_V_NUMBER = "SDL.texture.opengl.texture_v"
	PROP_TEXTURE_OPENGL_TEXTURE_TARGET_NUMBER = "SDL.texture.opengl.target"
	PROP_TEXTURE_OPENGL_TEX_W_FLOAT = "SDL.texture.opengl.tex_w"
	PROP_TEXTURE_OPENGL_TEX_H_FLOAT = "SDL.texture.opengl.tex_h"
	PROP_TEXTURE_OPENGLES2_TEXTURE_NUMBER = "SDL.texture.opengles2.texture"
	PROP_TEXTURE_OPENGLES2_TEXTURE_UV_NUMBER = "SDL.texture.opengles2.texture_uv"
	PROP_TEXTURE_OPENGLES2_TEXTURE_U_NUMBER = "SDL.texture.opengles2.texture_u"
	PROP_TEXTURE_OPENGLES2_TEXTURE_V_NUMBER = "SDL.texture.opengles2.texture_v"
	PROP_TEXTURE_OPENGLES2_TEXTURE_TARGET_NUMBER = "SDL.texture.opengles2.target"
	PROP_TEXTURE_VULKAN_TEXTURE_NUMBER = "SDL.texture.vulkan.texture"
	PROP_TEXTURE_GPU_TEXTURE_POINTER = "SDL.texture.gpu.texture"
	PROP_TEXTURE_GPU_TEXTURE_UV_POINTER = "SDL.texture.gpu.texture_uv"
	PROP_TEXTURE_GPU_TEXTURE_U_POINTER = "SDL.texture.gpu.texture_u"
	PROP_TEXTURE_GPU_TEXTURE_V_POINTER = "SDL.texture.gpu.texture_v"
)

// Property constants from SDL_gpu.h.
const (
	PROP_GPU_DEVICE_CREATE_DEBUGMODE_BOOLEAN = "SDL.gpu.device.create.debugmode"
	PROP_GPU_DEVICE_CREATE_PREFERLOWPOWER_BOOLEAN = "SDL.gpu.device.create.preferlowpower"
	PROP_GPU_DEVICE_CREATE_VERBOSE_BOOLEAN = "SDL.gpu.device.create.verbose"
	PROP_GPU_DEVICE_CREATE_NAME_STRING = "SDL.gpu.device.create.name"
	PROP_GPU_DEVICE_CREATE_FEATURE_CLIP_DISTANCE_BOOLEAN = "SDL.gpu.device.create.feature.clip_distance"
	PROP_GPU_DEVICE_CREATE_FEATURE_DEPTH_CLAMPING_BOOLEAN = "SDL.gpu.device.create.feature.depth_clamping"
	PROP_GPU_DEVICE_CREATE_FEATURE_INDIRECT_DRAW_FIRST_INSTANCE_BOOLEAN = "SDL.gpu.device.create.feature.indirect_draw_first_instance"
	PROP_GPU_DEVICE_CREATE_FEATURE_ANISOTROPY_BOOLEAN = "SDL.gpu.device.create.feature.anisotropy"
	PROP_GPU_DEVICE_CREATE_SHADERS_PRIVATE_BOOLEAN = "SDL.gpu.device.create.shaders.private"
	PROP_GPU_DEVICE_CREATE_SHADERS_SPIRV_BOOLEAN = "SDL.gpu.device.create.shaders.spirv"
	PROP_GPU_DEVICE_CREATE_SHADERS_DXBC_BOOLEAN = "SDL.gpu.device.create.shaders.dxbc"
	PROP_GPU_DEVICE_CREATE_SHADERS_DXIL_BOOLEAN = "SDL.gpu.device.create.shaders.dxil"
	PROP_GPU_DEVICE_CREATE_SHADERS_MSL_BOOLEAN = "SDL.gpu.device.create.shaders.msl"
	PROP_GPU_DEVICE_CREATE_SHADERS_METALLIB_BOOLEAN = "SDL.gpu.device.create.shaders.metallib"
	PROP_GPU_DEVICE_CREATE_D3D12_ALLOW_FEWER_RESOURCE_SLOTS_BOOLEAN = "SDL.gpu.device.create.d3d12.allowtier1resourcebinding"
	PROP_GPU_DEVICE_CREATE_D3D12_SEMANTIC_NAME_STRING = "SDL.gpu.device.create.d3d12.semantic"
	PROP_GPU_DEVICE_CREATE_D3D12_AGILITY_SDK_VERSION_NUMBER = "SDL.gpu.device.create.d3d12.agility_sdk_version"
	PROP_GPU_DEVICE_CREATE_D3D12_AGILITY_SDK_PATH_STRING = "SDL.gpu.device.create.d3d12.agility_sdk_path"
	PROP_GPU_DEVICE_CREATE_VULKAN_REQUIRE_HARDWARE_ACCELERATION_BOOLEAN = "SDL.gpu.device.create.vulkan.requirehardwareacceleration"
	PROP_GPU_DEVICE_CREATE_VULKAN_OPTIONS_POINTER = "SDL.gpu.device.create.vulkan.options"
	PROP_GPU_DEVICE_CREATE_METAL_ALLOW_MACFAMILY1_BOOLEAN = "SDL.gpu.device.create.metal.allowmacfamily1"
	PROP_GPU_DEVICE_CREATE_XR_ENABLE_BOOLEAN = "SDL.gpu.device.create.xr.enable"
	PROP_GPU_DEVICE_CREATE_XR_INSTANCE_POINTER = "SDL.gpu.device.create.xr.instance_out"
	PROP_GPU_DEVICE_CREATE_XR_SYSTEM_ID_POINTER = "SDL.gpu.device.create.xr.system_id_out"
	PROP_GPU_DEVICE_CREATE_XR_VERSION_NUMBER = "SDL.gpu.device.create.xr.version"
	PROP_GPU_DEVICE_CREATE_XR_FORM_FACTOR_NUMBER = "SDL.gpu.device.create.xr.form_factor"
	PROP_GPU_DEVICE_CREATE_XR_EXTENSION_COUNT_NUMBER = "SDL.gpu.device.create.xr.extensions.count"
	PROP_GPU_DEVICE_CREATE_XR_EXTENSION_NAMES_POINTER = "SDL.gpu.device.create.xr.extensions.names"
	PROP_GPU_DEVICE_CREATE_XR_LAYER_COUNT_NUMBER = "SDL.gpu.device.create.xr.layers.count"
	PROP_GPU_DEVICE_CREATE_XR_LAYER_NAMES_POINTER = "SDL.gpu.device.create.xr.layers.names"
	PROP_GPU_DEVICE_CREATE_XR_APPLICATION_NAME_STRING = "SDL.gpu.device.create.xr.application.name"
	PROP_GPU_DEVICE_CREATE_XR_APPLICATION_VERSION_NUMBER = "SDL.gpu.device.create.xr.application.version"
	PROP_GPU_DEVICE_CREATE_XR_ENGINE_NAME_STRING = "SDL.gpu.device.create.xr.engine.name"
	PROP_GPU_DEVICE_CREATE_XR_ENGINE_VERSION_NUMBER = "SDL.gpu.device.create.xr.engine.version"
	PROP_GPU_DEVICE_NAME_STRING = "SDL.gpu.device.name"
	PROP_GPU_DEVICE_DRIVER_NAME_STRING = "SDL.gpu.device.driver_name"
	PROP_GPU_DEVICE_DRIVER_VERSION_STRING = "SDL.gpu.device.driver_version"
	PROP_GPU_DEVICE_DRIVER_INFO_STRING = "SDL.gpu.device.driver_info"
	PROP_GPU_COMPUTEPIPELINE_CREATE_NAME_STRING = "SDL.gpu.computepipeline.create.name"
	PROP_GPU_GRAPHICSPIPELINE_CREATE_NAME_STRING = "SDL.gpu.graphicspipeline.create.name"
	PROP_GPU_SAMPLER_CREATE_NAME_STRING = "SDL.gpu.sampler.create.name"
	PROP_GPU_SHADER_CREATE_NAME_STRING = "SDL.gpu.shader.create.name"
	PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_R_FLOAT = "SDL.gpu.texture.create.d3d12.clear.r"
	PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_G_FLOAT = "SDL.gpu.texture.create.d3d12.clear.g"
	PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_B_FLOAT = "SDL.gpu.texture.create.d3d12.clear.b"
	PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_A_FLOAT = "SDL.gpu.texture.create.d3d12.clear.a"
	PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_DEPTH_FLOAT = "SDL.gpu.texture.create.d3d12.clear.depth"
	PROP_GPU_TEXTURE_CREATE_D3D12_CLEAR_STENCIL_NUMBER = "SDL.gpu.texture.create.d3d12.clear.stencil"
	PROP_GPU_TEXTURE_CREATE_NAME_STRING = "SDL.gpu.texture.create.name"
	PROP_GPU_BUFFER_CREATE_NAME_STRING = "SDL.gpu.buffer.create.name"
	PROP_GPU_TRANSFERBUFFER_CREATE_NAME_STRING = "SDL.gpu.transferbuffer.create.name"
)

// Property constants from SDL_audio.h.
const (
	PROP_AUDIOSTREAM_AUTO_CLEANUP_BOOLEAN = "SDL.audiostream.auto_cleanup"
)

// Property constants from SDL_iostream.h.
const (
	PROP_IOSTREAM_WINDOWS_HANDLE_POINTER = "SDL.iostream.windows.handle"
	PROP_IOSTREAM_STDIO_FILE_POINTER = "SDL.iostream.stdio.file"
	PROP_IOSTREAM_FILE_DESCRIPTOR_NUMBER = "SDL.iostream.file_descriptor"
	PROP_IOSTREAM_ANDROID_AASSET_POINTER = "SDL.iostream.android.aasset"
	PROP_IOSTREAM_MEMORY_POINTER = "SDL.iostream.memory.base"
	PROP_IOSTREAM_MEMORY_SIZE_NUMBER = "SDL.iostream.memory.size"
	PROP_IOSTREAM_MEMORY_FREE_FUNC_POINTER = "SDL.iostream.memory.free"
	PROP_IOSTREAM_DYNAMIC_MEMORY_POINTER = "SDL.iostream.dynamic.memory"
	PROP_IOSTREAM_DYNAMIC_CHUNKSIZE_NUMBER = "SDL.iostream.dynamic.chunksize"
)

// Property constants from SDL_surface.h.
const (
	PROP_SURFACE_SDR_WHITE_POINT_FLOAT = "SDL.surface.SDR_white_point"
	PROP_SURFACE_HDR_HEADROOM_FLOAT = "SDL.surface.HDR_headroom"
	PROP_SURFACE_TONEMAP_OPERATOR_STRING = "SDL.surface.tonemap"
	PROP_SURFACE_HOTSPOT_X_NUMBER = "SDL.surface.hotspot.x"
	PROP_SURFACE_HOTSPOT_Y_NUMBER = "SDL.surface.hotspot.y"
	PROP_SURFACE_ROTATION_FLOAT = "SDL.surface.rotation"
)

// Property constants from SDL_joystick.h.
const (
	PROP_JOYSTICK_CAP_MONO_LED_BOOLEAN = "SDL.joystick.cap.mono_led"
	PROP_JOYSTICK_CAP_RGB_LED_BOOLEAN = "SDL.joystick.cap.rgb_led"
	PROP_JOYSTICK_CAP_PLAYER_LED_BOOLEAN = "SDL.joystick.cap.player_led"
	PROP_JOYSTICK_CAP_RUMBLE_BOOLEAN = "SDL.joystick.cap.rumble"
	PROP_JOYSTICK_CAP_TRIGGER_RUMBLE_BOOLEAN = "SDL.joystick.cap.trigger_rumble"
)

// Property constants from SDL_gamepad.h (aliases to joystick properties).
const (
	PROP_GAMEPAD_CAP_MONO_LED_BOOLEAN       = PROP_JOYSTICK_CAP_MONO_LED_BOOLEAN
	PROP_GAMEPAD_CAP_RGB_LED_BOOLEAN        = PROP_JOYSTICK_CAP_RGB_LED_BOOLEAN
	PROP_GAMEPAD_CAP_PLAYER_LED_BOOLEAN     = PROP_JOYSTICK_CAP_PLAYER_LED_BOOLEAN
	PROP_GAMEPAD_CAP_RUMBLE_BOOLEAN         = PROP_JOYSTICK_CAP_RUMBLE_BOOLEAN
	PROP_GAMEPAD_CAP_TRIGGER_RUMBLE_BOOLEAN = PROP_JOYSTICK_CAP_TRIGGER_RUMBLE_BOOLEAN
)

// Property constants from SDL_keyboard.h.
const (
	PROP_TEXTINPUT_TYPE_NUMBER = "SDL.textinput.type"
	PROP_TEXTINPUT_CAPITALIZATION_NUMBER = "SDL.textinput.capitalization"
	PROP_TEXTINPUT_AUTOCORRECT_BOOLEAN = "SDL.textinput.autocorrect"
	PROP_TEXTINPUT_MULTILINE_BOOLEAN = "SDL.textinput.multiline"
	PROP_TEXTINPUT_ANDROID_INPUTTYPE_NUMBER = "SDL.textinput.android.inputtype"
)

// Property constants from SDL_hidapi.h.
const (
	PROP_HIDAPI_LIBUSB_DEVICE_HANDLE_POINTER = "SDL.hidapi.libusb.device.handle"
)

// Property constants from SDL_dialog.h.
const (
	PROP_FILE_DIALOG_FILTERS_POINTER = "SDL.filedialog.filters"
	PROP_FILE_DIALOG_NFILTERS_NUMBER = "SDL.filedialog.nfilters"
	PROP_FILE_DIALOG_WINDOW_POINTER = "SDL.filedialog.window"
	PROP_FILE_DIALOG_LOCATION_STRING = "SDL.filedialog.location"
	PROP_FILE_DIALOG_MANY_BOOLEAN = "SDL.filedialog.many"
	PROP_FILE_DIALOG_TITLE_STRING = "SDL.filedialog.title"
	PROP_FILE_DIALOG_ACCEPT_STRING = "SDL.filedialog.accept"
	PROP_FILE_DIALOG_CANCEL_STRING = "SDL.filedialog.cancel"
)

// Property constants from SDL_process.h.
const (
	PROP_PROCESS_CREATE_ARGS_POINTER = "SDL.process.create.args"
	PROP_PROCESS_CREATE_ENVIRONMENT_POINTER = "SDL.process.create.environment"
	PROP_PROCESS_CREATE_WORKING_DIRECTORY_STRING = "SDL.process.create.working_directory"
	PROP_PROCESS_CREATE_STDIN_NUMBER = "SDL.process.create.stdin_option"
	PROP_PROCESS_CREATE_STDIN_POINTER = "SDL.process.create.stdin_source"
	PROP_PROCESS_CREATE_STDOUT_NUMBER = "SDL.process.create.stdout_option"
	PROP_PROCESS_CREATE_STDOUT_POINTER = "SDL.process.create.stdout_source"
	PROP_PROCESS_CREATE_STDERR_NUMBER = "SDL.process.create.stderr_option"
	PROP_PROCESS_CREATE_STDERR_POINTER = "SDL.process.create.stderr_source"
	PROP_PROCESS_CREATE_STDERR_TO_STDOUT_BOOLEAN = "SDL.process.create.stderr_to_stdout"
	PROP_PROCESS_CREATE_BACKGROUND_BOOLEAN = "SDL.process.create.background"
	PROP_PROCESS_CREATE_CMDLINE_STRING = "SDL.process.create.cmdline"
	PROP_PROCESS_PID_NUMBER = "SDL.process.pid"
	PROP_PROCESS_STDIN_POINTER = "SDL.process.stdin"
	PROP_PROCESS_STDOUT_POINTER = "SDL.process.stdout"
	PROP_PROCESS_STDERR_POINTER = "SDL.process.stderr"
	PROP_PROCESS_BACKGROUND_BOOLEAN = "SDL.process.background"
)

// Property constants from SDL_properties.h.
const (
	PROP_NAME_STRING = "SDL.name"
)

// Property constants from SDL_thread.h.
const (
	PROP_THREAD_CREATE_ENTRY_FUNCTION_POINTER = "SDL.thread.create.entry_function"
	PROP_THREAD_CREATE_NAME_STRING = "SDL.thread.create.name"
	PROP_THREAD_CREATE_USERDATA_POINTER = "SDL.thread.create.userdata"
	PROP_THREAD_CREATE_STACKSIZE_NUMBER = "SDL.thread.create.stacksize"
)

// Property constants from SDL_tray.h.
const (
	PROP_TRAY_CREATE_ICON_POINTER = "SDL.tray.create.icon"
	PROP_TRAY_CREATE_TOOLTIP_STRING = "SDL.tray.create.tooltip"
	PROP_TRAY_CREATE_USERDATA_POINTER = "SDL.tray.create.userdata"
	PROP_TRAY_CREATE_LEFTCLICK_CALLBACK_POINTER = "SDL.tray.create.leftclick_callback"
	PROP_TRAY_CREATE_RIGHTCLICK_CALLBACK_POINTER = "SDL.tray.create.rightclick_callback"
	PROP_TRAY_CREATE_MIDDLECLICK_CALLBACK_POINTER = "SDL.tray.create.middleclick_callback"
)
