package store

import (
	"context"
	"errors"
	"fmt"
	"personal_blog/backend/internal/models"

	"gorm.io/gorm"
)

var ErrArticleNotFound = errors.New("article not found")

type ArticleStore interface {
	CreateArticle(ctx context.Context, article *models.Article) (*models.Article, error)
	GetArticleByID(ctx context.Context, articleID uint) (*models.Article, error)
	UpdateArticle(ctx context.Context, article *models.Article) error
	DeleteArticle(ctx context.Context, articleID uint) error
}

type articleStore struct {
	db *gorm.DB
}

func NewArticleStore(db *gorm.DB) ArticleStore {
	return &articleStore{db}
}

func (as *articleStore) CreateArticle(ctx context.Context, article *models.Article) (*models.Article, error) {
	if err := as.db.WithContext(ctx).Create(article).Error; err != nil {
		return nil, fmt.Errorf("failed to create article: %w", err)
	}
	return article, nil
}

func (as *articleStore) GetArticleByID(ctx context.Context, articleID uint) (*models.Article, error) {
	var article models.Article
	if err := as.db.WithContext(ctx).First(&article, articleID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrArticleNotFound
		}
		return nil, fmt.Errorf("failed to get article by ID %d: %w", articleID, err)
	}
	return &article, nil
}

func (as *articleStore) UpdateArticle(ctx context.Context, article *models.Article) error {
	existingArticle, err := as.GetArticleByID(ctx, article.ID)
	if err != nil {
		return fmt.Errorf("failed to find article for update: %w", err)
	}

	if article.Title != "" {
		existingArticle.Title = article.Title
	}
	if article.Content != "" {
		existingArticle.Content = article.Content
	}

	if err := as.db.WithContext(ctx).Save(existingArticle).Error; err != nil {
		return fmt.Errorf("failed to save updated article: %w", err)
	}
	return nil
}

func (as *articleStore) DeleteArticle(ctx context.Context, articleID uint) error {
	result := as.db.WithContext(ctx).Delete(&models.Article{}, articleID)
	if result.Error != nil {
		return fmt.Errorf("failed to delete article: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return ErrArticleNotFound
	}
	return nil
}
