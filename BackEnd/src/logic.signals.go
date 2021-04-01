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
	"unsafe"
)

var DOORS = [4]string{"5", "6", "16", "26"}
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
		setPin(door, false)
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

func turnOffPin(pin int) error {
	err := checkLight(pin)

	if err != nil {
		return err
	}
	fmt.Println("Apangando" + LIGHTS[pin])
	pinNum := C.CString(LIGHTS[pin])
	ceroV := C.CString("0")
	C.digitalWrite(pinNum, ceroV)
	C.free(unsafe.Pointer(ceroV))
	C.free(unsafe.Pointer(pinNum))

	return nil
}

func turnOnAllLights() ([]resources.StateResource, error) {
	// Lights set as Write
	for i, _ := range LIGHTS {
		err := turnOnPin(i)
		if err != nil {
			return nil, err
		}
	}

	return readAllLights(), nil

}

func turnOffAllLights() ([]resources.StateResource, error) {
	for i, _ := range LIGHTS {

		err := turnOffPin(i)
		if err != nil {
			return nil, err
		}
	}

	return readAllLights(), nil
}

func takePhoto() string {
	// var ptr *C.char
	C.takePhoto()
	// ptr = C.getPhoto(ar)

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

	// if you create a new image instead of loading from file, encode the image to buffer instead with png.Encode()

	// png.Encode(&buf, image)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	encodedStr := base64.StdEncoding.EncodeToString(buf)

	// data := C.GoString(ptr)

	// encodedStr := base64.StdEncoding.EncodeToString([]byte(data))

	return encodedStr

}

func readPin(pin string) string {

	pinNum := C.CString(pin)
	state := C.digitalRead(pinNum)
	C.free(unsafe.Pointer(pinNum))

	result := C.GoString(state)

	return result

}

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
