package config

var Config = make(map[string]string)

// Exportable struct
// Capitalize the type and the fields to make it exportable
type AppConf struct {
	Debugging     string `json:"debugging,omitempty"`
	Directory     string `json:"directory,omitempty"`
	MgoURL        string `json:"mgoURL,omitempty"`
	Port          string `json:"port,omitempty"`
	SmtpServerUrl string `json:"smtp-server-url,omitempty"`
	SmtpPort      string `json:"smtp-port,omitempty"`
	SmtpUsername  string `json:"smtp-username,omitempty"`
	SmtpPassword  string `json:"smtp-password,omitempty"`
	Smtp          map[string]string
}
