package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>

static size_t _sdl_ioprintf(SDL_IOStream *ctx, const char *str) {
	return SDL_IOprintf(ctx, "%s", str);
}

extern Sint64 goIOSize(void *userdata);
extern Sint64 goIOSeek(void *userdata, Sint64 offset, SDL_IOWhence whence);
extern size_t goIORead(void *userdata, void *ptr, size_t size, SDL_IOStatus *status);
extern size_t goIOWrite(void *userdata, void *ptr, size_t size, SDL_IOStatus *status);
extern bool   goIOFlush(void *userdata, SDL_IOStatus *status);
extern bool   goIOClose(void *userdata);

static SDL_IOStream *_sdl_open_io(void *userdata) {
	SDL_IOStreamInterface iface;
	SDL_INIT_INTERFACE(&iface);
	iface.size  = goIOSize;
	iface.seek  = goIOSeek;
	iface.read  = goIORead;
	iface.write = (size_t (*)(void *, const void *, size_t, SDL_IOStatus *))goIOWrite;
	iface.flush = goIOFlush;
	iface.close = goIOClose;
	return SDL_OpenIO(&iface, userdata);
}
*/
import "C"

import "unsafe"

// IOStream represents an SDL I/O stream.
type IOStream struct {
	c *C.SDL_IOStream
}

// IOStatus represents the status of an I/O stream.
type IOStatus int

// I/O stream status values.
const (
	IO_STATUS_READY IOStatus = C.SDL_IO_STATUS_READY
	IO_STATUS_ERROR IOStatus = C.SDL_IO_STATUS_ERROR
	IO_STATUS_EOF       IOStatus = C.SDL_IO_STATUS_EOF
	IO_STATUS_NOT_READY IOStatus = C.SDL_IO_STATUS_NOT_READY
	IO_STATUS_READONLY  IOStatus = C.SDL_IO_STATUS_READONLY
	IO_STATUS_WRITEONLY IOStatus = C.SDL_IO_STATUS_WRITEONLY
)

// IOWhence represents the seek reference point.
type IOWhence int

// I/O seek reference points.
const (
	IO_SEEK_SET IOWhence = C.SDL_IO_SEEK_SET
	IO_SEEK_CUR IOWhence = C.SDL_IO_SEEK_CUR
	IO_SEEK_END IOWhence = C.SDL_IO_SEEK_END
)

// IOFromFile creates an I/O stream from a file.
func IOFromFile(file, mode string) (*IOStream, error) {
	cf := C.CString(file)
	cm := C.CString(mode)
	defer C.free(unsafe.Pointer(cf))
	defer C.free(unsafe.Pointer(cm))
	cs := C.SDL_IOFromFile(cf, cm)
	if cs == nil {
		return nil, getError()
	}
	return &IOStream{c: cs}, nil
}

// IOFromMem creates an I/O stream from a memory buffer.
func IOFromMem(mem unsafe.Pointer, size int) (*IOStream, error) {
	cs := C.SDL_IOFromMem(mem, C.size_t(size))
	if cs == nil {
		return nil, getError()
	}
	return &IOStream{c: cs}, nil
}

// IOFromConstMem creates a read-only I/O stream from a memory buffer.
func IOFromConstMem(mem unsafe.Pointer, size int) (*IOStream, error) {
	cs := C.SDL_IOFromConstMem(mem, C.size_t(size))
	if cs == nil {
		return nil, getError()
	}
	return &IOStream{c: cs}, nil
}

// IOFromDynamicMem creates an I/O stream that grows as data is written.
func IOFromDynamicMem() (*IOStream, error) {
	cs := C.SDL_IOFromDynamicMem()
	if cs == nil {
		return nil, getError()
	}
	return &IOStream{c: cs}, nil
}

// Close closes the I/O stream.
func (s *IOStream) Close() error {
	if s.c != nil {
		if !C.SDL_CloseIO(s.c) {
			return getError()
		}
		s.c = nil
	}
	return nil
}

// Status returns the current status of the stream.
func (s *IOStream) Status() IOStatus {
	return IOStatus(C.SDL_GetIOStatus(s.c))
}

// Properties returns the properties of the stream.
func (s *IOStream) Properties() PropertiesID {
	return PropertiesID(C.SDL_GetIOProperties(s.c))
}

// SeekIO seeks to a position in the stream.
func (s *IOStream) SeekIO(offset int64, whence IOWhence) (int64, error) {
	r := C.SDL_SeekIO(s.c, C.Sint64(offset), C.SDL_IOWhence(whence))
	if r < 0 {
		return 0, getError()
	}
	return int64(r), nil
}

// Tell returns the current position in the stream.
func (s *IOStream) Tell() int64 {
	return int64(C.SDL_TellIO(s.c))
}

