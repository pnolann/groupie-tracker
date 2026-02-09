package handlers

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound, "La page que vous cherchez semble avoir été dévorée par un tigre.")
		return
	}

	RenderTemplate(w, "index.html", nil)
}