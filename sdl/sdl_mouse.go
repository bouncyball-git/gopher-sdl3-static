package sdl

/*
#include <SDL3/SDL.h>

extern void goMouseMotionTransformCallback(void *userdata, Uint64 timestamp, SDL_Window *window, SDL_MouseID mouseID, float *x, float *y);

static void cgoMouseMotionTransformTrampoline(void *userdata, Uint64 timestamp, SDL_Window *window, SDL_MouseID mouseID, float *x, float *y) {
	goMouseMotionTransformCallback(userdata, timestamp, window, mouseID, x, y);
}
*/
import "C"

import "unsafe"

// MouseID is a unique identifier for a mouse device.
type MouseID uint32

// Cursor represents an SDL cursor.
type Cursor struct {
	c *C.SDL_Cursor
}

// SystemCursor identifies a system cursor type.
type SystemCursor int

const (
	SYSTEM_CURSOR_DEFAULT     SystemCursor = C.SDL_SYSTEM_CURSOR_DEFAULT
	SYSTEM_CURSOR_TEXT        SystemCursor = C.SDL_SYSTEM_CURSOR_TEXT
	SYSTEM_CURSOR_WAIT        SystemCursor = C.SDL_SYSTEM_CURSOR_WAIT
	SYSTEM_CURSOR_CROSSHAIR   SystemCursor = C.SDL_SYSTEM_CURSOR_CROSSHAIR
	SYSTEM_CURSOR_PROGRESS    SystemCursor = C.SDL_SYSTEM_CURSOR_PROGRESS
	SYSTEM_CURSOR_NWSE_RESIZE SystemCursor = C.SDL_SYSTEM_CURSOR_NWSE_RESIZE
	SYSTEM_CURSOR_NESW_RESIZE SystemCursor = C.SDL_SYSTEM_CURSOR_NESW_RESIZE
	SYSTEM_CURSOR_EW_RESIZE   SystemCursor = C.SDL_SYSTEM_CURSOR_EW_RESIZE
	SYSTEM_CURSOR_NS_RESIZE   SystemCursor = C.SDL_SYSTEM_CURSOR_NS_RESIZE
	SYSTEM_CURSOR_MOVE        SystemCursor = C.SDL_SYSTEM_CURSOR_MOVE
	SYSTEM_CURSOR_NOT_ALLOWED SystemCursor = C.SDL_SYSTEM_CURSOR_NOT_ALLOWED
	SYSTEM_CURSOR_POINTER     SystemCursor = C.SDL_SYSTEM_CURSOR_POINTER
	SYSTEM_CURSOR_NW_RESIZE   SystemCursor = C.SDL_SYSTEM_CURSOR_NW_RESIZE
	SYSTEM_CURSOR_N_RESIZE    SystemCursor = C.SDL_SYSTEM_CURSOR_N_RESIZE
	SYSTEM_CURSOR_NE_RESIZE   SystemCursor = C.SDL_SYSTEM_CURSOR_NE_RESIZE
	SYSTEM_CURSOR_E_RESIZE    SystemCursor = C.SDL_SYSTEM_CURSOR_E_RESIZE
	SYSTEM_CURSOR_SE_RESIZE   SystemCursor = C.SDL_SYSTEM_CURSOR_SE_RESIZE
	SYSTEM_CURSOR_S_RESIZE    SystemCursor = C.SDL_SYSTEM_CURSOR_S_RESIZE
	SYSTEM_CURSOR_SW_RESIZE   SystemCursor = C.SDL_SYSTEM_CURSOR_SW_RESIZE
	SYSTEM_CURSOR_W_RESIZE    SystemCursor = C.SDL_SYSTEM_CURSOR_W_RESIZE
	SYSTEM_CURSOR_COUNT       SystemCursor = C.SDL_SYSTEM_CURSOR_COUNT
)

// MouseWheelDirection specifies the direction of a mouse wheel event.
type MouseWheelDirection int

const (
	MOUSEWHEEL_NORMAL  MouseWheelDirection = C.SDL_MOUSEWHEEL_NORMAL
	MOUSEWHEEL_FLIPPED MouseWheelDirection = C.SDL_MOUSEWHEEL_FLIPPED
)

// MouseButtonFlags represents the state of mouse buttons.
type MouseButtonFlags uint32

