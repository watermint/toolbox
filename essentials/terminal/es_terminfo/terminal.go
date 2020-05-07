package es_terminfo

import (
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func IsInTerminal() bool {
	return terminal.IsTerminal(int(os.Stdin.Fd()))
}

func IsOutTerminal() bool {
	return terminal.IsTerminal(int(os.Stdout.Fd()))
}

// Determine weather the terminal supports colors or not.
func IsOutColorTerminal() bool {
	if !IsOutTerminal() {
		return false
	}

	// May execute command `tput colors` (not on Win) to determine supported number of colors.
	// But this time just returns true.
	return true
}
