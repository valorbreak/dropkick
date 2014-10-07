package routes

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
	"github.com/gorilla/mux"
	"github.com/valorbreak/dropkick/model"
	"encoding/json"
)

var frontPage = template.Must(template.ParseFiles(
	"web/sites/themes/ice/layout.tmpl",
	"web/sites/themes/ice/index.tmpl",
	"web/sites/themes/ice/header.tmpl",
))

func adminHandler(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	body := vars["args"]

	content := model.Post{
		Page: model.Page{
			Title: "test",
			Body: body,
		},
		Date: "102304",
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
		log.Printf("Debug: %s", content.Title)
		frontPage.Execute(w, content)
	}
}
