package routes

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
)

/**
 *	Private Functions always have the first letter lowercase
 *	These functions are only viewable inside the package
 *
 *  Status Codes are in the net/http package 
 *  http://golang.org/src/pkg/net/http/status.go
 */
type Page struct {
	Header string
	Title string
	Body string
	Type string
	Id string
	Info string
	Debug string
}

func debugMessage(r *http.Request){
	log.Printf("Accessed: %s", r.URL.Path)
	if coreAppConf.Debugging == "1" {
	log.Printf("Debug %s", r)
	}
}

func welcomeHandler(w http.ResponseWriter, r *http.Request){
	debugMessage(r)
	fmt.Fprintf(w,"Welcome %s", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	debugMessage(r)
	fmt.Fprintf(w,"<h1>Hi there, I love %s</h1>", r.URL.Path[1:])
}

func jsonHandler(w http.ResponseWriter, r *http.Request){
	const jsonStream = `
		{"json":"test"}
	`
	debugMessage(r)
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
	debugMessage(r)
	fmt.Fprint(w,jsonStream)
}

func userEditHandler(w http.ResponseWriter, r *http.Request){
	jsonStream := `{"Edit User":"`+ r.URL.Path[1:] +`"}`
	debugMessage(r)
	fmt.Fprint(w,jsonStream)
}

func apiHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	url := vars["url"];
	log.Printf("Accessed: %s", url)
	debugMessage(r)
	fmt.Fprintf(w,"Welcome %s", r.URL.Path[1:])
}

func entityHandler(w http.ResponseWriter, r *http.Request){
	//vars := mux.Vars(r)
	content := Page{
		Title: "Golang",
		Debug: "string",
		Info: "Information about the entity",
		Body: "Body",
	}
	b, err := json.Marshal(content)
	if(err != nil){
		return
	}
	log.Printf("Accessed: %s", b)
	stringB := string(b)
	fmt.Fprint(w,stringB)
}

func entityTypeHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	content := Page{
		Title: "Golang",
		Type: vars["type"],
		Id: vars["id"],
		Body: "Body",
	}
	b, err := json.Marshal(content)
	if(err != nil){
		return
	}
	log.Printf("Accessed: %s", b)
	stringB := string(b)
	fmt.Fprint(w,stringB)
}
