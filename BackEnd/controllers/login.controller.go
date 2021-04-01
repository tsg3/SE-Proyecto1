package controllers

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logged")
}
