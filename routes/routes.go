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
	Debugging string
}

// make the configuration available for handlers
var coreAppConf AppConf

/*
 * Since we are returning the router, we could replace this with any mux routing libraries	
 */
func GetRouter(conf AppConf) *mux.Router{

	// Set this variable globally for the package
	coreAppConf = conf;

	// Set the Urls here
	fs := http.Dir(conf.Dir)
	fileHandler := http.FileServer(fs)

	// Gorilla Web Toolkit
	// Mux - HTTP Request multiplexer
	r := mux.NewRouter()
	
	// Trailing slash Redirects to non-trailing slash
	r.StrictSlash(true)
	
	// Set Routing Here
	apiRoute := r.PathPrefix("/api").Subrouter()
	apiRoute.HandleFunc("/", apiHandler)
	apiRoute.HandleFunc("/{key}/", apiHandler)
	apiRoute.HandleFunc("/{key}/details", apiHandler)

	r.HandleFunc("/", adminHandler)
	r.HandleFunc("/admin", adminHandler)
	r.HandleFunc("/admin{ext}", adminHandler)

	entityRoute := r.PathPrefix("/entity").Subrouter()
	entityRoute.HandleFunc("/", entityTypeHandler)
	entityRoute.HandleFunc("/{type}/", entityTypeHandler)
	entityRoute.HandleFunc("/{type}/{nid}", entityTypeHandler)

	//r.HandleFunc("/", jsonHandler)

	// Static Files are handled here
	// If the Request URL didn't match any routes
	// default to static folder
	// Order for which this function is declared is important
	r.PathPrefix("/").Handler(fileHandler)

	return r
}

