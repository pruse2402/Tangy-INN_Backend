package route

import (
	"tangy-inn/tangy-inn_backend/handlers"

	"github.com/go-chi/chi"
)

func Routes(h *handlers.Provider) chi.Router {
	r := chi.NewRouter()
	// Login API
	r.Post("/login", h.Authenticate)

	// MenuItem API
	r.Get("/store/{id}/menuItems", h.GetMenuItemList)
	r.Post("/menuItem", h.InsertMenuItem)
	r.Put("/menuItem/{id}", h.UpdateMenuItem)

	return r
}
