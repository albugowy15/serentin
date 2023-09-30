package auth

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Claims struct {
	ID    string
	Email string
	jwt.RegisteredClaims
}

func CreateJWTToken(secret string, claims *Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		log.Printf("Error signing token: %v\n", err)
	}
	return tokenStr
}

func DecodeJWTToken(tokenStr string, secret string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		jwtKey := []byte(secret)
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, status.Error(codes.Unauthenticated, "Unauthorized")
		}
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}
	if !token.Valid {
		return nil, status.Error(codes.Unauthenticated, "Unathorized")
	}
	return claims, nil
}
