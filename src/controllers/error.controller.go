package controllers

import (
	"net/http"
	"projet-igdb/src/models"
	"projet-igdb/src/templates"
)

func ErrorDisplay(w http.ResponseWriter, r *http.Request) {

	data := models.Error{
		Code:    r.FormValue("code"),
		Message: r.FormValue("message"),
	}

	templates.RenderTemplate(w, r, "error", data)
}
