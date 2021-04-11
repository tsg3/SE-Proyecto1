package routes

import (
	"serverHome/controllers"

	"goji.io"
	"goji.io/pat"
)

func set_camera_routes(camera *goji.Mux) {
	camera.HandleFunc(pat.Get("/take"), controllers.TakePhoto)
}
