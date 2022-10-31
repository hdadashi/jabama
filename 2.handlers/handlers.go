package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	render "github.com/hdadashi/jabama/3.render"
	"github.com/hdadashi/jabama/config"
	"github.com/hdadashi/jabama/driver"
	"github.com/hdadashi/jabama/forms"
	"github.com/hdadashi/jabama/models"
	"github.com/justinas/nosurf"
)

// middlewares funcs ---------------------------------------------------------------
func CSRF(next http.Handler) http.Handler {
	csrf := nosurf.New(next)
	csrf.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrf
}

func SessionLoad(next http.Handler) http.Handler {
	x := *Repo.App.Session
	return x.LoadAndSave(next)
}

// END------------------------------------------------------------------------------

// routes funcs --------------------------------------------------------------------

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		//DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Repository

var data *render.TemplateData = new(render.TemplateData)

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Renderer(w, r, "home.page.html", &render.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.Renderer(w, r, "about.page.html", &render.TemplateData{})
}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.Renderer(w, r, "general.page.html", &render.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.Renderer(w, r, "VIP.page.html", &render.TemplateData{})
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Renderer(w, r, "contact.page.html", &render.TemplateData{})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.Renderer(w, r, "book.page.html", &render.TemplateData{
		Form: forms.New(nil),
		Data: data.Data,
	})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.Renderer(w, r, "availability.page.html", &render.TemplateData{})
}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	render.Scream(err)

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		render.Renderer(w, r, "make-reservation.page.html", &render.TemplateData{
			Form: form,
			Data: data.Data,
		})
		return
	}
	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}
	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// AvailabilityJSON handles request for availability and sends JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := JSONresponse{
		OK:      true,
		Message: "Available!",
	}

	out, _ := json.MarshalIndent(resp, "", "     ")

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// PostAvailability handles post
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end is %s", start, end)))
}

// END------------------------------------------------------------------------------
type JSONresponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}
