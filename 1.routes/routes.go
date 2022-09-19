package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	handlers "github.com/hdadashi/jabama/2.handlers"
)

func RouteHome() http.Handler {
	mux := chi.NewRouter()

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
	return mux
}
