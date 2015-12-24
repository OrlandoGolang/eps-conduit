/*
main.go

Description:
	Eps-Conduit is a light-weight load balancer.

Source Code:
	https://github.com/orlandogolang/eps-conduit
*/
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

// Global variable of the next backend to be sent.  This is for round-robin load balancing
var nextHost int = 0

// Handling user flags
// User flags must be package globals they can be easily worked on by Config member functions
// and avoid passing each command line option as a parameter.
var configFile = flag.String("config", "/etc/conduit.conf", "Path to config file. Default is /etc/conduit.conf")
var backendStr = flag.String("b", "", "Host strings for the backend services (comma separated)")
var bind = flag.String("bind", "", "The port the load balancer should listen to")
var mode = flag.String("mode", "", "Balancing Mode")
var certFile = flag.String("cert", "", "Path to rsa private key")
var keyFile = flag.String("key", "", "Path to rsa public key")

func main() {

	flag.Parse()
	config := GetConfig(*configFile)
	config.handleUserInput()
	config.printConfigInfo()
	hostCount := len(config.Backends)

	// Function for handling the http requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		nextHost = pickHost(nextHost, hostCount)
		config.Proxies[nextHost].ServeHTTP(w, r)
	})

	// Start the http(s) listener depending on user's selected mode
	if config.Mode == "http" {
		http.ListenAndServe(":"+config.Bind, nil)
	} else if config.Mode == "https" {
		http.ListenAndServeTLS(":"+config.Bind, config.Certfile, config.Keyfile, nil)
	} else {
		fmt.Fprintf(os.Stderr, "unknown mode or mode not set")
		os.Exit(1)
	}
}

// pickHost determines the next backend host to forward the request to - according to round-robin
// It returns an integer, which represents the host's index in config.Backends
func pickHost(lastHost, hostCount int) int {
	x := lastHost + 1
	if x >= hostCount {
		return 0
	}
	return x
}
