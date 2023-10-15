package auth_test

import (
	"api/internal/service/auth"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createToken() string {
	claims := &auth.Claims{
		ID:    3004,
		Email: "tes@gmail.com",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	secret := "the-most-strong-secret"
	tokenStr := auth.CreateJWTToken(secret, claims)
	return tokenStr
}

func TestCreateJWTToken(t *testing.T) {
	token := createToken()
	if len(token) != 0 {
		t.Logf("token: %s\n", token)
	} else {
		t.Error("Error signing token")
	}
}

func TestDecodeJWTToken(t *testing.T) {
	// create token
	token := createToken()
	secret := "the-most-strong-secret"
	claims, err := auth.DecodeJWTToken(token, secret)
	t.Logf("claims: %v\n", claims)
	if err != nil {
		t.Errorf("Error decode token: %v\n", err)
	}
}
