package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

//TemplateData sends data from handlers to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	Data      map[string]interface{}
	CSRFToken string
	FlashMsg  string
	Warning   string
	Error     string
}

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	IsProd        bool
	Session       *scs.SessionManager
}
