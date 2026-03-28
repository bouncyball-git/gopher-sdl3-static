package sdl

/*
#include <stdlib.h>
#include <SDL3/SDL.h>
*/
import "C"

import "unsafe"

// --- Opaque handle types ---

// GPUDevice represents the SDL GPU context.
type GPUDevice struct{ c *C.SDL_GPUDevice }

// GPUBuffer represents a GPU buffer for vertices, indices, indirect draw commands, etc.
type GPUBuffer struct{ c *C.SDL_GPUBuffer }

// GPUTransferBuffer is used for transferring data to and from the device.
type GPUTransferBuffer struct{ c *C.SDL_GPUTransferBuffer }

// GPUTexture represents a GPU texture.
type GPUTexture struct{ c *C.SDL_GPUTexture }

// GPUSampler represents a GPU sampler object.
type GPUSampler struct{ c *C.SDL_GPUSampler }

// GPUShader represents a compiled shader object.
type GPUShader struct{ c *C.SDL_GPUShader }

// GPUComputePipeline represents a compute pipeline.
type GPUComputePipeline struct{ c *C.SDL_GPUComputePipeline }

// GPUGraphicsPipeline represents a graphics pipeline.
type GPUGraphicsPipeline struct{ c *C.SDL_GPUGraphicsPipeline }

// GPUCommandBuffer represents a command buffer.
type GPUCommandBuffer struct{ c *C.SDL_GPUCommandBuffer }

// GPURenderPass represents a render pass.
type GPURenderPass struct{ c *C.SDL_GPURenderPass }

// GPUComputePass represents a compute pass.
type GPUComputePass struct{ c *C.SDL_GPUComputePass }

// GPUCopyPass represents a copy pass.
type GPUCopyPass struct{ c *C.SDL_GPUCopyPass }

// GPUFence represents a GPU fence for synchronization.
type GPUFence struct{ c *C.SDL_GPUFence }

// --- Enum types ---

// GPUPrimitiveType specifies the primitive topology of a graphics pipeline.
type GPUPrimitiveType int

// GPU primitive type constants.
const (
	GPU_PRIMITIVETYPE_TRIANGLELIST  GPUPrimitiveType = C.SDL_GPU_PRIMITIVETYPE_TRIANGLELIST
	GPU_PRIMITIVETYPE_TRIANGLESTRIP GPUPrimitiveType = C.SDL_GPU_PRIMITIVETYPE_TRIANGLESTRIP
	GPU_PRIMITIVETYPE_LINELIST      GPUPrimitiveType = C.SDL_GPU_PRIMITIVETYPE_LINELIST
	GPU_PRIMITIVETYPE_LINESTRIP     GPUPrimitiveType = C.SDL_GPU_PRIMITIVETYPE_LINESTRIP
	GPU_PRIMITIVETYPE_POINTLIST     GPUPrimitiveType = C.SDL_GPU_PRIMITIVETYPE_POINTLIST
)

// GPULoadOp specifies how texture contents are treated at the beginning of a render pass.
type GPULoadOp int

// GPU load operation constants.
const (
	GPU_LOADOP_LOAD      GPULoadOp = C.SDL_GPU_LOADOP_LOAD
	GPU_LOADOP_CLEAR     GPULoadOp = C.SDL_GPU_LOADOP_CLEAR
	GPU_LOADOP_DONT_CARE GPULoadOp = C.SDL_GPU_LOADOP_DONT_CARE
)

// GPUStoreOp specifies how texture contents are treated at the end of a render pass.
type GPUStoreOp int

// GPU store operation constants.
const (
	GPU_STOREOP_STORE             GPUStoreOp = C.SDL_GPU_STOREOP_STORE
	GPU_STOREOP_DONT_CARE         GPUStoreOp = C.SDL_GPU_STOREOP_DONT_CARE
	GPU_STOREOP_RESOLVE           GPUStoreOp = C.SDL_GPU_STOREOP_RESOLVE
	GPU_STOREOP_RESOLVE_AND_STORE GPUStoreOp = C.SDL_GPU_STOREOP_RESOLVE_AND_STORE
)

// GPUIndexElementSize specifies the size of elements in an index buffer.
type GPUIndexElementSize int

// GPU index element size constants.
const (
	GPU_INDEXELEMENTSIZE_16BIT GPUIndexElementSize = C.SDL_GPU_INDEXELEMENTSIZE_16BIT
	GPU_INDEXELEMENTSIZE_32BIT GPUIndexElementSize = C.SDL_GPU_INDEXELEMENTSIZE_32BIT
)

// GPUTextureFormat specifies the pixel format of a texture.
type GPUTextureFormat int

// GPU texture format constants.
const (
	GPU_TEXTUREFORMAT_INVALID GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_INVALID

	// Unsigned Normalized Float Color Formats
	GPU_TEXTUREFORMAT_A8_UNORM           GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_A8_UNORM
	GPU_TEXTUREFORMAT_R8_UNORM           GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R8_UNORM
	GPU_TEXTUREFORMAT_R8G8_UNORM         GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R8G8_UNORM
	GPU_TEXTUREFORMAT_R8G8B8A8_UNORM     GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R8G8B8A8_UNORM
	GPU_TEXTUREFORMAT_R16_UNORM          GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R16_UNORM
	GPU_TEXTUREFORMAT_R16G16_UNORM       GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R16G16_UNORM
	GPU_TEXTUREFORMAT_R16G16B16A16_UNORM GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R16G16B16A16_UNORM
	GPU_TEXTUREFORMAT_R10G10B10A2_UNORM  GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R10G10B10A2_UNORM
	GPU_TEXTUREFORMAT_B5G6R5_UNORM       GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_B5G6R5_UNORM
	GPU_TEXTUREFORMAT_B5G5R5A1_UNORM     GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_B5G5R5A1_UNORM
	GPU_TEXTUREFORMAT_B4G4R4A4_UNORM     GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_B4G4R4A4_UNORM
	GPU_TEXTUREFORMAT_B8G8R8A8_UNORM     GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_B8G8R8A8_UNORM

	// Compressed Unsigned Normalized Float Color Formats
	GPU_TEXTUREFORMAT_BC1_RGBA_UNORM GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_BC1_RGBA_UNORM
	GPU_TEXTUREFORMAT_BC2_RGBA_UNORM GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_BC2_RGBA_UNORM
	GPU_TEXTUREFORMAT_BC3_RGBA_UNORM GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_BC3_RGBA_UNORM
	GPU_TEXTUREFORMAT_BC4_R_UNORM    GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_BC4_R_UNORM
	GPU_TEXTUREFORMAT_BC5_RG_UNORM   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_BC5_RG_UNORM
	GPU_TEXTUREFORMAT_BC7_RGBA_UNORM GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_BC7_RGBA_UNORM

	// Compressed Signed/Unsigned Float Color Formats
	GPU_TEXTUREFORMAT_BC6H_RGB_FLOAT  GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_BC6H_RGB_FLOAT
	GPU_TEXTUREFORMAT_BC6H_RGB_UFLOAT GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_BC6H_RGB_UFLOAT

	// Signed Normalized Float Color Formats
	GPU_TEXTUREFORMAT_R8_SNORM           GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R8_SNORM
	GPU_TEXTUREFORMAT_R8G8_SNORM         GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R8G8_SNORM
	GPU_TEXTUREFORMAT_R8G8B8A8_SNORM     GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R8G8B8A8_SNORM
	GPU_TEXTUREFORMAT_R16_SNORM          GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R16_SNORM
	GPU_TEXTUREFORMAT_R16G16_SNORM       GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R16G16_SNORM
	GPU_TEXTUREFORMAT_R16G16B16A16_SNORM GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R16G16B16A16_SNORM

	// Signed Float Color Formats
	GPU_TEXTUREFORMAT_R16_FLOAT          GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R16_FLOAT
	GPU_TEXTUREFORMAT_R16G16_FLOAT       GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R16G16_FLOAT
	GPU_TEXTUREFORMAT_R16G16B16A16_FLOAT GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R16G16B16A16_FLOAT
	GPU_TEXTUREFORMAT_R32_FLOAT          GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R32_FLOAT
	GPU_TEXTUREFORMAT_R32G32_FLOAT       GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R32G32_FLOAT
	GPU_TEXTUREFORMAT_R32G32B32A32_FLOAT GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R32G32B32A32_FLOAT

	// Unsigned Float Color Formats
	GPU_TEXTUREFORMAT_R11G11B10_UFLOAT GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R11G11B10_UFLOAT

	// Unsigned Integer Color Formats
	GPU_TEXTUREFORMAT_R8_UINT            GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R8_UINT
	GPU_TEXTUREFORMAT_R8G8_UINT          GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R8G8_UINT
	GPU_TEXTUREFORMAT_R8G8B8A8_UINT      GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R8G8B8A8_UINT
	GPU_TEXTUREFORMAT_R16_UINT           GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R16_UINT
	GPU_TEXTUREFORMAT_R16G16_UINT        GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R16G16_UINT
	GPU_TEXTUREFORMAT_R16G16B16A16_UINT  GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R16G16B16A16_UINT
	GPU_TEXTUREFORMAT_R32_UINT           GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R32_UINT
	GPU_TEXTUREFORMAT_R32G32_UINT        GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R32G32_UINT
	GPU_TEXTUREFORMAT_R32G32B32A32_UINT  GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R32G32B32A32_UINT

	// Signed Integer Color Formats
	GPU_TEXTUREFORMAT_R8_INT            GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R8_INT
	GPU_TEXTUREFORMAT_R8G8_INT          GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R8G8_INT
	GPU_TEXTUREFORMAT_R8G8B8A8_INT      GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R8G8B8A8_INT
	GPU_TEXTUREFORMAT_R16_INT           GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R16_INT
	GPU_TEXTUREFORMAT_R16G16_INT        GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R16G16_INT
	GPU_TEXTUREFORMAT_R16G16B16A16_INT  GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R16G16B16A16_INT
	GPU_TEXTUREFORMAT_R32_INT           GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R32_INT
	GPU_TEXTUREFORMAT_R32G32_INT        GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R32G32_INT
	GPU_TEXTUREFORMAT_R32G32B32A32_INT  GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R32G32B32A32_INT

	// SRGB Unsigned Normalized Color Formats
	GPU_TEXTUREFORMAT_R8G8B8A8_UNORM_SRGB GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_R8G8B8A8_UNORM_SRGB
	GPU_TEXTUREFORMAT_B8G8R8A8_UNORM_SRGB GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_B8G8R8A8_UNORM_SRGB

	// Compressed SRGB Unsigned Normalized Color Formats
	GPU_TEXTUREFORMAT_BC1_RGBA_UNORM_SRGB GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_BC1_RGBA_UNORM_SRGB
	GPU_TEXTUREFORMAT_BC2_RGBA_UNORM_SRGB GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_BC2_RGBA_UNORM_SRGB
	GPU_TEXTUREFORMAT_BC3_RGBA_UNORM_SRGB GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_BC3_RGBA_UNORM_SRGB
	GPU_TEXTUREFORMAT_BC7_RGBA_UNORM_SRGB GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_BC7_RGBA_UNORM_SRGB

	// Depth Formats
	GPU_TEXTUREFORMAT_D16_UNORM          GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_D16_UNORM
	GPU_TEXTUREFORMAT_D24_UNORM          GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_D24_UNORM
	GPU_TEXTUREFORMAT_D32_FLOAT          GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_D32_FLOAT
	GPU_TEXTUREFORMAT_D24_UNORM_S8_UINT  GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_D24_UNORM_S8_UINT
	GPU_TEXTUREFORMAT_D32_FLOAT_S8_UINT  GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_D32_FLOAT_S8_UINT

	// ASTC formats
	GPU_TEXTUREFORMAT_ASTC_4x4_UNORM   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_4x4_UNORM
	GPU_TEXTUREFORMAT_ASTC_5x4_UNORM   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_5x4_UNORM
	GPU_TEXTUREFORMAT_ASTC_5x5_UNORM   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_5x5_UNORM
	GPU_TEXTUREFORMAT_ASTC_6x5_UNORM   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_6x5_UNORM
	GPU_TEXTUREFORMAT_ASTC_6x6_UNORM   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_6x6_UNORM
	GPU_TEXTUREFORMAT_ASTC_8x5_UNORM   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_8x5_UNORM
	GPU_TEXTUREFORMAT_ASTC_8x6_UNORM   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_8x6_UNORM
	GPU_TEXTUREFORMAT_ASTC_8x8_UNORM   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_8x8_UNORM
	GPU_TEXTUREFORMAT_ASTC_10x5_UNORM  GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_10x5_UNORM
	GPU_TEXTUREFORMAT_ASTC_10x6_UNORM  GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_10x6_UNORM
	GPU_TEXTUREFORMAT_ASTC_10x8_UNORM  GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_10x8_UNORM
	GPU_TEXTUREFORMAT_ASTC_10x10_UNORM GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_10x10_UNORM
	GPU_TEXTUREFORMAT_ASTC_12x10_UNORM GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_12x10_UNORM
	GPU_TEXTUREFORMAT_ASTC_12x12_UNORM GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_12x12_UNORM

	// ASTC SRGB formats
	GPU_TEXTUREFORMAT_ASTC_4x4_UNORM_SRGB   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_4x4_UNORM_SRGB
	GPU_TEXTUREFORMAT_ASTC_5x4_UNORM_SRGB   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_5x4_UNORM_SRGB
	GPU_TEXTUREFORMAT_ASTC_5x5_UNORM_SRGB   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_5x5_UNORM_SRGB
	GPU_TEXTUREFORMAT_ASTC_6x5_UNORM_SRGB   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_6x5_UNORM_SRGB
	GPU_TEXTUREFORMAT_ASTC_6x6_UNORM_SRGB   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_6x6_UNORM_SRGB
	GPU_TEXTUREFORMAT_ASTC_8x5_UNORM_SRGB   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_8x5_UNORM_SRGB
	GPU_TEXTUREFORMAT_ASTC_8x6_UNORM_SRGB   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_8x6_UNORM_SRGB
	GPU_TEXTUREFORMAT_ASTC_8x8_UNORM_SRGB   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_8x8_UNORM_SRGB
	GPU_TEXTUREFORMAT_ASTC_10x5_UNORM_SRGB  GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_10x5_UNORM_SRGB
	GPU_TEXTUREFORMAT_ASTC_10x6_UNORM_SRGB  GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_10x6_UNORM_SRGB
	GPU_TEXTUREFORMAT_ASTC_10x8_UNORM_SRGB  GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_10x8_UNORM_SRGB
	GPU_TEXTUREFORMAT_ASTC_10x10_UNORM_SRGB GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_10x10_UNORM_SRGB
	GPU_TEXTUREFORMAT_ASTC_12x10_UNORM_SRGB GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_12x10_UNORM_SRGB
	GPU_TEXTUREFORMAT_ASTC_12x12_UNORM_SRGB GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_12x12_UNORM_SRGB

	// ASTC Float formats
	GPU_TEXTUREFORMAT_ASTC_4x4_FLOAT   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_4x4_FLOAT
	GPU_TEXTUREFORMAT_ASTC_5x4_FLOAT   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_5x4_FLOAT
	GPU_TEXTUREFORMAT_ASTC_5x5_FLOAT   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_5x5_FLOAT
	GPU_TEXTUREFORMAT_ASTC_6x5_FLOAT   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_6x5_FLOAT
	GPU_TEXTUREFORMAT_ASTC_6x6_FLOAT   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_6x6_FLOAT
	GPU_TEXTUREFORMAT_ASTC_8x5_FLOAT   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_8x5_FLOAT
	GPU_TEXTUREFORMAT_ASTC_8x6_FLOAT   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_8x6_FLOAT
	GPU_TEXTUREFORMAT_ASTC_8x8_FLOAT   GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_8x8_FLOAT
	GPU_TEXTUREFORMAT_ASTC_10x5_FLOAT  GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_10x5_FLOAT
	GPU_TEXTUREFORMAT_ASTC_10x6_FLOAT  GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_10x6_FLOAT
	GPU_TEXTUREFORMAT_ASTC_10x8_FLOAT  GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_10x8_FLOAT
	GPU_TEXTUREFORMAT_ASTC_10x10_FLOAT GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_10x10_FLOAT
	GPU_TEXTUREFORMAT_ASTC_12x10_FLOAT GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_12x10_FLOAT
	GPU_TEXTUREFORMAT_ASTC_12x12_FLOAT GPUTextureFormat = C.SDL_GPU_TEXTUREFORMAT_ASTC_12x12_FLOAT
)

