package sdl

/*
#include <SDL3/SDL.h>

extern bool goEventFilter(void *userdata, SDL_Event *event);

static bool cgoEventFilterTrampoline(void *userdata, SDL_Event *event) {
	return goEventFilter(userdata, event);
}

static SDL_EventFilter _get_event_filter_trampoline(void) {
	return cgoEventFilterTrampoline;
}
*/
import "C"

import "unsafe"

// ---------------------------------------------------------------------------
// EventType
// ---------------------------------------------------------------------------

// EventType identifies the type of an SDL event.
type EventType uint32

const (
	EVENT_FIRST EventType = C.SDL_EVENT_FIRST

	// Application events
	EVENT_QUIT                  EventType = C.SDL_EVENT_QUIT
	EVENT_TERMINATING           EventType = C.SDL_EVENT_TERMINATING
	EVENT_LOW_MEMORY            EventType = C.SDL_EVENT_LOW_MEMORY
	EVENT_WILL_ENTER_BACKGROUND EventType = C.SDL_EVENT_WILL_ENTER_BACKGROUND
	EVENT_DID_ENTER_BACKGROUND  EventType = C.SDL_EVENT_DID_ENTER_BACKGROUND
	EVENT_WILL_ENTER_FOREGROUND EventType = C.SDL_EVENT_WILL_ENTER_FOREGROUND
	EVENT_DID_ENTER_FOREGROUND  EventType = C.SDL_EVENT_DID_ENTER_FOREGROUND
	EVENT_LOCALE_CHANGED        EventType = C.SDL_EVENT_LOCALE_CHANGED
	EVENT_SYSTEM_THEME_CHANGED  EventType = C.SDL_EVENT_SYSTEM_THEME_CHANGED

	// Display events
	EVENT_DISPLAY_ORIENTATION          EventType = C.SDL_EVENT_DISPLAY_ORIENTATION
	EVENT_DISPLAY_ADDED                EventType = C.SDL_EVENT_DISPLAY_ADDED
	EVENT_DISPLAY_REMOVED              EventType = C.SDL_EVENT_DISPLAY_REMOVED
	EVENT_DISPLAY_MOVED                EventType = C.SDL_EVENT_DISPLAY_MOVED
	EVENT_DISPLAY_DESKTOP_MODE_CHANGED EventType = C.SDL_EVENT_DISPLAY_DESKTOP_MODE_CHANGED
	EVENT_DISPLAY_CURRENT_MODE_CHANGED EventType = C.SDL_EVENT_DISPLAY_CURRENT_MODE_CHANGED
	EVENT_DISPLAY_CONTENT_SCALE_CHANGED EventType = C.SDL_EVENT_DISPLAY_CONTENT_SCALE_CHANGED
	EVENT_DISPLAY_USABLE_BOUNDS_CHANGED EventType = C.SDL_EVENT_DISPLAY_USABLE_BOUNDS_CHANGED
	EVENT_DISPLAY_FIRST                EventType = C.SDL_EVENT_DISPLAY_FIRST
	EVENT_DISPLAY_LAST                 EventType = C.SDL_EVENT_DISPLAY_LAST

	// Window events
	EVENT_WINDOW_SHOWN                EventType = C.SDL_EVENT_WINDOW_SHOWN
	EVENT_WINDOW_HIDDEN               EventType = C.SDL_EVENT_WINDOW_HIDDEN
	EVENT_WINDOW_EXPOSED              EventType = C.SDL_EVENT_WINDOW_EXPOSED
	EVENT_WINDOW_MOVED                EventType = C.SDL_EVENT_WINDOW_MOVED
	EVENT_WINDOW_RESIZED              EventType = C.SDL_EVENT_WINDOW_RESIZED
	EVENT_WINDOW_PIXEL_SIZE_CHANGED   EventType = C.SDL_EVENT_WINDOW_PIXEL_SIZE_CHANGED
	EVENT_WINDOW_METAL_VIEW_RESIZED   EventType = C.SDL_EVENT_WINDOW_METAL_VIEW_RESIZED
	EVENT_WINDOW_MINIMIZED            EventType = C.SDL_EVENT_WINDOW_MINIMIZED
	EVENT_WINDOW_MAXIMIZED            EventType = C.SDL_EVENT_WINDOW_MAXIMIZED
	EVENT_WINDOW_RESTORED             EventType = C.SDL_EVENT_WINDOW_RESTORED
	EVENT_WINDOW_MOUSE_ENTER          EventType = C.SDL_EVENT_WINDOW_MOUSE_ENTER
	EVENT_WINDOW_MOUSE_LEAVE          EventType = C.SDL_EVENT_WINDOW_MOUSE_LEAVE
	EVENT_WINDOW_FOCUS_GAINED         EventType = C.SDL_EVENT_WINDOW_FOCUS_GAINED
	EVENT_WINDOW_FOCUS_LOST           EventType = C.SDL_EVENT_WINDOW_FOCUS_LOST
	EVENT_WINDOW_CLOSE_REQUESTED      EventType = C.SDL_EVENT_WINDOW_CLOSE_REQUESTED
	EVENT_WINDOW_HIT_TEST             EventType = C.SDL_EVENT_WINDOW_HIT_TEST
	EVENT_WINDOW_ICCPROF_CHANGED      EventType = C.SDL_EVENT_WINDOW_ICCPROF_CHANGED
	EVENT_WINDOW_DISPLAY_CHANGED      EventType = C.SDL_EVENT_WINDOW_DISPLAY_CHANGED
	EVENT_WINDOW_DISPLAY_SCALE_CHANGED EventType = C.SDL_EVENT_WINDOW_DISPLAY_SCALE_CHANGED
	EVENT_WINDOW_SAFE_AREA_CHANGED    EventType = C.SDL_EVENT_WINDOW_SAFE_AREA_CHANGED
	EVENT_WINDOW_OCCLUDED             EventType = C.SDL_EVENT_WINDOW_OCCLUDED
	EVENT_WINDOW_ENTER_FULLSCREEN     EventType = C.SDL_EVENT_WINDOW_ENTER_FULLSCREEN
	EVENT_WINDOW_LEAVE_FULLSCREEN     EventType = C.SDL_EVENT_WINDOW_LEAVE_FULLSCREEN
	EVENT_WINDOW_DESTROYED            EventType = C.SDL_EVENT_WINDOW_DESTROYED
	EVENT_WINDOW_HDR_STATE_CHANGED    EventType = C.SDL_EVENT_WINDOW_HDR_STATE_CHANGED
	EVENT_WINDOW_FIRST                EventType = C.SDL_EVENT_WINDOW_FIRST
	EVENT_WINDOW_LAST                 EventType = C.SDL_EVENT_WINDOW_LAST

	// Keyboard events
	EVENT_KEY_DOWN                 EventType = C.SDL_EVENT_KEY_DOWN
	EVENT_KEY_UP                   EventType = C.SDL_EVENT_KEY_UP
	EVENT_TEXT_EDITING              EventType = C.SDL_EVENT_TEXT_EDITING
	EVENT_TEXT_INPUT                EventType = C.SDL_EVENT_TEXT_INPUT
	EVENT_KEYMAP_CHANGED           EventType = C.SDL_EVENT_KEYMAP_CHANGED
	EVENT_KEYBOARD_ADDED           EventType = C.SDL_EVENT_KEYBOARD_ADDED
	EVENT_KEYBOARD_REMOVED         EventType = C.SDL_EVENT_KEYBOARD_REMOVED
	EVENT_TEXT_EDITING_CANDIDATES  EventType = C.SDL_EVENT_TEXT_EDITING_CANDIDATES
	EVENT_SCREEN_KEYBOARD_SHOWN   EventType = C.SDL_EVENT_SCREEN_KEYBOARD_SHOWN
	EVENT_SCREEN_KEYBOARD_HIDDEN  EventType = C.SDL_EVENT_SCREEN_KEYBOARD_HIDDEN

	// Mouse events
	EVENT_MOUSE_MOTION      EventType = C.SDL_EVENT_MOUSE_MOTION
	EVENT_MOUSE_BUTTON_DOWN EventType = C.SDL_EVENT_MOUSE_BUTTON_DOWN
	EVENT_MOUSE_BUTTON_UP   EventType = C.SDL_EVENT_MOUSE_BUTTON_UP
	EVENT_MOUSE_WHEEL       EventType = C.SDL_EVENT_MOUSE_WHEEL
	EVENT_MOUSE_ADDED       EventType = C.SDL_EVENT_MOUSE_ADDED
	EVENT_MOUSE_REMOVED     EventType = C.SDL_EVENT_MOUSE_REMOVED

	// Joystick events
	EVENT_JOYSTICK_AXIS_MOTION     EventType = C.SDL_EVENT_JOYSTICK_AXIS_MOTION
	EVENT_JOYSTICK_BALL_MOTION     EventType = C.SDL_EVENT_JOYSTICK_BALL_MOTION
	EVENT_JOYSTICK_HAT_MOTION      EventType = C.SDL_EVENT_JOYSTICK_HAT_MOTION
	EVENT_JOYSTICK_BUTTON_DOWN     EventType = C.SDL_EVENT_JOYSTICK_BUTTON_DOWN
	EVENT_JOYSTICK_BUTTON_UP       EventType = C.SDL_EVENT_JOYSTICK_BUTTON_UP
	EVENT_JOYSTICK_ADDED           EventType = C.SDL_EVENT_JOYSTICK_ADDED
	EVENT_JOYSTICK_REMOVED         EventType = C.SDL_EVENT_JOYSTICK_REMOVED
	EVENT_JOYSTICK_BATTERY_UPDATED EventType = C.SDL_EVENT_JOYSTICK_BATTERY_UPDATED
	EVENT_JOYSTICK_UPDATE_COMPLETE EventType = C.SDL_EVENT_JOYSTICK_UPDATE_COMPLETE

	// Gamepad events
	EVENT_GAMEPAD_AXIS_MOTION        EventType = C.SDL_EVENT_GAMEPAD_AXIS_MOTION
	EVENT_GAMEPAD_BUTTON_DOWN        EventType = C.SDL_EVENT_GAMEPAD_BUTTON_DOWN
	EVENT_GAMEPAD_BUTTON_UP          EventType = C.SDL_EVENT_GAMEPAD_BUTTON_UP
	EVENT_GAMEPAD_ADDED              EventType = C.SDL_EVENT_GAMEPAD_ADDED
	EVENT_GAMEPAD_REMOVED            EventType = C.SDL_EVENT_GAMEPAD_REMOVED
	EVENT_GAMEPAD_REMAPPED           EventType = C.SDL_EVENT_GAMEPAD_REMAPPED
	EVENT_GAMEPAD_TOUCHPAD_DOWN      EventType = C.SDL_EVENT_GAMEPAD_TOUCHPAD_DOWN
	EVENT_GAMEPAD_TOUCHPAD_MOTION    EventType = C.SDL_EVENT_GAMEPAD_TOUCHPAD_MOTION
	EVENT_GAMEPAD_TOUCHPAD_UP        EventType = C.SDL_EVENT_GAMEPAD_TOUCHPAD_UP
	EVENT_GAMEPAD_SENSOR_UPDATE      EventType = C.SDL_EVENT_GAMEPAD_SENSOR_UPDATE
	EVENT_GAMEPAD_UPDATE_COMPLETE    EventType = C.SDL_EVENT_GAMEPAD_UPDATE_COMPLETE
	EVENT_GAMEPAD_STEAM_HANDLE_UPDATED EventType = C.SDL_EVENT_GAMEPAD_STEAM_HANDLE_UPDATED

	// Touch events
	EVENT_FINGER_DOWN      EventType = C.SDL_EVENT_FINGER_DOWN
	EVENT_FINGER_UP        EventType = C.SDL_EVENT_FINGER_UP
	EVENT_FINGER_MOTION    EventType = C.SDL_EVENT_FINGER_MOTION
	EVENT_FINGER_CANCELED  EventType = C.SDL_EVENT_FINGER_CANCELED

	// Pinch events
	EVENT_PINCH_BEGIN  EventType = C.SDL_EVENT_PINCH_BEGIN
	EVENT_PINCH_UPDATE EventType = C.SDL_EVENT_PINCH_UPDATE
	EVENT_PINCH_END    EventType = C.SDL_EVENT_PINCH_END

	// Clipboard events
	EVENT_CLIPBOARD_UPDATE EventType = C.SDL_EVENT_CLIPBOARD_UPDATE

	// Drag and drop events
	EVENT_DROP_FILE     EventType = C.SDL_EVENT_DROP_FILE
	EVENT_DROP_TEXT     EventType = C.SDL_EVENT_DROP_TEXT
	EVENT_DROP_BEGIN    EventType = C.SDL_EVENT_DROP_BEGIN
	EVENT_DROP_COMPLETE EventType = C.SDL_EVENT_DROP_COMPLETE
	EVENT_DROP_POSITION EventType = C.SDL_EVENT_DROP_POSITION

	// Audio hotplug events
	EVENT_AUDIO_DEVICE_ADDED          EventType = C.SDL_EVENT_AUDIO_DEVICE_ADDED
	EVENT_AUDIO_DEVICE_REMOVED        EventType = C.SDL_EVENT_AUDIO_DEVICE_REMOVED
	EVENT_AUDIO_DEVICE_FORMAT_CHANGED EventType = C.SDL_EVENT_AUDIO_DEVICE_FORMAT_CHANGED

	// Sensor events
	EVENT_SENSOR_UPDATE EventType = C.SDL_EVENT_SENSOR_UPDATE

	// Pen events
	EVENT_PEN_PROXIMITY_IN  EventType = C.SDL_EVENT_PEN_PROXIMITY_IN
	EVENT_PEN_PROXIMITY_OUT EventType = C.SDL_EVENT_PEN_PROXIMITY_OUT
	EVENT_PEN_DOWN          EventType = C.SDL_EVENT_PEN_DOWN
	EVENT_PEN_UP            EventType = C.SDL_EVENT_PEN_UP
	EVENT_PEN_BUTTON_DOWN   EventType = C.SDL_EVENT_PEN_BUTTON_DOWN
	EVENT_PEN_BUTTON_UP     EventType = C.SDL_EVENT_PEN_BUTTON_UP
	EVENT_PEN_MOTION        EventType = C.SDL_EVENT_PEN_MOTION
	EVENT_PEN_AXIS          EventType = C.SDL_EVENT_PEN_AXIS

	// Camera hotplug events
	EVENT_CAMERA_DEVICE_ADDED    EventType = C.SDL_EVENT_CAMERA_DEVICE_ADDED
	EVENT_CAMERA_DEVICE_REMOVED  EventType = C.SDL_EVENT_CAMERA_DEVICE_REMOVED
	EVENT_CAMERA_DEVICE_APPROVED EventType = C.SDL_EVENT_CAMERA_DEVICE_APPROVED
	EVENT_CAMERA_DEVICE_DENIED   EventType = C.SDL_EVENT_CAMERA_DEVICE_DENIED

	// Render events
	EVENT_RENDER_TARGETS_RESET EventType = C.SDL_EVENT_RENDER_TARGETS_RESET
	EVENT_RENDER_DEVICE_RESET  EventType = C.SDL_EVENT_RENDER_DEVICE_RESET
	EVENT_RENDER_DEVICE_LOST   EventType = C.SDL_EVENT_RENDER_DEVICE_LOST

	// Internal / reserved
	EVENT_PRIVATE0      EventType = C.SDL_EVENT_PRIVATE0
	EVENT_PRIVATE1      EventType = C.SDL_EVENT_PRIVATE1
	EVENT_PRIVATE2      EventType = C.SDL_EVENT_PRIVATE2
	EVENT_PRIVATE3      EventType = C.SDL_EVENT_PRIVATE3
	EVENT_POLL_SENTINEL EventType = C.SDL_EVENT_POLL_SENTINEL

	EVENT_USER EventType = C.SDL_EVENT_USER
	EVENT_LAST EventType = C.SDL_EVENT_LAST
)

