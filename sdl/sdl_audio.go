package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>

extern void goAudioStreamCallback(void *userdata, SDL_AudioStream *stream, int additional_amount, int total_amount);
extern void goAudioPostmixCallback(void *userdata, SDL_AudioSpec *spec, float *buffer, int buflen);
extern void goAudioStreamDataCompleteCallback(void *userdata, void *buf, int buflen);

static void cgoAudioStreamCallbackTrampoline(void *userdata, SDL_AudioStream *stream, int additional_amount, int total_amount) {
	goAudioStreamCallback(userdata, stream, additional_amount, total_amount);
}

static void cgoAudioPostmixCallbackTrampoline(void *userdata, const SDL_AudioSpec *spec, float *buffer, int buflen) {
	goAudioPostmixCallback(userdata, (SDL_AudioSpec *)spec, buffer, buflen);
}

static void cgoAudioStreamDataCompleteTrampoline(void *userdata, const void *buf, int buflen) {
	goAudioStreamDataCompleteCallback(userdata, (void *)buf, buflen);
}

static SDL_AudioStreamCallback _get_audio_stream_callback_trampoline(void) {
	return cgoAudioStreamCallbackTrampoline;
}
*/
import "C"

import "unsafe"

// AudioDeviceID is a unique identifier for an audio device.
type AudioDeviceID uint32

// Default audio device IDs for playback and recording.
const (
	AUDIO_DEVICE_DEFAULT_PLAYBACK  AudioDeviceID = 0xFFFFFFFF
	AUDIO_DEVICE_DEFAULT_RECORDING AudioDeviceID = 0xFFFFFFFE
)

// AudioFormat represents an audio sample format.
type AudioFormat uint32

// Audio sample format constants.
const (
	AUDIO_UNKNOWN AudioFormat = C.SDL_AUDIO_UNKNOWN
	AUDIO_U8      AudioFormat = C.SDL_AUDIO_U8
	AUDIO_S8      AudioFormat = C.SDL_AUDIO_S8
	AUDIO_S16LE   AudioFormat = C.SDL_AUDIO_S16LE
	AUDIO_S16BE   AudioFormat = C.SDL_AUDIO_S16BE
	AUDIO_S32LE   AudioFormat = C.SDL_AUDIO_S32LE
	AUDIO_S32BE   AudioFormat = C.SDL_AUDIO_S32BE
	AUDIO_F32LE   AudioFormat = C.SDL_AUDIO_F32LE
	AUDIO_F32BE   AudioFormat = C.SDL_AUDIO_F32BE
	AUDIO_S16     AudioFormat = C.SDL_AUDIO_S16
	AUDIO_S32     AudioFormat = C.SDL_AUDIO_S32
	AUDIO_F32     AudioFormat = C.SDL_AUDIO_F32
)

// AudioSpec describes an audio format.
type AudioSpec struct {
	Format   AudioFormat
	Channels int32
	Freq     int32
}

func (s *AudioSpec) cptr() *C.SDL_AudioSpec {
	return (*C.SDL_AudioSpec)(unsafe.Pointer(s))
}

// AudioStream represents an audio data stream.
type AudioStream struct {
	c *C.SDL_AudioStream
}

// --- Driver functions ---

// GetNumAudioDrivers returns the number of built-in audio drivers.
func GetNumAudioDrivers() int {
	return int(C.SDL_GetNumAudioDrivers())
}

// GetAudioDriver returns the name of a built-in audio driver.
func GetAudioDriver(index int) string {
	return C.GoString(C.SDL_GetAudioDriver(C.int(index)))
}

// GetCurrentAudioDriver returns the name of the current audio driver.
func GetCurrentAudioDriver() string {
	return C.GoString(C.SDL_GetCurrentAudioDriver())
}

// --- Device enumeration ---

// GetAudioPlaybackDevices returns a list of playback audio device IDs.
func GetAudioPlaybackDevices() []AudioDeviceID {
	var count C.int
	cids := C.SDL_GetAudioPlaybackDevices(&count)
	if cids == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cids))
	n := int(count)
	result := make([]AudioDeviceID, n)
	slice := unsafe.Slice((*C.SDL_AudioDeviceID)(cids), n)
	for i, id := range slice {
		result[i] = AudioDeviceID(id)
	}
	return result
}

