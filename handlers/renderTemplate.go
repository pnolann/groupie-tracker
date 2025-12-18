package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate is a helper function to render a single template file.
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {

	templatePath := "templates/" + tmpl

	// Parse the specified template file.
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		// Log the error
		log.Printf("Error parsing template %s: %v", templatePath, err)
		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template
	err = t.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template %s: %v", tmpl, err)
	}
}
