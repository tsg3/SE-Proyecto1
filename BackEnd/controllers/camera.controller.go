package controllers

import (
	"encoding/json"
	"net/http"
	"serverHome/resources"
	"serverHome/signals"
)

func TakePhoto(w http.ResponseWriter, r *http.Request) {
	setCORS(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	data := signals.TakePhoto()

	res := resources.CameraResource{
		Data: data,
	}

	json, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendResponse(w, json)

}