// GPUTextureUsageFlags specifies how a texture is intended to be used.
type GPUTextureUsageFlags uint32

// GPU texture usage flag constants.
const (
	GPU_TEXTUREUSAGE_SAMPLER                                 GPUTextureUsageFlags = C.SDL_GPU_TEXTUREUSAGE_SAMPLER
	GPU_TEXTUREUSAGE_COLOR_TARGET                            GPUTextureUsageFlags = C.SDL_GPU_TEXTUREUSAGE_COLOR_TARGET
	GPU_TEXTUREUSAGE_DEPTH_STENCIL_TARGET                    GPUTextureUsageFlags = C.SDL_GPU_TEXTUREUSAGE_DEPTH_STENCIL_TARGET
	GPU_TEXTUREUSAGE_GRAPHICS_STORAGE_READ                   GPUTextureUsageFlags = C.SDL_GPU_TEXTUREUSAGE_GRAPHICS_STORAGE_READ
	GPU_TEXTUREUSAGE_COMPUTE_STORAGE_READ                    GPUTextureUsageFlags = C.SDL_GPU_TEXTUREUSAGE_COMPUTE_STORAGE_READ
	GPU_TEXTUREUSAGE_COMPUTE_STORAGE_WRITE                   GPUTextureUsageFlags = C.SDL_GPU_TEXTUREUSAGE_COMPUTE_STORAGE_WRITE
	GPU_TEXTUREUSAGE_COMPUTE_STORAGE_SIMULTANEOUS_READ_WRITE GPUTextureUsageFlags = C.SDL_GPU_TEXTUREUSAGE_COMPUTE_STORAGE_SIMULTANEOUS_READ_WRITE
)

// GPUTextureType specifies the type of a texture.
type GPUTextureType int

// GPU texture type constants.
const (
	GPU_TEXTURETYPE_2D         GPUTextureType = C.SDL_GPU_TEXTURETYPE_2D
	GPU_TEXTURETYPE_2D_ARRAY   GPUTextureType = C.SDL_GPU_TEXTURETYPE_2D_ARRAY
	GPU_TEXTURETYPE_3D         GPUTextureType = C.SDL_GPU_TEXTURETYPE_3D
	GPU_TEXTURETYPE_CUBE       GPUTextureType = C.SDL_GPU_TEXTURETYPE_CUBE
	GPU_TEXTURETYPE_CUBE_ARRAY GPUTextureType = C.SDL_GPU_TEXTURETYPE_CUBE_ARRAY
)

// GPUSampleCount specifies the sample count of a texture.
type GPUSampleCount int

// GPU sample count constants.
const (
	GPU_SAMPLECOUNT_1 GPUSampleCount = C.SDL_GPU_SAMPLECOUNT_1
	GPU_SAMPLECOUNT_2 GPUSampleCount = C.SDL_GPU_SAMPLECOUNT_2
	GPU_SAMPLECOUNT_4 GPUSampleCount = C.SDL_GPU_SAMPLECOUNT_4
	GPU_SAMPLECOUNT_8 GPUSampleCount = C.SDL_GPU_SAMPLECOUNT_8
)

// GPUCubeMapFace specifies the face of a cube map.
type GPUCubeMapFace int

// GPU cube map face constants.
const (
	GPU_CUBEMAPFACE_POSITIVEX GPUCubeMapFace = C.SDL_GPU_CUBEMAPFACE_POSITIVEX
	GPU_CUBEMAPFACE_NEGATIVEX GPUCubeMapFace = C.SDL_GPU_CUBEMAPFACE_NEGATIVEX
	GPU_CUBEMAPFACE_POSITIVEY GPUCubeMapFace = C.SDL_GPU_CUBEMAPFACE_POSITIVEY
	GPU_CUBEMAPFACE_NEGATIVEY GPUCubeMapFace = C.SDL_GPU_CUBEMAPFACE_NEGATIVEY
	GPU_CUBEMAPFACE_POSITIVEZ GPUCubeMapFace = C.SDL_GPU_CUBEMAPFACE_POSITIVEZ
	GPU_CUBEMAPFACE_NEGATIVEZ GPUCubeMapFace = C.SDL_GPU_CUBEMAPFACE_NEGATIVEZ
)

// GPUBufferUsageFlags specifies how a buffer is intended to be used.
type GPUBufferUsageFlags uint32

// GPU buffer usage flag constants.
const (
	GPU_BUFFERUSAGE_VERTEX                GPUBufferUsageFlags = C.SDL_GPU_BUFFERUSAGE_VERTEX
	GPU_BUFFERUSAGE_INDEX                 GPUBufferUsageFlags = C.SDL_GPU_BUFFERUSAGE_INDEX
	GPU_BUFFERUSAGE_INDIRECT              GPUBufferUsageFlags = C.SDL_GPU_BUFFERUSAGE_INDIRECT
	GPU_BUFFERUSAGE_GRAPHICS_STORAGE_READ GPUBufferUsageFlags = C.SDL_GPU_BUFFERUSAGE_GRAPHICS_STORAGE_READ
	GPU_BUFFERUSAGE_COMPUTE_STORAGE_READ  GPUBufferUsageFlags = C.SDL_GPU_BUFFERUSAGE_COMPUTE_STORAGE_READ
	GPU_BUFFERUSAGE_COMPUTE_STORAGE_WRITE GPUBufferUsageFlags = C.SDL_GPU_BUFFERUSAGE_COMPUTE_STORAGE_WRITE
)

// GPUTransferBufferUsage specifies how a transfer buffer is intended to be used.
type GPUTransferBufferUsage int

// GPU transfer buffer usage constants.
const (
	GPU_TRANSFERBUFFERUSAGE_UPLOAD   GPUTransferBufferUsage = C.SDL_GPU_TRANSFERBUFFERUSAGE_UPLOAD
	GPU_TRANSFERBUFFERUSAGE_DOWNLOAD GPUTransferBufferUsage = C.SDL_GPU_TRANSFERBUFFERUSAGE_DOWNLOAD
)

// GPUShaderStage specifies which stage a shader program corresponds to.
type GPUShaderStage int

// GPU shader stage constants.
const (
	GPU_SHADERSTAGE_VERTEX   GPUShaderStage = C.SDL_GPU_SHADERSTAGE_VERTEX
	GPU_SHADERSTAGE_FRAGMENT GPUShaderStage = C.SDL_GPU_SHADERSTAGE_FRAGMENT
)

// GPUShaderFormat specifies the format of shader code.
type GPUShaderFormat uint32

// GPU shader format constants.
const (
	GPU_SHADERFORMAT_INVALID  GPUShaderFormat = C.SDL_GPU_SHADERFORMAT_INVALID
	GPU_SHADERFORMAT_PRIVATE  GPUShaderFormat = C.SDL_GPU_SHADERFORMAT_PRIVATE
	GPU_SHADERFORMAT_SPIRV    GPUShaderFormat = C.SDL_GPU_SHADERFORMAT_SPIRV
	GPU_SHADERFORMAT_DXBC     GPUShaderFormat = C.SDL_GPU_SHADERFORMAT_DXBC
	GPU_SHADERFORMAT_DXIL     GPUShaderFormat = C.SDL_GPU_SHADERFORMAT_DXIL
	GPU_SHADERFORMAT_MSL      GPUShaderFormat = C.SDL_GPU_SHADERFORMAT_MSL
	GPU_SHADERFORMAT_METALLIB GPUShaderFormat = C.SDL_GPU_SHADERFORMAT_METALLIB
)

// GPUVertexElementFormat specifies the format of a vertex attribute.
type GPUVertexElementFormat int

// GPU vertex element format constants.
const (
	GPU_VERTEXELEMENTFORMAT_INVALID GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_INVALID

	// Integer formats
	GPU_VERTEXELEMENTFORMAT_INT   GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_INT
	GPU_VERTEXELEMENTFORMAT_INT2  GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_INT2
	GPU_VERTEXELEMENTFORMAT_INT3  GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_INT3
	GPU_VERTEXELEMENTFORMAT_INT4  GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_INT4
	GPU_VERTEXELEMENTFORMAT_UINT  GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_UINT
	GPU_VERTEXELEMENTFORMAT_UINT2 GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_UINT2
	GPU_VERTEXELEMENTFORMAT_UINT3 GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_UINT3
	GPU_VERTEXELEMENTFORMAT_UINT4 GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_UINT4

	// Float formats
	GPU_VERTEXELEMENTFORMAT_FLOAT  GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_FLOAT
	GPU_VERTEXELEMENTFORMAT_FLOAT2 GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_FLOAT2
	GPU_VERTEXELEMENTFORMAT_FLOAT3 GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_FLOAT3
	GPU_VERTEXELEMENTFORMAT_FLOAT4 GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_FLOAT4

	// Byte formats
	GPU_VERTEXELEMENTFORMAT_BYTE2  GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_BYTE2
	GPU_VERTEXELEMENTFORMAT_BYTE4  GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_BYTE4
	GPU_VERTEXELEMENTFORMAT_UBYTE2 GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_UBYTE2
	GPU_VERTEXELEMENTFORMAT_UBYTE4 GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_UBYTE4

	// Normalized byte formats
	GPU_VERTEXELEMENTFORMAT_BYTE2_NORM  GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_BYTE2_NORM
	GPU_VERTEXELEMENTFORMAT_BYTE4_NORM  GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_BYTE4_NORM
	GPU_VERTEXELEMENTFORMAT_UBYTE2_NORM GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_UBYTE2_NORM
	GPU_VERTEXELEMENTFORMAT_UBYTE4_NORM GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_UBYTE4_NORM

	// Short formats
	GPU_VERTEXELEMENTFORMAT_SHORT2  GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_SHORT2
	GPU_VERTEXELEMENTFORMAT_SHORT4  GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_SHORT4
	GPU_VERTEXELEMENTFORMAT_USHORT2 GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_USHORT2
	GPU_VERTEXELEMENTFORMAT_USHORT4 GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_USHORT4

	// Normalized short formats
	GPU_VERTEXELEMENTFORMAT_SHORT2_NORM  GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_SHORT2_NORM
	GPU_VERTEXELEMENTFORMAT_SHORT4_NORM  GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_SHORT4_NORM
	GPU_VERTEXELEMENTFORMAT_USHORT2_NORM GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_USHORT2_NORM
	GPU_VERTEXELEMENTFORMAT_USHORT4_NORM GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_USHORT4_NORM

	// Half-precision float formats
	GPU_VERTEXELEMENTFORMAT_HALF2 GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_HALF2
	GPU_VERTEXELEMENTFORMAT_HALF4 GPUVertexElementFormat = C.SDL_GPU_VERTEXELEMENTFORMAT_HALF4
)

