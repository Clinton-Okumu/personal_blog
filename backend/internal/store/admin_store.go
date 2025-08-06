package store

import (
	"context"
	"errors"
	"personal_blog/backend/internal/models"

	"gorm.io/gorm"
)

var ErrAdminNotFound = errors.New("admin not found")

type AdminStore interface {
	CreateAdmin(ctx context.Context, admin *models.Admin) error
	GetAdminByUsername(ctx context.Context, username string) (*models.Admin, error)
}

type adminStore struct {
	db *gorm.DB
}

func NewAdminStore(db *gorm.DB) AdminStore {
	return &adminStore{db}
}

func (as *adminStore) CreateAdmin(ctx context.Context, admin *models.Admin) error {
	return as.db.WithContext(ctx).Create(admin).Error
}

func (as *adminStore) GetAdminByUsername(ctx context.Context, username string) (*models.Admin, error) {
	var admin models.Admin
	err := as.db.WithContext(ctx).
		Where("username = ?", username).
		First(&admin).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrAdminNotFound
	}
	return &admin, err
}
