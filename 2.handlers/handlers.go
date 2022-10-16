package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	render "github.com/hdadashi/jabama/3.render"
	"github.com/hdadashi/jabama/config"
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
	return config.PassingSession.LoadAndSave(next)
}

// END------------------------------------------------------------------------------

// routes funcs --------------------------------------------------------------------
func RouteFinder(w http.ResponseWriter, r *http.Request) {

	data, err := config.GlobVar("input")
	render.Scream(err)
	var csrf *render.TemplateData = new(render.TemplateData)
	csrf.CSRF = nosurf.Token(r)
	requestURL := r.URL.String()

	if requestURL == "/" {
		render.Renderer(w, r, "home.page.html", data)
	}
	if requestURL == "/contact" {
		render.Renderer(w, r, "contact.page.html", nil)
	}
	if requestURL == "/about" {
		render.Renderer(w, r, "about.page.html", nil)
	}
	if requestURL == "/rooms/general" {
		render.Renderer(w, r, "general.page.html", nil)
	}
	if requestURL == "/rooms/vip" {
		render.Renderer(w, r, "VIP.page.html", nil)
	}
	if requestURL == "/availability" {
		render.Renderer(w, r, "availability.page.html", csrf)
	}
	if requestURL == "/availabilityJSON" {
		resp := new(JSONresponse)
		resp.OK = true
		resp.Message = "Available!"
		output, err := json.MarshalIndent(resp, "", "     ")

		render.Scream(err)

		w.Write(output)
	}
	if requestURL == "/postAvailability" {
		w.Write([]byte("yo, sup bro?????"))
	}
	if requestURL == "/book" {
		render.Renderer(w, r, "book.page.html", nil)
	}
}

func PostRoute(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("name")
	fmt.Println(name)
}

// END------------------------------------------------------------------------------
type JSONresponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}
