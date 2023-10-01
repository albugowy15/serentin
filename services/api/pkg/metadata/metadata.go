package metadata

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

func GetToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("metadata invalid")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return "", fmt.Errorf("authorization token is not supplied")
	}
	token := authHeader[0]
	return token, nil
}
