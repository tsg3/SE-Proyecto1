package routes

import (
	"home_server/controllers"

	"goji.io"
	"goji.io/pat"
)

func set_login_routes(login *goji.Mux) {

	login.HandleFunc(pat.Get("/login"), controllers.Login)

}
