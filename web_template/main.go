package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template // declare pointer to a template

// Using the init function to make sure the template is only parsed once in the program
func init() {
	// template.Must takes the reponse of template.ParseFiles and does error checking
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
	//parses all gohtml files
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func main() {
	http.HandleFunc("/", index_handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
