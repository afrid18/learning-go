package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/afrid18/basicwebapp/pkg/config"
	"github.com/afrid18/basicwebapp/pkg/handlers"
	"github.com/afrid18/basicwebapp/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application at portNumber %s\n", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}
