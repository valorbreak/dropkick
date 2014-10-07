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
	vars := mux.Vars(r)
	body := vars["args"]

	content := Page{
		Title: "Golang",
		Body: body,
	}

	if(".json" == vars["ext"]){
		b, err := json.Marshal(content)
		if(err != nil){
			return
		}
		debugMessage(r)
		stringB := string(b)
		fmt.Fprint(w,stringB)
	} else{
		debugMessage(r)
		lp := coreAppConf.Dir + "/sites/themes/ice/layout.html";
		fp := coreAppConf.Dir + "/sites/themes/ice/index.html";
		hp := coreAppConf.Dir + "/sites/themes/ice/header.html";

		t, err := template.ParseFiles(lp,fp,hp)
		if err != nil {
			log.Printf("File not found: %s", err)
			return
		}
		t.Execute(w, content)
	}
}
