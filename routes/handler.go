package routes

import (
	"fmt"
	http "net/http"
	log "log"
)

/**
 *	Private Functions always have the first letter lowercased
 *	These functions are only viewable inside the package
 *
 *  Status Codes are in the net/http package 
 *  http://golang.org/src/pkg/net/http/status.go
 */

func welcomeHandler(w http.ResponseWriter, r *http.Request){
	log.Printf("Accessed: /%s", r.URL.Path[1:])
	fmt.Fprintf(w,"Welcome %s", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	log.Printf("Accessed: /%s", r.URL.Path[1:])
	fmt.Fprintf(w,"<h1>Hi there, I love %s</h1>", r.URL.Path[1:])
}

func jsonHandler(w http.ResponseWriter, r *http.Request){
	const jsonStream = `
		{"json":"test"}
	`
	log.Printf("Accessed: /%s", r.URL.Path[1:])
	//log.Printf("W: /%s", w)
	
	//resp := &http.Response{
  	//	Status:     "200 OK YES",
  	//	StatusCode: 200,
  	//}
  	//w.Write(resp)
  	w.Header().Add("New Custom Header",r.URL.Path[1:])
  	//w.Header().Set("Access-Control-Allow-Origin", "*")
	// Custom Status Code
	w.WriteHeader(http.StatusCreated)

	fmt.Fprint(w,jsonStream)
}

func userHandler(w http.ResponseWriter, r *http.Request){
	jsonStream := `{"User":"`+ r.URL.Path[1:] +`"}`
	log.Printf("Accessed: /%s", r.URL.Path[1:])
	log.Printf("Debug %s", r)
	fmt.Fprint(w,jsonStream)
}

func userEditHandler(w http.ResponseWriter, r *http.Request){
	jsonStream := `{"Edit User":"`+ r.URL.Path[1:] +`"}`
	log.Printf("Accessed: /%s", r.URL.Path[1:])
	log.Printf("Debug %s", r)
	fmt.Fprint(w,jsonStream)
}


