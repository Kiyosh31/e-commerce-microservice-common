package grpcvalidators

import (
	"fmt"
	"net/mail"
	"regexp"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	isValidName = regexp.MustCompile(`^[a-z ,.'-]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("Must contain from: %d-%d characters", minLength, maxLength)
	}

	return nil
}

func ValidateName(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}

	if !isValidName(value) {
		return fmt.Errorf("Must contain only letters")
	}

	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 3, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("Is not a valid email")
	}

	return nil
}

func UnauthenticatedError(err error) error {
	return status.Errorf(codes.Unauthenticated, "Unauthorized: %v", err)
}
