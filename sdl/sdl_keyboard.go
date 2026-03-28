package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// KeyboardID is a unique identifier for a keyboard device.
type KeyboardID uint32

// HasKeyboard returns true if a keyboard is connected.
func HasKeyboard() bool {
	return bool(C.SDL_HasKeyboard())
}

// GetKeyboardFocus returns the window which currently has keyboard focus.
func GetKeyboardFocus() *Window {
	cw := C.SDL_GetKeyboardFocus()
	if cw == nil {
		return nil
	}
	return &Window{c: cw}
}

// GetKeyboardState returns the current state of the keyboard.
// The returned slice is indexed by Scancode values.
func GetKeyboardState() []bool {
	var numkeys C.int
	cstate := C.SDL_GetKeyboardState(&numkeys)
	if cstate == nil {
		return nil
	}
	n := int(numkeys)
	result := make([]bool, n)
	slice := unsafe.Slice((*C.bool)(cstate), n)
	for i, v := range slice {
		result[i] = bool(v)
	}
	return result
}

// GetModState returns the current key modifier state.
func GetModState() Keymod {
	return Keymod(C.SDL_GetModState())
}

// SetModState sets the current key modifier state.
func SetModState(modstate Keymod) {
	C.SDL_SetModState(C.SDL_Keymod(modstate))
}

// GetKeyFromScancode returns the keycode corresponding to a scancode with the given modifiers.
func GetKeyFromScancode(scancode Scancode, modstate Keymod, keyEvent bool) Keycode {
	return Keycode(C.SDL_GetKeyFromScancode(C.SDL_Scancode(scancode), C.SDL_Keymod(modstate), C.bool(keyEvent)))
}

// GetScancodeFromKey returns the scancode corresponding to a keycode.
func GetScancodeFromKey(key Keycode) (Scancode, Keymod) {
	var modstate C.SDL_Keymod
	sc := C.SDL_GetScancodeFromKey(C.SDL_Keycode(key), &modstate)
	return Scancode(sc), Keymod(modstate)
}

// GetScancodeName returns a human-readable name for a scancode.
func GetScancodeName(scancode Scancode) string {
	return C.GoString(C.SDL_GetScancodeName(C.SDL_Scancode(scancode)))
}

// GetScancodeFromName returns the scancode for a human-readable name.
func GetScancodeFromName(name string) Scancode {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	return Scancode(C.SDL_GetScancodeFromName(cn))
}

// GetKeyName returns a human-readable name for a keycode.
func GetKeyName(key Keycode) string {
	return C.GoString(C.SDL_GetKeyName(C.SDL_Keycode(key)))
}

// GetKeyFromName returns the keycode for a human-readable name.
func GetKeyFromName(name string) Keycode {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	return Keycode(C.SDL_GetKeyFromName(cn))
}

// StartTextInput starts accepting Unicode text input events for a window.
func StartTextInput(window *Window) error {
	if !C.SDL_StartTextInput(window.c) {
		return getError()
	}
	return nil
}

// StopTextInput stops receiving text input events for a window.
func StopTextInput(window *Window) error {
	if !C.SDL_StopTextInput(window.c) {
		return getError()
	}
	return nil
}

// TextInputActive returns true if Unicode text input events are enabled for a window.
func TextInputActive(window *Window) bool {
	return bool(C.SDL_TextInputActive(window.c))
}

// TextInputType represents the type of text being input.
type TextInputType int

const (
	TEXTINPUT_TYPE_TEXT                    TextInputType = C.SDL_TEXTINPUT_TYPE_TEXT
	TEXTINPUT_TYPE_TEXT_NAME               TextInputType = C.SDL_TEXTINPUT_TYPE_TEXT_NAME
	TEXTINPUT_TYPE_TEXT_EMAIL              TextInputType = C.SDL_TEXTINPUT_TYPE_TEXT_EMAIL
	TEXTINPUT_TYPE_TEXT_USERNAME           TextInputType = C.SDL_TEXTINPUT_TYPE_TEXT_USERNAME
	TEXTINPUT_TYPE_TEXT_PASSWORD_HIDDEN    TextInputType = C.SDL_TEXTINPUT_TYPE_TEXT_PASSWORD_HIDDEN
	TEXTINPUT_TYPE_TEXT_PASSWORD_VISIBLE   TextInputType = C.SDL_TEXTINPUT_TYPE_TEXT_PASSWORD_VISIBLE
	TEXTINPUT_TYPE_NUMBER                  TextInputType = C.SDL_TEXTINPUT_TYPE_NUMBER
	TEXTINPUT_TYPE_NUMBER_PASSWORD_HIDDEN  TextInputType = C.SDL_TEXTINPUT_TYPE_NUMBER_PASSWORD_HIDDEN
	TEXTINPUT_TYPE_NUMBER_PASSWORD_VISIBLE TextInputType = C.SDL_TEXTINPUT_TYPE_NUMBER_PASSWORD_VISIBLE
)

