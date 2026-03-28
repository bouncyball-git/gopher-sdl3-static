package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>

static bool _sdl_render_debug_text_format(SDL_Renderer *renderer, float x, float y, const char *str) {
	return SDL_RenderDebugTextFormat(renderer, x, y, "%s", str);
}
*/
import "C"

import "unsafe"

// Renderer represents an SDL rendering context.
type Renderer struct {
	c *C.SDL_Renderer
}

// Texture represents an SDL texture.
type Texture struct {
	c *C.SDL_Texture
}

// TextureAccess defines how a texture is accessed.
type TextureAccess int

const (
	TEXTUREACCESS_STATIC    TextureAccess = C.SDL_TEXTUREACCESS_STATIC
	TEXTUREACCESS_STREAMING TextureAccess = C.SDL_TEXTUREACCESS_STREAMING
	TEXTUREACCESS_TARGET    TextureAccess = C.SDL_TEXTUREACCESS_TARGET
)

// RendererLogicalPresentation specifies the logical presentation mode.
type RendererLogicalPresentation int

const (
	LOGICAL_PRESENTATION_DISABLED      RendererLogicalPresentation = C.SDL_LOGICAL_PRESENTATION_DISABLED
	LOGICAL_PRESENTATION_STRETCH       RendererLogicalPresentation = C.SDL_LOGICAL_PRESENTATION_STRETCH
	LOGICAL_PRESENTATION_LETTERBOX     RendererLogicalPresentation = C.SDL_LOGICAL_PRESENTATION_LETTERBOX
	LOGICAL_PRESENTATION_OVERSCAN      RendererLogicalPresentation = C.SDL_LOGICAL_PRESENTATION_OVERSCAN
	LOGICAL_PRESENTATION_INTEGER_SCALE RendererLogicalPresentation = C.SDL_LOGICAL_PRESENTATION_INTEGER_SCALE
)

// Renderer name constants.
const (
	SOFTWARE_RENDERER = "software"
	GPU_RENDERER      = "gpu"
)

// Renderer VSync constants.
const (
	RENDERER_VSYNC_DISABLED = 0
	RENDERER_VSYNC_ADAPTIVE = -1
)

// Debug text font character size in pixels.
const DEBUG_TEXT_FONT_CHARACTER_SIZE = C.SDL_DEBUG_TEXT_FONT_CHARACTER_SIZE

// Vertex contains vertex information for rendering geometry.
type Vertex struct {
	Position FPoint
	Color    FColor
	TexCoord FPoint
}

// CreateWindowAndRenderer creates a window and a renderer.
func CreateWindowAndRenderer(title string, width, height int, windowFlags WindowFlags) (*Window, *Renderer, error) {
	ct := C.CString(title)
	defer C.free(unsafe.Pointer(ct))
	var cw *C.SDL_Window
	var cr *C.SDL_Renderer
	if !C.SDL_CreateWindowAndRenderer(ct, C.int(width), C.int(height), C.SDL_WindowFlags(windowFlags), &cw, &cr) {
		return nil, nil, getError()
	}
	return &Window{c: cw}, &Renderer{c: cr}, nil
}

// CreateRenderer creates a renderer for the given window.
func CreateRenderer(window *Window, name string) (*Renderer, error) {
	var cn *C.char
	if name != "" {
		cn = C.CString(name)
		defer C.free(unsafe.Pointer(cn))
	}
	cr := C.SDL_CreateRenderer(window.c, cn)
	if cr == nil {
		return nil, getError()
	}
	return &Renderer{c: cr}, nil
}

// Destroy destroys the renderer.
func (r *Renderer) Destroy() {
	if r.c != nil {
		C.SDL_DestroyRenderer(r.c)
		r.c = nil
	}
}

// Name returns the name of the renderer.
func (r *Renderer) Name() string {
	return C.GoString(C.SDL_GetRendererName(r.c))
}

// Window returns the window associated with the renderer via properties.
func (r *Renderer) Window() *Window {
	props := C.SDL_GetRendererProperties(r.c)
	if props == 0 {
		return nil
	}
	cname := C.CString("SDL.renderer.window")
	defer C.free(unsafe.Pointer(cname))
	p := C.SDL_GetPointerProperty(props, cname, nil)
	if p == nil {
		return nil
	}
	return &Window{c: (*C.SDL_Window)(p)}
}

// OutputSize returns the output size of the renderer.
func (r *Renderer) OutputSize() (int, int, error) {
	var w, h C.int
	if !C.SDL_GetRenderOutputSize(r.c, &w, &h) {
		return 0, 0, getError()
	}
	return int(w), int(h), nil
}

// CurrentOutputSize returns the current output size of the renderer.
func (r *Renderer) CurrentOutputSize() (int, int, error) {
	var w, h C.int
	if !C.SDL_GetCurrentRenderOutputSize(r.c, &w, &h) {
		return 0, 0, getError()
	}
	return int(w), int(h), nil
}

// SetDrawColor sets the drawing color for the renderer.
func (r *Renderer) SetDrawColor(red, green, blue, alpha uint8) error {
	if !C.SDL_SetRenderDrawColor(r.c, C.Uint8(red), C.Uint8(green), C.Uint8(blue), C.Uint8(alpha)) {
		return getError()
	}
	return nil
}

