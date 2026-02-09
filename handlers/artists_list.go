package handlers

import (
	"encoding/json"
	"groupie-tracker/models"
	"net/http"
	"strconv"
	"strings"
)

type PageData struct {
	Artists []models.Artists
	Query   string
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed, "Méthode non autorisée.")
		return
	}

	respRoot, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Impossible de contacter l'API Groupie Trackers (Root).")
		return
	}
	defer respRoot.Body.Close()

	var index models.ApiIndex
	if err := json.NewDecoder(respRoot.Body).Decode(&index); err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Erreur lors de la lecture de l'index API.")
		return
	}

	respArtists, err := http.Get(index.Artists)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Impossible de récupérer la liste des artistes.")
		return
	}
	defer respArtists.Body.Close()

	var allArtists []models.Artists
	if err := json.NewDecoder(respArtists.Body).Decode(&allArtists); err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "Erreur lors du décodage des données artistes.")
		return
	}

	query := r.URL.Query().Get("query")
	minDateStr := r.URL.Query().Get("minDate")
	maxDateStr := r.URL.Query().Get("maxDate")
	members := r.URL.Query()["members"]

	var finalArtists []models.Artists

	for _, artist := range allArtists {
		if query != "" && !strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			continue
		}

		if minDateStr != "" {
			min, err := strconv.Atoi(minDateStr)
			if err == nil && artist.CreatedDate < min {
				continue
			}
		}

		if maxDateStr != "" {
			max, err := strconv.Atoi(maxDateStr)
			if err == nil && artist.CreatedDate > max {
				continue
			}
		}

		if len(members) > 0 {
			keep := false
			nb := strconv.Itoa(len(artist.Members))
			for _, m := range members {
				if m == nb {
					keep = true
					break
				}
			}
			if !keep {
				continue
			}
		}

		finalArtists = append(finalArtists, artist)
	}

	data := PageData{
		Artists: finalArtists,
		Query:   query,
	}

	RenderTemplate(w, "artists.html", data)
}