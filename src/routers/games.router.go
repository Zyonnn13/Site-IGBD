package routers

import (
	"net/http"
	"projet-igdb/src/controllers"
)

func InitGamesRoutes(router *http.ServeMux) {
	router.HandleFunc("/sega", controllers.SegaHandler)
	router.HandleFunc("/nintendo", controllers.NintendoHandler)
	router.HandleFunc("/level5", controllers.Level5Handler)
	router.HandleFunc("/ubisoft", controllers.UbisoftHandler)
	router.HandleFunc("/sony", controllers.SonyHandler)

	router.HandleFunc("/game", controllers.GameDetailsHandler)
}