// SetDrawColorFloat sets the drawing color for the renderer using float values.
func (r *Renderer) SetDrawColorFloat(red, green, blue, alpha float32) error {
	if !C.SDL_SetRenderDrawColorFloat(r.c, C.float(red), C.float(green), C.float(blue), C.float(alpha)) {
		return getError()
	}
	return nil
}

// GetDrawColor returns the drawing color of the renderer.
func (r *Renderer) GetDrawColor() (uint8, uint8, uint8, uint8, error) {
	var red, green, blue, alpha C.Uint8
	if !C.SDL_GetRenderDrawColor(r.c, &red, &green, &blue, &alpha) {
		return 0, 0, 0, 0, getError()
	}
	return uint8(red), uint8(green), uint8(blue), uint8(alpha), nil
}

// SetDrawBlendMode sets the blend mode for drawing operations.
func (r *Renderer) SetDrawBlendMode(blendMode BlendMode) error {
	if !C.SDL_SetRenderDrawBlendMode(r.c, C.SDL_BlendMode(blendMode)) {
		return getError()
	}
	return nil
}

// Clear clears the current rendering target.
func (r *Renderer) Clear() error {
	if !C.SDL_RenderClear(r.c) {
		return getError()
	}
	return nil
}

// Present presents the rendered content to the screen.
func (r *Renderer) Present() error {
	if !C.SDL_RenderPresent(r.c) {
		return getError()
	}
	return nil
}

// DrawPoint draws a point on the renderer.
func (r *Renderer) DrawPoint(x, y float32) error {
	if !C.SDL_RenderPoint(r.c, C.float(x), C.float(y)) {
		return getError()
	}
	return nil
}

// DrawLine draws a line on the renderer.
func (r *Renderer) DrawLine(x1, y1, x2, y2 float32) error {
	if !C.SDL_RenderLine(r.c, C.float(x1), C.float(y1), C.float(x2), C.float(y2)) {
		return getError()
	}
	return nil
}

// DrawLines draws a series of connected lines on the renderer.
func (r *Renderer) DrawLines(points []FPoint) error {
	if len(points) == 0 {
		return nil
	}
	if !C.SDL_RenderLines(r.c, (*C.SDL_FPoint)(unsafe.Pointer(&points[0])), C.int(len(points))) {
		return getError()
	}
	return nil
}

// DrawRect draws a rectangle outline on the renderer.
func (r *Renderer) DrawRect(rect *FRect) error {
	var cr *C.SDL_FRect
	if rect != nil {
		cr = rect.cptr()
	}
	if !C.SDL_RenderRect(r.c, cr) {
		return getError()
	}
	return nil
}

// DrawRects draws multiple rectangle outlines on the renderer.
func (r *Renderer) DrawRects(rects []FRect) error {
	if len(rects) == 0 {
		return nil
	}
	if !C.SDL_RenderRects(r.c, (*C.SDL_FRect)(unsafe.Pointer(&rects[0])), C.int(len(rects))) {
		return getError()
	}
	return nil
}

// FillRect fills a rectangle on the renderer.
func (r *Renderer) FillRect(rect *FRect) error {
	var cr *C.SDL_FRect
	if rect != nil {
		cr = rect.cptr()
	}
	if !C.SDL_RenderFillRect(r.c, cr) {
		return getError()
	}
	return nil
}

// FillRects fills multiple rectangles on the renderer.
func (r *Renderer) FillRects(rects []FRect) error {
	if len(rects) == 0 {
		return nil
	}
	if !C.SDL_RenderFillRects(r.c, (*C.SDL_FRect)(unsafe.Pointer(&rects[0])), C.int(len(rects))) {
		return getError()
	}
	return nil
}

// SetTarget sets the rendering target to a texture.
func (r *Renderer) SetTarget(texture *Texture) error {
	var ct *C.SDL_Texture
	if texture != nil {
		ct = texture.c
	}
	if !C.SDL_SetRenderTarget(r.c, ct) {
		return getError()
	}
	return nil
}

// GetTarget returns the current rendering target.
func (r *Renderer) GetTarget() *Texture {
	ct := C.SDL_GetRenderTarget(r.c)
	if ct == nil {
		return nil
	}
	return &Texture{c: ct}
}

// SetLogicalPresentation sets the logical presentation mode.
func (r *Renderer) SetLogicalPresentation(w, h int, mode RendererLogicalPresentation) error {
	if !C.SDL_SetRenderLogicalPresentation(r.c, C.int(w), C.int(h), C.SDL_RendererLogicalPresentation(mode)) {
		return getError()
	}
	return nil
}

// SetViewport sets the drawing area for rendering.
func (r *Renderer) SetViewport(rect *Rect) error {
	var cr *C.SDL_Rect
	if rect != nil {
		cr = rect.cptr()
	}
	if !C.SDL_SetRenderViewport(r.c, cr) {
		return getError()
	}
	return nil
}

