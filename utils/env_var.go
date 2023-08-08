package main

import (
	"errors"
	"os"
)

func GetEnvVar(envVar string) (string, error) {
	value := os.Getenv(envVar)
	if value == "" {
		return "", errors.New("You must provide " + envVar + " in environment variables")
	}

	return value, nil
}
