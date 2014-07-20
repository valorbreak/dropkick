package routes

import (
	"fmt"
	"net/http"
	"log"
)


func AppStart(port string) error{
	// Logging
	fmt.Printf("Hey Route is working\n")
	
	// Set the Urls here
	http.HandleFunc("/", WelcomeHandler)
	http.HandleFunc("/json", JsonHandler)
	http.HandleFunc("/view", ViewHandler)

	// Start Listening on port
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return err;
}