// GPUVertexInputRate specifies the rate at which vertex attributes are pulled from buffers.
type GPUVertexInputRate int

// GPU vertex input rate constants.
const (
	GPU_VERTEXINPUTRATE_VERTEX   GPUVertexInputRate = C.SDL_GPU_VERTEXINPUTRATE_VERTEX
	GPU_VERTEXINPUTRATE_INSTANCE GPUVertexInputRate = C.SDL_GPU_VERTEXINPUTRATE_INSTANCE
)

// GPUFillMode specifies the fill mode of the graphics pipeline.
type GPUFillMode int

// GPU fill mode constants.
const (
	GPU_FILLMODE_FILL GPUFillMode = C.SDL_GPU_FILLMODE_FILL
	GPU_FILLMODE_LINE GPUFillMode = C.SDL_GPU_FILLMODE_LINE
)

// GPUCullMode specifies the facing direction in which triangles will be culled.
type GPUCullMode int

// GPU cull mode constants.
const (
	GPU_CULLMODE_NONE  GPUCullMode = C.SDL_GPU_CULLMODE_NONE
	GPU_CULLMODE_FRONT GPUCullMode = C.SDL_GPU_CULLMODE_FRONT
	GPU_CULLMODE_BACK  GPUCullMode = C.SDL_GPU_CULLMODE_BACK
)

// GPUFrontFace specifies vertex winding for front-facing triangles.
type GPUFrontFace int

// GPU front face constants.
const (
	GPU_FRONTFACE_COUNTER_CLOCKWISE GPUFrontFace = C.SDL_GPU_FRONTFACE_COUNTER_CLOCKWISE
	GPU_FRONTFACE_CLOCKWISE         GPUFrontFace = C.SDL_GPU_FRONTFACE_CLOCKWISE
)

// GPUCompareOp specifies a comparison operator.
type GPUCompareOp int

// GPU comparison operator constants.
const (
	GPU_COMPAREOP_INVALID          GPUCompareOp = C.SDL_GPU_COMPAREOP_INVALID
	GPU_COMPAREOP_NEVER            GPUCompareOp = C.SDL_GPU_COMPAREOP_NEVER
	GPU_COMPAREOP_LESS             GPUCompareOp = C.SDL_GPU_COMPAREOP_LESS
	GPU_COMPAREOP_EQUAL            GPUCompareOp = C.SDL_GPU_COMPAREOP_EQUAL
	GPU_COMPAREOP_LESS_OR_EQUAL    GPUCompareOp = C.SDL_GPU_COMPAREOP_LESS_OR_EQUAL
	GPU_COMPAREOP_GREATER          GPUCompareOp = C.SDL_GPU_COMPAREOP_GREATER
	GPU_COMPAREOP_NOT_EQUAL        GPUCompareOp = C.SDL_GPU_COMPAREOP_NOT_EQUAL
	GPU_COMPAREOP_GREATER_OR_EQUAL GPUCompareOp = C.SDL_GPU_COMPAREOP_GREATER_OR_EQUAL
	GPU_COMPAREOP_ALWAYS           GPUCompareOp = C.SDL_GPU_COMPAREOP_ALWAYS
)

// GPUStencilOp specifies what happens to a stored stencil value.
type GPUStencilOp int

// GPU stencil operation constants.
const (
	GPU_STENCILOP_INVALID             GPUStencilOp = C.SDL_GPU_STENCILOP_INVALID
	GPU_STENCILOP_KEEP                GPUStencilOp = C.SDL_GPU_STENCILOP_KEEP
	GPU_STENCILOP_ZERO                GPUStencilOp = C.SDL_GPU_STENCILOP_ZERO
	GPU_STENCILOP_REPLACE             GPUStencilOp = C.SDL_GPU_STENCILOP_REPLACE
	GPU_STENCILOP_INCREMENT_AND_CLAMP GPUStencilOp = C.SDL_GPU_STENCILOP_INCREMENT_AND_CLAMP
	GPU_STENCILOP_DECREMENT_AND_CLAMP GPUStencilOp = C.SDL_GPU_STENCILOP_DECREMENT_AND_CLAMP
	GPU_STENCILOP_INVERT              GPUStencilOp = C.SDL_GPU_STENCILOP_INVERT
	GPU_STENCILOP_INCREMENT_AND_WRAP  GPUStencilOp = C.SDL_GPU_STENCILOP_INCREMENT_AND_WRAP
	GPU_STENCILOP_DECREMENT_AND_WRAP  GPUStencilOp = C.SDL_GPU_STENCILOP_DECREMENT_AND_WRAP
)

// GPUBlendOp specifies the blend operation for pixels in a render target.
type GPUBlendOp int

// GPU blend operation constants.
const (
	GPU_BLENDOP_INVALID          GPUBlendOp = C.SDL_GPU_BLENDOP_INVALID
	GPU_BLENDOP_ADD              GPUBlendOp = C.SDL_GPU_BLENDOP_ADD
	GPU_BLENDOP_SUBTRACT         GPUBlendOp = C.SDL_GPU_BLENDOP_SUBTRACT
	GPU_BLENDOP_REVERSE_SUBTRACT GPUBlendOp = C.SDL_GPU_BLENDOP_REVERSE_SUBTRACT
	GPU_BLENDOP_MIN              GPUBlendOp = C.SDL_GPU_BLENDOP_MIN
	GPU_BLENDOP_MAX              GPUBlendOp = C.SDL_GPU_BLENDOP_MAX
)

// GPUBlendFactor specifies a blending factor.
type GPUBlendFactor int

// GPU blend factor constants.
const (
	GPU_BLENDFACTOR_INVALID                GPUBlendFactor = C.SDL_GPU_BLENDFACTOR_INVALID
	GPU_BLENDFACTOR_ZERO                   GPUBlendFactor = C.SDL_GPU_BLENDFACTOR_ZERO
	GPU_BLENDFACTOR_ONE                    GPUBlendFactor = C.SDL_GPU_BLENDFACTOR_ONE
	GPU_BLENDFACTOR_SRC_COLOR              GPUBlendFactor = C.SDL_GPU_BLENDFACTOR_SRC_COLOR
	GPU_BLENDFACTOR_ONE_MINUS_SRC_COLOR    GPUBlendFactor = C.SDL_GPU_BLENDFACTOR_ONE_MINUS_SRC_COLOR
	GPU_BLENDFACTOR_DST_COLOR              GPUBlendFactor = C.SDL_GPU_BLENDFACTOR_DST_COLOR
	GPU_BLENDFACTOR_ONE_MINUS_DST_COLOR    GPUBlendFactor = C.SDL_GPU_BLENDFACTOR_ONE_MINUS_DST_COLOR
	GPU_BLENDFACTOR_SRC_ALPHA              GPUBlendFactor = C.SDL_GPU_BLENDFACTOR_SRC_ALPHA
	GPU_BLENDFACTOR_ONE_MINUS_SRC_ALPHA    GPUBlendFactor = C.SDL_GPU_BLENDFACTOR_ONE_MINUS_SRC_ALPHA
	GPU_BLENDFACTOR_DST_ALPHA              GPUBlendFactor = C.SDL_GPU_BLENDFACTOR_DST_ALPHA
	GPU_BLENDFACTOR_ONE_MINUS_DST_ALPHA    GPUBlendFactor = C.SDL_GPU_BLENDFACTOR_ONE_MINUS_DST_ALPHA
	GPU_BLENDFACTOR_CONSTANT_COLOR         GPUBlendFactor = C.SDL_GPU_BLENDFACTOR_CONSTANT_COLOR
	GPU_BLENDFACTOR_ONE_MINUS_CONSTANT_COLOR GPUBlendFactor = C.SDL_GPU_BLENDFACTOR_ONE_MINUS_CONSTANT_COLOR
	GPU_BLENDFACTOR_SRC_ALPHA_SATURATE     GPUBlendFactor = C.SDL_GPU_BLENDFACTOR_SRC_ALPHA_SATURATE
)

// GPUColorComponentFlags specifies which color components are written.
type GPUColorComponentFlags uint8

// GPU color component flag constants.
const (
	GPU_COLORCOMPONENT_R GPUColorComponentFlags = C.SDL_GPU_COLORCOMPONENT_R
	GPU_COLORCOMPONENT_G GPUColorComponentFlags = C.SDL_GPU_COLORCOMPONENT_G
	GPU_COLORCOMPONENT_B GPUColorComponentFlags = C.SDL_GPU_COLORCOMPONENT_B
	GPU_COLORCOMPONENT_A GPUColorComponentFlags = C.SDL_GPU_COLORCOMPONENT_A
)

// GPUFilter specifies a filter operation used by a sampler.
type GPUFilter int

// GPU filter constants.
const (
	GPU_FILTER_NEAREST GPUFilter = C.SDL_GPU_FILTER_NEAREST
	GPU_FILTER_LINEAR  GPUFilter = C.SDL_GPU_FILTER_LINEAR
)

// GPUSamplerMipmapMode specifies a mipmap mode used by a sampler.
type GPUSamplerMipmapMode int

// GPU sampler mipmap mode constants.
const (
	GPU_SAMPLERMIPMAPMODE_NEAREST GPUSamplerMipmapMode = C.SDL_GPU_SAMPLERMIPMAPMODE_NEAREST
	GPU_SAMPLERMIPMAPMODE_LINEAR  GPUSamplerMipmapMode = C.SDL_GPU_SAMPLERMIPMAPMODE_LINEAR
)

// GPUSamplerAddressMode specifies behavior of texture sampling outside 0-1 range.
type GPUSamplerAddressMode int

// GPU sampler address mode constants.
const (
	GPU_SAMPLERADDRESSMODE_REPEAT          GPUSamplerAddressMode = C.SDL_GPU_SAMPLERADDRESSMODE_REPEAT
	GPU_SAMPLERADDRESSMODE_MIRRORED_REPEAT GPUSamplerAddressMode = C.SDL_GPU_SAMPLERADDRESSMODE_MIRRORED_REPEAT
	GPU_SAMPLERADDRESSMODE_CLAMP_TO_EDGE   GPUSamplerAddressMode = C.SDL_GPU_SAMPLERADDRESSMODE_CLAMP_TO_EDGE
)

// GPUPresentMode specifies the timing for presenting swapchain textures.
type GPUPresentMode int

// GPU present mode constants.
const (
	GPU_PRESENTMODE_VSYNC     GPUPresentMode = C.SDL_GPU_PRESENTMODE_VSYNC
	GPU_PRESENTMODE_IMMEDIATE GPUPresentMode = C.SDL_GPU_PRESENTMODE_IMMEDIATE
	GPU_PRESENTMODE_MAILBOX   GPUPresentMode = C.SDL_GPU_PRESENTMODE_MAILBOX
)

// GPUSwapchainComposition specifies the texture format and colorspace of swapchain textures.
type GPUSwapchainComposition int

// GPU swapchain composition constants.
const (
	GPU_SWAPCHAINCOMPOSITION_SDR                 GPUSwapchainComposition = C.SDL_GPU_SWAPCHAINCOMPOSITION_SDR
	GPU_SWAPCHAINCOMPOSITION_SDR_LINEAR          GPUSwapchainComposition = C.SDL_GPU_SWAPCHAINCOMPOSITION_SDR_LINEAR
	GPU_SWAPCHAINCOMPOSITION_HDR_EXTENDED_LINEAR  GPUSwapchainComposition = C.SDL_GPU_SWAPCHAINCOMPOSITION_HDR_EXTENDED_LINEAR
	GPU_SWAPCHAINCOMPOSITION_HDR10_ST2084        GPUSwapchainComposition = C.SDL_GPU_SWAPCHAINCOMPOSITION_HDR10_ST2084
)

// --- Struct types ---

// GPUViewport specifies a viewport.
type GPUViewport struct {
	X        float32
	Y        float32
	W        float32
	H        float32
	MinDepth float32
	MaxDepth float32
}

// GPUTextureRegion specifies a region of a texture.
type GPUTextureRegion struct {
	Texture  *GPUTexture
	MipLevel uint32
	Layer    uint32
	X        uint32
	Y        uint32
	Z        uint32
	W        uint32
	H        uint32
	D        uint32
}

// GPUBlitRegion specifies a region of a texture used in a blit operation.
type GPUBlitRegion struct {
	Texture           *GPUTexture
	MipLevel          uint32
	LayerOrDepthPlane uint32
	X                 uint32
	Y                 uint32
	W                 uint32
	H                 uint32
}

// GPUBufferRegion specifies a region of a buffer.
type GPUBufferRegion struct {
	Buffer *GPUBuffer
	Offset uint32
	Size   uint32
}

