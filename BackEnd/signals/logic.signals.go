package signals

/*
#cgo CFLAGS: -I../src/include
#cgo LDFLAGS: -L../src/lib -lpinControl
#include <pinControl.h>
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

// import "unsafe"

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
		setPin(door, false)
	}

	err := setPinValues()

	return err
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

func switchPinValue(pin int, mode bool) error {
	err := checkLight(pin)

	if err != nil {
		return err
	}

	var value *C.char
	if mode {
		fmt.Println("Encendiendo" + LIGHTS[pin])
		value = C.CString("1")
	} else {
		fmt.Println("Apangando" + LIGHTS[pin])
		value = C.CString("0")
	}

	pinNum := C.CString(LIGHTS[pin])
	C.digitalWrite(pinNum, value)
	C.free(unsafe.Pointer(value))
	C.free(unsafe.Pointer(pinNum))

	return nil
}

func switchAllPinValues(mode bool) error {
	for i, _ := range LIGHTS {
		err := switchPinValue(i, mode)
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

	lightsObjs := []resources.StateResource{}

	for i, light := range LIGHTS {
		state := readPin(light)
		lightObj := resources.StateResource{
			Id:    i,
			State: state,
		}

		lightsObjs = append(lightsObjs, lightObj)
	}

	return lightsObjs
}

func saveLightsState(states []resources.StateResource) error {

	file, err := json.MarshalIndent(states, "", " ")

	if err != nil {
		fmt.Printf("There is an error saving a file: %s\n", err)
		return err
	}

	err = ioutil.WriteFile("lightsStates.json", file, 0644)

	if err != nil {
		fmt.Printf("There is an error saving a file: %s\n", err)
	}

	return err
}

func readStatesFromFile() ([]resources.StateResource, error) {
	file, err := ioutil.ReadFile("lightsStates.json")

	var states []resources.StateResource
	if err != nil {
		fmt.Printf("There is an error saving a file: %s\n", err)
		return states, err
	}

	err = json.Unmarshal([]byte(file), &states)

	if err != nil {
		fmt.Printf("There is an error saving a file: %s\n", err)
		return states, err
	}

	return states, nil
}

// This function sets the value of each pin, when the system starts
func setPinValues() error {
	if _, err := os.Stat("lightsStates.json"); os.IsNotExist(err) {
		fmt.Println("File lightsStates.json doesn't exists")
		return err
	}

	states, err := readStatesFromFile()

	if err != nil {
		return err
	}
	var mode bool
	for _, state := range states {
		if state.State == "1" {
			mode = true
		} else {
			mode = false
		}
		err = switchPinValue(state.Id, mode)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateLightsState(err error) error {
	if err != nil {
		return err
	}

	states := readAllLights()
	err_s := saveLightsState(states)

	return err_s
}
