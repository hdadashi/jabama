package handlers

import (
	"fmt"
	"net/http"
	"strings"

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
	var csrf *render.TemplateData = new(render.TemplateData)
	csrf.CSRF = nosurf.Token(r)

	render.Scream(err)

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
	if requestURL == "/postAvailability" {
		start := r.Form.Get("sdate")
		end := r.Form.Get("edate")
		start = strings.ReplaceAll(start, "-", "")
		end = strings.ReplaceAll(end, "-", "")
		w.Write([]byte(fmt.Sprintf("start is: %s and end is: %s and ", start, end)))
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
