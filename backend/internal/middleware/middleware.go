package middleware

import (
	"context"
	"net/http"
	"personal_blog/backend/internal/models"
	"personal_blog/backend/internal/store"
	"personal_blog/backend/internal/tokens"
	"personal_blog/backend/internal/utils"
	"strings"
)

type AdminMiddleware struct {
	AdminStore store.AdminStore
	TokenStore store.TokenStore
}

type contextKey string

const userContextKey = contextKey("admin")

func SetAdmin(r *http.Request, admin *models.Admin) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, admin)
	return r.WithContext(ctx)
}

func GetAdmin(r *http.Request) *models.Admin {
	user, ok := r.Context().Value(userContextKey).(*models.Admin)
	if !ok {
		panic("missing user in request context")
	}
	return user
}

func (am *AdminMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "Missing authorization header"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "Invalid authorization header format"})
			return
		}

		token := parts[1]
		admin, err := am.TokenStore.GetAdminByToken(r.Context(), token, tokens.ScopeAuth)
		if err != nil || admin == nil {
			utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "Invalid or expired token"})
			return
		}

		r = SetAdmin(r, admin)
		next.ServeHTTP(w, r)
	})
}
