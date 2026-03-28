package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
#include <wchar.h>

static int wchar_to_utf8(const wchar_t *src, char *dst, int dstlen) {
	return (int)wcstombs(dst, src, (size_t)dstlen);
}
*/
import "C"

import "unsafe"

// HIDBusType represents HID underlying bus types.
type HIDBusType int

// HID bus type values.
const (
	HID_BUS_UNKNOWN   HIDBusType = C.SDL_HID_API_BUS_UNKNOWN
	HID_BUS_USB       HIDBusType = C.SDL_HID_API_BUS_USB
	HID_BUS_BLUETOOTH HIDBusType = C.SDL_HID_API_BUS_BLUETOOTH
	HID_BUS_I2C       HIDBusType = C.SDL_HID_API_BUS_I2C
	HID_BUS_SPI       HIDBusType = C.SDL_HID_API_BUS_SPI
)

// HIDDevice represents an open HID device.
type HIDDevice struct {
	c *C.SDL_hid_device
}

// HIDDeviceInfo contains information about a connected HID device.
type HIDDeviceInfo struct {
	Path               string
	VendorID           uint16
	ProductID          uint16
	SerialNumber       string
	ReleaseNumber      uint16
	ManufacturerString string
	ProductString      string
	UsagePage          uint16
	Usage              uint16
	InterfaceNumber    int
	InterfaceClass     int
	InterfaceSubclass  int
	InterfaceProtocol  int
	BusType            HIDBusType
}

// wcharToString converts a C wchar_t string to a Go string.
func wcharToString(wcs *C.wchar_t) string {
	if wcs == nil {
		return ""
	}
	buf := make([]C.char, 1024)
	n := C.wchar_to_utf8(wcs, &buf[0], 1024)
	if n < 0 {
		return ""
	}
	return C.GoStringN(&buf[0], n)
}

// HID_Init initializes the HIDAPI library.
func HID_Init() error {
	if C.SDL_hid_init() != 0 {
		return getError()
	}
	return nil
}

// HID_Exit finalizes the HIDAPI library.
func HID_Exit() error {
	if C.SDL_hid_exit() != 0 {
		return getError()
	}
	return nil
}

// HID_DeviceChangeCount checks to see if devices may have been added or removed.
func HID_DeviceChangeCount() uint32 {
	return uint32(C.SDL_hid_device_change_count())
}

// HID_Enumerate enumerates the HID devices.
// If vendorID and productID are both 0, all HID devices are returned.
func HID_Enumerate(vendorID, productID uint16) []HIDDeviceInfo {
	head := C.SDL_hid_enumerate(C.ushort(vendorID), C.ushort(productID))
	if head == nil {
		return nil
	}
	defer C.SDL_hid_free_enumeration(head)

	var result []HIDDeviceInfo
	for info := head; info != nil; info = info.next {
		result = append(result, HIDDeviceInfo{
			Path:               C.GoString(info.path),
			VendorID:           uint16(info.vendor_id),
			ProductID:          uint16(info.product_id),
			SerialNumber:       wcharToString(info.serial_number),
			ReleaseNumber:      uint16(info.release_number),
			ManufacturerString: wcharToString(info.manufacturer_string),
			ProductString:      wcharToString(info.product_string),
			UsagePage:          uint16(info.usage_page),
			Usage:              uint16(info.usage),
			InterfaceNumber:    int(info.interface_number),
			InterfaceClass:     int(info.interface_class),
			InterfaceSubclass:  int(info.interface_subclass),
			InterfaceProtocol:  int(info.interface_protocol),
			BusType:            HIDBusType(info.bus_type),
		})
	}
	return result
}

// HID_Open opens a HID device using a Vendor ID, Product ID and optionally a serial number.
// If serialNumber is empty, the first device with the specified VID and PID is opened.
func HID_Open(vendorID, productID uint16, serialNumber string) (*HIDDevice, error) {
	var cserial *C.wchar_t
	if serialNumber != "" {
		// Convert Go string to wchar_t via C string
		cs := C.CString(serialNumber)
		defer C.free(unsafe.Pointer(cs))
		// Allocate wchar_t buffer and convert
		wlen := C.size_t(len(serialNumber) + 1)
		cserial = (*C.wchar_t)(C.malloc(wlen * C.size_t(unsafe.Sizeof(C.wchar_t(0)))))
		defer C.free(unsafe.Pointer(cserial))
		C.mbstowcs(cserial, cs, wlen)
	}
	dev := C.SDL_hid_open(C.ushort(vendorID), C.ushort(productID), cserial)
	if dev == nil {
		return nil, getError()
	}
	return &HIDDevice{c: dev}, nil
}

// HID_OpenPath opens a HID device by its path name.
func HID_OpenPath(path string) (*HIDDevice, error) {
	cp := C.CString(path)
	defer C.free(unsafe.Pointer(cp))
	dev := C.SDL_hid_open_path(cp)
	if dev == nil {
		return nil, getError()
	}
	return &HIDDevice{c: dev}, nil
}

// Properties returns the properties associated with the HID device.
func (d *HIDDevice) Properties() PropertiesID {
	return PropertiesID(C.SDL_hid_get_properties(d.c))
}

// Write writes an Output report to the HID device.
// The first byte of data must contain the Report ID.
// Returns the actual number of bytes written.
func (d *HIDDevice) Write(data []byte) (int, error) {
	n := C.SDL_hid_write(d.c, (*C.uchar)(unsafe.Pointer(&data[0])), C.size_t(len(data)))
	if n < 0 {
		return 0, getError()
	}
	return int(n), nil
}

// Read reads an Input report from the HID device.
// Returns the actual number of bytes read.
func (d *HIDDevice) Read(data []byte) (int, error) {
	n := C.SDL_hid_read(d.c, (*C.uchar)(unsafe.Pointer(&data[0])), C.size_t(len(data)))
	if n < 0 {
		return 0, getError()
	}
	return int(n), nil
}

