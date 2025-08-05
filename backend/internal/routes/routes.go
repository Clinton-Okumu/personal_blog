package routes

import (
	"personal_blog/backend/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetUpRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", app.Welcome)
	r.Get("/health", app.HealthChecker)
	r.Route("/articles", func(r chi.Router) {
		r.Post("/", app.ArticleHandler.CreateArticle)
		r.Get("/{id}", app.ArticleHandler.GetArticleByID)
		r.Put("/{id}", app.ArticleHandler.UpdateArticle)
		r.Delete("/{id}", app.ArticleHandler.DeleteArticle)
	})

	return r
}
