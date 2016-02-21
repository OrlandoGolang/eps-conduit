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

// Handling user flags
// User flags must be package globals they can be easily worked on by Config member functions
// and avoid passing each command line option as a parameter.
var configFile = flag.String("config", "/etc/conduit.conf", "Path to config file. Default is /etc/conduit.conf")
var backendStr = flag.String("b", "", "Host strings for the backend services (comma separated)")
var bind = flag.String("bind", "", "The port the load balancer should listen to")
var mode = flag.String("mode", "", "Balancing Mode")
var certFile = flag.String("cert", "", "Path to rsa private key")
var keyFile = flag.String("key", "", "Path to rsa public key")
var accessLog = flag.String("log", "", "Path to store access logs")

func main() {
	flag.Parse()
	config := GetConfig(*configFile)

	// send requests to proxies via config.handle
	http.HandleFunc("/", LoggingMiddleware(config.handle))

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