// ---------------------------------------------------------------------------
// Event interface
// ---------------------------------------------------------------------------

// Event is implemented by every concrete event type.
type Event interface {
	GetType() EventType
	GetTimestamp() uint64
}

// ---------------------------------------------------------------------------
// Concrete event types
// ---------------------------------------------------------------------------

// CommonEvent contains fields shared by every event.
type CommonEvent struct {
	Type      EventType
	Timestamp uint64
}

func (e *CommonEvent) GetType() EventType    { return e.Type }
func (e *CommonEvent) GetTimestamp() uint64   { return e.Timestamp }

// QuitEvent is generated when the user requests the application to quit.
type QuitEvent struct {
	CommonEvent
}

// DisplayEvent carries display state change data.
type DisplayEvent struct {
	CommonEvent
	DisplayID DisplayID
	Data1     int32
	Data2     int32
}

// WindowEvent carries window state change data.
type WindowEvent struct {
	CommonEvent
	WindowID WindowID
	Data1    int32
	Data2    int32
}

// KeyboardEvent carries keyboard button press/release data.
type KeyboardEvent struct {
	CommonEvent
	WindowID WindowID
	Which    KeyboardID
	Scancode Scancode
	Key      Keycode
	Mod      Keymod
	Raw      uint16
	Down     bool
	Repeat   bool
}

