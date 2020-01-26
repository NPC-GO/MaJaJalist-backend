package middleware

import (
	"github.com/NPC-GO/MaJaJalist-backend/database"
	"net/http"
)

func BeforeLoginAuth(user database.User) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := r.Cookie("session")
			if err == nil {
				if _, err := user.GetUserByToken(session.Value); err == nil {
					http.Redirect(w, r, "https://"+r.URL.Host, 302)
				}
				http.SetCookie(w, &http.Cookie{Name: "session", MaxAge: -1})
			}
			next.ServeHTTP(w, r)
		})
	}
}