// GPUIndirectDrawCommand specifies parameters of an indirect draw command.
type GPUIndirectDrawCommand struct {
	NumVertices   uint32
	NumInstances  uint32
	FirstVertex   uint32
	FirstInstance uint32
}

// GPUIndexedIndirectDrawCommand specifies parameters of an indexed indirect draw command.
type GPUIndexedIndirectDrawCommand struct {
	NumIndices    uint32
	NumInstances  uint32
	FirstIndex    uint32
	VertexOffset  int32
	FirstInstance uint32
}

// GPUIndirectDispatchCommand specifies parameters of an indirect dispatch command.
type GPUIndirectDispatchCommand struct {
	GroupCountX uint32
	GroupCountY uint32
	GroupCountZ uint32
}

// GPUStencilOpState specifies the stencil operation state.
type GPUStencilOpState struct {
	FailOp      GPUStencilOp
	PassOp      GPUStencilOp
	DepthFailOp GPUStencilOp
	CompareOp   GPUCompareOp
}

// GPUColorTargetBlendState specifies the blend state of a color target.
type GPUColorTargetBlendState struct {
	SrcColorBlendFactor GPUBlendFactor
	DstColorBlendFactor GPUBlendFactor
	ColorBlendOp        GPUBlendOp
	SrcAlphaBlendFactor GPUBlendFactor
	DstAlphaBlendFactor GPUBlendFactor
	AlphaBlendOp        GPUBlendOp
	ColorWriteMask      GPUColorComponentFlags
	EnableBlend         bool
	EnableColorWriteMask bool
}

// GPUSamplerCreateInfo specifies the parameters of a sampler.
type GPUSamplerCreateInfo struct {
	MinFilter        GPUFilter
	MagFilter        GPUFilter
	MipmapMode       GPUSamplerMipmapMode
	AddressModeU     GPUSamplerAddressMode
	AddressModeV     GPUSamplerAddressMode
	AddressModeW     GPUSamplerAddressMode
	MipLODBias       float32
	MaxAnisotropy    float32
	CompareOp        GPUCompareOp
	MinLOD           float32
	MaxLOD           float32
	EnableAnisotropy bool
	EnableCompare    bool
	Props            PropertiesID
}

// GPUVertexBufferDescription specifies parameters of vertex buffers used in a graphics pipeline.
type GPUVertexBufferDescription struct {
	Slot             uint32
	Pitch            uint32
	InputRate        GPUVertexInputRate
	InstanceStepRate uint32
}

// GPUVertexAttribute specifies a vertex attribute.
type GPUVertexAttribute struct {
	Location   uint32
	BufferSlot uint32
	Format     GPUVertexElementFormat
	Offset     uint32
}

// GPUVertexInputState specifies the vertex input state of a graphics pipeline.
type GPUVertexInputState struct {
	VertexBufferDescriptions []GPUVertexBufferDescription
	VertexAttributes         []GPUVertexAttribute
}

// GPURasterizerState specifies the rasterizer state.
type GPURasterizerState struct {
	FillMode              GPUFillMode
	CullMode              GPUCullMode
	FrontFace             GPUFrontFace
	DepthBiasConstantFactor float32
	DepthBiasClamp        float32
	DepthBiasSlopeFactor  float32
	EnableDepthBias       bool
	EnableDepthClip       bool
}

// GPUMultisampleState specifies the multisample state.
type GPUMultisampleState struct {
	SampleCount          GPUSampleCount
	SampleMask           uint32
	EnableMask           bool
	EnableAlphaToCoverage bool
}

// GPUDepthStencilState specifies the depth stencil state.
type GPUDepthStencilState struct {
	CompareOp         GPUCompareOp
	BackStencilState  GPUStencilOpState
	FrontStencilState GPUStencilOpState
	CompareMask       uint8
	WriteMask         uint8
	EnableDepthTest   bool
	EnableDepthWrite  bool
	EnableStencilTest bool
}

// GPUColorTargetDescription specifies parameters of color targets in a graphics pipeline.
type GPUColorTargetDescription struct {
	Format     GPUTextureFormat
	BlendState GPUColorTargetBlendState
}

// GPUGraphicsPipelineTargetInfo specifies render target descriptions.
type GPUGraphicsPipelineTargetInfo struct {
	ColorTargetDescriptions []GPUColorTargetDescription
	DepthStencilFormat      GPUTextureFormat
	HasDepthStencilTarget   bool
}

// GPUGraphicsPipelineCreateInfo specifies the parameters of a graphics pipeline.
type GPUGraphicsPipelineCreateInfo struct {
	VertexShader     *GPUShader
	FragmentShader   *GPUShader
	VertexInputState GPUVertexInputState
	PrimitiveType    GPUPrimitiveType
	RasterizerState  GPURasterizerState
	MultisampleState GPUMultisampleState
	DepthStencilState GPUDepthStencilState
	TargetInfo       GPUGraphicsPipelineTargetInfo
	Props            PropertiesID
}

// GPUComputePipelineCreateInfo specifies the parameters of a compute pipeline.
type GPUComputePipelineCreateInfo struct {
	Code                       []byte
	Entrypoint                 string
	Format                     GPUShaderFormat
	NumSamplers                uint32
	NumReadonlyStorageTextures uint32
	NumReadonlyStorageBuffers  uint32
	NumReadWriteStorageTextures uint32
	NumReadWriteStorageBuffers uint32
	NumUniformBuffers          uint32
	ThreadCountX               uint32
	ThreadCountY               uint32
	ThreadCountZ               uint32
	Props                      PropertiesID
}

// GPUShaderCreateInfo specifies code and metadata for creating a shader.
type GPUShaderCreateInfo struct {
	Code               []byte
	Entrypoint         string
	Format             GPUShaderFormat
	Stage              GPUShaderStage
	NumSamplers        uint32
	NumStorageTextures uint32
	NumStorageBuffers  uint32
	NumUniformBuffers  uint32
	Props              PropertiesID
}

// GPUTextureCreateInfo specifies the parameters of a texture.
type GPUTextureCreateInfo struct {
	Type              GPUTextureType
	Format            GPUTextureFormat
	Usage             GPUTextureUsageFlags
	Width             uint32
	Height            uint32
	LayerCountOrDepth uint32
	NumLevels         uint32
	SampleCount       GPUSampleCount
	Props             PropertiesID
}

// GPUBufferCreateInfo specifies the parameters of a buffer.
type GPUBufferCreateInfo struct {
	Usage GPUBufferUsageFlags
	Size  uint32
	Props PropertiesID
}

// GPUTransferBufferCreateInfo specifies the parameters of a transfer buffer.
type GPUTransferBufferCreateInfo struct {
	Usage GPUTransferBufferUsage
	Size  uint32
	Props PropertiesID
}

// GPUColorTargetInfo specifies a color target used by a render pass.
type GPUColorTargetInfo struct {
	Texture             *GPUTexture
	MipLevel            uint32
	LayerOrDepthPlane   uint32
	ClearColor          FColor
	LoadOp              GPULoadOp
	StoreOp             GPUStoreOp
	ResolveTexture      *GPUTexture
	ResolveMipLevel     uint32
	ResolveLayer        uint32
	Cycle               bool
	CycleResolveTexture bool
}

// GPUDepthStencilTargetInfo specifies a depth-stencil target used by a render pass.
type GPUDepthStencilTargetInfo struct {
	Texture        *GPUTexture
	ClearDepth     float32
	LoadOp         GPULoadOp
	StoreOp        GPUStoreOp
	StencilLoadOp  GPULoadOp
	StencilStoreOp GPUStoreOp
	Cycle          bool
	ClearStencil   uint8
	MipLevel       uint8
	Layer          uint8
}

// GPUBlitInfo contains parameters for a blit command.
type GPUBlitInfo struct {
	Source      GPUBlitRegion
	Destination GPUBlitRegion
	LoadOp      GPULoadOp
	ClearColor  FColor
	FlipMode    FlipMode
	Filter      GPUFilter
	Cycle       bool
}

// GPUBufferBinding specifies parameters in a buffer binding call.
type GPUBufferBinding struct {
	Buffer *GPUBuffer
	Offset uint32
}

// GPUTextureSamplerBinding specifies a texture-sampler pair for binding.
type GPUTextureSamplerBinding struct {
	Texture *GPUTexture
	Sampler *GPUSampler
}

// GPUStorageBufferReadWriteBinding specifies parameters for binding a buffer in a compute pass.
type GPUStorageBufferReadWriteBinding struct {
	Buffer *GPUBuffer
	Cycle  bool
}

// GPUStorageTextureReadWriteBinding specifies parameters for binding a texture in a compute pass.
type GPUStorageTextureReadWriteBinding struct {
	Texture  *GPUTexture
	MipLevel uint32
	Layer    uint32
	Cycle    bool
}

// GPUTextureTransferInfo specifies parameters for transferring data to/from a texture.
type GPUTextureTransferInfo struct {
	TransferBuffer *GPUTransferBuffer
	Offset         uint32
	PixelsPerRow   uint32
	RowsPerLayer   uint32
}

// GPUTransferBufferLocation specifies a location in a transfer buffer.
type GPUTransferBufferLocation struct {
	TransferBuffer *GPUTransferBuffer
	Offset         uint32
}

// GPUTextureLocation specifies a location in a texture.
type GPUTextureLocation struct {
	Texture  *GPUTexture
	MipLevel uint32
	Layer    uint32
	X        uint32
	Y        uint32
	Z        uint32
}

// GPUBufferLocation specifies a location in a buffer.
type GPUBufferLocation struct {
	Buffer *GPUBuffer
	Offset uint32
}

// --- Helper conversion functions ---

func (info *GPUColorTargetInfo) ctype() C.SDL_GPUColorTargetInfo {
	var ci C.SDL_GPUColorTargetInfo
	if info.Texture != nil {
		ci.texture = info.Texture.c
	}
	ci.mip_level = C.Uint32(info.MipLevel)
	ci.layer_or_depth_plane = C.Uint32(info.LayerOrDepthPlane)
	ci.clear_color = *(*C.SDL_FColor)(unsafe.Pointer(&info.ClearColor))
	ci.load_op = C.SDL_GPULoadOp(info.LoadOp)
	ci.store_op = C.SDL_GPUStoreOp(info.StoreOp)
	if info.ResolveTexture != nil {
		ci.resolve_texture = info.ResolveTexture.c
	}
	ci.resolve_mip_level = C.Uint32(info.ResolveMipLevel)
	ci.resolve_layer = C.Uint32(info.ResolveLayer)
	ci.cycle = C.bool(info.Cycle)
	ci.cycle_resolve_texture = C.bool(info.CycleResolveTexture)
	return ci
}

func (info *GPUDepthStencilTargetInfo) ctype() C.SDL_GPUDepthStencilTargetInfo {
	var ci C.SDL_GPUDepthStencilTargetInfo
	if info.Texture != nil {
		ci.texture = info.Texture.c
	}
	ci.clear_depth = C.float(info.ClearDepth)
	ci.load_op = C.SDL_GPULoadOp(info.LoadOp)
	ci.store_op = C.SDL_GPUStoreOp(info.StoreOp)
	ci.stencil_load_op = C.SDL_GPULoadOp(info.StencilLoadOp)
	ci.stencil_store_op = C.SDL_GPUStoreOp(info.StencilStoreOp)
	ci.cycle = C.bool(info.Cycle)
	ci.clear_stencil = C.Uint8(info.ClearStencil)
	ci.mip_level = C.Uint8(info.MipLevel)
	ci.layer = C.Uint8(info.Layer)
	return ci
}

func (info *GPUTextureRegion) ctype() C.SDL_GPUTextureRegion {
	var cr C.SDL_GPUTextureRegion
	if info.Texture != nil {
		cr.texture = info.Texture.c
	}
	cr.mip_level = C.Uint32(info.MipLevel)
	cr.layer = C.Uint32(info.Layer)
	cr.x = C.Uint32(info.X)
	cr.y = C.Uint32(info.Y)
	cr.z = C.Uint32(info.Z)
	cr.w = C.Uint32(info.W)
	cr.h = C.Uint32(info.H)
	cr.d = C.Uint32(info.D)
	return cr
}

func (info *GPUBlitRegion) ctype() C.SDL_GPUBlitRegion {
	var cr C.SDL_GPUBlitRegion
	if info.Texture != nil {
		cr.texture = info.Texture.c
	}
	cr.mip_level = C.Uint32(info.MipLevel)
	cr.layer_or_depth_plane = C.Uint32(info.LayerOrDepthPlane)
	cr.x = C.Uint32(info.X)
	cr.y = C.Uint32(info.Y)
	cr.w = C.Uint32(info.W)
	cr.h = C.Uint32(info.H)
	return cr
}

func (info *GPUBufferRegion) ctype() C.SDL_GPUBufferRegion {
	var cr C.SDL_GPUBufferRegion
	if info.Buffer != nil {
		cr.buffer = info.Buffer.c
	}
	cr.offset = C.Uint32(info.Offset)
	cr.size = C.Uint32(info.Size)
	return cr
}

func (info *GPUBufferLocation) ctype() C.SDL_GPUBufferLocation {
	var cl C.SDL_GPUBufferLocation
	if info.Buffer != nil {
		cl.buffer = info.Buffer.c
	}
	cl.offset = C.Uint32(info.Offset)
	return cl
}