// TextEditingEvent carries text composition data.
type TextEditingEvent struct {
	CommonEvent
	WindowID WindowID
	Text     string
	Start    int32
	Length   int32
}

// TextInputEvent carries committed text input data.
type TextInputEvent struct {
	CommonEvent
	WindowID WindowID
	Text     string
}

// MouseMotionEvent carries mouse movement data.
type MouseMotionEvent struct {
	CommonEvent
	WindowID WindowID
	Which    MouseID
	State    MouseButtonFlags
	X, Y     float32
	XRel     float32
	YRel     float32
}

// MouseButtonEvent carries mouse button press/release data.
type MouseButtonEvent struct {
	CommonEvent
	WindowID WindowID
	Which    MouseID
	Button   uint8
	Down     bool
	Clicks   uint8
	X, Y     float32
}

// MouseWheelEvent carries mouse wheel scroll data.
type MouseWheelEvent struct {
	CommonEvent
	WindowID  WindowID
	Which     MouseID
	X, Y      float32
	Direction MouseWheelDirection
}

// DropEvent carries drag-and-drop data.
type DropEvent struct {
	CommonEvent
	WindowID WindowID
	X, Y     float32
	Source   string
	Data     string
}

// AudioDeviceEvent carries audio device hotplug data.
type AudioDeviceEvent struct {
	CommonEvent
	Which     uint32
	Recording bool
}

// JoystickAxisEvent carries joystick axis motion data.
type JoystickAxisEvent struct {
	CommonEvent
	Which JoystickID
	Axis  uint8
	Value int16
}

// JoystickButtonEvent carries joystick button press/release data.
type JoystickButtonEvent struct {
	CommonEvent
	Which  JoystickID
	Button uint8
	Down   bool
}

// JoystickHatEvent carries joystick hat position change data.
type JoystickHatEvent struct {
	CommonEvent
	Which JoystickID
	Hat   uint8
	Value uint8
}

// GamepadAxisEvent carries gamepad axis motion data.
type GamepadAxisEvent struct {
	CommonEvent
	Which JoystickID
	Axis  int32
	Value int16
}

// GamepadButtonEvent carries gamepad button press/release data.
type GamepadButtonEvent struct {
	CommonEvent
	Which  JoystickID
	Button int32
	Down   bool
}

