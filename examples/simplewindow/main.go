// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

// This example shows how to set up a minimal GLFW application.
package main

import (
	"fmt"
	"os"
	"github.com/jteeuwen/glfw"
)

func main() {
	var err os.Error
	if err = glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	// Ensure glfw is cleanly terminated on exit.
	defer glfw.Terminate()

	if err = glfw.OpenWindow(300, 300, 0, 0, 0, 0, 0, 0, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	// Ensure window is cleanly closed on exit.
	defer glfw.CloseWindow()

	// Enable vertical sync on cards that support it.
	glfw.SetSwapInterval(1)

	// Set window title
	glfw.SetWindowTitle("Simple GLFW window")

	// Hook some events to demonstrate use of callbacks.
	glfw.SetWindowSizeCallback(onResize)
	glfw.SetWindowCloseCallback(onClose)
	glfw.SetMouseButtonCallback(onMouseBtn)
	glfw.SetMouseWheelCallback(onMouseWheel)
	glfw.SetKeyCallback(onKey)
	glfw.SetCharCallback(onChar)

	running := true
	for running {
		// OpenGL rendering goes here.

		// Swap front and back rendering buffers. This also implicitly calls
		// glfw.PollEvents(), so we have valid key/mouse/joystick states after
		// this. This behavior can be disabled by calling glfw.Disable with the
		// argument glfw.AutoPollEvents. You must be sure to manually call
		// PollEvents() or WaitEvents() in this case.
		glfw.SwapBuffers()

		// Break out of loop when Escape key is pressed, or window is closed.
		running = glfw.Key(glfw.KeyEsc) == 0 && glfw.WindowParam(glfw.Opened) == 1
	}
}

func onResize(w, h int) {
	println("resized:", w, h)
}

func onClose() int {
	println("closed")
	return 1 // return 0 to keep window open.
}

func onMouseBtn(button, state int) {
	println("mouse button:", button, state)
}

func onMouseWheel(delta int) {
	println("mouse wheel:", delta)
}

func onKey(key, state int) {
	println("key:", key, state)
}

func onChar(key, state int) {
	println("char:", key, state)
}
