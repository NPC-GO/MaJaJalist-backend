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
	}) //連接到database的設定
	defer DB.Close() //在最後斷開連接
	userDatabaseCtrl := database.User{DB: DB}
	server := router.InitRouter(userDatabaseCtrl) //把db傳進router就不用每次要使用db都連接一次
	err := http.ListenAndServeTLS(":443", "./certs/server.crt", "./certs/server.key", server)
	if err != nil {
		fmt.Println(err)
		http.ListenAndServe(":80", server) //無法使用https則用http
	}
}
