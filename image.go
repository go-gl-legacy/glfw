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
	"errors"
	"unsafe"
)

// LoadTexture2D reads an image from the file specified by the parameter name
// and uploads the image to OpenGL texture memory (using the glTexImage2D function).
//
// Note: If the glfw.BuildMipmapsBit flag is set, all mipmap levels for the
// loaded texture are generated and uploaded to texture memory.
//
// Note: The read image is always rescaled to the nearest larger power of
// two resolution using bilinear interpolation.
//
// Note: Unless the glfw.OriginUlBit flag is set, the first pixel in img.Data
// is the lower left corner of the image. Otherwise the first pixel is the upper
// left corner.
//
// Note: For single component images (i.e. gray scale), Format is set to
// GL_ALPHA if the glfw.AlphaMapBit flag is set, otherwise Format is set to GL_LUMINANCE.
//
// Note: ReadImage supports the Truevision Targa version 1 file format
// (.TGA). Supported pixel formats are: 8-bit gray scale, 8-bit paletted
// (24/32-bit color), 24-bit true color and 32-bit true color + alpha.
func LoadTexture2D(name string, flags int) bool {
	cn := C.CString(name)
	defer C.free(unsafe.Pointer(cn))
	return int(C.glfwLoadTexture2D(cn, C.int(flags))) == 1
}

// LoadMemoryTexture2D reads an image from the memory buffer specified by the
// parameter data and uploads the image to OpenG texture memory (using the
// glTexImage2D function).
//
// Note: If the glfw.BuildMipmapsBit flag is set, all mipmap levels for the
// loaded texture are generated and uploaded to texture memory.
//
// Note: The read image is always rescaled to the nearest larger power of
// two resolution using bilinear interpolation.
//
// Note: Unless the glfw.OriginUlBit flag is set, the first pixel in img.Data
// is the lower left corner of the image. Otherwise the first pixel is the upper
// left corner.
//
// Note: For single component images (i.e. gray scale), Format is set to
// GL_ALPHA if the glfw.AlphaMapBit flag is set, otherwise Format is set to GL_LUMINANCE.
//
// Note: ReadImage supports the Truevision Targa version 1 file format
// (.TGA). Supported pixel formats are: 8-bit gray scale, 8-bit paletted
// (24/32-bit color), 24-bit true color and 32-bit true color + alpha.
func LoadMemoryTexture2D(data []byte, flags int) bool {
	if len(data) == 0 {
		return false
	}

	return int(C.glfwLoadMemoryTexture2D(
		unsafe.Pointer(&data[0]), C.long(len(data)), C.int(flags))) == 1
}

// The Image type holds data for loaded images.
type Image struct {
	img C.GLFWimage
}

// ReadImage reads an image from the file specified by the parameter name and
// returns the image or an error.
//
// Note: By default the read image is rescaled to the nearest larger power of
// two resolution using bilinear interpolation, if necessary. Which is useful
// if the image is to be used as an OpenGL texture. This behavior can be
// disabled by setting the glfw.NoRescaleBit flag.
//
// Note: Unless the glfw.OriginUlBit flag is set, the first pixel in img.Data
// is the lower left corner of the image. Otherwise the first pixel is the upper
// left corner.
//
// Note: For single component images (i.e. gray scale), Format is set to
// GL_ALPHA if the glfw.AlphaMapBit flag is set, otherwise Format is set to GL_LUMINANCE.
//
// Note: ReadImage supports the Truevision Targa version 1 file format
// (.TGA). Supported pixel formats are: 8-bit gray scale, 8-bit paletted
// (24/32-bit color), 24-bit true color and 32-bit true color + alpha.
func ReadImage(name string, flags int) (i *Image, err error) {
	i = new(Image)

	str := C.CString(name)
	defer C.free(unsafe.Pointer(str))

	if C.glfwReadImage(str, &i.img, C.int(flags)) != 1 {
		return nil, errors.New("Failed to read image " + name)
	}

	return
}

// ReadMemoryImage reads an image file from the memory buffer specified by the
// parameter data and returns the image information or an error.
//
// Note: ReadMemoryImage supports the Truevision Targa version 1 file format
// (.TGA). Supported pixel formats are: 8-bit gray scale, 8-bit paletted
// (24/32-bit color), 24-bit true color and 32-bit true color + alpha.
func ReadMemoryImage(data []byte, flags int) (i *Image, err error) {
	if len(data) == 0 {
		return nil, errors.New("No image data was supplied")
	}

	i = new(Image)

	if C.glfwReadMemoryImage(unsafe.Pointer(&data[0]), C.long(len(data)), &i.img, C.int(flags)) != 1 {
		return nil, errors.New("Failed to read image from memory")
	}

	return
}

// LoadTextureImage2D uploads the image specified by the parameter img to 
// OpenGL texture memory (using the glTexImage2D function).
//
// Note: If the glfw.BuildMipmapsBit flag is set, all mipmap levels for the
// loaded texture are generated and uploaded to texture memory.
//
// Note: The read image is always rescaled to the nearest larger power of
// two resolution using bilinear interpolation.
//
// Note: Unless the glfw.OriginUlBit flag is set, the first pixel in img.Data
// is the lower left corner of the image. Otherwise the first pixel is the upper
// left corner.
//
// Note: For single component images (i.e. gray scale), Format is set to
// GL_ALPHA if the glfw.AlphaMapBit flag is set, otherwise Format is set to GL_LUMINANCE.
//
// Note: ReadImage supports the Truevision Targa version 1 file format
// (.TGA). Supported pixel formats are: 8-bit gray scale, 8-bit paletted
// (24/32-bit color), 24-bit true color and 32-bit true color + alpha.
func (i *Image) LoadTextureImage2D(flags int) bool {
	return int(C.glfwLoadTextureImage2D(&i.img, C.int(flags))) == 1
}

// Release frees any memory occupied by a loaded image, and clears all the
// fields of the GLFWimage struct. Any image that has been loaded by the
// glfw.ReadImage function should be deallocated using this function once the
// image is no longer needed.
func (i *Image) Release() { C.glfwFreeImage(&i.img) }

func (i *Image) Width() int { return int(i.img.Width) }

func (i *Image) SetWidth(v int) { i.img.Width = C.int(v) }

func (i *Image) Height() int { return int(i.img.Height) }

func (i *Image) SetHeight(v int) { i.img.Width = C.int(v) }

// Format specifies an OpenGL pixel format, which can be GL_LUMINANCE or
// GL_ALPHA (for gray scale images), GL_RGB or GL_RGBA.
func (i *Image) Format() int {
	return int(i.img.Format)
}

func (i *Image) SetFormat(v int) { i.img.Format = C.int(v) }

// BytesPerPixel specifies the number of bytes per pixel.
func (i *Image) BytesPerPixel() int { return int(i.img.BytesPerPixel) }

func (i *Image) SetBytesPerPixel(v int) { i.img.BytesPerPixel = C.int(v) }

// Data returns the actual pixel data.
//
// FIXME(jimt): This assumes images <= 2gb in size.
func (i *Image) Data() []byte {
	size := i.img.BytesPerPixel * i.img.Width * i.img.Height
	return (*(*[1<<31 - 1]byte)(unsafe.Pointer(i.img.Data)))[:size]
}

// SetData sets the actual pixel data.
//
// FIXME(jimt): This assumes images <= 2gb in size.
func (i *Image) SetData(v []byte) {
	i.img.Data = (*_Ctype_unsignedchar)(unsafe.Pointer(&v))
}
