package routes

import (
	"github.com/gorilla/mux"
	"github.com/valorbreak/dropkick/core/config"
	"net/http"
	"html/template"
	"log"
)

// Make the configuration available for handlers.
var coreAppConf config.AppConf

var contactPage *template.Template
var frontPage *template.Template
//var templates map[*template.Template]

func SetConfig(conf config.AppConf){
	// Set this variable globally for the package
	coreAppConf = conf
	setTemplate(coreAppConf.Directory)
}

/*
 Defines the templates that can be available on the site
*/
func setTemplate(directory string){
	contactPage = template.Must(template.ParseFiles(
			directory + "/sites/themes/ice/layout.tmpl",
			directory + "/sites/themes/ice/templates/contact.tmpl",
			directory + "/sites/themes/ice/header.tmpl",
	))
	frontPage = template.Must(template.ParseFiles(
			directory + "/sites/themes/ice/layout.tmpl",
			directory + "/sites/themes/ice/index.tmpl",
			directory + "/sites/themes/ice/header.tmpl",
	))
	log.Printf("%s",coreAppConf)
}

/*
 Since we are returning the router, we could replace this with any mux routing libraries
 */
func GetRouter() *mux.Router {

	// Gorilla Web Toolkit
	// Mux - HTTP Request multiplexer
	router := mux.NewRouter()

	// Router configuration
	// Trailing slash Redirects to non-trailing slash
	router.StrictSlash(true)

	// Set Routing Here
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/admin", adminHandler)
	router.HandleFunc("/admin{ext}", adminHandler)
	router.HandleFunc("/admin/{args}", adminHandler)

	// API Routes
	apiRoute := router.PathPrefix("/api").Subrouter()
	apiRoute.HandleFunc("/", apiHandler)

	// Entity Routes
	entityRoute := router.PathPrefix("/entity").Subrouter()
	entityRoute.HandleFunc("/", entityTypeHandler)
	entityRoute.HandleFunc("/{type}/", entityTypeHandler)
	entityRoute.HandleFunc("/{type}/{nid}", entityTypeHandler)

	// Contact Routes
	contactRoute := router.PathPrefix("/contact").Subrouter()
	contactRoute.HandleFunc("/", contactHandler).Methods("GET")
	contactRoute.HandleFunc("/", contactHandlerPost).Methods("POST")

	/*
		Static Files are handled here
		If the Request URL didn't match any routes
		default to static folder
		Order for which this function is declared is important
		Set the default front page here
	*/

	// /web is the public folder with limitation to several file extensions
	fs := http.Dir(coreAppConf.Directory)
	fileHandler := http.FileServer(fs)

	// Include these files for static access
	router.PathPrefix("/{anything:.+}.{ext:(css|html|js|png|jpg|woff|eot|svg|tff)}").Handler(fileHandler)

	return router
}
