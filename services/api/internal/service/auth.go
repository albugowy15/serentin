package service

import (
	"api/internal/db"
	"api/pkg/validator"
	pb "api/proto/auth"
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServiceServer struct {
	db db.Database

	*pb.UnimplementedAuthServiceServer
}

func NewAuthServer(db *db.Database) *AuthServiceServer {
	return &AuthServiceServer{
		db: *db,
	}
}

func (s *AuthServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// validation
	if err := validator.ValidateEmail(req.Email); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "email: %v", err)
	}
	if err := validator.ValidatePassword(req.Password); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "password: %v", err)
	}
	if err := validator.ValidateFullName(req.Fullname); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "fullname: %v", err)
	}
	birthDateinDate, err := time.Parse("2006-01-02", req.Birthdate)
	if err != nil {
		log.Fatalf("Error parsing date: %v\n", err)
	}
	if err := validator.ValidateBirtdate(&birthDateinDate); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "birthdate: %v", err)
	}

	// TODO: check if email already exists
	row := s.db.Conn.QueryRow("SELECT id FROM users WHERE email=?", req.Email)
	var id string
	err = row.Scan(&id)
	if err != sql.ErrNoRows {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "email already exist")
	}

	// TODO: hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	// TODO: save user to database
	userStmt := `INSERT INTO users (id, email, fullname, password, gender, birthdate, address, role) 
				VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	userId := uuid.New().String()
	_, err = s.db.Conn.Exec(userStmt, userId, req.Email, req.Fullname, hashedPassword, req.Gender, req.Birthdate, req.Address, "employee")
	if err != nil {
		log.Printf("Error inserting user: %v\n", err)
	}

	// TODO: send user id created
	return &pb.RegisterResponse{UserId: userId}, nil
}

func (s *AuthServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func (s *AuthServiceServer) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
