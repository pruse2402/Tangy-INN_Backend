package route

import (
	"tangy-inn/tangy-inn_backend/handlers"

	"github.com/go-chi/chi"
)

func NewRouter(h *handlers.Provider) *chi.Mux {

	Router := chi.NewRouter()

	Router.Route("/v1", func(r chi.Router) {

		r.Group(func(r chi.Router) {
			r.Mount("/", Routes(h))
		})

		r.Get("/ping", h.Ping)

	})

	return Router
}
