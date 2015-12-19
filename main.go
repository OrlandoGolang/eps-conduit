package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	//"strings"
)

//Global variable of the next backend to be sent.  This is for round-robin load balancing
var nextHost int = 0

func main() {

	//Handling user flags
	configFile := flag.String("config","/etc/conduit.conf", "Path to config file. Default is /etc/conduit.conf")
	//backendStr := flag.String("b", "", "target ips for backend servers")
	bind := flag.String("bind", "", "port to bind to")
	mode := flag.String("mode", "", "Balancing Mode")
	certFile := flag.String("cert", "", "cert")
	keyFile := flag.String("key", "", "key")

	flag.Parse()

	log.Println(*configFile)

	config := GetConfig(*configFile)

	//if *backendStr != "" {
		//Remove whitespace from backends
	//	*backendStr = strings.Replace(*backendStr, " ", "", -1)
		//Throwing backends into an array
	//	config.Backends = strings.Split(*backendStr, ",")
	//}

	totesHost := len(config.Backends)

	if *bind != "" {
		config.Bind = *bind
	}

	if *mode != "" {
		config.Mode = *mode
	}

	if *certFile != "" {
		config.Certfile = *certFile
	}

	if *keyFile != "" {
		config.Keyfile = *keyFile
	}

	//Create a proxy for each backend
	proxies := make([]*httputil.ReverseProxy, totesHost)
	for i := range config.Backends {
		host := config.Backends[i]
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = host
		}
		proxies[i] = &httputil.ReverseProxy{Director: director}
	}

	//Function for handling the http requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		nextHost = pickHost(nextHost, totesHost)
		proxies[nextHost].ServeHTTP(w, r)
	})

	//tell the user what info the load balancer is using
	for _, v := range config.Backends {
		log.Println("using " + v + " as a backend")
	}
	log.Println("listening on port " + *bind)

	//Start the http(s) listener depending on user's selected mode
	if *mode == "http" {
		http.ListenAndServe(":"+*bind, nil)
	} else if *mode == "https" {
		http.ListenAndServeTLS(":"+config.Bind, config.Certfile, config.Keyfile, nil)
	} else {
		fmt.Fprintf(os.Stderr, "unknown mode or mode not set")
		os.Exit(1)
	}
}

//Function for determining what the next backend host should be
func pickHost(lastHost, totesHost int) int {
	x := lastHost + 1
	if x >= totesHost {
		return 0
	}
	return x
}
