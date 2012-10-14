// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package glfw

//#include "callback.h"
import "C"

// =============================================================================

type WindowSizeHandler func(width, height int)

var windowSize WindowSizeHandler

//export goWindowSizeCB
func goWindowSizeCB(width, height C.int) {
	if windowSize == nil {
		return
	}
	windowSize(int(width), int(height))
}

// SetWindowSizeCallback sets the callback for window size change events.
func SetWindowSizeCallback(f WindowSizeHandler) {
	windowSize = f
	C.glfwSetWindowSizeCB()
}

// =============================================================================

type WindowCloseHandler func() int

var windowClose WindowCloseHandler

//export goWindowCloseCB
func goWindowCloseCB() C.int {
	if windowClose == nil {
		return 0
	}
	return C.int(windowClose())
}

// SetWindowCloseCallback sets the callback for window close events.
// A window has to be opened for this function to have any effect.
func SetWindowCloseCallback(f WindowCloseHandler) {
	windowClose = f
	C.glfwSetWindowCloseCB()
}

// =============================================================================

type WindowRefreshHandler func()

var windowRefresh WindowRefreshHandler

//export goWindowRefreshCB
func goWindowRefreshCB() {
	if windowRefresh == nil {
		return
	}
	windowRefresh()
}

// SetWindowRefreshCallback sets the callback for window refresh events, which
// occurs when any part of the window client area has been damaged, and needs to
// be repainted (for instance, if a part of the window that was previously
// occluded by another window has become visible).
func SetWindowRefreshCallback(f WindowRefreshHandler) {
	windowRefresh = f
	C.glfwSetWindowRefreshCB()
}

// =============================================================================

type MouseButtonHandler func(button, state int)

var mouseButton []MouseButtonHandler

//export goMouseButtonCB
func goMouseButtonCB(button, state C.int) {
	for _, f := range mouseButton {
		f(int(button), int(state))
	}
}

// SetMouseButtonCallback sets the callback for mouse button events.
// There can be more than one handler.
func SetMouseButtonCallback(f MouseButtonHandler) {
	mouseButton = append(mouseButton, f)
	C.glfwSetMouseButtonCB()
}

// =============================================================================

type MousePosHandler func(x, y int)

var mousePos []MousePosHandler

//export goMousePosCB
func goMousePosCB(x, y C.int) {
	for _, f := range mousePos {
		f(int(x), int(y))
	}
}

// SetMousePosCallback sets a callback for mouse motion events.
// There can be more than one handler.
func SetMousePosCallback(f MousePosHandler) {
	mousePos = append(mousePos, f)
	C.glfwSetMousePosCB()
}

// =============================================================================

type MouseWheelHandler func(delta int)

var mouseWheel []MouseWheelHandler

//export goMouseWheelCB
func goMouseWheelCB(delta C.int) {
	for _, f := range mouseWheel {
		f(int(delta))
	}
}

// This function sets the callback for mouse wheel events.
// There can be more than one handler.
func SetMouseWheelCallback(f MouseWheelHandler) {
	mouseWheel = append(mouseWheel, f)
	C.glfwSetMouseWheelCB()
}

// =============================================================================

type KeyHandler func(key, state int)

var key []KeyHandler

//export goKeyCB
func goKeyCB(k, state C.int) {
	for _, f := range key {
		f(int(k), int(state))
	}
}

// SetKeyCallback sets the callback for keyboard key events. The callback
// function is called every time the state of a single key is changed (from
// released to pressed or vice versa). The reported keys are unaffected by any
// modifiers (such as shift or alt) and each modifier is reported as a separate key.
//
// There can be more than one handler.
func SetKeyCallback(f KeyHandler) {
	key = append(key, f)
	C.glfwSetKeyCB()
}

// =============================================================================

type CharHandler func(int, int)

var char []CharHandler

//export goCharCB
func goCharCB(x, y C.int) {
	for _, f := range char {
		f(int(x), int(y))
	}
}

// SetCharCallback sets the callback for keyboard character events. The callback
// function is called every time a key that results in a printable Unicode
// character is pressed or released. Characters are affected by modifiers
// (such as shift or alt).
func SetCharCallback(f CharHandler) {
	char = append(char, f)
	C.glfwSetCharCB()
}