const (
	BUTTON_LEFT   = C.SDL_BUTTON_LEFT
	BUTTON_MIDDLE = C.SDL_BUTTON_MIDDLE
	BUTTON_RIGHT  = C.SDL_BUTTON_RIGHT
	BUTTON_X1     = C.SDL_BUTTON_X1
	BUTTON_X2     = C.SDL_BUTTON_X2

	BUTTON_LMASK  MouseButtonFlags = C.SDL_BUTTON_LMASK
	BUTTON_MMASK  MouseButtonFlags = C.SDL_BUTTON_MMASK
	BUTTON_RMASK  MouseButtonFlags = C.SDL_BUTTON_RMASK
	BUTTON_X1MASK MouseButtonFlags = C.SDL_BUTTON_X1MASK
	BUTTON_X2MASK MouseButtonFlags = C.SDL_BUTTON_X2MASK
)

// HasMouse returns true if a mouse is connected.
func HasMouse() bool {
	return bool(C.SDL_HasMouse())
}

// GetMouseFocus returns the window which currently has mouse focus.
func GetMouseFocus() *Window {
	cw := C.SDL_GetMouseFocus()
	if cw == nil {
		return nil
	}
	return &Window{c: cw}
}

// GetMouseState returns the current state of the mouse.
func GetMouseState() (x, y float32, buttons MouseButtonFlags) {
	var cx, cy C.float
	b := C.SDL_GetMouseState(&cx, &cy)
	return float32(cx), float32(cy), MouseButtonFlags(b)
}

// GetGlobalMouseState returns the current state of the mouse in global coordinates.
func GetGlobalMouseState() (x, y float32, buttons MouseButtonFlags) {
	var cx, cy C.float
	b := C.SDL_GetGlobalMouseState(&cx, &cy)
	return float32(cx), float32(cy), MouseButtonFlags(b)
}

// GetRelativeMouseState returns the relative state of the mouse.
func GetRelativeMouseState() (x, y float32, buttons MouseButtonFlags) {
	var cx, cy C.float
	b := C.SDL_GetRelativeMouseState(&cx, &cy)
	return float32(cx), float32(cy), MouseButtonFlags(b)
}

// WarpMouseInWindow moves the mouse cursor to the given position within the window.
func WarpMouseInWindow(window *Window, x, y float32) {
	var cw *C.SDL_Window
	if window != nil {
		cw = window.c
	}
	C.SDL_WarpMouseInWindow(cw, C.float(x), C.float(y))
}

// WarpMouseGlobal moves the mouse cursor to the given position in global screen space.
func WarpMouseGlobal(x, y float32) error {
	if !C.SDL_WarpMouseGlobal(C.float(x), C.float(y)) {
		return getError()
	}
	return nil
}

// SetWindowRelativeMouseMode sets relative mouse mode for a window.
func SetWindowRelativeMouseMode(window *Window, enabled bool) error {
	if !C.SDL_SetWindowRelativeMouseMode(window.c, C.bool(enabled)) {
		return getError()
	}
	return nil
}

// GetWindowRelativeMouseMode returns whether relative mouse mode is enabled for a window.
func GetWindowRelativeMouseMode(window *Window) bool {
	return bool(C.SDL_GetWindowRelativeMouseMode(window.c))
}

// CaptureMouse captures or releases the mouse.
func CaptureMouse(enabled bool) error {
	if !C.SDL_CaptureMouse(C.bool(enabled)) {
		return getError()
	}
	return nil
}

// CreateSystemCursor creates a system cursor.
func CreateSystemCursor(id SystemCursor) (*Cursor, error) {
	cc := C.SDL_CreateSystemCursor(C.SDL_SystemCursor(id))
	if cc == nil {
		return nil, getError()
	}
	return &Cursor{c: cc}, nil
}

// CreateColorCursor creates a color cursor from a surface.
func CreateColorCursor(surface *Surface, hotX, hotY int) (*Cursor, error) {
	cc := C.SDL_CreateColorCursor(surface.c, C.int(hotX), C.int(hotY))
	if cc == nil {
		return nil, getError()
	}
	return &Cursor{c: cc}, nil
}

// SetCursor sets the active cursor.
func SetCursor(cursor *Cursor) error {
	var cc *C.SDL_Cursor
	if cursor != nil {
		cc = cursor.c
	}
	if !C.SDL_SetCursor(cc) {
		return getError()
	}
	return nil
}

// GetCursor returns the active cursor.
func GetCursor() *Cursor {
	cc := C.SDL_GetCursor()
	if cc == nil {
		return nil
	}
	return &Cursor{c: cc}
}

