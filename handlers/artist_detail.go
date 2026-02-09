package handlers

import (
	"encoding/json"
	"groupie-tracker/models"
	"net/http"
)

type ArtistDetailData struct {
	Artist    models.Artists
	Relations models.Relations
}

func ArtistDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	resp, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	defer resp.Body.Close()

	var artist models.Artists
	json.NewDecoder(resp.Body).Decode(&artist)

	respRel, _ := http.Get(artist.Relations)
	defer respRel.Body.Close()

	var relations models.Relations
	json.NewDecoder(respRel.Body).Decode(&relations)

	data := ArtistDetailData{
		Artist:    artist,
		Relations: relations,
	}

	RenderTemplate(w, "artist_detail.html", data)
}