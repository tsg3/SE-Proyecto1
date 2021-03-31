package main

import (
	"fmt"
	"home_server/routes"
	"net/http"
)

const PORT = "5000"
const IP = "127.0.0.1"

func main() {
	// Initialize the routes and the server
	mux := routes.Server_init()

	fmt.Println("Starting server")
	// Starting the server in the IP and PORT assigned before
	http.ListenAndServe(IP+":"+PORT, mux)
	fmt.Println("Starting server")
}
