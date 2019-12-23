package router

import (
	"github.com/NPC-GO/MaJaJalist-backend/router/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InitRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
	)
	router.Get("/", handler.HtmlHandler)
	return router
}