// GetViewport returns the drawing area for the current target.
func (r *Renderer) GetViewport() (Rect, error) {
	var rect Rect
	if !C.SDL_GetRenderViewport(r.c, rect.cptr()) {
		return rect, getError()
	}
	return rect, nil
}

// SetClipRect sets the clip rectangle for rendering.
func (r *Renderer) SetClipRect(rect *Rect) error {
	var cr *C.SDL_Rect
	if rect != nil {
		cr = rect.cptr()
	}
	if !C.SDL_SetRenderClipRect(r.c, cr) {
		return getError()
	}
	return nil
}

// SetScale sets the drawing scale for the renderer.
func (r *Renderer) SetScale(scaleX, scaleY float32) error {
	if !C.SDL_SetRenderScale(r.c, C.float(scaleX), C.float(scaleY)) {
		return getError()
	}
	return nil
}

// GetScale returns the drawing scale for the renderer.
func (r *Renderer) GetScale() (float32, float32, error) {
	var scaleX, scaleY C.float
	if !C.SDL_GetRenderScale(r.c, &scaleX, &scaleY) {
		return 0, 0, getError()
	}
	return float32(scaleX), float32(scaleY), nil
}

// RenderTexture copies a texture to the rendering target.
func (r *Renderer) RenderTexture(texture *Texture, srcRect, dstRect *FRect) error {
	var sr, dr *C.SDL_FRect
	if srcRect != nil {
		sr = srcRect.cptr()
	}
	if dstRect != nil {
		dr = dstRect.cptr()
	}
	if !C.SDL_RenderTexture(r.c, texture.c, sr, dr) {
		return getError()
	}
	return nil
}

// RenderTextureRotated copies a texture to the rendering target with rotation and flipping.
func (r *Renderer) RenderTextureRotated(texture *Texture, srcRect, dstRect *FRect, angle float64, center *FPoint, flip FlipMode) error {
	var sr, dr *C.SDL_FRect
	if srcRect != nil {
		sr = srcRect.cptr()
	}
	if dstRect != nil {
		dr = dstRect.cptr()
	}
	var cp *C.SDL_FPoint
	if center != nil {
		cp = center.cptr()
	}
	if !C.SDL_RenderTextureRotated(r.c, texture.c, sr, dr, C.double(angle), cp, C.SDL_FlipMode(flip)) {
		return getError()
	}
	return nil
}

// RenderGeometry renders a list of triangles.
func (r *Renderer) RenderGeometry(texture *Texture, vertices []Vertex, indices []int32) error {
	var ct *C.SDL_Texture
	if texture != nil {
		ct = texture.c
	}
	var cv *C.SDL_Vertex
	if len(vertices) > 0 {
		cv = (*C.SDL_Vertex)(unsafe.Pointer(&vertices[0]))
	}
	var ci *C.int
	if len(indices) > 0 {
		ci = (*C.int)(unsafe.Pointer(&indices[0]))
	}
	if !C.SDL_RenderGeometry(r.c, ct, cv, C.int(len(vertices)), ci, C.int(len(indices))) {
		return getError()
	}
	return nil
}

// SetVSync sets VSync for the renderer.
func (r *Renderer) SetVSync(vsync int) error {
	if !C.SDL_SetRenderVSync(r.c, C.int(vsync)) {
		return getError()
	}
	return nil
}

// --- Texture functions ---

// CreateTexture creates a texture for the renderer.
func CreateTexture(renderer *Renderer, format PixelFormat, access TextureAccess, w, h int) (*Texture, error) {
	ct := C.SDL_CreateTexture(renderer.c, C.SDL_PixelFormat(format), C.SDL_TextureAccess(access), C.int(w), C.int(h))
	if ct == nil {
		return nil, getError()
	}
	return &Texture{c: ct}, nil
}

// CreateTextureFromSurface creates a texture from a surface.
func CreateTextureFromSurface(renderer *Renderer, surface *Surface) (*Texture, error) {
	ct := C.SDL_CreateTextureFromSurface(renderer.c, surface.c)
	if ct == nil {
		return nil, getError()
	}
	return &Texture{c: ct}, nil
}

// Destroy destroys the texture.
func (t *Texture) Destroy() {
	if t.c != nil {
		C.SDL_DestroyTexture(t.c)
		t.c = nil
	}
}

// Size returns the size of the texture.
func (t *Texture) Size() (float32, float32, error) {
	var w, h C.float
	if !C.SDL_GetTextureSize(t.c, &w, &h) {
		return 0, 0, getError()
	}
	return float32(w), float32(h), nil
}

// SetColorMod sets the color modulation for the texture.
func (t *Texture) SetColorMod(r, g, b uint8) error {
	if !C.SDL_SetTextureColorMod(t.c, C.Uint8(r), C.Uint8(g), C.Uint8(b)) {
		return getError()
	}
	return nil
}

// SetAlphaMod sets the alpha modulation for the texture.
func (t *Texture) SetAlphaMod(alpha uint8) error {
	if !C.SDL_SetTextureAlphaMod(t.c, C.Uint8(alpha)) {
		return getError()
	}
	return nil
}

