package handlers

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	RenderTemplate(w, "index.html", nil)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
}
