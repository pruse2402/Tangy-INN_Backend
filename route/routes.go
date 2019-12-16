package route

import (
	"tangy-inn/tangy-inn_backend/handlers"

	"github.com/go-chi/chi"
)

func Routes(h *handlers.Provider) chi.Router {
	r := chi.NewRouter()
	return r
}