func (info *GPUTextureLocation) ctype() C.SDL_GPUTextureLocation {
	var cl C.SDL_GPUTextureLocation
	if info.Texture != nil {
		cl.texture = info.Texture.c
	}
	cl.mip_level = C.Uint32(info.MipLevel)
	cl.layer = C.Uint32(info.Layer)
	cl.x = C.Uint32(info.X)
	cl.y = C.Uint32(info.Y)
	cl.z = C.Uint32(info.Z)
	return cl
}

func (info *GPUTextureTransferInfo) ctype() C.SDL_GPUTextureTransferInfo {
	var ct C.SDL_GPUTextureTransferInfo
	if info.TransferBuffer != nil {
		ct.transfer_buffer = info.TransferBuffer.c
	}
	ct.offset = C.Uint32(info.Offset)
	ct.pixels_per_row = C.Uint32(info.PixelsPerRow)
	ct.rows_per_layer = C.Uint32(info.RowsPerLayer)
	return ct
}

func (info *GPUTransferBufferLocation) ctype() C.SDL_GPUTransferBufferLocation {
	var cl C.SDL_GPUTransferBufferLocation
	if info.TransferBuffer != nil {
		cl.transfer_buffer = info.TransferBuffer.c
	}
	cl.offset = C.Uint32(info.Offset)
	return cl
}

func (info *GPUBufferBinding) ctype() C.SDL_GPUBufferBinding {
	var cb C.SDL_GPUBufferBinding
	if info.Buffer != nil {
		cb.buffer = info.Buffer.c
	}
	cb.offset = C.Uint32(info.Offset)
	return cb
}

func (info *GPUTextureSamplerBinding) ctype() C.SDL_GPUTextureSamplerBinding {
	var cb C.SDL_GPUTextureSamplerBinding
	if info.Texture != nil {
		cb.texture = info.Texture.c
	}
	if info.Sampler != nil {
		cb.sampler = info.Sampler.c
	}
	return cb
}

func (info *GPUStorageBufferReadWriteBinding) ctype() C.SDL_GPUStorageBufferReadWriteBinding {
	var cb C.SDL_GPUStorageBufferReadWriteBinding
	if info.Buffer != nil {
		cb.buffer = info.Buffer.c
	}
	cb.cycle = C.bool(info.Cycle)
	return cb
}

func (info *GPUStorageTextureReadWriteBinding) ctype() C.SDL_GPUStorageTextureReadWriteBinding {
	var cb C.SDL_GPUStorageTextureReadWriteBinding
	if info.Texture != nil {
		cb.texture = info.Texture.c
	}
	cb.mip_level = C.Uint32(info.MipLevel)
	cb.layer = C.Uint32(info.Layer)
	cb.cycle = C.bool(info.Cycle)
	return cb
}

// buildGraphicsPipelineCreateInfo converts Go struct to C struct.
func buildGraphicsPipelineCreateInfo(info *GPUGraphicsPipelineCreateInfo) (C.SDL_GPUGraphicsPipelineCreateInfo, []C.SDL_GPUVertexBufferDescription, []C.SDL_GPUVertexAttribute, []C.SDL_GPUColorTargetDescription) {
	var ci C.SDL_GPUGraphicsPipelineCreateInfo

	if info.VertexShader != nil {
		ci.vertex_shader = info.VertexShader.c
	}
	if info.FragmentShader != nil {
		ci.fragment_shader = info.FragmentShader.c
	}

	// Vertex input state
	var cVertBufDescs []C.SDL_GPUVertexBufferDescription
	for _, vbd := range info.VertexInputState.VertexBufferDescriptions {
		cVertBufDescs = append(cVertBufDescs, C.SDL_GPUVertexBufferDescription{
			slot:               C.Uint32(vbd.Slot),
			pitch:              C.Uint32(vbd.Pitch),
			input_rate:         C.SDL_GPUVertexInputRate(vbd.InputRate),
			instance_step_rate: C.Uint32(vbd.InstanceStepRate),
		})
	}
	var cVertAttrs []C.SDL_GPUVertexAttribute
	for _, va := range info.VertexInputState.VertexAttributes {
		cVertAttrs = append(cVertAttrs, C.SDL_GPUVertexAttribute{
			location:    C.Uint32(va.Location),
			buffer_slot: C.Uint32(va.BufferSlot),
			format:      C.SDL_GPUVertexElementFormat(va.Format),
			offset:      C.Uint32(va.Offset),
		})
	}
	if len(cVertBufDescs) > 0 {
		ci.vertex_input_state.vertex_buffer_descriptions = &cVertBufDescs[0]
	}
	ci.vertex_input_state.num_vertex_buffers = C.Uint32(len(cVertBufDescs))
	if len(cVertAttrs) > 0 {
		ci.vertex_input_state.vertex_attributes = &cVertAttrs[0]
	}
	ci.vertex_input_state.num_vertex_attributes = C.Uint32(len(cVertAttrs))

	ci.primitive_type = C.SDL_GPUPrimitiveType(info.PrimitiveType)

	// Rasterizer state
	ci.rasterizer_state.fill_mode = C.SDL_GPUFillMode(info.RasterizerState.FillMode)
	ci.rasterizer_state.cull_mode = C.SDL_GPUCullMode(info.RasterizerState.CullMode)
	ci.rasterizer_state.front_face = C.SDL_GPUFrontFace(info.RasterizerState.FrontFace)
	ci.rasterizer_state.depth_bias_constant_factor = C.float(info.RasterizerState.DepthBiasConstantFactor)
	ci.rasterizer_state.depth_bias_clamp = C.float(info.RasterizerState.DepthBiasClamp)
	ci.rasterizer_state.depth_bias_slope_factor = C.float(info.RasterizerState.DepthBiasSlopeFactor)
	ci.rasterizer_state.enable_depth_bias = C.bool(info.RasterizerState.EnableDepthBias)
	ci.rasterizer_state.enable_depth_clip = C.bool(info.RasterizerState.EnableDepthClip)

	// Multisample state
	ci.multisample_state.sample_count = C.SDL_GPUSampleCount(info.MultisampleState.SampleCount)
	ci.multisample_state.sample_mask = C.Uint32(info.MultisampleState.SampleMask)
	ci.multisample_state.enable_mask = C.bool(info.MultisampleState.EnableMask)
	ci.multisample_state.enable_alpha_to_coverage = C.bool(info.MultisampleState.EnableAlphaToCoverage)

	// Depth stencil state
	ci.depth_stencil_state.compare_op = C.SDL_GPUCompareOp(info.DepthStencilState.CompareOp)
	ci.depth_stencil_state.back_stencil_state.fail_op = C.SDL_GPUStencilOp(info.DepthStencilState.BackStencilState.FailOp)
	ci.depth_stencil_state.back_stencil_state.pass_op = C.SDL_GPUStencilOp(info.DepthStencilState.BackStencilState.PassOp)
	ci.depth_stencil_state.back_stencil_state.depth_fail_op = C.SDL_GPUStencilOp(info.DepthStencilState.BackStencilState.DepthFailOp)
	ci.depth_stencil_state.back_stencil_state.compare_op = C.SDL_GPUCompareOp(info.DepthStencilState.BackStencilState.CompareOp)
	ci.depth_stencil_state.front_stencil_state.fail_op = C.SDL_GPUStencilOp(info.DepthStencilState.FrontStencilState.FailOp)
	ci.depth_stencil_state.front_stencil_state.pass_op = C.SDL_GPUStencilOp(info.DepthStencilState.FrontStencilState.PassOp)
	ci.depth_stencil_state.front_stencil_state.depth_fail_op = C.SDL_GPUStencilOp(info.DepthStencilState.FrontStencilState.DepthFailOp)
	ci.depth_stencil_state.front_stencil_state.compare_op = C.SDL_GPUCompareOp(info.DepthStencilState.FrontStencilState.CompareOp)
	ci.depth_stencil_state.compare_mask = C.Uint8(info.DepthStencilState.CompareMask)
	ci.depth_stencil_state.write_mask = C.Uint8(info.DepthStencilState.WriteMask)
	ci.depth_stencil_state.enable_depth_test = C.bool(info.DepthStencilState.EnableDepthTest)
	ci.depth_stencil_state.enable_depth_write = C.bool(info.DepthStencilState.EnableDepthWrite)
	ci.depth_stencil_state.enable_stencil_test = C.bool(info.DepthStencilState.EnableStencilTest)

	// Target info
	var cColorTargetDescs []C.SDL_GPUColorTargetDescription
	for _, ctd := range info.TargetInfo.ColorTargetDescriptions {
		var cbs C.SDL_GPUColorTargetBlendState
		cbs.src_color_blendfactor = C.SDL_GPUBlendFactor(ctd.BlendState.SrcColorBlendFactor)
		cbs.dst_color_blendfactor = C.SDL_GPUBlendFactor(ctd.BlendState.DstColorBlendFactor)
		cbs.color_blend_op = C.SDL_GPUBlendOp(ctd.BlendState.ColorBlendOp)
		cbs.src_alpha_blendfactor = C.SDL_GPUBlendFactor(ctd.BlendState.SrcAlphaBlendFactor)
		cbs.dst_alpha_blendfactor = C.SDL_GPUBlendFactor(ctd.BlendState.DstAlphaBlendFactor)
		cbs.alpha_blend_op = C.SDL_GPUBlendOp(ctd.BlendState.AlphaBlendOp)
		cbs.color_write_mask = C.SDL_GPUColorComponentFlags(ctd.BlendState.ColorWriteMask)
		cbs.enable_blend = C.bool(ctd.BlendState.EnableBlend)
		cbs.enable_color_write_mask = C.bool(ctd.BlendState.EnableColorWriteMask)
		cColorTargetDescs = append(cColorTargetDescs, C.SDL_GPUColorTargetDescription{
			format:      C.SDL_GPUTextureFormat(ctd.Format),
			blend_state: cbs,
		})
	}
	if len(cColorTargetDescs) > 0 {
		ci.target_info.color_target_descriptions = &cColorTargetDescs[0]
	}
	ci.target_info.num_color_targets = C.Uint32(len(cColorTargetDescs))
	ci.target_info.depth_stencil_format = C.SDL_GPUTextureFormat(info.TargetInfo.DepthStencilFormat)
	ci.target_info.has_depth_stencil_target = C.bool(info.TargetInfo.HasDepthStencilTarget)

	ci.props = C.SDL_PropertiesID(info.Props)

	return ci, cVertBufDescs, cVertAttrs, cColorTargetDescs
}

// --- Device functions ---

// GPUSupportsShaderFormats checks for GPU runtime support.
func GPUSupportsShaderFormats(formatFlags GPUShaderFormat, name string) bool {
	var cname *C.char
	if name != "" {
		cname = C.CString(name)
		defer C.free(unsafe.Pointer(cname))
	}
	return bool(C.SDL_GPUSupportsShaderFormats(C.SDL_GPUShaderFormat(formatFlags), cname))
}

// GPUSupportsProperties checks for GPU runtime support via properties.
func GPUSupportsProperties(props PropertiesID) bool {
	return bool(C.SDL_GPUSupportsProperties(C.SDL_PropertiesID(props)))
}

// CreateGPUDevice creates a GPU context.
func CreateGPUDevice(formatFlags GPUShaderFormat, debugMode bool, name string) (*GPUDevice, error) {
	var cname *C.char
	if name != "" {
		cname = C.CString(name)
		defer C.free(unsafe.Pointer(cname))
	}
	dev := C.SDL_CreateGPUDevice(C.SDL_GPUShaderFormat(formatFlags), C.bool(debugMode), cname)
	if dev == nil {
		return nil, getError()
	}
	return &GPUDevice{c: dev}, nil
}

// Destroy destroys the GPU device.
func (d *GPUDevice) Destroy() {
	if d.c != nil {
		C.SDL_DestroyGPUDevice(d.c)
		d.c = nil
	}
}

// DestroyGPUDevice destroys a GPU device (package-level function).
func DestroyGPUDevice(device *GPUDevice) {
	device.Destroy()
}

// GetGPUDeviceDriver returns the name of the backend used for this device.
func (d *GPUDevice) GetDriver() string {
	return C.GoString(C.SDL_GetGPUDeviceDriver(d.c))
}

// GetGPUShaderFormats returns the supported shader formats for this device.
func (d *GPUDevice) GetShaderFormats() GPUShaderFormat {
	return GPUShaderFormat(C.SDL_GetGPUShaderFormats(d.c))
}

// --- Window / Swapchain ---

// ClaimWindowForGPUDevice claims a window for GPU rendering.
func ClaimWindowForGPUDevice(device *GPUDevice, window *Window) error {
	if !C.SDL_ClaimWindowForGPUDevice(device.c, window.c) {
		return getError()
	}
	return nil
}

// ReleaseWindowFromGPUDevice unclaims a window.
func ReleaseWindowFromGPUDevice(device *GPUDevice, window *Window) {
	C.SDL_ReleaseWindowFromGPUDevice(device.c, window.c)
}

// SetGPUSwapchainParameters changes swapchain parameters for a claimed window.
func SetGPUSwapchainParameters(device *GPUDevice, window *Window, composition GPUSwapchainComposition, presentMode GPUPresentMode) error {
	if !C.SDL_SetGPUSwapchainParameters(device.c, window.c, C.SDL_GPUSwapchainComposition(composition), C.SDL_GPUPresentMode(presentMode)) {
		return getError()
	}
	return nil
}

