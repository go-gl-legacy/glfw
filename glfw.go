// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package glfw

//#cgo   linux  CFLAGS: -I/usr/local/include -pthread
//#cgo   linux LDFLAGS: -L/usr/local/lib -pthread -lX11 -lXrandr -lm -lGL -lrt -lglfw
//#cgo  darwin  CFLAGS: -I/usr/local/include
//#cgo  darwin LDFLAGS: -L/usr/local/lib -framework Cocoa -framework OpenGL -framework IOKit -lglfw
//#cgo windows LDFLAGS: -lglu32 -lopengl32 -lglfwdll
//#include <stdlib.h>
//#ifdef _WIN32
//  #define GLFW_DLL
//#endif
//#include <GL/glfw.h>
import "C"
import (
	"errors"
	"unsafe"
)

// Init initializes GLFW. No other GLFW functions may be called before this
// function has succeeded.
func Init() (err error) {
	if C.glfwInit() != 1 {
		err = errors.New("Failed to initialize GLFW")
	}
	return
}

// Terminate closes the window, if open, and kills any running threads. This
// function should be called before a program exits
func Terminate() { C.glfwTerminate() }

// OpenWindowHint sets additional properties for a window that is to be opened.
// For a hint to be registered, the function must be called /before/ calling
// OpenWindow. When the OpenWindow function is called, any hints that
// were registered with the OpenWindowHint function are used for setting the
// corresponding window properties, and then all hints are reset to their default values.
//
//     RefreshRate:    Vertical monitor refresh rate in Hz (only used
//                     for fullscreen windows). Zero means system default.
//     AccumRedBits:   Number of bits for the red channel of the accumulation buffer.
//     AccumGreenBits: Number of bits for the green channel of the accumulation buffer.
//     AccumBlueBits:  Number of bits for the blue channel of the accumulation buffer.
//     AccumAlphaBits: Number of bits for the alpha channel of the accumulation buffer.
//     AuxBuffers:     Number of auxiliary buffers.
//     Stereo:         Specify if stereo rendering should be supported (can be 
//                     GL_TRUE or GL_FALSE).
//     WindowNoResize: Specify whether the window can be resized by the user
//                     (not used for fullscreen windows)
//     FsaaSamples:    Number of samples to use for the multisampling buffer.
//                     Zero disables multisampling.
//     
//     OpenGLVersionMajor:  Major number of the desired minimum OpenGL version.
//     OpenGLVersionMinor:  Minor number of the desired minimum OpenGL version.
//     OpenGLForwardCompat: Specify whether the OpenGL context should be
//                          forward-compatible (i.e. disallow legacy
//                          functionality). This should only be used when
//                          requesting OpenGL version 3.0 or above.
//     OpenGLDebugContext:  Specify whether a debug context should be created.
//     OpenGLProfile:       The OpenGL profile the context should implement, or
//                          zero to let the system choose. Available profiles are
//                          OpenGLCoreProfile and OpenGLCompatProfile.
func OpenWindowHint(target, hint int) { C.glfwOpenWindowHint(C.int(target), C.int(hint)) }

