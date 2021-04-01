package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	signals "serverHome/src"
	"strconv"

	"goji.io/pat"
)

func TurnOnLight(w http.ResponseWriter, r *http.Request) {

	setCORS(&w, r)

	id, err := strconv.Atoi(pat.Param(r, "id"))

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

	id, err := strconv.Atoi(pat.Param(r, "id"))

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

	ligths, err := signals.TurnOnAllLights()
	if err != nil {
		fmt.Fprintf(w, "Error while turning lights: %s!", err)
		return
	}

	json, err := json.Marshal(ligths)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendResponse(w, json)
}

func TurnOffAllLights(w http.ResponseWriter, r *http.Request) {

	setCORS(&w, r)
	ligths, err := signals.TurnOnAllLights()
	if err != nil {
		fmt.Fprintf(w, "Error while turning lights: %s!", err)
		return
	}

	json, err := json.Marshal(ligths)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendResponse(w, json)

}
