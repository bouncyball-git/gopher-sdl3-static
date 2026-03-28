package sdl

/*
#include <SDL3/SDL.h>
*/
import "C"

import (
	"math"
	"unsafe"
)

type Point struct{ X, Y int32 }
type FPoint struct{ X, Y float32 }
type Rect struct{ X, Y, W, H int32 }
type FRect struct{ X, Y, W, H float32 }

func (r *Rect) cptr() *C.SDL_Rect     { return (*C.SDL_Rect)(unsafe.Pointer(r)) }
func (r *FRect) cptr() *C.SDL_FRect   { return (*C.SDL_FRect)(unsafe.Pointer(r)) }
func (p *Point) cptr() *C.SDL_Point   { return (*C.SDL_Point)(unsafe.Pointer(p)) }
func (p *FPoint) cptr() *C.SDL_FPoint { return (*C.SDL_FPoint)(unsafe.Pointer(p)) }

func PointInRect(p Point, r Rect) bool {
	return p.X >= r.X && p.X < r.X+r.W && p.Y >= r.Y && p.Y < r.Y+r.H
}

func RectEmpty(r Rect) bool    { return r.W <= 0 || r.H <= 0 }
func RectsEqual(a, b Rect) bool { return a == b }

func HasRectIntersection(a, b Rect) bool {
	return bool(C.SDL_HasRectIntersection(a.cptr(), b.cptr()))
}

func GetRectIntersection(a, b Rect) (Rect, bool) {
	var result Rect
	ok := bool(C.SDL_GetRectIntersection(a.cptr(), b.cptr(), result.cptr()))
	return result, ok
}

func GetRectUnion(a, b Rect) Rect {
	var result Rect
	C.SDL_GetRectUnion(a.cptr(), b.cptr(), result.cptr())
	return result
}

func HasRectIntersectionFloat(a, b FRect) bool {
	return bool(C.SDL_HasRectIntersectionFloat(a.cptr(), b.cptr()))
}

func GetRectIntersectionFloat(a, b FRect) (FRect, bool) {
	var result FRect
	ok := bool(C.SDL_GetRectIntersectionFloat(a.cptr(), b.cptr(), result.cptr()))
	return result, ok
}

func GetRectUnionFloat(a, b FRect) FRect {
	var result FRect
	C.SDL_GetRectUnionFloat(a.cptr(), b.cptr(), result.cptr())
	return result
}

// GetRectEnclosingPoints calculates a minimal rectangle enclosing a set of points.
// If clip is not nil, only points inside the clipping rectangle are considered.
func GetRectEnclosingPoints(points []Point, clip *Rect) (Rect, bool) {
	var result Rect
	var cclip *C.SDL_Rect
	if clip != nil {
		cclip = clip.cptr()
	}
	ok := bool(C.SDL_GetRectEnclosingPoints((*C.SDL_Point)(unsafe.Pointer(&points[0])), C.int(len(points)), cclip, result.cptr()))
	return result, ok
}

// GetRectAndLineIntersection calculates the intersection of a rectangle and line segment.
func GetRectAndLineIntersection(rect Rect, x1, y1, x2, y2 *int32) bool {
	return bool(C.SDL_GetRectAndLineIntersection(rect.cptr(), (*C.int)(unsafe.Pointer(x1)), (*C.int)(unsafe.Pointer(y1)), (*C.int)(unsafe.Pointer(x2)), (*C.int)(unsafe.Pointer(y2))))
}

// PointInRectFloat determines whether a point resides inside a floating point rectangle.
func PointInRectFloat(p FPoint, r FRect) bool {
	return p.X >= r.X && p.X <= r.X+r.W && p.Y >= r.Y && p.Y <= r.Y+r.H
}

// RectEmptyFloat determines whether a floating point rectangle takes no space.
func RectEmptyFloat(r FRect) bool {
	return r.W < 0.0 || r.H < 0.0
}

// RectsEqualFloat determines whether two floating point rectangles are equal,
// within a default epsilon (FLT_EPSILON).
func RectsEqualFloat(a, b FRect) bool {
	const fltEpsilon = 1.1920928955078125e-07
	return RectsEqualEpsilon(a, b, fltEpsilon)
}

// RectsEqualEpsilon determines whether two floating point rectangles are equal,
// within the given epsilon.
func RectsEqualEpsilon(a, b FRect, epsilon float32) bool {
	return (a == b) ||
		(float32(math.Abs(float64(a.X-b.X))) <= epsilon &&
			float32(math.Abs(float64(a.Y-b.Y))) <= epsilon &&
			float32(math.Abs(float64(a.W-b.W))) <= epsilon &&
			float32(math.Abs(float64(a.H-b.H))) <= epsilon)
}

// GetRectEnclosingPointsFloat calculates a minimal rectangle enclosing a set of points with float precision.
func GetRectEnclosingPointsFloat(points []FPoint, clip *FRect) (FRect, bool) {
	var result FRect
	var cclip *C.SDL_FRect
	if clip != nil {
		cclip = clip.cptr()
	}
	ok := bool(C.SDL_GetRectEnclosingPointsFloat((*C.SDL_FPoint)(unsafe.Pointer(&points[0])), C.int(len(points)), cclip, result.cptr()))
	return result, ok
}

// GetRectAndLineIntersectionFloat calculates the intersection of a rectangle and line segment with float precision.
func GetRectAndLineIntersectionFloat(rect FRect, x1, y1, x2, y2 *float32) bool {
	return bool(C.SDL_GetRectAndLineIntersectionFloat(rect.cptr(), (*C.float)(unsafe.Pointer(x1)), (*C.float)(unsafe.Pointer(y1)), (*C.float)(unsafe.Pointer(x2)), (*C.float)(unsafe.Pointer(y2))))
}

// RectToFRect converts a Rect to an FRect.
func RectToFRect(rect Rect) FRect {
	return FRect{
		X: float32(rect.X),
		Y: float32(rect.Y),
		W: float32(rect.W),
		H: float32(rect.H),
	}
}