// OpenWindow opens a window that best matches the parameters given to the
// function. How well the resulting window matches the desired window depends
// mostly on the available hardware and OpenGL drivers. In general, selecting
// a fullscreen mode has better chances of generating a close match of
// buffers and channel sizes than does a normal desktop window, since GLFW can
// freely select from all the available video modes. A desktop window is
// normally restricted to the video mode of the desktop.
//
// Note: For additional control of window properties, see glfw.OpenWindowHint.
// 
//     width:   The width of the window. If width is zero, it will be calculated
//              as width = 4/3 height, if height is not zero. If both width and
//              height are zero, width will be set to 640.
//     
//     height:  The height of the window. If height is zero, it will be calculated
//              as height = 4/3 width, if width is not zero. If both width and
//              height are zero, height will be set to 480.
//     
//     r, g, b: The number of bits to use for each color component of the color
//              buffer (0 means default color depth). For instance, setting
//              r=5, g=6 and b=5 will create a 16-bit color buffer, if possible.
//              
//     a:       The number of bits to use for the alpha channel of the color
//              buffer (0 means no alpha channel).
//     
//     depth:   The number of bits to use for the depth buffer (0 means no depth buffer).
//     
//     stencil: The number of bits to use for the stencil buffer (0 means no
//              stencil buffer).
//      
//     mode:    Selects which type of OpenGL window to use. Mode must be either
//              glfw.Windowed, which will generate a normal desktop window, or
//              glfw.Fullscreen, which will generate a window that covers the
//              entire screen. When glfw.Fullscreen is selected, the video mode
//              will be changed to the resolution that closest matches the width
//              and height parameters.
func OpenWindow(width, height, r, g, b, a, depth, stencil, mode int) (err error) {
	if C.glfwOpenWindow(
		C.int(width), C.int(height),
		C.int(r), C.int(g), C.int(b), C.int(a),
		C.int(depth), C.int(stencil), C.int(mode),
	) != 1 {
		err = errors.New("Failed to open window")
	}
	return
}

// CloseWindow closes an opened window and destroys the associated OpenGL context.
func CloseWindow() { C.glfwCloseWindow() }

// IconifyWindow Iconifies (minimizes) a window. If the window is in fullscreen
// mode, then the desktop video mode will be restored.
func IconifyWindow() { C.glfwIconifyWindow() }

// RestoreWindow restores an iconified window. If the window that is restored is
// in fullscreen mode, then the fullscreen video mode will be restored.
func RestoreWindow() { C.glfwRestoreWindow() }

// WindowParam is used for acquiring various properties of an opened window.
//
//     Opened:         GL_TRUE if window is opened, else GL_FALSE.
//     Active:         GL_TRUE if window has focus, else GL_FALSE.
//     Iconified:      GL_TRUE if window is iconified, else GL_FALSE.
//     Accelerated:    GL_TRUE if window is hardware accelerated, else GL_FALSE.
//     RedBits:        Number of bits for the red color component.
//     GreenBits:      Number of bits for the green color component.
//     BlueBits:       Number of bits for the blue color component.
//     AlphaBits:      Number of bits for the alpha buffer.
//     DepthBits:      Number of bits for the depth buffer.
//     StencilBits:    Number of bits for the stencil buffer.
//     RefreshRate:    Vertical monitor refresh rate in Hz. Zero indicates an
//                     unknown or a default refresh rate.
//     AccumRedBits:   Number of bits for the red channel of the accumulation buffer.
//     AccumGreenBits: Number of bits for the green channel of the accumulation buffer.
//     AccumBlueBits:  Number of bits for the blue channel of the accumulation buffer.
//     AccumAlphaBits: Number of bits for the alpha channel of the accumulation buffer.
//     AuxBuffers:     Number of auxiliary buffers.
//     Stereo:         GL_TRUE if stereo rendering is supported, else GL_FALSE.
//     WindowNoResize: GL_TRUE if the window cannot be resized by the user, else GL_FALSE.
//     FsaaSamples:    Number of multisampling buffer samples. Zero indicated
//                     multisampling is disabled.
//     
//     OpenGLVersionMajor:  Major number of the actual version of the context.
//     OpenGLVersionMinor:  Minor number of the actual version of the context.
//     OpenGLForwardCompat: GL_TRUE if the context is forward-compatible, else GL_FALSE.
//     OpenGLDebugContext:  GL_TRUE if the context is a debug context.
//     OpenGLProfile:       The profile implemented by the context, or zero.
//
// Note: 'Accelerated' is only supported under Windows. Other systems will always
// return GL_TRUE. Under Windows, Accelerated means that the OpenGL renderer is
// a 3rd party renderer, rather than the fallback Microsoft software OpenGL
// renderer. In other words, it is not a real guarantee that the OpenGL renderer
// is actually hardware accelerated.
func WindowParam(param int) int { return int(C.glfwGetWindowParam(C.int(param))) }

