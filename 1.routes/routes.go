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

	mux.Get("/home", handlers.Home)
	return mux
}
