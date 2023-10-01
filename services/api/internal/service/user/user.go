package user

import (
	"api/internal/db"
	pb "api/proto/user"
	"context"
	"database/sql"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	db db.Database
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer(db *db.Database) *UserServiceServer {
	return &UserServiceServer{
		db: *db,
	}
}

func (s *UserServiceServer) Profile(ctx context.Context, req *pb.ProfileRequest) (*pb.UserData, error) {
	user := &User{}
	err := s.db.Conn.Get(
		user,
		`SELECT u.fullname, u.email, u.gender, u.birthdate, u.role,
		u.address, m.personality, j.position AS job_position 
		FROM users u INNER JOIN mbti m ON u.id_mbti = m.id 
		INNER JOIN job_positions j ON j.id = u.id_job_position 
		WHERE u.id=?`,
		req.GetIdUser(),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		log.Printf("Error get user with id: %s :%v\n", req.GetIdUser(), err)
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

	return &pb.UserData{
		Fullname:    user.Fullname,
		Email:       user.Email,
		Gender:      gender,
		Birthdate:   user.Birthdate,
		Address:     user.Address,
		Role:        user.Role,
		Personality: user.Personality,
		JobPosition: user.JobPosition,
	}, nil
}
func (s *UserServiceServer) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (s *UserServiceServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (s *UserServiceServer) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}

type User struct {
	Fullname    string
	Email       string
	Gender      string
	Birthdate   string
	Role        string
	Address     string
	Personality string
	JobPosition string
}