// WindowSize is used for determining the size of an opened window. The returned
// values are dimensions of the client area of the window (i.e. excluding any
// window borders and decorations).
//
// Note: Even if the size of a fullscreen window does not change once the window
// has been opened, it does not necessarily have to be the same as the size that
// was requested using glfw.OpenWindow. Therefor it is wise to use this function
// to determine the true size of the window once it has been opened.
func WindowSize() (int, int) {
	var w, h C.int
	C.glfwGetWindowSize(&w, &h)
	return int(w), int(h)
}

// SetWindowSize changes the size of an opened window. The width and height
// parameters denote the size of the client area of the window (i.e. excluding
// any window borders and decorations). If the window is in fullscreen mode, the
// video mode will be changed to a resolution that closest matches the width and
// height parameters (the number of color bits will not be changed).
//
// Note: This function has no effect if the window is iconified.
//
// Note: The OpenGL context is guaranteed to be preserved after calling
// SetWindowSize, even if the video mode is changed.
func SetWindowSize(width, height int) { C.glfwSetWindowSize(C.int(width), C.int(height)) }

// SetWindowPos changes the position of an opened window. It does not have any
// effect on a fullscreen window.
//
// Note: The behaviour of this function on multi-monitor systems is ill-defined.
func SetWindowPos(x, y int) { C.glfwSetWindowPos(C.int(x), C.int(y)) }

// SetWindowTitle changes the title of the opened window.
func SetWindowTitle(title string) { C.glfwSetWindowTitle(C.CString(title)) }

// SetSwapInterval selects the minimum number of monitor vertical retraces that
// should occur between two buffer swaps. If the selected swap interval is one,
// the rate of buffer swaps will never be higher than the vertical refresh rate
// of the monitor. If the selected swap interval is zero, the rate of buffer
// swaps is only limited by the speed of the software and the hardware.
func SetSwapInterval(interval int) { C.glfwSwapInterval(C.int(interval)) }

// SwapBuffers swaps the back and front color buffers of the window. If
// glfw.AutoPollEvents is enabled (which is the default), glfw.PollEvents() is
// called after swapping the front and back buffers.
func SwapBuffers() { C.glfwSwapBuffers() }

// PollEvents is used for polling for events, such as user input and window
// resize events. Upon calling this function, all window states, keyboard states
// and mouse states are updated. If any related callback functions are
// registered, these are called during the call to glfw.PollEvents.
//
// Note: glfw.PollEvents is called implicitly from glfw.SwapBuffers if 
// glfw.AutoPollEvents is enabled (as it is by default). Thus, if glfw.SwapBuffers
// is called frequently, which is normally the case, there is no need to call
// glfw.PollEvents.
func PollEvents() { C.glfwPollEvents() }

// WaitEvents is used for waiting for events, such as user input and window
// resize events. Upon calling this function, the calling thread will be put to
// sleep until any event appears in the event queue. When events are available,
// they will be processed just as they are processed by glfw.PollEvents.
//
// Note: It is guaranteed that glfw.WaitEvents will wake up on any event that
// can be processed by glfw.PollEvents. However, GLFW receives many events that
// are only processed internally and the function may behave differently on
// different systems. Do not make any assumptions about when or why
// glfw.WaitEvents will return.
func WaitEvents() { C.glfwWaitEvents() }

// Key returns glfw.KeyPress if the given key is held down, or glfw.KeyRelease
// otherwise.
// 
// Note: Not all key codes are supported on all systems. Also, while some keys
// are available on some keyboard layouts, they may not be available on other
// keyboard layouts.
//
// Note: For systems that do not distinguish between left and right versions of
// modifier keys (shift, alt and control), the left version is used (e.g. glfw.KeyLshift).
func Key(key int) int { return int(C.glfwGetKey(C.int(key))) }

// MouseButton returns glfw.KeyPress if the given mouse button is held down.
// glfw.KeyRelease otherwise.
func MouseButton(btn int) int { return int(C.glfwGetMouseButton(C.int(btn))) }

