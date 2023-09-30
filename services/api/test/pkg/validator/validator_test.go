package validator_test

import (
	"api/pkg/validator"
	"log"
	"testing"
)

type BasicTestCase struct {
	name  string
	value string
	valid bool
}

// test validation
func TestEmailValidation(t *testing.T) {
	tests := []BasicTestCase{
		{
			name:  "No email domain",
			value: "khfeooldiie",
			valid: false,
		},
		{
			name:  "No @ symbol",
			value: "jajanggmail.com",
			valid: false,
		},
		{
			name:  "Whitespace",
			value: "jaga @gmail.com",
			valid: false,
		},
		{
			name:  "Valid email",
			value: "heloa@gmail.com",
			valid: true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			err := validator.ValidateEmail(c.value)
			var validated bool
			if err != nil {
				validated = false
			} else {
				validated = true
			}
			if validated != c.valid {
				log.Println(c)
				t.Fatalf("invalid email: %v\n", err)
			}
		})
	}

}

func TestPasswordValidator(t *testing.T) {
	tests := []BasicTestCase{
		{
			name:  "No Numeric",
			value: "nonumeric",
			valid: false,
		},
		{
			name:  "No Alphabet",
			value: "12473662",
			valid: false,
		},
		{
			name:  "No Lowercase",
			value: "1247H3662",
			valid: false,
		},
		{
			name:  "No Uppercase",
			value: "1247h3662",
			valid: false,
		},
		{
			name:  "Correct Password",
			value: "Str0ngp44ass",
			valid: true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			err := validator.ValidatePassword(c.value)
			var validated bool
			if err != nil {
				validated = false
			} else {
				validated = true
			}
			if validated != c.valid {
				t.Fatalf("invalid password: %v\n", err)
			}
		})
	}

}

func TestFullNameValidator(t *testing.T) {
	tests := []BasicTestCase{
		{
			name:  "Invalid length",
			value: "Hey",
			valid: false,
		},
		{
			name:  "Valid fullname",
			value: "Kholid Bughowi",
			valid: true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			err := validator.ValidateFullName(c.value)
			var validated bool
			if err != nil {
				validated = false
			} else {
				validated = true
			}
			if validated != c.valid {
				t.Fatalf("invalid fullname: %v\n", err)
			}
		})
	}
}

func TestBirtdateValidator(t *testing.T) {
	tests := []BasicTestCase{
		{
			name:  "Invalid birthdate Now",
			value: "2023-09-15",
			valid: false,
		},
		{
			name:  "Invalid birthdate After",
			value: "2024-10-15",
			valid: false,
		},
		{
			name:  "Valid birthdate",
			value: "2000-02-01",
			valid: true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			err := validator.ValidateBirtdate(c.value)
			var validated bool
			if err != nil {
				validated = false
			} else {
				validated = true
			}
			if validated != c.valid {
				t.Log(c.name, c.value)
				t.Fatalf("invalid birthdate: %v\n", err)
			}
		})
	}
}
