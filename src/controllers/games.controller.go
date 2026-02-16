package controllers

import (
	"fmt"
	"net/http"
	"projet-igdb/src/models"
	"projet-igdb/src/services"
	"projet-igdb/src/templates"
	"slices"
	"strconv"
	"strings"
)

type GamePageData struct {
	StudioName         string
	Games              []models.Game
	AllGenres          []string
	AllPlatforms       []string
	TotalCount         int
	FilteredCount      int
	SelectedCategories []string
	SelectedGenres     []string
	SelectedPlatforms  []string
	Page               int
	Next               int
	Prev               int
	ShowPagination     bool
	FilterQuery        string
}

func SegaHandler(w http.ResponseWriter, r *http.Request) {
	getGamesByStudioID(w, r, "112", "SEGA", "sega")
}

func NintendoHandler(w http.ResponseWriter, r *http.Request) {
	getGamesByStudioID(w, r, "70", "Nintendo", "nintendo")
}

func Level5Handler(w http.ResponseWriter, r *http.Request) {
	getGamesByStudioID(w, r, "837", "Level-5", "level-5")
}

func UbisoftHandler(w http.ResponseWriter, r *http.Request) {
	getGamesByStudioID(w, r, "104", "Ubisoft", "ubisoft")
}

func SonyHandler(w http.ResponseWriter, r *http.Request) {
	getGamesByStudioID(w, r, "10100", "Sony Interactive Entertainment", "games_list")
}

func getGamesByStudioID(w http.ResponseWriter, r *http.Request, studioID string, studioName string, templateName string) {

	games, statusCode, err := services.GetGamesByStudio(studioID)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?code=%d&message=%s", statusCode, err.Error()), http.StatusSeeOther)
		return
	}

	r.ParseForm()

	categories := r.Form["categories"]
	genres := r.Form["genres"]
	platforms := r.Form["platforms"]

	filteredGames := []models.Game{}

	for _, game := range games {
		checkCategory := len(categories) == 0 || slices.Contains(categories, strconv.Itoa(game.Category))
		checkGenre := len(genres) == 0 || slices.ContainsFunc(game.Genres, func(g models.Genre) bool {
			return slices.Contains(genres, strings.ToLower(g.Name))
		})
		checkPlatform := len(platforms) == 0 || slices.ContainsFunc(game.Platforms, func(p models.Platform) bool {
			return slices.Contains(platforms, strings.ToLower(p.Name))
		})
		if checkCategory && checkGenre && checkPlatform {
			filteredGames = append(filteredGames, game)
		}
	}

	allGenres := collectUniqueGenres(games)
	allPlatforms := collectUniquePlatforms(games)

	pageStr := r.FormValue("page")
	pageNbr, _ := strconv.Atoi(pageStr)
	if pageNbr < 0 {
		pageNbr = 0
	}

	perPage := 20
	startIndex := pageNbr * perPage
	endIndex := startIndex + perPage
	totalFiltered := len(filteredGames)

	if startIndex >= totalFiltered {
		pageNbr = 0
		startIndex = 0
		endIndex = perPage
	}
	if endIndex > totalFiltered {
		endIndex = totalFiltered
	}

	paginatedGames := filteredGames[startIndex:endIndex]

	filterQuery := ""
	for _, c := range categories {
		filterQuery += "&categories=" + c
	}
	for _, g := range genres {
		filterQuery += "&genres=" + g
	}
	for _, p := range platforms {
		filterQuery += "&platforms=" + p
	}

	viewData := GamePageData{
		StudioName:         studioName,
		Games:              paginatedGames,
		AllGenres:          allGenres,
		AllPlatforms:       allPlatforms,
		TotalCount:         len(games),
		FilteredCount:      totalFiltered,
		SelectedCategories: categories,
		SelectedGenres:     genres,
		SelectedPlatforms:  platforms,
		Page:               pageNbr,
		Next:               pageNbr + 1,
		Prev:               pageNbr - 1,
		ShowPagination:     totalFiltered > perPage,
		FilterQuery:        filterQuery,
	}

	templates.RenderTemplate(w, r, templateName, viewData)
}

func collectUniqueGenres(games []models.Game) []string {
	genreMap := make(map[string]bool)
	for _, game := range games {
		for _, genre := range game.Genres {
			genreMap[genre.Name] = true
		}
	}

	genres := []string{}
	for genre := range genreMap {
		genres = append(genres, genre)
	}
	return genres
}

func collectUniquePlatforms(games []models.Game) []string {
	platformMap := make(map[string]bool)
	for _, game := range games {
		for _, platform := range game.Platforms {
			platformMap[platform.Name] = true
		}
	}

	platforms := []string{}
	for platform := range platformMap {
		platforms = append(platforms, platform)
	}
	return platforms
}

func GameDetailsHandler(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("id")
	if gameID == "" {
		http.Redirect(w, r, "/error?code=400&message=ID du jeu manquant", http.StatusSeeOther)
		return
	}

	game, statusCode, err := services.GetGameDetails(gameID)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?code=%d&message=%s", statusCode, err.Error()), http.StatusSeeOther)
		return
	}

	dlcs, _, _ := services.GetDLCs(gameID)

	isFavorite := services.IsFavorite(gameID, "game")

	data := struct {
		Game       *models.Game
		DLCs       []models.Game
		IsFavorite bool
	}{
		Game:       game,
		DLCs:       dlcs,
		IsFavorite: isFavorite,
	}

	templates.RenderTemplate(w, r, "game_details", data)
}
