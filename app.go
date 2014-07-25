package main

import (
	"fmt"
	"github.com/valorbreak/dropkick/core"
	"github.com/valorbreak/dropkick/model"
	"github.com/valorbreak/dropkick/routes"
	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	"log"
	"flag"
)

// this is the main function
type dkConf struct{
	Port string
	MgoURL string
}

func main() {

	// Command Line Flags 
	mgoURL := flag.String("mongo","mongo://localhost:27017","Mongo DB URL")
	port := flag.String("port","8080","Default port is set to 8080")
	dir := flag.String("directory", "web/", "directory of files")
	flag.Parse()

	conf := dkConf{*port, *mgoURL}		
	routeConfig := routes.AppConf{*port, *dir}
	
	defer routes.AppStart(routeConfig)

	x := model.Bad()
	core.Execute()
	fmt.Println(x)
	
	// Logging 
	log.Printf("About to listen on " + conf.Port)

	go func(url string) {
		session, err := mgo.Dial(url)	
		log.Printf("Successfully connected to Mongo")	
		if(err != nil){
			c := session.DB("test").C("goapp")
			fmt.Println(c)
			//test := c.Find(bson.M{"name": "Bob"})
			//fmt.Println(test)
		}
	}(conf.MgoURL)


}
// phpstorm commit
