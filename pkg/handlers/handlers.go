package handlers

import (
	"net/http"

	"github.com/jbrainz/go-web/pkg/config"
	"github.com/jbrainz/go-web/pkg/models"
	"github.com/jbrainz/go-web/pkg/render"
)

var Repo *Repository

// Repository type for the repo
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates the repository for the
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the homepage handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.html", &models.TemplateData{})
}

// About :  the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Check this out"
	render.Template(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