// Read reads data from the stream.
func (s *IOStream) Read(buf []byte) (int, error) {
	if len(buf) == 0 {
		return 0, nil
	}
	n := C.SDL_ReadIO(s.c, unsafe.Pointer(&buf[0]), C.size_t(len(buf)))
	if n == 0 && s.Status() == IO_STATUS_ERROR {
		return 0, getError()
	}
	return int(n), nil
}

// Write writes data to the stream.
func (s *IOStream) Write(buf []byte) (int, error) {
	if len(buf) == 0 {
		return 0, nil
	}
	n := C.SDL_WriteIO(s.c, unsafe.Pointer(&buf[0]), C.size_t(len(buf)))
	if n < C.size_t(len(buf)) {
		return int(n), getError()
	}
	return int(n), nil
}

// Flush flushes any buffered data.
func (s *IOStream) Flush() error {
	if !C.SDL_FlushIO(s.c) {
		return getError()
	}
	return nil
}

// LoadFile loads an entire file into memory. Caller must handle the returned bytes.
func LoadFile(file string) ([]byte, error) {
	cf := C.CString(file)
	defer C.free(unsafe.Pointer(cf))
	var size C.size_t
	data := C.SDL_LoadFile(cf, &size)
	if data == nil {
		return nil, getError()
	}
	defer C.SDL_free(data)
	result := make([]byte, int(size))
	copy(result, unsafe.Slice((*byte)(data), int(size)))
	return result, nil
}

// SaveFile saves data to a file.
func SaveFile(file string, data []byte) error {
	cf := C.CString(file)
	defer C.free(unsafe.Pointer(cf))
	var ptr unsafe.Pointer
	if len(data) > 0 {
		ptr = unsafe.Pointer(&data[0])
	}
	if !C.SDL_SaveFile(cf, ptr, C.size_t(len(data))) {
		return getError()
	}
	return nil
}

// GetSize returns the size of the data stream, or -1 if unknown.
func (s *IOStream) GetSize() int64 {
	return int64(C.SDL_GetIOSize(s.c))
}

// LoadAll loads all remaining data from the I/O stream into memory.
func (s *IOStream) LoadAll() ([]byte, error) {
	var size C.size_t
	data := C.SDL_LoadFile_IO(s.c, &size, C.bool(false))
	if data == nil {
		return nil, getError()
	}
	defer C.SDL_free(data)
	result := make([]byte, int(size))
	copy(result, unsafe.Slice((*byte)(data), int(size)))
	return result, nil
}

// SaveAll writes all data to the I/O stream.
func (s *IOStream) SaveAll(data []byte) error {
	var ptr unsafe.Pointer
	if len(data) > 0 {
		ptr = unsafe.Pointer(&data[0])
	}
	if !C.SDL_SaveFile_IO(s.c, ptr, C.size_t(len(data)), C.bool(false)) {
		return getError()
	}
	return nil
}

// ---------------------------------------------------------------------------
// Binary read operations
// ---------------------------------------------------------------------------

// ReadU8 reads a byte from the I/O stream.
func (s *IOStream) ReadU8() (uint8, error) {
	var v C.Uint8
	if !C.SDL_ReadU8(s.c, &v) {
		return 0, getError()
	}
	return uint8(v), nil
}

// ReadS8 reads a signed byte from the I/O stream.
func (s *IOStream) ReadS8() (int8, error) {
	var v C.Sint8
	if !C.SDL_ReadS8(s.c, &v) {
		return 0, getError()
	}
	return int8(v), nil
}

// ReadU16LE reads a little-endian 16-bit unsigned integer from the I/O stream.
func (s *IOStream) ReadU16LE() (uint16, error) {
	var v C.Uint16
	if !C.SDL_ReadU16LE(s.c, &v) {
		return 0, getError()
	}
	return uint16(v), nil
}

// ReadU16BE reads a big-endian 16-bit unsigned integer from the I/O stream.
func (s *IOStream) ReadU16BE() (uint16, error) {
	var v C.Uint16
	if !C.SDL_ReadU16BE(s.c, &v) {
		return 0, getError()
	}
	return uint16(v), nil
}

// ReadS16LE reads a little-endian 16-bit signed integer from the I/O stream.
func (s *IOStream) ReadS16LE() (int16, error) {
	var v C.Sint16
	if !C.SDL_ReadS16LE(s.c, &v) {
		return 0, getError()
	}
	return int16(v), nil
}

// ReadS16BE reads a big-endian 16-bit signed integer from the I/O stream.
func (s *IOStream) ReadS16BE() (int16, error) {
	var v C.Sint16
	if !C.SDL_ReadS16BE(s.c, &v) {
		return 0, getError()
	}
	return int16(v), nil
}