// SetBlendMode sets the blend mode for the texture.
func (t *Texture) SetBlendMode(blendMode BlendMode) error {
	if !C.SDL_SetTextureBlendMode(t.c, C.SDL_BlendMode(blendMode)) {
		return getError()
	}
	return nil
}

// SetScaleMode sets the scale mode for the texture.
func (t *Texture) SetScaleMode(scaleMode ScaleMode) error {
	if !C.SDL_SetTextureScaleMode(t.c, C.SDL_ScaleMode(scaleMode)) {
		return getError()
	}
	return nil
}

// Update updates the texture with new pixel data.
func (t *Texture) Update(rect *Rect, pixels unsafe.Pointer, pitch int) error {
	var cr *C.SDL_Rect
	if rect != nil {
		cr = rect.cptr()
	}
	if !C.SDL_UpdateTexture(t.c, cr, pixels, C.int(pitch)) {
		return getError()
	}
	return nil
}

// Lock locks the texture for write-only pixel access.
func (t *Texture) Lock(rect *Rect) (unsafe.Pointer, int, error) {
	var cr *C.SDL_Rect
	if rect != nil {
		cr = rect.cptr()
	}
	var pixels unsafe.Pointer
	var pitch C.int
	if !C.SDL_LockTexture(t.c, cr, &pixels, &pitch) {
		return nil, 0, getError()
	}
	return pixels, int(pitch), nil
}

// Unlock unlocks a previously locked texture.
func (t *Texture) Unlock() {
	C.SDL_UnlockTexture(t.c)
}

// --- Additional renderer creation functions ---

// CreateRendererWithProperties creates a renderer with the specified properties.
func CreateRendererWithProperties(props PropertiesID) (*Renderer, error) {
	cr := C.SDL_CreateRendererWithProperties(C.SDL_PropertiesID(props))
	if cr == nil {
		return nil, getError()
	}
	return &Renderer{c: cr}, nil
}

// CreateSoftwareRenderer creates a 2D software rendering context for a surface.
func CreateSoftwareRenderer(surface *Surface) (*Renderer, error) {
	cr := C.SDL_CreateSoftwareRenderer(surface.c)
	if cr == nil {
		return nil, getError()
	}
	return &Renderer{c: cr}, nil
}

// GetNumRenderDrivers returns the number of 2D rendering drivers available.
func GetNumRenderDrivers() int {
	return int(C.SDL_GetNumRenderDrivers())
}

// GetRenderDriver returns the name of a built-in 2D rendering driver.
func GetRenderDriver(index int) string {
	return C.GoString(C.SDL_GetRenderDriver(C.int(index)))
}

// GetRendererFromTexture returns the renderer that created a texture.
func GetRendererFromTexture(texture *Texture) *Renderer {
	cr := C.SDL_GetRendererFromTexture(texture.c)
	if cr == nil {
		return nil
	}
	return &Renderer{c: cr}
}

// --- Renderer properties ---

// GetProperties returns the properties associated with the renderer.
func (r *Renderer) GetProperties() PropertiesID {
	return PropertiesID(C.SDL_GetRendererProperties(r.c))
}

// --- Renderer logical presentation ---

// GetLogicalPresentation returns the logical presentation settings.
func (r *Renderer) GetLogicalPresentation() (w, h int, mode RendererLogicalPresentation, err error) {
	var cw, ch C.int
	var cm C.SDL_RendererLogicalPresentation
	if !C.SDL_GetRenderLogicalPresentation(r.c, &cw, &ch, &cm) {
		return 0, 0, 0, getError()
	}
	return int(cw), int(ch), RendererLogicalPresentation(cm), nil
}

// GetLogicalPresentationRect returns the final presentation rectangle for rendering.
func (r *Renderer) GetLogicalPresentationRect() (FRect, error) {
	var rect FRect
	if !C.SDL_GetRenderLogicalPresentationRect(r.c, rect.cptr()) {
		return rect, getError()
	}
	return rect, nil
}

// --- Coordinate conversion ---

// RenderCoordinatesFromWindow converts window coordinates to render coordinates.
func (r *Renderer) RenderCoordinatesFromWindow(windowX, windowY float32) (float32, float32, error) {
	var x, y C.float
	if !C.SDL_RenderCoordinatesFromWindow(r.c, C.float(windowX), C.float(windowY), &x, &y) {
		return 0, 0, getError()
	}
	return float32(x), float32(y), nil
}

// RenderCoordinatesToWindow converts render coordinates to window coordinates.
func (r *Renderer) RenderCoordinatesToWindow(x, y float32) (float32, float32, error) {
	var wx, wy C.float
	if !C.SDL_RenderCoordinatesToWindow(r.c, C.float(x), C.float(y), &wx, &wy) {
		return 0, 0, getError()
	}
	return float32(wx), float32(wy), nil
}

// ConvertEventToRenderCoordinates converts the coordinates in a raw SDL_Event
// to render coordinates. Pass a pointer to a C.SDL_Event as unsafe.Pointer.
func (r *Renderer) ConvertEventToRenderCoordinates(event unsafe.Pointer) error {
	if !C.SDL_ConvertEventToRenderCoordinates(r.c, (*C.SDL_Event)(event)) {
		return getError()
	}
	return nil
}

