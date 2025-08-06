package store

import (
	"context"
	"personal_blog/backend/internal/models"
	"personal_blog/backend/internal/tokens"
	"time"

	"gorm.io/gorm"
)

type TokenStore interface {
	CreateNewToken(ctx context.Context, adminID uint, ttl time.Duration, scope string) (*models.Token, string, error)
	DeleteAllTokensForAdmin(ctx context.Context, adminID uint, scope string) error
	GetAdminByToken(ctx context.Context, tokenPlaintext string, scope string) (*models.Admin, error)
}

type tokenStore struct {
	db *gorm.DB
}

func NewTokenStore(db *gorm.DB) TokenStore {
	return &tokenStore{db}
}

func (ts *tokenStore) CreateNewToken(ctx context.Context, adminID uint, ttl time.Duration, scope string) (*models.Token, string, error) {
	token, plaintext, err := tokens.GenerateToken(adminID, ttl, scope)
	if err != nil {
		return nil, "", err
	}
	if err := ts.db.WithContext(ctx).Create(token).Error; err != nil {
		return nil, "", err
	}
	return token, plaintext, nil
}

func (ts *tokenStore) DeleteAllTokensForAdmin(ctx context.Context, adminID uint, scope string) error {
	return ts.db.WithContext(ctx).
		Where("user_id = ? AND scope = ?", adminID, scope).
		Delete(&models.Token{}).Error
}

func (ts *tokenStore) GetAdminByToken(ctx context.Context, tokenPlaintext string, scope string) (*models.Admin, error) {
	hash := tokens.HashToken(tokenPlaintext)

	var token models.Token
	if err := ts.db.WithContext(ctx).
		Where("hash = ? AND scope = ? AND expiry > ?", hash, scope, time.Now()).
		First(&token).Error; err != nil {
		return nil, err
	}

	var admin models.Admin
	if err := ts.db.WithContext(ctx).
		First(&admin, token.AdminID).Error; err != nil {
		return nil, err
	}

	return &admin, nil
}
