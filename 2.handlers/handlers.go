package handlers

import (
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
		Path:     "/home",
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
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

	requestURL := r.URL.String()
	if requestURL == "/" {
		render.Renderer(w, "home.page.html", data)
	}
	if requestURL == "/contact" {
		render.Renderer(w, "contact.page.html", nil)
	}
	if requestURL == "/about" {
		render.Renderer(w, "about.page.html", nil)
	}
	if requestURL == "/rooms/general" {
		render.Renderer(w, "general.page.html", nil)
	}
	if requestURL == "/rooms/vip" {
		render.Renderer(w, "VIP.page.html", nil)
	}
	if requestURL == "/book" {
		render.Renderer(w, "book.page.html", nil)
	}
}

// END------------------------------------------------------------------------------
