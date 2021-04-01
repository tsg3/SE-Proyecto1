package routes

import (
	"serverHome/controllers"

	"goji.io"
	"goji.io/pat"
)

func set_door_routes(door *goji.Mux) {

	door.HandleFunc(pat.Get("/getState"), controllers.GetDoorsState)

}
