package main

import (
	"flag"
	"github.com/valorbreak/dropkick/core"
	"github.com/valorbreak/dropkick/core/config"
	"log"
	"io/ioutil"
	"encoding/json"
)

func main() {
	// Command Line Flags
	// $ dropkick -h
	var config = config.AppConf{}
	config.MgoURL = *flag.String("mongo", "mongodb://localhost:27017", "Mongo DB URL")
	config.Port = *flag.String("port", "8080", "Default port is set to 8080")
	config.Directory = *flag.String("directory", "web", "directory of files")
	config.Debugging = *flag.String("debug", "0", "Enable Debugging")
	flag.Parse()

	// Logging
	//log.Printf("Static File Directory: " + config.Directory)
	//log.Printf("About to listen on " + config.Port)

	// Read config.json for site specific configuration
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Printf("Config File Not Found: %s",err)
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil{
		log.Printf("Error Reading the JSON file: %s",err)
	}

	log.Printf("Listening on Port: %s", config.Port)
	log.Printf("Program Directory: %s", config.Directory)
	log.Printf("Debugging Setting: %s", config.Debugging)
	log.Printf("MongoDB URL: %s", config.MgoURL)

	core.AppStart(config)
}