// --- Safe area ---

// GetSafeArea returns the safe area for rendering within the current viewport.
func (r *Renderer) GetSafeArea() (Rect, error) {
	var rect Rect
	if !C.SDL_GetRenderSafeArea(r.c, rect.cptr()) {
		return rect, getError()
	}
	return rect, nil
}

// --- Clip rect ---

// GetClipRect returns the clip rectangle for the current target.
func (r *Renderer) GetClipRect() (Rect, error) {
	var rect Rect
	if !C.SDL_GetRenderClipRect(r.c, rect.cptr()) {
		return rect, getError()
	}
	return rect, nil
}

// ClipEnabled returns whether clipping is enabled on the renderer.
func (r *Renderer) ClipEnabled() bool {
	return bool(C.SDL_RenderClipEnabled(r.c))
}

// --- Color scale ---

// SetColorScale sets the color scale used for render operations.
func (r *Renderer) SetColorScale(scale float32) error {
	if !C.SDL_SetRenderColorScale(r.c, C.float(scale)) {
		return getError()
	}
	return nil
}

// GetColorScale returns the color scale used for render operations.
func (r *Renderer) GetColorScale() (float32, error) {
	var scale C.float
	if !C.SDL_GetRenderColorScale(r.c, &scale) {
		return 0, getError()
	}
	return float32(scale), nil
}

// --- Draw color float getter ---

// GetDrawColorFloat returns the drawing color of the renderer as float values.
func (r *Renderer) GetDrawColorFloat() (float32, float32, float32, float32, error) {
	var red, green, blue, alpha C.float
	if !C.SDL_GetRenderDrawColorFloat(r.c, &red, &green, &blue, &alpha) {
		return 0, 0, 0, 0, getError()
	}
	return float32(red), float32(green), float32(blue), float32(alpha), nil
}

// --- Additional render functions ---

// RenderPoints draws multiple points on the renderer.
func (r *Renderer) RenderPoints(points []FPoint) error {
	if len(points) == 0 {
		return nil
	}
	if !C.SDL_RenderPoints(r.c, (*C.SDL_FPoint)(unsafe.Pointer(&points[0])), C.int(len(points))) {
		return getError()
	}
	return nil
}

// RenderTextureAffine copies a portion of the texture with an affine transform.
func (r *Renderer) RenderTextureAffine(texture *Texture, srcrect *FRect, origin, right, down *FPoint) error {
	var sr *C.SDL_FRect
	if srcrect != nil {
		sr = srcrect.cptr()
	}
	var op, rp, dp *C.SDL_FPoint
	if origin != nil {
		op = origin.cptr()
	}
	if right != nil {
		rp = right.cptr()
	}
	if down != nil {
		dp = down.cptr()
	}
	if !C.SDL_RenderTextureAffine(r.c, texture.c, sr, op, rp, dp) {
		return getError()
	}
	return nil
}

// RenderTextureTiled tiles a portion of the texture to the rendering target.
func (r *Renderer) RenderTextureTiled(texture *Texture, srcrect *FRect, scale float32, dstrect *FRect) error {
	var sr, dr *C.SDL_FRect
	if srcrect != nil {
		sr = srcrect.cptr()
	}
	if dstrect != nil {
		dr = dstrect.cptr()
	}
	if !C.SDL_RenderTextureTiled(r.c, texture.c, sr, C.float(scale), dr) {
		return getError()
	}
	return nil
}

// RenderTexture9Grid performs a scaled copy using the 9-grid algorithm.
func (r *Renderer) RenderTexture9Grid(texture *Texture, srcrect *FRect, leftWidth, rightWidth, topHeight, bottomHeight, scale float32, dstrect *FRect) error {
	var sr, dr *C.SDL_FRect
	if srcrect != nil {
		sr = srcrect.cptr()
	}
	if dstrect != nil {
		dr = dstrect.cptr()
	}
	if !C.SDL_RenderTexture9Grid(r.c, texture.c, sr, C.float(leftWidth), C.float(rightWidth), C.float(topHeight), C.float(bottomHeight), C.float(scale), dr) {
		return getError()
	}
	return nil
}

// RenderReadPixels reads pixels from the current rendering target.
func (r *Renderer) RenderReadPixels(rect *Rect) (*Surface, error) {
	var cr *C.SDL_Rect
	if rect != nil {
		cr = rect.cptr()
	}
	cs := C.SDL_RenderReadPixels(r.c, cr)
	if cs == nil {
		return nil, getError()
	}
	return &Surface{c: cs}, nil
}

// --- VSync getter ---

// GetVSync returns the current VSync setting for the renderer.
func (r *Renderer) GetVSync() (int, error) {
	var vsync C.int
	if !C.SDL_GetRenderVSync(r.c, &vsync) {
		return 0, getError()
	}
	return int(vsync), nil
}

// --- Flush ---

// Flush forces the rendering context to flush any pending commands and state.
func (r *Renderer) Flush() error {
	if !C.SDL_FlushRenderer(r.c) {
		return getError()
	}
	return nil
}

