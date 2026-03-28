package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// JoystickID is a unique identifier for a joystick for the time it is connected.
type JoystickID uint32

// Joystick represents an opened SDL joystick device.
type Joystick struct {
	c *C.SDL_Joystick
}

// JoystickType identifies common joystick types.
type JoystickType int32

// JoystickType constants.
const (
	JOYSTICK_TYPE_UNKNOWN      JoystickType = C.SDL_JOYSTICK_TYPE_UNKNOWN
	JOYSTICK_TYPE_GAMEPAD      JoystickType = C.SDL_JOYSTICK_TYPE_GAMEPAD
	JOYSTICK_TYPE_WHEEL        JoystickType = C.SDL_JOYSTICK_TYPE_WHEEL
	JOYSTICK_TYPE_ARCADE_STICK JoystickType = C.SDL_JOYSTICK_TYPE_ARCADE_STICK
	JOYSTICK_TYPE_FLIGHT_STICK JoystickType = C.SDL_JOYSTICK_TYPE_FLIGHT_STICK
	JOYSTICK_TYPE_DANCE_PAD    JoystickType = C.SDL_JOYSTICK_TYPE_DANCE_PAD
	JOYSTICK_TYPE_GUITAR       JoystickType = C.SDL_JOYSTICK_TYPE_GUITAR
	JOYSTICK_TYPE_DRUM_KIT     JoystickType = C.SDL_JOYSTICK_TYPE_DRUM_KIT
	JOYSTICK_TYPE_ARCADE_PAD   JoystickType = C.SDL_JOYSTICK_TYPE_ARCADE_PAD
	JOYSTICK_TYPE_THROTTLE     JoystickType = C.SDL_JOYSTICK_TYPE_THROTTLE
)

// JoystickConnectionState describes how a joystick is connected.
type JoystickConnectionState int32

// JoystickConnectionState constants.
const (
	JOYSTICK_CONNECTION_INVALID  JoystickConnectionState = C.SDL_JOYSTICK_CONNECTION_INVALID
	JOYSTICK_CONNECTION_UNKNOWN  JoystickConnectionState = C.SDL_JOYSTICK_CONNECTION_UNKNOWN
	JOYSTICK_CONNECTION_WIRED    JoystickConnectionState = C.SDL_JOYSTICK_CONNECTION_WIRED
	JOYSTICK_CONNECTION_WIRELESS JoystickConnectionState = C.SDL_JOYSTICK_CONNECTION_WIRELESS
)

// Hat position constants.
const (
	HAT_CENTERED  = C.SDL_HAT_CENTERED
	HAT_UP        = C.SDL_HAT_UP
	HAT_RIGHT     = C.SDL_HAT_RIGHT
	HAT_DOWN      = C.SDL_HAT_DOWN
	HAT_LEFT      = C.SDL_HAT_LEFT
	HAT_RIGHTUP   = C.SDL_HAT_RIGHTUP
	HAT_RIGHTDOWN = C.SDL_HAT_RIGHTDOWN
	HAT_LEFTUP    = C.SDL_HAT_LEFTUP
	HAT_LEFTDOWN  = C.SDL_HAT_LEFTDOWN
)

// Axis value limits.
const (
	JOYSTICK_AXIS_MAX = C.SDL_JOYSTICK_AXIS_MAX
	JOYSTICK_AXIS_MIN = C.SDL_JOYSTICK_AXIS_MIN
)

// PowerState represents the power supply state.
type PowerState int32

// PowerState constants.
const (
	POWERSTATE_ERROR      PowerState = C.SDL_POWERSTATE_ERROR
	POWERSTATE_UNKNOWN    PowerState = C.SDL_POWERSTATE_UNKNOWN
	POWERSTATE_ON_BATTERY PowerState = C.SDL_POWERSTATE_ON_BATTERY
	POWERSTATE_NO_BATTERY PowerState = C.SDL_POWERSTATE_NO_BATTERY
	POWERSTATE_CHARGING   PowerState = C.SDL_POWERSTATE_CHARGING
	POWERSTATE_CHARGED    PowerState = C.SDL_POWERSTATE_CHARGED
)

