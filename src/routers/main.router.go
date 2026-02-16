package routers

import (
	"net/http"
	"projet-igdb/src/controllers"
)

func MainRouter() *http.ServeMux {

	mainRouter := http.NewServeMux()

	errorRouter(mainRouter)

	InitGamesRoutes(mainRouter)

	FavoritesRoutes(mainRouter)

	SearchRoutes(mainRouter)

	fileServer := http.FileServer(http.Dir("./assets"))
	mainRouter.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mainRouter.HandleFunc("/dino", func(w http.ResponseWriter, r *http.Request) {
		controllers.DinoDisplay(w, r)
	})

	mainRouter.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		controllers.AboutDisplay(w, r)
	})

	mainRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			controllers.NotFoundDisplay(w, r)
			return
		}
		controllers.DisplayHome(w, r)
	})

	return mainRouter
}
