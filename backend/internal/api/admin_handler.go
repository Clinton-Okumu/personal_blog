package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"personal_blog/backend/internal/models"
	"personal_blog/backend/internal/store"
	"personal_blog/backend/internal/utils"
	"strings"
)

type AdminHandler struct {
	adminStore store.AdminStore
	logger     *log.Logger
}

func NewAdminHandler(store store.AdminStore, logger *log.Logger) *AdminHandler {
	return &AdminHandler{
		adminStore: store,
		logger:     logger,
	}
}

type adminRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AdminHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input adminRegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}
	input.Username = strings.TrimSpace(input.Username)
	input.Password = strings.TrimSpace(input.Password)

	if input.Username == "" || input.Password == "" {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{
			"error": "username and password are required",
		})
		return
	}
	admin := &models.Admin{
		Username: input.Username,
		Password: input.Password,
	}
	if err := admin.SetPassword(input.Password); err != nil {
		h.logger.Printf("ERROR: hashing password: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal error"})
		return
	}

	_, err := h.adminStore.GetAdminByUsername(r.Context(), admin.Username)

	if err == nil {
		utils.WriteJSON(w, http.StatusConflict, utils.Envelope{"error": "username already taken"})
		return
	}

	if !errors.Is(err, store.ErrAdminNotFound) {
		h.logger.Printf("ERROR: checking for existing admin: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal error"})
		return
	}

	err = h.adminStore.CreateAdmin(r.Context(), admin)
	if err != nil {
		h.logger.Printf("ERROR: creating admin: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not create admin"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{
		"admin": map[string]any{
			"id":       admin.ID,
			"username": admin.Username,
		},
	})
}
