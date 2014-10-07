package model

/**
 *	Private Functions always have the first letter lowercase
 *	These functions are only viewable inside the package
 *
 *  Status Codes are in the net/http package
 *  http://golang.org/src/pkg/net/http/status.go
 */
type Page struct {
	Header string
	Title string
	Body string
	Type string
	Id string
	Info string
	Debug string
}

type Post struct {
	Page
	Date string
}
