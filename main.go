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
var backendStr = flag.String("b", "", "target ips for backend servers")
var bind = flag.String("bind", "", "port to bind to")
var mode = flag.String("mode", "", "Balancing Mode")
var certFile = flag.String("cert", "", "cert")
var keyFile = flag.String("key", "", "key")

func main() {

	flag.Parse()
	config := GetConfig(*configFile)
	config.handleUserInput()
	config.printConfigInfo()
	proxies := config.makeProxies()
	hostCount := len(config.Backends)

	// Function for handling the http requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		nextHost = pickHost(nextHost, hostCount)
		proxies[nextHost].ServeHTTP(w, r)
	})

	//Start the http(s) listener depending on user's selected mode
	if config.Mode == "http" {
		http.ListenAndServe(":"+config.Bind, nil)
	} else if config.Mode == "https" {
		http.ListenAndServeTLS(":"+config.Bind, config.Certfile, config.Keyfile, nil)
	} else {
		fmt.Fprintf(os.Stderr, "unknown mode or mode not set")
		os.Exit(1)
	}
}

// Function for determining what the next backend host should be
func pickHost(lastHost, hostCount int) int {
	x := lastHost + 1
	if x >= hostCount {
		return 0
	}
	return x
}
