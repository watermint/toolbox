package app_ui

import (
	"fmt"
	"io"
	"os"
)

const (
	// https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
	ColorBlack = iota + 30
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
	ColorBrightBlack
	ColorBrightRed
	ColorBrightGreen
	ColorBrightYellow
	ColorBrightBlue
	ColorBrightMagenta
	ColorBrightCyan
	ColorBrightWhite
)

func NewColorCUI() UI {
	return &ColorCUI{
		Out: os.Stdout,
		In:  os.Stdin,
		cui: &CUI{
			Out: os.Stdout,
			In:  os.Stdin,
		},
	}
}

type ColorCUI struct {
	Out       io.Writer
	In        io.Reader
	debugMode bool
	cui       UI
}

func (z *ColorCUI) DebugMode(debug bool) {
	z.debugMode = debug
}

func (z *ColorCUI) print(t string, color int) {
	fmt.Fprintf(z.Out, "\x1b[%dm%s\x1b[0m", color, t)
}

func (z *ColorCUI) Tell(msg UIMessage) {
	if z.debugMode {
		z.print("TELL\t", ColorCyan)
	}
	fmt.Fprintln(z.Out, msg.T())
}

func (z *ColorCUI) TellError(msg UIMessage) {
	z.print("ERROR\t", ColorBrightRed)
	fmt.Fprintln(z.Out, msg.T())
}

func (z *ColorCUI) TellSuccess(msg UIMessage) {
	z.print("SUCCESS\t", ColorGreen)
	fmt.Fprintln(z.Out, msg.T())
}

func (z *ColorCUI) TellFailure(msg UIMessage) {
	z.print("FAILURE\t", ColorBrightRed)
	fmt.Fprintln(z.Out, msg.T())
}

func (z *ColorCUI) AskRetry(msg UIMessage) bool {
	return z.cui.AskRetry(msg)
}

func (z *ColorCUI) AskText(msg UIMessage) string {
	return z.cui.AskText(msg)
}
