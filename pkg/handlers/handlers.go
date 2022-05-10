package handlers

import (
	"net/http"

	"github.com/sandeepkhannavp/htmltest/pkg/config"
	"github.com/sandeepkhannavp/htmltest/pkg/models"
	"github.com/sandeepkhannavp/htmltest/pkg/render"
)


var Repo *Respository //call the associated functions using the repo which is pointer to the structure -- see main

//handlers has it own pointer to application wide config AppConfig which is implemented as struct so it can have associated functions
type Respository struct{
	App *config.AppConfig
}


func NewRepo(a *config.AppConfig) *Respository{
	return &Respository{
		 App:a,
	}
}

//new handlers sets the repositpry for the handlers
func NewHandlers(r *Respository){
	Repo = r
}

//Home is the home page handler
func (m *Respository)Home(w http.ResponseWriter,r *http.Request){
	render.RenderTemplate(w,"home.page.html",&models.TemplateData{})
}

//About is the about page handler
func (m *Respository)About(w http.ResponseWriter , r *http.Request){
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello"
	render.RenderTemplate(w,"about.page.html",&models.TemplateData{
		StringMap: stringMap,
	})
}
//render the templates
