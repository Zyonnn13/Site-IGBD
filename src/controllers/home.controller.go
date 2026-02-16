package controllers

import (
	"net/http"
	"projet-igdb/src/templates"
)

func DisplayHome(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, r, "index", nil)
}
