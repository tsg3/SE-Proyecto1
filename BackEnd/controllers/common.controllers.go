package controllers

import (
	"encoding/json"
	"net/http"
	"serverHome/resources"
)

// Common function to send the response
func sendResponse(w http.ResponseWriter, infoToSend []byte) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(infoToSend)

}

func sendTurnedOnOK(w http.ResponseWriter, id int) {

	light := resources.StateResource{
		Id:    id,
		State: "1",
	}

	json, err := json.Marshal(light)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendResponse(w, json)
}

func sendTurnedOffOK(w http.ResponseWriter, id int) {

	light := resources.StateResource{
		Id:    id,
		State: "0",
	}

	json, err := json.Marshal(light)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendResponse(w, json)
}

func setCORS(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Authorization")
	(*w).Header().Set("Content-Type", "application/json")
}
