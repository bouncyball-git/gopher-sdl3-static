package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// Surface represents an SDL surface.
type Surface struct {
	c *C.SDL_Surface
}

// SurfaceFlags are flags for a surface.
type SurfaceFlags uint32

// Surface flag constants.
const (
	SURFACE_PREALLOCATED SurfaceFlags = C.SDL_SURFACE_PREALLOCATED
	SURFACE_LOCK_NEEDED  SurfaceFlags = C.SDL_SURFACE_LOCK_NEEDED
	SURFACE_LOCKED       SurfaceFlags = C.SDL_SURFACE_LOCKED
	SURFACE_SIMD_ALIGNED SurfaceFlags = C.SDL_SURFACE_SIMD_ALIGNED
)

// ScaleMode specifies the scaling mode for surface operations.
type ScaleMode int

const (
	SCALEMODE_INVALID ScaleMode = C.SDL_SCALEMODE_INVALID
	SCALEMODE_NEAREST ScaleMode = C.SDL_SCALEMODE_NEAREST
	SCALEMODE_LINEAR  ScaleMode = C.SDL_SCALEMODE_LINEAR
)

// FlipMode specifies the flip mode for surface operations.
type FlipMode int

const (
	FLIP_NONE                   FlipMode = C.SDL_FLIP_NONE
	FLIP_HORIZONTAL             FlipMode = C.SDL_FLIP_HORIZONTAL
	FLIP_VERTICAL               FlipMode = C.SDL_FLIP_VERTICAL
	FLIP_HORIZONTAL_AND_VERTICAL FlipMode = C.SDL_FLIP_HORIZONTAL_AND_VERTICAL
)

// CreateSurface creates a new surface with the specified dimensions and format.
func CreateSurface(width, height int, format PixelFormat) (*Surface, error) {
	cs := C.SDL_CreateSurface(C.int(width), C.int(height), C.SDL_PixelFormat(format))
	if cs == nil {
		return nil, getError()
	}
	return &Surface{c: cs}, nil
}

// CreateSurfaceFrom creates a surface from existing pixel data.
func CreateSurfaceFrom(width, height int, format PixelFormat, pixels unsafe.Pointer, pitch int) (*Surface, error) {
	cs := C.SDL_CreateSurfaceFrom(C.int(width), C.int(height), C.SDL_PixelFormat(format), pixels, C.int(pitch))
	if cs == nil {
		return nil, getError()
	}
	return &Surface{c: cs}, nil
}

// Destroy frees the surface.
func (s *Surface) Destroy() {
	if s.c != nil {
		C.SDL_DestroySurface(s.c)
		s.c = nil
	}
}

// Width returns the width of the surface.
func (s *Surface) Width() int {
	return int(s.c.w)
}

// Height returns the height of the surface.
func (s *Surface) Height() int {
	return int(s.c.h)
}

// Pitch returns the pitch (bytes per row) of the surface.
func (s *Surface) Pitch() int {
	return int(s.c.pitch)
}

// Format returns the pixel format of the surface.
func (s *Surface) Format() PixelFormat {
	return PixelFormat(s.c.format)
}

// Pixels returns a pointer to the pixel data of the surface.
func (s *Surface) Pixels() unsafe.Pointer {
	return unsafe.Pointer(s.c.pixels)
}

// Lock locks the surface for direct pixel access.
func (s *Surface) Lock() error {
	if !C.SDL_LockSurface(s.c) {
		return getError()
	}
	return nil
}

// Unlock unlocks the surface.
func (s *Surface) Unlock() {
	C.SDL_UnlockSurface(s.c)
}

// SetColorKey sets the color key (transparent pixel) for the surface.
func (s *Surface) SetColorKey(enabled bool, key uint32) error {
	if !C.SDL_SetSurfaceColorKey(s.c, C.bool(enabled), C.Uint32(key)) {
		return getError()
	}
	return nil
}

// FillRect fills a rectangle on the surface with a color.
func (s *Surface) FillRect(rect *Rect, color uint32) error {
	var cr *C.SDL_Rect
	if rect != nil {
		cr = rect.cptr()
	}
	if !C.SDL_FillSurfaceRect(s.c, cr, C.Uint32(color)) {
		return getError()
	}
	return nil
}

