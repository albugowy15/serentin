package main

import (
	"flag"
	"log"
	"time"

	authPb "api/proto/auth"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := authPb.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// test register
	r, err := c.Register(ctx, &authPb.RegisterRequest{
		Email:       "kholidbughowi@gmail.com",
		Fullname:    "Mohamad kholid bughowi",
		Password:    "Yhf7763hfe",
		Birthdate:   "2001-10-15",
		Gender:      "M",
		Address:     "ITS Student dormitory",
		MbtiType:    "INTJ",
		JobPosition: "Software Developer",
	})
	if err != nil {
		log.Fatalf("Could not register user: %v", err)
	}
	log.Printf("UserId %s", r.GetUserId())

	// test login
	res, ok := c.Login(ctx, &authPb.LoginRequest{
		Email:    "kholidbughowi@gmail.com",
		Password: "Yhf7763hfe",
	})
	if ok != nil {
		log.Fatalf("Could not login user: %v", err)
	}
	log.Printf("Token %s", res.GetToken())
}
