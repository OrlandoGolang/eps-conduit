package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

// Config Struct represents the load balancer's configuration
type Config struct {

	// the backend services to balance
	Backends []string `toml:"backends"`

	// The port the load balancer is bound to
	Bind     string   `toml:"bind"`

	// The load balancing mode (i.e. round-robin)
	// Currently, only round-robin is used, and configurable mode is not supported.
	Mode     string   `toml:"mode"`

	// Path to certificate file
	Certfile string   `toml:certFile`

	// Path to private key file related to certificate
	Keyfile  string   `toml:keyFile`
}

// singleton Config instance initially set to nil
var config *Config = nil

// GetConfig implements a singleton pattern for access the Config instance.
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
}

// handleUserInput checks command line input and overrides config file settings
// Backends is parsed from a raw string to a slice of strings
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
func (c *Config) makeProxies() []*httputil.ReverseProxy {
	// Create a proxy for each backend
	proxies := make([]*httputil.ReverseProxy, len(c.Backends))
	for i := range c.Backends {
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = c.Backends[i]
		}
		proxies[i] = &httputil.ReverseProxy{Director: director}
	}
	return proxies
}
