package routers

import (
	"net/http"
	"projet-igdb/src/controllers"
)

func errorRouter(router *http.ServeMux) {
	router.HandleFunc("/error", controllers.ErrorDisplay)
}