// --- Debug text ---

// RenderDebugText draws debug text to the renderer.
func (r *Renderer) RenderDebugText(x, y float32, text string) error {
	ct := C.CString(text)
	defer C.free(unsafe.Pointer(ct))
	if !C.SDL_RenderDebugText(r.c, C.float(x), C.float(y), ct) {
		return getError()
	}
	return nil
}

// RenderDebugTextFormat renders debug text (pre-formatted string).
// Use fmt.Sprintf for formatting before calling this.
func (r *Renderer) RenderDebugTextFormat(x, y float32, text string) error {
	ct := C.CString(text)
	defer C.free(unsafe.Pointer(ct))
	if !C._sdl_render_debug_text_format(r.c, C.float(x), C.float(y), ct) {
		return getError()
	}
	return nil
}

// --- Additional texture functions ---

// GetProperties returns the properties associated with the texture.
func (t *Texture) GetProperties() PropertiesID {
	return PropertiesID(C.SDL_GetTextureProperties(t.c))
}

// SetColorModFloat sets the color modulation for the texture using float values.
func (t *Texture) SetColorModFloat(r, g, b float32) error {
	if !C.SDL_SetTextureColorModFloat(t.c, C.float(r), C.float(g), C.float(b)) {
		return getError()
	}
	return nil
}

// GetColorMod returns the color modulation of the texture.
func (t *Texture) GetColorMod() (uint8, uint8, uint8, error) {
	var r, g, b C.Uint8
	if !C.SDL_GetTextureColorMod(t.c, &r, &g, &b) {
		return 0, 0, 0, getError()
	}
	return uint8(r), uint8(g), uint8(b), nil
}

// GetColorModFloat returns the color modulation of the texture as float values.
func (t *Texture) GetColorModFloat() (float32, float32, float32, error) {
	var r, g, b C.float
	if !C.SDL_GetTextureColorModFloat(t.c, &r, &g, &b) {
		return 0, 0, 0, getError()
	}
	return float32(r), float32(g), float32(b), nil
}

// SetAlphaModFloat sets the alpha modulation for the texture using a float value.
func (t *Texture) SetAlphaModFloat(alpha float32) error {
	if !C.SDL_SetTextureAlphaModFloat(t.c, C.float(alpha)) {
		return getError()
	}
	return nil
}

// GetAlphaMod returns the alpha modulation of the texture.
func (t *Texture) GetAlphaMod() (uint8, error) {
	var alpha C.Uint8
	if !C.SDL_GetTextureAlphaMod(t.c, &alpha) {
		return 0, getError()
	}
	return uint8(alpha), nil
}

// GetAlphaModFloat returns the alpha modulation of the texture as a float value.
func (t *Texture) GetAlphaModFloat() (float32, error) {
	var alpha C.float
	if !C.SDL_GetTextureAlphaModFloat(t.c, &alpha) {
		return 0, getError()
	}
	return float32(alpha), nil
}

// GetBlendMode returns the blend mode of the texture.
func (t *Texture) GetBlendMode() (BlendMode, error) {
	var mode C.SDL_BlendMode
	if !C.SDL_GetTextureBlendMode(t.c, &mode) {
		return 0, getError()
	}
	return BlendMode(mode), nil
}

// GetScaleMode returns the scale mode of the texture.
func (t *Texture) GetScaleMode() (ScaleMode, error) {
	var mode C.SDL_ScaleMode
	if !C.SDL_GetTextureScaleMode(t.c, &mode) {
		return 0, getError()
	}
	return ScaleMode(mode), nil
}

// UpdateYUV updates a rectangle within a planar YV12 or IYUV texture with new pixel data.
func (t *Texture) UpdateYUV(rect *Rect, yPlane []byte, yPitch int, uPlane []byte, uPitch int, vPlane []byte, vPitch int) error {
	var cr *C.SDL_Rect
	if rect != nil {
		cr = rect.cptr()
	}
	if !C.SDL_UpdateYUVTexture(t.c, cr,
		(*C.Uint8)(unsafe.Pointer(&yPlane[0])), C.int(yPitch),
		(*C.Uint8)(unsafe.Pointer(&uPlane[0])), C.int(uPitch),
		(*C.Uint8)(unsafe.Pointer(&vPlane[0])), C.int(vPitch)) {
		return getError()
	}
	return nil
}

// UpdateNV updates a rectangle within a planar NV12 or NV21 texture with new pixels.
func (t *Texture) UpdateNV(rect *Rect, yPlane []byte, yPitch int, uvPlane []byte, uvPitch int) error {
	var cr *C.SDL_Rect
	if rect != nil {
		cr = rect.cptr()
	}
	if !C.SDL_UpdateNVTexture(t.c, cr,
		(*C.Uint8)(unsafe.Pointer(&yPlane[0])), C.int(yPitch),
		(*C.Uint8)(unsafe.Pointer(&uvPlane[0])), C.int(uvPitch)) {
		return getError()
	}
	return nil
}

