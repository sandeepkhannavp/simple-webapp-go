package render

import (
	//"errors"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{
	 
}

//render template using html template
func RenderTemplate (w http.ResponseWriter, tmpl string){
	tc,err := CreateTemplateCache()
	if err!=nil{
		log.Fatal(err)
	}
	t,ok := tc[tmpl]
	if !ok{
		log.Fatal(err)
	}

	buf:=new(bytes.Buffer)

	_= t.Execute(buf,nil)

	_ ,err = buf.WriteTo(w)
	if err!=nil{
		fmt.Println("error writing template to browser",err)
	}
	parsedTemplate,_ := template.ParseFiles("./templates/"+tmpl)
	err = parsedTemplate.Execute(w,nil)
	if err!=nil{
		fmt.Println("error while parsing templates")
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