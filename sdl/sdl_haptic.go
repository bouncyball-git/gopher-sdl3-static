package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// Haptic represents an opened SDL haptic device.
type Haptic struct {
	c *C.SDL_Haptic
}

// HapticID is a unique identifier for a haptic device.
type HapticID uint32

// HapticEffectType represents the type of a haptic effect.
type HapticEffectType uint16

// HapticDirectionType represents the type of coordinates used for haptic direction.
type HapticDirectionType uint8

// HapticEffectID is an identifier for a haptic effect on a device.
type HapticEffectID int

// Haptic effect type constants.
const (
	HAPTIC_CONSTANT     HapticEffectType = C.SDL_HAPTIC_CONSTANT
	HAPTIC_SINE         HapticEffectType = C.SDL_HAPTIC_SINE
	HAPTIC_SQUARE       HapticEffectType = C.SDL_HAPTIC_SQUARE
	HAPTIC_TRIANGLE     HapticEffectType = C.SDL_HAPTIC_TRIANGLE
	HAPTIC_SAWTOOTHUP   HapticEffectType = C.SDL_HAPTIC_SAWTOOTHUP
	HAPTIC_SAWTOOTHDOWN HapticEffectType = C.SDL_HAPTIC_SAWTOOTHDOWN
	HAPTIC_RAMP         HapticEffectType = C.SDL_HAPTIC_RAMP
	HAPTIC_SPRING       HapticEffectType = C.SDL_HAPTIC_SPRING
	HAPTIC_DAMPER       HapticEffectType = C.SDL_HAPTIC_DAMPER
	HAPTIC_INERTIA      HapticEffectType = C.SDL_HAPTIC_INERTIA
	HAPTIC_FRICTION     HapticEffectType = C.SDL_HAPTIC_FRICTION
	HAPTIC_LEFTRIGHT    HapticEffectType = C.SDL_HAPTIC_LEFTRIGHT
	HAPTIC_RESERVED1    HapticEffectType = C.SDL_HAPTIC_RESERVED1
	HAPTIC_RESERVED2    HapticEffectType = C.SDL_HAPTIC_RESERVED2
	HAPTIC_RESERVED3    HapticEffectType = C.SDL_HAPTIC_RESERVED3
	HAPTIC_CUSTOM       HapticEffectType = C.SDL_HAPTIC_CUSTOM
)

// Haptic feature constants.
const (
	HAPTIC_GAIN       = C.SDL_HAPTIC_GAIN
	HAPTIC_AUTOCENTER = C.SDL_HAPTIC_AUTOCENTER
	HAPTIC_STATUS     = C.SDL_HAPTIC_STATUS
	HAPTIC_PAUSE      = C.SDL_HAPTIC_PAUSE
)

// Direction encoding constants.
const (
	HAPTIC_POLAR        HapticDirectionType = C.SDL_HAPTIC_POLAR
	HAPTIC_CARTESIAN    HapticDirectionType = C.SDL_HAPTIC_CARTESIAN
	HAPTIC_SPHERICAL    HapticDirectionType = C.SDL_HAPTIC_SPHERICAL
	HAPTIC_STEERING_AXIS HapticDirectionType = C.SDL_HAPTIC_STEERING_AXIS
)

// HAPTIC_INFINITY is used to play a device an infinite number of times.
const HAPTIC_INFINITY = C.SDL_HAPTIC_INFINITY

// HapticDirection represents a haptic direction.
type HapticDirection struct {
	Type HapticDirectionType
	Dir  [3]int32
}

