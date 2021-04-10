package main

import (
	"fmt"
	"net/http"
	"serverHome/controllers"
	"serverHome/network"
	"serverHome/routes"
	"serverHome/signals"
)

const PORT = "5000"

func main() {
	// Get the IP
	IP, err := network.GetPersonalIp()

	if IP == "" || err != nil {
		panic(err)
	}

	// Init AdminUser
	controllers.InitAdminUser()
	// Initialize the routes and the server
	mux := routes.Server_init()
	signals.SignalsInit()

	fmt.Println("Starting server on: " + IP + ":" + PORT)

	// Starting the server in the IP and PORT assigned before
	err_s := http.ListenAndServe(IP+":"+PORT, mux)

	signals.SignalsOff()

	if err_s != nil {
		fmt.Printf("There is a error %s\n", err_s)
		panic(err_s)
	}
}
