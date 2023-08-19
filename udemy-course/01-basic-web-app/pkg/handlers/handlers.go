package handlers

import (
	"net/http"

	"github.com/afrid18/basicwebapp/pkg/config"
	"github.com/afrid18/basicwebapp/pkg/models"
	"github.com/afrid18/basicwebapp/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	td := models.TemplateData{
		StringMap: map[string]string{"test": "Hello Bro"},
	}
	render.RenderTemplate(w, "about.page.tmpl", &td)
}
