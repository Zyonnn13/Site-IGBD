package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projet-igdb/src/models"
	"strings"
	"time"
)

type SearchResult struct {
	Game models.Game
}

func GetGamesByStudio(studioID string) ([]models.Game, int, error) {
	client := http.Client{
		Timeout: 15 * time.Second,
	}
	url := "https://api.igdb.com/v4/games"

	token, errToken := GetIGDBToken()
	if errToken != nil {
		return nil, http.StatusInternalServerError, errToken
	}

	queryBody := fmt.Sprintf(`
		fields 
			id,
			name, 
			summary,
			first_release_date,
			rating,
			slug,
			category,
			cover.url,
			genres.name,
			platforms.name;
		where involved_companies.company = %s 
		& cover != null;
		sort first_release_date desc;
		limit 500;
	`, studioID)

	request, err := http.NewRequest("POST", url, strings.NewReader(queryBody))
	if err != nil {
		return nil, 0, fmt.Errorf("erreur création de la requête: %w", err)
	}

	request.Header.Add("Client-ID", ClientID)
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := client.Do(request)
	if err != nil {
		return nil, 0, fmt.Errorf("erreur lors de l'appel API: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, response.StatusCode, fmt.Errorf("l'API a retourné le code %d: %s", response.StatusCode, response.Status)
	}

	var games []models.Game
	err = json.NewDecoder(response.Body).Decode(&games)
	if err != nil {
		return nil, response.StatusCode, fmt.Errorf("erreur décodage JSON: %w", err)
	}

	for i := range games {
		if !strings.HasPrefix(games[i].Cover.URL, "https:") {
			games[i].Cover.URL = "https:" + games[i].Cover.URL
		}
		games[i].Cover.URL = strings.Replace(games[i].Cover.URL, "t_thumb", "t_cover_big", 1)

		for j := range games[i].Screenshots {
			if !strings.HasPrefix(games[i].Screenshots[j].URL, "https:") {
				games[i].Screenshots[j].URL = "https:" + games[i].Screenshots[j].URL
			}
			games[i].Screenshots[j].URL = strings.Replace(games[i].Screenshots[j].URL, "t_thumb", "t_screenshot_big", 1)
		}
	}

	return games, response.StatusCode, nil
}

func SearchGames(query string) ([]models.Game, int, error) {
	client := http.Client{
		Timeout: 15 * time.Second,
	}
	url := "https://api.igdb.com/v4/games"

	token, errToken := GetIGDBToken()
	if errToken != nil {
		return nil, http.StatusInternalServerError, errToken
	}

	queryBody := fmt.Sprintf(`
		search "%s";
		fields 
			id,
			name, 
			summary,
			first_release_date,
			rating,
			slug,
			category,
			cover.url,
			genres.name,
			platforms.name;
		where cover != null;
		limit 50;
	`, query)

	request, err := http.NewRequest("POST", url, strings.NewReader(queryBody))
	if err != nil {
		return nil, 0, fmt.Errorf("erreur création de la requête: %w", err)
	}

	request.Header.Add("Client-ID", ClientID)
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := client.Do(request)
	if err != nil {
		return nil, 0, fmt.Errorf("erreur lors de l'appel API: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, response.StatusCode, fmt.Errorf("l'API a retourné le code %d: %s", response.StatusCode, response.Status)
	}

	var games []models.Game
	err = json.NewDecoder(response.Body).Decode(&games)
	if err != nil {
		return nil, response.StatusCode, fmt.Errorf("erreur décodage JSON: %w", err)
	}

	for i := range games {
		if !strings.HasPrefix(games[i].Cover.URL, "https:") {
			games[i].Cover.URL = "https:" + games[i].Cover.URL
		}
		games[i].Cover.URL = strings.Replace(games[i].Cover.URL, "t_thumb", "t_cover_big", 1)
	}

	return games, response.StatusCode, nil
}

func GetGameDetails(gameID string) (*models.Game, int, error) {
	client := http.Client{
		Timeout: 15 * time.Second,
	}

	url := "https://api.igdb.com/v4/games"

	token, errToken := GetIGDBToken()
	if errToken != nil {
		return nil, http.StatusInternalServerError, errToken
	}

	queryBody := fmt.Sprintf(`
		fields 
			id,
			name, 
			summary,
			storyline,
			first_release_date,
			rating,
			slug,
			cover.url,
			screenshots.url,
			genres.name,
			platforms.name;
		where id = %s;
	`, gameID)

	request, err := http.NewRequest("POST", url, strings.NewReader(queryBody))
	if err != nil {
		return nil, 0, fmt.Errorf("erreur création de la requête: %w", err)
	}

	request.Header.Add("Client-ID", ClientID)
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := client.Do(request)
	if err != nil {
		return nil, 0, fmt.Errorf("erreur lors de l'appel API: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, response.StatusCode, fmt.Errorf("l'API a retourné le code %d: %s", response.StatusCode, response.Status)
	}

	var games []models.Game
	err = json.NewDecoder(response.Body).Decode(&games)
	if err != nil {
		return nil, response.StatusCode, fmt.Errorf("erreur décodage JSON: %w", err)
	}

	if len(games) == 0 {
		return nil, response.StatusCode, fmt.Errorf("jeu introuvable")
	}

	game := &games[0]

	if !strings.HasPrefix(game.Cover.URL, "https:") {
		game.Cover.URL = "https:" + game.Cover.URL
	}
	game.Cover.URL = strings.Replace(game.Cover.URL, "t_thumb", "t_cover_big", 1)

	for j := range game.Screenshots {
		if !strings.HasPrefix(game.Screenshots[j].URL, "https:") {
			game.Screenshots[j].URL = "https:" + game.Screenshots[j].URL
		}
		game.Screenshots[j].URL = strings.Replace(game.Screenshots[j].URL, "t_thumb", "t_screenshot_big", 1)
	}

	return game, response.StatusCode, nil
}

func GetDLCs(parentGameID string) ([]models.Game, int, error) {
	client := http.Client{
		Timeout: 15 * time.Second,
	}
	url := "https://api.igdb.com/v4/games"

	token, errToken := GetIGDBToken()
	if errToken != nil {
		return nil, http.StatusInternalServerError, errToken
	}

	queryBody := fmt.Sprintf(`
		fields 
			id,
			name, 
			cover.url,
			first_release_date,
			category;
		where parent_game = %s 
		& cover != null;
		sort first_release_date asc;
	`, parentGameID)

	request, err := http.NewRequest("POST", url, strings.NewReader(queryBody))
	if err != nil {
		return nil, 0, fmt.Errorf("erreur création de la requête: %w", err)
	}

	request.Header.Add("Client-ID", ClientID)
	request.Header.Add("Authorization", "Bearer "+token)

	response, err := client.Do(request)
	if err != nil {
		return nil, 0, fmt.Errorf("erreur lors de l'appel API: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, response.StatusCode, fmt.Errorf("l'API a retourné le code %d: %s", response.StatusCode, response.Status)
	}

	var dlcs []models.Game
	err = json.NewDecoder(response.Body).Decode(&dlcs)
	if err != nil {
		return nil, response.StatusCode, fmt.Errorf("erreur décodage JSON: %w", err)
	}

	for i := range dlcs {
		if !strings.HasPrefix(dlcs[i].Cover.URL, "https:") {
			dlcs[i].Cover.URL = "https:" + dlcs[i].Cover.URL
		}
		dlcs[i].Cover.URL = strings.Replace(dlcs[i].Cover.URL, "t_thumb", "t_cover_small", 1)
	}

	return dlcs, response.StatusCode, nil
}
