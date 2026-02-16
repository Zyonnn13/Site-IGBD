package controllers

import (
	"net/http"
	"projet-igdb/src/templates"
)

func NotFoundDisplay(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	templates.RenderTemplate(w, r, "error404", nil)
}

func DinoDisplay(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, r, "dino", nil)
}

func AboutDisplay(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, r, "about", nil)
}
