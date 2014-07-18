package routes

import (
	"fmt"
	"net/http"
)

func Execute() {
	fmt.Printf("Hey Route is working\n")
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w,"Welcome %s", r.URL.Path[1:])
}

func ViewHandler(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w,"<h1>Hi there, I love %s</h1>", r.URL.Path[1:])
}

func JsonHandler(w http.ResponseWriter, r *http.Request){
	const jsonStream = `
		{"json":"test"}
	`

	fmt.Fprint(w,jsonStream)
}