// GetAudioRecordingDevices returns a list of recording audio device IDs.
func GetAudioRecordingDevices() []AudioDeviceID {
	var count C.int
	cids := C.SDL_GetAudioRecordingDevices(&count)
	if cids == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cids))
	n := int(count)
	result := make([]AudioDeviceID, n)
	slice := unsafe.Slice((*C.SDL_AudioDeviceID)(cids), n)
	for i, id := range slice {
		result[i] = AudioDeviceID(id)
	}
	return result
}

// GetAudioDeviceName returns the name of an audio device.
func GetAudioDeviceName(devid AudioDeviceID) string {
	return C.GoString(C.SDL_GetAudioDeviceName(C.SDL_AudioDeviceID(devid)))
}

// GetAudioDeviceFormat returns the format of an audio device.
func GetAudioDeviceFormat(devid AudioDeviceID) (AudioSpec, int, error) {
	var spec AudioSpec
	var sampleFrames C.int
	if !C.SDL_GetAudioDeviceFormat(C.SDL_AudioDeviceID(devid), spec.cptr(), &sampleFrames) {
		return spec, 0, getError()
	}
	return spec, int(sampleFrames), nil
}

// --- Device management ---

// OpenAudioDevice opens an audio device.
func OpenAudioDevice(devid AudioDeviceID, spec *AudioSpec) (AudioDeviceID, error) {
	var cs *C.SDL_AudioSpec
	if spec != nil {
		cs = spec.cptr()
	}
	id := C.SDL_OpenAudioDevice(C.SDL_AudioDeviceID(devid), cs)
	if id == 0 {
		return 0, getError()
	}
	return AudioDeviceID(id), nil
}

// CloseAudioDevice closes an audio device.
func CloseAudioDevice(devid AudioDeviceID) {
	C.SDL_CloseAudioDevice(C.SDL_AudioDeviceID(devid))
}

// PauseAudioDevice pauses an audio device.
func PauseAudioDevice(devid AudioDeviceID) error {
	if !C.SDL_PauseAudioDevice(C.SDL_AudioDeviceID(devid)) {
		return getError()
	}
	return nil
}

// ResumeAudioDevice resumes an audio device.
func ResumeAudioDevice(devid AudioDeviceID) error {
	if !C.SDL_ResumeAudioDevice(C.SDL_AudioDeviceID(devid)) {
		return getError()
	}
	return nil
}

// AudioDevicePaused returns true if the audio device is paused.
func AudioDevicePaused(devid AudioDeviceID) bool {
	return bool(C.SDL_AudioDevicePaused(C.SDL_AudioDeviceID(devid)))
}

// GetAudioDeviceGain returns the gain of an audio device.
func GetAudioDeviceGain(devid AudioDeviceID) float32 {
	return float32(C.SDL_GetAudioDeviceGain(C.SDL_AudioDeviceID(devid)))
}

// SetAudioDeviceGain sets the gain of an audio device.
func SetAudioDeviceGain(devid AudioDeviceID, gain float32) error {
	if !C.SDL_SetAudioDeviceGain(C.SDL_AudioDeviceID(devid), C.float(gain)) {
		return getError()
	}
	return nil
}

// --- Stream creation ---

// CreateAudioStream creates a new audio stream for converting audio data.
func CreateAudioStream(srcSpec, dstSpec *AudioSpec) (*AudioStream, error) {
	cs := C.SDL_CreateAudioStream(srcSpec.cptr(), dstSpec.cptr())
	if cs == nil {
		return nil, getError()
	}
	return &AudioStream{c: cs}, nil
}

// Destroy destroys the audio stream.
func (s *AudioStream) Destroy() {
	if s.c != nil {
		C.SDL_DestroyAudioStream(s.c)
		s.c = nil
	}
}

// --- Stream format ---

// GetFormat returns the current input and output format of the stream.
func (s *AudioStream) GetFormat() (src, dst AudioSpec, err error) {
	if !C.SDL_GetAudioStreamFormat(s.c, src.cptr(), dst.cptr()) {
		return src, dst, getError()
	}
	return src, dst, nil
}

// SetFormat sets the input and/or output format of the stream.
func (s *AudioStream) SetFormat(srcSpec, dstSpec *AudioSpec) error {
	var cs, cd *C.SDL_AudioSpec
	if srcSpec != nil {
		cs = srcSpec.cptr()
	}
	if dstSpec != nil {
		cd = dstSpec.cptr()
	}
	if !C.SDL_SetAudioStreamFormat(s.c, cs, cd) {
		return getError()
	}
	return nil
}

