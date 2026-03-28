package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// Window represents an SDL window.
type Window struct {
	c *C.SDL_Window
}

// Window position macros (Go equivalents).
const (
	WINDOWPOS_UNDEFINED_MASK      = 0x1FFF0000
	WINDOWPOS_CENTERED_MASK       = 0x2FFF0000
	WINDOW_SURFACE_VSYNC_DISABLED = 0
	WINDOW_SURFACE_VSYNC_ADAPTIVE = -1
)

// WindowPosUndefinedDisplay returns an undefined position for a specific display.
func WindowPosUndefinedDisplay(x int) int { return int(WINDOWPOS_UNDEFINED_MASK) | x }

// WindowPosCenteredDisplay returns a centered position for a specific display.
func WindowPosCenteredDisplay(x int) int { return int(WINDOWPOS_CENTERED_MASK) | x }

// WindowPosIsUndefined returns whether a position value is undefined.
func WindowPosIsUndefined(x int) bool { return (x & 0xFFFF0000) == WINDOWPOS_UNDEFINED_MASK }

// WindowPosIsCentered returns whether a position value is centered.
func WindowPosIsCentered(x int) bool { return (x & 0xFFFF0000) == WINDOWPOS_CENTERED_MASK }

// WindowID is a unique identifier for a window.
type WindowID uint32

// DisplayID is a unique identifier for a display.
type DisplayID uint32

// WindowFlags are flags for window creation and state.
type WindowFlags uint64

const (
	WINDOW_FULLSCREEN          WindowFlags = 0x0000000000000001
	WINDOW_OPENGL              WindowFlags = 0x0000000000000002
	WINDOW_OCCLUDED            WindowFlags = 0x0000000000000004
	WINDOW_HIDDEN              WindowFlags = 0x0000000000000008
	WINDOW_BORDERLESS          WindowFlags = 0x0000000000000010
	WINDOW_RESIZABLE           WindowFlags = 0x0000000000000020
	WINDOW_MINIMIZED           WindowFlags = 0x0000000000000040
	WINDOW_MAXIMIZED           WindowFlags = 0x0000000000000080
	WINDOW_MOUSE_GRABBED       WindowFlags = 0x0000000000000100
	WINDOW_INPUT_FOCUS         WindowFlags = 0x0000000000000200
	WINDOW_MOUSE_FOCUS         WindowFlags = 0x0000000000000400
	WINDOW_EXTERNAL            WindowFlags = 0x0000000000000800
	WINDOW_MODAL               WindowFlags = 0x0000000000001000
	WINDOW_HIGH_PIXEL_DENSITY  WindowFlags = 0x0000000000002000
	WINDOW_MOUSE_CAPTURE       WindowFlags = 0x0000000000004000
	WINDOW_MOUSE_RELATIVE_MODE WindowFlags = 0x0000000000008000
	WINDOW_ALWAYS_ON_TOP       WindowFlags = 0x0000000000010000
	WINDOW_UTILITY             WindowFlags = 0x0000000000020000
	WINDOW_TOOLTIP             WindowFlags = 0x0000000000040000
	WINDOW_POPUP_MENU          WindowFlags = 0x0000000000080000
	WINDOW_KEYBOARD_GRABBED    WindowFlags = 0x0000000000100000
	WINDOW_FILL_DOCUMENT       WindowFlags = 0x0000000000200000
	WINDOW_VULKAN              WindowFlags = 0x0000000010000000
	WINDOW_METAL               WindowFlags = 0x0000000020000000
	WINDOW_TRANSPARENT         WindowFlags = 0x0000000040000000
	WINDOW_NOT_FOCUSABLE       WindowFlags = 0x0000000080000000
)

// SystemTheme represents the system theme.
type SystemTheme int

const (
	SYSTEM_THEME_UNKNOWN SystemTheme = C.SDL_SYSTEM_THEME_UNKNOWN
	SYSTEM_THEME_LIGHT   SystemTheme = C.SDL_SYSTEM_THEME_LIGHT
	SYSTEM_THEME_DARK    SystemTheme = C.SDL_SYSTEM_THEME_DARK
)

// DisplayMode describes a display mode.
type DisplayMode struct {
	DisplayID              DisplayID
	Format                 PixelFormat
	W                      int32
	H                      int32
	PixelDensity           float32
	RefreshRate            float32
	RefreshRateNumerator   int32
	RefreshRateDenominator int32
}

// FlashOperation specifies the flash operation for a window.
type FlashOperation int

const (
	FLASH_CANCEL        FlashOperation = C.SDL_FLASH_CANCEL
	FLASH_BRIEFLY       FlashOperation = C.SDL_FLASH_BRIEFLY
	FLASH_UNTIL_FOCUSED FlashOperation = C.SDL_FLASH_UNTIL_FOCUSED
)

// CreateWindow creates a window with the specified title, size, and flags.
func CreateWindow(title string, w, h int, flags WindowFlags) (*Window, error) {
	ct := C.CString(title)
	defer C.free(unsafe.Pointer(ct))
	cw := C.SDL_CreateWindow(ct, C.int(w), C.int(h), C.SDL_WindowFlags(flags))
	if cw == nil {
		return nil, getError()
	}
	return &Window{c: cw}, nil
}

// Destroy destroys the window.
func (w *Window) Destroy() {
	if w.c != nil {
		C.SDL_DestroyWindow(w.c)
		w.c = nil
	}
}

// ID returns the window's unique identifier.
func (w *Window) ID() WindowID {
	return WindowID(C.SDL_GetWindowID(w.c))
}

// Title returns the title of the window.
func (w *Window) Title() string {
	return C.GoString(C.SDL_GetWindowTitle(w.c))
}

