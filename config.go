package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

// Config Struct represents the load balancer's configuration
type Config struct {

	// the backend services to balance
	Backends []string `toml:"backends"`

	// The port the load balancer is bound to
	Bind string `toml:"bind"`

	// Secure or unsecure http protocol
	Mode string `toml:"mode"`

	// Path to certificate file
	Certfile string `toml:certFile`

	// Path to private key file related to certificate
	Keyfile string `toml:keyFile`

	// Path to access logs
	AccessLog string `toml:accessLog`

	// Revserse Proxies to forward requests to
	Proxies []*httputil.ReverseProxy

	// Number of proxies available
	HostCount int

	// The index of the next proxy to forward a request to
	NextHost int
}

// singleton Config instance initially set to nil
var config *Config = nil

// GetConfig implements a singleton pattern to access the Config singleton
func GetConfig(configFile string) *Config {
	if config == nil {
		config = new(Config)
		config.init(configFile)
	}
	return config
}

// init initializes a new Config instance by reading from the config file
// It will unmarshal the toml file into the Config struct
func (c *Config) init(configFile string) {
	_, err := os.Stat(configFile)

	if err != nil {
		log.Fatal("Config file not found: ", configFile)
	}

	if _, err := toml.DecodeFile(configFile, c); err != nil {
		log.Fatal(err)
	}

	c.handleUserInput()
	c.printConfigInfo()
	c.makeProxies()
	c.HostCount = len(c.Backends)
	c.NextHost = 0

	if err := c.configureAccessLog(); err != nil {
		log.Fatal(err)
	}
}

// handleUserInput checks command line input and overrides config file settings
// Backends is parsed from a raw string to a slice of strings
// TODO: Better input validation
func (c *Config) handleUserInput() {
	if *backendStr != "" {
		// Remove whitespace from backends
		*backendStr = strings.Replace(*backendStr, " ", "", -1)
		// Throwing backends into an array
		c.Backends = strings.Split(*backendStr, ",")
	}
	if *bind != "" {
		c.Bind = *bind
	}
	if *mode != "" {
		c.Mode = *mode
	}
	if *certFile != "" {
		c.Certfile = *certFile
	}
	if *keyFile != "" {
		c.Keyfile = *keyFile
	}
	if *accessLog != "" {
		c.AccessLog = *accessLog
	}
}

// printConfigInfo prints to stderr host and port settings applied to current process
func (c *Config) printConfigInfo() {
	// tell the user what info the load balancer is using
	for _, v := range c.Backends {
		log.Println("using " + v + " as a backend")
	}
	log.Println("listening on port " + c.Bind)
}

// makeProxies creates slice of ReverseProxies based on the Config's backend hosts
// It returns a slice of httputil.ReverseProxy
func (c *Config) makeProxies() {
	// Create a proxy for each backend
	c.Proxies = make([]*httputil.ReverseProxy, len(c.Backends))
	for i := range c.Backends {
		// host must be defined here, and not within the anonymous function.
		// Otherwise, you'll run into scoping issues
		host := c.Backends[i]
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = host
		}
		c.Proxies[i] = &httputil.ReverseProxy{Director: director}
	}
}

// configureAccessLog checks if the log path exists and if not, sets it up so
// logging can successfully take place
func (c *Config) configureAccessLog() error {
	if _, err := os.Stat(config.AccessLog); os.IsNotExist(err) {
		file, err := os.OpenFile(config.AccessLog, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}

// Function for handling the http requests
func (c *Config) handle(w http.ResponseWriter, r *http.Request) {
	c.pickHost()
	c.Proxies[c.NextHost].ServeHTTP(w, r)
}

// pickHost determines the next backend host to forward the request to - according to round-robin
// It returns an integer, which represents the host's index in config.Backends
func (c *Config) pickHost() {
	nextHost := c.NextHost + 1
	if nextHost >= c.HostCount {
		c.NextHost = 0
	} else {
		c.NextHost = nextHost
	}
}
