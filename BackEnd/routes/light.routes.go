package routes

import (
	"home_server/controllers"

	"goji.io"
	"goji.io/pat"
)

func set_light_routes(lights *goji.Mux) {

	lights.HandleFunc(pat.Post("/turnOn/:light"), controllers.TurnOnLight)
	lights.HandleFunc(pat.Post("/turnOff/:light"), controllers.TurnOffLight)
	lights.HandleFunc(pat.Post("/turnOnAll"), controllers.TurnOnAllLights)
	lights.HandleFunc(pat.Post("/turnOffAll"), controllers.TurnOffAllLights)

}
