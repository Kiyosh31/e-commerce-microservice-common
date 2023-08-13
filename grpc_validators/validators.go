package grpcvalidators

import (
	"fmt"
	"net/mail"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	isValidName = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	isValidDate = regexp.MustCompile(`^\d{4}\-(0?[1-9]|1[012])\-(0?[1-9]|[12][0-9]|3[01])$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("Must contain from: %d-%d characters", minLength, maxLength)
	}

	return nil
}

func ValidateMongoId(value string) error {
	_, err := primitive.ObjectIDFromHex(value)
	if err != nil {
		return err
	}

	return nil
}

func ValidateName(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}

	if !isValidName(value) {
		return fmt.Errorf("Must contain letters or numbers")
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

func ValidateBirth(value string) error {
	if err := ValidateString(value, 5, 10); err != nil {
		return err
	}

	if !isValidDate(value) {
		return fmt.Errorf("Must be a valid date format: yyyy-mm-dd")
	}

	return nil
}

func ValidateRole(value string) error {
	if err := ValidateString(value, 5, 10); err != nil {
		return err
	}

	if value == "customer" || value == "seller" {
		return fmt.Errorf("Role muts have format: customer|seller")
	}

	return nil
}

func UnauthenticatedError(err error) error {
	return status.Errorf(codes.Unauthenticated, "Unauthorized: %v", err)
}
