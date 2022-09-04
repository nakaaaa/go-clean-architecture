package env

import "strconv"

func GetBool(key string) (bool, error) {
	s, err := GetString(key)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(s)
}
