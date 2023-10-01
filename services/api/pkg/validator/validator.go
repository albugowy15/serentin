package validator

import (
	"fmt"
	"log"
	"regexp"
	"time"
	"unicode"
)

var (
	emailPattern = `^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+[a-zA-Z0-9-.]+$`
)

func ValidateEmail(email string) error {
	if len(email) < 5 || len(email) > 200 {
		return fmt.Errorf("email must be between 5 and 200 characters")
	}
	re := regexp.MustCompile(emailPattern)
	if !re.MatchString(email) {
		return fmt.Errorf("invalid email")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 || len(password) > 30 {
		return fmt.Errorf("password must be between 8 and 30 characters")
	}
	isValid := true
	var (
		upp, low, num bool
		tot           uint8
	)
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		default:
			isValid = false
		}
	}
	if !upp || !low || !num || tot < 8 {
		isValid = false
	}

	if isValid {
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

func ValidateBirtdate(birthdate string) error {
	dateFormat := "2006-01-02"
	parsedBirthdate, err := time.Parse(dateFormat, birthdate)
	if err != nil {
		log.Fatal(err)
	}
	minBirthdate := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	maxBirthdate := time.Now().AddDate(-18, 0, 0)

	isValid := !parsedBirthdate.Before(minBirthdate) && !parsedBirthdate.After(maxBirthdate)

	if !isValid {
		return fmt.Errorf("invalid birthdate")
	}
	return nil
}