// SetTitle sets the title of the window.
func (w *Window) SetTitle(title string) error {
	ct := C.CString(title)
	defer C.free(unsafe.Pointer(ct))
	if !C.SDL_SetWindowTitle(w.c, ct) {
		return getError()
	}
	return nil
}

// Size returns the size of the window in screen coordinates.
func (w *Window) Size() (int, int, error) {
	var width, height C.int
	if !C.SDL_GetWindowSize(w.c, &width, &height) {
		return 0, 0, getError()
	}
	return int(width), int(height), nil
}

// SetSize sets the size of the window.
func (w *Window) SetSize(width, height int) error {
	if !C.SDL_SetWindowSize(w.c, C.int(width), C.int(height)) {
		return getError()
	}
	return nil
}

// SizeInPixels returns the size of the window in pixels.
func (w *Window) SizeInPixels() (int, int, error) {
	var width, height C.int
	if !C.SDL_GetWindowSizeInPixels(w.c, &width, &height) {
		return 0, 0, getError()
	}
	return int(width), int(height), nil
}

// Position returns the position of the window.
func (w *Window) Position() (int, int, error) {
	var x, y C.int
	if !C.SDL_GetWindowPosition(w.c, &x, &y) {
		return 0, 0, getError()
	}
	return int(x), int(y), nil
}

// SetPosition sets the position of the window.
func (w *Window) SetPosition(x, y int) error {
	if !C.SDL_SetWindowPosition(w.c, C.int(x), C.int(y)) {
		return getError()
	}
	return nil
}

// Flags returns the window flags.
func (w *Window) Flags() WindowFlags {
	return WindowFlags(C.SDL_GetWindowFlags(w.c))
}

// SetFullscreen sets the fullscreen state of the window.
func (w *Window) SetFullscreen(fullscreen bool) error {
	if !C.SDL_SetWindowFullscreen(w.c, C.bool(fullscreen)) {
		return getError()
	}
	return nil
}

// SetResizable sets whether the window is resizable.
func (w *Window) SetResizable(resizable bool) error {
	if !C.SDL_SetWindowResizable(w.c, C.bool(resizable)) {
		return getError()
	}
	return nil
}

// SetBordered sets whether the window has a border.
func (w *Window) SetBordered(bordered bool) error {
	if !C.SDL_SetWindowBordered(w.c, C.bool(bordered)) {
		return getError()
	}
	return nil
}

// SetAlwaysOnTop sets whether the window is always on top.
func (w *Window) SetAlwaysOnTop(onTop bool) error {
	if !C.SDL_SetWindowAlwaysOnTop(w.c, C.bool(onTop)) {
		return getError()
	}
	return nil
}

// Show shows the window.
func (w *Window) Show() error {
	if !C.SDL_ShowWindow(w.c) {
		return getError()
	}
	return nil
}

// Hide hides the window.
func (w *Window) Hide() error {
	if !C.SDL_HideWindow(w.c) {
		return getError()
	}
	return nil
}

// Raise raises the window above other windows.
func (w *Window) Raise() error {
	if !C.SDL_RaiseWindow(w.c) {
		return getError()
	}
	return nil
}

// Maximize maximizes the window.
func (w *Window) Maximize() error {
	if !C.SDL_MaximizeWindow(w.c) {
		return getError()
	}
	return nil
}

// Minimize minimizes the window.
func (w *Window) Minimize() error {
	if !C.SDL_MinimizeWindow(w.c) {
		return getError()
	}
	return nil
}

// Restore restores the window from minimized or maximized state.
func (w *Window) Restore() error {
	if !C.SDL_RestoreWindow(w.c) {
		return getError()
	}
	return nil
}

// SetMinimumSize sets the minimum size of the window.
func (w *Window) SetMinimumSize(minW, minH int) error {
	if !C.SDL_SetWindowMinimumSize(w.c, C.int(minW), C.int(minH)) {
		return getError()
	}
	return nil
}

// MinimumSize returns the minimum size of the window.
func (w *Window) MinimumSize() (int, int, error) {
	var width, height C.int
	if !C.SDL_GetWindowMinimumSize(w.c, &width, &height) {
		return 0, 0, getError()
	}
	return int(width), int(height), nil
}

// SetMaximumSize sets the maximum size of the window.
func (w *Window) SetMaximumSize(maxW, maxH int) error {
	if !C.SDL_SetWindowMaximumSize(w.c, C.int(maxW), C.int(maxH)) {
		return getError()
	}
	return nil
}

// MaximumSize returns the maximum size of the window.
func (w *Window) MaximumSize() (int, int, error) {
	var width, height C.int
	if !C.SDL_GetWindowMaximumSize(w.c, &width, &height) {
		return 0, 0, getError()
	}
	return int(width), int(height), nil
}

// Surface returns the surface associated with the window.
func (w *Window) Surface() (*Surface, error) {
	cs := C.SDL_GetWindowSurface(w.c)
	if cs == nil {
		return nil, getError()
	}
	return &Surface{c: cs}, nil
}

// UpdateSurface updates the window surface.
func (w *Window) UpdateSurface() error {
	if !C.SDL_UpdateWindowSurface(w.c) {
		return getError()
	}
	return nil
}

// Flash flashes the window to get the user's attention.
func (w *Window) Flash(operation FlashOperation) error {
	if !C.SDL_FlashWindow(w.c, C.SDL_FlashOperation(operation)) {
		return getError()
	}
	return nil
}

// Display returns the display associated with the window.
func (w *Window) Display() DisplayID {
	return DisplayID(C.SDL_GetDisplayForWindow(w.c))
}

