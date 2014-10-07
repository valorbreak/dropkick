package core

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"github.com/valorbreak/dropkick/routes"
	mgo "gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

// Exportable struct
// Capitalize the type and the fields to make it exportable
type AppConf struct{
	Port string
	Dir string
	MgoURL string
	Debugging string
}

func AppStart(conf AppConf) error{
	// Logging
	log.Printf("Initialize Core\n")
	
	// This is a GoRoutine
	// go accepts a func that runs concurrently
	go connectMgo(conf.MgoURL)

	routeConf := routes.AppConf{conf.Port, conf.Dir, conf.MgoURL, conf.Debugging}

	r := routes.GetRouter(routeConf)
	// Handle Mux
	http.Handle("/",r)
	// http.Handle("/",fileHandler)

	f, err := os.OpenFile("logs/messages.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
	    log.Fatal("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Printf("Resources are now available")
	log.Printf("Directory %s", conf.Dir)
	// Start Listening on port
	port := ":" + conf.Port
	err2 := http.ListenAndServe(port, r)
	if err2 != nil {
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
