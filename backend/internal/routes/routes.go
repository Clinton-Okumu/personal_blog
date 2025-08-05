package routes

import (
	"personal_blog/backend/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetUpRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", app.Welcome)
	r.Get("/health", app.HealthChecker)

	// Mount feature-specific routes
	r.Mount("/articles", ArticleRoutes(app))
	// r.Mount("/admin", AdminRoutes(app))

	return r
}
