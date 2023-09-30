package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"api/internal/db"
	"api/internal/service/auth"
	authPb "api/proto/auth"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	db := db.Connect("127.0.0.1", 3306, "root", "root", "serentin_db")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	gRPCServer := grpc.NewServer()
	authServer := auth.NewAuthServer(db)
	authPb.RegisterAuthServiceServer(gRPCServer, authServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
