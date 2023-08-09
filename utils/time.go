package utils

import "time"

func stringToTimeDuration(definedTime string) (time.Duration, error) {
	duration, err := time.ParseDuration(definedTime)
	if err != nil {
		return time.Duration(0), err
	}

	return duration, nil
}
