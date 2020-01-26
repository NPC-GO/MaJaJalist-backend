package router

import (
	"github.com/NPC-GO/MaJaJalist-backend/database"
	"github.com/NPC-GO/MaJaJalist-backend/middleware"
	"github.com/NPC-GO/MaJaJalist-backend/router/handler"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v9"
	"net/http"
)

func InitRouter(db *pg.DB) chi.Router {
	router := chi.NewRouter()
	router.Use(
		chiMiddleware.Logger,
		chiMiddleware.Recoverer,
	)
	router.With(middleware.BeforeLoginAuth(database.User{DB: db})).Post("/login", handler.Login(database.User{DB: db}))
	//用middleware擋住已經登錄的
	router.Get("/*", http.StripPrefix("/", http.FileServer(http.Dir("dist"))).ServeHTTP)
	router.Get("/", handler.HtmlHandler)
	return router
}
