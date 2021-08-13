package handlers

import (
	"net/http"

	"github.com/jbrainz/go-web/pkg/render"
)

// Home is the homepage handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

// About :  the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}
