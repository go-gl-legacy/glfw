// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package glfw

//#include <stdlib.h>
//#include <GL/glfw.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type VidMode struct {
	W int // width
	H int // height
	R int // red bits
	G int // green bits
	B int // blue bits
}

func vidModeFromPtr(ptr *C.GLFWvidmode) *VidMode {
	v := new(VidMode)
	v.W = int(ptr.Width)
	v.H = int(ptr.Height)
	v.R = int(ptr.RedBits)
	v.G = int(ptr.GreenBits)
	v.B = int(ptr.BlueBits)
	return v
}

func (this *VidMode) String() string {
	return fmt.Sprintf("%dx%d @ %d bpp", this.W, this.H, this.R+this.G+this.B)
}

func (this *VidMode) toPtr() (ptr *C.GLFWvidmode) {
	ptr = (*C.GLFWvidmode)(C.malloc(C.size_t(unsafe.Sizeof(ptr))))
	ptr.Width = C.int(this.W)
	ptr.Height = C.int(this.H)
	ptr.RedBits = C.int(this.R)
	ptr.GreenBits = C.int(this.G)
	ptr.BlueBits = C.int(this.B)
	return
}
