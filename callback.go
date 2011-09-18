// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package glfw

//#include "glue.h"
//
// extern void goWindowSizeCB(int, int);
// extern int  goWindowCloseCB(void);
// extern void goWindowRefreshCB(void);
// extern void goMouseButtonCB(int, int);
// extern void goMousePosCB(int, int);
// extern void goMouseWheelCB(int);
// extern void goKeyCB(int, int);
// extern void goCharCB(int, int);
//
// void setWindowSizeCB()    { glfwSetWindowSizeCallback(goWindowSizeCB); }
// void setWindowCloseCB()   { glfwSetWindowCloseCallback(goWindowCloseCB); }
// void setWindowRefreshCB() { glfwSetWindowRefreshCallback(goWindowRefreshCB); }
// void setMouseButtonCB()   { glfwSetMouseButtonCallback(goMouseButtonCB); }
// void setMousePosCB()      { glfwSetMousePosCallback(goMousePosCB); }
// void setMouseWheelCB()    { glfwSetMouseWheelCallback(goMouseWheelCB); }
// void setKeyCB()           { glfwSetKeyCallback(goKeyCB); }
// void setCharCB()          { glfwSetCharCallback(goCharCB); }
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
	C.setWindowSizeCB()
}

// =============================================================================

type WindowCloseHandler func() int

var windowClose WindowCloseHandler

//export goWindowCloseCB
func goWindowCloseCB() C.int {
	if goWindowCloseCB == nil {
		return 0
	}
	return C.int(windowClose())
}

// SetWindowCloseCallback sets the callback for window close events.
// A window has to be opened for this function to have any effect.
func SetWindowCloseCallback(f WindowCloseHandler) {
	windowClose = f
	C.setWindowCloseCB()
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
	C.setWindowRefreshCB()
}

// =============================================================================

type MouseButtonHandler func(button, state int)

var mouseButton MouseButtonHandler

//export goMouseButtonCB
func goMouseButtonCB(button, state C.int) {
	if mouseButton == nil {
		return
	}
	mouseButton(int(button), int(state))
}

// SetMouseButtonCallback sets the callback for mouse button events.
func SetMouseButtonCallback(f MouseButtonHandler) {
	mouseButton = f
	C.setMouseButtonCB()
}

// =============================================================================

type MousePosHandler func(x, y int)

var mousePos MousePosHandler

//export goMousePosCB
func goMousePosCB(x, y C.int) {
	if mousePos == nil {
		return
	}
	mousePos(int(x), int(y))
}

// SetMousePosCallback sets the callback for mouse motion events.
func SetMousePosCallback(f MousePosHandler) {
	mousePos = f
	C.setMousePosCB()
}

// =============================================================================

type MouseWheelHandler func(delta int)

var mouseWheel MouseWheelHandler

//export goMouseWheelCB
func goMouseWheelCB(delta C.int) {
	if mouseWheel == nil {
		return
	}
	mouseWheel(int(delta))
}

// This function sets the callback for mouse wheel events.
func SetMouseWheelCallback(f MouseWheelHandler) {
	mouseWheel = f
	C.setMouseWheelCB()
}

// =============================================================================

type KeyHandler func(key, state int)

var key KeyHandler

//export goKeyCB
func goKeyCB(k, state C.int) {
	if key == nil {
		return
	}
	key(int(k), int(state))
}

// SetKeyCallback sets the callback for keyboard key events. The callback
// function is called every time the state of a single key is changed (from
// released to pressed or vice versa). The reported keys are unaffected by any
// modifiers (such as shift or alt) and each modifier is reported as a separate key.
func SetKeyCallback(f KeyHandler) {
	key = f
	C.setKeyCB()
}

// =============================================================================

type CharHandler func(int, int)

var char CharHandler

//export goCharCB
func goCharCB(x, y C.int) {
	if char == nil {
		return
	}
	char(int(x), int(y))
}

// SetCharCallback sets the callback for keyboard character events. The callback
// function is called every time a key that results in a printable Unicode
// character is pressed or released. Characters are affected by modifiers
// (such as shift or alt).
func SetCharCallback(f CharHandler) {
	char = f
	C.setCharCB()
}
