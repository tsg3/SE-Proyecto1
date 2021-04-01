package controllers

import (
	"encoding/json"
	"net/http"
	signals "serverHome/src"
)

func GetDoorsState(w http.ResponseWriter, r *http.Request) {
	setCORS(&w, r)
	states := signals.ReadAllDoors()

	json, err := json.Marshal(states)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendResponse(w, json)
}
