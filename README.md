# gopher-sdl3

GO bindings for [SDL3](https://github.com/libsdl-org/SDL). Includes static SDL lib v3.5.0 (linux only)

## Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [API Reference](#api-reference)
  - [Core](#core)
  - [Initialization](#initialization)
  - [Error Handling](#error-handling)
  - [Video / Windows](#video--windows)
  - [2D Rendering](#2d-rendering)
  - [Events](#events)
  - [Pixel Formats & Colors](#pixel-formats--colors)
  - [Surfaces](#surfaces)
  - [Rectangles & Points](#rectangles--points)
  - [Blend Modes](#blend-modes)
  - [Keyboard](#keyboard)
  - [Keycodes](#keycodes)
  - [Scancodes](#scancodes)
  - [Mouse](#mouse)
  - [Touch](#touch)
  - [Pen](#pen)
  - [Gamepad](#gamepad)
  - [Joystick](#joystick)
  - [Haptic / Force Feedback](#haptic--force-feedback)
  - [Sensor](#sensor)
  - [Camera](#camera)
  - [Audio](#audio)
  - [GPU](#gpu)
  - [Vulkan](#vulkan)
  - [Metal](#metal)
  - [Filesystem](#filesystem)
  - [I/O Streams](#io-streams)
  - [Storage](#storage)
  - [Process](#process)
  - [File Dialogs](#file-dialogs)
  - [Clipboard](#clipboard)
  - [Timer](#timer)
  - [Hints](#hints)
  - [Log](#log)
  - [Properties](#properties)
  - [Thread](#thread)
  - [Mutex / Sync](#mutex--sync)
  - [Atomic](#atomic)
  - [HID API](#hid-api)
  - [Async I/O](#async-io)
  - [System Tray](#system-tray)
  - [Locale](#locale)
  - [Power](#power)
  - [Shared Objects](#shared-objects)
  - [Misc](#misc)
  - [Time](#time)
  - [GUID](#guid)
  - [CPU Info](#cpu-info)
  - [Version](#version)
  - [Message Box](#message-box)

## Installation

```bash
go get github.com/bouncyball-git/gopher-sdl3
```

Requires SDL3 installed with pkg-config support:

```bash
export PKG_CONFIG_PATH=/path/to/sdl3/install/lib/pkgconfig
go build ./...
```

## Quick Start

```go
package main

import (
    "fmt"
    "os"
    "github.com/bouncyball-git/gopher-sdl3/sdl"
)

func main() {
    sdl.Init(sdl.INIT_VIDEO)
    defer sdl.Quit()

    window, renderer, _ := sdl.CreateWindowAndRenderer("Hello", 800, 600, sdl.WINDOW_RESIZABLE)
    defer renderer.Destroy()
    defer window.Destroy()

    running := true
    for running {
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
            case *sdl.QuitEvent:
                running = false
            }
        }
        renderer.SetDrawColor(0, 0, 0, 255)
        renderer.Clear()
        renderer.Present()
    }
}
```

## API Reference

### Core

File: `sdl.go`

| Constants | Description |
|-----------|-------------|
| `WINDOWPOS_UNDEFINED` | Undefined window position |
| `WINDOWPOS_CENTERED` | Centered window position |

---

### Initialization

File: `sdl_init.go`

**Types:** `InitFlags`, `AppResult`, `MainThreadCallback`

| Function | Description |
|----------|-------------|
| `Init(flags)` | Initialize SDL subsystems |
| `InitSubSystem(flags)` | Initialize specific subsystems |
| `QuitSubSystem(flags)` | Shut down specific subsystems |
| `WasInit(flags)` | Query initialized subsystems |
| `Quit()` | Shut down all subsystems |
| `IsMainThread()` | Check if on main thread |
| `RunOnMainThread(callback, wait)` | Run callback on main thread |
| `SetAppMetadata(name, version, id)` | Set application metadata |
| `SetAppMetadataProperty(name, value)` | Set metadata property |
| `GetAppMetadataProperty(name)` | Get metadata property |

**Constants:** `INIT_AUDIO`, `INIT_VIDEO`, `INIT_JOYSTICK`, `INIT_HAPTIC`, `INIT_GAMEPAD`, `INIT_EVENTS`, `INIT_SENSOR`, `INIT_CAMERA`, `APP_CONTINUE`, `APP_SUCCESS`, `APP_FAILURE`

---

### Error Handling

File: `sdl_error.go`

| Function | Description |
|----------|-------------|
| `GetError()` | Get last error message |
| `ClearError()` | Clear error message |
| `SetError(msg)` | Set error message |
| `OutOfMemory()` | Set out-of-memory error |

---

### Video / Windows

File: `sdl_video.go`

**Types:** `Window`, `WindowID`, `DisplayID`, `WindowFlags`, `DisplayMode`, `SystemTheme`, `DisplayOrientation`, `FlashOperation`, `GLContext`, `GLAttr`, `HitTestResult`, `ProgressState`

**Window creation:**

| Function | Description |
|----------|-------------|
| `CreateWindow(title, w, h, flags)` | Create a window |
| `CreatePopupWindow(parent, x, y, w, h, flags)` | Create a popup window |
| `CreateWindowWithProperties(props)` | Create window from properties |

**Window methods:**

| Method | Description |
|--------|-------------|
| `Destroy()` | Destroy the window |
| `ID()` | Get window ID |
| `Title()` / `SetTitle()` | Get/set title |
| `Size()` / `SetSize()` | Get/set size |
| `SizeInPixels()` | Get size in pixels |
| `Position()` / `SetPosition()` | Get/set position |
| `Flags()` | Get window flags |
| `SetFullscreen(bool)` | Set fullscreen |
| `SetResizable(bool)` | Set resizable |
| `SetBordered(bool)` | Set bordered |
| `SetAlwaysOnTop(bool)` | Set always-on-top |
| `Show()` / `Hide()` | Show/hide |
| `Raise()` / `Maximize()` / `Minimize()` / `Restore()` | Window state |
| `SetMinimumSize()` / `MinimumSize()` | Min size |
| `SetMaximumSize()` / `MaximumSize()` | Max size |
| `SetAspectRatio()` / `AspectRatio()` | Aspect ratio |
| `Surface()` / `UpdateSurface()` | Window surface |
| `Flash(op)` | Flash window |
| `Display()` / `PixelDensity()` / `DisplayScale()` | Display info |
| `SetIcon(surface)` | Set window icon |
| `SetOpacity()` / `Opacity()` | Opacity |
| `SetParent()` / `Parent()` | Parent window |
| `SetModal(bool)` / `SetFocusable(bool)` | Modal/focus |
| `SetKeyboardGrab()` / `SetMouseGrab()` | Input grab |
| `SetMouseRect()` / `MouseRect()` | Mouse confinement |
| `Properties()` | Window properties |
| `SetProgressState()` / `SetProgressValue()` | Taskbar progress |
| `SetShape()` / `SetHitTest()` | Shape/hit test |
| `SafeArea()` / `BordersSize()` | Layout info |
| `HasSurface()` / `DestroySurface()` | Surface management |
| `SetSurfaceVSync()` / `SurfaceVSync()` | Surface VSync |
| `Sync()` | Synchronize window state |

**Display functions:**

| Function | Description |
|----------|-------------|
| `GetDisplays()` | List displays |
| `GetPrimaryDisplay()` | Get primary display |
| `GetDisplayName(id)` | Display name |
| `GetDisplayBounds(id)` | Display bounds |
| `GetDisplayUsableBounds(id)` | Usable bounds |
| `GetDisplayContentScale(id)` | Content scale |
| `GetDisplayProperties(id)` | Display properties |
| `GetFullscreenDisplayModes(id)` | Fullscreen modes |
| `GetDesktopDisplayMode(id)` | Desktop mode |
| `GetCurrentDisplayMode(id)` | Current mode |
| `GetDisplayForPoint()` / `GetDisplayForRect()` | Display from coords |

**OpenGL functions:** `GL_LoadLibrary`, `GL_UnloadLibrary`, `GL_GetProcAddress`, `GL_ExtensionSupported`, `GL_ResetAttributes`, `GL_SetAttribute`, `GL_GetAttribute`, `GL_CreateContext`, `GL_MakeCurrent`, `GL_GetCurrentWindow`, `GL_GetCurrentContext`, `GL_SetSwapInterval`, `GL_GetSwapInterval`, `GL_SwapWindow`, `GL_DestroyContext`

**EGL functions:** `EGL_GetProcAddress`, `EGL_GetCurrentDisplay`, `EGL_GetCurrentConfig`, `EGL_GetWindowSurface`, `EGL_SetAttributeCallbacks`

**Screen saver:** `ScreenSaverEnabled()`, `EnableScreenSaver()`, `DisableScreenSaver()`

**Window flags:** `WINDOW_FULLSCREEN`, `WINDOW_OPENGL`, `WINDOW_HIDDEN`, `WINDOW_BORDERLESS`, `WINDOW_RESIZABLE`, `WINDOW_MINIMIZED`, `WINDOW_MAXIMIZED`, `WINDOW_MOUSE_GRABBED`, `WINDOW_INPUT_FOCUS`, `WINDOW_MOUSE_FOCUS`, `WINDOW_EXTERNAL`, `WINDOW_MODAL`, `WINDOW_HIGH_PIXEL_DENSITY`, `WINDOW_MOUSE_CAPTURE`, `WINDOW_MOUSE_RELATIVE_MODE`, `WINDOW_ALWAYS_ON_TOP`, `WINDOW_UTILITY`, `WINDOW_TOOLTIP`, `WINDOW_POPUP_MENU`, `WINDOW_KEYBOARD_GRABBED`, `WINDOW_FILL_DOCUMENT`, `WINDOW_VULKAN`, `WINDOW_METAL`, `WINDOW_TRANSPARENT`, `WINDOW_NOT_FOCUSABLE`

---

### 2D Rendering

File: `sdl_render.go`

**Types:** `Renderer`, `Texture`, `TextureAccess`, `RendererLogicalPresentation`, `Vertex`, `TextureAddressMode`, `GPURenderState`

**Renderer creation:**

| Function | Description |
|----------|-------------|
| `CreateWindowAndRenderer(title, w, h, flags)` | Create window + renderer |
| `CreateRenderer(window, name)` | Create renderer for window |
| `CreateRendererWithProperties(props)` | Create from properties |
| `CreateSoftwareRenderer(surface)` | Software renderer |
| `CreateGPURenderer(device, window)` | GPU renderer |

**Drawing methods:**

| Method | Description |
|--------|-------------|
| `Clear()` | Clear target |
| `Present()` | Show rendered content |
| `SetDrawColor(r, g, b, a)` | Set draw color |
| `DrawPoint(x, y)` | Draw point |
| `DrawLine(x1, y1, x2, y2)` | Draw line |
| `DrawLines(points)` | Draw connected lines |
| `DrawRect(rect)` / `DrawRects(rects)` | Draw rectangles |
| `FillRect(rect)` / `FillRects(rects)` | Fill rectangles |
| `RenderTexture(tex, src, dst)` | Copy texture |
| `RenderTextureRotated(...)` | Copy with rotation |
| `RenderTextureAffine(...)` | Affine transform |
| `RenderTextureTiled(...)` | Tiled rendering |
| `RenderTexture9Grid(...)` | 9-grid rendering |
| `RenderTexture9GridTiled(...)` | 9-grid tiled |
| `RenderGeometry(...)` | Triangle geometry |
| `RenderGeometryRaw(...)` | Raw vertex data |
| `RenderPoints(points)` | Multiple points |
| `RenderDebugText(x, y, text)` | Debug text |
| `RenderReadPixels(rect)` | Read pixels |

**Renderer state:** `SetDrawBlendMode`, `GetDrawBlendMode`, `SetScale`, `GetScale`, `SetViewport`, `GetViewport`, `SetClipRect`, `GetClipRect`, `ClipEnabled`, `SetTarget`, `GetTarget`, `SetLogicalPresentation`, `GetLogicalPresentation`, `SetColorScale`, `GetColorScale`, `SetVSync`, `GetVSync`, `Flush`, `SetTextureAddressMode`, `GetTextureAddressMode`, `SetDefaultTextureScaleMode`, `GetDefaultTextureScaleMode`

**Texture functions:** `CreateTexture`, `CreateTextureFromSurface`, `CreateTextureWithProperties`, `Destroy`, `Size`, `SetColorMod`, `GetColorMod`, `SetAlphaMod`, `GetAlphaMod`, `SetBlendMode`, `GetBlendMode`, `SetScaleMode`, `GetScaleMode`, `SetPalette`, `GetPalette`, `Update`, `UpdateYUV`, `UpdateNV`, `Lock`, `LockToSurface`, `Unlock`, `GetProperties`

**Constants:** `TEXTUREACCESS_STATIC`, `TEXTUREACCESS_STREAMING`, `TEXTUREACCESS_TARGET`, `SOFTWARE_RENDERER`, `GPU_RENDERER`, `RENDERER_VSYNC_DISABLED`, `RENDERER_VSYNC_ADAPTIVE`, `DEBUG_TEXT_FONT_CHARACTER_SIZE`

---

### Events

File: `sdl_events.go`

**Types:** `EventType`, `Event` (interface), `CommonEvent`, `QuitEvent`, `DisplayEvent`, `WindowEvent`, `KeyboardEvent`, `TextEditingEvent`, `TextInputEvent`, `TextEditingCandidatesEvent`, `MouseMotionEvent`, `MouseButtonEvent`, `MouseWheelEvent`, `JoystickAxisEvent`, `JoystickButtonEvent`, `JoystickHatEvent`, `JoystickDeviceEvent`, `JoystickBallEvent`, `JoystickBatteryEvent`, `GamepadAxisEvent`, `GamepadButtonEvent`, `GamepadDeviceEvent`, `GamepadTouchpadEvent`, `GamepadSensorEvent`, `TouchFingerEvent`, `PenMotionEvent`, `PenButtonEvent`, `PenTouchEvent`, `PenProximityEvent`, `PenAxisEvent`, `PinchFingerEvent`, `DropEvent`, `AudioDeviceEvent`, `CameraDeviceEvent`, `ClipboardEvent`, `RenderEvent`, `SensorEvent`, `MouseDeviceEvent`, `UserEvent`, `KeyboardDeviceEvent`, `EventFilterFunc`

| Function | Description |
|----------|-------------|
| `PollEvent()` | Poll for event (returns nil if none) |
| `WaitEvent()` | Wait for next event |
| `WaitEventTimeout(ms)` | Wait with timeout |
| `PushEvent(event)` | Push event to queue |
| `PumpEvents()` | Pump event queue |
| `FlushEvent(type)` | Flush events by type |
| `FlushEvents(min, max)` | Flush event range |
| `HasEvent(type)` / `HasEvents(min, max)` | Check for events |
| `SetEventEnabled(type, bool)` / `EventEnabled(type)` | Enable/disable events |
| `RegisterEvents(n)` | Allocate user event types |
| `PeepEvents(...)` | Peek at event queue |
| `GetWindowFromEvent(event)` | Get window from event |
| `GetEventDescription(event)` | Describe event |
| `SetEventFilter(filter)` | Set event filter |
| `GetEventFilter()` | Query event filter |
| `AddEventWatch(filter)` | Add event watch |
| `RemoveEventWatch(id)` | Remove event watch |
| `FilterEvents(filter)` | Filter current queue |

---

### Pixel Formats & Colors

File: `sdl_pixels.go`

**Types:** `PixelFormat`, `Color`, `FColor`, `Colorspace`, `Palette`, `PixelFormatDetails`, `PixelType`, `BitmapOrder`, `PackedOrder`, `ArrayOrder`, `PackedLayout`, `ColorType`, `ColorRange`, `ColorPrimaries`, `TransferCharacteristics`, `MatrixCoefficients`, `ChromaLocation`

| Function | Description |
|----------|-------------|
| `GetPixelFormatName(format)` | Format name string |
| `GetPixelFormatDetails(format)` | Format details |
| `GetMasksForPixelFormat(format)` | Get RGBA masks |
| `GetPixelFormatForMasks(...)` | Get format from masks |
| `CreatePalette(n)` | Create palette |
| `SetPaletteColors(palette, colors, first)` | Set palette colors |
| `DestroyPalette(palette)` | Destroy palette |
| `MapRGB(format, palette, r, g, b)` | Map RGB to pixel |
| `MapRGBA(format, palette, r, g, b, a)` | Map RGBA to pixel |
| `GetRGB(pixel, format, palette)` | Get RGB from pixel |
| `GetRGBA(pixel, format, palette)` | Get RGBA from pixel |

Includes 73+ pixel format constants, 13 colorspace constants, and all pixel type/order/layout enums.

---

### Surfaces

File: `sdl_surface.go`

**Types:** `Surface`, `SurfaceFlags`, `ScaleMode`, `FlipMode`

**Creation:** `CreateSurface`, `CreateSurfaceFrom`, `LoadBMP`, `LoadBMP_IO`, `LoadPNG`, `LoadPNG_IO`, `LoadSurface`, `LoadSurface_IO`

**Methods:** `Destroy`, `Width`, `Height`, `Pitch`, `Format`, `Pixels`, `Lock`, `Unlock`, `SetColorKey`, `HasColorKey`, `GetColorKey`, `FillRect`, `FillRects`, `Blit`, `BlitScaled`, `BlitTiled`, `BlitTiledWithScale`, `Blit9Grid`, `SetBlendMode`, `GetBlendMode`, `SetColorMod`, `GetColorMod`, `SetAlphaMod`, `GetAlphaMod`, `SetClipRect`, `GetClipRect`, `SetRLE`, `HasRLE`, `Properties`, `SetColorspace`, `GetColorspace`, `CreatePalette`, `SetPalette`, `GetPalette`, `Flip`, `Rotate`, `Duplicate`, `Scale`, `Convert`, `Clear`, `MapRGB`, `MapRGBA`, `ReadPixel`, `ReadPixelFloat`, `WritePixel`, `WritePixelFloat`, `SaveBMP`, `SaveBMP_IO`, `SavePNG`, `SavePNG_IO`, `PremultiplyAlpha`

---

### Rectangles & Points

File: `sdl_rect.go`

**Types:** `Point`, `FPoint`, `Rect`, `FRect`

**Functions:** `PointInRect`, `RectEmpty`, `RectsEqual`, `HasRectIntersection`, `GetRectIntersection`, `GetRectUnion`, `GetRectEnclosingPoints`, `GetRectAndLineIntersection`, and float variants.

---

### Blend Modes

File: `sdl_blendmode.go`

**Types:** `BlendMode`, `BlendOperation`, `BlendFactor`

**Constants:** `BLENDMODE_NONE`, `BLENDMODE_BLEND`, `BLENDMODE_ADD`, `BLENDMODE_MOD`, `BLENDMODE_MUL`, `BLENDMODE_INVALID`

**Function:** `ComposeCustomBlendMode(...)`

---

### Keyboard

File: `sdl_keyboard.go`

**Types:** `KeyboardID`, `TextInputType`, `Capitalization`

**Functions:** `HasKeyboard`, `GetKeyboards`, `GetKeyboardNameForID`, `GetKeyboardFocus`, `GetKeyboardState`, `GetModState`, `SetModState`, `ResetKeyboard`, `GetKeyFromScancode`, `GetScancodeFromKey`, `GetScancodeName`, `GetScancodeFromName`, `SetScancodeName`, `GetKeyName`, `GetKeyFromName`, `StartTextInput`, `StopTextInput`, `TextInputActive`, `HasScreenKeyboardSupport`, `ScreenKeyboardShown`, `ClearComposition`, `SetTextInputArea`, `GetTextInputArea`

---

### Keycodes

File: `sdl_keycode.go`

**Types:** `Keycode`, `Keymod`

256 keycode constants (`K_RETURN`, `K_ESCAPE`, `K_SPACE`, `K_A`..`K_Z`, `K_F1`..`K_F24`, etc.), 18 keymod constants (`KMOD_LSHIFT`, `KMOD_CTRL`, etc.), `K_SCANCODE_MASK`, `K_EXTENDED_MASK`.

**Function:** `ScancodeToKeycode(sc)`

---

### Scancodes

File: `sdl_scancode.go`

**Type:** `Scancode`

~250 scancode constants (`SCANCODE_A`..`SCANCODE_Z`, `SCANCODE_1`..`SCANCODE_0`, `SCANCODE_F1`..`SCANCODE_F24`, `SCANCODE_RETURN`, `SCANCODE_ESCAPE`, etc.)

---

### Mouse

File: `sdl_mouse.go`

**Types:** `MouseID`, `Cursor`, `SystemCursor`, `MouseWheelDirection`, `MouseButtonFlags`, `MouseMotionTransformFunc`

**Functions:** `HasMouse`, `GetMice`, `GetMouseFocus`, `GetMouseState`, `GetGlobalMouseState`, `GetRelativeMouseState`, `WarpMouseInWindow`, `WarpMouseGlobal`, `SetWindowRelativeMouseMode`, `GetWindowRelativeMouseMode`, `CaptureMouse`, `CreateSystemCursor`, `CreateColorCursor`, `CreateCursor`, `GetDefaultCursor`, `SetCursor`, `GetCursor`, `ShowCursor`, `HideCursor`, `CursorVisible`, `GetMouseNameForID`, `SetRelativeMouseTransform`

**Cursor.** `Destroy()`

**Constants:** `BUTTON_LEFT`, `BUTTON_MIDDLE`, `BUTTON_RIGHT`, `BUTTON_X1`, `BUTTON_X2`, `BUTTON_LMASK`..`BUTTON_X2MASK`, 20 `SYSTEM_CURSOR_*` values

---

### Touch

File: `sdl_touch.go`

**Types:** `TouchID`, `FingerID`, `TouchDeviceType`, `Finger`

**Functions:** `GetTouchDevices`, `GetTouchDeviceName`, `GetTouchDeviceType`, `GetTouchFingers`

**Constants:** `TOUCH_MOUSEID`, `MOUSE_TOUCHID`

---

### Pen

File: `sdl_pen.go`

**Types:** `PenID`, `PenInputFlags`, `PenAxis`, `PenDeviceType`

**Function:** `GetPenDeviceType(id)`

**Constants:** `PEN_AXIS_*` (8), `PEN_INPUT_*` (8), `PEN_DEVICE_TYPE_*` (3), `PEN_MOUSEID`, `PEN_TOUCHID`

---

### Gamepad

File: `sdl_gamepad.go`

**Types:** `Gamepad`, `GamepadType`, `GamepadButton`, `GamepadAxis`, `GamepadButtonLabel`, `GamepadBindingType`

**Functions:** `HasGamepad`, `GetGamepads`, `IsGamepad`, `OpenGamepad`, `GetGamepadFromID`, `GetGamepadFromPlayerIndex`, `AddGamepadMapping`, `AddGamepadMappingsFromFile`, `AddGamepadMappingsFromIO`, `ReloadGamepadMappings`, `GetGamepadMappings`, `GetGamepadMapping`, `SetGamepadMapping`, `GetGamepadMappingForGUID`, `GetGamepadMappingForID`, `GetGamepadNameForID`, `GetGamepadPathForID`, `GetGamepadTypeForID`, `GetRealGamepadTypeForID`, and 15+ more ForID functions.

**Gamepad methods:** `Close`, `ID`, `Name`, `Path`, `Type`, `RealType`, `PlayerIndex`, `SetPlayerIndex`, `Vendor`, `Product`, `Serial`, `Connected`, `ConnectionState`, `PowerInfo`, `GetGamepadAxis`, `GetGamepadButton`, `GamepadHasAxis`, `GamepadHasButton`, `RumbleGamepad`, `RumbleGamepadTriggers`, `SetGamepadLED`, `SendGamepadEffect`, sensor/touchpad support.

**String conversion:** `GetGamepadTypeFromString`, `GetGamepadStringForType`, `GetGamepadAxisFromString`, `GetGamepadStringForAxis`, `GetGamepadButtonFromString`, `GetGamepadStringForButton`

---

### Joystick

File: `sdl_joystick.go`

**Types:** `Joystick`, `JoystickID`, `JoystickType`, `JoystickConnectionState`, `PowerState`, `VirtualJoystickTouchpadDesc`, `VirtualJoystickSensorDesc`

**Functions:** `HasJoystick`, `GetJoysticks`, `OpenJoystick`, `GetJoystickFromID`, `GetJoystickFromPlayerIndex`, `LockJoysticks`, `UnlockJoysticks`, `TryLockJoysticks`, `SetJoystickEventsEnabled`, `JoystickEventsEnabled`, `UpdateJoysticks`, `AttachVirtualJoystick`, `DetachVirtualJoystick`, `IsJoystickVirtual`, and ForID variants.

**Joystick methods:** `Close`, `Name`, `Path`, `Type`, `ID`, `Connected`, `Vendor`, `Product`, `Serial`, `GUID`, `PlayerIndex`, `SetPlayerIndex`, `ConnectionState`, `PowerInfo`, `Properties`, `GetJoystickAxis`, `GetJoystickHat`, `GetJoystickButton`, `GetJoystickBall`, `AxisInitialState`, `RumbleJoystick`, `RumbleTriggers`, `SetJoystickLED`, `SendJoystickEffect`, virtual joystick setters.

---

### Haptic / Force Feedback

File: `sdl_haptic.go`

**Types:** `Haptic`, `HapticID`, `HapticEffectType`, `HapticDirectionType`, `HapticEffectID`, `HapticDirection`, `HapticConstant`, `HapticPeriodic`, `HapticCondition`, `HapticRamp`, `HapticLeftRight`, `HapticCustom`

**Functions:** `GetHaptics`, `OpenHaptic`, `GetHapticFromID`, `CloseHaptic`, `IsMouseHaptic`, `OpenHapticFromMouse`, `IsJoystickHaptic`, `OpenHapticFromJoystick`, `GetMaxHapticEffects`, `GetMaxHapticEffectsPlaying`, `GetHapticFeatures`, `GetNumHapticAxes`, `HapticEffectSupported`, `CreateHapticEffect`, `UpdateHapticEffect`, `RunHapticEffect`, `StopHapticEffect`, `DestroyHapticEffect`, `SetHapticGain`, `SetHapticAutocenter`, `PauseHaptic`, `ResumeHaptic`, `StopHapticEffects`, `HapticRumbleSupported`, `InitHapticRumble`, `PlayHapticRumble`, `StopHapticRumble`

---

### Sensor

File: `sdl_sensor.go`

**Types:** `Sensor`, `SensorID`, `SensorType`

**Functions:** `GetSensors`, `GetSensorNameForID`, `GetSensorTypeForID`, `OpenSensor`, `GetSensorFromID`, `UpdateSensors`

**Methods:** `Close`, `Name`, `Type`, `ID`, `Properties`, `GetData`

---

### Camera

File: `sdl_camera.go`

**Types:** `Camera`, `CameraID`, `CameraSpec`, `CameraPosition`, `CameraPermissionState`

**Functions:** `GetNumCameraDrivers`, `GetCameraDriver`, `GetCurrentCameraDriver`, `GetCameras`, `GetCameraSupportedFormats`, `GetCameraName`, `GetCameraPosition`, `OpenCamera`

**Methods:** `Close`, `ID`, `Properties`, `GetFormat`, `PermissionState`, `AcquireFrame`, `ReleaseFrame`

---

### Audio

File: `sdl_audio.go`

**Types:** `AudioDeviceID`, `AudioFormat`, `AudioSpec`, `AudioStream`, `AudioStreamCallbackFunc`, `AudioPostmixCallbackFunc`, `AudioStreamDataCompleteCallbackFunc`

**Device functions:** `GetNumAudioDrivers`, `GetAudioDriver`, `GetCurrentAudioDriver`, `GetAudioPlaybackDevices`, `GetAudioRecordingDevices`, `GetAudioDeviceName`, `GetAudioDeviceFormat`, `OpenAudioDevice`, `CloseAudioDevice`, `PauseAudioDevice`, `ResumeAudioDevice`, `AudioDevicePaused`, `GetAudioDeviceGain`, `SetAudioDeviceGain`, `IsAudioDevicePhysical`, `IsAudioDevicePlayback`, `GetAudioDeviceChannelMap`, `SetAudioPostmixCallback`, `OpenAudioDeviceStream`

**Stream creation:** `CreateAudioStream`

**Stream methods:** `Destroy`, `GetFormat`, `SetFormat`, `GetFrequencyRatio`, `SetFrequencyRatio`, `GetGain`, `SetGain`, `PutData`, `PutDataBytes`, `PutDataNoCopy`, `PutPlanarData`, `GetData`, `Available`, `Queued`, `Flush`, `Clear`, `Lock`, `Unlock`, `BindToDevice`, `Unbind`, `Device`, `PauseDevice`, `ResumeDevice`, `DevicePaused`, `Properties`, `SetGetCallback`, `SetPutCallback`, channel map functions.

**Utility:** `LoadWAV`, `LoadWAV_IO`, `MixAudio`, `GetAudioFormatName`, `ConvertAudioSamples`, `GetSilenceValueForFormat`, `BindAudioStreams`, `UnbindAudioStreams`

---

### GPU

File: `sdl_gpu.go`

**Types:** 77 types including `GPUDevice`, `GPUCommandBuffer`, `GPURenderPass`, `GPUComputePass`, `GPUCopyPass`, `GPUShader`, `GPUGraphicsPipeline`, `GPUComputePipeline`, `GPUTexture`, `GPUBuffer`, `GPUSampler`, `GPUTransferBuffer`, `GPUFence`, `GPUTextureFormat`, `GPUVulkanOptions`, and many create-info/state structs.

92 package functions and 7 methods covering device creation, command buffers, render/compute/copy passes, shader/pipeline/resource creation, drawing, dispatching, data transfer, synchronization, format queries, and debug labels.

---

### Vulkan

File: `sdl_vulkan.go`

**Types:** `VkInstance`, `VkPhysicalDevice`, `VkSurfaceKHR`

**Functions:** `Vulkan_LoadLibrary`, `Vulkan_UnloadLibrary`, `Vulkan_GetInstanceExtensions`, `Vulkan_GetVkGetInstanceProcAddr`, `Vulkan_CreateSurface`, `Vulkan_DestroySurface`, `Vulkan_GetPresentationSupport`

---

### Metal

File: `sdl_metal.go`

**Type:** `MetalView`

**Functions:** `Metal_CreateView`, `Metal_DestroyView`, `Metal_GetLayer`

---

### Filesystem

File: `sdl_filesystem.go`

**Types:** `Folder`, `PathType`, `PathInfo`, `GlobFlags`, `EnumerationResult`, `EnumerateDirectoryFunc`

**Functions:** `GetBasePath`, `GetPrefPath`, `GetUserFolder`, `CreateDirectory`, `RemovePath`, `RenamePath`, `CopyFile`, `GetPathInfo`, `GetCurrentDirectory`, `GlobDirectory`, `EnumerateDirectory`

---

### I/O Streams

File: `sdl_iostream.go`

**Types:** `IOStream`, `IOStatus`, `IOWhence`, `IOStreamImpl`

**Creation:** `IOFromFile`, `IOFromMem`, `IOFromConstMem`, `IOFromDynamicMem`, `OpenIO` (custom), `LoadFile`, `SaveFile`

**Methods:** `Close`, `Status`, `Properties`, `SeekIO`, `Tell`, `GetSize`, `Read`, `Write`, `Flush`, `Printf`, `LoadAll`, `SaveAll`, plus 28 binary I/O methods (`ReadU8`, `ReadU16LE`, `WriteU32BE`, etc.)

---

### Storage

File: `sdl_storage.go`

**Type:** `Storage`

**Functions:** `OpenTitleStorage`, `OpenUserStorage`, `OpenFileStorage`, `OpenStorage`

**Methods:** `Close`, `Ready`, `GetFileSize`, `ReadFile`, `WriteFile`, `CreateDirectory`, `RemovePath`, `RenamePath`, `CopyFile`, `GetPathInfo`, `GetSpaceRemaining`, `GlobStorageDirectory`, `EnumerateDirectory`

---

### Process

File: `sdl_process.go`

**Types:** `Process`, `ProcessIO`

**Functions:** `CreateProcess`, `CreateProcessWithProperties`

**Methods:** `Properties`, `GetInput`, `GetOutput`, `Kill`, `Wait`, `ReadAll`, `Destroy`

---

### File Dialogs

File: `sdl_dialog.go`

**Types:** `FileDialogType`, `DialogFileFilter`, `DialogFileCallback`

**Functions:** `ShowOpenFileDialog`, `ShowSaveFileDialog`, `ShowOpenFolderDialog`, `ShowFileDialogWithProperties`

---

### Clipboard

File: `sdl_clipboard.go`

**Functions:** `SetClipboardText`, `GetClipboardText`, `HasClipboardText`, `SetPrimarySelectionText`, `GetPrimarySelectionText`, `HasPrimarySelectionText`, `ClearClipboardData`, `GetClipboardData`, `HasClipboardData`, `GetClipboardMimeTypes`, `SetClipboardData`

---

### Timer

File: `sdl_timer.go`

**Types:** `TimerID`, `TimerCallback`, `TimerNSCallback`

**Functions:** `GetTicks`, `GetTicksNS`, `GetPerformanceCounter`, `GetPerformanceFrequency`, `Delay`, `DelayNS`, `DelayPrecise`, `AddTimer`, `AddTimerNS`, `RemoveTimer`

**Constants:** `MS_PER_SECOND`, `NS_PER_SECOND`, `NS_PER_MS`, `NS_PER_US`, etc.

---

### Hints

File: `sdl_hints.go`

**Types:** `HintPriority`, `HintCallbackFunc`

**Functions:** `SetHint`, `GetHint`, `GetHintBoolean`, `ResetHint`, `ResetHints`, `SetHintWithPriority`, `AddHintCallback`, `RemoveHintCallback`

266 `HINT_*` string constants for all SDL hints.

---

### Log

File: `sdl_log.go`

**Types:** `LogPriority`, `LogCategory`, `LogOutputFunc`

**Functions:** `Log`, `LogMessage`, `LogVerbose`, `LogDebug`, `LogInfo`, `LogWarn`, `LogError`, `LogCritical`, `LogTrace`, `SetLogPriority`, `GetLogPriority`, `SetLogPriorities`, `ResetLogPriorities`, `SetLogPriorityPrefix`, `SetLogOutputFunction`, `GetLogOutputFunction`, `GetDefaultLogOutputFunction`

---

### Properties

File: `sdl_properties.go`

**Types:** `PropertiesID`, `PropertyType`, `EnumeratePropertiesFunc`, `CleanupPropertyFunc`

**Functions:** `CreateProperties`, `DestroyProperties`, `GetGlobalProperties`, `CopyProperties`, `LockProperties`, `UnlockProperties`, `SetStringProperty`, `GetStringProperty`, `SetNumberProperty`, `GetNumberProperty`, `SetFloatProperty`, `GetFloatProperty`, `SetBooleanProperty`, `GetBooleanProperty`, `SetPointerProperty`, `GetPointerProperty`, `SetPointerPropertyWithCleanup`, `HasProperty`, `ClearProperty`, `GetPropertyType`, `EnumerateProperties`

272 `PROP_*` string constants for all SDL property names.

---

### Thread

File: `sdl_thread.go`

**Types:** `Thread`, `ThreadID`, `ThreadPriority`, `ThreadState`, `TLSID`, `ThreadFunction`

**Functions:** `GetCurrentThreadID`, `SetCurrentThreadPriority`, `CreateThread`, `CreateThreadWithProperties`, `GetTLS`, `SetTLS`, `CleanupTLS`

**Methods:** `Name`, `ID`, `State`, `Wait`, `Detach`

---

### Mutex / Sync

File: `sdl_mutex.go`

**Types:** `Mutex`, `RWLock`, `Semaphore`, `Condition`, `InitState`, `InitStatus`

- **Mutex:** `CreateMutex`, `Lock`, `TryLock`, `Unlock`, `Destroy`
- **RWLock:** `CreateRWLock`, `LockForReading`, `LockForWriting`, `TryLockForReading`, `TryLockForWriting`, `Unlock`, `Destroy`
- **Semaphore:** `CreateSemaphore`, `Wait`, `TryWait`, `WaitTimeout`, `Signal`, `Value`, `Destroy`
- **Condition:** `CreateCondition`, `Signal`, `Broadcast`, `Wait`, `WaitTimeout`, `Destroy`
- **InitState:** `ShouldInit`, `ShouldQuit`, `SetInitialized`

---

### Atomic

File: `sdl_atomic.go`

**Types:** `SpinLock`, `AtomicInt`, `AtomicU32`

**Functions:** `TryLockSpinlock`, `LockSpinlock`, `UnlockSpinlock`, `CompareAndSwapAtomicPointer`, `SetAtomicPointer`, `GetAtomicPointer`, `MemoryBarrierReleaseFunction`, `MemoryBarrierAcquireFunction`

**Methods:** `AtomicInt.CompareAndSwap`, `Set`, `Get`, `Add` (same for `AtomicU32`)

---

### HID API

File: `sdl_hidapi.go`

**Types:** `HIDDevice`, `HIDDeviceInfo`, `HIDBusType`

**Functions:** `HID_Init`, `HID_Exit`, `HID_DeviceChangeCount`, `HID_Enumerate`, `HID_Open`, `HID_OpenPath`, `HID_BLEScan`

**Methods:** `Close`, `Properties`, `Write`, `Read`, `ReadTimeout`, `SetNonblocking`, `SendFeatureReport`, `GetFeatureReport`, `GetInputReport`, `GetManufacturerString`, `GetProductString`, `GetSerialNumberString`, `GetIndexedString`, `GetDeviceInfo`, `GetReportDescriptor`

---

### Async I/O

File: `sdl_asyncio.go`

**Types:** `AsyncIO`, `AsyncIOQueue`, `AsyncIOTaskType`, `AsyncIOResult`, `AsyncIOOutcome`

**Functions:** `AsyncIOFromFile`, `CreateAsyncIOQueue`, `LoadFileAsync`

**Methods:** `AsyncIO.GetSize`, `Read`, `Write`, `Close`; `AsyncIOQueue.Destroy`, `GetResult`, `WaitResult`, `Signal`

---

### System Tray

File: `sdl_tray.go`

**Types:** `Tray`, `TrayMenu`, `TrayEntry`, `TrayEntryFlags`, `TrayCallbackFunc`

**Functions:** `CreateTray`, `CreateTrayWithProperties`, `UpdateTrays`

**Tray methods:** `Destroy`, `SetIcon`, `SetTooltip`, `CreateMenu`, `GetMenu`

**TrayMenu methods:** `InsertEntry`, `GetEntries`, `ParentEntry`, `ParentTray`

**TrayEntry methods:** `Remove`, `SetLabel`, `Label`, `SetChecked`, `Checked`, `SetEnabled`, `Enabled`, `CreateSubmenu`, `GetSubmenu`, `Parent`, `Click`, `SetCallback`

---

### Locale

File: `sdl_locale.go`

**Type:** `Locale`

**Function:** `GetPreferredLocales()`

---

### Power

File: `sdl_power.go`

**Function:** `GetPowerInfo()` returns `PowerState`, seconds, percent

---

### Shared Objects

File: `sdl_loadso.go`

**Type:** `SharedObject`

**Functions:** `LoadObject(path)`, `SharedObject.LoadFunction(name)`, `SharedObject.Unload()`

---

### Misc

File: `sdl_misc.go`

**Function:** `OpenURL(url)`

---

### Time

File: `sdl_time.go`

**Types:** `Time`, `DateFormat`, `TimeFormat`, `DateTime`

**Functions:** `GetCurrentTime`, `TimeToDateTime`, `DateTimeToTime`, `GetDaysInMonth`, `GetDayOfYear`, `GetDayOfWeek`, `GetDateTimeLocalePreferences`, `TimeToWindows`, `TimeFromWindows`

---

### GUID

File: `sdl_guid.go`

**Type:** `GUID`

**Functions:** `GUIDFromString(s)`, `GUID.String()`

---

### CPU Info

File: `sdl_cpuinfo.go`

**Functions:** `GetNumLogicalCPUCores`, `GetCPUCacheLineSize`, `HasAltiVec`, `HasMMX`, `HasSSE`, `HasSSE2`, `HasSSE3`, `HasSSE41`, `HasSSE42`, `HasAVX`, `HasAVX2`, `HasAVX512F`, `HasARMSIMD`, `HasNEON`, `HasLSX`, `HasLASX`, `GetSystemRAM`, `GetSIMDAlignment`, `GetSystemPageSize`

**Constant:** `CACHELINE_SIZE`

---

### Version

File: `sdl_version.go`

**Functions:** `GetVersion`, `GetRevision`, `VersionNum`, `VersionMajor`, `VersionMinor`, `VersionMicro`, `VersionAtLeast`

**Constants:** `MAJOR_VERSION`, `MINOR_VERSION`, `MICRO_VERSION`

---

### Message Box

File: `sdl_messagebox.go`

**Types:** `MessageBoxFlags`, `MessageBoxButtonFlags`, `MessageBoxButtonData`, `MessageBoxColor`, `MessageBoxColorScheme`, `MessageBoxData`, `MessageBoxColorType`

**Functions:** `ShowSimpleMessageBox`, `ShowMessageBox`

**Constants:** `MESSAGEBOX_ERROR`, `MESSAGEBOX_WARNING`, `MESSAGEBOX_INFORMATION`, `MESSAGEBOX_COLOR_*`
