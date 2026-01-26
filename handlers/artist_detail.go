package handlers

import (
	"encoding/json"
	"groupie-tracker/models"
	"net/http"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID manquant", http.StatusBadRequest)
		return
	}

	urlArtist := "https://groupietrackers.herokuapp.com/api/artists/" + id
	resp, err := http.Get(urlArtist)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération de l'artiste", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var artist models.Artists
	if err := json.NewDecoder(resp.Body).Decode(&artist); err != nil {
		http.Error(w, "Erreur lors du décodage de l'artiste", http.StatusInternalServerError)
		return
	}

	respRel, err := http.Get(artist.Relations)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des relations", http.StatusInternalServerError)
		return
	}
	defer respRel.Body.Close()

	var relations models.Relations
	if err := json.NewDecoder(respRel.Body).Decode(&relations); err != nil {
		http.Error(w, "Erreur lors du décodage des relations", http.StatusInternalServerError)
		return
	}

	artist.RelationsData = relations

	RenderTemplate(w, "artist.html", artist)
}
