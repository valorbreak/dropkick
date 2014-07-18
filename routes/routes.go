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
        fmt.Fprintf(w,"Hi there, I love %s", r.URL.Path[1:])
}
