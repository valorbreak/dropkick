package main

import (
	"flag"
	"github.com/valorbreak/dropkick/core"
	"log"
)

func main() {

	// Command Line Flags
	// $ dropkick -h
	mgoURL := flag.String("mongo", "mongodb://localhost:27017", "Mongo DB URL")
	port := flag.String("port", "8080", "Default port is set to 8080")
	dir := flag.String("directory", "web", "directory of files")
	debugging := flag.String("debug", "0", "Enable Debugging")

	flag.Parse()

	// Logging
	log.Printf("Static File Directory: " + *dir)
	log.Printf("About to listen on " + *port)

	// Start the application
	routeConfig := core.AppConf{*port, *dir, *mgoURL, *debugging}
	core.AppStart(routeConfig)

}
