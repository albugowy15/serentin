package auth

import (
	"api/pkg/validator"
	pb "api/proto/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateRegister(req *pb.RegisterRequest) error {
	if err := validator.ValidateEmail(req.Email); err != nil {
		status.Errorf(codes.InvalidArgument, "email: %v", err)
	}
	if err := validator.ValidatePassword(req.Password); err != nil {
		status.Errorf(codes.InvalidArgument, "password: %v", err)
	}
	if err := validator.ValidateFullName(req.Fullname); err != nil {
		status.Errorf(codes.InvalidArgument, "fullname: %v", err)
	}
	if err := validator.ValidateBirtdate(req.Birthdate); err != nil {
		status.Errorf(codes.InvalidArgument, "birthdate: %v", err)
	}
	return nil
}
