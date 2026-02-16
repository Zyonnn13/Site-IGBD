package controllers

import (
	"net/http"
	"projet-igdb/src/services"
	"projet-igdb/src/templates"
)


func DisplayFavorites(w http.ResponseWriter, r *http.Request) {
	favorites, err := services.GetAllFavorites()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des favoris", http.StatusInternalServerError)
		return
	}

	data := struct {
		Favorites []services.Favorite
	}{
		Favorites: favorites,
	}

	templates.RenderTemplate(w, r, "favorites", data)
}


func AddToFavorites(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}


	id := r.FormValue("id")
	favType := r.FormValue("type")
	name := r.FormValue("name")
	image := r.FormValue("image")

	if id == "" || favType == "" || name == "" {
		http.Error(w, "Paramètres manquants", http.StatusBadRequest)
		return
	}


	err := services.AddFavorite(id, favType, name, image)
	if err != nil {
		http.Error(w, "Erreur lors de l'ajout aux favoris", http.StatusInternalServerError)
		return
	}

	
	referer := r.Header.Get("Referer")
	if referer == "" {
		referer = "/"
	}
	http.Redirect(w, r, referer, http.StatusSeeOther)
}


func RemoveFromFavorites(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	
	id := r.FormValue("id")
	favType := r.FormValue("type")

	if id == "" || favType == "" {
		http.Error(w, "Paramètres manquants", http.StatusBadRequest)
		return
	}

	
	err := services.RemoveFavorite(id, favType)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression des favoris", http.StatusInternalServerError)
		return
	}

	
	http.Redirect(w, r, "/favorites", http.StatusSeeOther)
}
