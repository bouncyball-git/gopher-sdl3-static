package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// Gamepad represents an opened SDL gamepad.
type Gamepad struct {
	c *C.SDL_Gamepad
}

// GamepadType identifies the type of a gamepad.
type GamepadType int32

// GamepadType constants.
const (
	GAMEPAD_TYPE_UNKNOWN                       GamepadType = C.SDL_GAMEPAD_TYPE_UNKNOWN
	GAMEPAD_TYPE_STANDARD                      GamepadType = C.SDL_GAMEPAD_TYPE_STANDARD
	GAMEPAD_TYPE_XBOX360                       GamepadType = C.SDL_GAMEPAD_TYPE_XBOX360
	GAMEPAD_TYPE_XBOXONE                       GamepadType = C.SDL_GAMEPAD_TYPE_XBOXONE
	GAMEPAD_TYPE_PS3                           GamepadType = C.SDL_GAMEPAD_TYPE_PS3
	GAMEPAD_TYPE_PS4                           GamepadType = C.SDL_GAMEPAD_TYPE_PS4
	GAMEPAD_TYPE_PS5                           GamepadType = C.SDL_GAMEPAD_TYPE_PS5
	GAMEPAD_TYPE_NINTENDO_SWITCH_PRO           GamepadType = C.SDL_GAMEPAD_TYPE_NINTENDO_SWITCH_PRO
	GAMEPAD_TYPE_NINTENDO_SWITCH_JOYCON_LEFT   GamepadType = C.SDL_GAMEPAD_TYPE_NINTENDO_SWITCH_JOYCON_LEFT
	GAMEPAD_TYPE_NINTENDO_SWITCH_JOYCON_RIGHT  GamepadType = C.SDL_GAMEPAD_TYPE_NINTENDO_SWITCH_JOYCON_RIGHT
	GAMEPAD_TYPE_NINTENDO_SWITCH_JOYCON_PAIR   GamepadType = C.SDL_GAMEPAD_TYPE_NINTENDO_SWITCH_JOYCON_PAIR
	GAMEPAD_TYPE_GAMECUBE                      GamepadType = C.SDL_GAMEPAD_TYPE_GAMECUBE
	GAMEPAD_TYPE_COUNT                         GamepadType = C.SDL_GAMEPAD_TYPE_COUNT
)

// GamepadButton identifies a button on a gamepad.
type GamepadButton int32