// --- Global functions ---

// LockJoysticks locks the joystick subsystem for atomic access.
func LockJoysticks() {
	C.SDL_LockJoysticks()
}

// UnlockJoysticks unlocks the joystick subsystem.
func UnlockJoysticks() {
	C.SDL_UnlockJoysticks()
}

// HasJoystick returns true if a joystick is currently connected.
func HasJoystick() bool {
	return bool(C.SDL_HasJoystick())
}

// GetJoysticks returns a list of currently connected joystick instance IDs.
func GetJoysticks() []JoystickID {
	var count C.int
	cids := C.SDL_GetJoysticks(&count)
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

// GetJoystickNameForID returns the name of a joystick by instance ID.
func GetJoystickNameForID(id JoystickID) string {
	return C.GoString(C.SDL_GetJoystickNameForID(C.SDL_JoystickID(id)))
}

// GetJoystickPathForID returns the path of a joystick by instance ID.
func GetJoystickPathForID(id JoystickID) string {
	return C.GoString(C.SDL_GetJoystickPathForID(C.SDL_JoystickID(id)))
}

// GetJoystickTypeForID returns the type of a joystick by instance ID.
func GetJoystickTypeForID(id JoystickID) JoystickType {
	return JoystickType(C.SDL_GetJoystickTypeForID(C.SDL_JoystickID(id)))
}

// OpenJoystick opens a joystick for use.
func OpenJoystick(id JoystickID) (*Joystick, error) {
	ptr := C.SDL_OpenJoystick(C.SDL_JoystickID(id))
	if ptr == nil {
		return nil, getError()
	}
	return &Joystick{c: ptr}, nil
}

// GetJoystickFromID returns the Joystick associated with an instance ID, if it has been opened.
func GetJoystickFromID(id JoystickID) (*Joystick, error) {
	ptr := C.SDL_GetJoystickFromID(C.SDL_JoystickID(id))
	if ptr == nil {
		return nil, getError()
	}
	return &Joystick{c: ptr}, nil
}

// GetJoystickFromPlayerIndex returns the Joystick associated with a player index.
func GetJoystickFromPlayerIndex(playerIndex int) (*Joystick, error) {
	ptr := C.SDL_GetJoystickFromPlayerIndex(C.int(playerIndex))
	if ptr == nil {
		return nil, getError()
	}
	return &Joystick{c: ptr}, nil
}

// SetJoystickEventsEnabled enables or disables joystick event processing.
func SetJoystickEventsEnabled(enabled bool) {
	C.SDL_SetJoystickEventsEnabled(C.bool(enabled))
}

// JoystickEventsEnabled returns true if joystick events are being processed.
func JoystickEventsEnabled() bool {
	return bool(C.SDL_JoystickEventsEnabled())
}

// UpdateJoysticks updates the current state of the open joysticks.
func UpdateJoysticks() {
	C.SDL_UpdateJoysticks()
}

// --- Joystick methods ---

// Name returns the name of the joystick.
func (j *Joystick) Name() string {
	return C.GoString(C.SDL_GetJoystickName(j.c))
}

// Path returns the implementation-dependent path of the joystick.
func (j *Joystick) Path() string {
	return C.GoString(C.SDL_GetJoystickPath(j.c))
}

// Type returns the type of the joystick.
func (j *Joystick) Type() JoystickType {
	return JoystickType(C.SDL_GetJoystickType(j.c))
}

// ID returns the instance ID of the joystick.
func (j *Joystick) ID() JoystickID {
	return JoystickID(C.SDL_GetJoystickID(j.c))
}

// Connected returns true if the joystick is currently connected.
func (j *Joystick) Connected() bool {
	return bool(C.SDL_JoystickConnected(j.c))
}

// Close closes the joystick.
func (j *Joystick) Close() {
	if j.c != nil {
		C.SDL_CloseJoystick(j.c)
		j.c = nil
	}
}

// --- Counts ---

// GetNumJoystickAxes returns the number of axes on the joystick.
func GetNumJoystickAxes(j *Joystick) (int, error) {
	n := C.SDL_GetNumJoystickAxes(j.c)
	if n < 0 {
		return 0, getError()
	}
	return int(n), nil
}

// GetNumJoystickBalls returns the number of trackballs on the joystick.
func GetNumJoystickBalls(j *Joystick) (int, error) {
	n := C.SDL_GetNumJoystickBalls(j.c)
	if n < 0 {
		return 0, getError()
	}
	return int(n), nil
}

// GetNumJoystickHats returns the number of POV hats on the joystick.
func GetNumJoystickHats(j *Joystick) (int, error) {
	n := C.SDL_GetNumJoystickHats(j.c)
	if n < 0 {
		return 0, getError()
	}
	return int(n), nil
}

// GetNumJoystickButtons returns the number of buttons on the joystick.
func GetNumJoystickButtons(j *Joystick) (int, error) {
	n := C.SDL_GetNumJoystickButtons(j.c)
	if n < 0 {
		return 0, getError()
	}
	return int(n), nil
}

// --- Input state ---

// GetJoystickAxis returns the current state of an axis on the joystick.
func GetJoystickAxis(j *Joystick, axis int) int16 {
	return int16(C.SDL_GetJoystickAxis(j.c, C.int(axis)))
}

// GetJoystickHat returns the current state of a POV hat on the joystick.
func GetJoystickHat(j *Joystick, hat int) uint8 {
	return uint8(C.SDL_GetJoystickHat(j.c, C.int(hat)))
}

// GetJoystickButton returns the current state of a button on the joystick.
func GetJoystickButton(j *Joystick, button int) bool {
	return bool(C.SDL_GetJoystickButton(j.c, C.int(button)))
}

// GetJoystickBall returns the ball axis change since the last poll.
func GetJoystickBall(j *Joystick, ball int) (dx, dy int, err error) {
	var cdx, cdy C.int
	if !C.SDL_GetJoystickBall(j.c, C.int(ball), &cdx, &cdy) {
		return 0, 0, getError()
	}
	return int(cdx), int(cdy), nil
}

// --- Rumble and effects ---

// RumbleJoystick starts a rumble effect on the joystick.
func RumbleJoystick(j *Joystick, lowFreq, highFreq uint16, durationMS uint32) error {
	if !C.SDL_RumbleJoystick(j.c, C.Uint16(lowFreq), C.Uint16(highFreq), C.Uint32(durationMS)) {
		return getError()
	}
	return nil
}

// SetJoystickLED sets the LED color on the joystick.
func SetJoystickLED(j *Joystick, red, green, blue uint8) error {
	if !C.SDL_SetJoystickLED(j.c, C.Uint8(red), C.Uint8(green), C.Uint8(blue)) {
		return getError()
	}
	return nil
}

// SendJoystickEffect sends a joystick-specific effect packet.
func SendJoystickEffect(j *Joystick, data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if !C.SDL_SendJoystickEffect(j.c, unsafe.Pointer(&data[0]), C.int(len(data))) {
		return getError()
	}
	return nil
}

// --- Property accessors ---

// Vendor returns the USB vendor ID of the joystick, or 0 if unavailable.
func (j *Joystick) Vendor() uint16 {
	return uint16(C.SDL_GetJoystickVendor(j.c))
}

// Product returns the USB product ID of the joystick, or 0 if unavailable.
func (j *Joystick) Product() uint16 {
	return uint16(C.SDL_GetJoystickProduct(j.c))
}

// ProductVersion returns the product version of the joystick, or 0 if unavailable.
func (j *Joystick) ProductVersion() uint16 {
	return uint16(C.SDL_GetJoystickProductVersion(j.c))
}

// FirmwareVersion returns the firmware version of the joystick, or 0 if unavailable.
func (j *Joystick) FirmwareVersion() uint16 {
	return uint16(C.SDL_GetJoystickFirmwareVersion(j.c))
}

// Serial returns the serial number of the joystick, or an empty string if unavailable.
func (j *Joystick) Serial() string {
	cs := C.SDL_GetJoystickSerial(j.c)
	if cs == nil {
		return ""
	}
	return C.GoString(cs)
}

// PlayerIndex returns the player index of the joystick, or -1 if unavailable.
func (j *Joystick) PlayerIndex() int {
	return int(C.SDL_GetJoystickPlayerIndex(j.c))
}

// SetPlayerIndex sets the player index of the joystick.
func (j *Joystick) SetPlayerIndex(playerIndex int) error {
	if !C.SDL_SetJoystickPlayerIndex(j.c, C.int(playerIndex)) {
		return getError()
	}
	return nil
}

// ConnectionState returns the connection state of the joystick.
func (j *Joystick) ConnectionState() JoystickConnectionState {
	return JoystickConnectionState(C.SDL_GetJoystickConnectionState(j.c))
}

// PowerInfo returns the battery state and percentage of the joystick.
// The percent value is between 0 and 100, or -1 if unknown.
func (j *Joystick) PowerInfo() (PowerState, int) {
	var percent C.int
	state := C.SDL_GetJoystickPowerInfo(j.c, &percent)
	return PowerState(state), int(percent)
}

// TryLockJoysticks attempts to lock the joystick subsystem.
func TryLockJoysticks() bool {
	return bool(C.SDL_TryLockJoysticks())
}

// Properties returns the properties associated with the joystick.
func (j *Joystick) Properties() PropertiesID {
	return PropertiesID(C.SDL_GetJoystickProperties(j.c))
}

// GUID returns the GUID of an opened joystick.
func (j *Joystick) GUID() GUID {
	return guidFromC(C.SDL_GetJoystickGUID(j.c))
}

// GetJoystickGUIDForID returns the GUID of a joystick by instance ID.
func GetJoystickGUIDForID(instanceID JoystickID) GUID {
	return guidFromC(C.SDL_GetJoystickGUIDForID(C.SDL_JoystickID(instanceID)))
}

// GetJoystickGUIDInfo extracts vendor/product/version/crc16 from a GUID.
func GetJoystickGUIDInfo(guid GUID) (vendor, product, version, crc16 uint16) {
	var cv, cp, cvr, cc C.Uint16
	C.SDL_GetJoystickGUIDInfo(guid.cval(), &cv, &cp, &cvr, &cc)
	return uint16(cv), uint16(cp), uint16(cvr), uint16(cc)
}

// GetJoystickPlayerIndexForID returns the player index of a joystick by instance ID.
func GetJoystickPlayerIndexForID(instanceID JoystickID) int {
	return int(C.SDL_GetJoystickPlayerIndexForID(C.SDL_JoystickID(instanceID)))
}

// GetJoystickVendorForID returns the USB vendor ID of a joystick by instance ID.
func GetJoystickVendorForID(instanceID JoystickID) uint16 {
	return uint16(C.SDL_GetJoystickVendorForID(C.SDL_JoystickID(instanceID)))
}

// GetJoystickProductForID returns the USB product ID of a joystick by instance ID.
func GetJoystickProductForID(instanceID JoystickID) uint16 {
	return uint16(C.SDL_GetJoystickProductForID(C.SDL_JoystickID(instanceID)))
}

// GetJoystickProductVersionForID returns the product version of a joystick by instance ID.
func GetJoystickProductVersionForID(instanceID JoystickID) uint16 {
	return uint16(C.SDL_GetJoystickProductVersionForID(C.SDL_JoystickID(instanceID)))
}

// AxisInitialState returns the initial state of an axis, or false if no initial state.
func (j *Joystick) AxisInitialState(axis int) (int16, bool) {
	var state C.Sint16
	ok := C.SDL_GetJoystickAxisInitialState(j.c, C.int(axis), &state)
	return int16(state), bool(ok)
}

// RumbleTriggers rumbles the joystick triggers.
func (j *Joystick) RumbleTriggers(leftRumble, rightRumble uint16, durationMS uint32) error {
	if !C.SDL_RumbleJoystickTriggers(j.c, C.Uint16(leftRumble), C.Uint16(rightRumble), C.Uint32(durationMS)) {
		return getError()
	}
	return nil
}

// AttachVirtualJoystick attaches a virtual joystick from a descriptor.
// The desc parameter should be a pointer to an SDL_VirtualJoystickDesc struct.
func AttachVirtualJoystick(desc unsafe.Pointer) (JoystickID, error) {
	id := C.SDL_AttachVirtualJoystick((*C.SDL_VirtualJoystickDesc)(desc))
	if id == 0 {
		return 0, getError()
	}
	return JoystickID(id), nil
}

// DetachVirtualJoystick detaches a virtual joystick.
func DetachVirtualJoystick(instanceID JoystickID) error {
	if !C.SDL_DetachVirtualJoystick(C.SDL_JoystickID(instanceID)) {
		return getError()
	}
	return nil
}

// IsJoystickVirtual returns whether a joystick is virtual.
func IsJoystickVirtual(instanceID JoystickID) bool {
	return bool(C.SDL_IsJoystickVirtual(C.SDL_JoystickID(instanceID)))
}

// SetVirtualAxis sets the value of a virtual joystick axis.
func (j *Joystick) SetVirtualAxis(axis int, value int16) error {
	if !C.SDL_SetJoystickVirtualAxis(j.c, C.int(axis), C.Sint16(value)) {
		return getError()
	}
	return nil
}

// SetVirtualBall sets the value of a virtual joystick trackball.
func (j *Joystick) SetVirtualBall(ball int, xrel, yrel int16) error {
	if !C.SDL_SetJoystickVirtualBall(j.c, C.int(ball), C.Sint16(xrel), C.Sint16(yrel)) {
		return getError()
	}
	return nil
}

// SetVirtualButton sets the state of a virtual joystick button.
func (j *Joystick) SetVirtualButton(button int, down bool) error {
	if !C.SDL_SetJoystickVirtualButton(j.c, C.int(button), C.bool(down)) {
		return getError()
	}
	return nil
}

// SetVirtualHat sets the value of a virtual joystick hat.
func (j *Joystick) SetVirtualHat(hat int, value uint8) error {
	if !C.SDL_SetJoystickVirtualHat(j.c, C.int(hat), C.Uint8(value)) {
		return getError()
	}
	return nil
}

// SetVirtualTouchpad sets the state of a virtual joystick touchpad finger.
func (j *Joystick) SetVirtualTouchpad(touchpad, finger int, down bool, x, y, pressure float32) error {
	if !C.SDL_SetJoystickVirtualTouchpad(j.c, C.int(touchpad), C.int(finger), C.bool(down), C.float(x), C.float(y), C.float(pressure)) {
		return getError()
	}
	return nil
}

// SendVirtualSensorData sends sensor data for a virtual joystick.
func (j *Joystick) SendVirtualSensorData(sensorType SensorType, sensorTimestamp uint64, data []float32) error {
	if len(data) == 0 {
		return nil
	}
	if !C.SDL_SendJoystickVirtualSensorData(j.c, C.SDL_SensorType(sensorType), C.Uint64(sensorTimestamp), (*C.float)(unsafe.Pointer(&data[0])), C.int(len(data))) {
		return getError()
	}
	return nil
}

// VirtualJoystickTouchpadDesc describes a virtual joystick touchpad.
type VirtualJoystickTouchpadDesc struct {
	NFingers uint16
}

// VirtualJoystickSensorDesc describes a virtual joystick sensor.
type VirtualJoystickSensorDesc struct {
	Type SensorType
	Rate float32
}
