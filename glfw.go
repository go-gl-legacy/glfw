// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package glfw

//#include "glue.h"
import "C"
import (
	"os"
	"unsafe"
)

func Terminate()                       { C.glfwTerminate() }
func OpenWindowHint(target, hint int)  { C.glfwOpenWindowHint(C.int(target), C.int(hint)) }
func CloseWindow()                     { C.glfwCloseWindow() }
func IconifyWindow()                   { C.glfwIconifyWindow() }
func RestoreWindow()                   { C.glfwRestoreWindow() }
func SwapBuffers()                     { C.glfwSwapBuffers() }
func WindowParam(param int) int        { return int(C.glfwGetWindowParam(C.int(param))) }
func SetSwapInterval(interval int)     { C.glfwSwapInterval(C.int(interval)) }
func SetWindowSize(width, height int)  { C.glfwSetWindowSize(C.int(width), C.int(height)) }
func SetWindowPos(x, y int)            { C.glfwSetWindowPos(C.int(x), C.int(y)) }
func SetWindowTitle(title string)      { C.glfwSetWindowTitle(C.CString(title)) }
func PollEvents()                      { C.glfwPollEvents() }
func WaitEvents()                      { C.glfwWaitEvents() }
func Key(key int) int                  { return int(C.glfwGetKey(C.int(key))) }
func MouseButton(btn int) int          { return int(C.glfwGetMouseButton(C.int(btn))) }
func SetMousePos(x, y int)             { C.glfwSetMousePos(C.int(x), C.int(y)) }
func MouseWheel() int                  { return int(C.glfwGetMouseWheel()) }
func SetMouseWheel(pos int)            { C.glfwSetMouseWheel(C.int(pos)) }
func JoystickParam(joy, param int) int { return int(C.glfwGetJoystickParam(C.int(joy), C.int(param))) }
func Time() float64                    { return float64(C.glfwGetTime()) }
func SetTime(t float64)                { C.glfwSetTime(C.double(t)) }
func Sleep(t float64)                  { C.glfwSleep(C.double(t)) }
func NumberOfProcessors() int          { return int(C.glfwGetNumberOfProcessors()) }
func Enable(token int)                 { C.glfwEnable(C.int(token)) }
func Disable(token int)                { C.glfwDisable(C.int(token)) }

func Init() (err os.Error) {
	if C.glfwInit() != 1 {
		err = os.NewError("Failed to initialize GLFW")
	}
	return
}

func OpenWindow(width, height, r, g, b, a, depth, stencil, mode int) (err os.Error) {
	if C.glfwOpenWindow(
		C.int(width), C.int(height),
		C.int(r), C.int(g), C.int(b), C.int(a),
		C.int(depth), C.int(stencil), C.int(mode),
	) != 1 {
		err = os.NewError("Failed to open window")
	}
	return
}

func WindowSize() (int, int) {
	var w, h C.int
	C.glfwGetWindowSize(&w, &h)
	return int(w), int(h)
}

func VideoModes(max int) []*VidMode {
	var vm C.GLFWvidmode

	size := unsafe.Sizeof(vm)
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

func DesktopMode() *VidMode {
	var vm C.GLFWvidmode
	C.glfwGetDesktopMode(&vm)
	return vidModeFromPtr(&vm)
}

func MousePos() (int, int) {
	var cx, cy C.int
	C.glfwGetMousePos(&cx, &cy)
	return int(cx), int(cy)
}

func JoystickPos(joy int, numaxes int) float32 {
	var pos C.float
	if int(C.glfwGetJoystickPos(C.int(joy), &pos, C.int(numaxes))) != 1 {
		return 0
	}
	return float32(pos)
}

func JoystickButtons(joy, numbuttons int) []byte {
	ptr := (*_Ctype_unsignedchar)(C.malloc(C.size_t(C.int(numbuttons))))
	defer C.free(unsafe.Pointer(ptr))

	var count C.int
	if count = C.glfwGetJoystickButtons(C.int(joy), ptr, C.int(numbuttons)); count == 0 {
		return nil
	}

	b := make([]byte, count)
	copy(b, (*(*[1<<31 - 1]byte)(unsafe.Pointer(ptr)))[:count])
	return b
}

func ExtensionSupported(name string) bool {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	return C.glfwExtensionSupported(cs) != 1
}

func ProcAddress(name string) uintptr {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))
	return uintptr(C.glfwGetProcAddress(cs))
}

func GLVersion() (int, int, int) {
	var major, minor, rev C.int
	C.glfwGetGLVersion(&major, &minor, &rev)
	return int(major), int(minor), int(rev)
}

func LoadTexture2D(name string, flags int) int {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	return int(C.glfwLoadTexture2D(cn, C.int(flags)))
}

func LoadMemoryTexture2D(data []byte, flags int) int {
	return int(C.glfwLoadMemoryTexture2D(unsafe.Pointer(&data), C.long(len(data)), C.int(flags)))
}