// LockToSurface locks a portion of the texture for write-only pixel access,
// exposing it as an SDL Surface. The returned surface is freed internally
// after calling Unlock or Destroy.
func (t *Texture) LockToSurface(rect *Rect) (*Surface, error) {
	var cr *C.SDL_Rect
	if rect != nil {
		cr = rect.cptr()
	}
	var cs *C.SDL_Surface
	if !C.SDL_LockTextureToSurface(t.c, cr, &cs) {
		return nil, getError()
	}
	return &Surface{c: cs}, nil
}

// TextureAddressMode represents how texture coordinates outside [0,1] are handled.
type TextureAddressMode int

// Texture address mode constants.
const (
	TEXTURE_ADDRESS_INVALID TextureAddressMode = C.SDL_TEXTURE_ADDRESS_INVALID
	TEXTURE_ADDRESS_AUTO    TextureAddressMode = C.SDL_TEXTURE_ADDRESS_AUTO
	TEXTURE_ADDRESS_CLAMP   TextureAddressMode = C.SDL_TEXTURE_ADDRESS_CLAMP
	TEXTURE_ADDRESS_WRAP    TextureAddressMode = C.SDL_TEXTURE_ADDRESS_WRAP
)

// GPURenderState represents a GPU render state object for the 2D renderer.
type GPURenderState struct {
	c *C.SDL_GPURenderState
}

// GetRenderWindow returns the window associated with the renderer.
func GetRenderWindow(renderer *Renderer) *Window {
	cw := C.SDL_GetRenderWindow(renderer.c)
	if cw == nil {
		return nil
	}
	return &Window{c: cw}
}

// CreateTextureWithProperties creates a texture from a properties set.
func CreateTextureWithProperties(renderer *Renderer, props PropertiesID) (*Texture, error) {
	ct := C.SDL_CreateTextureWithProperties(renderer.c, C.SDL_PropertiesID(props))
	if ct == nil {
		return nil, getError()
	}
	return &Texture{c: ct}, nil
}

// GetDrawBlendMode returns the blend mode used for drawing operations.
func (r *Renderer) GetDrawBlendMode() (BlendMode, error) {
	var mode C.SDL_BlendMode
	if !C.SDL_GetRenderDrawBlendMode(r.c, &mode) {
		return 0, getError()
	}
	return BlendMode(mode), nil
}

// ViewportSet returns whether an explicit viewport is set.
func (r *Renderer) ViewportSet() bool {
	return bool(C.SDL_RenderViewportSet(r.c))
}

// SetDefaultTextureScaleMode sets the default scale mode for new textures.
func (r *Renderer) SetDefaultTextureScaleMode(mode ScaleMode) error {
	if !C.SDL_SetDefaultTextureScaleMode(r.c, C.SDL_ScaleMode(mode)) {
		return getError()
	}
	return nil
}

// GetDefaultTextureScaleMode returns the default scale mode for new textures.
func (r *Renderer) GetDefaultTextureScaleMode() (ScaleMode, error) {
	var mode C.SDL_ScaleMode
	if !C.SDL_GetDefaultTextureScaleMode(r.c, &mode) {
		return 0, getError()
	}
	return ScaleMode(mode), nil
}

// SetTextureAddressMode sets the texture address mode for rendering.
func (r *Renderer) SetTextureAddressMode(uMode, vMode TextureAddressMode) error {
	if !C.SDL_SetRenderTextureAddressMode(r.c, C.SDL_TextureAddressMode(uMode), C.SDL_TextureAddressMode(vMode)) {
		return getError()
	}
	return nil
}

// GetTextureAddressMode returns the texture address mode for rendering.
func (r *Renderer) GetTextureAddressMode() (TextureAddressMode, TextureAddressMode, error) {
	var u, v C.SDL_TextureAddressMode
	if !C.SDL_GetRenderTextureAddressMode(r.c, &u, &v) {
		return 0, 0, getError()
	}
	return TextureAddressMode(u), TextureAddressMode(v), nil
}

// SetPalette sets the palette for a paletted texture.
func (t *Texture) SetPalette(palette *Palette) error {
	var cp *C.SDL_Palette
	if palette != nil {
		cp = palette.c
	}
	if !C.SDL_SetTexturePalette(t.c, cp) {
		return getError()
	}
	return nil
}

// GetPalette returns the palette of a paletted texture.
func (t *Texture) GetPalette() *Palette {
	cp := C.SDL_GetTexturePalette(t.c)
	if cp == nil {
		return nil
	}
	return &Palette{c: cp}
}

// RenderTexture9GridTiled renders a texture as a 9-grid with tiled center.
func (r *Renderer) RenderTexture9GridTiled(texture *Texture, srcrect *FRect, leftWidth, rightWidth, topHeight, bottomHeight, scale float32, dstrect *FRect, tileScale float32) error {
	var csr, cdr *C.SDL_FRect
	if srcrect != nil {
		csr = srcrect.cptr()
	}
	if dstrect != nil {
		cdr = dstrect.cptr()
	}
	if !C.SDL_RenderTexture9GridTiled(r.c, texture.c, csr, C.float(leftWidth), C.float(rightWidth), C.float(topHeight), C.float(bottomHeight), C.float(scale), cdr, C.float(tileScale)) {
		return getError()
	}
	return nil
}