// Blit performs a blit from the source surface to this surface.
func (s *Surface) Blit(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	var sr, dr *C.SDL_Rect
	if srcRect != nil {
		sr = srcRect.cptr()
	}
	if dstRect != nil {
		dr = dstRect.cptr()
	}
	if !C.SDL_BlitSurface(s.c, sr, dst.c, dr) {
		return getError()
	}
	return nil
}

// SetBlendMode sets the blend mode for the surface.
func (s *Surface) SetBlendMode(blendMode BlendMode) error {
	if !C.SDL_SetSurfaceBlendMode(s.c, C.SDL_BlendMode(blendMode)) {
		return getError()
	}
	return nil
}

// LoadBMP loads a BMP image from a file.
func LoadBMP(file string) (*Surface, error) {
	cf := C.CString(file)
	defer C.free(unsafe.Pointer(cf))
	cs := C.SDL_LoadBMP(cf)
	if cs == nil {
		return nil, getError()
	}
	return &Surface{c: cs}, nil
}

// Properties returns the properties associated with a surface.
func (s *Surface) Properties() PropertiesID {
	return PropertiesID(C.SDL_GetSurfaceProperties(s.c))
}

// SetColorspace sets the colorspace used by a surface.
func (s *Surface) SetColorspace(colorspace Colorspace) error {
	if !C.SDL_SetSurfaceColorspace(s.c, C.SDL_Colorspace(colorspace)) {
		return getError()
	}
	return nil
}

// GetColorspace returns the colorspace used by a surface.
func (s *Surface) GetColorspace() Colorspace {
	return Colorspace(C.SDL_GetSurfaceColorspace(s.c))
}

// CreatePalette creates a palette and associates it with the surface.
func (s *Surface) CreatePalette() (*Palette, error) {
	cp := C.SDL_CreateSurfacePalette(s.c)
	if cp == nil {
		return nil, getError()
	}
	return &Palette{c: cp}, nil
}

// SetPalette sets the palette used by the surface.
func (s *Surface) SetPalette(palette *Palette) error {
	if !C.SDL_SetSurfacePalette(s.c, palette.c) {
		return getError()
	}
	return nil
}

// GetPalette returns the palette used by the surface.
func (s *Surface) GetPalette() *Palette {
	cp := C.SDL_GetSurfacePalette(s.c)
	if cp == nil {
		return nil
	}
	return &Palette{c: cp}
}

// AddAlternateImage adds an alternate version of a surface.
func (s *Surface) AddAlternateImage(image *Surface) error {
	if !C.SDL_AddSurfaceAlternateImage(s.c, image.c) {
		return getError()
	}
	return nil
}

// HasAlternateImages returns whether a surface has alternate versions available.
func (s *Surface) HasAlternateImages() bool {
	return bool(C.SDL_SurfaceHasAlternateImages(s.c))
}

// GetImages returns an array including all versions of a surface.
func (s *Surface) GetImages() []*Surface {
	var count C.int
	cimages := C.SDL_GetSurfaceImages(s.c, &count)
	if cimages == nil {
		return nil
	}
	defer C.SDL_free(unsafe.Pointer(cimages))
	images := make([]*Surface, int(count))
	cslice := unsafe.Slice(cimages, int(count))
	for i := 0; i < int(count); i++ {
		images[i] = &Surface{c: cslice[i]}
	}
	return images
}

// RemoveAlternateImages removes all alternate versions of a surface.
func (s *Surface) RemoveAlternateImages() {
	C.SDL_RemoveSurfaceAlternateImages(s.c)
}

// SetRLE sets the RLE acceleration hint for a surface.
func (s *Surface) SetRLE(enabled bool) error {
	if !C.SDL_SetSurfaceRLE(s.c, C.bool(enabled)) {
		return getError()
	}
	return nil
}

// HasRLE returns whether the surface is RLE enabled.
func (s *Surface) HasRLE() bool {
	return bool(C.SDL_SurfaceHasRLE(s.c))
}

// HasColorKey returns whether the surface has a color key.
func (s *Surface) HasColorKey() bool {
	return bool(C.SDL_SurfaceHasColorKey(s.c))
}

