package main

import (
	"flag"
	"log"
	"time"

	authPb "api/proto/auth"
	userPb "api/proto/user"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func TryRegister(c authPb.AuthServiceClient, ctx context.Context) {
	r, err := c.Register(ctx, &authPb.RegisterRequest{
		Email:         "kholidbughowi@gmail.com",
		Fullname:      "Mohamad kholid bughowi",
		Password:      "Yhf7763hfe",
		Birthdate:     "2001-10-15",
		Gender:        "M",
		Address:       "ITS Student dormitory",
		IdMbti:        3,
		IdJobPosition: 1,
	})
	if err != nil {
		log.Fatalf("Could not register user: %v", err)
	}
	log.Printf("UserId %s", r.GetUserId())
}

func TryLogin(c authPb.AuthServiceClient, ctx context.Context) string {
	res, err := c.Login(ctx, &authPb.LoginRequest{
		Email:    "kholidbughowi@gmail.com",
		Password: "Yhf7763hfe",
	})
	if err != nil {
		log.Fatalf("Could not login user: %v", err)
	}
	log.Printf("Token %s", res.GetToken())
	return res.GetToken()
}

func GetUserProfile(c userPb.UserServiceClient, ctx context.Context) {
	res, err := c.Profile(ctx, &userPb.UserId{
		IdUser: "6baa6634-1e28-4e89-a36a-aa31c32ff482",
	})
	if err != nil {
		log.Fatalf("Could not get user profile: %v", err)
	}
	log.Println(res)
}

func TryUpdateProfile(c userPb.UserServiceClient, ctx context.Context) {
	res, err := c.Update(ctx, &userPb.UpdateRequest{
		IdUser: "6baa6634-1e28-4e89-a36a-aa31c32ff482",
		UserData: &userPb.EditableData{
			Fullname:      "Mohamad Kholid Bughowi",
			Email:         "kholidbughowi@gmail.com",
			Birthdate:     "2001-10-15",
			Address:       "Informatics ITS",
			IdMbti:        1,
			IdJobPosition: 1,
		},
	})
	if err != nil {
		log.Fatalf("Could not update user profile: %v", err)
	}
	log.Println(res)
}

func TryChangePassword(c userPb.UserServiceClient, ctx context.Context) {
	res, err := c.ChangePassword(ctx, &userPb.ChangePasswordRequest{
		IdUser:      "6baa6634-1e28-4e89-a36a-aa31c32ff482",
		OldPassword: "Yhf7763hfe",
		NewPassword: "Yeye2y12345",
	})
	if err != nil {
		log.Fatalf("Could not change user password: %v", err)
	}
	log.Println(res)
}

func TryDeleteProfile(c userPb.UserServiceClient, ctx context.Context) {
	res, err := c.Delete(ctx, &userPb.UserId{
		IdUser: "6baa6634-1e28-4e89-a36a-aa31c32ff482",
	})
	if err != nil {
		log.Fatalf("Could not delete user profile: %v", err)
	}
	log.Println(res)
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	authClient := authPb.NewAuthServiceClient(conn)
	userClient := userPb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	TryRegister(authClient, ctx)
	token := TryLogin(authClient, ctx)
	md := metadata.Pairs("authorization", token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	GetUserProfile(userClient, ctx)
	TryUpdateProfile(userClient, ctx)
	TryChangePassword(userClient, ctx)
	TryDeleteProfile(userClient, ctx)
}
