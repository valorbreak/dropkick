package routes

import (
	"fmt"
	"net/http"
	"log"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request){
	log.Printf("Accessed: /%s", r.URL.Path[1:])
	fmt.Fprintf(w,"Welcome %s", r.URL.Path[1:])
}

func ViewHandler(w http.ResponseWriter, r *http.Request){
	log.Printf("Accessed: /%s", r.URL.Path[1:])
	fmt.Fprintf(w,"<h1>Hi there, I love %s</h1>", r.URL.Path[1:])
}

func JsonHandler(w http.ResponseWriter, r *http.Request){
	const jsonStream = `
		{"json":"test"}
	`
	log.Printf("Accessed: /%s", r.URL.Path[1:])
	fmt.Fprint(w,jsonStream)
}