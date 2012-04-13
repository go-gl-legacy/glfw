// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package glfw

import "runtime"

// Fix issue #4 (https://github.com/jteeuwen/glfw/issues/4) on OSX.
//
// This used to be fixed in the Go runtime, but the patch was reverted.
// See http://code.google.com/p/go/issues/detail?id=3125 for info.
func init() {
	runtime.LockOSThread()
}