// TouchFingerEvent carries touch finger data.
type TouchFingerEvent struct {
	CommonEvent
	TouchID    uint64
	FingerID   uint64
	X, Y       float32
	DX, DY     float32
	Pressure   float32
	WindowID   WindowID
}

// PenMotionEvent carries pen motion data.
type PenMotionEvent struct {
	CommonEvent
	WindowID WindowID
	Which    PenID
	X, Y     float32
}

// PenButtonEvent carries pen button press/release data.
type PenButtonEvent struct {
	CommonEvent
	WindowID WindowID
	Which    PenID
	Button   uint8
	Down     bool
}

// SensorEvent carries sensor update data.
type SensorEvent struct {
	CommonEvent
	Which           SensorID
	Data            [6]float32
	SensorTimestamp uint64
}

// UserEvent carries user-defined event data.
type UserEvent struct {
	CommonEvent
	WindowID WindowID
	Code     int32
	Data1    unsafe.Pointer
	Data2    unsafe.Pointer
}

// JoystickDeviceEvent is generated when a joystick is added or removed.
type JoystickDeviceEvent struct {
	CommonEvent
	Which JoystickID
}

// JoystickBallEvent is generated for joystick trackball motion.
type JoystickBallEvent struct {
	CommonEvent
	Which JoystickID
	Ball  uint8
	XRel  int16
	YRel  int16
}

// JoystickBatteryEvent is generated when a joystick battery state changes.
type JoystickBatteryEvent struct {
	CommonEvent
	Which   JoystickID
	State   PowerState
	Percent int
}

// GamepadDeviceEvent is generated when a gamepad is added, removed, or remapped.
type GamepadDeviceEvent struct {
	CommonEvent
	Which JoystickID
}

// GamepadTouchpadEvent is generated for gamepad touchpad activity.
type GamepadTouchpadEvent struct {
	CommonEvent
	Which    JoystickID
	Touchpad int32
	Finger   int32
	X, Y     float32
	Pressure float32
}

// GamepadSensorEvent is generated for gamepad sensor updates.
type GamepadSensorEvent struct {
	CommonEvent
	Which           JoystickID
	Sensor          int32
	Data            [3]float32
	SensorTimestamp uint64
}

// MouseDeviceEvent is generated when a mouse is added or removed.
type MouseDeviceEvent struct {
	CommonEvent
	Which MouseID
}

// CameraDeviceEvent is generated when a camera device is added, removed, approved, or denied.
type CameraDeviceEvent struct {
	CommonEvent
	Which CameraID
}

// RenderEvent is generated for render target/device state changes.
type RenderEvent struct {
	CommonEvent
	WindowID WindowID
}

// ClipboardEvent is generated when clipboard contents change.
type ClipboardEvent struct {
	CommonEvent
	Owner        bool
	NumMimeTypes int32
}

// PenProximityEvent is generated when a pen enters or leaves proximity.
type PenProximityEvent struct {
	CommonEvent
	WindowID WindowID
	Which    PenID
}

// PenTouchEvent is generated when a pen touches or lifts from a surface.
type PenTouchEvent struct {
	CommonEvent
	WindowID WindowID
	Which    PenID
	PenState PenInputFlags
	X, Y     float32
	Eraser   bool
	Down     bool
}

// PenAxisEvent is generated when a pen axis value changes.
type PenAxisEvent struct {
	CommonEvent
	WindowID WindowID
	Which    PenID
	PenState PenInputFlags
	X, Y     float32
	Axis     PenAxis
	Value    float32
}

// PinchFingerEvent is generated for pinch gestures.
type PinchFingerEvent struct {
	CommonEvent
	Scale    float32
	WindowID WindowID
}

// TextEditingCandidatesEvent is generated for IME candidate lists.
type TextEditingCandidatesEvent struct {
	CommonEvent
	WindowID          WindowID
	Candidates        []string
	SelectedCandidate int32
	Horizontal        bool
}

// ---------------------------------------------------------------------------
// convertEvent turns a C SDL_Event into the appropriate Go Event struct.
// ---------------------------------------------------------------------------

