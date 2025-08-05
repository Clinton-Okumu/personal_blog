package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"personal_blog/backend/internal/models"
	"personal_blog/backend/internal/store"
	"personal_blog/backend/internal/utils"
)

type ArticleHandler struct {
	articleStore store.ArticleStore
	logger       *log.Logger
}

func NewArticleHandler(articleStore store.ArticleStore, logger *log.Logger) *ArticleHandler {
	return &ArticleHandler{
		articleStore: articleStore,
		logger:       logger,
	}
}

func (ah *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		ah.logger.Printf("ERROR: decoding article: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request sent"})
		return
	}

	createArticle, err := ah.articleStore.CreateArticle(r.Context(), &article)
	if err != nil {
		ah.logger.Printf("ERROR: creating article: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not create article"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"article": createArticle})
}

func (ah *ArticleHandler) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid article id"})
		return
	}

	booking, err := ah.articleStore.GetArticleByID(r.Context(), uint(id))
	if err != nil {
		if errors.Is(err, store.ErrArticleNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "article not found"})
			return
		}
		ah.logger.Printf("ERROR: fetching article: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not retrieve article"})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"article": booking})
}

func (ah *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid article id"})
		return
	}

	var article models.Article
	err = json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		ah.logger.Printf("ERROR: decoding article: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request sent"})
		return
	}

	article.ID = uint(id)
	err = ah.articleStore.UpdateArticle(r.Context(), &article)
	if err != nil {
		ah.logger.Printf("ERROR: updating article: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not update article"})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "article updated"})
}

func (ah *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid article id"})
		return
	}

	err = ah.articleStore.DeleteArticle(r.Context(), uint(id))
	if err != nil {
		ah.logger.Printf("ERROR: deleting article: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not delete article"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "article deleted"})
}
