package routes

import (
	"goji.io"
	"goji.io/pat"
)

// Init the routes of the API
func Server_init() *goji.Mux {

	// Set the root route
	root := goji.NewMux()
	// Establish the routes for each elemento of the API
	camera := goji.SubMux()
	lights := goji.SubMux()
	doors := goji.SubMux()
	login := goji.SubMux()

	// Set the subroutes
	root.Handle(pat.New("/api/camera/*"), camera)
	root.Handle(pat.New("/api/lights/*"), lights)
	root.Handle(pat.New("/api/doors/*"), doors)
	root.Handle(pat.New("/api/login/*"), login)

	// Set the routes for each element
	set_camera_routes(camera)
	set_light_routes(lights)
	set_door_routes(doors)
	set_login_routes(login)

	return root
}
