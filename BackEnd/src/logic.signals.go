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
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"serverHome/resources"
	"strconv"
	"unsafe"
)

var DOORS = [4]string{"26", "9", "16", "11"}
var LIGHTS = [5]string{"17", "18", "22", "23", "27"}

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

func signalsInit() error {

	// Lights set as Write
	for _, light := range LIGHTS {
		setPin(light, true)
	}

	//  Doors set as Only Read
	for _, door := range DOORS {
		setPin(door, true)
	}

	return nil
}

func signalsOff() error {
	// Lights set as Write
	for _, light := range LIGHTS {
		unSetPin(light)
	}

	//  Doors set as Only Read
	for _, door := range DOORS {
		unSetPin(door)
	}

	return nil

}

func checkLight(pin int) error {
	if pin >= len(LIGHTS) {
		return errors.New("ID OUT of RANGE")
	}
	return nil
}

func turnOnPin(pin int) error {
	err := checkLight(pin)
	if err != nil {
		return err
	}
	fmt.Println("Prendiendo" + LIGHTS[pin])
	pinNum := C.CString(LIGHTS[pin])
	oneV := C.CString("1")
	C.digitalWrite(pinNum, oneV)
	C.free(unsafe.Pointer(oneV))
	C.free(unsafe.Pointer(pinNum))

	return nil
}

func turnOffPin(pin int, mode bool) error {
	err := checkLight(pin)

	if err != nil {
		return err
	}

	var pinNum *C.char
	if mode {
		fmt.Println("Apangando" + LIGHTS[pin])
		pinNum = C.CString(LIGHTS[pin])
	} else {
		fmt.Println("Apangando" + DOORS[pin])
		pinNum = C.CString(DOORS[pin])
	}

	ceroV := C.CString("0")
	C.digitalWrite(pinNum, ceroV)
	C.free(unsafe.Pointer(ceroV))
	C.free(unsafe.Pointer(pinNum))

	return nil
}

func turnOnAllLights() error {
	// Lights set as Write
	for i, _ := range LIGHTS {
		err := turnOnPin(i)
		if err != nil {
			return err
		}
	}

	return nil

}

func turnOffAllLights() error {
	for i, _ := range LIGHTS {

		err := turnOffPin(i, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func takePhoto() string {
	C.takePhoto()

	imgFile, err := os.Open("photo.jpeg") // a QR code image

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)
	encodedStr := base64.StdEncoding.EncodeToString(buf)

	return encodedStr

}

func readPin(pin string) string {

	pinNum := C.CString(pin)

	state := C.digitalRead(pinNum)

	result := strconv.Itoa(int(state))

	C.free(unsafe.Pointer(pinNum))

	return result

}

// func readPinDoor(pin string) string {

// 	setPin(pin, false)
// 	pinNum := C.CString(pin)

// 	state := C.digitalRead(pinNum)

// 	result := C.GoString(state)

// 	C.free(unsafe.Pointer(pinNum))
// 	unSetPin(pin)

// 	return result

// }

func readAllDoors() []resources.StateResource {

	doorsObjs := []resources.StateResource{}

	for i, door := range DOORS {
		state := readPin(door)
		doorObj := resources.StateResource{
			Id:    i,
			State: state,
		}

		doorsObjs = append(doorsObjs, doorObj)
	}

	return doorsObjs
}

func readAllLights() []resources.StateResource {

	doorsObjs := []resources.StateResource{}

	for i, light := range LIGHTS {
		state := readPin(light)
		lightObj := resources.StateResource{
			Id:    i,
			State: state,
		}

		doorsObjs = append(doorsObjs, lightObj)
	}

	return doorsObjs
}
