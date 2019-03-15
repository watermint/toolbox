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

func ColorPrint(out io.Writer, t string, color int) {
	fmt.Fprintf(out, "\x1b[%dm%s\x1b[0m", color, t)
}

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

func (z *ColorCUI) Tell(msg UIMessage) {
	if z.debugMode {
		ColorPrint(z.Out, "TELL\t", ColorCyan)
	}
	fmt.Fprintln(z.Out, msg.T())
}

func (z *ColorCUI) TellError(msg UIMessage) {
	ColorPrint(z.Out, "ERROR\t", ColorBrightRed)
	fmt.Fprintln(z.Out, msg.T())
}

func (z *ColorCUI) TellSuccess(msg UIMessage) {
	ColorPrint(z.Out, "SUCCESS\t", ColorGreen)
	fmt.Fprintln(z.Out, msg.T())
}

func (z *ColorCUI) TellFailure(msg UIMessage) {
	ColorPrint(z.Out, "FAILURE\t", ColorRed)
	fmt.Fprintln(z.Out, msg.T())
}

func (z *ColorCUI) AskRetry(msg UIMessage) bool {
	return z.cui.AskRetry(msg)
}

func (z *ColorCUI) AskText(msg UIMessage) string {
	return z.cui.AskText(msg)
}

func (z *ColorCUI) AskConfirm(msg UIMessage) bool {
	return z.cui.AskConfirm(msg)
}