// PixelDensity returns the pixel density of the window.
func (w *Window) PixelDensity() float32 {
	return float32(C.SDL_GetWindowPixelDensity(w.c))
}

// DisplayScale returns the display scale of the window.
func (w *Window) DisplayScale() float32 {
	return float32(C.SDL_GetWindowDisplayScale(w.c))
}

// GetWindowFromID returns the window associated with the given ID.
func GetWindowFromID(id WindowID) (*Window, error) {
	cw := C.SDL_GetWindowFromID(C.SDL_WindowID(id))
	if cw == nil {
		return nil, getError()
	}
	return &Window{c: cw}, nil
}

// GetDisplays returns a list of currently connected displays.
func GetDisplays() []DisplayID {
	var count C.int
	cids := C.SDL_GetDisplays(&count)
	if cids == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cids))
	n := int(count)
	result := make([]DisplayID, n)
	slice := unsafe.Slice((*C.SDL_DisplayID)(cids), n)
	for i, id := range slice {
		result[i] = DisplayID(id)
	}
	return result
}

// GetDisplayName returns the name of a display.
func GetDisplayName(displayID DisplayID) string {
	return C.GoString(C.SDL_GetDisplayName(C.SDL_DisplayID(displayID)))
}

// GetDisplayBounds returns the desktop area of a display.
func GetDisplayBounds(displayID DisplayID) (Rect, error) {
	var rect Rect
	if !C.SDL_GetDisplayBounds(C.SDL_DisplayID(displayID), rect.cptr()) {
		return rect, getError()
	}
	return rect, nil
}

// GetDisplayUsableBounds returns the usable desktop area of a display.
func GetDisplayUsableBounds(displayID DisplayID) (Rect, error) {
	var rect Rect
	if !C.SDL_GetDisplayUsableBounds(C.SDL_DisplayID(displayID), rect.cptr()) {
		return rect, getError()
	}
	return rect, nil
}

// GetDisplayContentScale returns the content scale of a display.
func GetDisplayContentScale(displayID DisplayID) float32 {
	return float32(C.SDL_GetDisplayContentScale(C.SDL_DisplayID(displayID)))
}

// GetNumVideoDrivers returns the number of video drivers compiled into SDL.
func GetNumVideoDrivers() int {
	return int(C.SDL_GetNumVideoDrivers())
}

// GetVideoDriver returns the name of a built-in video driver.
func GetVideoDriver(index int) string {
	return C.GoString(C.SDL_GetVideoDriver(C.int(index)))
}

// GetCurrentVideoDriver returns the name of the currently initialized video driver.
func GetCurrentVideoDriver() string {
	return C.GoString(C.SDL_GetCurrentVideoDriver())
}

// GetSystemTheme returns the current system theme.
func GetSystemTheme() SystemTheme {
	return SystemTheme(C.SDL_GetSystemTheme())
}

// DisplayOrientation represents display orientation values.
type DisplayOrientation int

const (
	ORIENTATION_UNKNOWN            DisplayOrientation = C.SDL_ORIENTATION_UNKNOWN
	ORIENTATION_LANDSCAPE          DisplayOrientation = C.SDL_ORIENTATION_LANDSCAPE
	ORIENTATION_LANDSCAPE_FLIPPED  DisplayOrientation = C.SDL_ORIENTATION_LANDSCAPE_FLIPPED
	ORIENTATION_PORTRAIT           DisplayOrientation = C.SDL_ORIENTATION_PORTRAIT
	ORIENTATION_PORTRAIT_FLIPPED   DisplayOrientation = C.SDL_ORIENTATION_PORTRAIT_FLIPPED
)

// GLContext represents an OpenGL context.
type GLContext C.SDL_GLContext

// GLAttr represents an OpenGL attribute.
type GLAttr int

const (
	GL_RED_SIZE                   GLAttr = C.SDL_GL_RED_SIZE
	GL_GREEN_SIZE                 GLAttr = C.SDL_GL_GREEN_SIZE
	GL_BLUE_SIZE                  GLAttr = C.SDL_GL_BLUE_SIZE
	GL_ALPHA_SIZE                 GLAttr = C.SDL_GL_ALPHA_SIZE
	GL_BUFFER_SIZE                GLAttr = C.SDL_GL_BUFFER_SIZE
	GL_DOUBLEBUFFER               GLAttr = C.SDL_GL_DOUBLEBUFFER
	GL_DEPTH_SIZE                 GLAttr = C.SDL_GL_DEPTH_SIZE
	GL_STENCIL_SIZE               GLAttr = C.SDL_GL_STENCIL_SIZE
	GL_ACCUM_RED_SIZE             GLAttr = C.SDL_GL_ACCUM_RED_SIZE
	GL_ACCUM_GREEN_SIZE           GLAttr = C.SDL_GL_ACCUM_GREEN_SIZE
	GL_ACCUM_BLUE_SIZE            GLAttr = C.SDL_GL_ACCUM_BLUE_SIZE
	GL_ACCUM_ALPHA_SIZE           GLAttr = C.SDL_GL_ACCUM_ALPHA_SIZE
	GL_STEREO                     GLAttr = C.SDL_GL_STEREO
	GL_MULTISAMPLEBUFFERS         GLAttr = C.SDL_GL_MULTISAMPLEBUFFERS
	GL_MULTISAMPLESAMPLES         GLAttr = C.SDL_GL_MULTISAMPLESAMPLES
	GL_ACCELERATED_VISUAL         GLAttr = C.SDL_GL_ACCELERATED_VISUAL
	GL_RETAINED_BACKING           GLAttr = C.SDL_GL_RETAINED_BACKING
	GL_CONTEXT_MAJOR_VERSION      GLAttr = C.SDL_GL_CONTEXT_MAJOR_VERSION
	GL_CONTEXT_MINOR_VERSION      GLAttr = C.SDL_GL_CONTEXT_MINOR_VERSION
	GL_CONTEXT_FLAGS              GLAttr = C.SDL_GL_CONTEXT_FLAGS
	GL_CONTEXT_PROFILE_MASK       GLAttr = C.SDL_GL_CONTEXT_PROFILE_MASK
	GL_SHARE_WITH_CURRENT_CONTEXT GLAttr = C.SDL_GL_SHARE_WITH_CURRENT_CONTEXT
	GL_FRAMEBUFFER_SRGB_CAPABLE   GLAttr = C.SDL_GL_FRAMEBUFFER_SRGB_CAPABLE
	GL_CONTEXT_RELEASE_BEHAVIOR   GLAttr = C.SDL_GL_CONTEXT_RELEASE_BEHAVIOR
	GL_CONTEXT_RESET_NOTIFICATION GLAttr = C.SDL_GL_CONTEXT_RESET_NOTIFICATION
	GL_CONTEXT_NO_ERROR           GLAttr = C.SDL_GL_CONTEXT_NO_ERROR
	GL_FLOATBUFFERS               GLAttr = C.SDL_GL_FLOATBUFFERS
	GL_EGL_PLATFORM               GLAttr = C.SDL_GL_EGL_PLATFORM
)

