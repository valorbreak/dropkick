package routes

import (
//	"fmt"
	"net/http"
	"log"
//	"flag"
	"github.com/gorilla/mux"
)

type AppConf struct{
	Port string
	Dir string
}

func AppStart(conf AppConf) error{
	// Logging
	log.Printf("Hey Route is working\n")
		
	// Set the Urls here
	fs := http.Dir(conf.Dir)
	fileHandler := http.FileServer(fs)

	/*
	app := http.NewServeMux()

	app.HandleFunc("/view", ViewHandler)
	app.HandleFunc("/json", JsonHandler)
	app.HandleFunc("/new/test", JsonHandler)
	app.HandleFunc("/file", fileHandler)
	
	http.Handle("/", app)
	http.Handle("/static", http.StripPrefix("/static/",fileHandler))
	*/

	// Gorilla Web Toolkit
	// Mux - HTTP Request multiplexer
	r := mux.NewRouter()
	
	// Trailing slash Redirects to non-trailing slash
	r.StrictSlash(true) 	
	
	// Set Routing Here
	r.HandleFunc("/api/view", ViewHandler)
	r.HandleFunc("/api/json", JsonHandler)
	r.HandleFunc("/api/user", UserHandler)
	r.HandleFunc("/api/user/edit", userEditHandler)
	r.HandleFunc("/", JsonHandler)

	// Static Files are handled here
	// If the Request URL didn't match any routes
	// default to static folder
	// Order for which this function is declared is importent
	// 
	r.PathPrefix("/").Handler(fileHandler)

	
	// Handle Mux
	http.Handle("/",r)
	//http.Handle("/",fileHandler)

	log.Printf("Resources are now available")
	
	// Start Listening on port
	port := ":" + conf.Port
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return err;
}


