// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package glfw

// GLFW Version
const (
	VersionMajor    = 2
	VersionMinor    = 7
	VersionRevision = 0
)

// Key and button state/action definitions
const (
	KeyRelease = iota
	KeyPress
)

// Keyboard key definitions: 8-bit ISO-8859-1 (Latin 1) encoding is used
// for printable keys (such as A-Z, 0-9 etc), and values above 256
// represent Special (non-printable) keys (e.g. F1, Page Up etc).
const (
	KeyUnknown = -1
	KeySpace   = 32
	KeySpecial = 256
)

const (
	_ = (KeySpecial + iota)
	KeyEsc
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyF13
	KeyF14
	KeyF15
	KeyF16
	KeyF17
	KeyF18
	KeyF19
	KeyF20
	KeyF21
	KeyF22
	KeyF23
	KeyF24
	KeyF25
	KeyUp
	KeyDown
	KeyLeft
	KeyRight
	KeyLshift
	KeyRshift
	KeyLctrl
	KeyRctrl
	KeyLalt
	KeyRalt
	KeyTab
	KeyEnter
	KeyBackspace
	KeyInsert
	KeyDel
	KeyPageup
	KeyPagedown
	KeyHome
	KeyEnd
	KeyKP0
	KeyKP1
	KeyKP2
	KeyKP3
	KeyKP4
	KeyKP5
	KeyKP6
	KeyKP7
	KeyKP8
	KeyKP9
	KeyKPDivide
	KeyKPMultiply
	KeyKPSubtract
	KeyKPAdd
	KeyKPDecimal
	KeyKPEqual
	KeyKPEnter
	KeyKPNumlock
	KeyCapslock
	KeyScrolllock
	KeyPause
	KeyLsuper
	KeyRsuper
	KeyMenu
	KeyLast = KeyMenu
)

// Mouse button definitions
const (
	Mouse1 = iota
	Mouse2
	Mouse3
	Mouse4
	Mouse5
	Mouse6
	Mouse7
	Mouse8
	MouseLast   = Mouse8
	MouseLeft   = Mouse1
	MouseRight  = Mouse2
	MouseMiddle = Mouse3
)

// Joystick identifiers
const (
	Joy1 = iota
	Joy2
	Joy3
	Joy4
	Joy5
	Joy6
	Joy7
	Joy8
	Joy9
	Joy10
	Joy11
	Joy12
	Joy13
	Joy14
	Joy15
	Joy16
	JoyLast = Joy16
)

// glfwOpenWindow modes
const (
	_ = 0x00010000 + iota
	Windowed
	Fullscreen
)

// glfwGetWindowParam / glfwOpenWindowHint tokens
const (
	_ = 0x00020000 + iota
	Opened
	Active
	Iconified
	Accelerated
	RedBits
	GreenBits
	BlueBits
	AlphaBits
	DepthBits
	StencilBits

	// The following constants are used for both glfwGetWindowParam
	// and glfwOpenWindowHint
	RefreshRate
	AccumRedBits
	AccumGreenBits
	AccumBlueBits
	AccumAlphaBits
	AuxBuffers
	Stereo
	WindowNoResize
	FsaaSamples
	OpenGLVersionMajor
	OpenGLVersionMinor
	OpenGLForwardCompat
	OpenGLDebugContext
	OpenGLProfile
)

// GLFW_OPENGL_PROFILE tokens
const (
	_ = 0x00050000 + iota
	OpenGLCoreProfile
	OpenGLCompatProfile
)

// glfwEnable/glfwDisable tokens 
const (
	_ = 0x00030000 + iota
	MouseCursor
	StickyKeys
	StickyMouseButtons
	SystemKeys
	KeyRepeat
	AutoPollEvents
)

// glfwWaitThread wait modes
const (
	_ = 0x00040000 + iota
	Wait
	NoWait
)

// glfwGetJoystickParam tokens 
const (
	_ = 0x00050000 + iota
	Present
	Axes
	Buttons
)

// glfwReadImage/glfwLoadTexture2D flags 
const (
	NoRescaleBit = 1 << iota // Only for glfwReadImage
	OriginUlBit
	BuildMipmapsBit // Only for glfwLoadTexture2D
	AlphaMapBit
)

// Time spans longer than this (seconds) are considered to be infinity
const Infinity = 100000.0
