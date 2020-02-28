package env

import (
	"os"
	"strconv"
)

//Getenv :
func Getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

//GetenvI :
func GetenvI(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		ival, err := strconv.Atoi(value)
		if err == nil {
			return ival
		}
	}
	return fallback
}
