//go:build darwin || linux
// +build darwin linux

package elocale

import (
	"errors"
	"os"
)

func currentLocaleString() (string, error) {
	if lcAll := os.Getenv("LC_ALL"); lcAll != "" {
		return lcAll, nil
	}
	if lang := os.Getenv("LANG"); lang != "" {
		return lang, nil
	}
	return "", errors.New("locale not found")
}
