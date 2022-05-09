package main

import (
	//"errors"
	"fmt"
	"net/http"
	"github.com/sandeepkhannavp/htmltest/pkg/handlers"
)

const portNumber = ":8080"


func main(){
	//this handler function listens for requests from a web browser
	http.HandleFunc("/",handlers.Home)
	http.HandleFunc("/about",handlers.About)
	//start a web server that listens for requests - listen to the port 8080
	fmt.Println(fmt.Sprintf("Starting Application on port %s",portNumber))
	_ = http.ListenAndServe(portNumber,nil)
}