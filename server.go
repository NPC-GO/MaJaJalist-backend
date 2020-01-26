package main

import (
	"fmt"
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
	}) //連接到database
	defer DB.Close() //在最後斷開連接
	server := router.InitRouter(DB)
	err := http.ListenAndServeTLS(":443", "./certs/server.crt", "./certs/server.key", server)
	if err != nil {
		fmt.Println(err)
		http.ListenAndServe(":80", server) //無法使用https則用http
	}
}
