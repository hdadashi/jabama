package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	render "github.com/hdadashi/jabama/3.render"
	"github.com/hdadashi/jabama/config"
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
	return config.PassingSession.LoadAndSave(next)
}

// END------------------------------------------------------------------------------

// routes funcs --------------------------------------------------------------------
func RouteFinder(w http.ResponseWriter, r *http.Request) {

	data, err := config.GlobVar("input")
	var csrf *render.TemplateData = new(render.TemplateData)
	csrf.CSRF = nosurf.Token(r)
	csrf.Form = forms.New(nil)

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
	if requestURL == "/availabilityJSON" {
		resp := new(JSONresponse)
		resp.OK = true
		resp.Message = r.PostFormValue("name") + ", its Available!"
		out, err := json.MarshalIndent(resp, "", "     ")
		render.Scream(err)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Write(out)
	}
	if requestURL == "/book" {
		var emptyReservation models.Reservation
		csrf.Data = emptyReservation
		render.Renderer(w, r, "book.page.html", csrf)
	}
	if requestURL == "/PostBook" {
		err := r.ParseForm()
		render.Scream(err)

		reservation := models.Reservation{
			Name:  r.Form.Get("name"),
			Lname: r.Form.Get("lname"),
			Email: r.Form.Get("email"),
			Phone: r.Form.Get("phone"),
		}
		form := forms.New(r.PostForm)
		form.Required("name", "lname", "email")
		form.MinLength("name", 3, r)
		form.IsEmail("email")
		var data *render.TemplateData = new(render.TemplateData)
		data.CSRF = nosurf.Token(r)
		data.Data = reservation
		data.Form = form
		if form.Valid() {
			render.Renderer(w, r, "reservationSummary.page.html", data)
		} else {
			render.Renderer(w, r, "book.page.html", data)
		}
		return
	}
}

// END------------------------------------------------------------------------------
type JSONresponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}
