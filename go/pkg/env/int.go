package env

import (
	"strconv"
)

func GetInt(key string) (int, error) {
	s, err := GetString(key)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(s)
}
