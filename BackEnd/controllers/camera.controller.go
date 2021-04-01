package controllers

import (
	"encoding/json"
	"net/http"
	"serverHome/resources"
	signals "serverHome/src"
)

func TakePhoto(w http.ResponseWriter, r *http.Request) {
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
