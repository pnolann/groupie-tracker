package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"log"
	"net/http"
)

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/artists" {
		http.NotFound(w, r)
		return
	}

	// Calls the api at the root
	respRoot, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		log.Println("Root API error:", err)
		http.Error(w, "Server error (API Root)", 500)
		return
	}
	defer respRoot.Body.Close()

	// Decoding the root response
	var index models.ApiIndex
	if err := json.NewDecoder(respRoot.Body).Decode(&index); err != nil {
		http.Error(w, "Erreur décodage Index", 500)
		return
	}

	// Calls the response
	respArtists, err := http.Get(index.Artists)
	if err != nil {
		log.Println("Erreur API Artists:", err)
		http.Error(w, "Erreur serveur (API Artists)", 500)
		return
	}
	defer respArtists.Body.Close()

	// Decoding the artists
	var artists []models.Artists
	if err := json.NewDecoder(respArtists.Body).Decode(&artists); err != nil {
		fmt.Println("Détail de l'erreur :", err) // Affiche la raison exacte dans le terminal
		http.Error(w, "Erreur décodage Artistes", 500)
		return
	}

	RenderTemplate(w, "artists.html", artists)
}
