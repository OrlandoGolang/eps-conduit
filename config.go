package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"strings"
	"net/http"
	"net/http/httputil"
)

type Config struct {
	Backends []string `toml:"backends"`
	Bind     string   `toml:"bind"`
	Mode     string   `toml:"mode"`
	Certfile string   `toml:certFile`
	Keyfile  string   `toml:keyFile`
}

var config *Config = nil

func GetConfig(configFile string) *Config {
	if config == nil {
		config = new(Config)
		config.init(configFile)
	}
	return config
}

func (c *Config) init(configFile string) {
	_, err := os.Stat(configFile)

	if err != nil {
		log.Fatal("Config file not found: ", configFile)
	}
	if _, err := toml.DecodeFile(configFile, c); err != nil {
		log.Fatal(err)
	}
}

func (c *Config) handleUserInput() {

	if *backendStr != "" {
		//Remove whitespace from backends
		*backendStr = strings.Replace(*backendStr, " ", "", -1)
		//Throwing backends into an array
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

func (c *Config) printConfigInfo() {
	// tell the user what info the load balancer is using
	for _, v := range c.Backends {
		log.Println("using " + v + " as a backend")
	}
	log.Println("listening on port " + c.Bind)
}

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
