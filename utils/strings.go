package utils

import "fmt"

func ParseInterfaceToString(word interface{}) (string, error) {
	s, ok := word.(string)
	if !ok {
		return "nil", fmt.Errorf("Error while parsing interface to string")
	}

	return s, nil
}
