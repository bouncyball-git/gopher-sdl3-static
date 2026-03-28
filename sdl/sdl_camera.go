package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// Camera represents an opened SDL camera device.
type Camera struct {
	c *C.SDL_Camera
}

// CameraID is a unique identifier for a camera device.
type CameraID uint32

// CameraSpec describes a camera output format.
type CameraSpec struct {
	Format               PixelFormat
	Colorspace           Colorspace
	Width                int32
	Height               int32
	FramerateNumerator   int32
	FramerateDenominator int32
}

func (s *CameraSpec) cptr() *C.SDL_CameraSpec {
	return (*C.SDL_CameraSpec)(unsafe.Pointer(s))
}

func cameraSpecFromC(cs *C.SDL_CameraSpec) CameraSpec {
	return CameraSpec{
		Format:               PixelFormat(cs.format),
		Colorspace:           Colorspace(cs.colorspace),
		Width:                int32(cs.width),
		Height:               int32(cs.height),
		FramerateNumerator:   int32(cs.framerate_numerator),
		FramerateDenominator: int32(cs.framerate_denominator),
	}
}

// CameraPosition represents the position of a camera on a device.
type CameraPosition int32

// Camera position constants.
const (
	CAMERA_POSITION_UNKNOWN      CameraPosition = C.SDL_CAMERA_POSITION_UNKNOWN
	CAMERA_POSITION_FRONT_FACING CameraPosition = C.SDL_CAMERA_POSITION_FRONT_FACING
	CAMERA_POSITION_BACK_FACING  CameraPosition = C.SDL_CAMERA_POSITION_BACK_FACING
)

// CameraPermissionState represents the current state of a camera access request.
type CameraPermissionState int32

// Camera permission state constants.
const (
	CAMERA_PERMISSION_STATE_DENIED   CameraPermissionState = C.SDL_CAMERA_PERMISSION_STATE_DENIED
	CAMERA_PERMISSION_STATE_PENDING  CameraPermissionState = C.SDL_CAMERA_PERMISSION_STATE_PENDING
	CAMERA_PERMISSION_STATE_APPROVED CameraPermissionState = C.SDL_CAMERA_PERMISSION_STATE_APPROVED
)

// --- Driver functions ---

// GetNumCameraDrivers returns the number of built-in camera drivers.
func GetNumCameraDrivers() int {
	return int(C.SDL_GetNumCameraDrivers())
}

// GetCameraDriver returns the name of a built-in camera driver by index.
func GetCameraDriver(index int) string {
	return C.GoString(C.SDL_GetCameraDriver(C.int(index)))
}

// GetCurrentCameraDriver returns the name of the current camera driver.
func GetCurrentCameraDriver() string {
	return C.GoString(C.SDL_GetCurrentCameraDriver())
}

// --- Device enumeration ---

// GetCameras returns a list of currently connected camera device IDs.
func GetCameras() []CameraID {
	var count C.int
	cids := C.SDL_GetCameras(&count)
	if cids == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cids))
	n := int(count)
	result := make([]CameraID, n)
	slice := unsafe.Slice((*C.SDL_CameraID)(cids), n)
	for i, id := range slice {
		result[i] = CameraID(id)
	}
	return result
}

// GetCameraSupportedFormats returns the list of native formats/sizes a camera supports.
func GetCameraSupportedFormats(instanceID CameraID) []CameraSpec {
	var count C.int
	cspecs := C.SDL_GetCameraSupportedFormats(C.SDL_CameraID(instanceID), &count)
	if cspecs == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cspecs))
	n := int(count)
	result := make([]CameraSpec, n)
	slice := unsafe.Slice((**C.SDL_CameraSpec)(cspecs), n)
	for i, cs := range slice {
		result[i] = cameraSpecFromC(cs)
	}
	return result
}

// GetCameraName returns the human-readable name of a camera device.
func GetCameraName(instanceID CameraID) string {
	return C.GoString(C.SDL_GetCameraName(C.SDL_CameraID(instanceID)))
}

