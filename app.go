package main

import (
	"fmt"
	"github.com/valorbreak/dropkick/core"
	"github.com/valorbreak/dropkick/model"
	"flag"
	"log"
)

func main() {

	// Command Line Flags
	// $ dropkick -h
	mgoURL := flag.String("mongo", "mongodb://localhost:27017", "Mongo DB URL")
	port := flag.String("port", "8080", "Default port is set to 8080")
	dir := flag.String("directory", "web", "directory of files")
	debugging := flag.String("debug", "0", "Enable Debugging")
	//dir := flag.String("directory", "web/static", "directory of files")
	flag.Parse()

	log.Printf("Static File Directory: "+*dir)

	routeConfig := core.AppConf{*port, *dir, *mgoURL, *debugging}
	// defer - runs the specified function before (this) function ends
	defer core.AppStart(routeConfig)
	
	x := model.Bad()
	//core.Execute()
	fmt.Println(x)

	// Logging
	log.Printf("About to listen on " + *port)
	// This is where the Deferred function executes

}
