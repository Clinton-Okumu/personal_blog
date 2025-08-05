package routes

import (
	"personal_blog/backend/internal/app"

	"github.com/go-chi/chi/v5"
)

func ArticleRoutes(app *app.Application) chi.Router {
	r := chi.NewRouter()
	r.Post("/", app.ArticleHandler.CreateArticle)
	r.Get("/{id}", app.ArticleHandler.GetArticleByID)
	r.Put("/{id}", app.ArticleHandler.UpdateArticle)
	r.Delete("/{id}", app.ArticleHandler.DeleteArticle)
	return r
}