// GamepadButton constants.
const (
	GAMEPAD_BUTTON_INVALID        GamepadButton = C.SDL_GAMEPAD_BUTTON_INVALID
	GAMEPAD_BUTTON_SOUTH          GamepadButton = C.SDL_GAMEPAD_BUTTON_SOUTH
	GAMEPAD_BUTTON_EAST           GamepadButton = C.SDL_GAMEPAD_BUTTON_EAST
	GAMEPAD_BUTTON_WEST           GamepadButton = C.SDL_GAMEPAD_BUTTON_WEST
	GAMEPAD_BUTTON_NORTH          GamepadButton = C.SDL_GAMEPAD_BUTTON_NORTH
	GAMEPAD_BUTTON_BACK           GamepadButton = C.SDL_GAMEPAD_BUTTON_BACK
	GAMEPAD_BUTTON_GUIDE          GamepadButton = C.SDL_GAMEPAD_BUTTON_GUIDE
	GAMEPAD_BUTTON_START          GamepadButton = C.SDL_GAMEPAD_BUTTON_START
	GAMEPAD_BUTTON_LEFT_STICK     GamepadButton = C.SDL_GAMEPAD_BUTTON_LEFT_STICK
	GAMEPAD_BUTTON_RIGHT_STICK    GamepadButton = C.SDL_GAMEPAD_BUTTON_RIGHT_STICK
	GAMEPAD_BUTTON_LEFT_SHOULDER  GamepadButton = C.SDL_GAMEPAD_BUTTON_LEFT_SHOULDER
	GAMEPAD_BUTTON_RIGHT_SHOULDER GamepadButton = C.SDL_GAMEPAD_BUTTON_RIGHT_SHOULDER
	GAMEPAD_BUTTON_DPAD_UP        GamepadButton = C.SDL_GAMEPAD_BUTTON_DPAD_UP
	GAMEPAD_BUTTON_DPAD_DOWN      GamepadButton = C.SDL_GAMEPAD_BUTTON_DPAD_DOWN
	GAMEPAD_BUTTON_DPAD_LEFT      GamepadButton = C.SDL_GAMEPAD_BUTTON_DPAD_LEFT
	GAMEPAD_BUTTON_DPAD_RIGHT     GamepadButton = C.SDL_GAMEPAD_BUTTON_DPAD_RIGHT
	GAMEPAD_BUTTON_MISC1          GamepadButton = C.SDL_GAMEPAD_BUTTON_MISC1
	GAMEPAD_BUTTON_RIGHT_PADDLE1  GamepadButton = C.SDL_GAMEPAD_BUTTON_RIGHT_PADDLE1
	GAMEPAD_BUTTON_LEFT_PADDLE1   GamepadButton = C.SDL_GAMEPAD_BUTTON_LEFT_PADDLE1
	GAMEPAD_BUTTON_RIGHT_PADDLE2  GamepadButton = C.SDL_GAMEPAD_BUTTON_RIGHT_PADDLE2
	GAMEPAD_BUTTON_LEFT_PADDLE2   GamepadButton = C.SDL_GAMEPAD_BUTTON_LEFT_PADDLE2
	GAMEPAD_BUTTON_TOUCHPAD       GamepadButton = C.SDL_GAMEPAD_BUTTON_TOUCHPAD
	GAMEPAD_BUTTON_MISC2          GamepadButton = C.SDL_GAMEPAD_BUTTON_MISC2
	GAMEPAD_BUTTON_MISC3          GamepadButton = C.SDL_GAMEPAD_BUTTON_MISC3
	GAMEPAD_BUTTON_MISC4          GamepadButton = C.SDL_GAMEPAD_BUTTON_MISC4
	GAMEPAD_BUTTON_MISC5          GamepadButton = C.SDL_GAMEPAD_BUTTON_MISC5
	GAMEPAD_BUTTON_MISC6          GamepadButton = C.SDL_GAMEPAD_BUTTON_MISC6
	GAMEPAD_BUTTON_COUNT          GamepadButton = C.SDL_GAMEPAD_BUTTON_COUNT
)

// GamepadAxis identifies an axis on a gamepad.
type GamepadAxis int32

// GamepadAxis constants.
const (
	GAMEPAD_AXIS_INVALID       GamepadAxis = C.SDL_GAMEPAD_AXIS_INVALID
	GAMEPAD_AXIS_LEFTX         GamepadAxis = C.SDL_GAMEPAD_AXIS_LEFTX
	GAMEPAD_AXIS_LEFTY         GamepadAxis = C.SDL_GAMEPAD_AXIS_LEFTY
	GAMEPAD_AXIS_RIGHTX        GamepadAxis = C.SDL_GAMEPAD_AXIS_RIGHTX
	GAMEPAD_AXIS_RIGHTY        GamepadAxis = C.SDL_GAMEPAD_AXIS_RIGHTY
	GAMEPAD_AXIS_LEFT_TRIGGER  GamepadAxis = C.SDL_GAMEPAD_AXIS_LEFT_TRIGGER
	GAMEPAD_AXIS_RIGHT_TRIGGER GamepadAxis = C.SDL_GAMEPAD_AXIS_RIGHT_TRIGGER
	GAMEPAD_AXIS_COUNT         GamepadAxis = C.SDL_GAMEPAD_AXIS_COUNT
)

// GamepadButtonLabel identifies the label on a gamepad face button.
type GamepadButtonLabel int32