func convertEvent(cEvent *C.SDL_Event) Event {
	typ := EventType(*(*C.Uint32)(unsafe.Pointer(cEvent)))
	common := func() CommonEvent {
		ce := (*C.SDL_CommonEvent)(unsafe.Pointer(cEvent))
		return CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)}
	}

	switch {
	// Quit / app lifecycle
	case typ == EVENT_QUIT,
		typ == EVENT_TERMINATING,
		typ == EVENT_LOW_MEMORY,
		typ == EVENT_WILL_ENTER_BACKGROUND,
		typ == EVENT_DID_ENTER_BACKGROUND,
		typ == EVENT_WILL_ENTER_FOREGROUND,
		typ == EVENT_DID_ENTER_FOREGROUND,
		typ == EVENT_LOCALE_CHANGED,
		typ == EVENT_SYSTEM_THEME_CHANGED:
		return &QuitEvent{CommonEvent: common()}

	// Display events
	case typ >= EVENT_DISPLAY_FIRST && typ <= EVENT_DISPLAY_LAST:
		ce := (*C.SDL_DisplayEvent)(unsafe.Pointer(cEvent))
		return &DisplayEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			DisplayID:   DisplayID(ce.displayID),
			Data1:       int32(ce.data1),
			Data2:       int32(ce.data2),
		}

	// Window events
	case typ >= EVENT_WINDOW_FIRST && typ <= EVENT_WINDOW_LAST:
		ce := (*C.SDL_WindowEvent)(unsafe.Pointer(cEvent))
		return &WindowEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			WindowID:    WindowID(ce.windowID),
			Data1:       int32(ce.data1),
			Data2:       int32(ce.data2),
		}

	// Keyboard key events
	case typ == EVENT_KEY_DOWN || typ == EVENT_KEY_UP:
		ce := (*C.SDL_KeyboardEvent)(unsafe.Pointer(cEvent))
		return &KeyboardEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			WindowID:    WindowID(ce.windowID),
			Which:       KeyboardID(ce.which),
			Scancode:    Scancode(ce.scancode),
			Key:         Keycode(ce.key),
			Mod:         Keymod(ce.mod),
			Raw:         uint16(ce.raw),
			Down:        bool(ce.down),
			Repeat:      bool(ce.repeat),
		}

	// Text editing
	case typ == EVENT_TEXT_EDITING:
		ce := (*C.SDL_TextEditingEvent)(unsafe.Pointer(cEvent))
		var text string
		if ce.text != nil {
			text = C.GoString(ce.text)
		}
		return &TextEditingEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			WindowID:    WindowID(ce.windowID),
			Text:        text,
			Start:       int32(ce.start),
			Length:       int32(ce.length),
		}

	// Text input
	case typ == EVENT_TEXT_INPUT:
		ce := (*C.SDL_TextInputEvent)(unsafe.Pointer(cEvent))
		var text string
		if ce.text != nil {
			text = C.GoString(ce.text)
		}
		return &TextInputEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			WindowID:    WindowID(ce.windowID),
			Text:        text,
		}

	// Mouse motion
	case typ == EVENT_MOUSE_MOTION:
		ce := (*C.SDL_MouseMotionEvent)(unsafe.Pointer(cEvent))
		return &MouseMotionEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			WindowID:    WindowID(ce.windowID),
			Which:       MouseID(ce.which),
			State:       MouseButtonFlags(ce.state),
			X:           float32(ce.x),
			Y:           float32(ce.y),
			XRel:        float32(ce.xrel),
			YRel:        float32(ce.yrel),
		}

	// Mouse button
	case typ == EVENT_MOUSE_BUTTON_DOWN || typ == EVENT_MOUSE_BUTTON_UP:
		ce := (*C.SDL_MouseButtonEvent)(unsafe.Pointer(cEvent))
		return &MouseButtonEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			WindowID:    WindowID(ce.windowID),
			Which:       MouseID(ce.which),
			Button:      uint8(ce.button),
			Down:        bool(ce.down),
			Clicks:      uint8(ce.clicks),
			X:           float32(ce.x),
			Y:           float32(ce.y),
		}

	// Mouse wheel
	case typ == EVENT_MOUSE_WHEEL:
		ce := (*C.SDL_MouseWheelEvent)(unsafe.Pointer(cEvent))
		return &MouseWheelEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			WindowID:    WindowID(ce.windowID),
			Which:       MouseID(ce.which),
			X:           float32(ce.x),
			Y:           float32(ce.y),
			Direction:   MouseWheelDirection(ce.direction),
		}

	// Joystick axis
	case typ == EVENT_JOYSTICK_AXIS_MOTION:
		ce := (*C.SDL_JoyAxisEvent)(unsafe.Pointer(cEvent))
		return &JoystickAxisEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			Which:       JoystickID(ce.which),
			Axis:        uint8(ce.axis),
			Value:       int16(ce.value),
		}

	// Joystick button
	case typ == EVENT_JOYSTICK_BUTTON_DOWN || typ == EVENT_JOYSTICK_BUTTON_UP:
		ce := (*C.SDL_JoyButtonEvent)(unsafe.Pointer(cEvent))
		return &JoystickButtonEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			Which:       JoystickID(ce.which),
			Button:      uint8(ce.button),
			Down:        bool(ce.down),
		}

	// Joystick hat
	case typ == EVENT_JOYSTICK_HAT_MOTION:
		ce := (*C.SDL_JoyHatEvent)(unsafe.Pointer(cEvent))
		return &JoystickHatEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			Which:       JoystickID(ce.which),
			Hat:         uint8(ce.hat),
			Value:       uint8(ce.value),
		}

	// Gamepad axis
	case typ == EVENT_GAMEPAD_AXIS_MOTION:
		ce := (*C.SDL_GamepadAxisEvent)(unsafe.Pointer(cEvent))
		return &GamepadAxisEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			Which:       JoystickID(ce.which),
			Axis:        int32(ce.axis),
			Value:       int16(ce.value),
		}

	// Gamepad button
	case typ == EVENT_GAMEPAD_BUTTON_DOWN || typ == EVENT_GAMEPAD_BUTTON_UP:
		ce := (*C.SDL_GamepadButtonEvent)(unsafe.Pointer(cEvent))
		return &GamepadButtonEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			Which:       JoystickID(ce.which),
			Button:      int32(ce.button),
			Down:        bool(ce.down),
		}

	// Touch finger events
	case typ == EVENT_FINGER_DOWN || typ == EVENT_FINGER_UP ||
		typ == EVENT_FINGER_MOTION || typ == EVENT_FINGER_CANCELED:
		ce := (*C.SDL_TouchFingerEvent)(unsafe.Pointer(cEvent))
		return &TouchFingerEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			TouchID:     uint64(ce.touchID),
			FingerID:    uint64(ce.fingerID),
			X:           float32(ce.x),
			Y:           float32(ce.y),
			DX:          float32(ce.dx),
			DY:          float32(ce.dy),
			Pressure:    float32(ce.pressure),
			WindowID:    WindowID(ce.windowID),
		}

	// Pen motion
	case typ == EVENT_PEN_MOTION:
		ce := (*C.SDL_PenMotionEvent)(unsafe.Pointer(cEvent))
		return &PenMotionEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			WindowID:    WindowID(ce.windowID),
			Which:       PenID(ce.which),
			X:           float32(ce.x),
			Y:           float32(ce.y),
		}

	// Pen button
	case typ == EVENT_PEN_BUTTON_DOWN || typ == EVENT_PEN_BUTTON_UP:
		ce := (*C.SDL_PenButtonEvent)(unsafe.Pointer(cEvent))
		return &PenButtonEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			WindowID:    WindowID(ce.windowID),
			Which:       PenID(ce.which),
			Button:      uint8(ce.button),
			Down:        bool(ce.down),
		}

	// Sensor update
	case typ == EVENT_SENSOR_UPDATE:
		ce := (*C.SDL_SensorEvent)(unsafe.Pointer(cEvent))
		var data [6]float32
		for i := 0; i < 6; i++ {
			data[i] = float32(ce.data[i])
		}
		return &SensorEvent{
			CommonEvent:     CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			Which:           SensorID(ce.which),
			Data:            data,
			SensorTimestamp: uint64(ce.sensor_timestamp),
		}

	// Audio device events
	case typ == EVENT_AUDIO_DEVICE_ADDED ||
		typ == EVENT_AUDIO_DEVICE_REMOVED ||
		typ == EVENT_AUDIO_DEVICE_FORMAT_CHANGED:
		ce := (*C.SDL_AudioDeviceEvent)(unsafe.Pointer(cEvent))
		return &AudioDeviceEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			Which:       uint32(ce.which),
			Recording:   bool(ce.recording),
		}

	// Drop events
	case typ == EVENT_DROP_FILE || typ == EVENT_DROP_TEXT ||
		typ == EVENT_DROP_BEGIN || typ == EVENT_DROP_COMPLETE ||
		typ == EVENT_DROP_POSITION:
		ce := (*C.SDL_DropEvent)(unsafe.Pointer(cEvent))
		var source, data string
		if ce.source != nil {
			source = C.GoString(ce.source)
		}
		if ce.data != nil {
			data = C.GoString(ce.data)
		}
		return &DropEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			WindowID:    WindowID(ce.windowID),
			X:           float32(ce.x),
			Y:           float32(ce.y),
			Source:      source,
			Data:        data,
		}

	// User events
	case typ >= EVENT_USER && typ <= EVENT_LAST:
		ce := (*C.SDL_UserEvent)(unsafe.Pointer(cEvent))
		return &UserEvent{
			CommonEvent: CommonEvent{Type: EventType(ce._type), Timestamp: uint64(ce.timestamp)},
			WindowID:    WindowID(ce.windowID),
			Code:        int32(ce.code),
			Data1:       ce.data1,
			Data2:       ce.data2,
		}

	case typ == EVENT_JOYSTICK_ADDED || typ == EVENT_JOYSTICK_REMOVED || typ == EVENT_JOYSTICK_UPDATE_COMPLETE:
		e := (*C.SDL_JoyDeviceEvent)(unsafe.Pointer(cEvent))
		return &JoystickDeviceEvent{CommonEvent: common(), Which: JoystickID(e.which)}

	case typ == EVENT_JOYSTICK_BALL_MOTION:
		e := (*C.SDL_JoyBallEvent)(unsafe.Pointer(cEvent))
		return &JoystickBallEvent{CommonEvent: common(), Which: JoystickID(e.which), Ball: uint8(e.ball), XRel: int16(e.xrel), YRel: int16(e.yrel)}

	case typ == EVENT_JOYSTICK_BATTERY_UPDATED:
		e := (*C.SDL_JoyBatteryEvent)(unsafe.Pointer(cEvent))
		return &JoystickBatteryEvent{CommonEvent: common(), Which: JoystickID(e.which), State: PowerState(e.state), Percent: int(e.percent)}

	case typ == EVENT_GAMEPAD_ADDED || typ == EVENT_GAMEPAD_REMOVED || typ == EVENT_GAMEPAD_REMAPPED || typ == EVENT_GAMEPAD_UPDATE_COMPLETE || typ == EVENT_GAMEPAD_STEAM_HANDLE_UPDATED:
		e := (*C.SDL_GamepadDeviceEvent)(unsafe.Pointer(cEvent))
		return &GamepadDeviceEvent{CommonEvent: common(), Which: JoystickID(e.which)}

	case typ >= EVENT_GAMEPAD_TOUCHPAD_DOWN && typ <= EVENT_GAMEPAD_TOUCHPAD_UP:
		e := (*C.SDL_GamepadTouchpadEvent)(unsafe.Pointer(cEvent))
		return &GamepadTouchpadEvent{CommonEvent: common(), Which: JoystickID(e.which), Touchpad: int32(e.touchpad), Finger: int32(e.finger), X: float32(e.x), Y: float32(e.y), Pressure: float32(e.pressure)}

	case typ == EVENT_GAMEPAD_SENSOR_UPDATE:
		e := (*C.SDL_GamepadSensorEvent)(unsafe.Pointer(cEvent))
		return &GamepadSensorEvent{CommonEvent: common(), Which: JoystickID(e.which), Sensor: int32(e.sensor), Data: [3]float32{float32(e.data[0]), float32(e.data[1]), float32(e.data[2])}, SensorTimestamp: uint64(e.sensor_timestamp)}

	case typ == EVENT_MOUSE_ADDED || typ == EVENT_MOUSE_REMOVED:
		e := (*C.SDL_MouseDeviceEvent)(unsafe.Pointer(cEvent))
		return &MouseDeviceEvent{CommonEvent: common(), Which: MouseID(e.which)}

	case typ >= EVENT_CAMERA_DEVICE_ADDED && typ <= EVENT_CAMERA_DEVICE_DENIED:
		e := (*C.SDL_CameraDeviceEvent)(unsafe.Pointer(cEvent))
		return &CameraDeviceEvent{CommonEvent: common(), Which: CameraID(e.which)}

	case typ >= EVENT_RENDER_TARGETS_RESET && typ <= EVENT_RENDER_DEVICE_LOST:
		e := (*C.SDL_RenderEvent)(unsafe.Pointer(cEvent))
		return &RenderEvent{CommonEvent: common(), WindowID: WindowID(e.windowID)}

	case typ == EVENT_CLIPBOARD_UPDATE:
		e := (*C.SDL_ClipboardEvent)(unsafe.Pointer(cEvent))
		return &ClipboardEvent{CommonEvent: common(), Owner: bool(e.owner), NumMimeTypes: int32(e.num_mime_types)}

	case typ == EVENT_PEN_PROXIMITY_IN || typ == EVENT_PEN_PROXIMITY_OUT:
		e := (*C.SDL_PenProximityEvent)(unsafe.Pointer(cEvent))
		return &PenProximityEvent{CommonEvent: common(), WindowID: WindowID(e.windowID), Which: PenID(e.which)}

	case typ == EVENT_PEN_DOWN || typ == EVENT_PEN_UP:
		e := (*C.SDL_PenTouchEvent)(unsafe.Pointer(cEvent))
		return &PenTouchEvent{CommonEvent: common(), WindowID: WindowID(e.windowID), Which: PenID(e.which), PenState: PenInputFlags(e.pen_state), X: float32(e.x), Y: float32(e.y), Eraser: bool(e.eraser), Down: bool(e.down)}

	case typ == EVENT_PEN_AXIS:
		e := (*C.SDL_PenAxisEvent)(unsafe.Pointer(cEvent))
		return &PenAxisEvent{CommonEvent: common(), WindowID: WindowID(e.windowID), Which: PenID(e.which), PenState: PenInputFlags(e.pen_state), X: float32(e.x), Y: float32(e.y), Axis: PenAxis(e.axis), Value: float32(e.value)}

	case typ >= EVENT_PINCH_BEGIN && typ <= EVENT_PINCH_END:
		e := (*C.SDL_PinchFingerEvent)(unsafe.Pointer(cEvent))
		return &PinchFingerEvent{CommonEvent: common(), Scale: float32(e.scale), WindowID: WindowID(e.windowID)}

	case typ == EVENT_TEXT_EDITING_CANDIDATES:
		e := (*C.SDL_TextEditingCandidatesEvent)(unsafe.Pointer(cEvent))
		n := int(e.num_candidates)
		cands := make([]string, n)
		if e.candidates != nil {
			slice := unsafe.Slice(e.candidates, n)
			for i, s := range slice {
				cands[i] = C.GoString(s)
			}
		}
		return &TextEditingCandidatesEvent{CommonEvent: common(), WindowID: WindowID(e.windowID), Candidates: cands, SelectedCandidate: int32(e.selected_candidate), Horizontal: bool(e.horizontal)}

	// Fallback: return a CommonEvent for any unhandled type.
	default:
		return &CommonEvent{Type: typ, Timestamp: common().Timestamp}
	}
}

