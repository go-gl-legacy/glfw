// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package glfw

//#include <stdlib.h>
//#ifdef _WIN32
//  #define GLFW_DLL
//#endif
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
