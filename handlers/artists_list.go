package handlers

import (
	"encoding/json"
	"groupie-tracker/models"
	"net/http"
	"strings"
)

type PageData struct {
	Artists []models.Artists
	Query   string
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	respRoot, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		http.Error(w, "Server error (API Root)", 500)
		return
	}
	defer respRoot.Body.Close()

	var index models.ApiIndex
	if err := json.NewDecoder(respRoot.Body).Decode(&index); err != nil {
		http.Error(w, "Erreur décodage Index", 500)
		return
	}

	respArtists, err := http.Get(index.Artists)
	if err != nil {
		http.Error(w, "Erreur serveur (API Artists)", 500)
		return
	}
	defer respArtists.Body.Close()

	var allArtists []models.Artists
	if err := json.NewDecoder(respArtists.Body).Decode(&allArtists); err != nil {
		http.Error(w, "Erreur décodage Artistes", 500)
		return
	}

	query := r.URL.Query().Get("query")
	var finalArtists []models.Artists

	if query == "" {
		finalArtists = allArtists
	} else {
		lowerQuery := strings.ToLower(query)
		for _, artist := range allArtists {
			if strings.Contains(strings.ToLower(artist.Name), lowerQuery) {
				finalArtists = append(finalArtists, artist)
			}
		}
	}

	data := PageData{
		Artists: finalArtists,
		Query:   query,
	}

	RenderTemplate(w, "artists.html", data)
}