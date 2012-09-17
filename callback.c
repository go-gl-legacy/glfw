// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

#include "callback.h"

void GLFWCALL windowSizeCallback(int width, int height) {
    goWindowSizeCB(width, height);
}

int GLFWCALL windowCloseCallback() {
    goWindowCloseCB();
}

void GLFWCALL windowRefreshCallback() {
    goWindowRefreshCB();
}

void GLFWCALL mouseButtonCallback(int button, int action) {
    goMouseButtonCB(button, action);
}

void GLFWCALL mousePosCallback(int x, int y) {
    goMousePosCB(x, y);
}

void GLFWCALL mouseWheelCallback(int pos) {
    goMouseWheelCB(pos);
}

void GLFWCALL keyCallback(int key, int action) {
    goKeyCB(key, action);
}

void GLFWCALL charCallback(int character, int action) {
    goCharCB(character, action);
}

void setWindowSizeCB()
{
	glfwSetWindowSizeCallback(windowSizeCallback);
}

void setWindowCloseCB()
{
	glfwSetWindowCloseCallback(windowCloseCallback);
}

void setWindowRefreshCB()
{
	glfwSetWindowRefreshCallback(windowRefreshCallback);
}

void setMouseButtonCB()
{
	glfwSetMouseButtonCallback(mouseButtonCallback);
}

void setMousePosCB()
{
	glfwSetMousePosCallback(mousePosCallback);
}

void setMouseWheelCB()
{
	glfwSetMouseWheelCallback(mouseWheelCallback);
}

void setKeyCB()
{
	glfwSetKeyCallback(keyCallback);
}

void setCharCB()
{
	glfwSetCharCallback(charCallback);
}
