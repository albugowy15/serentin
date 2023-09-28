package validator

import (
	"fmt"
	"regexp"
	"time"
)

var (
	emailPattern    = `^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+[a-zA-Z0-9-.]+$`
	passwordPattern = `^[A-Za-z]*[A-Z]+[A-Za-z0-9]*$|^[A-Za-z]*[a-z]+[A-Za-z0-9]*$|^[A-Za-z]*[0-9]+[A-Za-z0-9]*$`
)

func ValidateEmail(email string) error {
	if len(email) < 5 || len(email) > 200 {
		return fmt.Errorf("email must be between 5 and 200 characters")
	}
	re := regexp.MustCompile(emailPattern)
	if re.MatchString(email) {
		return nil
	} else {
		return fmt.Errorf("invalid email")
	}
}

func ValidatePassword(password string) error {
	if len(password) < 8 || len(password) > 30 {
		return fmt.Errorf("password must be between 8 and 30 characters")
	}
	re := regexp.MustCompile(passwordPattern)
	if re.MatchString(password) {
		return nil
	} else {
		return fmt.Errorf("password must contain at least one uppercase letter, one lowercase letter and one number")
	}
}

func ValidateFullName(fullname string) error {
	if len(fullname) < 5 || len(fullname) > 200 {
		return fmt.Errorf("fullname must be between 5 and 200 characters")
	}
	return nil
}

func ValidateBirtdate(birthdate *time.Time) error {
	if birthdate == nil {
		return fmt.Errorf("birthdate is required")
	}
	if birthdate.After(time.Now()) {
		return fmt.Errorf("birtdate must be before today")
	}

	return nil
}
