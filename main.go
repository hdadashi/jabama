package main

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	routes "github.com/hdadashi/jabama/1.routes"
	render "github.com/hdadashi/jabama/3.render"
	"github.com/hdadashi/jabama/config"
	"github.com/hdadashi/jabama/models"
)

var sessionManager *scs.SessionManager

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

	config.PassingSession = sessionManager

	fmt.Println("Server is running on port 8080")
	server := &http.Server{
		Addr:    ":8080",
		Handler: routes.Routes(),
	}
	err := server.ListenAndServe()
	render.Scream(err)
}
