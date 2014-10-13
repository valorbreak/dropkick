package routes

import (
	"github.com/valorbreak/dropkick/model"
	"html/template"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	debugMessage(r)
	content := model.Page{
		Title: "Golang",
		Body:  "content",
	}
	t, err := template.ParseFiles(coreAppConf.Directory + "/sites/themes/agency/index.html")
	if err != nil {
		log.Printf("File not found: %s", err)
		return
	}
	t.Execute(w, content)
}
