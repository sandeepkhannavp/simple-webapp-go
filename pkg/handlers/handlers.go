package handlers

import (
	"net/http"

	"github.com/sandeepkhannavp/htmltest/pkg/config"
	"github.com/sandeepkhannavp/htmltest/pkg/models"
	"github.com/sandeepkhannavp/htmltest/pkg/render"
)

//Template data holds data sent from handlers to Templates

//repository used by the handlers
var Repo *Respository
//
type Respository struct{
	App *config.AppConfig

}

//creates a new repository and return it
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
