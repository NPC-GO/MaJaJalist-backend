package handler

import (
	"github.com/NPC-GO/MaJaJalist-backend/database"
	"net/http"
)

func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./dist/index.html")
}
func Login(user database.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
} //post /login 的handler