// GetFrequencyRatio returns the frequency ratio of the stream.
func (s *AudioStream) GetFrequencyRatio() float32 {
	return float32(C.SDL_GetAudioStreamFrequencyRatio(s.c))
}

// SetFrequencyRatio sets the frequency ratio of the stream.
func (s *AudioStream) SetFrequencyRatio(ratio float32) error {
	if !C.SDL_SetAudioStreamFrequencyRatio(s.c, C.float(ratio)) {
		return getError()
	}
	return nil
}

// GetGain returns the gain of the stream.
func (s *AudioStream) GetGain() float32 {
	return float32(C.SDL_GetAudioStreamGain(s.c))
}

// SetGain sets the gain of the stream.
func (s *AudioStream) SetGain(gain float32) error {
	if !C.SDL_SetAudioStreamGain(s.c, C.float(gain)) {
		return getError()
	}
	return nil
}

// --- Stream data ---

// PutData queues audio data for conversion/output.
func (s *AudioStream) PutData(buf unsafe.Pointer, length int) error {
	if !C.SDL_PutAudioStreamData(s.c, buf, C.int(length)) {
		return getError()
	}
	return nil
}

// PutDataBytes queues audio data from a byte slice.
func (s *AudioStream) PutDataBytes(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	return s.PutData(unsafe.Pointer(&data[0]), len(data))
}

// GetData gets converted audio data from the stream.
func (s *AudioStream) GetData(buf unsafe.Pointer, length int) (int, error) {
	n := C.SDL_GetAudioStreamData(s.c, buf, C.int(length))
	if n < 0 {
		return 0, getError()
	}
	return int(n), nil
}

// Available returns the number of bytes available for reading.
func (s *AudioStream) Available() int {
	return int(C.SDL_GetAudioStreamAvailable(s.c))
}

// Queued returns the number of bytes queued for output.
func (s *AudioStream) Queued() int {
	return int(C.SDL_GetAudioStreamQueued(s.c))
}

// Flush flushes any pending data in the stream.
func (s *AudioStream) Flush() error {
	if !C.SDL_FlushAudioStream(s.c) {
		return getError()
	}
	return nil
}

// Clear clears all data from the stream.
func (s *AudioStream) Clear() error {
	if !C.SDL_ClearAudioStream(s.c) {
		return getError()
	}
	return nil
}

// Lock locks the stream for thread-safe access.
func (s *AudioStream) Lock() error {
	if !C.SDL_LockAudioStream(s.c) {
		return getError()
	}
	return nil
}

// Unlock unlocks the stream.
func (s *AudioStream) Unlock() error {
	if !C.SDL_UnlockAudioStream(s.c) {
		return getError()
	}
	return nil
}

// --- Stream binding ---

// BindToDevice binds the stream to an audio device.
func (s *AudioStream) BindToDevice(devid AudioDeviceID) error {
	if !C.SDL_BindAudioStream(C.SDL_AudioDeviceID(devid), s.c) {
		return getError()
	}
	return nil
}

// Unbind unbinds the stream from its audio device.
func (s *AudioStream) Unbind() {
	C.SDL_UnbindAudioStream(s.c)
}

// Device returns the device the stream is bound to.
func (s *AudioStream) Device() AudioDeviceID {
	return AudioDeviceID(C.SDL_GetAudioStreamDevice(s.c))
}

// PauseDevice pauses the device associated with the stream.
func (s *AudioStream) PauseDevice() error {
	if !C.SDL_PauseAudioStreamDevice(s.c) {
		return getError()
	}
	return nil
}

// ResumeDevice resumes the device associated with the stream.
func (s *AudioStream) ResumeDevice() error {
	if !C.SDL_ResumeAudioStreamDevice(s.c) {
		return getError()
	}
	return nil
}

// --- WAV loading ---

// LoadWAV loads a WAV file. Returns the audio spec, audio buffer, and length.
// The caller must free the buffer with SDL_free via FreeWAVBuffer.
func LoadWAV(path string) (AudioSpec, []byte, error) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	var spec AudioSpec
	var audioBuf *C.Uint8
	var audioLen C.Uint32
	if !C.SDL_LoadWAV(cpath, spec.cptr(), &audioBuf, &audioLen) {
		return spec, nil, getError()
	}
	length := int(audioLen)
	data := make([]byte, length)
	copy(data, unsafe.Slice((*byte)(unsafe.Pointer(audioBuf)), length))
	C.SDL_free(unsafe.Pointer(audioBuf))
	return spec, data, nil
}