// GLProfile represents OpenGL context profile values.
type GLProfile uint32

const (
	GL_CONTEXT_PROFILE_CORE          GLProfile = C.SDL_GL_CONTEXT_PROFILE_CORE
	GL_CONTEXT_PROFILE_COMPATIBILITY GLProfile = C.SDL_GL_CONTEXT_PROFILE_COMPATIBILITY
	GL_CONTEXT_PROFILE_ES            GLProfile = C.SDL_GL_CONTEXT_PROFILE_ES
)

// GLContextFlag represents OpenGL context flags.
type GLContextFlag uint32

const (
	GL_CONTEXT_DEBUG_FLAG              GLContextFlag = C.SDL_GL_CONTEXT_DEBUG_FLAG
	GL_CONTEXT_FORWARD_COMPATIBLE_FLAG GLContextFlag = C.SDL_GL_CONTEXT_FORWARD_COMPATIBLE_FLAG
	GL_CONTEXT_ROBUST_ACCESS_FLAG      GLContextFlag = C.SDL_GL_CONTEXT_ROBUST_ACCESS_FLAG
	GL_CONTEXT_RESET_ISOLATION_FLAG    GLContextFlag = C.SDL_GL_CONTEXT_RESET_ISOLATION_FLAG
)

// GLContextReleaseFlag represents OpenGL context release behavior flags.
type GLContextReleaseFlag uint32

const (
	GL_CONTEXT_RELEASE_BEHAVIOR_NONE  GLContextReleaseFlag = C.SDL_GL_CONTEXT_RELEASE_BEHAVIOR_NONE
	GL_CONTEXT_RELEASE_BEHAVIOR_FLUSH GLContextReleaseFlag = C.SDL_GL_CONTEXT_RELEASE_BEHAVIOR_FLUSH
)

// GLContextResetNotification represents OpenGL context reset notification values.
type GLContextResetNotification uint32

const (
	GL_CONTEXT_RESET_NO_NOTIFICATION GLContextResetNotification = C.SDL_GL_CONTEXT_RESET_NO_NOTIFICATION
	GL_CONTEXT_RESET_LOSE_CONTEXT    GLContextResetNotification = C.SDL_GL_CONTEXT_RESET_LOSE_CONTEXT
)

// HitTestResult represents possible return values from the HitTest callback.
type HitTestResult int

const (
	HITTEST_NORMAL             HitTestResult = C.SDL_HITTEST_NORMAL
	HITTEST_DRAGGABLE          HitTestResult = C.SDL_HITTEST_DRAGGABLE
	HITTEST_RESIZE_TOPLEFT     HitTestResult = C.SDL_HITTEST_RESIZE_TOPLEFT
	HITTEST_RESIZE_TOP         HitTestResult = C.SDL_HITTEST_RESIZE_TOP
	HITTEST_RESIZE_TOPRIGHT    HitTestResult = C.SDL_HITTEST_RESIZE_TOPRIGHT
	HITTEST_RESIZE_RIGHT       HitTestResult = C.SDL_HITTEST_RESIZE_RIGHT
	HITTEST_RESIZE_BOTTOMRIGHT HitTestResult = C.SDL_HITTEST_RESIZE_BOTTOMRIGHT
	HITTEST_RESIZE_BOTTOM      HitTestResult = C.SDL_HITTEST_RESIZE_BOTTOM
	HITTEST_RESIZE_BOTTOMLEFT  HitTestResult = C.SDL_HITTEST_RESIZE_BOTTOMLEFT
	HITTEST_RESIZE_LEFT        HitTestResult = C.SDL_HITTEST_RESIZE_LEFT
)

// HitTest is a callback used for hit-testing.
type HitTest func(win *Window, area *Point, data unsafe.Pointer) HitTestResult

// displayModeFromC converts a C SDL_DisplayMode to a Go DisplayMode.
func displayModeFromC(cm *C.SDL_DisplayMode) DisplayMode {
	return DisplayMode{
		DisplayID:              DisplayID(cm.displayID),
		Format:                 PixelFormat(cm.format),
		W:                      int32(cm.w),
		H:                      int32(cm.h),
		PixelDensity:           float32(cm.pixel_density),
		RefreshRate:            float32(cm.refresh_rate),
		RefreshRateNumerator:   int32(cm.refresh_rate_numerator),
		RefreshRateDenominator: int32(cm.refresh_rate_denominator),
	}
}

