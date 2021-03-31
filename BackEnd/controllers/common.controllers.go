package controllers

import (
	"net/http"
)

// Common function to send the response
func sendResponse(w http.ResponseWriter, infoToSend []byte) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(infoToSend)

}