// --- Utility ---

// MixAudio mixes audio data.
func MixAudio(dst, src []byte, format AudioFormat, volume float32) error {
	if len(dst) == 0 || len(src) == 0 {
		return nil
	}
	if !C.SDL_MixAudio((*C.Uint8)(unsafe.Pointer(&dst[0])), (*C.Uint8)(unsafe.Pointer(&src[0])),
		C.SDL_AudioFormat(format), C.Uint32(len(src)), C.float(volume)) {
		return getError()
	}
	return nil
}

// GetAudioFormatName returns a human-readable name for an audio format.
func GetAudioFormatName(format AudioFormat) string {
	return C.GoString(C.SDL_GetAudioFormatName(C.SDL_AudioFormat(format)))
}

// IsAudioDevicePhysical returns true if the audio device is a physical device (not logical).
func IsAudioDevicePhysical(devid AudioDeviceID) bool {
	return bool(C.SDL_IsAudioDevicePhysical(C.SDL_AudioDeviceID(devid)))
}

// IsAudioDevicePlayback returns true if the audio device is a playback device.
func IsAudioDevicePlayback(devid AudioDeviceID) bool {
	return bool(C.SDL_IsAudioDevicePlayback(C.SDL_AudioDeviceID(devid)))
}

// GetAudioDeviceChannelMap returns the current channel map of an audio device.
func GetAudioDeviceChannelMap(devid AudioDeviceID) []int {
	var count C.int
	cmap := C.SDL_GetAudioDeviceChannelMap(C.SDL_AudioDeviceID(devid), &count)
	if cmap == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cmap))
	n := int(count)
	result := make([]int, n)
	slice := unsafe.Slice((*C.int)(cmap), n)
	for i, v := range slice {
		result[i] = int(v)
	}
	return result
}

// GetInputChannelMap returns the current input channel map of the audio stream.
func (s *AudioStream) GetInputChannelMap() []int {
	var count C.int
	cmap := C.SDL_GetAudioStreamInputChannelMap(s.c, &count)
	if cmap == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cmap))
	n := int(count)
	result := make([]int, n)
	slice := unsafe.Slice((*C.int)(cmap), n)
	for i, v := range slice {
		result[i] = int(v)
	}
	return result
}

// GetOutputChannelMap returns the current output channel map of the audio stream.
func (s *AudioStream) GetOutputChannelMap() []int {
	var count C.int
	cmap := C.SDL_GetAudioStreamOutputChannelMap(s.c, &count)
	if cmap == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cmap))
	n := int(count)
	result := make([]int, n)
	slice := unsafe.Slice((*C.int)(cmap), n)
	for i, v := range slice {
		result[i] = int(v)
	}
	return result
}

// SetInputChannelMap sets the current input channel map of the audio stream.
func (s *AudioStream) SetInputChannelMap(chmap []int) error {
	var cmap *C.int
	if len(chmap) > 0 {
		tmp := make([]C.int, len(chmap))
		for i, v := range chmap {
			tmp[i] = C.int(v)
		}
		cmap = &tmp[0]
	}
	if !C.SDL_SetAudioStreamInputChannelMap(s.c, cmap, C.int(len(chmap))) {
		return getError()
	}
	return nil
}

// SetOutputChannelMap sets the current output channel map of the audio stream.
func (s *AudioStream) SetOutputChannelMap(chmap []int) error {
	var cmap *C.int
	if len(chmap) > 0 {
		tmp := make([]C.int, len(chmap))
		for i, v := range chmap {
			tmp[i] = C.int(v)
		}
		cmap = &tmp[0]
	}
	if !C.SDL_SetAudioStreamOutputChannelMap(s.c, cmap, C.int(len(chmap))) {
		return getError()
	}
	return nil
}