// ---------------------------------------------------------------------------
// Public event functions
// ---------------------------------------------------------------------------

// PollEvent polls for currently pending events.
// Returns the next event if available, or nil if the queue is empty.
func PollEvent() Event {
	var cEvent C.SDL_Event
	if !bool(C.SDL_PollEvent(&cEvent)) {
		return nil
	}
	return convertEvent(&cEvent)
}

// WaitEvent waits indefinitely for the next available event.
// Returns the event on success, or nil on error.
func WaitEvent() Event {
	var cEvent C.SDL_Event
	if !bool(C.SDL_WaitEvent(&cEvent)) {
		return nil
	}
	return convertEvent(&cEvent)
}

// WaitEventTimeout waits until the specified timeout in milliseconds for an event.
// Returns the event if available, or nil if the timeout elapsed or an error occurred.
func WaitEventTimeout(timeoutMS int32) Event {
	var cEvent C.SDL_Event
	if !bool(C.SDL_WaitEventTimeout(&cEvent, C.Sint32(timeoutMS))) {
		return nil
	}
	return convertEvent(&cEvent)
}

// PushEvent adds a user event to the event queue.
// Returns true on success or false on failure.
func PushEvent(event *UserEvent) bool {
	var cEvent C.SDL_Event
	cUser := (*C.SDL_UserEvent)(unsafe.Pointer(&cEvent))
	cUser._type = C.Uint32(event.Type)
	cUser.timestamp = C.Uint64(event.Timestamp)
	cUser.windowID = C.SDL_WindowID(event.WindowID)
	cUser.code = C.Sint32(event.Code)
	cUser.data1 = event.Data1
	cUser.data2 = event.Data2
	return bool(C.SDL_PushEvent(&cEvent))
}

