package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	handlers "github.com/hdadashi/jabama/2.handlers"
)

var mux = chi.NewRouter()

func Routes() http.Handler {
	//use middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(handlers.CSRF)
	mux.Use(handlers.SessionLoad)
	//------------end

	//let application to use the main root files (eg. images)
	dir := http.Dir("./")
	fileServer := http.FileServer(dir)
	stripPrefix := http.StripPrefix("/", fileServer)
	mux.Handle("/*", stripPrefix)
	//------------end

	//adding routes--
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/rooms/general", handlers.Repo.Generals)
	mux.Get("/rooms/vip", handlers.Repo.Majors)
	mux.Get("/book", handlers.Repo.Reservation)
	mux.Get("/availability", handlers.Repo.Availability)

	mux.Post("/PostBook", handlers.Repo.PostReservation)
	mux.Post("/availabilityJSON", handlers.Repo.AvailabilityJSON)
	mux.Post("/postAvailability", handlers.Repo.PostAvailability)
	//------------end

	return mux
}
