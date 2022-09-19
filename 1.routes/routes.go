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

	mux.Get("/", handlers.RouteFinder)
	mux.Get("/about", handlers.RouteFinder)
	mux.Get("/contact", handlers.RouteFinder)
	mux.Get("/rooms/general", handlers.RouteFinder)
	mux.Get("/rooms/vip", handlers.RouteFinder)
	return mux
}
