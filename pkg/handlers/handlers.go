package handlers

import (
	"net/http"

	"github.com/tsawler/go-course/pkg/config"
	"github.com/tsawler/go-course/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// in order for a function to respond to a request from a web browser it needs two parameters
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is the home page")

	//-----
	render.RenderTemplate(w, "home.page.html")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// sum := addValues(2, 2)
	// // fmt.Printf(w, "This is the about page")
	// _, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))

	// -----------
	render.RenderTemplate(w, "about.page.html")
}