// ReadTimeout reads an Input report from the HID device with a timeout.
// milliseconds is the timeout in milliseconds, or -1 for blocking wait.
// Returns the actual number of bytes read.
func (d *HIDDevice) ReadTimeout(data []byte, milliseconds int) (int, error) {
	n := C.SDL_hid_read_timeout(d.c, (*C.uchar)(unsafe.Pointer(&data[0])), C.size_t(len(data)), C.int(milliseconds))
	if n < 0 {
		return 0, getError()
	}
	return int(n), nil
}

// SetNonblocking sets the device handle to be non-blocking.
// nonblock: 1 to enable nonblocking, 0 to disable.
func (d *HIDDevice) SetNonblocking(nonblock int) error {
	if C.SDL_hid_set_nonblocking(d.c, C.int(nonblock)) != 0 {
		return getError()
	}
	return nil
}

// SendFeatureReport sends a Feature report to the device.
// The first byte of data must contain the Report ID.
// Returns the actual number of bytes written.
func (d *HIDDevice) SendFeatureReport(data []byte) (int, error) {
	n := C.SDL_hid_send_feature_report(d.c, (*C.uchar)(unsafe.Pointer(&data[0])), C.size_t(len(data)))
	if n < 0 {
		return 0, getError()
	}
	return int(n), nil
}

// GetFeatureReport gets a feature report from the HID device.
// Set the first byte of data to the Report ID of the report to be read.
// Returns the number of bytes read plus one for the report ID.
func (d *HIDDevice) GetFeatureReport(data []byte) (int, error) {
	n := C.SDL_hid_get_feature_report(d.c, (*C.uchar)(unsafe.Pointer(&data[0])), C.size_t(len(data)))
	if n < 0 {
		return 0, getError()
	}
	return int(n), nil
}

// GetInputReport gets an input report from the HID device.
// Set the first byte of data to the Report ID of the report to be read.
// Returns the number of bytes read plus one for the report ID.
func (d *HIDDevice) GetInputReport(data []byte) (int, error) {
	n := C.SDL_hid_get_input_report(d.c, (*C.uchar)(unsafe.Pointer(&data[0])), C.size_t(len(data)))
	if n < 0 {
		return 0, getError()
	}
	return int(n), nil
}

// Close closes the HID device.
func (d *HIDDevice) Close() error {
	if C.SDL_hid_close(d.c) != 0 {
		return getError()
	}
	d.c = nil
	return nil
}

// GetManufacturerString gets the manufacturer string from the HID device.
func (d *HIDDevice) GetManufacturerString() (string, error) {
	var wbuf [256]C.wchar_t
	if C.SDL_hid_get_manufacturer_string(d.c, &wbuf[0], 256) != 0 {
		return "", getError()
	}
	return wcharToString(&wbuf[0]), nil
}

// GetProductString gets the product string from the HID device.
func (d *HIDDevice) GetProductString() (string, error) {
	var wbuf [256]C.wchar_t
	if C.SDL_hid_get_product_string(d.c, &wbuf[0], 256) != 0 {
		return "", getError()
	}
	return wcharToString(&wbuf[0]), nil
}

// GetSerialNumberString gets the serial number string from the HID device.
func (d *HIDDevice) GetSerialNumberString() (string, error) {
	var wbuf [256]C.wchar_t
	if C.SDL_hid_get_serial_number_string(d.c, &wbuf[0], 256) != 0 {
		return "", getError()
	}
	return wcharToString(&wbuf[0]), nil
}

// GetDeviceInfo gets the device info from the HID device.
func (d *HIDDevice) GetDeviceInfo() (*HIDDeviceInfo, error) {
	info := C.SDL_hid_get_device_info(d.c)
	if info == nil {
		return nil, getError()
	}
	return &HIDDeviceInfo{
		Path:               C.GoString(info.path),
		VendorID:           uint16(info.vendor_id),
		ProductID:          uint16(info.product_id),
		SerialNumber:       wcharToString(info.serial_number),
		ReleaseNumber:      uint16(info.release_number),
		ManufacturerString: wcharToString(info.manufacturer_string),
		ProductString:      wcharToString(info.product_string),
		UsagePage:          uint16(info.usage_page),
		Usage:              uint16(info.usage),
		InterfaceNumber:    int(info.interface_number),
		InterfaceClass:     int(info.interface_class),
		InterfaceSubclass:  int(info.interface_subclass),
		InterfaceProtocol:  int(info.interface_protocol),
		BusType:            HIDBusType(info.bus_type),
	}, nil
}

// HID_BLEScan starts or stops a BLE scan for HID devices.
func HID_BLEScan(active bool) {
	C.SDL_hid_ble_scan(C.bool(active))
}

// GetIndexedString returns a string from a HID device by string index.
func (d *HIDDevice) GetIndexedString(stringIndex int) (string, error) {
	buf := make([]C.wchar_t, 256)
	if C.SDL_hid_get_indexed_string(d.c, C.int(stringIndex), &buf[0], 256) < 0 {
		return "", getError()
	}
	return wcharToString(&buf[0]), nil
}

// GetReportDescriptor returns the report descriptor from a HID device.
func (d *HIDDevice) GetReportDescriptor(bufSize int) ([]byte, error) {
	buf := make([]C.uchar, bufSize)
	n := C.SDL_hid_get_report_descriptor(d.c, &buf[0], C.size_t(bufSize))
	if n < 0 {
		return nil, getError()
	}
	result := make([]byte, int(n))
	for i := 0; i < int(n); i++ {
		result[i] = byte(buf[i])
	}
	return result, nil
}
