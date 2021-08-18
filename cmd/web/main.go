package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jbrainz/go-web/pkg/config"
	"github.com/jbrainz/go-web/pkg/handlers"
	"github.com/jbrainz/go-web/pkg/render"
)

const port = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template file")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)
	http.HandleFunc("/", repo.Home)
	http.HandleFunc("/about", repo.About)

	fmt.Println(fmt.Sprintf("starting application o port %s", port))
	_ = http.ListenAndServe(port, nil)
}