// GamepadButtonLabel constants.
const (
	GAMEPAD_BUTTON_LABEL_UNKNOWN  GamepadButtonLabel = C.SDL_GAMEPAD_BUTTON_LABEL_UNKNOWN
	GAMEPAD_BUTTON_LABEL_A        GamepadButtonLabel = C.SDL_GAMEPAD_BUTTON_LABEL_A
	GAMEPAD_BUTTON_LABEL_B        GamepadButtonLabel = C.SDL_GAMEPAD_BUTTON_LABEL_B
	GAMEPAD_BUTTON_LABEL_X        GamepadButtonLabel = C.SDL_GAMEPAD_BUTTON_LABEL_X
	GAMEPAD_BUTTON_LABEL_Y        GamepadButtonLabel = C.SDL_GAMEPAD_BUTTON_LABEL_Y
	GAMEPAD_BUTTON_LABEL_CROSS    GamepadButtonLabel = C.SDL_GAMEPAD_BUTTON_LABEL_CROSS
	GAMEPAD_BUTTON_LABEL_CIRCLE   GamepadButtonLabel = C.SDL_GAMEPAD_BUTTON_LABEL_CIRCLE
	GAMEPAD_BUTTON_LABEL_SQUARE   GamepadButtonLabel = C.SDL_GAMEPAD_BUTTON_LABEL_SQUARE
	GAMEPAD_BUTTON_LABEL_TRIANGLE GamepadButtonLabel = C.SDL_GAMEPAD_BUTTON_LABEL_TRIANGLE
)

// --- Mapping functions ---

// AddGamepadMapping adds or updates a gamepad mapping.
// Returns 1 if a new mapping is added, 0 if an existing mapping is updated, -1 on failure.
func AddGamepadMapping(mapping string) (int, error) {
	cmapping := C.CString(mapping)
	defer C.free(unsafe.Pointer(cmapping))
	rc := int(C.SDL_AddGamepadMapping(cmapping))
	if rc < 0 {
		return rc, getError()
	}
	return rc, nil
}

// AddGamepadMappingsFromFile loads gamepad mappings from a file.
// Returns the number of mappings added or -1 on failure.
func AddGamepadMappingsFromFile(file string) (int, error) {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))
	rc := int(C.SDL_AddGamepadMappingsFromFile(cfile))
	if rc < 0 {
		return rc, getError()
	}
	return rc, nil
}

// ReloadGamepadMappings reinitializes the SDL mapping database to its initial state.
func ReloadGamepadMappings() error {
	if !C.SDL_ReloadGamepadMappings() {
		return getError()
	}
	return nil
}

// GetGamepadMappings returns the current gamepad mappings.
func GetGamepadMappings() []string {
	var count C.int
	cmappings := C.SDL_GetGamepadMappings(&count)
	if cmappings == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cmappings))
	n := int(count)
	result := make([]string, n)
	slice := unsafe.Slice((**C.char)(unsafe.Pointer(cmappings)), n)
	for i, s := range slice {
		result[i] = C.GoString(s)
	}
	return result
}

// GetGamepadMapping returns the current mapping string of an opened gamepad.
// The caller does not need to free the returned string.
func GetGamepadMapping(gamepad *Gamepad) string {
	cmapping := C.SDL_GetGamepadMapping(gamepad.c)
	if cmapping == nil {
		return ""
	}
	defer C.SDL_free(unsafe.Pointer(cmapping))
	return C.GoString(cmapping)
}

// SetGamepadMapping sets the mapping for a joystick or gamepad by instance ID.
func SetGamepadMapping(instanceID JoystickID, mapping string) error {
	cmapping := C.CString(mapping)
	defer C.free(unsafe.Pointer(cmapping))
	if !C.SDL_SetGamepadMapping(C.SDL_JoystickID(instanceID), cmapping) {
		return getError()
	}
	return nil
}

// --- Enumeration ---

// HasGamepad returns true if a gamepad is currently connected.
func HasGamepad() bool {
	return bool(C.SDL_HasGamepad())
}