// ReadU32LE reads a little-endian 32-bit unsigned integer from the I/O stream.
func (s *IOStream) ReadU32LE() (uint32, error) {
	var v C.Uint32
	if !C.SDL_ReadU32LE(s.c, &v) {
		return 0, getError()
	}
	return uint32(v), nil
}

// ReadU32BE reads a big-endian 32-bit unsigned integer from the I/O stream.
func (s *IOStream) ReadU32BE() (uint32, error) {
	var v C.Uint32
	if !C.SDL_ReadU32BE(s.c, &v) {
		return 0, getError()
	}
	return uint32(v), nil
}

// ReadS32LE reads a little-endian 32-bit signed integer from the I/O stream.
func (s *IOStream) ReadS32LE() (int32, error) {
	var v C.Sint32
	if !C.SDL_ReadS32LE(s.c, &v) {
		return 0, getError()
	}
	return int32(v), nil
}

// ReadS32BE reads a big-endian 32-bit signed integer from the I/O stream.
func (s *IOStream) ReadS32BE() (int32, error) {
	var v C.Sint32
	if !C.SDL_ReadS32BE(s.c, &v) {
		return 0, getError()
	}
	return int32(v), nil
}

// ReadU64LE reads a little-endian 64-bit unsigned integer from the I/O stream.
func (s *IOStream) ReadU64LE() (uint64, error) {
	var v C.Uint64
	if !C.SDL_ReadU64LE(s.c, &v) {
		return 0, getError()
	}
	return uint64(v), nil
}

// ReadU64BE reads a big-endian 64-bit unsigned integer from the I/O stream.
func (s *IOStream) ReadU64BE() (uint64, error) {
	var v C.Uint64
	if !C.SDL_ReadU64BE(s.c, &v) {
		return 0, getError()
	}
	return uint64(v), nil
}

// ReadS64LE reads a little-endian 64-bit signed integer from the I/O stream.
func (s *IOStream) ReadS64LE() (int64, error) {
	var v C.Sint64
	if !C.SDL_ReadS64LE(s.c, &v) {
		return 0, getError()
	}
	return int64(v), nil
}

// ReadS64BE reads a big-endian 64-bit signed integer from the I/O stream.
func (s *IOStream) ReadS64BE() (int64, error) {
	var v C.Sint64
	if !C.SDL_ReadS64BE(s.c, &v) {
		return 0, getError()
	}
	return int64(v), nil
}

// ---------------------------------------------------------------------------
// Binary write operations
// ---------------------------------------------------------------------------

// WriteU8 writes a byte to the I/O stream.
func (s *IOStream) WriteU8(v uint8) error {
	if !C.SDL_WriteU8(s.c, C.Uint8(v)) {
		return getError()
	}
	return nil
}

// WriteS8 writes a signed byte to the I/O stream.
func (s *IOStream) WriteS8(v int8) error {
	if !C.SDL_WriteS8(s.c, C.Sint8(v)) {
		return getError()
	}
	return nil
}

// WriteU16LE writes a little-endian 16-bit unsigned integer to the I/O stream.
func (s *IOStream) WriteU16LE(v uint16) error {
	if !C.SDL_WriteU16LE(s.c, C.Uint16(v)) {
		return getError()
	}
	return nil
}

// WriteU16BE writes a big-endian 16-bit unsigned integer to the I/O stream.
func (s *IOStream) WriteU16BE(v uint16) error {
	if !C.SDL_WriteU16BE(s.c, C.Uint16(v)) {
		return getError()
	}
	return nil
}

// WriteS16LE writes a little-endian 16-bit signed integer to the I/O stream.
func (s *IOStream) WriteS16LE(v int16) error {
	if !C.SDL_WriteS16LE(s.c, C.Sint16(v)) {
		return getError()
	}
	return nil
}

// WriteS16BE writes a big-endian 16-bit signed integer to the I/O stream.
func (s *IOStream) WriteS16BE(v int16) error {
	if !C.SDL_WriteS16BE(s.c, C.Sint16(v)) {
		return getError()
	}
	return nil
}

// WriteU32LE writes a little-endian 32-bit unsigned integer to the I/O stream.
func (s *IOStream) WriteU32LE(v uint32) error {
	if !C.SDL_WriteU32LE(s.c, C.Uint32(v)) {
		return getError()
	}
	return nil
}

// WriteU32BE writes a big-endian 32-bit unsigned integer to the I/O stream.
func (s *IOStream) WriteU32BE(v uint32) error {
	if !C.SDL_WriteU32BE(s.c, C.Uint32(v)) {
		return getError()
	}
	return nil
}

// WriteS32LE writes a little-endian 32-bit signed integer to the I/O stream.
func (s *IOStream) WriteS32LE(v int32) error {
	if !C.SDL_WriteS32LE(s.c, C.Sint32(v)) {
		return getError()
	}
	return nil
}

