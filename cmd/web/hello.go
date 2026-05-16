package web

import (
	"net/http"
)

func HelloWebHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "no endpoint yet", http.StatusNotFound)
}