// GetCameraPosition returns the position of the camera in relation to the system.
func GetCameraPosition(instanceID CameraID) CameraPosition {
	return CameraPosition(C.SDL_GetCameraPosition(C.SDL_CameraID(instanceID)))
}

// --- Open/Close ---

// OpenCamera opens a camera device for use.
func OpenCamera(instanceID CameraID, spec *CameraSpec) (*Camera, error) {
	var cs *C.SDL_CameraSpec
	if spec != nil {
		cs = spec.cptr()
	}
	c := C.SDL_OpenCamera(C.SDL_CameraID(instanceID), cs)
	if c == nil {
		return nil, getError()
	}
	return &Camera{c: c}, nil
}

// Close closes the camera device.
func (cam *Camera) Close() {
	if cam.c != nil {
		C.SDL_CloseCamera(cam.c)
		cam.c = nil
	}
}

// --- Camera methods ---

// ID returns the instance ID of the opened camera.
func (cam *Camera) ID() (CameraID, error) {
	id := C.SDL_GetCameraID(cam.c)
	if id == 0 {
		return 0, getError()
	}
	return CameraID(id), nil
}

// Properties returns the properties associated with the camera.
func (cam *Camera) Properties() (PropertiesID, error) {
	id := C.SDL_GetCameraProperties(cam.c)
	if id == 0 {
		return 0, getError()
	}
	return PropertiesID(id), nil
}

// GetFormat returns the format the camera is using when generating images.
func (cam *Camera) GetFormat() (CameraSpec, error) {
	var spec CameraSpec
	if !C.SDL_GetCameraFormat(cam.c, spec.cptr()) {
		return spec, getError()
	}
	return spec, nil
}

// PermissionState returns the current camera permission state.
func (cam *Camera) PermissionState() CameraPermissionState {
	return CameraPermissionState(C.SDL_GetCameraPermissionState(cam.c))
}

// AcquireFrame acquires a frame of video from the camera.
// Returns the surface and timestamp in nanoseconds.
// The returned surface must be released with ReleaseFrame. Do not call Surface.Destroy on it.
func (cam *Camera) AcquireFrame() (*Surface, uint64) {
	var timestampNS C.Uint64
	cs := C.SDL_AcquireCameraFrame(cam.c, &timestampNS)
	if cs == nil {
		return nil, 0
	}
	return &Surface{c: cs}, uint64(timestampNS)
}

// ReleaseFrame releases a frame previously acquired with AcquireFrame.
func (cam *Camera) ReleaseFrame(surface *Surface) {
	if surface != nil && surface.c != nil {
		C.SDL_ReleaseCameraFrame(cam.c, surface.c)
		surface.c = nil
	}
}

// --- Package-level function aliases ---

// GetCameraID returns the instance ID of an opened camera.
func GetCameraID(camera *Camera) (CameraID, error) {
	return camera.ID()
}

// GetCameraProperties returns the properties associated with a camera.
func GetCameraProperties(camera *Camera) (PropertiesID, error) {
	return camera.Properties()
}

// GetCameraFormat returns the format the camera is using.
func GetCameraFormat(camera *Camera) (CameraSpec, error) {
	return camera.GetFormat()
}

// GetCameraPermissionState returns the camera permission state.
func GetCameraPermissionState(camera *Camera) CameraPermissionState {
	return camera.PermissionState()
}

// AcquireCameraFrame acquires a frame from the camera.
func AcquireCameraFrame(camera *Camera) (*Surface, uint64) {
	return camera.AcquireFrame()
}

// ReleaseCameraFrame releases a frame back to the camera subsystem.
func ReleaseCameraFrame(camera *Camera, surface *Surface) {
	camera.ReleaseFrame(surface)
}

// CloseCamera closes a camera device.
func CloseCamera(camera *Camera) {
	camera.Close()
}