// HapticConstant contains a template for a constant haptic effect.
type HapticConstant struct {
	Type      HapticEffectType
	Direction HapticDirection

	Length uint32
	Delay  uint16

	Button   uint16
	Interval uint16

	Level int16

	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

// HapticPeriodic contains a template for a periodic haptic effect.
type HapticPeriodic struct {
	Type      HapticEffectType
	Direction HapticDirection

	Length uint32
	Delay  uint16

	Button   uint16
	Interval uint16

	Period    uint16
	Magnitude int16
	Offset    int16
	Phase     uint16

	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

// HapticCondition contains a template for a condition haptic effect.
type HapticCondition struct {
	Type      HapticEffectType
	Direction HapticDirection

	Length uint32
	Delay  uint16

	Button   uint16
	Interval uint16

	RightSat   [3]uint16
	LeftSat    [3]uint16
	RightCoeff [3]int16
	LeftCoeff  [3]int16
	Deadband   [3]uint16
	Center     [3]int16
}

// HapticRamp contains a template for a ramp haptic effect.
type HapticRamp struct {
	Type      HapticEffectType
	Direction HapticDirection

	Length uint32
	Delay  uint16

	Button   uint16
	Interval uint16

	Start int16
	End   int16

	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

// HapticLeftRight contains a template for a left/right haptic effect.
type HapticLeftRight struct {
	Type HapticEffectType

	Length uint32

	LargeMagnitude uint16
	SmallMagnitude uint16
}

// HapticCustom contains a template for a custom haptic effect.
type HapticCustom struct {
	Type      HapticEffectType
	Direction HapticDirection

	Length uint32
	Delay  uint16

	Button   uint16
	Interval uint16

	Channels uint8
	Period   uint16
	Samples  uint16
	Data     *uint16

	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

// --- Device enumeration ---

// GetHaptics returns a list of currently connected haptic device IDs.
func GetHaptics() []HapticID {
	var count C.int
	cids := C.SDL_GetHaptics(&count)
	if cids == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cids))
	n := int(count)
	result := make([]HapticID, n)
	slice := unsafe.Slice((*C.SDL_HapticID)(cids), n)
	for i, id := range slice {
		result[i] = HapticID(id)
	}
	return result
}

// GetHapticNameForID returns the name of a haptic device by instance ID.
func GetHapticNameForID(instanceID HapticID) string {
	return C.GoString(C.SDL_GetHapticNameForID(C.SDL_HapticID(instanceID)))
}

// OpenHaptic opens a haptic device for use.
func OpenHaptic(instanceID HapticID) (*Haptic, error) {
	c := C.SDL_OpenHaptic(C.SDL_HapticID(instanceID))
	if c == nil {
		return nil, getError()
	}
	return &Haptic{c: c}, nil
}

// GetHapticFromID returns the Haptic associated with an instance ID, if opened.
func GetHapticFromID(instanceID HapticID) (*Haptic, error) {
	c := C.SDL_GetHapticFromID(C.SDL_HapticID(instanceID))
	if c == nil {
		return nil, getError()
	}
	return &Haptic{c: c}, nil
}

// Close closes the haptic device.
func (h *Haptic) Close() {
	if h.c != nil {
		C.SDL_CloseHaptic(h.c)
		h.c = nil
	}
}

// Name returns the implementation dependent name of the haptic device.
func (h *Haptic) Name() string {
	return C.GoString(C.SDL_GetHapticName(h.c))
}

// ID returns the instance ID of the haptic device.
func (h *Haptic) ID() (HapticID, error) {
	id := C.SDL_GetHapticID(h.c)
	if id == 0 {
		return 0, getError()
	}
	return HapticID(id), nil
}

// --- Mouse/Joystick haptic ---

// IsMouseHaptic returns true if the current mouse has haptic capabilities.
func IsMouseHaptic() bool {
	return bool(C.SDL_IsMouseHaptic())
}

// OpenHapticFromMouse opens a haptic device from the current mouse.
func OpenHapticFromMouse() (*Haptic, error) {
	c := C.SDL_OpenHapticFromMouse()
	if c == nil {
		return nil, getError()
	}
	return &Haptic{c: c}, nil
}

// IsJoystickHaptic returns true if a joystick has haptic features.
func IsJoystickHaptic(joystick *Joystick) bool {
	return bool(C.SDL_IsJoystickHaptic(joystick.c))
}

// OpenHapticFromJoystick opens a haptic device from a joystick.
func OpenHapticFromJoystick(joystick *Joystick) (*Haptic, error) {
	c := C.SDL_OpenHapticFromJoystick(joystick.c)
	if c == nil {
		return nil, getError()
	}
	return &Haptic{c: c}, nil
}

// --- Device queries ---

// MaxEffects returns the number of effects a haptic device can store.
func (h *Haptic) MaxEffects() (int, error) {
	n := C.SDL_GetMaxHapticEffects(h.c)
	if n < 0 {
		return 0, getError()
	}
	return int(n), nil
}

// MaxEffectsPlaying returns the number of effects a haptic device can play at once.
func (h *Haptic) MaxEffectsPlaying() (int, error) {
	n := C.SDL_GetMaxHapticEffectsPlaying(h.c)
	if n < 0 {
		return 0, getError()
	}
	return int(n), nil
}

// Features returns the haptic device's supported features as a bitmask.
func (h *Haptic) Features() uint32 {
	return uint32(C.SDL_GetHapticFeatures(h.c))
}

// NumAxes returns the number of haptic axes the device has.
func (h *Haptic) NumAxes() (int, error) {
	n := C.SDL_GetNumHapticAxes(h.c)
	if n < 0 {
		return 0, getError()
	}
	return int(n), nil
}

// --- Effect management ---

// EffectSupported checks if an effect is supported by the haptic device.
// The effect parameter should be a pointer to an SDL_HapticEffect union (use unsafe.Pointer).
func (h *Haptic) EffectSupported(effect unsafe.Pointer) bool {
	return bool(C.SDL_HapticEffectSupported(h.c, (*C.SDL_HapticEffect)(effect)))
}

// CreateEffect creates a new haptic effect on the device.
// The effect parameter should be a pointer to an SDL_HapticEffect union (use unsafe.Pointer).
func (h *Haptic) CreateEffect(effect unsafe.Pointer) (HapticEffectID, error) {
	id := C.SDL_CreateHapticEffect(h.c, (*C.SDL_HapticEffect)(effect))
	if id < 0 {
		return HapticEffectID(id), getError()
	}
	return HapticEffectID(id), nil
}

// UpdateEffect updates the properties of an existing haptic effect.
// The data parameter should be a pointer to an SDL_HapticEffect union (use unsafe.Pointer).
func (h *Haptic) UpdateEffect(effect HapticEffectID, data unsafe.Pointer) error {
	if !C.SDL_UpdateHapticEffect(h.c, C.SDL_HapticEffectID(effect), (*C.SDL_HapticEffect)(data)) {
		return getError()
	}
	return nil
}

// RunEffect runs a haptic effect on the device.
func (h *Haptic) RunEffect(effect HapticEffectID, iterations uint32) error {
	if !C.SDL_RunHapticEffect(h.c, C.SDL_HapticEffectID(effect), C.Uint32(iterations)) {
		return getError()
	}
	return nil
}

// StopEffect stops a haptic effect on the device.
func (h *Haptic) StopEffect(effect HapticEffectID) error {
	if !C.SDL_StopHapticEffect(h.c, C.SDL_HapticEffectID(effect)) {
		return getError()
	}
	return nil
}

// DestroyEffect destroys a haptic effect on the device.
func (h *Haptic) DestroyEffect(effect HapticEffectID) {
	C.SDL_DestroyHapticEffect(h.c, C.SDL_HapticEffectID(effect))
}

// EffectStatus returns true if the effect is currently playing.
func (h *Haptic) EffectStatus(effect HapticEffectID) bool {
	return bool(C.SDL_GetHapticEffectStatus(h.c, C.SDL_HapticEffectID(effect)))
}

// --- Device settings ---

// SetGain sets the global gain of the haptic device (0-100).
func (h *Haptic) SetGain(gain int) error {
	if !C.SDL_SetHapticGain(h.c, C.int(gain)) {
		return getError()
	}
	return nil
}

// SetAutocenter sets the global autocenter of the haptic device (0-100).
func (h *Haptic) SetAutocenter(autocenter int) error {
	if !C.SDL_SetHapticAutocenter(h.c, C.int(autocenter)) {
		return getError()
	}
	return nil
}

// Pause pauses the haptic device.
func (h *Haptic) Pause() error {
	if !C.SDL_PauseHaptic(h.c) {
		return getError()
	}
	return nil
}

// Resume resumes the haptic device.
func (h *Haptic) Resume() error {
	if !C.SDL_ResumeHaptic(h.c) {
		return getError()
	}
	return nil
}

// StopEffects stops all currently playing effects on the haptic device.
func (h *Haptic) StopEffects() error {
	if !C.SDL_StopHapticEffects(h.c) {
		return getError()
	}
	return nil
}

// --- Rumble ---

// RumbleSupported checks whether rumble is supported on the haptic device.
func (h *Haptic) RumbleSupported() bool {
	return bool(C.SDL_HapticRumbleSupported(h.c))
}

// InitRumble initializes the haptic device for simple rumble playback.
func (h *Haptic) InitRumble() error {
	if !C.SDL_InitHapticRumble(h.c) {
		return getError()
	}
	return nil
}

// PlayRumble runs a simple rumble effect on the haptic device.
func (h *Haptic) PlayRumble(strength float32, length uint32) error {
	if !C.SDL_PlayHapticRumble(h.c, C.float(strength), C.Uint32(length)) {
		return getError()
	}
	return nil
}

// StopRumble stops the simple rumble on the haptic device.
func (h *Haptic) StopRumble() error {
	if !C.SDL_StopHapticRumble(h.c) {
		return getError()
	}
	return nil
}

// --- Package-level function aliases ---

// GetMaxHapticEffects returns the number of effects a haptic device can store.
func GetMaxHapticEffects(haptic *Haptic) (int, error) {
	return haptic.MaxEffects()
}

// GetMaxHapticEffectsPlaying returns the number of effects that can play simultaneously.
func GetMaxHapticEffectsPlaying(haptic *Haptic) (int, error) {
	return haptic.MaxEffectsPlaying()
}

// GetHapticFeatures returns the device's supported features as a bitmask.
func GetHapticFeatures(haptic *Haptic) uint32 {
	return haptic.Features()
}

// GetNumHapticAxes returns the number of haptic axes the device has.
func GetNumHapticAxes(haptic *Haptic) (int, error) {
	return haptic.NumAxes()
}

// HapticEffectSupported checks if an effect is supported by the device.
func HapticEffectSupported(haptic *Haptic, effect unsafe.Pointer) bool {
	return haptic.EffectSupported(effect)
}

// CreateHapticEffect creates a new haptic effect on the device.
func CreateHapticEffect(haptic *Haptic, effect unsafe.Pointer) (HapticEffectID, error) {
	return haptic.CreateEffect(effect)
}

// UpdateHapticEffect updates the properties of an existing haptic effect.
func UpdateHapticEffect(haptic *Haptic, effect HapticEffectID, data unsafe.Pointer) error {
	return haptic.UpdateEffect(effect, data)
}

// RunHapticEffect runs a haptic effect on the device.
func RunHapticEffect(haptic *Haptic, effect HapticEffectID, iterations uint32) error {
	return haptic.RunEffect(effect, iterations)
}

// StopHapticEffect stops a haptic effect on the device.
func StopHapticEffect(haptic *Haptic, effect HapticEffectID) error {
	return haptic.StopEffect(effect)
}

// DestroyHapticEffect destroys a haptic effect on the device.
func DestroyHapticEffect(haptic *Haptic, effect HapticEffectID) {
	haptic.DestroyEffect(effect)
}

// GetHapticEffectStatus returns true if the effect is currently playing.
func GetHapticEffectStatus(haptic *Haptic, effect HapticEffectID) bool {
	return haptic.EffectStatus(effect)
}

// SetHapticGain sets the global gain of the haptic device.
func SetHapticGain(haptic *Haptic, gain int) error {
	return haptic.SetGain(gain)
}

// SetHapticAutocenter sets the global autocenter of the haptic device.
func SetHapticAutocenter(haptic *Haptic, autocenter int) error {
	return haptic.SetAutocenter(autocenter)
}

// PauseHaptic pauses the haptic device.
func PauseHaptic(haptic *Haptic) error {
	return haptic.Pause()
}

// ResumeHaptic resumes the haptic device.
func ResumeHaptic(haptic *Haptic) error {
	return haptic.Resume()
}

// StopHapticEffects stops all effects on the haptic device.
func StopHapticEffects(haptic *Haptic) error {
	return haptic.StopEffects()
}

// HapticRumbleSupported checks whether rumble is supported.
func HapticRumbleSupported(haptic *Haptic) bool {
	return haptic.RumbleSupported()
}

// InitHapticRumble initializes the haptic device for simple rumble playback.
func InitHapticRumble(haptic *Haptic) error {
	return haptic.InitRumble()
}

// PlayHapticRumble runs a simple rumble effect.
func PlayHapticRumble(haptic *Haptic, strength float32, length uint32) error {
	return haptic.PlayRumble(strength, length)
}

// StopHapticRumble stops the simple rumble.
func StopHapticRumble(haptic *Haptic) error {
	return haptic.StopRumble()
}

// CloseHaptic closes a haptic device.
func CloseHaptic(haptic *Haptic) {
	haptic.Close()
}

// GetHapticName returns the name of the haptic device.
func GetHapticName(haptic *Haptic) string {
	return haptic.Name()
}

// GetHapticID returns the instance ID of the haptic device.
func GetHapticID(haptic *Haptic) (HapticID, error) {
	return haptic.ID()
}
