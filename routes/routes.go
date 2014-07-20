package routes

import (
	"fmt"
	"net/http"
	"log"
	"flag"
)

func AppStart(port string) error{
	// Logging
	fmt.Printf("Hey Route is working\n")
	
	// Set the Urls here
	dir := flag.String("directory", "web/", "directory of files")
	flag.Parse()

	fs := http.Dir(*dir)
	fileHandler := http.FileServer(fs)

	http.Handle("/",fileHandler)
	http.Handle("/static/", fileHandler)
	
	//http.HandleFunc("/", WelcomeHandler)
	http.HandleFunc("/view", ViewHandler)
	http.HandleFunc("/json", JsonHandler)
	http.HandleFunc("/new", JsonHandler)
	http.HandleFunc("/new/", JsonHandler)
	
	log.Printf("/, /json, /view is available")
	
	// Start Listening on port
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return err;
}