// GetGamepads returns a list of currently connected gamepad instance IDs.
func GetGamepads() []JoystickID {
	var count C.int
	cids := C.SDL_GetGamepads(&count)
	if cids == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cids))
	n := int(count)
	result := make([]JoystickID, n)
	slice := unsafe.Slice((*C.SDL_JoystickID)(cids), n)
	for i, id := range slice {
		result[i] = JoystickID(id)
	}
	return result
}

// IsGamepad returns true if the given joystick is supported by the gamepad interface.
func IsGamepad(instanceID JoystickID) bool {
	return bool(C.SDL_IsGamepad(C.SDL_JoystickID(instanceID)))
}

// --- Open / lookup ---

// OpenGamepad opens a gamepad for use.
func OpenGamepad(instanceID JoystickID) (*Gamepad, error) {
	c := C.SDL_OpenGamepad(C.SDL_JoystickID(instanceID))
	if c == nil {
		return nil, getError()
	}
	return &Gamepad{c: c}, nil
}

// GetGamepadFromID returns the Gamepad associated with a joystick instance ID, if it has been opened.
func GetGamepadFromID(instanceID JoystickID) (*Gamepad, error) {
	c := C.SDL_GetGamepadFromID(C.SDL_JoystickID(instanceID))
	if c == nil {
		return nil, getError()
	}
	return &Gamepad{c: c}, nil
}

// GetGamepadFromPlayerIndex returns the Gamepad associated with a player index.
func GetGamepadFromPlayerIndex(playerIndex int) *Gamepad {
	c := C.SDL_GetGamepadFromPlayerIndex(C.int(playerIndex))
	if c == nil {
		return nil
	}
	return &Gamepad{c: c}
}

// --- Gamepad properties ---

// Properties returns the property ID for this gamepad.
func (g *Gamepad) Properties() PropertiesID {
	return PropertiesID(C.SDL_GetGamepadProperties(g.c))
}

// --- Gamepad instance info ---

// Close closes the gamepad.
func (g *Gamepad) Close() {
	if g.c != nil {
		C.SDL_CloseGamepad(g.c)
		g.c = nil
	}
}

// ID returns the instance ID of the gamepad.
func (g *Gamepad) ID() JoystickID {
	return JoystickID(C.SDL_GetGamepadID(g.c))
}

// Name returns the implementation-dependent name of the gamepad.
func (g *Gamepad) Name() string {
	return C.GoString(C.SDL_GetGamepadName(g.c))
}

// Path returns the implementation-dependent path of the gamepad.
func (g *Gamepad) Path() string {
	return C.GoString(C.SDL_GetGamepadPath(g.c))
}

// Type returns the type of the gamepad.
func (g *Gamepad) Type() GamepadType {
	return GamepadType(C.SDL_GetGamepadType(g.c))
}

// RealType returns the type of the gamepad, ignoring any mapping override.
func (g *Gamepad) RealType() GamepadType {
	return GamepadType(C.SDL_GetRealGamepadType(g.c))
}

// PlayerIndex returns the player index of the gamepad, or -1 if not available.
func (g *Gamepad) PlayerIndex() int {
	return int(C.SDL_GetGamepadPlayerIndex(g.c))
}

// SetPlayerIndex sets the player index of the gamepad.
func (g *Gamepad) SetPlayerIndex(playerIndex int) error {
	if !C.SDL_SetGamepadPlayerIndex(g.c, C.int(playerIndex)) {
		return getError()
	}
	return nil
}

// Vendor returns the USB vendor ID of the gamepad, or 0 if unavailable.
func (g *Gamepad) Vendor() uint16 {
	return uint16(C.SDL_GetGamepadVendor(g.c))
}

// Product returns the USB product ID of the gamepad, or 0 if unavailable.
func (g *Gamepad) Product() uint16 {
	return uint16(C.SDL_GetGamepadProduct(g.c))
}

// ProductVersion returns the product version of the gamepad, or 0 if unavailable.
func (g *Gamepad) ProductVersion() uint16 {
	return uint16(C.SDL_GetGamepadProductVersion(g.c))
}

