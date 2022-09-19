package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	handlers "github.com/hdadashi/jabama/2.handlers"
)

var mux = chi.NewRouter()

func RouteHome() http.Handler {
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

	mux.Get("/", handlers.Home)
	mux.Get("/about", handlers.Home)
	mux.Get("/contact", handlers.Home)
	return mux
}
