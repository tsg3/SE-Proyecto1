package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"serverHome/resources"
	"serverHome/signals"
	"strconv"

	"goji.io/pat"
)

func TurnOnLight(w http.ResponseWriter, r *http.Request) {

	setCORS(&w, r)

	id, err := strconv.Atoi(pat.Param(r, "light"))

	if err != nil {
		fmt.Fprintf(w, "Conversion error: %s!", err)
		return
	}

	err = signals.TurnOnPin(id)

	if err != nil {
		fmt.Fprintf(w, "Error while turning pin %d: %s!", id, err)
		return
	}

	sendTurnedOnOK(w, id)

}

func TurnOffLight(w http.ResponseWriter, r *http.Request) {

	setCORS(&w, r)

	id, err := strconv.Atoi(pat.Param(r, "light"))

	if err != nil {
		fmt.Fprintf(w, "Conversion error: %s!", err)
		return
	}

	err = signals.TurnOffPin(id)

	if err != nil {
		fmt.Fprintf(w, "Error while turning pin %d: %s!", id, err)
		return
	}

	sendTurnedOffOK(w, id)

}

func TurnOnAllLights(w http.ResponseWriter, r *http.Request) {

	setCORS(&w, r)

	err := signals.TurnOnAllLights()
	if err != nil {
		fmt.Fprintf(w, "Error while turning lights: %s!", err)
		return
	}

	respones := resources.StateResource{
		Id:    -1,
		State: "ALLON",
	}

	json, err := json.Marshal(respones)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendResponse(w, json)
}

func TurnOffAllLights(w http.ResponseWriter, r *http.Request) {

	setCORS(&w, r)
	err := signals.TurnOffAllLights()
	if err != nil {
		fmt.Fprintf(w, "Error while turning lights: %s!", err)
		return
	}
	respones := resources.StateResource{
		Id:    -1,
		State: "ALLOFF",
	}

	json, err := json.Marshal(respones)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendResponse(w, json)

}

func GetAllLights(w http.ResponseWriter, r *http.Request) {

	setCORS(&w, r)
	lightsState := signals.ReadAllLights()

	json, err := json.Marshal(lightsState)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendResponse(w, json)

}
