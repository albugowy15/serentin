package auth

import (
	"api/pkg/validator"
	pb "api/proto"
	"fmt"
)

func ValidateRegister(req *pb.RegisterRequest) error {
	if err := validator.ValidateEmail(req.Email); err != nil {
		return fmt.Errorf("invalid email: %v", err)
	}
	if err := validator.ValidatePassword(req.Password); err != nil {
		return fmt.Errorf("invalid password: %v", err)
	}
	if err := validator.ValidateFullName(req.Fullname); err != nil {
		return fmt.Errorf("invalid full name: %v", err)
	}
	if err := validator.ValidateBirtdate(req.Birthdate); err != nil {
		return fmt.Errorf("invalid birthdate: %v", err)
	}
	return nil
}
