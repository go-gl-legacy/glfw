// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package glfw

//#include "glue.h"
import "C"
import (
	"unsafe"
	"os"
)

type Image struct {
	ptr *C.GLFWimage
}

func NewImage() *Image {
	v := new(Image)
	v.ptr = (*C.GLFWimage)(C.malloc(C.size_t(unsafe.Sizeof(*v.ptr))))
	return v
}

func (this *Image) Release() {
	C.glfwFreeImage(this.ptr)
	this.ptr = nil
}

func (this *Image) Width() int {
	return int(this.ptr.Width)
}

func (this *Image) SetWidth(v int) {
	this.ptr.Width = C.int(v)
}

func (this *Image) Height() int {
	return int(this.ptr.Height)
}

func (this *Image) SetHeight(v int) {
	this.ptr.Width = C.int(v)
}

func (this *Image) Format() int {
	return int(this.ptr.Format)
}

func (this *Image) SetFormat(v int) {
	this.ptr.Format = C.int(v)
}

func (this *Image) BytesPerPixel() int {
	return int(this.ptr.BytesPerPixel)
}

func (this *Image) SetBytesPerPixel(v int) {
	this.ptr.BytesPerPixel = C.int(v)
}

func (this *Image) Data() []byte {
	size := this.ptr.BytesPerPixel * this.ptr.Width * this.ptr.Height
	return (*(*[1<<31 - 1]byte)(unsafe.Pointer(this.ptr.Data)))[:size]
}

func (this *Image) SetData(v []byte) {
	this.ptr.Data = (*_Ctype_unsignedchar)(unsafe.Pointer(&v))
}

func (this *Image) Read(name string, flags int) (err os.Error) {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	if C.glfwReadImage(cn, this.ptr, C.int(flags)) != 1 {
		err = os.NewError("Failed to read image " + name)
	}
	return
}

func (this *Image) ReadMemoryImage(data []byte, flags int) (err os.Error) {
	if C.glfwReadMemoryImage(unsafe.Pointer(&data), C.long(len(data)), this.ptr, C.int(flags)) != 1 {
		err = os.NewError("Failed to read image from memory")
	}
	return
}

func (this *Image) LoadTextureImage2D(flags int) int {
	return int(C.glfwLoadTextureImage2D(this.ptr, C.int(flags)))
}
