package router

import (
	"github.com/NPC-GO/MaJaJalist-backend/database"
	"github.com/NPC-GO/MaJaJalist-backend/middleware"
	"github.com/NPC-GO/MaJaJalist-backend/router/handler"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v9"
	"net/http"
)

func InitRouter(db *pg.DB) chi.Router {
	router := chi.NewRouter()
	router.Use(
		chimiddleware.Logger,
		chimiddleware.Recoverer,
	)
	router.With(middleware.BeforeLoginAuth(database.User{DB: db})).Post("/login")
	router.Get("/*", http.StripPrefix("/", http.FileServer(http.Dir("dist"))).ServeHTTP)
	router.Get("/", handler.HtmlHandler)
	return router
}