// GetGPUSwapchainTextureFormat obtains the texture format of the swapchain.
func GetGPUSwapchainTextureFormat(device *GPUDevice, window *Window) GPUTextureFormat {
	return GPUTextureFormat(C.SDL_GetGPUSwapchainTextureFormat(device.c, window.c))
}

// --- Command buffer ---

// AcquireGPUCommandBuffer acquires a command buffer from the device.
func AcquireGPUCommandBuffer(device *GPUDevice) (*GPUCommandBuffer, error) {
	cb := C.SDL_AcquireGPUCommandBuffer(device.c)
	if cb == nil {
		return nil, getError()
	}
	return &GPUCommandBuffer{c: cb}, nil
}

// Submit submits the command buffer.
func (cb *GPUCommandBuffer) Submit() error {
	if !C.SDL_SubmitGPUCommandBuffer(cb.c) {
		return getError()
	}
	return nil
}

// SubmitGPUCommandBuffer submits a command buffer (package-level function).
func SubmitGPUCommandBuffer(commandBuffer *GPUCommandBuffer) error {
	return commandBuffer.Submit()
}

// SubmitAndAcquireFence submits the command buffer and returns a fence.
func (cb *GPUCommandBuffer) SubmitAndAcquireFence() (*GPUFence, error) {
	f := C.SDL_SubmitGPUCommandBufferAndAcquireFence(cb.c)
	if f == nil {
		return nil, getError()
	}
	return &GPUFence{c: f}, nil
}

// SubmitGPUCommandBufferAndAcquireFence submits and acquires a fence (package-level).
func SubmitGPUCommandBufferAndAcquireFence(commandBuffer *GPUCommandBuffer) (*GPUFence, error) {
	return commandBuffer.SubmitAndAcquireFence()
}

// Cancel cancels the command buffer.
func (cb *GPUCommandBuffer) Cancel() error {
	if !C.SDL_CancelGPUCommandBuffer(cb.c) {
		return getError()
	}
	return nil
}

// CancelGPUCommandBuffer cancels a command buffer (package-level).
func CancelGPUCommandBuffer(commandBuffer *GPUCommandBuffer) error {
	return commandBuffer.Cancel()
}

// AcquireGPUSwapchainTexture acquires a swapchain texture for presentation.
func AcquireGPUSwapchainTexture(commandBuffer *GPUCommandBuffer, window *Window) (*GPUTexture, uint32, uint32, error) {
	var tex *C.SDL_GPUTexture
	var w, h C.Uint32
	if !C.SDL_AcquireGPUSwapchainTexture(commandBuffer.c, window.c, &tex, &w, &h) {
		return nil, 0, 0, getError()
	}
	if tex == nil {
		return nil, uint32(w), uint32(h), nil
	}
	return &GPUTexture{c: tex}, uint32(w), uint32(h), nil
}

// WaitAndAcquireGPUSwapchainTexture blocks until a swapchain texture is available, then acquires it.
func WaitAndAcquireGPUSwapchainTexture(commandBuffer *GPUCommandBuffer, window *Window) (*GPUTexture, uint32, uint32, error) {
	var tex *C.SDL_GPUTexture
	var w, h C.Uint32
	if !C.SDL_WaitAndAcquireGPUSwapchainTexture(commandBuffer.c, window.c, &tex, &w, &h) {
		return nil, 0, 0, getError()
	}
	if tex == nil {
		return nil, uint32(w), uint32(h), nil
	}
	return &GPUTexture{c: tex}, uint32(w), uint32(h), nil
}

// WaitForGPUIdle blocks the thread until the GPU is completely idle.
func WaitForGPUIdle(device *GPUDevice) error {
	if !C.SDL_WaitForGPUIdle(device.c) {
		return getError()
	}
	return nil
}

// --- Shader ---

// CreateGPUShader creates a shader object.
func CreateGPUShader(device *GPUDevice, info *GPUShaderCreateInfo) (*GPUShader, error) {
	centrypoint := C.CString(info.Entrypoint)
	defer C.free(unsafe.Pointer(centrypoint))

	var ccode *C.Uint8
	if len(info.Code) > 0 {
		ccode = (*C.Uint8)(unsafe.Pointer(&info.Code[0]))
	}

	ci := C.SDL_GPUShaderCreateInfo{
		code_size:            C.size_t(len(info.Code)),
		code:                 ccode,
		entrypoint:           centrypoint,
		format:               C.SDL_GPUShaderFormat(info.Format),
		stage:                C.SDL_GPUShaderStage(info.Stage),
		num_samplers:         C.Uint32(info.NumSamplers),
		num_storage_textures: C.Uint32(info.NumStorageTextures),
		num_storage_buffers:  C.Uint32(info.NumStorageBuffers),
		num_uniform_buffers:  C.Uint32(info.NumUniformBuffers),
		props:                C.SDL_PropertiesID(info.Props),
	}

	shader := C.SDL_CreateGPUShader(device.c, &ci)
	if shader == nil {
		return nil, getError()
	}
	return &GPUShader{c: shader}, nil
}

// ReleaseGPUShader frees a shader.
func ReleaseGPUShader(device *GPUDevice, shader *GPUShader) {
	C.SDL_ReleaseGPUShader(device.c, shader.c)
}

// --- Graphics Pipeline ---

// CreateGPUGraphicsPipeline creates a graphics pipeline.
func CreateGPUGraphicsPipeline(device *GPUDevice, info *GPUGraphicsPipelineCreateInfo) (*GPUGraphicsPipeline, error) {
	ci, _, _, _ := buildGraphicsPipelineCreateInfo(info)
	pipeline := C.SDL_CreateGPUGraphicsPipeline(device.c, &ci)
	if pipeline == nil {
		return nil, getError()
	}
	return &GPUGraphicsPipeline{c: pipeline}, nil
}

// ReleaseGPUGraphicsPipeline frees a graphics pipeline.
func ReleaseGPUGraphicsPipeline(device *GPUDevice, pipeline *GPUGraphicsPipeline) {
	C.SDL_ReleaseGPUGraphicsPipeline(device.c, pipeline.c)
}

// --- Compute Pipeline ---

// CreateGPUComputePipeline creates a compute pipeline.
func CreateGPUComputePipeline(device *GPUDevice, info *GPUComputePipelineCreateInfo) (*GPUComputePipeline, error) {
	centrypoint := C.CString(info.Entrypoint)
	defer C.free(unsafe.Pointer(centrypoint))

	var ccode *C.Uint8
	if len(info.Code) > 0 {
		ccode = (*C.Uint8)(unsafe.Pointer(&info.Code[0]))
	}

	ci := C.SDL_GPUComputePipelineCreateInfo{
		code_size:                      C.size_t(len(info.Code)),
		code:                           ccode,
		entrypoint:                     centrypoint,
		format:                         C.SDL_GPUShaderFormat(info.Format),
		num_samplers:                   C.Uint32(info.NumSamplers),
		num_readonly_storage_textures:  C.Uint32(info.NumReadonlyStorageTextures),
		num_readonly_storage_buffers:   C.Uint32(info.NumReadonlyStorageBuffers),
		num_readwrite_storage_textures: C.Uint32(info.NumReadWriteStorageTextures),
		num_readwrite_storage_buffers:  C.Uint32(info.NumReadWriteStorageBuffers),
		num_uniform_buffers:            C.Uint32(info.NumUniformBuffers),
		threadcount_x:                  C.Uint32(info.ThreadCountX),
		threadcount_y:                  C.Uint32(info.ThreadCountY),
		threadcount_z:                  C.Uint32(info.ThreadCountZ),
		props:                          C.SDL_PropertiesID(info.Props),
	}

	pipeline := C.SDL_CreateGPUComputePipeline(device.c, &ci)
	if pipeline == nil {
		return nil, getError()
	}
	return &GPUComputePipeline{c: pipeline}, nil
}

// ReleaseGPUComputePipeline frees a compute pipeline.
func ReleaseGPUComputePipeline(device *GPUDevice, pipeline *GPUComputePipeline) {
	C.SDL_ReleaseGPUComputePipeline(device.c, pipeline.c)
}

// --- Texture ---

// CreateGPUTexture creates a texture object.
func CreateGPUTexture(device *GPUDevice, info *GPUTextureCreateInfo) (*GPUTexture, error) {
	ci := C.SDL_GPUTextureCreateInfo{
		_type:                C.SDL_GPUTextureType(info.Type),
		format:               C.SDL_GPUTextureFormat(info.Format),
		usage:                C.SDL_GPUTextureUsageFlags(info.Usage),
		width:                C.Uint32(info.Width),
		height:               C.Uint32(info.Height),
		layer_count_or_depth: C.Uint32(info.LayerCountOrDepth),
		num_levels:           C.Uint32(info.NumLevels),
		sample_count:         C.SDL_GPUSampleCount(info.SampleCount),
		props:                C.SDL_PropertiesID(info.Props),
	}
	tex := C.SDL_CreateGPUTexture(device.c, &ci)
	if tex == nil {
		return nil, getError()
	}
	return &GPUTexture{c: tex}, nil
}

// ReleaseGPUTexture frees a texture.
func ReleaseGPUTexture(device *GPUDevice, texture *GPUTexture) {
	C.SDL_ReleaseGPUTexture(device.c, texture.c)
}

// --- Buffer ---

// CreateGPUBuffer creates a buffer object.
func CreateGPUBuffer(device *GPUDevice, info *GPUBufferCreateInfo) (*GPUBuffer, error) {
	ci := C.SDL_GPUBufferCreateInfo{
		usage: C.SDL_GPUBufferUsageFlags(info.Usage),
		size:  C.Uint32(info.Size),
		props: C.SDL_PropertiesID(info.Props),
	}
	buf := C.SDL_CreateGPUBuffer(device.c, &ci)
	if buf == nil {
		return nil, getError()
	}
	return &GPUBuffer{c: buf}, nil
}

// ReleaseGPUBuffer frees a buffer.
func ReleaseGPUBuffer(device *GPUDevice, buffer *GPUBuffer) {
	C.SDL_ReleaseGPUBuffer(device.c, buffer.c)
}

// --- Sampler ---

// CreateGPUSampler creates a sampler object.
func CreateGPUSampler(device *GPUDevice, info *GPUSamplerCreateInfo) (*GPUSampler, error) {
	ci := C.SDL_GPUSamplerCreateInfo{
		min_filter:        C.SDL_GPUFilter(info.MinFilter),
		mag_filter:        C.SDL_GPUFilter(info.MagFilter),
		mipmap_mode:       C.SDL_GPUSamplerMipmapMode(info.MipmapMode),
		address_mode_u:    C.SDL_GPUSamplerAddressMode(info.AddressModeU),
		address_mode_v:    C.SDL_GPUSamplerAddressMode(info.AddressModeV),
		address_mode_w:    C.SDL_GPUSamplerAddressMode(info.AddressModeW),
		mip_lod_bias:      C.float(info.MipLODBias),
		max_anisotropy:    C.float(info.MaxAnisotropy),
		compare_op:        C.SDL_GPUCompareOp(info.CompareOp),
		min_lod:           C.float(info.MinLOD),
		max_lod:           C.float(info.MaxLOD),
		enable_anisotropy: C.bool(info.EnableAnisotropy),
		enable_compare:    C.bool(info.EnableCompare),
		props:             C.SDL_PropertiesID(info.Props),
	}
	sampler := C.SDL_CreateGPUSampler(device.c, &ci)
	if sampler == nil {
		return nil, getError()
	}
	return &GPUSampler{c: sampler}, nil
}

// ReleaseGPUSampler frees a sampler.
func ReleaseGPUSampler(device *GPUDevice, sampler *GPUSampler) {
	C.SDL_ReleaseGPUSampler(device.c, sampler.c)
}

// --- Transfer Buffer ---

// CreateGPUTransferBuffer creates a transfer buffer.
func CreateGPUTransferBuffer(device *GPUDevice, info *GPUTransferBufferCreateInfo) (*GPUTransferBuffer, error) {
	ci := C.SDL_GPUTransferBufferCreateInfo{
		usage: C.SDL_GPUTransferBufferUsage(info.Usage),
		size:  C.Uint32(info.Size),
		props: C.SDL_PropertiesID(info.Props),
	}
	tb := C.SDL_CreateGPUTransferBuffer(device.c, &ci)
	if tb == nil {
		return nil, getError()
	}
	return &GPUTransferBuffer{c: tb}, nil
}

// ReleaseGPUTransferBuffer frees a transfer buffer.
func ReleaseGPUTransferBuffer(device *GPUDevice, transferBuffer *GPUTransferBuffer) {
	C.SDL_ReleaseGPUTransferBuffer(device.c, transferBuffer.c)
}

// --- Debug Naming ---

// SetGPUBufferName sets an arbitrary string constant to label a buffer.
func SetGPUBufferName(device *GPUDevice, buffer *GPUBuffer, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.SDL_SetGPUBufferName(device.c, buffer.c, ctext)
}

// SetGPUTextureName sets an arbitrary string constant to label a texture.
func SetGPUTextureName(device *GPUDevice, texture *GPUTexture, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.SDL_SetGPUTextureName(device.c, texture.c, ctext)
}

// --- Render Pass ---

