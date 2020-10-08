package middleware

import (
	"github.com/NPC-GO/MaJaJalist-backend/database"
	"net/http"
)

func BeforeLoginAuth(userDatabaseCtrl database.User) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := r.Cookie("session")
			if err == nil {
				if _, err := userDatabaseCtrl.GetUserByToken(session.Value); err == nil {
					if r.Method == http.MethodGet {
						http.Redirect(w, r, "https://"+r.URL.Host, 302)
						return
					}
					http.Error(w, "you has login", 403)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
