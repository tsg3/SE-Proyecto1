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
	"fmt"
	"serverHome/resources"
	"unsafe"
)

var DOORS = [4]string{"5", "6", "16", "26"}
var LIGHTS = [5]string{"17", "18", "27", "22", "23"}

func setPin(pin string, mode bool) {

	modeC := C.bool(mode)
	pinNum := C.CString(pin)
	C.pinMode(pinNum, modeC)
	C.free(unsafe.Pointer(pinNum))

}

func unSetPin(pin string) {
	pinNum := C.CString(pin)
	C.unExportPin(pinNum)
	C.free(unsafe.Pointer(pinNum))
}

func signalsInit() {

	// Lights set as Write
	for _, light := range LIGHTS {
		setPin(light, true)
	}

	//  Doors set as Only Read
	for _, door := range DOORS {
		setPin(door, false)
	}
}

func signalsOff() {
	// Lights set as Write
	unSetPin("17")
	unSetPin("18")
	unSetPin("27")
	unSetPin("22")
	unSetPin("23")

	//  Doors set as Only Read
	unSetPin("5")
	unSetPin("6")
	unSetPin("16")
	unSetPin("26")
}

func turnOnPin(pin string) {
	fmt.Println("Prendiendo" + pin)
	pinNum := C.CString(pin)
	oneV := C.CString("1")
	C.digitalWrite(pinNum, oneV)
	C.free(unsafe.Pointer(oneV))
	C.free(unsafe.Pointer(pinNum))
}

func turnOffPin(pin string) {
	fmt.Println("Apangando" + pin)
	pinNum := C.CString(pin)
	ceroV := C.CString("0")
	C.digitalWrite(pinNum, ceroV)
	C.free(unsafe.Pointer(ceroV))
	C.free(unsafe.Pointer(pinNum))
}

func turnOnAllLights() {
	// Lights set as Write
	for _, light := range LIGHTS {
		turnOnPin(light)
	}

}

func turnOffAllLights() {
	for _, light := range LIGHTS {
		turnOffPin(light)
	}

}

func takePhoto() string {
	var ptr *C.char
	ar := C.takePhoto()
	ptr = C.getPhoto(ar)

	data := C.GoString(ptr)

	return data

}

func readPin(pin string) string {

	pinNum := C.CString(pin)
	state := C.digitalRead(pinNum)
	C.free(unsafe.Pointer(pinNum))

	return state

}

func readAllDoors() []resources.DoorResource {

	doorsObjs := []resources.DoorResource{}

	for i, door := range DOORS {
		state := readPin(door)
		doorObj := resources.DoorResource{
			Id:    i,
			State: state,
		}

		doorsObjs = append(doorsObjs, doorObj)
	}

	return doorsObjs
}