// BeginGPURenderPass begins a render pass on a command buffer.
func BeginGPURenderPass(commandBuffer *GPUCommandBuffer, colorTargetInfos []GPUColorTargetInfo, depthStencilTargetInfo *GPUDepthStencilTargetInfo) *GPURenderPass {
	var cColorTargets []C.SDL_GPUColorTargetInfo
	for i := range colorTargetInfos {
		cColorTargets = append(cColorTargets, colorTargetInfos[i].ctype())
	}
	var pColorTargets *C.SDL_GPUColorTargetInfo
	if len(cColorTargets) > 0 {
		pColorTargets = &cColorTargets[0]
	}

	var pDepthStencil *C.SDL_GPUDepthStencilTargetInfo
	var cDepthStencil C.SDL_GPUDepthStencilTargetInfo
	if depthStencilTargetInfo != nil {
		cDepthStencil = depthStencilTargetInfo.ctype()
		pDepthStencil = &cDepthStencil
	}

	rp := C.SDL_BeginGPURenderPass(commandBuffer.c, pColorTargets, C.Uint32(len(cColorTargets)), pDepthStencil)
	if rp == nil {
		return nil
	}
	return &GPURenderPass{c: rp}
}

// EndGPURenderPass ends the given render pass.
func EndGPURenderPass(renderPass *GPURenderPass) {
	C.SDL_EndGPURenderPass(renderPass.c)
}

// --- Render Pass Bindings ---

// BindGPUGraphicsPipeline binds a graphics pipeline on a render pass.
func BindGPUGraphicsPipeline(renderPass *GPURenderPass, pipeline *GPUGraphicsPipeline) {
	C.SDL_BindGPUGraphicsPipeline(renderPass.c, pipeline.c)
}

// SetGPUViewport sets the current viewport state on a render pass.
func SetGPUViewport(renderPass *GPURenderPass, viewport *GPUViewport) {
	C.SDL_SetGPUViewport(renderPass.c, (*C.SDL_GPUViewport)(unsafe.Pointer(viewport)))
}

// SetGPUScissor sets the current scissor state on a render pass.
func SetGPUScissor(renderPass *GPURenderPass, scissor *Rect) {
	C.SDL_SetGPUScissor(renderPass.c, scissor.cptr())
}

// SetGPUStencilReference sets the current stencil reference value.
func SetGPUStencilReference(renderPass *GPURenderPass, reference uint8) {
	C.SDL_SetGPUStencilReference(renderPass.c, C.Uint8(reference))
}

// BindGPUVertexBuffers binds vertex buffers on a render pass.
func BindGPUVertexBuffers(renderPass *GPURenderPass, firstSlot uint32, bindings []GPUBufferBinding) {
	if len(bindings) == 0 {
		return
	}
	cBindings := make([]C.SDL_GPUBufferBinding, len(bindings))
	for i := range bindings {
		cBindings[i] = bindings[i].ctype()
	}
	C.SDL_BindGPUVertexBuffers(renderPass.c, C.Uint32(firstSlot), &cBindings[0], C.Uint32(len(cBindings)))
}

// BindGPUIndexBuffer binds an index buffer on a render pass.
func BindGPUIndexBuffer(renderPass *GPURenderPass, binding *GPUBufferBinding, indexElementSize GPUIndexElementSize) {
	cb := binding.ctype()
	C.SDL_BindGPUIndexBuffer(renderPass.c, &cb, C.SDL_GPUIndexElementSize(indexElementSize))
}

// BindGPUVertexSamplers binds texture-sampler pairs for use on the vertex shader.
func BindGPUVertexSamplers(renderPass *GPURenderPass, firstSlot uint32, bindings []GPUTextureSamplerBinding) {
	if len(bindings) == 0 {
		return
	}
	cBindings := make([]C.SDL_GPUTextureSamplerBinding, len(bindings))
	for i := range bindings {
		cBindings[i] = bindings[i].ctype()
	}
	C.SDL_BindGPUVertexSamplers(renderPass.c, C.Uint32(firstSlot), &cBindings[0], C.Uint32(len(cBindings)))
}

// BindGPUVertexStorageTextures binds storage textures for use on the vertex shader.
func BindGPUVertexStorageTextures(renderPass *GPURenderPass, firstSlot uint32, textures []*GPUTexture) {
	if len(textures) == 0 {
		return
	}
	cTextures := make([]*C.SDL_GPUTexture, len(textures))
	for i, t := range textures {
		cTextures[i] = t.c
	}
	C.SDL_BindGPUVertexStorageTextures(renderPass.c, C.Uint32(firstSlot), &cTextures[0], C.Uint32(len(cTextures)))
}

// BindGPUVertexStorageBuffers binds storage buffers for use on the vertex shader.
func BindGPUVertexStorageBuffers(renderPass *GPURenderPass, firstSlot uint32, buffers []*GPUBuffer) {
	if len(buffers) == 0 {
		return
	}
	cBuffers := make([]*C.SDL_GPUBuffer, len(buffers))
	for i, b := range buffers {
		cBuffers[i] = b.c
	}
	C.SDL_BindGPUVertexStorageBuffers(renderPass.c, C.Uint32(firstSlot), &cBuffers[0], C.Uint32(len(cBuffers)))
}

// BindGPUFragmentSamplers binds texture-sampler pairs for use on the fragment shader.
func BindGPUFragmentSamplers(renderPass *GPURenderPass, firstSlot uint32, bindings []GPUTextureSamplerBinding) {
	if len(bindings) == 0 {
		return
	}
	cBindings := make([]C.SDL_GPUTextureSamplerBinding, len(bindings))
	for i := range bindings {
		cBindings[i] = bindings[i].ctype()
	}
	C.SDL_BindGPUFragmentSamplers(renderPass.c, C.Uint32(firstSlot), &cBindings[0], C.Uint32(len(cBindings)))
}

// BindGPUFragmentStorageTextures binds storage textures for use on the fragment shader.
func BindGPUFragmentStorageTextures(renderPass *GPURenderPass, firstSlot uint32, textures []*GPUTexture) {
	if len(textures) == 0 {
		return
	}
	cTextures := make([]*C.SDL_GPUTexture, len(textures))
	for i, t := range textures {
		cTextures[i] = t.c
	}
	C.SDL_BindGPUFragmentStorageTextures(renderPass.c, C.Uint32(firstSlot), &cTextures[0], C.Uint32(len(cTextures)))
}

// BindGPUFragmentStorageBuffers binds storage buffers for use on the fragment shader.
func BindGPUFragmentStorageBuffers(renderPass *GPURenderPass, firstSlot uint32, buffers []*GPUBuffer) {
	if len(buffers) == 0 {
		return
	}
	cBuffers := make([]*C.SDL_GPUBuffer, len(buffers))
	for i, b := range buffers {
		cBuffers[i] = b.c
	}
	C.SDL_BindGPUFragmentStorageBuffers(renderPass.c, C.Uint32(firstSlot), &cBuffers[0], C.Uint32(len(cBuffers)))
}

// --- Drawing ---

// DrawGPUPrimitives draws data using bound graphics state.
func DrawGPUPrimitives(renderPass *GPURenderPass, numVertices, numInstances, firstVertex, firstInstance uint32) {
	C.SDL_DrawGPUPrimitives(renderPass.c, C.Uint32(numVertices), C.Uint32(numInstances), C.Uint32(firstVertex), C.Uint32(firstInstance))
}

// DrawGPUPrimitivesIndirect draws data using bound graphics state with parameters from a buffer.
func DrawGPUPrimitivesIndirect(renderPass *GPURenderPass, buffer *GPUBuffer, offset, drawCount uint32) {
	C.SDL_DrawGPUPrimitivesIndirect(renderPass.c, buffer.c, C.Uint32(offset), C.Uint32(drawCount))
}

// DrawGPUIndexedPrimitives draws indexed data using bound graphics state.
func DrawGPUIndexedPrimitives(renderPass *GPURenderPass, numIndices, numInstances, firstIndex uint32, vertexOffset int32, firstInstance uint32) {
	C.SDL_DrawGPUIndexedPrimitives(renderPass.c, C.Uint32(numIndices), C.Uint32(numInstances), C.Uint32(firstIndex), C.Sint32(vertexOffset), C.Uint32(firstInstance))
}

// DrawGPUIndexedPrimitivesIndirect draws indexed data with parameters from a buffer.
func DrawGPUIndexedPrimitivesIndirect(renderPass *GPURenderPass, buffer *GPUBuffer, offset, drawCount uint32) {
	C.SDL_DrawGPUIndexedPrimitivesIndirect(renderPass.c, buffer.c, C.Uint32(offset), C.Uint32(drawCount))
}

// --- Compute Pass ---

// BeginGPUComputePass begins a compute pass on a command buffer.
func BeginGPUComputePass(commandBuffer *GPUCommandBuffer, storageTextureBindings []GPUStorageTextureReadWriteBinding, storageBufferBindings []GPUStorageBufferReadWriteBinding) *GPUComputePass {
	var cTexBindings []C.SDL_GPUStorageTextureReadWriteBinding
	for i := range storageTextureBindings {
		cTexBindings = append(cTexBindings, storageTextureBindings[i].ctype())
	}
	var pTexBindings *C.SDL_GPUStorageTextureReadWriteBinding
	if len(cTexBindings) > 0 {
		pTexBindings = &cTexBindings[0]
	}

	var cBufBindings []C.SDL_GPUStorageBufferReadWriteBinding
	for i := range storageBufferBindings {
		cBufBindings = append(cBufBindings, storageBufferBindings[i].ctype())
	}
	var pBufBindings *C.SDL_GPUStorageBufferReadWriteBinding
	if len(cBufBindings) > 0 {
		pBufBindings = &cBufBindings[0]
	}

	cp := C.SDL_BeginGPUComputePass(commandBuffer.c, pTexBindings, C.Uint32(len(cTexBindings)), pBufBindings, C.Uint32(len(cBufBindings)))
	if cp == nil {
		return nil
	}
	return &GPUComputePass{c: cp}
}

// EndGPUComputePass ends the current compute pass.
func EndGPUComputePass(computePass *GPUComputePass) {
	C.SDL_EndGPUComputePass(computePass.c)
}

// BindGPUComputePipeline binds a compute pipeline.
func BindGPUComputePipeline(computePass *GPUComputePass, pipeline *GPUComputePipeline) {
	C.SDL_BindGPUComputePipeline(computePass.c, pipeline.c)
}

// BindGPUComputeSamplers binds texture-sampler pairs for use on the compute shader.
func BindGPUComputeSamplers(computePass *GPUComputePass, firstSlot uint32, bindings []GPUTextureSamplerBinding) {
	if len(bindings) == 0 {
		return
	}
	cBindings := make([]C.SDL_GPUTextureSamplerBinding, len(bindings))
	for i := range bindings {
		cBindings[i] = bindings[i].ctype()
	}
	C.SDL_BindGPUComputeSamplers(computePass.c, C.Uint32(firstSlot), &cBindings[0], C.Uint32(len(cBindings)))
}

// BindGPUComputeStorageTextures binds storage textures for use on the compute pipeline.
func BindGPUComputeStorageTextures(computePass *GPUComputePass, firstSlot uint32, textures []*GPUTexture) {
	if len(textures) == 0 {
		return
	}
	cTextures := make([]*C.SDL_GPUTexture, len(textures))
	for i, t := range textures {
		cTextures[i] = t.c
	}
	C.SDL_BindGPUComputeStorageTextures(computePass.c, C.Uint32(firstSlot), &cTextures[0], C.Uint32(len(cTextures)))
}

// BindGPUComputeStorageBuffers binds storage buffers for use on the compute pipeline.
func BindGPUComputeStorageBuffers(computePass *GPUComputePass, firstSlot uint32, buffers []*GPUBuffer) {
	if len(buffers) == 0 {
		return
	}
	cBuffers := make([]*C.SDL_GPUBuffer, len(buffers))
	for i, b := range buffers {
		cBuffers[i] = b.c
	}
	C.SDL_BindGPUComputeStorageBuffers(computePass.c, C.Uint32(firstSlot), &cBuffers[0], C.Uint32(len(cBuffers)))
}

// DispatchGPUCompute dispatches compute work.
func DispatchGPUCompute(computePass *GPUComputePass, groupCountX, groupCountY, groupCountZ uint32) {
	C.SDL_DispatchGPUCompute(computePass.c, C.Uint32(groupCountX), C.Uint32(groupCountY), C.Uint32(groupCountZ))
}

// DispatchGPUComputeIndirect dispatches compute work with parameters from a buffer.
func DispatchGPUComputeIndirect(computePass *GPUComputePass, buffer *GPUBuffer, offset uint32) {
	C.SDL_DispatchGPUComputeIndirect(computePass.c, buffer.c, C.Uint32(offset))
}

// --- Copy Pass ---

// BeginGPUCopyPass begins a copy pass on a command buffer.
func BeginGPUCopyPass(commandBuffer *GPUCommandBuffer) *GPUCopyPass {
	cp := C.SDL_BeginGPUCopyPass(commandBuffer.c)
	if cp == nil {
		return nil
	}
	return &GPUCopyPass{c: cp}
}

// EndGPUCopyPass ends the current copy pass.
func EndGPUCopyPass(copyPass *GPUCopyPass) {
	C.SDL_EndGPUCopyPass(copyPass.c)
}

