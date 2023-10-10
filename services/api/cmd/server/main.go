package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"api/internal/db"
	"api/internal/service/auth"
	"api/internal/service/common"
	"api/internal/service/logbook"
	"api/internal/service/user"
	pb "api/proto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	port          = flag.Int("port", 50051, "The server port")
	publicMethods = []string{
		"/api.AuthService",
		"/api.CommonService",
	}
)

func serverInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	isPublic := false
	for _, method := range publicMethods {
		if strings.HasPrefix(info.FullMethod, method) {
			isPublic = true
			break
		}
	}
	if !isPublic {
		if err := authorize(ctx); err != nil {
			log.Printf("Request - Method:%s\tDuration:%s\tError:%v\n", info.FullMethod, time.Since(start), err)
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
	secret := os.Getenv("JWT_ACCESS_TOKEN_SECRET")
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
	// load env and flag
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	flag.Parse()

	// initialize db
	dbHost := os.Getenv("DB_HOST")
	dbPortStr := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		log.Fatal(err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	db := db.Connect(dbHost, dbPort, dbUser, dbPass, dbName)
	defer db.Close()

	// open tcp connection
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
	defer lis.Close()

	// create gRPC server
	gRPCServer := grpc.NewServer(withServerUnaryInterceptor())
	authServer := auth.NewAuthServer(db)
	userServer := user.NewUserServiceServer(db)
	commonServer := common.NewCommonServer(db)
	logbookServer := logbook.NewLogbookServer(db)

	// register gRPC server
	pb.RegisterAuthServiceServer(gRPCServer, authServer)
	pb.RegisterUserServiceServer(gRPCServer, userServer)
	pb.RegisterCommonServiceServer(gRPCServer, commonServer)
	pb.RegisterLogbookServiceServer(gRPCServer, logbookServer)

	// start gRPC server
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	defer gRPCServer.Stop()
}
