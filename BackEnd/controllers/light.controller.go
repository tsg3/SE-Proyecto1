package controllers

import (
	signals "home_server/src"
	"net/http"
)

func TurnOnLight(w http.ResponseWriter, r *http.Request) {

	signals.TurnOnPin("18")

}

func TurnOffLight(w http.ResponseWriter, r *http.Request) {

	signals.TurnOffPin("18")

}

func TurnOffAllLights(w http.ResponseWriter, r *http.Request) {

}

func TurnOnAllLights(w http.ResponseWriter, r *http.Request) {

}
