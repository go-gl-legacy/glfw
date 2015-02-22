// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package glfw

import "runtime"

// Fix issue #4 (https://github.com/jteeuwen/glfw/issues/4) on OSX.
//
// This used to be fixed in the Go runtime, but the patch was reverted.
// See http://code.google.com/p/go/issues/detail?id=3125 for info.
func init() {
	runtime.LockOSThread()
}
