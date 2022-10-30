package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	routes "github.com/hdadashi/jabama/1.routes"
	render "github.com/hdadashi/jabama/3.render"
	"github.com/hdadashi/jabama/config"
	"github.com/hdadashi/jabama/models"
)

var sessionManager *scs.SessionManager
var app config.AppConfig
var InfoLog *log.Logger
var errorLog *log.Logger

func main() {
	//Adding the reservation data to the session
	gob.Register(models.Reservation{})

	sessionManager = scs.New()

	//the time that session lives
	sessionManager.Lifetime = 1 * time.Hour

	//do not encrypt the cookies
	sessionManager.Cookie.Secure = false

	//sessions do not persist after the browser is closed
	sessionManager.Cookie.Persist = true

	//کوکی‌ها فقط زمانی که کاربر به صورت معمولی سایت را باز کند در دسترس خواهند بود و هدایت کاربر به سایت از هر طریقی، بدون کوکی انجام می‌شود
	sessionManager.Cookie.SameSite = http.SameSiteStrictMode

	app.Session = sessionManager

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return
	}

	app.TemplateCache = tc
	app.UseCache = false

	fmt.Println("Server is running on port 8080")
	server := &http.Server{
		Addr:    ":8080",
		Handler: routes.Routes(),
	}
	err = server.ListenAndServe()
	render.Scream(err)
}
