package routes

import (
	"personal_blog/backend/internal/app"

	"github.com/go-chi/chi/v5"
)

func AdminRoutes(app *app.Application) chi.Router {
	r := chi.NewRouter()
	r.Post("/register", app.AdminHandler.Register)
	r.Post("/login", app.TokenHandler.HandleCreateToken)

	return r
}
