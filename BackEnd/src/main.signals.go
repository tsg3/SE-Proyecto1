package signals

import "serverHome/resources"

func SignalsInit() {
	signalsInit()
}

func SignalsOff() {
	signalsOff()
}

func TurnOnPin(pin string) {
	turnOnPin(pin)
}

func TurnOffPin(pin string) {
	turnOffPin(pin)
}

func TurnOnAllLights() {
	turnOffAllLights()
}

func TurnOffAllLights() {
	turnOffAllLights()
}

func TakePhoto() string {

	data := takePhoto()

	return data

}

func ReadAllDoors() []resources.DoorResource {
	return readAllDoors()
}
