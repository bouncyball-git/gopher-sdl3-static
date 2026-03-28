package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// MessageBoxFlags specifies the type of message box.
type MessageBoxFlags uint32

const (
	MESSAGEBOX_ERROR       MessageBoxFlags = C.SDL_MESSAGEBOX_ERROR
	MESSAGEBOX_WARNING     MessageBoxFlags = C.SDL_MESSAGEBOX_WARNING
	MESSAGEBOX_INFORMATION MessageBoxFlags = C.SDL_MESSAGEBOX_INFORMATION
)

// ShowSimpleMessageBox displays a simple message box.
func ShowSimpleMessageBox(flags MessageBoxFlags, title, message string, window *Window) error {
	ct := C.CString(title)
	cm := C.CString(message)
	defer C.free(unsafe.Pointer(ct))
	defer C.free(unsafe.Pointer(cm))
	var cw *C.SDL_Window
	if window != nil {
		cw = window.c
	}
	if !C.SDL_ShowSimpleMessageBox(C.SDL_MessageBoxFlags(flags), ct, cm, cw) {
		return getError()
	}
	return nil
}

// Additional MessageBoxFlags constants.
const (
	MESSAGEBOX_BUTTONS_LEFT_TO_RIGHT MessageBoxFlags = C.SDL_MESSAGEBOX_BUTTONS_LEFT_TO_RIGHT
	MESSAGEBOX_BUTTONS_RIGHT_TO_LEFT MessageBoxFlags = C.SDL_MESSAGEBOX_BUTTONS_RIGHT_TO_LEFT
)

// MessageBoxButtonFlags represents flags for message box buttons.
type MessageBoxButtonFlags uint32

const (
	MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT MessageBoxButtonFlags = C.SDL_MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT
	MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT MessageBoxButtonFlags = C.SDL_MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT
)

// MessageBoxButtonData describes a button in a message box.
type MessageBoxButtonData struct {
	Flags    MessageBoxButtonFlags
	ButtonID int
	Text     string
}

// MessageBoxColor represents an RGB color for message box color schemes.
type MessageBoxColor struct {
	R, G, B uint8
}

// MessageBoxColorScheme defines a set of colors for a message box.
type MessageBoxColorScheme struct {
	Background       MessageBoxColor
	Text             MessageBoxColor
	ButtonBorder     MessageBoxColor
	ButtonBackground MessageBoxColor
	ButtonSelected   MessageBoxColor
}

// MessageBoxData describes a message box with title, text, buttons, and colors.
type MessageBoxData struct {
	Flags       MessageBoxFlags
	Window      *Window
	Title       string
	Message     string
	Buttons     []MessageBoxButtonData
	ColorScheme *MessageBoxColorScheme
}

// ShowMessageBox creates a modal message box and returns the button ID that was pressed.
func ShowMessageBox(data *MessageBoxData) (int, error) {
	ct := C.CString(data.Title)
	cm := C.CString(data.Message)
	defer C.free(unsafe.Pointer(ct))
	defer C.free(unsafe.Pointer(cm))

	var cw *C.SDL_Window
	if data.Window != nil {
		cw = data.Window.c
	}

	cbuttons := make([]C.SDL_MessageBoxButtonData, len(data.Buttons))
	cbuttonTexts := make([]*C.char, len(data.Buttons))
	for i, b := range data.Buttons {
		cbuttonTexts[i] = C.CString(b.Text)
		cbuttons[i].flags = C.SDL_MessageBoxButtonFlags(b.Flags)
		cbuttons[i].buttonID = C.int(b.ButtonID)
		cbuttons[i].text = cbuttonTexts[i]
	}
	defer func() {
		for _, t := range cbuttonTexts {
			C.free(unsafe.Pointer(t))
		}
	}()

	var ccs *C.SDL_MessageBoxColorScheme
	var colorScheme C.SDL_MessageBoxColorScheme
	if data.ColorScheme != nil {
		cs := data.ColorScheme
		colorScheme.colors[C.SDL_MESSAGEBOX_COLOR_BACKGROUND] = C.SDL_MessageBoxColor{r: C.Uint8(cs.Background.R), g: C.Uint8(cs.Background.G), b: C.Uint8(cs.Background.B)}
		colorScheme.colors[C.SDL_MESSAGEBOX_COLOR_TEXT] = C.SDL_MessageBoxColor{r: C.Uint8(cs.Text.R), g: C.Uint8(cs.Text.G), b: C.Uint8(cs.Text.B)}
		colorScheme.colors[C.SDL_MESSAGEBOX_COLOR_BUTTON_BORDER] = C.SDL_MessageBoxColor{r: C.Uint8(cs.ButtonBorder.R), g: C.Uint8(cs.ButtonBorder.G), b: C.Uint8(cs.ButtonBorder.B)}
		colorScheme.colors[C.SDL_MESSAGEBOX_COLOR_BUTTON_BACKGROUND] = C.SDL_MessageBoxColor{r: C.Uint8(cs.ButtonBackground.R), g: C.Uint8(cs.ButtonBackground.G), b: C.Uint8(cs.ButtonBackground.B)}
		colorScheme.colors[C.SDL_MESSAGEBOX_COLOR_BUTTON_SELECTED] = C.SDL_MessageBoxColor{r: C.Uint8(cs.ButtonSelected.R), g: C.Uint8(cs.ButtonSelected.G), b: C.Uint8(cs.ButtonSelected.B)}
		ccs = &colorScheme
	}

	var cbd *C.SDL_MessageBoxButtonData
	if len(cbuttons) > 0 {
		cbd = &cbuttons[0]
	}

	cdata := C.SDL_MessageBoxData{
		flags:       C.SDL_MessageBoxFlags(data.Flags),
		window:      cw,
		title:       ct,
		message:     cm,
		numbuttons:  C.int(len(data.Buttons)),
		buttons:     cbd,
		colorScheme: ccs,
	}

	var buttonID C.int
	if !C.SDL_ShowMessageBox(&cdata, &buttonID) {
		return 0, getError()
	}
	return int(buttonID), nil
}

// MessageBoxColorType represents a message box color index.
type MessageBoxColorType int

const (
	MESSAGEBOX_COLOR_BACKGROUND        MessageBoxColorType = C.SDL_MESSAGEBOX_COLOR_BACKGROUND
	MESSAGEBOX_COLOR_TEXT              MessageBoxColorType = C.SDL_MESSAGEBOX_COLOR_TEXT
	MESSAGEBOX_COLOR_BUTTON_BORDER     MessageBoxColorType = C.SDL_MESSAGEBOX_COLOR_BUTTON_BORDER
	MESSAGEBOX_COLOR_BUTTON_BACKGROUND MessageBoxColorType = C.SDL_MESSAGEBOX_COLOR_BUTTON_BACKGROUND
	MESSAGEBOX_COLOR_BUTTON_SELECTED   MessageBoxColorType = C.SDL_MESSAGEBOX_COLOR_BUTTON_SELECTED
)
