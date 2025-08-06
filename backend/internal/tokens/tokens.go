package tokens

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"personal_blog/backend/internal/models"
	"time"
)

const (
	ScopeAuth          = "auth"
	ScopeResetPassword = "reset_password"
	TokenLength        = 16
)

func GenerateToken(adminID uint, ttl time.Duration, scope string) (*models.Token, string, error) {
	b := make([]byte, TokenLength)
	_, err := rand.Read(b)
	if err != nil {
		return nil, "", fmt.Errorf("could not generate random bytes: %w", err)
	}
	plaintext := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b)

	hash := sha256.Sum256([]byte(plaintext))

	token := &models.Token{
		AdminID: adminID,
		Hash:    hash[:],
		Scope:   scope,
		Expiry:  time.Now().Add(ttl),
	}

	return token, plaintext, nil
}

func HashToken(token string) []byte {
	hash := sha256.Sum256([]byte(token))
	return hash[:]
}
