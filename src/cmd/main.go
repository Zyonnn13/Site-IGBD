package main

import (
	"fmt"
	"log"
	"net/http"
	"projet-igdb/src/routers"
	"projet-igdb/src/templates"
)

func main() {

	templates.Load()

	mux := routers.MainRouter()

	addr := "localhost:8080"

	fmt.Printf("Serveur prêt sur http://%s\n", addr)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatalf("Erreur lancement serveur : %s\n", err.Error())
	}
}
