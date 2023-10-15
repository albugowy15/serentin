package user

import (
	"api/internal/db"
	"api/pkg/session"
	"api/pkg/validator"
	pb "api/proto"
	"context"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	Fullname    string
	Email       string
	Gender      string
	Birthdate   string
	Address     string
	Personality string
	Position    string
}

type UserServiceServer struct {
	db db.Database
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer(db *db.Database) *UserServiceServer {
	return &UserServiceServer{
		db: *db,
	}
}

func (s *UserServiceServer) Profile(ctx context.Context, req *pb.UserRequest) (*pb.ProfileResponse, error) {
	// validate user id
	session := session.GetSession(ctx)
	user := &User{}
	err := s.db.Conn.Get(
		user,
		`SELECT u.fullname, u.email, u.gender, u.birthdate,
		u.address, m.personality, j.position 
		FROM users u INNER JOIN mbti m ON u.id_mbti = m.id 
		INNER JOIN job_positions j ON j.id = u.id_job_position 
		WHERE u.id=?`,
		session.ID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		log.Printf("Error get user with id: %d :%v\n", session.ID, err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	genderMap := map[string]string{
		"M": "Male",
		"F": "Female",
	}
	gender, ok := genderMap[user.Gender]
	if !ok {
		log.Fatal("Unrecognized gender type")
	}

	return &pb.ProfileResponse{
		Fullname:    user.Fullname,
		Email:       user.Email,
		Gender:      gender,
		Birthdate:   user.Birthdate,
		Address:     user.Address,
		Personality: user.Personality,
		Position:    user.Position,
	}, nil
}

func (s *UserServiceServer) Update(ctx context.Context, req *pb.EditableData) (*pb.MessageResponse, error) {
	session := session.GetSession(ctx)
	// validate request
	if err := validator.ValidateFullName(req.GetFullname()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid fullname: %v", err)
	}
	if err := validator.ValidateEmail(req.GetEmail()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid email: %v", err)
	}
	if err := validator.ValidateBirtdate(req.GetBirthdate()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid birthdate: %v", err)
	}
	// check idMti
	var idMbti int64
	if err := s.db.Conn.Get(&idMbti, "SELECT id from mbti WHERE id=?", req.GetIdMbti()); err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.InvalidArgument, "id_mbti not found")
		}
		log.Printf("Error get mbti with id %d: %v\n", req.GetIdMbti(), err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	// check idJobPosition
	var idJobPosition int64
	if err := s.db.Conn.Get(&idJobPosition, "SELECT id from job_positions WHERE id=?", req.GetIdJobPosition()); err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.InvalidArgument, "id_job_position not found")
		}
		log.Printf("Error get job position with id %d: %v\n", req.GetIdJobPosition(), err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	updateUserStmt := `UPDATE users SET fullname=?, email=?, birthdate=?, address=?, id_mbti=?, id_job_position=? WHERE id=?`
	if _, err := s.db.Conn.Exec(
		updateUserStmt,
		req.GetFullname(),
		req.GetEmail(),
		req.GetBirthdate(),
		req.GetAddress(),
		req.GetIdMbti(),
		req.GetIdJobPosition(),
		session.ID,
	); err != nil {
		log.Printf("Erorr update user %d: %v\n", session.ID, err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &pb.MessageResponse{
		IdUser:  session.ID,
		Message: "successfully update user",
	}, nil
}

func (s *UserServiceServer) Delete(ctx context.Context, req *pb.UserRequest) (*pb.MessageResponse, error) {
	session := session.GetSession(ctx)
	result := s.db.Conn.MustExec("DELETE FROM users WHERE id=?", session.ID)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Erorr delete user %d: %v\n", session.ID, err)
		return nil, status.Error(codes.Internal, "internal error")
	}
	if rowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &pb.MessageResponse{
		IdUser:  session.ID,
		Message: "successfully delete user",
	}, nil
}

func (s *UserServiceServer) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.MessageResponse, error) {
	session := session.GetSession(ctx)
	if err := validator.ValidatePassword(req.GetNewPassword()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid password: %v", err)
	}

	var oldPassword string
	err := s.db.Conn.Get(&oldPassword, "SELECT password FROM users WHERE id=?", session.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		log.Printf("Error change password for user %d: %v\n", session.ID, err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(oldPassword), []byte(req.GetOldPassword())); err != nil {
		log.Printf("Error compare password: %v\n", err)
		return nil, status.Error(codes.InvalidArgument, "Old Password not match")
	}

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetNewPassword()), 10)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	updatePasswordStmt := `UPDATE users SET password=? WHERE id=?`
	result := s.db.Conn.MustExec(updatePasswordStmt, hashedNewPassword, session.ID)
	if _, err := result.RowsAffected(); err != nil {
		log.Printf("Erorr update password for user %d: %v\n", session.ID, err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &pb.MessageResponse{
		IdUser:  session.ID,
		Message: "successfully change password",
	}, nil
}
