package handlers

import(
	"net/http"
)

func Home(w http.ResponseWriter, r http.Request){
	renderTemplate(w, "home.html")
}