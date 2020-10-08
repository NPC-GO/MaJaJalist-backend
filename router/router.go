package router

import (
	"github.com/NPC-GO/MaJaJalist-backend/database"
	"github.com/NPC-GO/MaJaJalist-backend/middleware"
	"github.com/NPC-GO/MaJaJalist-backend/router/handler"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"net/http"
)

func InitRouter(userDatabaseCtrl database.User) chi.Router {
	router := chi.NewRouter()
	router.Use(
		chiMiddleware.Logger,
		chiMiddleware.Recoverer,
	)
	router.With(middleware.BeforeLoginAuth(userDatabaseCtrl)).Post("/login", handler.Login(userDatabaseCtrl))
	//用middleware擋住已經登錄的
	router.Get("/*", http.StripPrefix("/", http.FileServer(http.Dir("dist"))).ServeHTTP) //用來提供js與css檔案
	router.Get("/", handler.HtmlHandler)
	router.Get("/unfinished", handler.HtmlHandler)
	router.Get("/finished", handler.HtmlHandler)
	router.Get("/trashcan", handler.HtmlHandler)
	router.Get("/settings", handler.HtmlHandler)
	router.Get("/login", handler.HtmlHandler)
	return router
}
