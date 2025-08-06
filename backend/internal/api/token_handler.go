package api

import (
	"encoding/json"
	"log"
	"net/http"
	"personal_blog/backend/internal/store"
	"personal_blog/backend/internal/tokens"
	"personal_blog/backend/internal/utils"
	"time"
)

type TokenHandler struct {
	tokenStore store.TokenStore
	adminStore store.AdminStore
	logger     *log.Logger
}

type createTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewTokenHandler(tokenStore store.TokenStore, adminStore store.AdminStore, logger *log.Logger) *TokenHandler {
	return &TokenHandler{
		tokenStore: tokenStore,
		adminStore: adminStore,
		logger:     logger,
	}
}

func (h *TokenHandler) HandleCreateToken(w http.ResponseWriter, r *http.Request) {
	var req createTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Printf("ERROR decoding createTokenRequest: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid request body"})
		return
	}

	admin, err := h.adminStore.GetAdminByUsername(r.Context(), req.Username)
	if err != nil {
		h.logger.Printf("ERROR finding admin: %v", err)
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "Invalid credentials"})
		return
	}
	match, err := admin.CheckPassword(req.Password)
	if err != nil || !match {
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "Invalid credentials"})
		return
	}
	tokenModel, plaintext, err := h.tokenStore.CreateNewToken(r.Context(), admin.ID, 24*time.Hour, tokens.ScopeAuth)
	if err != nil {
		h.logger.Printf("ERROR creating token: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Failed to generate token"})
		return
	}
	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{
		"token":  plaintext,
		"expiry": tokenModel.Expiry,
		"admin": map[string]any{
			"id":       admin.ID,
			"username": admin.Username,
		},
	})
}
