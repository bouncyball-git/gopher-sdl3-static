package sdl

/*
#include <SDL3/SDL.h>
*/
import "C"

type BlendMode uint32

const (
	BLENDMODE_NONE               BlendMode = C.SDL_BLENDMODE_NONE
	BLENDMODE_BLEND              BlendMode = C.SDL_BLENDMODE_BLEND
	BLENDMODE_BLEND_PREMULTIPLIED BlendMode = C.SDL_BLENDMODE_BLEND_PREMULTIPLIED
	BLENDMODE_ADD                BlendMode = C.SDL_BLENDMODE_ADD
	BLENDMODE_ADD_PREMULTIPLIED  BlendMode = C.SDL_BLENDMODE_ADD_PREMULTIPLIED
	BLENDMODE_MOD                BlendMode = C.SDL_BLENDMODE_MOD
	BLENDMODE_MUL                BlendMode = C.SDL_BLENDMODE_MUL
	BLENDMODE_INVALID            BlendMode = C.SDL_BLENDMODE_INVALID
)

type BlendOperation int

const (
	BLENDOPERATION_ADD          BlendOperation = C.SDL_BLENDOPERATION_ADD
	BLENDOPERATION_SUBTRACT     BlendOperation = C.SDL_BLENDOPERATION_SUBTRACT
	BLENDOPERATION_REV_SUBTRACT BlendOperation = C.SDL_BLENDOPERATION_REV_SUBTRACT
	BLENDOPERATION_MINIMUM      BlendOperation = C.SDL_BLENDOPERATION_MINIMUM
	BLENDOPERATION_MAXIMUM      BlendOperation = C.SDL_BLENDOPERATION_MAXIMUM
)

type BlendFactor int

const (
	BLENDFACTOR_ZERO                BlendFactor = C.SDL_BLENDFACTOR_ZERO
	BLENDFACTOR_ONE                 BlendFactor = C.SDL_BLENDFACTOR_ONE
	BLENDFACTOR_SRC_COLOR           BlendFactor = C.SDL_BLENDFACTOR_SRC_COLOR
	BLENDFACTOR_ONE_MINUS_SRC_COLOR BlendFactor = C.SDL_BLENDFACTOR_ONE_MINUS_SRC_COLOR
	BLENDFACTOR_SRC_ALPHA           BlendFactor = C.SDL_BLENDFACTOR_SRC_ALPHA
	BLENDFACTOR_ONE_MINUS_SRC_ALPHA BlendFactor = C.SDL_BLENDFACTOR_ONE_MINUS_SRC_ALPHA
	BLENDFACTOR_DST_COLOR           BlendFactor = C.SDL_BLENDFACTOR_DST_COLOR
	BLENDFACTOR_ONE_MINUS_DST_COLOR BlendFactor = C.SDL_BLENDFACTOR_ONE_MINUS_DST_COLOR
	BLENDFACTOR_DST_ALPHA           BlendFactor = C.SDL_BLENDFACTOR_DST_ALPHA
	BLENDFACTOR_ONE_MINUS_DST_ALPHA BlendFactor = C.SDL_BLENDFACTOR_ONE_MINUS_DST_ALPHA
)

func ComposeCustomBlendMode(srcColorFactor BlendFactor, dstColorFactor BlendFactor, colorOperation BlendOperation, srcAlphaFactor BlendFactor, dstAlphaFactor BlendFactor, alphaOperation BlendOperation) BlendMode {
	return BlendMode(C.SDL_ComposeCustomBlendMode(C.SDL_BlendFactor(srcColorFactor), C.SDL_BlendFactor(dstColorFactor), C.SDL_BlendOperation(colorOperation), C.SDL_BlendFactor(srcAlphaFactor), C.SDL_BlendFactor(dstAlphaFactor), C.SDL_BlendOperation(alphaOperation)))
}
