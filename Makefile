# This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
# license. Its contents can be found at:
# http://creativecommons.org/publicdomain/zero/1.0/

include $(GOROOT)/src/Make.inc

TARG = github.com/jteeuwen/glfw
GOFILES = constants.go
CGOFILES = glfw.go vidmode.go image.go
CGO_CFLAGS = `pkg-config --cflags libglfw`
CGO_LDFLAGS = `pkg-config --libs libglfw`

include $(GOROOT)/src/Make.pkg
