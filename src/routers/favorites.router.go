package routers

import (
	"net/http"
	"projet-igdb/src/controllers"
)

func FavoritesRoutes(router *http.ServeMux) {

	router.HandleFunc("/favorites", controllers.DisplayFavorites)

	router.HandleFunc("/favorites/add", controllers.AddToFavorites)

	router.HandleFunc("/favorites/remove", controllers.RemoveFromFavorites)
}
