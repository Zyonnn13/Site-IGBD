package templates

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var listTemp *template.Template


var funcMap = template.FuncMap{
	"lower": strings.ToLower,
	"contains": func(list []string, val string) bool {
		for _, v := range list {
			if v == val {
				return true
			}
		}
		return false
	},
	"itoa": func(i int) string {
		return fmt.Sprintf("%d", i)
	},
}

func Load() {
	listTemplates, errTemplates := template.New("").Funcs(funcMap).ParseGlob("./templates/*.html")
	if errTemplates != nil {
		log.Fatalf("Erreur chargement des templates : %s", errTemplates.Error())
	}
	listTemp = listTemplates
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	var buffer bytes.Buffer

	errRender := listTemp.ExecuteTemplate(&buffer, name, data)
	if errRender != nil {
		msg := url.QueryEscape("Erreur d'affichage du template : " + errRender.Error())
		http.Redirect(w, r, fmt.Sprintf("/error?code=500&message=%s", msg), http.StatusSeeOther)
		return
	}

	_, _ = buffer.WriteTo(w)
}