// WriteS32BE writes a big-endian 32-bit signed integer to the I/O stream.
func (s *IOStream) WriteS32BE(v int32) error {
	if !C.SDL_WriteS32BE(s.c, C.Sint32(v)) {
		return getError()
	}
	return nil
}

// WriteU64LE writes a little-endian 64-bit unsigned integer to the I/O stream.
func (s *IOStream) WriteU64LE(v uint64) error {
	if !C.SDL_WriteU64LE(s.c, C.Uint64(v)) {
		return getError()
	}
	return nil
}

// WriteU64BE writes a big-endian 64-bit unsigned integer to the I/O stream.
func (s *IOStream) WriteU64BE(v uint64) error {
	if !C.SDL_WriteU64BE(s.c, C.Uint64(v)) {
		return getError()
	}
	return nil
}

// WriteS64LE writes a little-endian 64-bit signed integer to the I/O stream.
func (s *IOStream) WriteS64LE(v int64) error {
	if !C.SDL_WriteS64LE(s.c, C.Sint64(v)) {
		return getError()
	}
	return nil
}

// WriteS64BE writes a big-endian 64-bit signed integer to the I/O stream.
func (s *IOStream) WriteS64BE(v int64) error {
	if !C.SDL_WriteS64BE(s.c, C.Sint64(v)) {
		return getError()
	}
	return nil
}

// IOStreamImpl is the interface to implement for custom I/O streams.
type IOStreamImpl interface {
	Size() int64
	SeekIO(offset int64, whence IOWhence) int64
	Read(buf []byte) (int, IOStatus)
	Write(buf []byte) (int, IOStatus)
	Flush() (bool, IOStatus)
	Close() bool
}

//export goIOSize
func goIOSize(userdata unsafe.Pointer) C.Sint64 {
	impl := getCallback(uintptr(userdata)).(IOStreamImpl)
	return C.Sint64(impl.Size())
}

//export goIOSeek
func goIOSeek(userdata unsafe.Pointer, offset C.Sint64, whence C.SDL_IOWhence) C.Sint64 {
	impl := getCallback(uintptr(userdata)).(IOStreamImpl)
	return C.Sint64(impl.SeekIO(int64(offset), IOWhence(whence)))
}

//export goIORead
func goIORead(userdata unsafe.Pointer, ptr unsafe.Pointer, size C.size_t, status *C.SDL_IOStatus) C.size_t {
	impl := getCallback(uintptr(userdata)).(IOStreamImpl)
	buf := unsafe.Slice((*byte)(ptr), int(size))
	n, st := impl.Read(buf)
	*status = C.SDL_IOStatus(st)
	return C.size_t(n)
}

//export goIOWrite
func goIOWrite(userdata unsafe.Pointer, ptr unsafe.Pointer, size C.size_t, status *C.SDL_IOStatus) C.size_t {
	impl := getCallback(uintptr(userdata)).(IOStreamImpl)
	buf := unsafe.Slice((*byte)(ptr), int(size))
	n, st := impl.Write(buf)
	*status = C.SDL_IOStatus(st)
	return C.size_t(n)
}

//export goIOFlush
func goIOFlush(userdata unsafe.Pointer, status *C.SDL_IOStatus) C.bool {
	impl := getCallback(uintptr(userdata)).(IOStreamImpl)
	ok, st := impl.Flush()
	*status = C.SDL_IOStatus(st)
	return C.bool(ok)
}

//export goIOClose
func goIOClose(userdata unsafe.Pointer) C.bool {
	impl := getCallback(uintptr(userdata)).(IOStreamImpl)
	ok := impl.Close()
	unregisterCallback(uintptr(userdata))
	return C.bool(ok)
}

// OpenIO creates a custom I/O stream backed by a Go implementation.
func OpenIO(impl IOStreamImpl) (*IOStream, error) {
	id := registerCallback(impl)
	cs := C._sdl_open_io(ptrFromID(id))
	if cs == nil {
		unregisterCallback(id)
		return nil, getError()
	}
	return &IOStream{c: cs}, nil
}

// Printf writes a formatted string to the I/O stream.
// Go format arguments should be resolved before calling (use fmt.Sprintf).
func (s *IOStream) Printf(str string) (int, error) {
	cs := C.CString(str)
	defer C.free(unsafe.Pointer(cs))
	n := C._sdl_ioprintf(s.c, cs)
	if n == 0 && s.Status() == IO_STATUS_ERROR {
		return 0, getError()
	}
	return int(n), nil
}

// Note: SDL_IOvprintf is not wrapped because it takes a C va_list argument
// which has no equivalent in Go. Use Printf instead.
