// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

#include "callback.h"

void setWindowSizeCB()
{
	glfwSetWindowSizeCallback(goWindowSizeCB);
}

void setWindowCloseCB()
{
	glfwSetWindowCloseCallback(goWindowCloseCB);
}

void setWindowRefreshCB()
{
	glfwSetWindowRefreshCallback(goWindowRefreshCB);
}

void setMouseButtonCB()
{
	glfwSetMouseButtonCallback(goMouseButtonCB);
}

void setMousePosCB()
{
	glfwSetMousePosCallback(goMousePosCB);
}

void setMouseWheelCB()
{
	glfwSetMouseWheelCallback(goMouseWheelCB);
}

void setKeyCB()
{
	glfwSetKeyCallback(goKeyCB);
}

void setCharCB()
{
	glfwSetCharCallback(goCharCB);
}