// MousePos returns the current mouse position. If the cursor is not hidden,
// the mouse position is the cursor position, relative to the upper left corner
// of the window and with the Y-axis down. If the cursor is hidden, the mouse
// position is a virtual absolute position, not limited to any boundaries except
// to those implied by the maximum number that can be represented by a signed integer.
func MousePos() (int, int) {
	var cx, cy C.int
	C.glfwGetMousePos(&cx, &cy)
	return int(cx), int(cy)
}

// SetMousePos changes the position of the mouse. If the cursor is visible
// (not disabled), the cursor will be moved to the specified position, relative
// to the upper left corner of the window client area and with the Y-axis down.
// If the cursor is hidden (disabled), only the mouse position that is reported
// by GLFW is changed.
func SetMousePos(x, y int) { C.glfwSetMousePos(C.int(x), C.int(y)) }

// MouseWheel returns the current mouse wheel position. The mouse wheel can be
// thought of as a third mouse axis, which is available as a separate wheel or
// up/down stick on some mice.
func MouseWheel() int { return int(C.glfwGetMouseWheel()) }

// This function changes the position of the mouse wheel.
func SetMouseWheel(pos int) { C.glfwSetMouseWheel(C.int(pos)) }

// JoystickParam is used for acquiring various properties of a joystick.
func JoystickParam(joy, param int) int { return int(C.glfwGetJoystickParam(C.int(joy), C.int(param))) }

// JoystickPos queries the current position of one or more axes of a joystick.
// The positional values are returned in an array, where the first element
// represents the first axis of the joystick (normally the X axis). Each
// position is in the range -1.0 to 1.0. Where applicable, the positive
// direction of an axis is right, forward or up, and the negative direction is
// left, back or down.
//
// Note: If len(axes) exceeds the number of axes supported by the joystick, or if
// the joystick is not available, the unused elements in the pos array will be
// set to 0.0 (zero).
// 
// Note: The function returns the number of actually returned axes. This is
// the minimum of len(axes) and the number of axes supported by the joystick.
// If the joystick is not supported or connected, the function will return 0 (zero).
func JoystickPos(joy int, axes []float32) int {
	if len(axes) == 0 {
		return 0
	}

	return int(C.glfwGetJoystickPos(C.int(joy), (*C.float)(unsafe.Pointer(&axes[0])), C.int(len(axes))))
}

// JoystickButtons queries the current state of one or more buttons of a joystick.
// The button states are returned in an array, where the first element
// represents the first button of the joystick. Each state can be either
// glfw.KeyPress or glfw.KeyRelease.
//
// Note: If len(buttons) exceeds the number of buttons supported by the joystick,
// or if the joystick is not available, the unused elements in the buttons array
// will be set to glfw.KeyRelease.
//
// Note: The function returns the number of actually returned buttons. This is
// the minimum of len(buttons) and the number of buttons supported by the
// joystick. If the joystick is not supported or connected, the function will
// return 0 (zero).
func JoystickButtons(joy int, buttons []byte) int {
	if len(buttons) == 0 {
		return 0
	}

	return int(C.glfwGetJoystickButtons(C.int(joy),
		(*C.uchar)(unsafe.Pointer(&buttons[0])), C.int(len(buttons))))
}

// Time returns the state of a high precision timer. Unless the timer has been
// set by the glfw.SetTime function, the time is measured as the number of
// seconds that have passed since glfwInit was called.
//
// Note: The resolution of the timer depends on which system the program is running on.
func Time() float64 { return float64(C.glfwGetTime()) }

// SetTime sets the current time of the high precision timer to the specified
// time. Subsequent calls to glfw.Time will be relative to this time. The time
// is given in seconds.
func SetTime(t float64) { C.glfwSetTime(C.double(t)) }

// Sleep puts the calling thread to sleep for the requested period of time.
// Only the calling thread is put to sleep. Other threads within the same
// process can still execute.
//
// Note: There is usually a system dependent minimum time for which it is
// possible to sleep. This time is generally in the range 1 ms to 20 ms,
// depending on thread sheduling time slot intervals etc. Using a shorter time
// as a parameter to glfwSleep can give one of two results: either the thread
// will sleep for the minimum possible sleep time, or the thread will not sleep
// at all (glfw.Sleep returns immediately). The latter should only happen when
// very short sleep times are specified, if at all.
func Sleep(t float64) { C.glfwSleep(C.double(t)) }

