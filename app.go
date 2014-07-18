package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/valorbreak/sitgo/model"
	"github.com/valorbreak/sitgo/core"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w,"Welcome %s", r.URL.Path[1:])
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hi there, I love %s", r.URL.Path[1:])
}

func main() {
	fmt.Println("Hello, test")
	port := ":8080"
	
	x := model.Bad()
	core.Execute()
	fmt.Println(x)
	log.Printf("About to listen on "+port)

	http.HandleFunc("/foo",handler)
	http.HandleFunc("/",welcomeHandler)

	// Start Listening to port
	err := http.ListenAndServe(port, nil)
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}
