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
	return turnOffPin(pin, true)
}

func TurnOnAllLights() error {
	return turnOnAllLights()

}

func TurnOffAllLights() error {
	return turnOffAllLights()
}

func TakePhoto() string {

	data := takePhoto()

	return data

}

func ReadAllDoors() []resources.StateResource {
	return readAllDoors()
}

func ReadAllLights() []resources.StateResource {
	return readAllLights()
}