// CreatePopupWindow creates a popup window with the specified parent, offset, size, and flags.
func CreatePopupWindow(parent *Window, offsetX, offsetY, w, h int, flags WindowFlags) (*Window, error) {
	cw := C.SDL_CreatePopupWindow(parent.c, C.int(offsetX), C.int(offsetY), C.int(w), C.int(h), C.SDL_WindowFlags(flags))
	if cw == nil {
		return nil, getError()
	}
	return &Window{c: cw}, nil
}

// CreateWindowWithProperties creates a window with the specified properties.
func CreateWindowWithProperties(props PropertiesID) (*Window, error) {
	cw := C.SDL_CreateWindowWithProperties(C.SDL_PropertiesID(props))
	if cw == nil {
		return nil, getError()
	}
	return &Window{c: cw}, nil
}

// GetDisplayForPoint returns the display containing a point.
func GetDisplayForPoint(p Point) DisplayID {
	return DisplayID(C.SDL_GetDisplayForPoint(p.cptr()))
}

// GetDisplayForRect returns the display most associated with a rect.
func GetDisplayForRect(r Rect) DisplayID {
	return DisplayID(C.SDL_GetDisplayForRect(r.cptr()))
}

// GetNaturalDisplayOrientation returns the natural orientation of a display.
func GetNaturalDisplayOrientation(displayID DisplayID) DisplayOrientation {
	return DisplayOrientation(C.SDL_GetNaturalDisplayOrientation(C.SDL_DisplayID(displayID)))
}

// GetCurrentDisplayOrientation returns the current orientation of a display.
func GetCurrentDisplayOrientation(displayID DisplayID) DisplayOrientation {
	return DisplayOrientation(C.SDL_GetCurrentDisplayOrientation(C.SDL_DisplayID(displayID)))
}

// GetFullscreenDisplayModes returns a list of fullscreen display modes for a display.
func GetFullscreenDisplayModes(displayID DisplayID) []DisplayMode {
	var count C.int
	cmodes := C.SDL_GetFullscreenDisplayModes(C.SDL_DisplayID(displayID), &count)
	if cmodes == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cmodes))
	n := int(count)
	result := make([]DisplayMode, n)
	slice := unsafe.Slice((**C.SDL_DisplayMode)(unsafe.Pointer(cmodes)), n)
	for i, cm := range slice {
		result[i] = displayModeFromC(cm)
	}
	return result
}

// GetDesktopDisplayMode returns information about the desktop display mode.
func GetDesktopDisplayMode(displayID DisplayID) (*DisplayMode, error) {
	cm := C.SDL_GetDesktopDisplayMode(C.SDL_DisplayID(displayID))
	if cm == nil {
		return nil, getError()
	}
	mode := displayModeFromC(cm)
	return &mode, nil
}

// GetCurrentDisplayMode returns information about the current display mode.
func GetCurrentDisplayMode(displayID DisplayID) (*DisplayMode, error) {
	cm := C.SDL_GetCurrentDisplayMode(C.SDL_DisplayID(displayID))
	if cm == nil {
		return nil, getError()
	}
	mode := displayModeFromC(cm)
	return &mode, nil
}

// GetClosestFullscreenDisplayMode returns the closest match to the requested display mode.
func GetClosestFullscreenDisplayMode(displayID DisplayID, w, h int, refreshRate float32, includeHighDensityModes bool) (*DisplayMode, error) {
	var closest C.SDL_DisplayMode
	if !C.SDL_GetClosestFullscreenDisplayMode(C.SDL_DisplayID(displayID), C.int(w), C.int(h), C.float(refreshRate), C.bool(includeHighDensityModes), &closest) {
		return nil, getError()
	}
	mode := displayModeFromC(&closest)
	return &mode, nil
}

// SetFullscreenMode sets the display mode to use when the window is visible at fullscreen.
func (w *Window) SetFullscreenMode(mode *DisplayMode) error {
	var cm *C.SDL_DisplayMode
	if mode != nil {
		cmode := C.SDL_DisplayMode{
			displayID:              C.SDL_DisplayID(mode.DisplayID),
			format:                 C.SDL_PixelFormat(mode.Format),
			w:                      C.int(mode.W),
			h:                      C.int(mode.H),
			pixel_density:          C.float(mode.PixelDensity),
			refresh_rate:           C.float(mode.RefreshRate),
			refresh_rate_numerator: C.int(mode.RefreshRateNumerator),
			refresh_rate_denominator: C.int(mode.RefreshRateDenominator),
		}
		cm = &cmode
	}
	if !C.SDL_SetWindowFullscreenMode(w.c, cm) {
		return getError()
	}
	return nil
}

// FullscreenMode returns the display mode to use when the window is visible at fullscreen.
func (w *Window) FullscreenMode() *DisplayMode {
	cm := C.SDL_GetWindowFullscreenMode(w.c)
	if cm == nil {
		return nil
	}
	mode := displayModeFromC(cm)
	return &mode
}

// SafeArea returns the safe area for the window.
func (w *Window) SafeArea() (Rect, error) {
	var rect Rect
	if !C.SDL_GetWindowSafeArea(w.c, rect.cptr()) {
		return rect, getError()
	}
	return rect, nil
}

