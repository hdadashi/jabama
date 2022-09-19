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
func Home(w http.ResponseWriter, r *http.Request) {

	data, err := config.GlobVar("input")
	render.Scream(err)
	//get the user ip address
	data.IP = r.RemoteAddr
	render.Renderer(w, "home.page.html", data)
}

// END------------------------------------------------------------------------------
