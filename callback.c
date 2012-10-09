// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

#include "callback.h"

void GLFWCALL glfwWindowSizeCB (int width, int height) {
    goWindowSizeCB(width, height);
}

int GLFWCALL glfwWindowCloseCB (void) {
    goWindowCloseCB();
}

void GLFWCALL glfwWindowRefreshCB (void) {
    goWindowRefreshCB();
}

void GLFWCALL glfwMouseButtonCB (int button, int action) {
    goMouseButtonCB(button, action);
}

void GLFWCALL glfwMousePosCB (int x, int y) {
    goMousePosCB(x, y);
}

void GLFWCALL glfwMouseWheelCB (int pos) {
    goMouseWheelCB(pos);
}

void GLFWCALL glfwKeyCB (int key, int action) {
    goKeyCB(key, action);
}

void GLFWCALL glfwCharCB (int character, int action) {
    goCharCB(character, action);
}

void glfwSetWindowSizeCB (void)
{
	glfwSetWindowSizeCallback(glfwWindowSizeCB);
}

void glfwSetWindowCloseCB (void)
{
	glfwSetWindowCloseCallback(glfwWindowCloseCB);
}

void glfwSetWindowRefreshCB (void)
{
	glfwSetWindowRefreshCallback(glfwWindowRefreshCB);
}

void glfwSetMouseButtonCB (void)
{
	glfwSetMouseButtonCallback(glfwMouseButtonCB);
}

void glfwSetMousePosCB (void)
{
	glfwSetMousePosCallback(glfwMousePosCB);
}

void glfwSetMouseWheelCB (void)
{
	glfwSetMouseWheelCallback(glfwMouseWheelCB);
}

void glfwSetKeyCB (void)
{
	glfwSetKeyCallback(glfwKeyCB);
}

void glfwSetCharCB (void)
{
	glfwSetCharCallback(glfwCharCB);
}
