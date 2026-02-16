package controllers

import (
	"fmt"
	"net/http"
	"projet-igdb/src/services"
	"projet-igdb/src/templates"
	"strings"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	query = strings.TrimSpace(query)

	if query == "" {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	games, statusCode, err := services.SearchGames(query)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?code=%d&message=%s", statusCode, err.Error()), http.StatusSeeOther)
		return
	}

	queryLower := strings.ToLower(query)
	var filteredGames []services.SearchResult
	for _, game := range games {
		checkName := strings.Contains(strings.ToLower(game.Name), queryLower)

		checkGenre := false
		for _, genre := range game.Genres {
			if strings.Contains(strings.ToLower(genre.Name), queryLower) {
				checkGenre = true
				break
			}
		}

		if checkName || checkGenre {
			filteredGames = append(filteredGames, services.SearchResult{Game: game})
		}
	}

	data := struct {
		Query   string
		Results []services.SearchResult
		Count   int
	}{
		Query:   query,
		Results: filteredGames,
		Count:   len(filteredGames),
	}

	templates.RenderTemplate(w, r, "search", data)
}
