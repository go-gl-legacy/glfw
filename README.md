## GLFW

This is a set of Go bindings for libGLFW.

GLFW is a small C library that lets you create and manage an OpenGL context and
its associated window, enumerate and change display modes, as well as handle
inputs such as keyboard, mouse, joystick and time.

GLFW provides a thin, multi-platform abstraction layer, primarily for
applications whose sole graphics output is through the OpenGL API. While GLFW is
very useful when developing multi-platform OpenGL applications, single-platform
developers can also benefit from avoiding the drudgery of kludgy platform-
specific APIs.

This wrapper targets GLFW 2.7

We do not implement the thread/mutex api.
Callback functions are working.

### Dependencies

* [libglfw](http://www.glfw.org/download.html)
 
_Important_: libglfw builds/installs itself as a static library by default.
This does not work well with go. Eventhough the building of this package may succeed
without problems, any application you use it in will likely throw up a range
of symbol lookup errors. It is therefor strongly recommended to build libglfw
as a SHARED library instead. See the libglfw README and makefiles on how to
do this.


### Usage

    go get github.com/go-gl/glfw


### License

Copyright 2012 The go-gl Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

GLFW is licensed under the zlib/libpng license.
Its contents can be found in the GLFW_LICENSE file.

