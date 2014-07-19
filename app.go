package main

import (
	"fmt"
	"net/http"
	//	"html/template"
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
	routes.Execute()

	fmt.Println(x)
	log.Printf("About to listen on " + port)

	welcome := routes.WelcomeHandler
	view := routes.ViewHandler
	json := routes.JsonHandler
	http.HandleFunc("/foo", view)
	http.HandleFunc("/", welcome)
	http.HandleFunc("/api", json)

	// Start Listening to port
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
