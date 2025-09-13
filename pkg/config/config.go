package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// ApplicationConfig holds the application wide configuration
type ApplicationConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InfoLog       *log.Logger
	IsProd        bool
	Session       *scs.SessionManager
}