// GetColorKey returns the color key (transparent pixel) for the surface.
func (s *Surface) GetColorKey() (uint32, error) {
	var key C.Uint32
	if !C.SDL_GetSurfaceColorKey(s.c, &key) {
		return 0, getError()
	}
	return uint32(key), nil
}

// SetColorMod sets an additional color value multiplied into blit operations.
func (s *Surface) SetColorMod(r, g, b uint8) error {
	if !C.SDL_SetSurfaceColorMod(s.c, C.Uint8(r), C.Uint8(g), C.Uint8(b)) {
		return getError()
	}
	return nil
}

// GetColorMod returns the additional color value multiplied into blit operations.
func (s *Surface) GetColorMod() (r, g, b uint8, err error) {
	var cr, cg, cb C.Uint8
	if !C.SDL_GetSurfaceColorMod(s.c, &cr, &cg, &cb) {
		return 0, 0, 0, getError()
	}
	return uint8(cr), uint8(cg), uint8(cb), nil
}

// SetAlphaMod sets an additional alpha value used in blit operations.
func (s *Surface) SetAlphaMod(alpha uint8) error {
	if !C.SDL_SetSurfaceAlphaMod(s.c, C.Uint8(alpha)) {
		return getError()
	}
	return nil
}

// GetAlphaMod returns the additional alpha value used in blit operations.
func (s *Surface) GetAlphaMod() (uint8, error) {
	var alpha C.Uint8
	if !C.SDL_GetSurfaceAlphaMod(s.c, &alpha) {
		return 0, getError()
	}
	return uint8(alpha), nil
}

// GetBlendMode returns the blend mode used for blit operations.
func (s *Surface) GetBlendMode() (BlendMode, error) {
	var blendMode C.SDL_BlendMode
	if !C.SDL_GetSurfaceBlendMode(s.c, &blendMode) {
		return 0, getError()
	}
	return BlendMode(blendMode), nil
}

// SetClipRect sets the clipping rectangle for a surface.
func (s *Surface) SetClipRect(rect *Rect) bool {
	var cr *C.SDL_Rect
	if rect != nil {
		cr = rect.cptr()
	}
	return bool(C.SDL_SetSurfaceClipRect(s.c, cr))
}

// GetClipRect returns the clipping rectangle for a surface.
func (s *Surface) GetClipRect() (Rect, error) {
	var rect Rect
	if !C.SDL_GetSurfaceClipRect(s.c, rect.cptr()) {
		return rect, getError()
	}
	return rect, nil
}

// Flip flips a surface vertically or horizontally.
func (s *Surface) Flip(flip FlipMode) error {
	if !C.SDL_FlipSurface(s.c, C.SDL_FlipMode(flip)) {
		return getError()
	}
	return nil
}

// Duplicate creates a new surface identical to the existing surface.
func (s *Surface) Duplicate() (*Surface, error) {
	cs := C.SDL_DuplicateSurface(s.c)
	if cs == nil {
		return nil, getError()
	}
	return &Surface{c: cs}, nil
}

// Scale creates a new surface identical to the existing surface, scaled to the desired size.
func (s *Surface) Scale(width, height int, scaleMode ScaleMode) (*Surface, error) {
	cs := C.SDL_ScaleSurface(s.c, C.int(width), C.int(height), C.SDL_ScaleMode(scaleMode))
	if cs == nil {
		return nil, getError()
	}
	return &Surface{c: cs}, nil
}

// Convert copies the surface to a new surface of the specified format.
func (s *Surface) Convert(format PixelFormat) (*Surface, error) {
	cs := C.SDL_ConvertSurface(s.c, C.SDL_PixelFormat(format))
	if cs == nil {
		return nil, getError()
	}
	return &Surface{c: cs}, nil
}

// ConvertAndColorspace copies the surface to a new surface of the specified format and colorspace.
func (s *Surface) ConvertAndColorspace(format PixelFormat, palette *Palette, colorspace Colorspace, props PropertiesID) (*Surface, error) {
	var cp *C.SDL_Palette
	if palette != nil {
		cp = palette.c
	}
	cs := C.SDL_ConvertSurfaceAndColorspace(s.c, C.SDL_PixelFormat(format), cp, C.SDL_Colorspace(colorspace), C.SDL_PropertiesID(props))
	if cs == nil {
		return nil, getError()
	}
	return &Surface{c: cs}, nil
}