// --- Transfer Buffer Data ---

// MapGPUTransferBuffer maps a transfer buffer into application address space.
func MapGPUTransferBuffer(device *GPUDevice, transferBuffer *GPUTransferBuffer, cycle bool) (unsafe.Pointer, error) {
	ptr := C.SDL_MapGPUTransferBuffer(device.c, transferBuffer.c, C.bool(cycle))
	if ptr == nil {
		return nil, getError()
	}
	return ptr, nil
}

// UnmapGPUTransferBuffer unmaps a previously mapped transfer buffer.
func UnmapGPUTransferBuffer(device *GPUDevice, transferBuffer *GPUTransferBuffer) {
	C.SDL_UnmapGPUTransferBuffer(device.c, transferBuffer.c)
}

// --- Upload / Download / Copy ---

// UploadToGPUTexture uploads data from a transfer buffer to a texture.
func UploadToGPUTexture(copyPass *GPUCopyPass, source *GPUTextureTransferInfo, destination *GPUTextureRegion, cycle bool) {
	cs := source.ctype()
	cd := destination.ctype()
	C.SDL_UploadToGPUTexture(copyPass.c, &cs, &cd, C.bool(cycle))
}

// UploadToGPUBuffer uploads data from a transfer buffer to a buffer.
func UploadToGPUBuffer(copyPass *GPUCopyPass, source *GPUTransferBufferLocation, destination *GPUBufferRegion, cycle bool) {
	cs := source.ctype()
	cd := destination.ctype()
	C.SDL_UploadToGPUBuffer(copyPass.c, &cs, &cd, C.bool(cycle))
}

// DownloadFromGPUTexture copies data from a texture to a transfer buffer.
func DownloadFromGPUTexture(copyPass *GPUCopyPass, source *GPUTextureRegion, destination *GPUTextureTransferInfo) {
	cs := source.ctype()
	cd := destination.ctype()
	C.SDL_DownloadFromGPUTexture(copyPass.c, &cs, &cd)
}

// DownloadFromGPUBuffer copies data from a buffer to a transfer buffer.
func DownloadFromGPUBuffer(copyPass *GPUCopyPass, source *GPUBufferRegion, destination *GPUTransferBufferLocation) {
	cs := source.ctype()
	cd := destination.ctype()
	C.SDL_DownloadFromGPUBuffer(copyPass.c, &cs, &cd)
}

// CopyGPUTextureToTexture performs a texture-to-texture copy.
func CopyGPUTextureToTexture(copyPass *GPUCopyPass, source *GPUTextureLocation, destination *GPUTextureLocation, w, h, d uint32, cycle bool) {
	cs := source.ctype()
	cd := destination.ctype()
	C.SDL_CopyGPUTextureToTexture(copyPass.c, &cs, &cd, C.Uint32(w), C.Uint32(h), C.Uint32(d), C.bool(cycle))
}

// CopyGPUBufferToBuffer performs a buffer-to-buffer copy.
func CopyGPUBufferToBuffer(copyPass *GPUCopyPass, source *GPUBufferLocation, destination *GPUBufferLocation, size uint32, cycle bool) {
	cs := source.ctype()
	cd := destination.ctype()
	C.SDL_CopyGPUBufferToBuffer(copyPass.c, &cs, &cd, C.Uint32(size), C.bool(cycle))
}

// --- Blit / Mipmaps ---

// BlitGPUTexture blits from a source texture region to a destination texture region.
func BlitGPUTexture(commandBuffer *GPUCommandBuffer, info *GPUBlitInfo) {
	var ci C.SDL_GPUBlitInfo
	ci.source = info.Source.ctype()
	ci.destination = info.Destination.ctype()
	ci.load_op = C.SDL_GPULoadOp(info.LoadOp)
	ci.clear_color = *(*C.SDL_FColor)(unsafe.Pointer(&info.ClearColor))
	ci.flip_mode = C.SDL_FlipMode(info.FlipMode)
	ci.filter = C.SDL_GPUFilter(info.Filter)
	ci.cycle = C.bool(info.Cycle)
	C.SDL_BlitGPUTexture(commandBuffer.c, &ci)
}

// GenerateGPUMipmaps generates mipmaps for a texture.
func GenerateGPUMipmaps(commandBuffer *GPUCommandBuffer, texture *GPUTexture) {
	C.SDL_GenerateMipmapsForGPUTexture(commandBuffer.c, texture.c)
}

// --- Fence ---

// QueryGPUFence checks the status of a fence.
func QueryGPUFence(device *GPUDevice, fence *GPUFence) bool {
	return bool(C.SDL_QueryGPUFence(device.c, fence.c))
}

// ReleaseGPUFence releases a fence.
func ReleaseGPUFence(device *GPUDevice, fence *GPUFence) {
	C.SDL_ReleaseGPUFence(device.c, fence.c)
}

// WaitForGPUFences blocks the thread until the given fences are signaled.
func WaitForGPUFences(device *GPUDevice, waitAll bool, fences []*GPUFence) error {
	if len(fences) == 0 {
		return nil
	}
	cFences := make([]*C.SDL_GPUFence, len(fences))
	for i, f := range fences {
		cFences[i] = f.c
	}
	if !C.SDL_WaitForGPUFences(device.c, C.bool(waitAll), &cFences[0], C.Uint32(len(cFences))) {
		return getError()
	}
	return nil
}

// --- Uniform Data ---

// PushGPUVertexUniformData pushes data to a vertex uniform slot.
func PushGPUVertexUniformData(commandBuffer *GPUCommandBuffer, slotIndex uint32, data unsafe.Pointer, length uint32) {
	C.SDL_PushGPUVertexUniformData(commandBuffer.c, C.Uint32(slotIndex), data, C.Uint32(length))
}

// PushGPUFragmentUniformData pushes data to a fragment uniform slot.
func PushGPUFragmentUniformData(commandBuffer *GPUCommandBuffer, slotIndex uint32, data unsafe.Pointer, length uint32) {
	C.SDL_PushGPUFragmentUniformData(commandBuffer.c, C.Uint32(slotIndex), data, C.Uint32(length))
}

// PushGPUComputeUniformData pushes data to a compute uniform slot.
func PushGPUComputeUniformData(commandBuffer *GPUCommandBuffer, slotIndex uint32, data unsafe.Pointer, length uint32) {
	C.SDL_PushGPUComputeUniformData(commandBuffer.c, C.Uint32(slotIndex), data, C.Uint32(length))
}

// --- Format Info ---

// GPUTextureSupportsFormat determines whether a texture format is supported.
func GPUTextureSupportsFormat(device *GPUDevice, format GPUTextureFormat, texType GPUTextureType, usage GPUTextureUsageFlags) bool {
	return bool(C.SDL_GPUTextureSupportsFormat(device.c, C.SDL_GPUTextureFormat(format), C.SDL_GPUTextureType(texType), C.SDL_GPUTextureUsageFlags(usage)))
}

// GPUTextureSupportsSampleCount determines if a sample count for a texture format is supported.
func GPUTextureSupportsSampleCount(device *GPUDevice, format GPUTextureFormat, sampleCount GPUSampleCount) bool {
	return bool(C.SDL_GPUTextureSupportsSampleCount(device.c, C.SDL_GPUTextureFormat(format), C.SDL_GPUSampleCount(sampleCount)))
}

// CreateGPUDeviceWithProperties creates a GPU device with properties.
func CreateGPUDeviceWithProperties(props PropertiesID) (*GPUDevice, error) {
	cd := C.SDL_CreateGPUDeviceWithProperties(C.SDL_PropertiesID(props))
	if cd == nil {
		return nil, getError()
	}
	return &GPUDevice{c: cd}, nil
}

// GetNumGPUDrivers returns the number of GPU drivers compiled into SDL.
func GetNumGPUDrivers() int {
	return int(C.SDL_GetNumGPUDrivers())
}

// GetGPUDriver returns the name of a built-in GPU driver.
func GetGPUDriver(index int) string {
	return C.GoString(C.SDL_GetGPUDriver(C.int(index)))
}

// GetGPUDeviceProperties returns the properties of a GPU device.
func (d *GPUDevice) Properties() PropertiesID {
	return PropertiesID(C.SDL_GetGPUDeviceProperties(d.c))
}

// SetGPUAllowedFramesInFlight sets the maximum number of frames allowed to be in-flight.
func SetGPUAllowedFramesInFlight(device *GPUDevice, allowedFrames uint32) error {
	if !C.SDL_SetGPUAllowedFramesInFlight(device.c, C.Uint32(allowedFrames)) {
		return getError()
	}
	return nil
}

// PushGPUDebugGroup pushes a GPU debug group with a name.
func PushGPUDebugGroup(commandBuffer *GPUCommandBuffer, name string) {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	C.SDL_PushGPUDebugGroup(commandBuffer.c, cn)
}

// PopGPUDebugGroup pops the last GPU debug group.
func PopGPUDebugGroup(commandBuffer *GPUCommandBuffer) {
	C.SDL_PopGPUDebugGroup(commandBuffer.c)
}

// InsertGPUDebugLabel inserts a GPU debug label.
func InsertGPUDebugLabel(commandBuffer *GPUCommandBuffer, text string) {
	ct := C.CString(text)
	defer C.free(unsafe.Pointer(ct))
	C.SDL_InsertGPUDebugLabel(commandBuffer.c, ct)
}

// SetGPUBlendConstants sets the blend constants for the render pass.
func SetGPUBlendConstants(renderPass *GPURenderPass, blendConstants FColor) {
	C.SDL_SetGPUBlendConstants(renderPass.c, C.SDL_FColor{r: C.float(blendConstants.R), g: C.float(blendConstants.G), b: C.float(blendConstants.B), a: C.float(blendConstants.A)})
}

// WaitForGPUSwapchain waits until a swapchain texture is available.
func WaitForGPUSwapchain(device *GPUDevice, window *Window) error {
	if !C.SDL_WaitForGPUSwapchain(device.c, window.c) {
		return getError()
	}
	return nil
}

// WindowSupportsGPUSwapchainComposition returns whether a window supports a swapchain composition mode.
func WindowSupportsGPUSwapchainComposition(device *GPUDevice, window *Window, composition GPUSwapchainComposition) bool {
	return bool(C.SDL_WindowSupportsGPUSwapchainComposition(device.c, window.c, C.SDL_GPUSwapchainComposition(composition)))
}

// WindowSupportsGPUPresentMode returns whether a window supports a present mode.
func WindowSupportsGPUPresentMode(device *GPUDevice, window *Window, mode GPUPresentMode) bool {
	return bool(C.SDL_WindowSupportsGPUPresentMode(device.c, window.c, C.SDL_GPUPresentMode(mode)))
}

// CalculateGPUTextureFormatSize calculates the total size of a texture format given dimensions.
func CalculateGPUTextureFormatSize(format GPUTextureFormat, w, h, depthOrLayers uint32) uint32 {
	return uint32(C.SDL_CalculateGPUTextureFormatSize(C.SDL_GPUTextureFormat(format), C.Uint32(w), C.Uint32(h), C.Uint32(depthOrLayers)))
}

// GPUTextureFormatTexelBlockSize returns the texel block size of a GPU texture format.
func GPUTextureFormatTexelBlockSize(format GPUTextureFormat) uint32 {
	return uint32(C.SDL_GPUTextureFormatTexelBlockSize(C.SDL_GPUTextureFormat(format)))
}

// GetGPUTextureFormatFromPixelFormat converts a pixel format to a GPU texture format.
func GetGPUTextureFormatFromPixelFormat(format PixelFormat) GPUTextureFormat {
	return GPUTextureFormat(C.SDL_GetGPUTextureFormatFromPixelFormat(C.SDL_PixelFormat(format)))
}

// GetPixelFormatFromGPUTextureFormat converts a GPU texture format to a pixel format.
func GetPixelFormatFromGPUTextureFormat(format GPUTextureFormat) PixelFormat {
	return PixelFormat(C.SDL_GetPixelFormatFromGPUTextureFormat(C.SDL_GPUTextureFormat(format)))
}

// Note: SDL_GDKSuspendGPU and SDL_GDKResumeGPU are only available on the GDK platform.

// GPUVulkanOptions specifies Vulkan-specific options for GPU device creation.
type GPUVulkanOptions struct {
	VulkanAPIVersion                uint32
	FeatureList                     unsafe.Pointer
	Vulkan10PhysicalDeviceFeatures  unsafe.Pointer
	DeviceExtensionCount            uint32
	DeviceExtensionNames            unsafe.Pointer
	InstanceExtensionCount          uint32
	InstanceExtensionNames          unsafe.Pointer
}

// Note: The following C #define macros have no Go equivalent:
// - SDL_mutex.h thread safety annotations (SDL_ACQUIRE, SDL_GUARDED_BY, etc.) - C compiler attributes
// - SDL_atomic.h: SDL_CPUP (asm pause instruction), SDL_MEMORY_BARRIER_USES_FUNCTION (compile-time flag)
// - SDL_events.h: SDL_EVENT_ENUM_PADDING (forces enum to uint32 size)
// - SDL_video.h: SDL_DisplayModeData, SDL_GLContextState (SDL-internal types)
// - SDL_video.h: SDL_EGLint (EGL integer type, available via unsafe.Pointer in EGL functions)
