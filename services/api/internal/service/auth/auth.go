package auth

import (
	"api/configs"
	"api/internal/db"
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

func (s *AuthServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if err := ValidateRegister(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	var id string
	err := s.db.Conn.Get(&id, "SELECT id FROM users WHERE email=?", req.Email)
	if err == nil {
		return nil, status.Error(codes.InvalidArgument, "email already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	userStmt := `INSERT INTO users 
				(id, email, fullname, password, gender, birthdate, address, role, id_mbti, id_job_position) 
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	userId := uuid.New().String()
	res := s.db.Conn.MustExec(userStmt,
		userId,
		req.Email,
		req.Fullname,
		hashedPassword,
		req.Gender,
		req.Birthdate,
		req.Address,
		"employee",
		req.IdMbti,
		req.IdJobPosition,
	)
	if rowsCount, err := res.RowsAffected(); rowsCount == 0 || err != nil {
		log.Printf("Rows affected: %v\n", rowsCount)
		log.Printf("Error inserting user: %v\n", err)
		return nil, status.Error(codes.Internal, "internal error")
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
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.Internal, "Email %s not found", req.Email)
		}
		log.Println(err)
		return nil, status.Error(codes.Internal, "internal error")
	}
	// check hash match password
	if err := bcrypt.CompareHashAndPassword([]byte(userCredential.Password), []byte(req.Password)); err != nil {
		log.Printf("Error compare password %v\n", err)
		return nil, status.Error(codes.InvalidArgument, "Password not match")
	}

	// create jwt
	tokenClaims := &Claims{
		Email: userCredential.Email,
		ID:    userCredential.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	jwtAccessKey := configs.ViperEnvVariable("JWT_ACCESS_TOKEN_SECRET")
	tokenStr := CreateJWTToken(jwtAccessKey, tokenClaims)

	refreshTokenClaims := &Claims{
		Email: userCredential.Email,
		ID:    userCredential.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	jwtRefreshKey := configs.ViperEnvVariable("JWT_REFRESH_TOKEN_SECRET")
	refreshTokenStr := CreateJWTToken(jwtRefreshKey, refreshTokenClaims)

	// send jwt
	return &pb.LoginResponse{
		Token:        tokenStr,
		RefreshToken: refreshTokenStr,
	}, nil
}

func (s *AuthServiceServer) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.LoginResponse, error) {
	tokenStr := req.GetRefreshToken()
	jwtRefreshKey := configs.ViperEnvVariable("JWT_REFRESH_TOKEN_SECRET")
	claims, err := DecodeJWTToken(tokenStr, jwtRefreshKey)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		return nil, status.Error(codes.InvalidArgument, "Refresh token expired")
	}

	newTokenClaims := &Claims{
		Email: claims.Email,
		ID:    claims.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	jwtAccessKey := configs.ViperEnvVariable("JWT_ACCESS_TOKEN_SECRET")
	newTokenStr := CreateJWTToken(jwtAccessKey, newTokenClaims)

	refreshTokenClaims := &Claims{
		Email: claims.Email,
		ID:    claims.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	refreshTokenStr := CreateJWTToken(jwtRefreshKey, refreshTokenClaims)

	return &pb.LoginResponse{
		Token:        newTokenStr,
		RefreshToken: refreshTokenStr,
	}, nil
}
