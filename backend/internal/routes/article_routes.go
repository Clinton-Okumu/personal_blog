package routes

import (
	"personal_blog/backend/internal/app"

	"github.com/go-chi/chi/v5"
)

func ArticleRoutes(app *app.Application) chi.Router {
	r := chi.NewRouter()

	// Public: only GET
	r.Get("/{id}", app.ArticleHandler.GetArticleByID)

	// Protected: POST, PUT, DELETE
	r.Group(func(protected chi.Router) {
		protected.Use(app.Middleware.Authenticate)

		protected.Post("/", app.ArticleHandler.CreateArticle)
		protected.Put("/{id}", app.ArticleHandler.UpdateArticle)
		protected.Delete("/{id}", app.ArticleHandler.DeleteArticle)
	})

	return r
}
