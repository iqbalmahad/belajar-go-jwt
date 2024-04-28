package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndVerifyJWTToken(t *testing.T) {
	secretKey := []byte("your_secret_key")
	username := "testuser"

	// Generate token
	token, err := GenerateJWTToken(username, secretKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify token
	claims, err := VerifyJWTToken(token, secretKey)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, username, claims.Username)
}
