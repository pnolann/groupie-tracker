package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {

	templatePath := "templates/" + tmpl

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("Error parsing template %s: %v", templatePath, err)
		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template %s: %v", tmpl, err)
	}
}
