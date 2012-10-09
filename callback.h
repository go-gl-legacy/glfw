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

void GLFWCALL glfwWindowSizeCB(int, int);
int  GLFWCALL glfwWindowCloseCB(void);
void GLFWCALL glfwWindowRefreshCB(void);
void GLFWCALL glfwMouseButtonCB(int, int);
void GLFWCALL glfwMousePosCB(int, int);
void GLFWCALL glfwMouseWheelCB(int);
void GLFWCALL glfwKeyCB(int, int);
void GLFWCALL glfwCharCB(int, int);

extern void goWindowSizeCB(int, int);
extern int  goWindowCloseCB(void);
extern void goWindowRefreshCB(void);
extern void goMouseButtonCB(int, int);
extern void goMousePosCB(int, int);
extern void goMouseWheelCB(int);
extern void goKeyCB(int, int);
extern void goCharCB(int, int);

void glfwSetWindowSizeCB(void);
void glfwSetWindowCloseCB(void);
void glfwSetWindowRefreshCB(void);
void glfwSetMouseButtonCB(void);
void glfwSetMousePosCB(void);
void glfwSetMouseWheelCB(void);
void glfwSetKeyCB(void);
void glfwSetCharCB(void);

#endif // _CALLBACK_H_
