package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Passing multiple template files as an argument
	tpl, err := template.ParseFiles("templates/one.gohtml", "templates/two.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	// Executing template one.gohtml using tpl of type *Template
	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
	// Passing one more file to the *Template
	tpl, err = tpl.ParseFiles("templates/three.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	// Executing template two.gohtml using tpl of type *Template
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
	// Executing template two.gohtml using tpl of type *Template
	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}
