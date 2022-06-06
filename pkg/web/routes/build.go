package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func BuildList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all builds"))
}

func BuildFetch(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get build #" + chi.URLParam(r, "id")))
}

func BuildStart(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("start a new build"))
}