// PumpEvents gathers events from input devices and places them in the event queue.
func PumpEvents() {
	C.SDL_PumpEvents()
}

// FlushEvent clears events of a specific type from the event queue.
func FlushEvent(typ EventType) {
	C.SDL_FlushEvent(C.Uint32(typ))
}

// FlushEvents clears events of a range of types from the event queue.
func FlushEvents(minType, maxType EventType) {
	C.SDL_FlushEvents(C.Uint32(minType), C.Uint32(maxType))
}

// HasEvent checks for the existence of a certain event type in the event queue.
func HasEvent(typ EventType) bool {
	return bool(C.SDL_HasEvent(C.Uint32(typ)))
}

// HasEvents checks for the existence of certain event types in the event queue.
func HasEvents(minType, maxType EventType) bool {
	return bool(C.SDL_HasEvents(C.Uint32(minType), C.Uint32(maxType)))
}

// SetEventEnabled sets the state of processing events by type.
func SetEventEnabled(typ EventType, enabled bool) {
	C.SDL_SetEventEnabled(C.Uint32(typ), C.bool(enabled))
}

// EventEnabled queries the state of processing events by type.
func EventEnabled(typ EventType) bool {
	return bool(C.SDL_EventEnabled(C.Uint32(typ)))
}

// RegisterEvents allocates a set of user-defined events and returns the
// beginning event number. Returns 0 on failure.
func RegisterEvents(numEvents int) uint32 {
	return uint32(C.SDL_RegisterEvents(C.int(numEvents)))
}

// ---------------------------------------------------------------------------
// EventAction
// ---------------------------------------------------------------------------

// EventAction is the type of action to request from PeepEvents.
type EventAction int

const (
	ADDEVENT  EventAction = C.SDL_ADDEVENT
	PEEKEVENT EventAction = C.SDL_PEEKEVENT
	GETEVENT  EventAction = C.SDL_GETEVENT
)

// PeepEvents checks the event queue for messages and optionally returns them.
// For ADDEVENT, the events slice contains events to add to the queue.
// For PEEKEVENT/GETEVENT, the function returns up to numEvents events matching
// the type range [minType, maxType].
// Returns the number of events actually stored, or an error on failure.
func PeepEvents(events []Event, numEvents int, action EventAction, minType, maxType EventType) (int, error) {
	if action == ADDEVENT {
		cEvents := make([]C.SDL_Event, len(events))
		for i, e := range events {
			if ue, ok := e.(*UserEvent); ok {
				cUser := (*C.SDL_UserEvent)(unsafe.Pointer(&cEvents[i]))
				cUser._type = C.Uint32(ue.Type)
				cUser.timestamp = C.Uint64(ue.Timestamp)
				cUser.windowID = C.SDL_WindowID(ue.WindowID)
				cUser.code = C.Sint32(ue.Code)
				cUser.data1 = ue.Data1
				cUser.data2 = ue.Data2
			}
		}
		n := C.SDL_PeepEvents(&cEvents[0], C.int(len(events)), C.SDL_EventAction(action), C.Uint32(minType), C.Uint32(maxType))
		if n < 0 {
			return 0, getError()
		}
		return int(n), nil
	}
	// PEEKEVENT or GETEVENT
	cEvents := make([]C.SDL_Event, numEvents)
	n := C.SDL_PeepEvents(&cEvents[0], C.int(numEvents), C.SDL_EventAction(action), C.Uint32(minType), C.Uint32(maxType))
	if n < 0 {
		return 0, getError()
	}
	// Replace the events slice contents with converted events
	for i := 0; i < int(n) && i < len(events); i++ {
		events[i] = convertEvent(&cEvents[i])
	}
	// Append if events slice is shorter
	for i := len(events); i < int(n); i++ {
		events = append(events, convertEvent(&cEvents[i]))
	}
	return int(n), nil
}

// PeepEventsSlice is a convenience wrapper around PeepEvents for PEEKEVENT and
// GETEVENT actions. It allocates and returns a slice of events.
func PeepEventsSlice(numEvents int, action EventAction, minType, maxType EventType) ([]Event, error) {
	cEvents := make([]C.SDL_Event, numEvents)
	n := C.SDL_PeepEvents(&cEvents[0], C.int(numEvents), C.SDL_EventAction(action), C.Uint32(minType), C.Uint32(maxType))
	if n < 0 {
		return nil, getError()
	}
	result := make([]Event, int(n))
	for i := 0; i < int(n); i++ {
		result[i] = convertEvent(&cEvents[i])
	}
	return result, nil
}

