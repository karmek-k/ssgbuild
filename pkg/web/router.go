package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/karmek-k/ssgbuild/pkg/web/routes"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/build", routes.BuildList)
	r.Get("/build/{id}", routes.BuildFetch)
	r.Post("/build", routes.BuildStart)

	return r
}
