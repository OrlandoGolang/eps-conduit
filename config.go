package main


import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type Config struct {
	Backends []string `toml:"backends"`
	Bind     string `toml:"bind"`
	Mode     string `toml:"mode"`
	Certfile string `toml:certFile`
	Keyfile  string `toml:keyFile`
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

	log.Println("Line 30")
	if err != nil {
		log.Fatal("Config file not found: ", configFile)
	}
	log.Println("Line 34")
	if _, err := toml.DecodeFile(configFile, c); err != nil {
		log.Fatal(err)
	log.Println("Line 37")
	}
}