// GetWindowFromEvent returns the window associated with an event, or nil.
// The caller must pass the raw C event; this function operates on a CommonEvent
// by re-encoding it to a C SDL_Event.
func GetWindowFromEvent(event Event) *Window {
	if event == nil {
		return nil
	}
	// We need the original C SDL_Event. Since we only have the Go Event interface,
	// we construct a minimal C event with just the type set so SDL can look up
	// the window. However, SDL_GetWindowFromEvent needs the full event data.
	// We use PeepEvents or reconstruct from known event types.
	// The simplest correct approach: build a C event from the Go type.
	var cEvent C.SDL_Event
	switch e := event.(type) {
	case *WindowEvent:
		ce := (*C.SDL_WindowEvent)(unsafe.Pointer(&cEvent))
		ce._type = C.SDL_EventType(e.Type)
		ce.timestamp = C.Uint64(e.Timestamp)
		ce.windowID = C.SDL_WindowID(e.WindowID)
		ce.data1 = C.Sint32(e.Data1)
		ce.data2 = C.Sint32(e.Data2)
	case *KeyboardEvent:
		ce := (*C.SDL_KeyboardEvent)(unsafe.Pointer(&cEvent))
		ce._type = C.SDL_EventType(e.Type)
		ce.timestamp = C.Uint64(e.Timestamp)
		ce.windowID = C.SDL_WindowID(e.WindowID)
	case *TextEditingEvent:
		ce := (*C.SDL_TextEditingEvent)(unsafe.Pointer(&cEvent))
		ce._type = C.SDL_EventType(e.Type)
		ce.timestamp = C.Uint64(e.Timestamp)
		ce.windowID = C.SDL_WindowID(e.WindowID)
	case *TextInputEvent:
		ce := (*C.SDL_TextInputEvent)(unsafe.Pointer(&cEvent))
		ce._type = C.SDL_EventType(e.Type)
		ce.timestamp = C.Uint64(e.Timestamp)
		ce.windowID = C.SDL_WindowID(e.WindowID)
	case *MouseMotionEvent:
		ce := (*C.SDL_MouseMotionEvent)(unsafe.Pointer(&cEvent))
		ce._type = C.SDL_EventType(e.Type)
		ce.timestamp = C.Uint64(e.Timestamp)
		ce.windowID = C.SDL_WindowID(e.WindowID)
	case *MouseButtonEvent:
		ce := (*C.SDL_MouseButtonEvent)(unsafe.Pointer(&cEvent))
		ce._type = C.SDL_EventType(e.Type)
		ce.timestamp = C.Uint64(e.Timestamp)
		ce.windowID = C.SDL_WindowID(e.WindowID)
	case *MouseWheelEvent:
		ce := (*C.SDL_MouseWheelEvent)(unsafe.Pointer(&cEvent))
		ce._type = C.SDL_EventType(e.Type)
		ce.timestamp = C.Uint64(e.Timestamp)
		ce.windowID = C.SDL_WindowID(e.WindowID)
	case *DropEvent:
		ce := (*C.SDL_DropEvent)(unsafe.Pointer(&cEvent))
		ce._type = C.SDL_EventType(e.Type)
		ce.timestamp = C.Uint64(e.Timestamp)
		ce.windowID = C.SDL_WindowID(e.WindowID)
	case *UserEvent:
		ce := (*C.SDL_UserEvent)(unsafe.Pointer(&cEvent))
		ce._type = C.Uint32(e.Type)
		ce.timestamp = C.Uint64(e.Timestamp)
		ce.windowID = C.SDL_WindowID(e.WindowID)
	default:
		// For events without a windowID, just set the type.
		ce := (*C.SDL_CommonEvent)(unsafe.Pointer(&cEvent))
		ce._type = C.Uint32(event.GetType())
		ce.timestamp = C.Uint64(event.GetTimestamp())
	}
	w := C.SDL_GetWindowFromEvent(&cEvent)
	if w == nil {
		return nil
	}
	return &Window{c: w}
}

// GetEventDescription returns an English description of an event for logging.
func GetEventDescription(event Event) string {
	if event == nil {
		return ""
	}
	var cEvent C.SDL_Event
	// Set at minimum the type field so SDL can describe it.
	ce := (*C.SDL_CommonEvent)(unsafe.Pointer(&cEvent))
	ce._type = C.Uint32(event.GetType())
	ce.timestamp = C.Uint64(event.GetTimestamp())
	var buf [512]C.char
	C.SDL_GetEventDescription(&cEvent, &buf[0], 512)
	return C.GoString(&buf[0])
}

// EventFilterFunc is a Go function that filters SDL events.
// Return true to keep the event, false to drop it.
type EventFilterFunc func(event Event) bool

//export goEventFilter
func goEventFilter(userdata unsafe.Pointer, cevent *C.SDL_Event) C.bool {
	id := uintptr(userdata)
	fn := getCallback(id).(EventFilterFunc)
	event := convertEvent(cevent)
	return C.bool(fn(event))
}

// SetEventFilter sets a function to filter events before they reach the queue.
// Pass nil to clear the filter.
func SetEventFilter(filter EventFilterFunc) {
	if filter == nil {
		C.SDL_SetEventFilter(nil, nil)
		return
	}
	id := registerCallback(filter)
	C.SDL_SetEventFilter(C._get_event_filter_trampoline(), ptrFromID(id))
}

// GetEventFilter returns whether an event filter is currently set.
func GetEventFilter() bool {
	return bool(C.SDL_GetEventFilter(nil, nil))
}

// AddEventWatch adds a callback triggered when events are added to the queue.
func AddEventWatch(filter EventFilterFunc) uintptr {
	id := registerCallback(filter)
	C.SDL_AddEventWatch(C._get_event_filter_trampoline(), ptrFromID(id))
	return id
}

// RemoveEventWatch removes an event watch callback. Pass the ID returned by AddEventWatch.
func RemoveEventWatch(id uintptr) {
	C.SDL_RemoveEventWatch(C._get_event_filter_trampoline(), ptrFromID(id))
	unregisterCallback(id)
}

// FilterEvents runs a filter function on the current event queue, removing events for which it returns false.
func FilterEvents(filter EventFilterFunc) {
	id := registerCallback(filter)
	C.SDL_FilterEvents(C._get_event_filter_trampoline(), ptrFromID(id))
	unregisterCallback(id)
}

// KeyboardDeviceEvent is generated when a keyboard is added or removed.
type KeyboardDeviceEvent struct {
	CommonEvent
	Which KeyboardID
}
