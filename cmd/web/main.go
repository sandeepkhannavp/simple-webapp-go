package main

import (
	//"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/sandeepkhannavp/htmltest/pkg/config"
	"github.com/sandeepkhannavp/htmltest/pkg/handlers"
	"github.com/sandeepkhannavp/htmltest/pkg/render"
)

const portNumber = ":8080"


func main(){
	var app config.AppConfig
	tc ,err := render.CreateTemplateCache()
	if err!=nil{
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache=tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	//this handler function listens for requests from a web browser
	http.HandleFunc("/",handlers.Repo.Home)
	http.HandleFunc("/about",handlers.Repo.About)
	//start a web server that listens for requests - listen to the port 8080
	fmt.Println(fmt.Sprintf("Starting Application on port %s",portNumber))
	_ = http.ListenAndServe(portNumber,nil)
}