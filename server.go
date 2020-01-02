package main

import (
	"github.com/NPC-GO/MaJaJalist-backend/database"
	"github.com/NPC-GO/MaJaJalist-backend/router"
	"github.com/go-pg/pg/v9"
	"net/http"
)

func main() {
	DB := database.New(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "MaJaJalist",
		Addr:     "database:5432",
	})
	defer DB.Close()
	server := router.InitRouter()
	http.ListenAndServe(":80", server)
}
