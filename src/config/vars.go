package config

import (
	"os"
)

var Env string

var Port string
var SiteHost string

func init() {
	// get ENV key from .env file and if not found then use "development" as default
	Env = os.Getenv("ENV")
	if Env == "" {
		Env = "development"
	}
	SiteHost = os.Getenv("SITE_HOST")
	Port = os.Getenv("PORT")
}
