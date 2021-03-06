package config

import (
	"log"
	"text/template"
)

//appconfig holds the application configurations
type AppConfig struct{
	TemplateCache map[string]*template.Template
	UseCache bool
	InfoLog *log.Logger
	
}