package session

import (
	"api/internal/service/auth"
	"context"
	"os"

	"google.golang.org/grpc/metadata"
)

func GetSession(ctx context.Context) *auth.Claims {
	md, _ := metadata.FromIncomingContext(ctx)
	token := md.Get("authorization")[0]
	secret := os.Getenv("JWT_ACCESS_TOKEN_SECRET")
	claims, _ := auth.DecodeJWTToken(token, secret)
	return claims
}
