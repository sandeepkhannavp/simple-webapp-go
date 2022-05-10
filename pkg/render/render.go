package render

import (
	//"errors"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/sandeepkhannavp/htmltest/pkg/config"
)

var functions = template.FuncMap{
	 
}
var app *config.AppConfig

//set the config for the template package
func NewTemplates(a *config.AppConfig){
	app=a
}

//render template using html template
func RenderTemplate (w http.ResponseWriter, tmpl string){

	var tc map[string]*template.Template

	//if im in development mode not production - don't use template cache
	//instead rebuild it on every request
	if app.UseCache{
		//get the template cache from app config
		//when we render a page we are pulling this value from the app config
		tc = app.TemplateCache
	}else{
		tc,_ = CreateTemplateCache()
	}

	t,ok := tc[tmpl]
	if !ok{
		log.Fatal("could not get template from template cache")
	}

	buf:=new(bytes.Buffer)

	_= t.Execute(buf,nil)

	_ ,err := buf.WriteTo(w)
	if err!=nil{
		fmt.Println("error writing template to browser",err)
	}

}
//creates the template cache as a map 

func CreateTemplateCache() (map[string]*template.Template,error){

	myCache := map[string]*template.Template{}
	pages,err := filepath.Glob("./templates/*.page.html") //find the pages
	if err!=nil{
		return myCache,err
	}
	for _,page := range pages{
		name:=filepath.Base(page) //it returns the full path - just find the base
		ts,err:= template.New(name).Funcs(functions).ParseFiles(page)
		if err!=nil{
			return myCache,err
		}

		matches ,err := filepath.Glob("./templates/*.layout.html")
		if err!=nil{
			return myCache,err
		}
		if len(matches)>0 {
			ts,err = ts.ParseGlob("./templates/*.layout.html")
			if err!=nil{
				return myCache,err
			}
		}
		myCache[name]=ts
	}

	return myCache,nil
}