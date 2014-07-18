package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/valorbreak/dropkick/model"
	"github.com/valorbreak/dropkick/core"
	"github.com/valorbreak/dropkick/routes"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hi there, I love %s", r.URL.Path[1:])
}

func main() {
	fmt.Println("Hello, test")
	port := ":8080"
	
	x := model.Bad()
	core.Execute()
	routes.Execute()
	
	fmt.Println(x)
	log.Printf("About to listen on "+port)

	welcome := routes.WelcomeHandler
	view := routes.ViewHandler
	http.HandleFunc("/foo", view)
	http.HandleFunc("/", welcome)

	// Start Listening to port
	err := http.ListenAndServe(port, nil)
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}