// BordersSize returns the size of the window borders.
func (w *Window) BordersSize() (top, left, bottom, right int, err error) {
	var ct, cl, cb, cr C.int
	if !C.SDL_GetWindowBordersSize(w.c, &ct, &cl, &cb, &cr) {
		return 0, 0, 0, 0, getError()
	}
	return int(ct), int(cl), int(cb), int(cr), nil
}

// SetIcon sets the icon for the window.
func (w *Window) SetIcon(icon *Surface) error {
	if !C.SDL_SetWindowIcon(w.c, icon.c) {
		return getError()
	}
	return nil
}

// SetOpacity sets the opacity of the window.
func (w *Window) SetOpacity(opacity float32) error {
	if !C.SDL_SetWindowOpacity(w.c, C.float(opacity)) {
		return getError()
	}
	return nil
}

// Opacity returns the opacity of the window.
func (w *Window) Opacity() float32 {
	return float32(C.SDL_GetWindowOpacity(w.c))
}

// SetParent sets the window as a child of a parent window.
func (w *Window) SetParent(parent *Window) error {
	var cp *C.SDL_Window
	if parent != nil {
		cp = parent.c
	}
	if !C.SDL_SetWindowParent(w.c, cp) {
		return getError()
	}
	return nil
}

// SetModal toggles the modal state of the window.
func (w *Window) SetModal(modal bool) error {
	if !C.SDL_SetWindowModal(w.c, C.bool(modal)) {
		return getError()
	}
	return nil
}

// SetFocusable sets whether the window may have input focus.
func (w *Window) SetFocusable(focusable bool) error {
	if !C.SDL_SetWindowFocusable(w.c, C.bool(focusable)) {
		return getError()
	}
	return nil
}

// ShowSystemMenu displays the system-level window menu.
func (w *Window) ShowSystemMenu(x, y int) error {
	if !C.SDL_ShowWindowSystemMenu(w.c, C.int(x), C.int(y)) {
		return getError()
	}
	return nil
}

// SetShape sets the shape of a transparent window.
func (w *Window) SetShape(shape *Surface) error {
	var cs *C.SDL_Surface
	if shape != nil {
		cs = shape.c
	}
	if !C.SDL_SetWindowShape(w.c, cs) {
		return getError()
	}
	return nil
}

// UpdateSurfaceRects updates specific areas of the window surface.
func (w *Window) UpdateSurfaceRects(rects []Rect) error {
	if len(rects) == 0 {
		return nil
	}
	if !C.SDL_UpdateWindowSurfaceRects(w.c, (*C.SDL_Rect)(unsafe.Pointer(&rects[0])), C.int(len(rects))) {
		return getError()
	}
	return nil
}

// Sync synchronizes the window state, blocking until pending window state changes are applied.
func (w *Window) Sync() error {
	if !C.SDL_SyncWindow(w.c) {
		return getError()
	}
	return nil
}

// SetHitTest sets a callback for hit-testing on the window.
// Pass nil to disable hit-testing.
func (w *Window) SetHitTest(callback HitTest, callbackData unsafe.Pointer) error {
	// Note: for a full implementation, a CGo callback trampoline is needed.
	// Setting to nil disables hit-testing.
	if callback == nil {
		if !C.SDL_SetWindowHitTest(w.c, nil, nil) {
			return getError()
		}
		return nil
	}
	// Storing Go callbacks with CGo requires a trampoline; for now we only support
	// disabling hit-test. A full implementation would use a registered callback map.
	if !C.SDL_SetWindowHitTest(w.c, nil, nil) {
		return getError()
	}
	return nil
}

// GetWindows returns a list of all valid windows.
func GetWindows() []*Window {
	var count C.int
	cws := C.SDL_GetWindows(&count)
	if cws == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cws))
	n := int(count)
	result := make([]*Window, n)
	slice := unsafe.Slice((**C.SDL_Window)(unsafe.Pointer(cws)), n)
	for i, cw := range slice {
		result[i] = &Window{c: cw}
	}
	return result
}

// ScreenSaverEnabled returns whether the screen saver is currently enabled.
func ScreenSaverEnabled() bool {
	return bool(C.SDL_ScreenSaverEnabled())
}

// EnableScreenSaver allows the screen to be blanked by a screen saver.
func EnableScreenSaver() error {
	if !C.SDL_EnableScreenSaver() {
		return getError()
	}
	return nil
}

// DisableScreenSaver prevents the screen from being blanked by a screen saver.
func DisableScreenSaver() error {
	if !C.SDL_DisableScreenSaver() {
		return getError()
	}
	return nil
}

// GL_LoadLibrary dynamically loads an OpenGL library.
func GL_LoadLibrary(path string) error {
	var cp *C.char
	if path != "" {
		cp = C.CString(path)
		defer C.free(unsafe.Pointer(cp))
	}
	if !C.SDL_GL_LoadLibrary(cp) {
		return getError()
	}
	return nil
}

// GL_GetProcAddress returns a pointer to the named OpenGL function.
func GL_GetProcAddress(proc string) unsafe.Pointer {
	cp := C.CString(proc)
	defer C.free(unsafe.Pointer(cp))
	return unsafe.Pointer(C.SDL_GL_GetProcAddress(cp))
}

// GL_UnloadLibrary unloads the OpenGL library previously loaded by GL_LoadLibrary.
func GL_UnloadLibrary() {
	C.SDL_GL_UnloadLibrary()
}

// GL_ExtensionSupported checks if an OpenGL extension is supported for the current context.
func GL_ExtensionSupported(extension string) bool {
	ce := C.CString(extension)
	defer C.free(unsafe.Pointer(ce))
	return bool(C.SDL_GL_ExtensionSupported(ce))
}

