package router

import (
	"github.com/NPC-GO/MaJaJalist-backend/router/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func InitRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
	)
	router.Get("/*", http.StripPrefix("/", http.FileServer(http.Dir("dist"))).ServeHTTP)
	router.Get("/", handler.HtmlHandler)
	return router
}