// Enable is used to enable a certain feature.
//
// Supported tokens are: MouseCursor, StickyKeys, StickyMouseButtons,
// SystemKeys, KeyRepeat, AutoPollEvents
func Enable(token int) { C.glfwEnable(C.int(token)) }

// Disable is used to diusable a certain feature.
//
// Supported tokens are: MouseCursor, StickyKeys, StickyMouseButtons,
// SystemKeys, KeyRepeat, AutoPollEvents
func Disable(token int) { C.glfwDisable(C.int(token)) }

// VideoModes returns a list of supported video modes.
//
// The returned list is sorted, first by color depth (RedBits + GreenBits + BlueBits),
// and then by resolution (W idth × Height), with the lowest resolution, fewest
// bits per pixel mode first.
func VideoModes(max int) []*VidMode {
	var vm C.GLFWvidmode

	size := int(unsafe.Sizeof(vm))
	ptr := (*C.GLFWvidmode)(C.malloc(C.size_t(size * max)))
	defer C.free(unsafe.Pointer(ptr))
	count := C.glfwGetVideoModes(ptr, C.int(max))

	if count == 0 {
		return nil
	}

	list := make([]*VidMode, count)
	for i := range list {
		p := (*C.GLFWvidmode)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(i*size)))
		list[i] = vidModeFromPtr(p)
	}

	return list
}

// DesktopMode returns the desktop video mode.
//
// Note: The color depth of the desktop display is always reported as the number
// of bits for each individual color component (red, green and blue), even if
// the desktop is not using an RGB or RGBA color format. For instance, an
// indexed 256 color display may report RedBits = 3, GreenBits = 3 and
// BlueBits = 2, which adds up to 8 bits in total.
//
// Note: The desktop video mode is the video mode used by the desktop at the
// time the GLFW window was opened, not the current video mode (which may differ
// from the desktop video mode if the GLFW window is a fullscreen window).
func DesktopMode() *VidMode {
	var vm C.GLFWvidmode
	C.glfwGetDesktopMode(&vm)
	return vidModeFromPtr(&vm)
}

// ExtensionSupported does a string search in the list of supported OpenGL 
// extensions to find if the specified extension is listed.
//
// Note: An OpenGL context must be created before this function can be called.
//
// Note: In addition to checking for OpenGL extensions, GLFW also checks for
// extensions in the operating system “glue API”, such as WGL extensions under
// Microsoft Windows and GLX extensions under the X Window System.
func ExtensionSupported(name string) bool {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	return C.glfwExtensionSupported(cs) == 1
}

// ProcAddress acquires the pointer to an OpenGL extension function. Some
// (but not all) OpenGL extensions define new API functions, which are usually
// not available through normal linking. It is therefore necessary to get access
// to those API functions at runtime.
//
// Note: An OpenGL context must be created before this function can be called.
//
// Note: Some systems do not support dynamic function pointer retrieval, in
// which case this function will always return 0 (NULL).
func ProcAddress(name string) uintptr {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	return uintptr(C.glfwGetProcAddress(cs))
}

// GLVersion returns the OpenGL implementation version. This is a convenient
// function that parses the version number information at the beginning of the
// string returned by calling glGetString( GL_VERSION ). The OpenGL version
// information can be used to determine what functionality is supported by the
// used OpenGL implementation.
//
// Note: An OpenGL context must be created before this function can be called.
func GLVersion() (int, int, int) {
	var major, minor, rev C.int
	C.glfwGetGLVersion(&major, &minor, &rev)
	return int(major), int(minor), int(rev)
}

// NumberOfProcessors determines the number of active processors in the system.
//
// Note: Systems with several logical processors per physical processor, also
// known as SMT (Symmetric Multi-Threading) processors, will report the number
// of logical processors.
func NumberOfProcessors() int { return int(C.glfwGetNumberOfProcessors()) }