// Capitalization represents auto-capitalization type.
type Capitalization int

const (
	CAPITALIZE_NONE      Capitalization = C.SDL_CAPITALIZE_NONE
	CAPITALIZE_SENTENCES Capitalization = C.SDL_CAPITALIZE_SENTENCES
	CAPITALIZE_WORDS     Capitalization = C.SDL_CAPITALIZE_WORDS
	CAPITALIZE_LETTERS   Capitalization = C.SDL_CAPITALIZE_LETTERS
)

// GetKeyboards returns a list of currently connected keyboard IDs.
func GetKeyboards() []KeyboardID {
	var count C.int
	cids := C.SDL_GetKeyboards(&count)
	if cids == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cids))
	n := int(count)
	result := make([]KeyboardID, n)
	slice := unsafe.Slice((*C.SDL_KeyboardID)(cids), n)
	for i, id := range slice {
		result[i] = KeyboardID(id)
	}
	return result
}

// GetKeyboardNameForID returns the name of a keyboard.
func GetKeyboardNameForID(instanceID KeyboardID) string {
	return C.GoString(C.SDL_GetKeyboardNameForID(C.SDL_KeyboardID(instanceID)))
}

// ResetKeyboard clears the state of the keyboard, generating key up events for all pressed keys.
func ResetKeyboard() {
	C.SDL_ResetKeyboard()
}

// SetScancodeName sets a human-readable name for a scancode.
func SetScancodeName(scancode Scancode, name string) error {
	cn := C.CString(name)
	// Note: SDL docs say the string must stay valid while SDL is being used,
	// so we intentionally do NOT free cn here.
	if !C.SDL_SetScancodeName(C.SDL_Scancode(scancode), cn) {
		C.free(unsafe.Pointer(cn))
		return getError()
	}
	return nil
}

// HasScreenKeyboardSupport returns true if the platform has screen keyboard support.
func HasScreenKeyboardSupport() bool {
	return bool(C.SDL_HasScreenKeyboardSupport())
}

// ScreenKeyboardShown returns true if the screen keyboard is shown for the given window.
func ScreenKeyboardShown(window *Window) bool {
	return bool(C.SDL_ScreenKeyboardShown(window.c))
}

// ClearComposition dismisses the composition window/IME without disabling the subsystem.
func ClearComposition(window *Window) error {
	if !C.SDL_ClearComposition(window.c) {
		return getError()
	}
	return nil
}

// SetTextInputArea sets the area used to type Unicode text input.
func SetTextInputArea(window *Window, rect *Rect, cursor int) error {
	var cr *C.SDL_Rect
	if rect != nil {
		cr = (*C.SDL_Rect)(unsafe.Pointer(rect))
	}
	if !C.SDL_SetTextInputArea(window.c, cr, C.int(cursor)) {
		return getError()
	}
	return nil
}

// GetTextInputArea returns the area used to type Unicode text input.
func GetTextInputArea(window *Window) (Rect, int, error) {
	var rect Rect
	var cursor C.int
	if !C.SDL_GetTextInputArea(window.c, (*C.SDL_Rect)(unsafe.Pointer(&rect)), &cursor) {
		return rect, 0, getError()
	}
	return rect, int(cursor), nil
}

// StartTextInputWithProperties starts accepting Unicode text input events with properties.
func StartTextInputWithProperties(window *Window, props PropertiesID) error {
	if !C.SDL_StartTextInputWithProperties(window.c, C.SDL_PropertiesID(props)) {
		return getError()
	}
	return nil
}
