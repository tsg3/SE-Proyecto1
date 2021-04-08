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
	err := switchPinValue(pin, true)

	return updateLightsState(err)
}

func TurnOffPin(pin int) error {
	err := switchPinValue(pin, false)
	return updateLightsState(err)
}

func TurnOnAllLights() error {
	err := switchAllPinValues(true)
	return updateLightsState(err)

}

func TurnOffAllLights() error {
	err := switchAllPinValues(false)
	return updateLightsState(err)
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
