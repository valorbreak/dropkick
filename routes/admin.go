package routes

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
	"github.com/gorilla/mux"
	"encoding/json"
)

func adminHandler(w http.ResponseWriter,r *http.Request){
	debugMessage(r)
	vars := mux.Vars(r)
	body := vars["args"]


	content := Page{
		Title: "Golang",
		Body: body,
	}
	t, err := template.ParseFiles(coreAppConf.Dir + "/sites/themes/ice/index.html")

	if err != nil{
		log.Printf("File not found: %s", err)
		return
	}

	if(".json" == vars["ext"]){
		b, err := json.Marshal(content)
		if(err != nil){
			return
		}
		log.Printf("Accessed: %s", b)
		stringB := string(b)
		fmt.Fprint(w,stringB)
	} else{
		log.Printf("Accessed: %s", r)
		t.Execute(w, content)
	}
}
