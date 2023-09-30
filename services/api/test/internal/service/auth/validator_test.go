package auth_test

import (
	"api/internal/service/auth"
	pb "api/proto/auth"
	"testing"
)

func TestAuthValidator(t *testing.T) {
	tests := []struct {
		name  string
		value *pb.RegisterRequest
		valid bool
	}{
		{
			name: "Valid request",
			value: &pb.RegisterRequest{
				Email:       "samplemail@gmail.com",
				Password:    "st0ng1Nd3pn3t",
				Fullname:    "Sample fullname",
				Birthdate:   "2001-10-01",
				Gender:      "F",
				Address:     "Noway street, California",
				MbtiType:    "1",
				JobPosition: "Software engineer",
			},
			valid: true,
		},
		{
			name: "Bad request",
			value: &pb.RegisterRequest{
				Email:       "samplemailgmail.com",
				Password:    "st0ng1d3pn3t",
				Fullname:    "Sample fullname",
				Birthdate:   "2006-10-01",
				Gender:      "F",
				Address:     "Noway street, California",
				MbtiType:    "1",
				JobPosition: "Software engineer",
			},
			valid: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := auth.ValidateRegister(test.value)
			isValid := true
			if err != nil {
				isValid = false
			}
			if isValid != test.valid {
				t.Errorf("Error validate request: %v\n", err)
			}
		})
	}

}
