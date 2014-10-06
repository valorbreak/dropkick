package routes


import (
	"net/http"
	"log"
	"html/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request){
	debugMessage(r)
	content := Page{
		Title: "Golang",
		Body: "content",
	}
	t, err := template.ParseFiles(coreAppConf.Dir + "/sites/themes/flame/index.html")
	if err != nil {
		log.Printf("File not found: %s", err)
		return
	}
	t.Execute(w, content)
}
