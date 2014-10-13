package routes

import (
	"fmt"
	"net/http"
	"net/smtp"
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


var contactPage = template.Must(template.ParseFiles(
	"web/sites/themes/ice/layout.tmpl",
	"web/sites/themes/ice/templates/contact.tmpl",
	"web/sites/themes/ice/header.tmpl",
))

func contactHandler(w http.ResponseWriter,r *http.Request){
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
		contactPage.Execute(w, content)
	}
}


func contactHandlerPost(w http.ResponseWriter,r *http.Request){
	// Golang Gotcha, run ParseForm() before using r.Form.Get()
	err := r.ParseForm()
	if err != nil {
		log.Printf("Form is invalid: %s",err)

	}

	email := r.Form.Get("InputEmail")
	inputBody := r.Form.Get("InputMessage")

	body := fmt.Sprintf("Reply-To: %v\r\nSubject: Afterofficeit - No Reply: \r\n\n%s", email,inputBody)
	to := []string{email}
	username := coreAppConf.SmtpUsername
	password := coreAppConf.SmtpPassword
	auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")
	smtpServerUrl := coreAppConf.SmtpServerUrl+":"+coreAppConf.SmtpPort
	smtp.SendMail(smtpServerUrl, auth, email, to, []byte(body))
}