// FirmwareVersion returns the firmware version of the gamepad, or 0 if unavailable.
func (g *Gamepad) FirmwareVersion() uint16 {
	return uint16(C.SDL_GetGamepadFirmwareVersion(g.c))
}

// Serial returns the serial number of the gamepad, or an empty string if unavailable.
func (g *Gamepad) Serial() string {
	return C.GoString(C.SDL_GetGamepadSerial(g.c))
}

// SteamHandle returns the Steam Input handle of the gamepad, or 0 if unavailable.
func (g *Gamepad) SteamHandle() uint64 {
	return uint64(C.SDL_GetGamepadSteamHandle(g.c))
}

// ConnectionState returns the connection state of the gamepad.
func (g *Gamepad) ConnectionState() JoystickConnectionState {
	return JoystickConnectionState(C.SDL_GetGamepadConnectionState(g.c))
}

// PowerInfo returns the battery state and percentage of the gamepad.
// Percent is between 0 and 100, or -1 if not determinable.
func (g *Gamepad) PowerInfo() (PowerState, int) {
	var percent C.int
	state := C.SDL_GetGamepadPowerInfo(g.c, &percent)
	return PowerState(state), int(percent)
}

// Connected returns true if the gamepad is currently connected.
func (g *Gamepad) Connected() bool {
	return bool(C.SDL_GamepadConnected(g.c))
}

// --- Joystick access ---

// GetGamepadJoystick returns the underlying Joystick for the gamepad.
// The returned Joystick is owned by the Gamepad; do not close it.
func GetGamepadJoystick(gamepad *Gamepad) *Joystick {
	c := C.SDL_GetGamepadJoystick(gamepad.c)
	if c == nil {
		return nil
	}
	return &Joystick{c: c}
}

// --- Events ---

// SetGamepadEventsEnabled sets whether gamepad event processing is enabled.
func SetGamepadEventsEnabled(enabled bool) {
	C.SDL_SetGamepadEventsEnabled(C.bool(enabled))
}

// GamepadEventsEnabled returns true if gamepad events are being processed.
func GamepadEventsEnabled() bool {
	return bool(C.SDL_GamepadEventsEnabled())
}

// UpdateGamepads manually pumps gamepad updates.
func UpdateGamepads() {
	C.SDL_UpdateGamepads()
}

// --- String conversion ---

// GetGamepadTypeFromString converts a string to a GamepadType.
func GetGamepadTypeFromString(str string) GamepadType {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return GamepadType(C.SDL_GetGamepadTypeFromString(cstr))
}

// GetGamepadStringForType converts a GamepadType to a string.
func GetGamepadStringForType(t GamepadType) string {
	return C.GoString(C.SDL_GetGamepadStringForType(C.SDL_GamepadType(t)))
}

// GetGamepadAxisFromString converts a string to a GamepadAxis.
func GetGamepadAxisFromString(str string) GamepadAxis {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return GamepadAxis(C.SDL_GetGamepadAxisFromString(cstr))
}

// GetGamepadStringForAxis converts a GamepadAxis to a string.
func GetGamepadStringForAxis(axis GamepadAxis) string {
	return C.GoString(C.SDL_GetGamepadStringForAxis(C.SDL_GamepadAxis(axis)))
}

// GetGamepadButtonFromString converts a string to a GamepadButton.
func GetGamepadButtonFromString(str string) GamepadButton {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return GamepadButton(C.SDL_GetGamepadButtonFromString(cstr))
}

// GetGamepadStringForButton converts a GamepadButton to a string.
func GetGamepadStringForButton(button GamepadButton) string {
	return C.GoString(C.SDL_GetGamepadStringForButton(C.SDL_GamepadButton(button)))
}

// --- Axis and button state ---

