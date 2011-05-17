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

type WindowSizeHandler func(width, height int)

var windowSize WindowSizeHandler

//export goWindowSizeCB
func goWindowSizeCB(width, height C.int) {
	windowSize(int(width), int(height))
}

func SetWindowSizeCallback(f WindowSizeHandler) {
	windowSize = f
	C.setWindowSizeCB()
}

type WindowCloseHandler func() int

var windowClose WindowCloseHandler

//export goWindowCloseCB
func goWindowCloseCB() C.int {
	return C.int(windowClose())
}

func SetWindowCloseCallback(f WindowCloseHandler) {
	windowClose = f
	C.setWindowCloseCB()
}

type WindowRefreshHandler func()

var windowRefresh WindowRefreshHandler

//export goWindowRefreshCB
func goWindowRefreshCB() {
	windowRefresh()
}

func SetWindowRefreshCallback(f WindowRefreshHandler) {
	windowRefresh = f
	C.setWindowRefreshCB()
}

type MouseButtonHandler func(button, state int)

var mouseButton MouseButtonHandler

//export goMouseButtonCB
func goMouseButtonCB(button, state C.int) {
	mouseButton(int(button), int(state))
}

func SetMouseButtonCallback(f MouseButtonHandler) {
	mouseButton = f
	C.setMouseButtonCB()
}

type MousePosHandler func(x, y int)

var mousePos MousePosHandler

//export goMousePosCB
func goMousePosCB(x, y C.int) {
	mousePos(int(x), int(y))
}

func SetMousePosCallback(f MousePosHandler) {
	mousePos = f
	C.setMousePosCB()
}

type MouseWheelHandler func(delta int)

var mouseWheel MouseWheelHandler

//export goMouseWheelCB
func goMouseWheelCB(delta C.int) {
	mouseWheel(int(delta))
}

func SetMouseWheelCallback(f MouseWheelHandler) {
	mouseWheel = f
	C.setMouseWheelCB()
}

type KeyHandler func(key, state int)

var key KeyHandler

//export goKeyCB
func goKeyCB(k, state C.int) {
	key(int(k), int(state))
}

func SetKeyCallback(f KeyHandler) {
	key = f
	C.setKeyCB()
}

type CharHandler func(int, int)

var char CharHandler

//export goCharCB
func goCharCB(x, y C.int) {
	char(int(x), int(y))
}

func SetCharCallback(f CharHandler) {
	char = f
	C.setCharCB()
}
