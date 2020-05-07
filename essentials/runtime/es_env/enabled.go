package es_env

import (
	"os"
	"strconv"
)

// Returns true, if a feature enabled by env variable.
func IsEnabled(envName string) bool {
	if b, err := strconv.ParseBool(os.Getenv(envName)); err != nil {
		return false
	} else {
		return b
	}
}
