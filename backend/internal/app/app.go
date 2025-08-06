package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"personal_blog/backend/internal/api"
	"personal_blog/backend/internal/middleware"
	"personal_blog/backend/internal/store"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Application struct {
	DB             *gorm.DB
	Logger         *log.Logger
	ArticleHandler *api.ArticleHandler
	AdminHandler   *api.AdminHandler
	TokenHandler   *api.TokenHandler
	Middleware     middleware.AdminMiddleware
}

func NewApplication() (*Application, error) {
	// load environment variables
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	// database connection
	db, err := store.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// automigrate models
	if err := store.RunMigrations(db); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	// setup logger
	logger := log.New(os.Stdout, "[Blog]", log.Ldate|log.Ltime)

	// stores
	articleStore := store.NewArticleStore(db)
	adminStore := store.NewAdminStore(db)
	tokenStore := store.NewTokenStore(db)

	// handlers
	articleHandler := api.NewArticleHandler(articleStore, logger)
	adminHandler := api.NewAdminHandler(adminStore, logger)
	tokenHandler := api.NewTokenHandler(tokenStore, adminStore, logger)
	middlwareHandler := middleware.AdminMiddleware{AdminStore: adminStore}

	app := &Application{
		DB:             db,
		Logger:         logger,
		ArticleHandler: articleHandler,
		AdminHandler:   adminHandler,
		TokenHandler:   tokenHandler,
		Middleware:     middlwareHandler,
	}
	return app, nil
}

func (a *Application) HealthChecker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}

func (a *Application) Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my personal blog API")
}
