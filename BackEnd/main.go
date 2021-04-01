package main

import (
	"fmt"
	"net"
	"net/http"
	"serverHome/controllers"
	"serverHome/routes"
	"strings"
)

const PORT = "5000"

func get_ip() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		s := strings.Split(addr.String(), "/")
		if s[1] == "24" {
			return s[0], nil
		}
	}

	return "", nil
}

func main() {
	// Get the IP
	IP, err := get_ip()

	if IP == "" || err != nil {
		panic(err)
	}

	// Init AdminUser
	controllers.InitAdminUser()
	// Initialize the routes and the server
	mux := routes.Server_init()
	// signals.SignalsInit()
	// signals.SignalsOff()

	fmt.Println("Starting server on: " + IP + ":" + PORT)

	// Starting the server in the IP and PORT assigned before
	err_s := http.ListenAndServe(IP+":"+PORT, mux)

	if err_s != nil {
		fmt.Println("There ir a error %v", err_s)
		panic(err_s)

	}
}