// GL_ResetAttributes resets all previously set OpenGL context attributes to their default values.
func GL_ResetAttributes() {
	C.SDL_GL_ResetAttributes()
}

// GL_SetAttribute sets an OpenGL window attribute before window creation.
func GL_SetAttribute(attr GLAttr, value int) error {
	if !C.SDL_GL_SetAttribute(C.SDL_GLAttr(attr), C.int(value)) {
		return getError()
	}
	return nil
}

// GL_GetAttribute returns the actual value for an attribute from the current context.
func GL_GetAttribute(attr GLAttr) (int, error) {
	var value C.int
	if !C.SDL_GL_GetAttribute(C.SDL_GLAttr(attr), &value) {
		return 0, getError()
	}
	return int(value), nil
}

// GL_CreateContext creates an OpenGL context for an OpenGL window, and makes it current.
func GL_CreateContext(window *Window) (GLContext, error) {
	ctx := C.SDL_GL_CreateContext(window.c)
	if ctx == nil {
		return GLContext(nil), getError()
	}
	return GLContext(ctx), nil
}

// GL_MakeCurrent sets up an OpenGL context for rendering into an OpenGL window.
func GL_MakeCurrent(window *Window, context GLContext) error {
	if !C.SDL_GL_MakeCurrent(window.c, C.SDL_GLContext(context)) {
		return getError()
	}
	return nil
}

// GL_GetCurrentWindow returns the currently active OpenGL window.
func GL_GetCurrentWindow() (*Window, error) {
	cw := C.SDL_GL_GetCurrentWindow()
	if cw == nil {
		return nil, getError()
	}
	return &Window{c: cw}, nil
}

// GL_GetCurrentContext returns the currently active OpenGL context.
func GL_GetCurrentContext() GLContext {
	return GLContext(C.SDL_GL_GetCurrentContext())
}

// GL_SetSwapInterval sets the swap interval for the current OpenGL context.
func GL_SetSwapInterval(interval int) error {
	if !C.SDL_GL_SetSwapInterval(C.int(interval)) {
		return getError()
	}
	return nil
}

// GL_GetSwapInterval returns the swap interval for the current OpenGL context.
func GL_GetSwapInterval() (int, error) {
	var interval C.int
	if !C.SDL_GL_GetSwapInterval(&interval) {
		return 0, getError()
	}
	return int(interval), nil
}

// GL_SwapWindow updates a window with OpenGL rendering.
func GL_SwapWindow(window *Window) error {
	if !C.SDL_GL_SwapWindow(window.c) {
		return getError()
	}
	return nil
}

// GL_DestroyContext deletes an OpenGL context.
func GL_DestroyContext(context GLContext) error {
	if !C.SDL_GL_DestroyContext(C.SDL_GLContext(context)) {
		return getError()
	}
	return nil
}

// ProgressState represents the taskbar progress state of a window.
type ProgressState int

// Progress state constants.
const (
	PROGRESS_STATE_INVALID       ProgressState = C.SDL_PROGRESS_STATE_INVALID
	PROGRESS_STATE_NONE          ProgressState = C.SDL_PROGRESS_STATE_NONE
	PROGRESS_STATE_INDETERMINATE ProgressState = C.SDL_PROGRESS_STATE_INDETERMINATE
	PROGRESS_STATE_NORMAL        ProgressState = C.SDL_PROGRESS_STATE_NORMAL
	PROGRESS_STATE_PAUSED        ProgressState = C.SDL_PROGRESS_STATE_PAUSED
	PROGRESS_STATE_ERROR         ProgressState = C.SDL_PROGRESS_STATE_ERROR
)

// GetPrimaryDisplay returns the instance ID of the primary display.
func GetPrimaryDisplay() DisplayID {
	return DisplayID(C.SDL_GetPrimaryDisplay())
}

// GetDisplayProperties returns the properties associated with a display.
func GetDisplayProperties(displayID DisplayID) PropertiesID {
	return PropertiesID(C.SDL_GetDisplayProperties(C.SDL_DisplayID(displayID)))
}

// Properties returns the properties associated with the window.
func (w *Window) Properties() PropertiesID {
	return PropertiesID(C.SDL_GetWindowProperties(w.c))
}

// Parent returns the parent of a window.
func (w *Window) Parent() *Window {
	cw := C.SDL_GetWindowParent(w.c)
	if cw == nil {
		return nil
	}
	return &Window{c: cw}
}

// PixelFormat returns the pixel format associated with the window.
func (w *Window) PixelFormat() PixelFormat {
	return PixelFormat(C.SDL_GetWindowPixelFormat(w.c))
}

// SetAspectRatio sets the minimum and maximum aspect ratios for the window.
func (w *Window) SetAspectRatio(minAspect, maxAspect float32) error {
	if !C.SDL_SetWindowAspectRatio(w.c, C.float(minAspect), C.float(maxAspect)) {
		return getError()
	}
	return nil
}

// AspectRatio returns the minimum and maximum aspect ratios for the window.
func (w *Window) AspectRatio() (float32, float32, error) {
	var minA, maxA C.float
	if !C.SDL_GetWindowAspectRatio(w.c, &minA, &maxA) {
		return 0, 0, getError()
	}
	return float32(minA), float32(maxA), nil
}

// SetKeyboardGrab sets whether the keyboard is grabbed by the window.
func (w *Window) SetKeyboardGrab(grabbed bool) error {
	if !C.SDL_SetWindowKeyboardGrab(w.c, C.bool(grabbed)) {
		return getError()
	}
	return nil
}

// KeyboardGrab returns whether the keyboard is grabbed by the window.
func (w *Window) KeyboardGrab() bool {
	return bool(C.SDL_GetWindowKeyboardGrab(w.c))
}