// ConvertPixels copies a block of pixels of one format to another format.
func ConvertPixels(width, height int, srcFormat PixelFormat, src unsafe.Pointer, srcPitch int, dstFormat PixelFormat, dst unsafe.Pointer, dstPitch int) error {
	if !C.SDL_ConvertPixels(C.int(width), C.int(height), C.SDL_PixelFormat(srcFormat), src, C.int(srcPitch), C.SDL_PixelFormat(dstFormat), dst, C.int(dstPitch)) {
		return getError()
	}
	return nil
}

// ConvertPixelsAndColorspace copies a block of pixels of one format and colorspace to another.
func ConvertPixelsAndColorspace(width, height int, srcFormat PixelFormat, srcColorspace Colorspace, srcProps PropertiesID, src unsafe.Pointer, srcPitch int, dstFormat PixelFormat, dstColorspace Colorspace, dstProps PropertiesID, dst unsafe.Pointer, dstPitch int) error {
	if !C.SDL_ConvertPixelsAndColorspace(C.int(width), C.int(height), C.SDL_PixelFormat(srcFormat), C.SDL_Colorspace(srcColorspace), C.SDL_PropertiesID(srcProps), src, C.int(srcPitch), C.SDL_PixelFormat(dstFormat), C.SDL_Colorspace(dstColorspace), C.SDL_PropertiesID(dstProps), dst, C.int(dstPitch)) {
		return getError()
	}
	return nil
}

// PremultiplyAlpha premultiplies the alpha on a block of pixels.
func PremultiplyAlpha(width, height int, srcFormat PixelFormat, src unsafe.Pointer, srcPitch int, dstFormat PixelFormat, dst unsafe.Pointer, dstPitch int, linear bool) error {
	if !C.SDL_PremultiplyAlpha(C.int(width), C.int(height), C.SDL_PixelFormat(srcFormat), src, C.int(srcPitch), C.SDL_PixelFormat(dstFormat), dst, C.int(dstPitch), C.bool(linear)) {
		return getError()
	}
	return nil
}

// PremultiplyAlpha premultiplies the alpha in a surface.
func (s *Surface) PremultiplyAlpha(linear bool) error {
	if !C.SDL_PremultiplySurfaceAlpha(s.c, C.bool(linear)) {
		return getError()
	}
	return nil
}

// Clear clears a surface with a specific color, with floating point precision.
func (s *Surface) Clear(r, g, b, a float32) error {
	if !C.SDL_ClearSurface(s.c, C.float(r), C.float(g), C.float(b), C.float(a)) {
		return getError()
	}
	return nil
}

// FillRects fills a set of rectangles on the surface with a color.
func (s *Surface) FillRects(rects []Rect, color uint32) error {
	if len(rects) == 0 {
		return nil
	}
	if !C.SDL_FillSurfaceRects(s.c, rects[0].cptr(), C.int(len(rects)), C.Uint32(color)) {
		return getError()
	}
	return nil
}

// BlitUnchecked performs low-level surface blitting only (no clipping).
func (s *Surface) BlitUnchecked(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	var sr, dr *C.SDL_Rect
	if srcRect != nil {
		sr = srcRect.cptr()
	}
	if dstRect != nil {
		dr = dstRect.cptr()
	}
	if !C.SDL_BlitSurfaceUnchecked(s.c, sr, dst.c, dr) {
		return getError()
	}
	return nil
}

// BlitScaled performs a scaled blit from the source surface to the destination surface.
func (s *Surface) BlitScaled(srcRect *Rect, dst *Surface, dstRect *Rect, scaleMode ScaleMode) error {
	var sr, dr *C.SDL_Rect
	if srcRect != nil {
		sr = srcRect.cptr()
	}
	if dstRect != nil {
		dr = dstRect.cptr()
	}
	if !C.SDL_BlitSurfaceScaled(s.c, sr, dst.c, dr, C.SDL_ScaleMode(scaleMode)) {
		return getError()
	}
	return nil
}

