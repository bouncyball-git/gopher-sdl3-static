package sdl

/*
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// PixelFormat represents an SDL pixel format.
type PixelFormat uint32

const (
	PIXELFORMAT_UNKNOWN     PixelFormat = C.SDL_PIXELFORMAT_UNKNOWN
	PIXELFORMAT_INDEX1LSB   PixelFormat = C.SDL_PIXELFORMAT_INDEX1LSB
	PIXELFORMAT_INDEX1MSB   PixelFormat = C.SDL_PIXELFORMAT_INDEX1MSB
	PIXELFORMAT_INDEX4LSB   PixelFormat = C.SDL_PIXELFORMAT_INDEX4LSB
	PIXELFORMAT_INDEX4MSB   PixelFormat = C.SDL_PIXELFORMAT_INDEX4MSB
	PIXELFORMAT_INDEX8      PixelFormat = C.SDL_PIXELFORMAT_INDEX8
	PIXELFORMAT_RGB332      PixelFormat = C.SDL_PIXELFORMAT_RGB332
	PIXELFORMAT_XRGB4444   PixelFormat = C.SDL_PIXELFORMAT_XRGB4444
	PIXELFORMAT_XBGR4444   PixelFormat = C.SDL_PIXELFORMAT_XBGR4444
	PIXELFORMAT_XRGB1555   PixelFormat = C.SDL_PIXELFORMAT_XRGB1555
	PIXELFORMAT_XBGR1555   PixelFormat = C.SDL_PIXELFORMAT_XBGR1555
	PIXELFORMAT_ARGB4444   PixelFormat = C.SDL_PIXELFORMAT_ARGB4444
	PIXELFORMAT_RGBA4444   PixelFormat = C.SDL_PIXELFORMAT_RGBA4444
	PIXELFORMAT_ABGR4444   PixelFormat = C.SDL_PIXELFORMAT_ABGR4444
	PIXELFORMAT_BGRA4444   PixelFormat = C.SDL_PIXELFORMAT_BGRA4444
	PIXELFORMAT_ARGB1555   PixelFormat = C.SDL_PIXELFORMAT_ARGB1555
	PIXELFORMAT_RGBA5551   PixelFormat = C.SDL_PIXELFORMAT_RGBA5551
	PIXELFORMAT_ABGR1555   PixelFormat = C.SDL_PIXELFORMAT_ABGR1555
	PIXELFORMAT_BGRA5551   PixelFormat = C.SDL_PIXELFORMAT_BGRA5551
	PIXELFORMAT_RGB565      PixelFormat = C.SDL_PIXELFORMAT_RGB565
	PIXELFORMAT_BGR565      PixelFormat = C.SDL_PIXELFORMAT_BGR565
	PIXELFORMAT_RGB24       PixelFormat = C.SDL_PIXELFORMAT_RGB24
	PIXELFORMAT_BGR24       PixelFormat = C.SDL_PIXELFORMAT_BGR24
	PIXELFORMAT_XRGB8888   PixelFormat = C.SDL_PIXELFORMAT_XRGB8888
	PIXELFORMAT_RGBX8888   PixelFormat = C.SDL_PIXELFORMAT_RGBX8888
	PIXELFORMAT_XBGR8888   PixelFormat = C.SDL_PIXELFORMAT_XBGR8888
	PIXELFORMAT_BGRX8888   PixelFormat = C.SDL_PIXELFORMAT_BGRX8888
	PIXELFORMAT_ARGB8888   PixelFormat = C.SDL_PIXELFORMAT_ARGB8888
	PIXELFORMAT_RGBA8888   PixelFormat = C.SDL_PIXELFORMAT_RGBA8888
	PIXELFORMAT_ABGR8888   PixelFormat = C.SDL_PIXELFORMAT_ABGR8888
	PIXELFORMAT_BGRA8888   PixelFormat = C.SDL_PIXELFORMAT_BGRA8888
	PIXELFORMAT_XRGB2101010 PixelFormat = C.SDL_PIXELFORMAT_XRGB2101010
	PIXELFORMAT_XBGR2101010 PixelFormat = C.SDL_PIXELFORMAT_XBGR2101010
	PIXELFORMAT_ARGB2101010 PixelFormat = C.SDL_PIXELFORMAT_ARGB2101010
	PIXELFORMAT_ABGR2101010 PixelFormat = C.SDL_PIXELFORMAT_ABGR2101010
	PIXELFORMAT_RGB48       PixelFormat = C.SDL_PIXELFORMAT_RGB48
	PIXELFORMAT_BGR48       PixelFormat = C.SDL_PIXELFORMAT_BGR48
	PIXELFORMAT_RGBA64      PixelFormat = C.SDL_PIXELFORMAT_RGBA64
	PIXELFORMAT_ARGB64      PixelFormat = C.SDL_PIXELFORMAT_ARGB64
	PIXELFORMAT_BGRA64      PixelFormat = C.SDL_PIXELFORMAT_BGRA64
	PIXELFORMAT_ABGR64      PixelFormat = C.SDL_PIXELFORMAT_ABGR64
	PIXELFORMAT_RGB48_FLOAT PixelFormat = C.SDL_PIXELFORMAT_RGB48_FLOAT
	PIXELFORMAT_BGR48_FLOAT PixelFormat = C.SDL_PIXELFORMAT_BGR48_FLOAT
	PIXELFORMAT_RGBA64_FLOAT PixelFormat = C.SDL_PIXELFORMAT_RGBA64_FLOAT
	PIXELFORMAT_ARGB64_FLOAT PixelFormat = C.SDL_PIXELFORMAT_ARGB64_FLOAT
	PIXELFORMAT_BGRA64_FLOAT PixelFormat = C.SDL_PIXELFORMAT_BGRA64_FLOAT
	PIXELFORMAT_ABGR64_FLOAT PixelFormat = C.SDL_PIXELFORMAT_ABGR64_FLOAT
	PIXELFORMAT_RGB96_FLOAT  PixelFormat = C.SDL_PIXELFORMAT_RGB96_FLOAT
	PIXELFORMAT_BGR96_FLOAT  PixelFormat = C.SDL_PIXELFORMAT_BGR96_FLOAT
	PIXELFORMAT_RGBA128_FLOAT PixelFormat = C.SDL_PIXELFORMAT_RGBA128_FLOAT
	PIXELFORMAT_ARGB128_FLOAT PixelFormat = C.SDL_PIXELFORMAT_ARGB128_FLOAT
	PIXELFORMAT_BGRA128_FLOAT PixelFormat = C.SDL_PIXELFORMAT_BGRA128_FLOAT
	PIXELFORMAT_ABGR128_FLOAT PixelFormat = C.SDL_PIXELFORMAT_ABGR128_FLOAT

	// Indexed formats (2-bit)
	PIXELFORMAT_INDEX2LSB PixelFormat = C.SDL_PIXELFORMAT_INDEX2LSB
	PIXELFORMAT_INDEX2MSB PixelFormat = C.SDL_PIXELFORMAT_INDEX2MSB

	// YUV formats
	PIXELFORMAT_YV12 PixelFormat = C.SDL_PIXELFORMAT_YV12
	PIXELFORMAT_IYUV PixelFormat = C.SDL_PIXELFORMAT_IYUV
	PIXELFORMAT_YUY2 PixelFormat = C.SDL_PIXELFORMAT_YUY2
	PIXELFORMAT_UYVY PixelFormat = C.SDL_PIXELFORMAT_UYVY
	PIXELFORMAT_YVYU PixelFormat = C.SDL_PIXELFORMAT_YVYU
	PIXELFORMAT_NV12 PixelFormat = C.SDL_PIXELFORMAT_NV12
	PIXELFORMAT_NV21 PixelFormat = C.SDL_PIXELFORMAT_NV21
	PIXELFORMAT_P010 PixelFormat = C.SDL_PIXELFORMAT_P010

	// Compressed/special formats
	PIXELFORMAT_EXTERNAL_OES PixelFormat = C.SDL_PIXELFORMAT_EXTERNAL_OES
	PIXELFORMAT_MJPG         PixelFormat = C.SDL_PIXELFORMAT_MJPG

	// Endian-dependent 32-bit aliases
	PIXELFORMAT_RGBA32 PixelFormat = C.SDL_PIXELFORMAT_RGBA32
	PIXELFORMAT_ARGB32 PixelFormat = C.SDL_PIXELFORMAT_ARGB32
	PIXELFORMAT_BGRA32 PixelFormat = C.SDL_PIXELFORMAT_BGRA32
	PIXELFORMAT_ABGR32 PixelFormat = C.SDL_PIXELFORMAT_ABGR32
	PIXELFORMAT_RGBX32 PixelFormat = C.SDL_PIXELFORMAT_RGBX32
	PIXELFORMAT_XRGB32 PixelFormat = C.SDL_PIXELFORMAT_XRGB32
	PIXELFORMAT_BGRX32 PixelFormat = C.SDL_PIXELFORMAT_BGRX32
	PIXELFORMAT_XBGR32 PixelFormat = C.SDL_PIXELFORMAT_XBGR32
)

// Color represents an RGBA color with 8-bit components.
type Color struct {
	R, G, B, A uint8
}

// FColor represents an RGBA color with floating-point components.
type FColor struct {
	R, G, B, A float32
}

// Colorspace represents an SDL color space.
type Colorspace uint32

const (
	ALPHA_OPAQUE            = 255
	ALPHA_OPAQUE_FLOAT      = 1.0
	ALPHA_TRANSPARENT       = 0
	ALPHA_TRANSPARENT_FLOAT = 0.0
)

// PixelType represents the type of pixel data.
type PixelType int

const (
	PIXELTYPE_UNKNOWN  PixelType = C.SDL_PIXELTYPE_UNKNOWN
	PIXELTYPE_INDEX1   PixelType = C.SDL_PIXELTYPE_INDEX1
	PIXELTYPE_INDEX4   PixelType = C.SDL_PIXELTYPE_INDEX4
	PIXELTYPE_INDEX8   PixelType = C.SDL_PIXELTYPE_INDEX8
	PIXELTYPE_PACKED8  PixelType = C.SDL_PIXELTYPE_PACKED8
	PIXELTYPE_PACKED16 PixelType = C.SDL_PIXELTYPE_PACKED16
	PIXELTYPE_PACKED32 PixelType = C.SDL_PIXELTYPE_PACKED32
	PIXELTYPE_ARRAYU8  PixelType = C.SDL_PIXELTYPE_ARRAYU8
	PIXELTYPE_ARRAYU16 PixelType = C.SDL_PIXELTYPE_ARRAYU16
	PIXELTYPE_ARRAYU32 PixelType = C.SDL_PIXELTYPE_ARRAYU32
	PIXELTYPE_ARRAYF16 PixelType = C.SDL_PIXELTYPE_ARRAYF16
	PIXELTYPE_ARRAYF32 PixelType = C.SDL_PIXELTYPE_ARRAYF32
	PIXELTYPE_INDEX2   PixelType = C.SDL_PIXELTYPE_INDEX2
)

// BitmapOrder represents bitmap pixel order.
type BitmapOrder int

const (
	BITMAPORDER_NONE BitmapOrder = C.SDL_BITMAPORDER_NONE
	BITMAPORDER_4321 BitmapOrder = C.SDL_BITMAPORDER_4321
	BITMAPORDER_1234 BitmapOrder = C.SDL_BITMAPORDER_1234
)

// PackedOrder represents packed pixel order.
type PackedOrder int

const (
	PACKEDORDER_NONE PackedOrder = C.SDL_PACKEDORDER_NONE
	PACKEDORDER_XRGB PackedOrder = C.SDL_PACKEDORDER_XRGB
	PACKEDORDER_RGBX PackedOrder = C.SDL_PACKEDORDER_RGBX
	PACKEDORDER_ARGB PackedOrder = C.SDL_PACKEDORDER_ARGB
	PACKEDORDER_RGBA PackedOrder = C.SDL_PACKEDORDER_RGBA
	PACKEDORDER_XBGR PackedOrder = C.SDL_PACKEDORDER_XBGR
	PACKEDORDER_BGRX PackedOrder = C.SDL_PACKEDORDER_BGRX
	PACKEDORDER_ABGR PackedOrder = C.SDL_PACKEDORDER_ABGR
	PACKEDORDER_BGRA PackedOrder = C.SDL_PACKEDORDER_BGRA
)

// ArrayOrder represents array pixel order.
type ArrayOrder int

const (
	ARRAYORDER_NONE ArrayOrder = C.SDL_ARRAYORDER_NONE
	ARRAYORDER_RGB  ArrayOrder = C.SDL_ARRAYORDER_RGB
	ARRAYORDER_RGBA ArrayOrder = C.SDL_ARRAYORDER_RGBA
	ARRAYORDER_ARGB ArrayOrder = C.SDL_ARRAYORDER_ARGB
	ARRAYORDER_BGR  ArrayOrder = C.SDL_ARRAYORDER_BGR
	ARRAYORDER_BGRA ArrayOrder = C.SDL_ARRAYORDER_BGRA
	ARRAYORDER_ABGR ArrayOrder = C.SDL_ARRAYORDER_ABGR
)

// PackedLayout represents packed pixel layout.
type PackedLayout int

const (
	PACKEDLAYOUT_NONE      PackedLayout = C.SDL_PACKEDLAYOUT_NONE
	PACKEDLAYOUT_332       PackedLayout = C.SDL_PACKEDLAYOUT_332
	PACKEDLAYOUT_4444      PackedLayout = C.SDL_PACKEDLAYOUT_4444
	PACKEDLAYOUT_1555      PackedLayout = C.SDL_PACKEDLAYOUT_1555
	PACKEDLAYOUT_5551      PackedLayout = C.SDL_PACKEDLAYOUT_5551
	PACKEDLAYOUT_565       PackedLayout = C.SDL_PACKEDLAYOUT_565
	PACKEDLAYOUT_8888      PackedLayout = C.SDL_PACKEDLAYOUT_8888
	PACKEDLAYOUT_2101010   PackedLayout = C.SDL_PACKEDLAYOUT_2101010
	PACKEDLAYOUT_1010102   PackedLayout = C.SDL_PACKEDLAYOUT_1010102
)

// Palette represents an SDL color palette.
type Palette struct {
	c *C.SDL_Palette
}

// ColorType represents a colorspace color type.
type ColorType uint32

const (
	COLOR_TYPE_UNKNOWN ColorType = C.SDL_COLOR_TYPE_UNKNOWN
	COLOR_TYPE_RGB     ColorType = C.SDL_COLOR_TYPE_RGB
	COLOR_TYPE_YCBCR   ColorType = C.SDL_COLOR_TYPE_YCBCR
)

// ColorRange represents a colorspace color range.
type ColorRange uint32

const (
	COLOR_RANGE_UNKNOWN ColorRange = C.SDL_COLOR_RANGE_UNKNOWN
	COLOR_RANGE_LIMITED ColorRange = C.SDL_COLOR_RANGE_LIMITED
	COLOR_RANGE_FULL    ColorRange = C.SDL_COLOR_RANGE_FULL
)

// ColorPrimaries represents colorspace color primaries.
type ColorPrimaries uint32

const (
	COLOR_PRIMARIES_UNKNOWN      ColorPrimaries = C.SDL_COLOR_PRIMARIES_UNKNOWN
	COLOR_PRIMARIES_BT709        ColorPrimaries = C.SDL_COLOR_PRIMARIES_BT709
	COLOR_PRIMARIES_UNSPECIFIED  ColorPrimaries = C.SDL_COLOR_PRIMARIES_UNSPECIFIED
	COLOR_PRIMARIES_BT470M       ColorPrimaries = C.SDL_COLOR_PRIMARIES_BT470M
	COLOR_PRIMARIES_BT470BG      ColorPrimaries = C.SDL_COLOR_PRIMARIES_BT470BG
	COLOR_PRIMARIES_BT601        ColorPrimaries = C.SDL_COLOR_PRIMARIES_BT601
	COLOR_PRIMARIES_SMPTE240     ColorPrimaries = C.SDL_COLOR_PRIMARIES_SMPTE240
	COLOR_PRIMARIES_GENERIC_FILM ColorPrimaries = C.SDL_COLOR_PRIMARIES_GENERIC_FILM
	COLOR_PRIMARIES_BT2020       ColorPrimaries = C.SDL_COLOR_PRIMARIES_BT2020
	COLOR_PRIMARIES_XYZ          ColorPrimaries = C.SDL_COLOR_PRIMARIES_XYZ
	COLOR_PRIMARIES_SMPTE431     ColorPrimaries = C.SDL_COLOR_PRIMARIES_SMPTE431
	COLOR_PRIMARIES_SMPTE432     ColorPrimaries = C.SDL_COLOR_PRIMARIES_SMPTE432
	COLOR_PRIMARIES_EBU3213      ColorPrimaries = C.SDL_COLOR_PRIMARIES_EBU3213
	COLOR_PRIMARIES_CUSTOM       ColorPrimaries = C.SDL_COLOR_PRIMARIES_CUSTOM
)

// TransferCharacteristics represents colorspace transfer characteristics.
type TransferCharacteristics uint32

const (
	TRANSFER_CHARACTERISTICS_UNKNOWN      TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_UNKNOWN
	TRANSFER_CHARACTERISTICS_BT709        TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_BT709
	TRANSFER_CHARACTERISTICS_UNSPECIFIED  TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_UNSPECIFIED
	TRANSFER_CHARACTERISTICS_GAMMA22      TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_GAMMA22
	TRANSFER_CHARACTERISTICS_GAMMA28      TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_GAMMA28
	TRANSFER_CHARACTERISTICS_BT601        TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_BT601
	TRANSFER_CHARACTERISTICS_SMPTE240     TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_SMPTE240
	TRANSFER_CHARACTERISTICS_LINEAR       TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_LINEAR
	TRANSFER_CHARACTERISTICS_LOG100       TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_LOG100
	TRANSFER_CHARACTERISTICS_LOG100_SQRT10 TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_LOG100_SQRT10
	TRANSFER_CHARACTERISTICS_IEC61966     TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_IEC61966
	TRANSFER_CHARACTERISTICS_BT1361       TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_BT1361
	TRANSFER_CHARACTERISTICS_SRGB         TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_SRGB
	TRANSFER_CHARACTERISTICS_BT2020_10BIT TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_BT2020_10BIT
	TRANSFER_CHARACTERISTICS_BT2020_12BIT TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_BT2020_12BIT
	TRANSFER_CHARACTERISTICS_PQ           TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_PQ
	TRANSFER_CHARACTERISTICS_SMPTE428     TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_SMPTE428
	TRANSFER_CHARACTERISTICS_HLG          TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_HLG
	TRANSFER_CHARACTERISTICS_CUSTOM       TransferCharacteristics = C.SDL_TRANSFER_CHARACTERISTICS_CUSTOM
)

// MatrixCoefficients represents colorspace matrix coefficients.
type MatrixCoefficients uint32

const (
	MATRIX_COEFFICIENTS_IDENTITY           MatrixCoefficients = C.SDL_MATRIX_COEFFICIENTS_IDENTITY
	MATRIX_COEFFICIENTS_BT709              MatrixCoefficients = C.SDL_MATRIX_COEFFICIENTS_BT709
	MATRIX_COEFFICIENTS_UNSPECIFIED        MatrixCoefficients = C.SDL_MATRIX_COEFFICIENTS_UNSPECIFIED
	MATRIX_COEFFICIENTS_FCC                MatrixCoefficients = C.SDL_MATRIX_COEFFICIENTS_FCC
	MATRIX_COEFFICIENTS_BT470BG            MatrixCoefficients = C.SDL_MATRIX_COEFFICIENTS_BT470BG
	MATRIX_COEFFICIENTS_BT601              MatrixCoefficients = C.SDL_MATRIX_COEFFICIENTS_BT601
	MATRIX_COEFFICIENTS_SMPTE240           MatrixCoefficients = C.SDL_MATRIX_COEFFICIENTS_SMPTE240
	MATRIX_COEFFICIENTS_YCGCO              MatrixCoefficients = C.SDL_MATRIX_COEFFICIENTS_YCGCO
	MATRIX_COEFFICIENTS_BT2020_NCL         MatrixCoefficients = C.SDL_MATRIX_COEFFICIENTS_BT2020_NCL
	MATRIX_COEFFICIENTS_BT2020_CL          MatrixCoefficients = C.SDL_MATRIX_COEFFICIENTS_BT2020_CL
	MATRIX_COEFFICIENTS_SMPTE2085          MatrixCoefficients = C.SDL_MATRIX_COEFFICIENTS_SMPTE2085
	MATRIX_COEFFICIENTS_CHROMA_DERIVED_NCL MatrixCoefficients = C.SDL_MATRIX_COEFFICIENTS_CHROMA_DERIVED_NCL
	MATRIX_COEFFICIENTS_CHROMA_DERIVED_CL  MatrixCoefficients = C.SDL_MATRIX_COEFFICIENTS_CHROMA_DERIVED_CL
	MATRIX_COEFFICIENTS_ICTCP              MatrixCoefficients = C.SDL_MATRIX_COEFFICIENTS_ICTCP
	MATRIX_COEFFICIENTS_CUSTOM             MatrixCoefficients = C.SDL_MATRIX_COEFFICIENTS_CUSTOM
)

// ChromaLocation represents a colorspace chroma sample location.
type ChromaLocation uint32

const (
	CHROMA_LOCATION_NONE    ChromaLocation = C.SDL_CHROMA_LOCATION_NONE
	CHROMA_LOCATION_LEFT    ChromaLocation = C.SDL_CHROMA_LOCATION_LEFT
	CHROMA_LOCATION_CENTER  ChromaLocation = C.SDL_CHROMA_LOCATION_CENTER
	CHROMA_LOCATION_TOPLEFT ChromaLocation = C.SDL_CHROMA_LOCATION_TOPLEFT
)

const (
	COLORSPACE_UNKNOWN        Colorspace = C.SDL_COLORSPACE_UNKNOWN
	COLORSPACE_SRGB           Colorspace = C.SDL_COLORSPACE_SRGB
	COLORSPACE_SRGB_LINEAR    Colorspace = C.SDL_COLORSPACE_SRGB_LINEAR
	COLORSPACE_HDR10          Colorspace = C.SDL_COLORSPACE_HDR10
	COLORSPACE_JPEG           Colorspace = C.SDL_COLORSPACE_JPEG
	COLORSPACE_BT601_LIMITED  Colorspace = C.SDL_COLORSPACE_BT601_LIMITED
	COLORSPACE_BT601_FULL     Colorspace = C.SDL_COLORSPACE_BT601_FULL
	COLORSPACE_BT709_LIMITED  Colorspace = C.SDL_COLORSPACE_BT709_LIMITED
	COLORSPACE_BT709_FULL     Colorspace = C.SDL_COLORSPACE_BT709_FULL
	COLORSPACE_BT2020_LIMITED Colorspace = C.SDL_COLORSPACE_BT2020_LIMITED
	COLORSPACE_BT2020_FULL    Colorspace = C.SDL_COLORSPACE_BT2020_FULL
	COLORSPACE_RGB_DEFAULT    Colorspace = C.SDL_COLORSPACE_RGB_DEFAULT
	COLORSPACE_YUV_DEFAULT    Colorspace = C.SDL_COLORSPACE_YUV_DEFAULT
)

// PixelFormatDetails contains details about the format of a pixel.
type PixelFormatDetails struct {
	c *C.SDL_PixelFormatDetails
}

// Format returns the pixel format.
func (d *PixelFormatDetails) Format() PixelFormat {
	return PixelFormat(d.c.format)
}

// BitsPerPixel returns the number of significant bits in a pixel value.
func (d *PixelFormatDetails) BitsPerPixel() uint8 {
	return uint8(d.c.bits_per_pixel)
}

// BytesPerPixel returns the number of bytes required to hold a pixel value.
func (d *PixelFormatDetails) BytesPerPixel() uint8 {
	return uint8(d.c.bytes_per_pixel)
}

// GetPixelFormatName returns a human-readable name for a pixel format.
func GetPixelFormatName(format PixelFormat) string {
	return C.GoString(C.SDL_GetPixelFormatName(C.SDL_PixelFormat(format)))
}

// GetPixelFormatDetails returns a PixelFormatDetails structure for the given pixel format.
func GetPixelFormatDetails(format PixelFormat) (*PixelFormatDetails, error) {
	c := C.SDL_GetPixelFormatDetails(C.SDL_PixelFormat(format))
	if c == nil {
		return nil, getError()
	}
	return &PixelFormatDetails{c: c}, nil
}

// GetMasksForPixelFormat converts a pixel format to a bpp value and RGBA masks.
func GetMasksForPixelFormat(format PixelFormat) (bpp int, rmask, gmask, bmask, amask uint32, ok bool) {
	var cbpp C.int
	var cr, cg, cb, ca C.Uint32
	ok = bool(C.SDL_GetMasksForPixelFormat(C.SDL_PixelFormat(format), &cbpp, &cr, &cg, &cb, &ca))
	return int(cbpp), uint32(cr), uint32(cg), uint32(cb), uint32(ca), ok
}

// GetPixelFormatForMasks converts a bpp value and RGBA masks to a pixel format.
func GetPixelFormatForMasks(bpp int, rmask, gmask, bmask, amask uint32) PixelFormat {
	return PixelFormat(C.SDL_GetPixelFormatForMasks(C.int(bpp), C.Uint32(rmask), C.Uint32(gmask), C.Uint32(bmask), C.Uint32(amask)))
}

// CreatePalette creates a palette structure with the specified number of color entries.
func CreatePalette(ncolors int) (*Palette, error) {
	c := C.SDL_CreatePalette(C.int(ncolors))
	if c == nil {
		return nil, getError()
	}
	return &Palette{c: c}, nil
}

// SetPaletteColors sets a range of colors in a palette.
func SetPaletteColors(palette *Palette, colors []Color, firstcolor int) error {
	if !C.SDL_SetPaletteColors(palette.c, (*C.SDL_Color)(unsafe.Pointer(&colors[0])), C.int(firstcolor), C.int(len(colors))) {
		return getError()
	}
	return nil
}

// DestroyPalette frees a palette created with CreatePalette.
func DestroyPalette(palette *Palette) {
	if palette != nil && palette.c != nil {
		C.SDL_DestroyPalette(palette.c)
		palette.c = nil
	}
}

// MapRGB maps an RGB triple to an opaque pixel value for the given pixel format.
func MapRGB(format *PixelFormatDetails, palette *Palette, r, g, b uint8) uint32 {
	var cp *C.SDL_Palette
	if palette != nil {
		cp = palette.c
	}
	return uint32(C.SDL_MapRGB(format.c, cp, C.Uint8(r), C.Uint8(g), C.Uint8(b)))
}

// MapRGBA maps an RGBA quadruple to a pixel value for the given pixel format.
func MapRGBA(format *PixelFormatDetails, palette *Palette, r, g, b, a uint8) uint32 {
	var cp *C.SDL_Palette
	if palette != nil {
		cp = palette.c
	}
	return uint32(C.SDL_MapRGBA(format.c, cp, C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)))
}

// GetRGB gets RGB values from a pixel in the specified format.
func GetRGB(pixel uint32, format *PixelFormatDetails, palette *Palette) (r, g, b uint8) {
	var cp *C.SDL_Palette
	if palette != nil {
		cp = palette.c
	}
	var cr, cg, cb C.Uint8
	C.SDL_GetRGB(C.Uint32(pixel), format.c, cp, &cr, &cg, &cb)
	return uint8(cr), uint8(cg), uint8(cb)
}

// GetRGBA gets RGBA values from a pixel in the specified format.
func GetRGBA(pixel uint32, format *PixelFormatDetails, palette *Palette) (r, g, b, a uint8) {
	var cp *C.SDL_Palette
	if palette != nil {
		cp = palette.c
	}
	var cr, cg, cb, ca C.Uint8
	C.SDL_GetRGBA(C.Uint32(pixel), format.c, cp, &cr, &cg, &cb, &ca)
	return uint8(cr), uint8(cg), uint8(cb), uint8(ca)
}