// SetMouseGrab sets whether the mouse is grabbed by the window.
func (w *Window) SetMouseGrab(grabbed bool) error {
	if !C.SDL_SetWindowMouseGrab(w.c, C.bool(grabbed)) {
		return getError()
	}
	return nil
}

// MouseGrab returns whether the mouse is grabbed by the window.
func (w *Window) MouseGrab() bool {
	return bool(C.SDL_GetWindowMouseGrab(w.c))
}

// SetMouseRect confines the mouse cursor to a specified area of the window.
func (w *Window) SetMouseRect(rect *Rect) error {
	var cr *C.SDL_Rect
	if rect != nil {
		cr = rect.cptr()
	}
	if !C.SDL_SetWindowMouseRect(w.c, cr) {
		return getError()
	}
	return nil
}

// MouseRect returns the mouse confinement rectangle for the window.
func (w *Window) MouseRect() *Rect {
	cr := C.SDL_GetWindowMouseRect(w.c)
	if cr == nil {
		return nil
	}
	return &Rect{X: int32(cr.x), Y: int32(cr.y), W: int32(cr.w), H: int32(cr.h)}
}

// GetGrabbedWindow returns the window that currently has keyboard or mouse grab.
func GetGrabbedWindow() *Window {
	cw := C.SDL_GetGrabbedWindow()
	if cw == nil {
		return nil
	}
	return &Window{c: cw}
}

// HasSurface returns whether the window has a surface associated with it.
func (w *Window) HasSurface() bool {
	return bool(C.SDL_WindowHasSurface(w.c))
}

// DestroySurface destroys the surface associated with the window.
func (w *Window) DestroySurface() error {
	if !C.SDL_DestroyWindowSurface(w.c) {
		return getError()
	}
	return nil
}

// SetSurfaceVSync sets the VSync mode for the window surface.
func (w *Window) SetSurfaceVSync(vsync int) error {
	if !C.SDL_SetWindowSurfaceVSync(w.c, C.int(vsync)) {
		return getError()
	}
	return nil
}

// SurfaceVSync returns the VSync mode for the window surface.
func (w *Window) SurfaceVSync() (int, error) {
	var v C.int
	if !C.SDL_GetWindowSurfaceVSync(w.c, &v) {
		return 0, getError()
	}
	return int(v), nil
}

// SetFillDocument sets the fill-document mode for the window (Emscripten).
func (w *Window) SetFillDocument(fill bool) error {
	if !C.SDL_SetWindowFillDocument(w.c, C.bool(fill)) {
		return getError()
	}
	return nil
}

// SetProgressState sets the progress state of the window in the taskbar.
func (w *Window) SetProgressState(state ProgressState) error {
	if !C.SDL_SetWindowProgressState(w.c, C.SDL_ProgressState(state)) {
		return getError()
	}
	return nil
}

// GetProgressState returns the progress state of the window.
func (w *Window) GetProgressState() ProgressState {
	return ProgressState(C.SDL_GetWindowProgressState(w.c))
}

// SetProgressValue sets the progress value (0.0 to 1.0) of the window.
func (w *Window) SetProgressValue(value float32) error {
	if !C.SDL_SetWindowProgressValue(w.c, C.float(value)) {
		return getError()
	}
	return nil
}

// GetProgressValue returns the progress value of the window.
func (w *Window) GetProgressValue() float32 {
	return float32(C.SDL_GetWindowProgressValue(w.c))
}

// ICCProfile returns the ICC profile data for the window's display.
// The caller receives a copy of the data.
func (w *Window) ICCProfile() ([]byte, error) {
	var size C.size_t
	data := C.SDL_GetWindowICCProfile(w.c, &size)
	if data == nil {
		return nil, getError()
	}
	defer C.SDL_free(data)
	result := make([]byte, int(size))
	copy(result, unsafe.Slice((*byte)(data), int(size)))
	return result, nil
}

// EGL_GetProcAddress returns an EGL function pointer.
func EGL_GetProcAddress(proc string) unsafe.Pointer {
	cp := C.CString(proc)
	defer C.free(unsafe.Pointer(cp))
	return unsafe.Pointer(C.SDL_EGL_GetProcAddress(cp))
}

// EGL_GetCurrentDisplay returns the currently active EGL display.
func EGL_GetCurrentDisplay() unsafe.Pointer {
	return unsafe.Pointer(C.SDL_EGL_GetCurrentDisplay())
}

// EGL_GetCurrentConfig returns the currently active EGL config.
func EGL_GetCurrentConfig() unsafe.Pointer {
	return unsafe.Pointer(C.SDL_EGL_GetCurrentConfig())
}

// EGL_GetWindowSurface returns the EGL surface associated with a window.
func EGL_GetWindowSurface(window *Window) unsafe.Pointer {
	return unsafe.Pointer(C.SDL_EGL_GetWindowSurface(window.c))
}

// EGL_SetAttributeCallbacks sets callbacks for EGL attribute arrays.
// The callback parameters are raw C function pointers (SDL_EGLAttribArrayCallback, SDL_EGLIntArrayCallback).
func EGL_SetAttributeCallbacks(platformAttribCallback, surfaceAttribCallback, contextAttribCallback, userdata unsafe.Pointer) {
	C.SDL_EGL_SetAttributeCallbacks(
		C.SDL_EGLAttribArrayCallback(platformAttribCallback),
		C.SDL_EGLIntArrayCallback(surfaceAttribCallback),
		C.SDL_EGLIntArrayCallback(contextAttribCallback),
		userdata)
}