// RenderGeometryRaw renders a list of triangles from raw vertex data.
func (r *Renderer) RenderGeometryRaw(texture *Texture, xyData unsafe.Pointer, xyStride int, colorData unsafe.Pointer, colorStride int, uvData unsafe.Pointer, uvStride int, numVertices int, indices unsafe.Pointer, numIndices int, sizeIndices int) error {
	var ct *C.SDL_Texture
	if texture != nil {
		ct = texture.c
	}
	if !C.SDL_RenderGeometryRaw(r.c, ct, (*C.float)(xyData), C.int(xyStride), (*C.SDL_FColor)(colorData), C.int(colorStride), (*C.float)(uvData), C.int(uvStride), C.int(numVertices), indices, C.int(numIndices), C.int(sizeIndices)) {
		return getError()
	}
	return nil
}

// GetMetalLayer returns the Metal layer associated with the renderer.
func (r *Renderer) GetMetalLayer() unsafe.Pointer {
	return C.SDL_GetRenderMetalLayer(r.c)
}

// GetMetalCommandEncoder returns the Metal command encoder for the current frame.
func (r *Renderer) GetMetalCommandEncoder() unsafe.Pointer {
	return C.SDL_GetRenderMetalCommandEncoder(r.c)
}

// AddVulkanRenderSemaphores adds Vulkan semaphores for the renderer to wait on and signal.
func (r *Renderer) AddVulkanRenderSemaphores(waitStageMask uint32, waitSemaphore, signalSemaphore int64) error {
	if !C.SDL_AddVulkanRenderSemaphores(r.c, C.Uint32(waitStageMask), C.Sint64(waitSemaphore), C.Sint64(signalSemaphore)) {
		return getError()
	}
	return nil
}

// CreateGPURenderer creates a renderer from an existing GPU device.
func CreateGPURenderer(device *GPUDevice, window *Window) (*Renderer, error) {
	cr := C.SDL_CreateGPURenderer(device.c, window.c)
	if cr == nil {
		return nil, getError()
	}
	return &Renderer{c: cr}, nil
}

// GetGPURendererDevice returns the GPU device used by a GPU renderer.
func (r *Renderer) GetGPURendererDevice() *GPUDevice {
	cd := C.SDL_GetGPURendererDevice(r.c)
	if cd == nil {
		return nil
	}
	return &GPUDevice{c: cd}
}

// CreateGPURenderState creates a GPU render state for the 2D renderer.
func CreateGPURenderState(renderer *Renderer, createinfo unsafe.Pointer) (*GPURenderState, error) {
	cs := C.SDL_CreateGPURenderState(renderer.c, (*C.SDL_GPURenderStateCreateInfo)(createinfo))
	if cs == nil {
		return nil, getError()
	}
	return &GPURenderState{c: cs}, nil
}

// SetSamplerBindings sets sampler bindings on the GPU render state.
func (s *GPURenderState) SetSamplerBindings(bindings unsafe.Pointer, numBindings int) error {
	if !C.SDL_SetGPURenderStateSamplerBindings(s.c, C.int(numBindings), (*C.SDL_GPUTextureSamplerBinding)(bindings)) {
		return getError()
	}
	return nil
}

// SetStorageTextures sets storage texture bindings on the GPU render state.
func (s *GPURenderState) SetStorageTextures(textures unsafe.Pointer, numTextures int) error {
	if !C.SDL_SetGPURenderStateStorageTextures(s.c, C.int(numTextures), (**C.SDL_GPUTexture)(textures)) {
		return getError()
	}
	return nil
}

// SetStorageBuffers sets storage buffer bindings on the GPU render state.
func (s *GPURenderState) SetStorageBuffers(buffers unsafe.Pointer, numBuffers int) error {
	if !C.SDL_SetGPURenderStateStorageBuffers(s.c, C.int(numBuffers), (**C.SDL_GPUBuffer)(buffers)) {
		return getError()
	}
	return nil
}

// SetFragmentUniforms sets fragment uniform data on the GPU render state.
func (s *GPURenderState) SetFragmentUniforms(slotIndex uint32, data unsafe.Pointer, length uint32) error {
	if !C.SDL_SetGPURenderStateFragmentUniforms(s.c, C.Uint32(slotIndex), data, C.Uint32(length)) {
		return getError()
	}
	return nil
}

// SetGPURenderState activates a GPU render state on the renderer.
func (r *Renderer) SetGPURenderState(state *GPURenderState) error {
	var cs *C.SDL_GPURenderState
	if state != nil {
		cs = state.c
	}
	if !C.SDL_SetGPURenderState(r.c, cs) {
		return getError()
	}
	return nil
}

// Destroy destroys the GPU render state.
func (s *GPURenderState) Destroy() {
	if s.c != nil {
		C.SDL_DestroyGPURenderState(s.c)
		s.c = nil
	}
}

// Note: SDL_GDKSuspendRenderer and SDL_GDKResumeRenderer are only available
// on the GDK platform (Xbox) and are not wrapped here.
