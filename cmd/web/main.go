package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sazzadnibir/go_modules/pkg/config"
	"github.com/sazzadnibir/go_modules/pkg/handlers"
	"github.com/sazzadnibir/go_modules/pkg/render"
)

const port = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can't create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandalers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", port))
	// http.ListenAndServe(port, nil)

	serve := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}
