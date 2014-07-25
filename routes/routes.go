package routes

import (
	"net/http"
	"github.com/gorilla/mux"
)

// Exportable struct
// Capitalize the type and the fields to make it exportable
type AppConf struct{
	Port string
	Dir string
	MgoURL string
}

func GetRouter(conf AppConf) *mux.Router{
	// Set the Urls here
	fs := http.Dir(conf.Dir)
	fileHandler := http.FileServer(fs)

	// Gorilla Web Toolkit
	// Mux - HTTP Request multiplexer
	r := mux.NewRouter()
	
	// Trailing slash Redirects to non-trailing slash
	r.StrictSlash(true)
	
	// Set Routing Here
	r.HandleFunc("/api/view", viewHandler)
	r.HandleFunc("/api/json", jsonHandler)
	r.HandleFunc("/api/user", userHandler)
	r.HandleFunc("/api/user/edit", userEditHandler)
	
	//r.HandleFunc("/", jsonHandler)

	// Static Files are handled here
	// If the Request URL didn't match any routes
	// default to static folder
	// Order for which this function is declared is importent
	r.PathPrefix("/").Handler(fileHandler)
	return r
}

