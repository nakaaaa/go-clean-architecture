package env

import (
	"fmt"
	"os"
)

func GetString(key string) (string, error) {
	s := os.Getenv(key)
	if s == "" {
		// return "", fmt.Errorf("getenv returns empty string: key=%s", key)
		return "", fmt.Errorf(ErrEmpty+": key=%s", key)
	}

	return s, nil
}
