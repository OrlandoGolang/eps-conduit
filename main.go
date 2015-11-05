package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"strings"
)

//Global variable of the next backend to be sent.  This is for round-robin load balancing
var nextHost int = 0

func main() {

	var config = ReadConfig()

	//Handling user flags
	backend := flag.String("b", config.Backends, "target ips for backend servers")
	bind := flag.String("bind", config.Bind, "port to bind to")
	mode := flag.String("mode", config.Mode, "Balancing Mode")
	certFile := flag.String("cert", config.Certfile, "cert")
	keyFile := flag.String("key", config.Keyfile, "key")
	flag.Parse()

	//Error out if backend is not set
	if *backend == "" {
		fmt.Fprintf(os.Stderr, "no backends chosen, use the -b flag.  ex:\n")
		fmt.Fprintf(os.Stderr, "eps-conduit -b 10.1.8.1,10.1.8.2")
		os.Exit(1)
	}

	//Remove whitespace from backends
	strings.Replace(*backend, " ", "", -1)
	//Throwing backends into an array
	backends := strings.Split(*backend, ",")
	totesHost := len(backends)

	//Create a proxy for each backend
	proxies := make([]*httputil.ReverseProxy, totesHost)
	for i := range backends {
		host := backends[i]
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = host
		}
		proxies[i] = &httputil.ReverseProxy{Director: director}
	}

	//Function for handling the http requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		go proxies[nextHost].ServeHTTP(w, r)
		nextHost = pickHost(nextHost, totesHost)
		log.Println("nextHost is " + strconv.Itoa(nextHost))
	})

	//tell the user what info the load balancer is using
	for _, v := range backends {
		log.Println("using " + v + " as a backend")
	}
	log.Println("listening on port " + *bind)

	//Start the http(s) listener depending on user's selected mode
	if *mode == "http" {
		http.ListenAndServe(":"+*bind, nil)
	} else if *mode == "https" {
		http.ListenAndServeTLS(":"+*bind, *certFile, *keyFile, nil)
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
