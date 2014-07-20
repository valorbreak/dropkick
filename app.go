package main

import (
	"fmt"
	"github.com/valorbreak/dropkick/core"
	"github.com/valorbreak/dropkick/model"
	"github.com/valorbreak/dropkick/routes"
	"log"
)

// this is the main function

func main() {
	fmt.Println("Hello, test")
	port := ":8080"

	x := model.Bad()
	core.Execute()
	fmt.Println(x)
	

	// Logging 
	log.Printf("About to listen on " + port)
	
	routes.AppStart(port)


}
// phpstorm commit
