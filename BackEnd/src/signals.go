package signals

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L./lib -lpinControl
#include <pinControl.h>
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

// import "unsafe"

import (
	"unsafe"
)

func setPin(pin string) {

	pinNum := C.CString(pin)
	C.pinMode(pinNum, true)
	C.free(unsafe.Pointer(pinNum))

}

func TurnOnPin(pin string) {
	pinNum := C.CString(pin)
	setPin(pin)
	oneV := C.CString("1")
	C.digitalWrite(pinNum, oneV)
	C.free(unsafe.Pointer(oneV))
	C.free(unsafe.Pointer(pinNum))
}

func TurnOffPin(pin string) {
	pinNum := C.CString(pin)
	setPin(pin)
	ceroV := C.CString("0")
	C.digitalWrite(pinNum, ceroV)
	C.free(unsafe.Pointer(ceroV))
	C.free(unsafe.Pointer(pinNum))
}
