package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"api/configs"
	"api/internal/db"
	"api/internal/service/auth"
	"api/internal/service/user"
	authPb "api/proto/auth"
	userPb "api/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func serverInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	// TODO : change the method
	if !strings.HasPrefix(info.FullMethod, "/auth") {
		if err := authorize(ctx); err != nil {
			return nil, err
		}
	}
	h, err := handler(ctx, req)
	log.Printf("Request - Method:%s\tDuration:%s\tError:%v\n", info.FullMethod, time.Since(start), err)

	return h, err
}

func authorize(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Metadata invalid")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}
	token := authHeader[0]
	secret := configs.ViperEnvVariable("JWT_ACCESS_TOKEN_SECRET")
	_, err := auth.DecodeJWTToken(token, secret)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}
	return nil
}

func withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

func main() {
	flag.Parse()

	dbHost := configs.ViperEnvVariable("DB_HOST")
	dbPortStr := configs.ViperEnvVariable("DB_PORT")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		log.Fatal(err)
	}
	dbUser := configs.ViperEnvVariable("DB_USER")
	dbPass := configs.ViperEnvVariable("DB_PASSWORD")
	dbName := configs.ViperEnvVariable("DB_DATABASE")
	db := db.Connect(dbHost, dbPort, dbUser, dbPass, dbName)
	defer db.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	gRPCServer := grpc.NewServer(withServerUnaryInterceptor())
	authServer := auth.NewAuthServer(db)
	userServer := user.NewUserServiceServer(db)
	authPb.RegisterAuthServiceServer(gRPCServer, authServer)
	userPb.RegisterUserServiceServer(gRPCServer, userServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
