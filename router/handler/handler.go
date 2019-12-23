package handler

import (
	"net/http"
)

func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./dist/index.html")
}
