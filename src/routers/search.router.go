package routers

import (
	"net/http"
	"projet-igdb/src/controllers"
)


func SearchRoutes(router *http.ServeMux) {
	router.HandleFunc("/search", controllers.SearchHandler)
}