// BlitUncheckedScaled performs low-level scaled surface blitting only (no clipping).
func (s *Surface) BlitUncheckedScaled(srcRect *Rect, dst *Surface, dstRect *Rect, scaleMode ScaleMode) error {
	var sr, dr *C.SDL_Rect
	if srcRect != nil {
		sr = srcRect.cptr()
	}
	if dstRect != nil {
		dr = dstRect.cptr()
	}
	if !C.SDL_BlitSurfaceUncheckedScaled(s.c, sr, dst.c, dr, C.SDL_ScaleMode(scaleMode)) {
		return getError()
	}
	return nil
}

// BlitTiled performs a tiled blit from the source surface to the destination surface.
func (s *Surface) BlitTiled(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	var sr, dr *C.SDL_Rect
	if srcRect != nil {
		sr = srcRect.cptr()
	}
	if dstRect != nil {
		dr = dstRect.cptr()
	}
	if !C.SDL_BlitSurfaceTiled(s.c, sr, dst.c, dr) {
		return getError()
	}
	return nil
}

// BlitTiledWithScale performs a scaled and tiled blit from the source surface to the destination surface.
func (s *Surface) BlitTiledWithScale(srcRect *Rect, scale float32, scaleMode ScaleMode, dst *Surface, dstRect *Rect) error {
	var sr, dr *C.SDL_Rect
	if srcRect != nil {
		sr = srcRect.cptr()
	}
	if dstRect != nil {
		dr = dstRect.cptr()
	}
	if !C.SDL_BlitSurfaceTiledWithScale(s.c, sr, C.float(scale), C.SDL_ScaleMode(scaleMode), dst.c, dr) {
		return getError()
	}
	return nil
}

// Blit9Grid performs a scaled blit using the 9-grid algorithm.
func (s *Surface) Blit9Grid(srcRect *Rect, leftWidth, rightWidth, topHeight, bottomHeight int, scale float32, scaleMode ScaleMode, dst *Surface, dstRect *Rect) error {
	var sr, dr *C.SDL_Rect
	if srcRect != nil {
		sr = srcRect.cptr()
	}
	if dstRect != nil {
		dr = dstRect.cptr()
	}
	if !C.SDL_BlitSurface9Grid(s.c, sr, C.int(leftWidth), C.int(rightWidth), C.int(topHeight), C.int(bottomHeight), C.float(scale), C.SDL_ScaleMode(scaleMode), dst.c, dr) {
		return getError()
	}
	return nil
}

// MapRGB maps an RGB triple to an opaque pixel value for a surface.
func (s *Surface) MapRGB(r, g, b uint8) uint32 {
	return uint32(C.SDL_MapSurfaceRGB(s.c, C.Uint8(r), C.Uint8(g), C.Uint8(b)))
}

// MapRGBA maps an RGBA quadruple to a pixel value for a surface.
func (s *Surface) MapRGBA(r, g, b, a uint8) uint32 {
	return uint32(C.SDL_MapSurfaceRGBA(s.c, C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)))
}

// ReadPixel retrieves a single pixel from a surface as 8-bit RGBA components.
func (s *Surface) ReadPixel(x, y int) (r, g, b, a uint8, err error) {
	var cr, cg, cb, ca C.Uint8
	if !C.SDL_ReadSurfacePixel(s.c, C.int(x), C.int(y), &cr, &cg, &cb, &ca) {
		return 0, 0, 0, 0, getError()
	}
	return uint8(cr), uint8(cg), uint8(cb), uint8(ca), nil
}

// ReadPixelFloat retrieves a single pixel from a surface as floating-point RGBA components.
func (s *Surface) ReadPixelFloat(x, y int) (r, g, b, a float32, err error) {
	var cr, cg, cb, ca C.float
	if !C.SDL_ReadSurfacePixelFloat(s.c, C.int(x), C.int(y), &cr, &cg, &cb, &ca) {
		return 0, 0, 0, 0, getError()
	}
	return float32(cr), float32(cg), float32(cb), float32(ca), nil
}

// WritePixel writes a single pixel to a surface with 8-bit RGBA components.
func (s *Surface) WritePixel(x, y int, r, g, b, a uint8) error {
	if !C.SDL_WriteSurfacePixel(s.c, C.int(x), C.int(y), C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)) {
		return getError()
	}
	return nil
}

