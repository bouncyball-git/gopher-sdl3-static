package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// Sensor represents an opened SDL sensor device.
type Sensor struct {
	c *C.SDL_Sensor
}

// SensorID is a unique identifier for a sensor device.
type SensorID uint32

// SensorType represents the type of a sensor.
type SensorType int32

// Sensor type constants.
const (
	SENSOR_INVALID SensorType = C.SDL_SENSOR_INVALID
	SENSOR_UNKNOWN SensorType = C.SDL_SENSOR_UNKNOWN
	SENSOR_ACCEL   SensorType = C.SDL_SENSOR_ACCEL
	SENSOR_GYRO    SensorType = C.SDL_SENSOR_GYRO
	SENSOR_ACCEL_L SensorType = C.SDL_SENSOR_ACCEL_L
	SENSOR_GYRO_L  SensorType = C.SDL_SENSOR_GYRO_L
	SENSOR_ACCEL_R SensorType = C.SDL_SENSOR_ACCEL_R
	SENSOR_GYRO_R  SensorType = C.SDL_SENSOR_GYRO_R
)

// STANDARD_GRAVITY is the standard gravity constant for accelerometer sensors.
const STANDARD_GRAVITY = 9.80665

// --- Enumeration ---

// GetSensors returns a list of currently connected sensor IDs.
func GetSensors() []SensorID {
	var count C.int
	cids := C.SDL_GetSensors(&count)
	if cids == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cids))
	n := int(count)
	result := make([]SensorID, n)
	slice := unsafe.Slice((*C.SDL_SensorID)(cids), n)
	for i, id := range slice {
		result[i] = SensorID(id)
	}
	return result
}

// GetSensorNameForID returns the name of a sensor by instance ID.
func GetSensorNameForID(instanceID SensorID) string {
	return C.GoString(C.SDL_GetSensorNameForID(C.SDL_SensorID(instanceID)))
}

// GetSensorTypeForID returns the type of a sensor by instance ID.
func GetSensorTypeForID(instanceID SensorID) SensorType {
	return SensorType(C.SDL_GetSensorTypeForID(C.SDL_SensorID(instanceID)))
}

// GetSensorNonPortableTypeForID returns the platform dependent type of a sensor.
func GetSensorNonPortableTypeForID(instanceID SensorID) int {
	return int(C.SDL_GetSensorNonPortableTypeForID(C.SDL_SensorID(instanceID)))
}

// --- Open/Close ---

// OpenSensor opens a sensor for use.
func OpenSensor(instanceID SensorID) (*Sensor, error) {
	c := C.SDL_OpenSensor(C.SDL_SensorID(instanceID))
	if c == nil {
		return nil, getError()
	}
	return &Sensor{c: c}, nil
}

// GetSensorFromID returns the Sensor associated with an instance ID, if opened.
func GetSensorFromID(instanceID SensorID) (*Sensor, error) {
	c := C.SDL_GetSensorFromID(C.SDL_SensorID(instanceID))
	if c == nil {
		return nil, getError()
	}
	return &Sensor{c: c}, nil
}

// --- Sensor methods ---

// Properties returns the properties associated with the sensor.
func (s *Sensor) Properties() (PropertiesID, error) {
	id := C.SDL_GetSensorProperties(s.c)
	if id == 0 {
		return 0, getError()
	}
	return PropertiesID(id), nil
}

// Name returns the implementation dependent name of the sensor.
func (s *Sensor) Name() string {
	return C.GoString(C.SDL_GetSensorName(s.c))
}

// Type returns the type of the sensor.
func (s *Sensor) Type() SensorType {
	return SensorType(C.SDL_GetSensorType(s.c))
}

// NonPortableType returns the platform dependent type of the sensor.
func (s *Sensor) NonPortableType() int {
	return int(C.SDL_GetSensorNonPortableType(s.c))
}

// ID returns the instance ID of the sensor.
func (s *Sensor) ID() (SensorID, error) {
	id := C.SDL_GetSensorID(s.c)
	if id == 0 {
		return 0, getError()
	}
	return SensorID(id), nil
}

// GetData reads the current state of the sensor into the provided slice.
func (s *Sensor) GetData(data []float32) error {
	if len(data) == 0 {
		return nil
	}
	if !C.SDL_GetSensorData(s.c, (*C.float)(unsafe.Pointer(&data[0])), C.int(len(data))) {
		return getError()
	}
	return nil
}

// Close closes the sensor.
func (s *Sensor) Close() {
	if s.c != nil {
		C.SDL_CloseSensor(s.c)
		s.c = nil
	}
}

// --- Global ---

// UpdateSensors updates the current state of all open sensors.
func UpdateSensors() {
	C.SDL_UpdateSensors()
}
