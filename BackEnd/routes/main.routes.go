package routes

import (
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func enableCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers", "authorization")
}

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
