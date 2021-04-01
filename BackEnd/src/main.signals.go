package signals

import (
	"serverHome/resources"
)

func SignalsInit() error {
	return signalsInit()
}

func SignalsOff() error {
	return signalsOff()
}

func TurnOnPin(pin int) error {
	return turnOnPin(pin)
}

func TurnOffPin(pin int) error {
	return turnOffPin(pin)
}

func TurnOnAllLights() ([]resources.StateResource, error) {
	return turnOffAllLights()

}

func TurnOffAllLights() ([]resources.StateResource, error) {
	return turnOffAllLights()
}

func TakePhoto() string {

	data := takePhoto()

	return data

}

func ReadAllDoors() []resources.StateResource {
	return readAllDoors()
}
