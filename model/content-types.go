package model

/**
 *	Private Functions always have the first letter lowercase
 *	These functions are only viewable inside the package
 *
 *  Status Codes are in the net/http package
 *  http://golang.org/src/pkg/net/http/status.go
 */
type Page struct {
	Header string `json:"header,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
	Type   string `json:"type,omitempty"`
	Id     string `json:"id,omitempty"`
	Info   string `json:"info,omitempty"`
	Debug  string `json:"debug,omitempty"`
}

type Post struct {
	Pagego
	Date string `json:"header"`
}

type entity struct {
	Uuid string
}