// ConvertAudioSamples converts audio samples from one format to another.
func ConvertAudioSamples(srcSpec *AudioSpec, srcData []byte, dstSpec *AudioSpec) ([]byte, error) {
	var dstBuf *C.Uint8
	var dstLen C.int
	var srcPtr *C.Uint8
	if len(srcData) > 0 {
		srcPtr = (*C.Uint8)(unsafe.Pointer(&srcData[0]))
	}
	if !C.SDL_ConvertAudioSamples(srcSpec.cptr(), srcPtr, C.int(len(srcData)), dstSpec.cptr(), &dstBuf, &dstLen) {
		return nil, getError()
	}
	length := int(dstLen)
	data := make([]byte, length)
	copy(data, unsafe.Slice((*byte)(unsafe.Pointer(dstBuf)), length))
	C.SDL_free(unsafe.Pointer(dstBuf))
	return data, nil
}

// GetSilenceValueForFormat returns the silence value for an audio format.
func GetSilenceValueForFormat(format AudioFormat) int {
	return int(C.SDL_GetSilenceValueForFormat(C.SDL_AudioFormat(format)))
}

// Properties returns the properties associated with the audio stream.
func (s *AudioStream) Properties() PropertiesID {
	return PropertiesID(C.SDL_GetAudioStreamProperties(s.c))
}

// DevicePaused returns true if the device bound to the audio stream is paused.
func (s *AudioStream) DevicePaused() bool {
	return bool(C.SDL_AudioStreamDevicePaused(s.c))
}

// LoadWAV_IO loads a WAV from an IOStream. Returns the audio spec, audio buffer, and length.
func LoadWAV_IO(src *IOStream, closeio bool) (AudioSpec, []byte, error) {
	var spec AudioSpec
	var audioBuf *C.Uint8
	var audioLen C.Uint32
	if !C.SDL_LoadWAV_IO(src.c, C.bool(closeio), spec.cptr(), &audioBuf, &audioLen) {
		return spec, nil, getError()
	}
	length := int(audioLen)
	data := make([]byte, length)
	copy(data, unsafe.Slice((*byte)(unsafe.Pointer(audioBuf)), length))
	C.SDL_free(unsafe.Pointer(audioBuf))
	if closeio {
		src.c = nil
	}
	return spec, data, nil
}

// BindAudioStreams binds multiple audio streams to a device at once.
func BindAudioStreams(devid AudioDeviceID, streams []*AudioStream) error {
	if len(streams) == 0 {
		return nil
	}
	cstreams := make([]*C.SDL_AudioStream, len(streams))
	for i, s := range streams {
		cstreams[i] = s.c
	}
	if !C.SDL_BindAudioStreams(C.SDL_AudioDeviceID(devid), &cstreams[0], C.int(len(streams))) {
		return getError()
	}
	return nil
}

// UnbindAudioStreams unbinds multiple audio streams at once.
func UnbindAudioStreams(streams []*AudioStream) {
	if len(streams) == 0 {
		return
	}
	cstreams := make([]*C.SDL_AudioStream, len(streams))
	for i, s := range streams {
		cstreams[i] = s.c
	}
	C.SDL_UnbindAudioStreams(&cstreams[0], C.int(len(streams)))
}

// PutPlanarData queues planar audio data into the stream.
func (s *AudioStream) PutPlanarData(channelBuffers []unsafe.Pointer, numChannels, numSamples int) error {
	if len(channelBuffers) == 0 {
		return nil
	}
	if !C.SDL_PutAudioStreamPlanarData(s.c, (*unsafe.Pointer)(unsafe.Pointer(&channelBuffers[0])), C.int(numChannels), C.int(numSamples)) {
		return getError()
	}
	return nil
}

// AudioStreamCallbackFunc is called when an audio stream needs or receives data.
type AudioStreamCallbackFunc func(stream *AudioStream, additionalAmount, totalAmount int)

// AudioPostmixCallbackFunc is called after audio mixing with access to the final buffer.
type AudioPostmixCallbackFunc func(spec *AudioSpec, buffer []float32)

// AudioStreamDataCompleteCallbackFunc is called when no-copy audio data has been consumed.
type AudioStreamDataCompleteCallbackFunc func(buf unsafe.Pointer, buflen int)

//export goAudioStreamCallback
func goAudioStreamCallback(userdata unsafe.Pointer, stream *C.SDL_AudioStream, additionalAmount, totalAmount C.int) {
	id := uintptr(userdata)
	fn := getCallback(id).(AudioStreamCallbackFunc)
	fn(&AudioStream{c: stream}, int(additionalAmount), int(totalAmount))
}