// GamepadHasAxis returns true if the gamepad has the given axis.
func GamepadHasAxis(gamepad *Gamepad, axis GamepadAxis) bool {
	return bool(C.SDL_GamepadHasAxis(gamepad.c, C.SDL_GamepadAxis(axis)))
}

// GetGamepadAxis returns the current state of an axis on the gamepad.
func GetGamepadAxis(gamepad *Gamepad, axis GamepadAxis) int16 {
	return int16(C.SDL_GetGamepadAxis(gamepad.c, C.SDL_GamepadAxis(axis)))
}

// GamepadHasButton returns true if the gamepad has the given button.
func GamepadHasButton(gamepad *Gamepad, button GamepadButton) bool {
	return bool(C.SDL_GamepadHasButton(gamepad.c, C.SDL_GamepadButton(button)))
}

// GetGamepadButton returns true if the button is pressed on the gamepad.
func GetGamepadButton(gamepad *Gamepad, button GamepadButton) bool {
	return bool(C.SDL_GetGamepadButton(gamepad.c, C.SDL_GamepadButton(button)))
}

// GetGamepadButtonLabel returns the label of a button on the gamepad.
func GetGamepadButtonLabel(gamepad *Gamepad, button GamepadButton) GamepadButtonLabel {
	return GamepadButtonLabel(C.SDL_GetGamepadButtonLabel(gamepad.c, C.SDL_GamepadButton(button)))
}

// GetGamepadButtonLabelForType returns the label of a button for a given gamepad type.
func GetGamepadButtonLabelForType(t GamepadType, button GamepadButton) GamepadButtonLabel {
	return GamepadButtonLabel(C.SDL_GetGamepadButtonLabelForType(C.SDL_GamepadType(t), C.SDL_GamepadButton(button)))
}

// --- Touchpad ---

// GetNumGamepadTouchpads returns the number of touchpads on the gamepad.
func GetNumGamepadTouchpads(gamepad *Gamepad) int {
	return int(C.SDL_GetNumGamepadTouchpads(gamepad.c))
}

// GetNumGamepadTouchpadFingers returns the number of supported simultaneous fingers on a touchpad.
func GetNumGamepadTouchpadFingers(gamepad *Gamepad, touchpad int) int {
	return int(C.SDL_GetNumGamepadTouchpadFingers(gamepad.c, C.int(touchpad)))
}

// GetGamepadTouchpadFinger returns the state of a finger on a touchpad.
func GetGamepadTouchpadFinger(gamepad *Gamepad, touchpad, finger int) (down bool, x, y, pressure float32, err error) {
	var cdown C.bool
	var cx, cy, cpressure C.float
	if !C.SDL_GetGamepadTouchpadFinger(gamepad.c, C.int(touchpad), C.int(finger), &cdown, &cx, &cy, &cpressure) {
		return false, 0, 0, 0, getError()
	}
	return bool(cdown), float32(cx), float32(cy), float32(cpressure), nil
}

// --- Sensor ---

// GamepadHasSensor returns true if the gamepad has the given sensor.
func GamepadHasSensor(gamepad *Gamepad, sensorType SensorType) bool {
	return bool(C.SDL_GamepadHasSensor(gamepad.c, C.SDL_SensorType(sensorType)))
}

// SetGamepadSensorEnabled enables or disables data reporting for a gamepad sensor.
func SetGamepadSensorEnabled(gamepad *Gamepad, sensorType SensorType, enabled bool) error {
	if !C.SDL_SetGamepadSensorEnabled(gamepad.c, C.SDL_SensorType(sensorType), C.bool(enabled)) {
		return getError()
	}
	return nil
}

// GamepadSensorEnabled returns true if the given sensor is enabled on the gamepad.
func GamepadSensorEnabled(gamepad *Gamepad, sensorType SensorType) bool {
	return bool(C.SDL_GamepadSensorEnabled(gamepad.c, C.SDL_SensorType(sensorType)))
}