// WritePixelFloat writes a single pixel to a surface with floating-point RGBA components.
func (s *Surface) WritePixelFloat(x, y int, r, g, b, a float32) error {
	if !C.SDL_WriteSurfacePixelFloat(s.c, C.int(x), C.int(y), C.float(r), C.float(g), C.float(b), C.float(a)) {
		return getError()
	}
	return nil
}

// SaveBMP saves the surface to a file in BMP format.
func (s *Surface) SaveBMP(file string) error {
	cf := C.CString(file)
	defer C.free(unsafe.Pointer(cf))
	if !C.SDL_SaveBMP(s.c, cf) {
		return getError()
	}
	return nil
}

// SaveBMP_IO saves the surface to an SDL I/O stream in BMP format.
func (s *Surface) SaveBMP_IO(dst *IOStream, closeio bool) error {
	if !C.SDL_SaveBMP_IO(s.c, dst.c, C.bool(closeio)) {
		return getError()
	}
	return nil
}

// LoadBMP_IO loads a BMP image from an SDL I/O stream.
func LoadBMP_IO(src *IOStream, closeio bool) (*Surface, error) {
	cs := C.SDL_LoadBMP_IO(src.c, C.bool(closeio))
	if cs == nil {
		return nil, getError()
	}
	return &Surface{c: cs}, nil
}

// RotateSurface returns a copy of a surface rotated clockwise by the given angle in degrees.
func (s *Surface) Rotate(angle float32) (*Surface, error) {
	cs := C.SDL_RotateSurface(s.c, C.float(angle))
	if cs == nil {
		return nil, getError()
	}
	return &Surface{c: cs}, nil
}

// LoadSurface loads a surface from a file (format auto-detected).
func LoadSurface(file string) (*Surface, error) {
	cf := C.CString(file)
	defer C.free(unsafe.Pointer(cf))
	cs := C.SDL_LoadSurface(cf)
	if cs == nil {
		return nil, getError()
	}
	return &Surface{c: cs}, nil
}

// LoadSurface_IO loads a surface from an I/O stream.
func LoadSurface_IO(src *IOStream, closeio bool) (*Surface, error) {
	cs := C.SDL_LoadSurface_IO(src.c, C.bool(closeio))
	if cs == nil {
		return nil, getError()
	}
	return &Surface{c: cs}, nil
}

// LoadPNG loads a PNG image from a file.
func LoadPNG(file string) (*Surface, error) {
	cf := C.CString(file)
	defer C.free(unsafe.Pointer(cf))
	cs := C.SDL_LoadPNG(cf)
	if cs == nil {
		return nil, getError()
	}
	return &Surface{c: cs}, nil
}

// LoadPNG_IO loads a PNG image from an I/O stream.
func LoadPNG_IO(src *IOStream, closeio bool) (*Surface, error) {
	cs := C.SDL_LoadPNG_IO(src.c, C.bool(closeio))
	if cs == nil {
		return nil, getError()
	}
	return &Surface{c: cs}, nil
}

// SavePNG saves the surface as a PNG file.
func (s *Surface) SavePNG(file string) error {
	cf := C.CString(file)
	defer C.free(unsafe.Pointer(cf))
	if !C.SDL_SavePNG(s.c, cf) {
		return getError()
	}
	return nil
}

// SavePNG_IO saves the surface as PNG to an I/O stream.
func (s *Surface) SavePNG_IO(dst *IOStream, closeio bool) error {
	if !C.SDL_SavePNG_IO(s.c, dst.c, C.bool(closeio)) {
		return getError()
	}
	return nil
}

// StretchSurface stretches a source surface rectangle to a destination surface rectangle.
func StretchSurface(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect, scaleMode ScaleMode) error {
	var csr, cdr *C.SDL_Rect
	if srcrect != nil {
		csr = srcrect.cptr()
	}
	if dstrect != nil {
		cdr = dstrect.cptr()
	}
	if !C.SDL_StretchSurface(src.c, csr, dst.c, cdr, C.SDL_ScaleMode(scaleMode)) {
		return getError()
	}
	return nil
}