// Destroy frees the cursor.
func (c *Cursor) Destroy() {
	if c.c != nil {
		C.SDL_DestroyCursor(c.c)
		c.c = nil
	}
}

// ShowCursor shows the cursor.
func ShowCursor() error {
	if !C.SDL_ShowCursor() {
		return getError()
	}
	return nil
}

// HideCursor hides the cursor.
func HideCursor() error {
	if !C.SDL_HideCursor() {
		return getError()
	}
	return nil
}

// CursorVisible returns true if the cursor is currently visible.
func CursorVisible() bool {
	return bool(C.SDL_CursorVisible())
}

// GetMouseNameForID returns the name of a mouse device.
func GetMouseNameForID(instanceID MouseID) string {
	return C.GoString(C.SDL_GetMouseNameForID(C.SDL_MouseID(instanceID)))
}

// GetMice returns a list of currently connected mouse IDs.
func GetMice() []MouseID {
	var count C.int
	cids := C.SDL_GetMice(&count)
	if cids == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cids))
	n := int(count)
	result := make([]MouseID, n)
	slice := unsafe.Slice((*C.SDL_MouseID)(cids), n)
	for i, id := range slice {
		result[i] = MouseID(id)
	}
	return result
}

// CreateCursor creates a cursor using the specified bitmap data and mask (in MSB format).
// The cursor width (w) must be a multiple of 8 bits.
func CreateCursor(data, mask []byte, w, h, hotX, hotY int) (*Cursor, error) {
	var cd, cm *C.Uint8
	if len(data) > 0 {
		cd = (*C.Uint8)(unsafe.Pointer(&data[0]))
	}
	if len(mask) > 0 {
		cm = (*C.Uint8)(unsafe.Pointer(&mask[0]))
	}
	cc := C.SDL_CreateCursor(cd, cm, C.int(w), C.int(h), C.int(hotX), C.int(hotY))
	if cc == nil {
		return nil, getError()
	}
	return &Cursor{c: cc}, nil
}

// CursorFrameInfo describes a single frame of an animated cursor.
type CursorFrameInfo struct {
	Surface  *Surface
	Duration uint32
}

// CreateAnimatedCursor creates an animated color cursor from a sequence of frames.
func CreateAnimatedCursor(frames []CursorFrameInfo, hotX, hotY int) (*Cursor, error) {
	if len(frames) == 0 {
		return nil, getError()
	}
	cframes := make([]C.SDL_CursorFrameInfo, len(frames))
	for i, f := range frames {
		cframes[i].surface = f.Surface.c
		cframes[i].duration = C.Uint32(f.Duration)
	}
	cc := C.SDL_CreateAnimatedCursor(&cframes[0], C.int(len(frames)), C.int(hotX), C.int(hotY))
	if cc == nil {
		return nil, getError()
	}
	return &Cursor{c: cc}, nil
}

// GetDefaultCursor returns the default cursor.
func GetDefaultCursor() *Cursor {
	cc := C.SDL_GetDefaultCursor()
	if cc == nil {
		return nil
	}
	return &Cursor{c: cc}
}

// MouseMotionTransformFunc transforms relative mouse motion values.
type MouseMotionTransformFunc func(timestamp uint64, window *Window, mouseID MouseID, x, y *float32)

//export goMouseMotionTransformCallback
func goMouseMotionTransformCallback(userdata unsafe.Pointer, timestamp C.Uint64, window *C.SDL_Window, mouseID C.SDL_MouseID, x, y *C.float) {
	id := uintptr(userdata)
	fn := getCallback(id).(MouseMotionTransformFunc)
	gx := float32(*x)
	gy := float32(*y)
	var w *Window
	if window != nil {
		w = &Window{c: window}
	}
	fn(uint64(timestamp), w, MouseID(mouseID), &gx, &gy)
	*x = C.float(gx)
	*y = C.float(gy)
}

// SetRelativeMouseTransform sets a callback to transform relative mouse motion.
// Pass nil to disable the transform.
func SetRelativeMouseTransform(callback MouseMotionTransformFunc) error {
	if callback == nil {
		if !C.SDL_SetRelativeMouseTransform(nil, nil) {
			return getError()
		}
		return nil
	}
	id := registerCallback(callback)
	if !C.SDL_SetRelativeMouseTransform(C.SDL_MouseMotionTransformCallback(C.cgoMouseMotionTransformTrampoline), ptrFromID(id)) {
		unregisterCallback(id)
		return getError()
	}
	return nil
}
