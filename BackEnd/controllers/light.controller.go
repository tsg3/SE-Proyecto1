package controllers

import (
	"net/http"
	signals "serverHome/src"

	"goji.io/pat"
)

func TurnOnLight(w http.ResponseWriter, r *http.Request) {
	id := pat.Param(r, "id")

	signals.TurnOnPin(id)

}

func TurnOffLight(w http.ResponseWriter, r *http.Request) {
	id := pat.Param(r, "id")

	signals.TurnOffPin(id)

}

func TurnOffAllLights(w http.ResponseWriter, r *http.Request) {
	signals.TurnOnAllLights()
}

func TurnOnAllLights(w http.ResponseWriter, r *http.Request) {
	signals.TurnOffAllLights()
}
