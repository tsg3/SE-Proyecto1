package routes

import (
	"serverHome/controllers"

	"goji.io"
	"goji.io/pat"
)

func set_login_routes(login *goji.Mux) {

	login.HandleFunc(pat.Post("/"), controllers.Login)

}