// GetGamepadSensorDataRate returns the data rate of a gamepad sensor in events per second.
func GetGamepadSensorDataRate(gamepad *Gamepad, sensorType SensorType) float32 {
	return float32(C.SDL_GetGamepadSensorDataRate(gamepad.c, C.SDL_SensorType(sensorType)))
}

// GetGamepadSensorData returns the current state of a gamepad sensor.
func GetGamepadSensorData(gamepad *Gamepad, sensorType SensorType, numValues int) ([]float32, error) {
	data := make([]float32, numValues)
	if !C.SDL_GetGamepadSensorData(gamepad.c, C.SDL_SensorType(sensorType), (*C.float)(unsafe.Pointer(&data[0])), C.int(numValues)) {
		return nil, getError()
	}
	return data, nil
}

// --- Rumble and effects ---

// RumbleGamepad starts a rumble effect on the gamepad.
func RumbleGamepad(gamepad *Gamepad, lowFrequencyRumble, highFrequencyRumble uint16, durationMS uint32) error {
	if !C.SDL_RumbleGamepad(gamepad.c, C.Uint16(lowFrequencyRumble), C.Uint16(highFrequencyRumble), C.Uint32(durationMS)) {
		return getError()
	}
	return nil
}

// RumbleGamepadTriggers starts a rumble effect in the gamepad's triggers.
func RumbleGamepadTriggers(gamepad *Gamepad, leftRumble, rightRumble uint16, durationMS uint32) error {
	if !C.SDL_RumbleGamepadTriggers(gamepad.c, C.Uint16(leftRumble), C.Uint16(rightRumble), C.Uint32(durationMS)) {
		return getError()
	}
	return nil
}

// SetGamepadLED updates the gamepad's LED color.
func SetGamepadLED(gamepad *Gamepad, red, green, blue uint8) error {
	if !C.SDL_SetGamepadLED(gamepad.c, C.Uint8(red), C.Uint8(green), C.Uint8(blue)) {
		return getError()
	}
	return nil
}

// SendGamepadEffect sends a gamepad-specific effect packet.
func SendGamepadEffect(gamepad *Gamepad, data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if !C.SDL_SendGamepadEffect(gamepad.c, unsafe.Pointer(&data[0]), C.int(len(data))) {
		return getError()
	}
	return nil
}

// GetGamepadNameForID returns the name of a gamepad by instance ID.
func GetGamepadNameForID(instanceID JoystickID) string {
	return C.GoString(C.SDL_GetGamepadNameForID(C.SDL_JoystickID(instanceID)))
}

// GetGamepadPathForID returns the path of a gamepad by instance ID.
func GetGamepadPathForID(instanceID JoystickID) string {
	return C.GoString(C.SDL_GetGamepadPathForID(C.SDL_JoystickID(instanceID)))
}

// GetGamepadPlayerIndexForID returns the player index of a gamepad by instance ID.
func GetGamepadPlayerIndexForID(instanceID JoystickID) int {
	return int(C.SDL_GetGamepadPlayerIndexForID(C.SDL_JoystickID(instanceID)))
}

// GetGamepadGUIDForID returns the GUID of a gamepad by instance ID.
func GetGamepadGUIDForID(instanceID JoystickID) GUID {
	return guidFromC(C.SDL_GetGamepadGUIDForID(C.SDL_JoystickID(instanceID)))
}

// GetGamepadVendorForID returns the USB vendor ID of a gamepad by instance ID.
func GetGamepadVendorForID(instanceID JoystickID) uint16 {
	return uint16(C.SDL_GetGamepadVendorForID(C.SDL_JoystickID(instanceID)))
}

// GetGamepadProductForID returns the USB product ID of a gamepad by instance ID.
func GetGamepadProductForID(instanceID JoystickID) uint16 {
	return uint16(C.SDL_GetGamepadProductForID(C.SDL_JoystickID(instanceID)))
}

// GetGamepadProductVersionForID returns the product version of a gamepad by instance ID.
func GetGamepadProductVersionForID(instanceID JoystickID) uint16 {
	return uint16(C.SDL_GetGamepadProductVersionForID(C.SDL_JoystickID(instanceID)))
}

