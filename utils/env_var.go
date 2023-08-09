package utils

import (
	"fmt"
	"os"
)

func GetEnvVar(envVar string) (string, error) {
	value := os.Getenv(envVar)
	if value == "" {
		return "",
			fmt.Errorf("You must provide" + envVar + "in environment variables")
	}

	return value, nil
}