//export goAudioPostmixCallback
func goAudioPostmixCallback(userdata unsafe.Pointer, spec *C.SDL_AudioSpec, buffer *C.float, buflen C.int) {
	id := uintptr(userdata)
	fn := getCallback(id).(AudioPostmixCallbackFunc)
	goSpec := &AudioSpec{Format: AudioFormat(spec.format), Channels: int32(spec.channels), Freq: int32(spec.freq)}
	goBuf := unsafe.Slice((*float32)(unsafe.Pointer(buffer)), int(buflen))
	fn(goSpec, goBuf)
}

//export goAudioStreamDataCompleteCallback
func goAudioStreamDataCompleteCallback(userdata unsafe.Pointer, buf unsafe.Pointer, buflen C.int) {
	id := uintptr(userdata)
	fn := getCallback(id).(AudioStreamDataCompleteCallbackFunc)
	fn(buf, int(buflen))
	unregisterCallback(id)
}

// SetGetCallback sets a callback that fires when data is requested from the stream.
func (s *AudioStream) SetGetCallback(callback AudioStreamCallbackFunc) error {
	if callback == nil {
		if !C.SDL_SetAudioStreamGetCallback(s.c, nil, nil) {
			return getError()
		}
		return nil
	}
	id := registerCallback(callback)
	if !C.SDL_SetAudioStreamGetCallback(s.c, C._get_audio_stream_callback_trampoline(), ptrFromID(id)) {
		unregisterCallback(id)
		return getError()
	}
	return nil
}

// SetPutCallback sets a callback that fires when data is added to the stream.
func (s *AudioStream) SetPutCallback(callback AudioStreamCallbackFunc) error {
	if callback == nil {
		if !C.SDL_SetAudioStreamPutCallback(s.c, nil, nil) {
			return getError()
		}
		return nil
	}
	id := registerCallback(callback)
	if !C.SDL_SetAudioStreamPutCallback(s.c, C._get_audio_stream_callback_trampoline(), ptrFromID(id)) {
		unregisterCallback(id)
		return getError()
	}
	return nil
}

// SetAudioPostmixCallback sets a callback that fires after audio mixing on a device.
func SetAudioPostmixCallback(devid AudioDeviceID, callback AudioPostmixCallbackFunc) error {
	if callback == nil {
		if !C.SDL_SetAudioPostmixCallback(C.SDL_AudioDeviceID(devid), nil, nil) {
			return getError()
		}
		return nil
	}
	id := registerCallback(callback)
	if !C.SDL_SetAudioPostmixCallback(C.SDL_AudioDeviceID(devid), C.SDL_AudioPostmixCallback(C.cgoAudioPostmixCallbackTrampoline), ptrFromID(id)) {
		unregisterCallback(id)
		return getError()
	}
	return nil
}

// OpenAudioDeviceStream opens an audio device and creates a stream with a callback.
func OpenAudioDeviceStream(devid AudioDeviceID, spec *AudioSpec, callback AudioStreamCallbackFunc) (*AudioStream, error) {
	var cs *C.SDL_AudioSpec
	if spec != nil {
		cs = spec.cptr()
	}
	var ccb C.SDL_AudioStreamCallback
	var cud unsafe.Pointer
	if callback != nil {
		id := registerCallback(callback)
		ccb = C._get_audio_stream_callback_trampoline()
		cud = ptrFromID(id)
	}
	as := C.SDL_OpenAudioDeviceStream(C.SDL_AudioDeviceID(devid), cs, ccb, cud)
	if as == nil {
		return nil, getError()
	}
	return &AudioStream{c: as}, nil
}

// PutDataNoCopy queues audio data without copying. The data must remain valid until the callback fires.
func (s *AudioStream) PutDataNoCopy(buf unsafe.Pointer, length int, callback AudioStreamDataCompleteCallbackFunc) error {
	var ccb C.SDL_AudioStreamDataCompleteCallback
	var cud unsafe.Pointer
	if callback != nil {
		id := registerCallback(callback)
		ccb = C.SDL_AudioStreamDataCompleteCallback(C.cgoAudioStreamDataCompleteTrampoline)
		cud = ptrFromID(id)
	}
	if !C.SDL_PutAudioStreamDataNoCopy(s.c, buf, C.int(length), ccb, cud) {
		return getError()
	}
	return nil
}
