package service

import (
	"api/configs"
	"api/internal/db"
	"api/pkg/validator"
	pb "api/proto/auth"
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func validateRegister(req *pb.RegisterRequest) error {
	if err := validator.ValidateEmail(req.Email); err != nil {
		status.Errorf(codes.InvalidArgument, "email: %v", err)
	}
	if err := validator.ValidatePassword(req.Password); err != nil {
		status.Errorf(codes.InvalidArgument, "password: %v", err)
	}
	if err := validator.ValidateFullName(req.Fullname); err != nil {
		status.Errorf(codes.InvalidArgument, "fullname: %v", err)
	}
	birthDateinDate, err := time.Parse("2006-01-02", req.Birthdate)
	if err != nil {
		log.Fatalf("Error parsing date: %v\n", err)
	}
	if err := validator.ValidateBirtdate(&birthDateinDate); err != nil {
		status.Errorf(codes.InvalidArgument, "birthdate: %v", err)
	}
	return nil
}

func (s *AuthServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if err := validateRegister(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation fails: %v", err)
	}
	row := s.db.Conn.QueryRow("SELECT id FROM users WHERE email=?", req.Email)
	var id string
	err := row.Scan(&id)
	if err != sql.ErrNoRows {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, "email already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	userStmt := `INSERT INTO users (id, email, fullname, password, gender, birthdate, address, role) 
				VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	userId := uuid.New().String()
	_, err = s.db.Conn.Exec(userStmt, userId, req.Email, req.Fullname, hashedPassword, req.Gender, req.Birthdate, req.Address, "employee")
	if err != nil {
		log.Printf("Error inserting user: %v\n", err)
	}

	return &pb.RegisterResponse{UserId: userId}, nil
}

func (s *AuthServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// get user id and password from email
	type UserCredentials struct {
		ID       string `db:"id"`
		Email    string `db:"email"`
		Password string `db:"password"`
	}
	userCredential := UserCredentials{}
	err := s.db.Conn.Get(&userCredential,
		"SELECT id, email, password FROM users WHERE email=?",
		req.Email)
	if err != nil {
		log.Println(err)
	}
	// check hash match password
	if err := bcrypt.CompareHashAndPassword([]byte(userCredential.Password), []byte(req.Password)); err != nil {
		log.Printf("Error compare password %v\n", err)
		return nil, status.Error(codes.InvalidArgument, "Password not match")
	}

	// create jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    userCredential.ID,
		"email": userCredential.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now().Unix(),
	})
	jwtKey := configs.ViperEnvVariable("JWT_SECRET")
	tokenStr, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		log.Printf("Error signing token: %v\n", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	// send jwt
	return &pb.LoginResponse{
		Token:        tokenStr,
		RefreshToken: tokenStr,
	}, nil
}

func (s *AuthServiceServer) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
