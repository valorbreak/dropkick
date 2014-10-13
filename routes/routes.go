package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/valorbreak/dropkick/core/config"
)

// Make the configuration available for handlers.
var coreAppConf config.AppConf

/**
 * Since we are returning the router, we could replace this with any mux routing libraries	
 */
func GetRouter(conf config.AppConf) *mux.Router{

	// Set this variable globally for the package
	coreAppConf = conf;

	// Set the Urls here
	fs := http.Dir(conf.Directory)
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

	r.HandleFunc("/admin", adminHandler)
	r.HandleFunc("/admin{ext}", adminHandler)
	r.HandleFunc("/admin/{args}", adminHandler)

	entityRoute := r.PathPrefix("/entity").Subrouter()
	entityRoute.HandleFunc("/", entityTypeHandler)
	entityRoute.HandleFunc("/{type}/", entityTypeHandler)
	entityRoute.HandleFunc("/{type}/{nid}", entityTypeHandler)

	contactRoute := r.PathPrefix("/contact").Subrouter()
	contactRoute.HandleFunc("/", contactHandler).Methods("GET")
	contactRoute.HandleFunc("/", contactHandlerPost).Methods("POST")

	/**
	 * Static Files are handled here
	 * If the Request URL didn't match any routes
	 * default to static folder
	 * Order for which this function is declared is important
	 * Set the default front page here
	 */
	r.HandleFunc("/", indexHandler)
	r.PathPrefix("/").Handler(fileHandler)

	return r
}