// GetGamepadTypeForID returns the type of a gamepad by instance ID.
func GetGamepadTypeForID(instanceID JoystickID) GamepadType {
	return GamepadType(C.SDL_GetGamepadTypeForID(C.SDL_JoystickID(instanceID)))
}

// GetRealGamepadTypeForID returns the real (physical) type of a gamepad by instance ID.
func GetRealGamepadTypeForID(instanceID JoystickID) GamepadType {
	return GamepadType(C.SDL_GetRealGamepadTypeForID(C.SDL_JoystickID(instanceID)))
}

// GetGamepadMappingForID returns the mapping string of a gamepad by instance ID.
// The caller receives a copy; the C string is freed.
func GetGamepadMappingForID(instanceID JoystickID) string {
	cs := C.SDL_GetGamepadMappingForID(C.SDL_JoystickID(instanceID))
	if cs == nil {
		return ""
	}
	s := C.GoString(cs)
	C.SDL_free(unsafe.Pointer(cs))
	return s
}

// GetGamepadMappingForGUID returns the mapping string for a gamepad GUID.
func GetGamepadMappingForGUID(guid GUID) string {
	cs := C.SDL_GetGamepadMappingForGUID(guid.cval())
	if cs == nil {
		return ""
	}
	s := C.GoString(cs)
	C.SDL_free(unsafe.Pointer(cs))
	return s
}

// AddGamepadMappingsFromIO loads gamepad mappings from an I/O stream.
func AddGamepadMappingsFromIO(src *IOStream, closeio bool) (int, error) {
	n := C.SDL_AddGamepadMappingsFromIO(src.c, C.bool(closeio))
	if n < 0 {
		return 0, getError()
	}
	return int(n), nil
}

// GetGamepadBindings returns the current bindings for a gamepad.
// Returns a slice of raw C binding pointers as unsafe.Pointer.
func GetGamepadBindings(gamepad *Gamepad) ([]unsafe.Pointer, error) {
	var count C.int
	cb := C.SDL_GetGamepadBindings(gamepad.c, &count)
	if cb == nil {
		return nil, getError()
	}
	defer C.SDL_free(unsafe.Pointer(cb))
	n := int(count)
	result := make([]unsafe.Pointer, n)
	slice := unsafe.Slice((**C.SDL_GamepadBinding)(unsafe.Pointer(cb)), n)
	for i, b := range slice {
		result[i] = unsafe.Pointer(b)
	}
	return result, nil
}

// GetGamepadAppleSFSymbolsNameForButton returns the Apple SF Symbols name for a button.
func GetGamepadAppleSFSymbolsNameForButton(gamepad *Gamepad, button GamepadButton) string {
	return C.GoString(C.SDL_GetGamepadAppleSFSymbolsNameForButton(gamepad.c, C.SDL_GamepadButton(button)))
}

// GetGamepadAppleSFSymbolsNameForAxis returns the Apple SF Symbols name for an axis.
func GetGamepadAppleSFSymbolsNameForAxis(gamepad *Gamepad, axis GamepadAxis) string {
	return C.GoString(C.SDL_GetGamepadAppleSFSymbolsNameForAxis(gamepad.c, C.SDL_GamepadAxis(axis)))
}

// GamepadBindingType represents the type of a gamepad binding.
type GamepadBindingType int

const (
	GAMEPAD_BINDTYPE_NONE   GamepadBindingType = C.SDL_GAMEPAD_BINDTYPE_NONE
	GAMEPAD_BINDTYPE_BUTTON GamepadBindingType = C.SDL_GAMEPAD_BINDTYPE_BUTTON
	GAMEPAD_BINDTYPE_AXIS   GamepadBindingType = C.SDL_GAMEPAD_BINDTYPE_AXIS
	GAMEPAD_BINDTYPE_HAT    GamepadBindingType = C.SDL_GAMEPAD_BINDTYPE_HAT
)
