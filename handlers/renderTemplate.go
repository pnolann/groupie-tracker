package handlers

import (
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string){
	t, err := template.ParseFiles("templates/"+ tmpl )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}