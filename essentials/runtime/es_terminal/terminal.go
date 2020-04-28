package es_terminal

import (
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func IsTerminal() bool {
	return terminal.IsTerminal(int(os.Stdout.Fd()))
}
