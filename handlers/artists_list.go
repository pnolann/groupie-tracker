package handlers

import (
	"net/http"
)

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {

	RenderTemplate(w, "artists.html", nil)

	if r.URL.Path != "/artists" {
		http.NotFound(w, r)
		return
	}
}
