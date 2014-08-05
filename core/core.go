package core

import (
	"fmt"
	"net/http"
	"log"
	"github.com/valorbreak/dropkick/routes"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

// Exportable struct
// Capitalize the type and the fields to make it exportable
type AppConf struct{
	Port string
	Dir string
	MgoURL string
}

func AppStart(conf AppConf) error{
	// Logging
	log.Printf("Hey Core is working\n")
	
	// This is a GoRoutine
	// go accepts a func that runs concurrently
	go connectMgo(conf.MgoURL)

	routeConf := routes.AppConf{conf.Port, conf.Dir, conf.MgoURL}
	routes.GetRouter(routeConf)

	
	r := routes.GetRouter(routeConf)
	// Handle Mux
	http.Handle("/",r)
	// http.Handle("/",fileHandler)

	log.Printf("Resources are now available")
	
	// Start Listening on port
	port := ":" + conf.Port
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return err;
}

func connectMgo(url string) {
	session, err := mgo.Dial(url)
	if err != nil {
		log.Printf("Cannot connect to MongoDB: %s,\n err: %s", url, err)
		c := session.DB("dropkick").C("users")
		fmt.Println(c)
		//test := c.Find(bson.M{"name": "Bob"})
		//fmt.Println(test)
	} else {
		log.Printf("Successfully connected to MongoDB:")
	}
}
