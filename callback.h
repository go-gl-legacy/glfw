// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

#ifndef _CALLBACK_H_
#define _CALLBACK_H_

#include <stdlib.h>
#ifdef _WIN32
  #define GLFW_DLL
#endif
#include <GL/glfw.h>

void GLFWCALL windowSizeCallback(int, int);
int GLFWCALL windowCloseCallback(void);
void GLFWCALL windowRefreshCallback(void);
void GLFWCALL mouseButtonCallback(int, int);
void GLFWCALL mousePosCallback(int, int);
void GLFWCALL mouseWheelCallback(int);
void GLFWCALL keyCallback(int, int);
void GLFWCALL charCallback(int, int);

extern void goWindowSizeCB(int, int);
extern int  goWindowCloseCB(void);
extern void goWindowRefreshCB(void);
extern void goMouseButtonCB(int, int);
extern void goMousePosCB(int, int);
extern void goMouseWheelCB(int);
extern void goKeyCB(int, int);
extern void goCharCB(int, int);

void setWindowSizeCB();
void setWindowCloseCB();
void setWindowRefreshCB();
void setMouseButtonCB();
void setMousePosCB();
void setMouseWheelCB();
void setKeyCB();
void setCharCB();

#endif // _CALLBACK_H_
