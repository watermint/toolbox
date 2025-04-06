package es_color

import (
	"fmt"
	"io"

	"github.com/watermint/toolbox/essentials/terminal/es_terminfo"
)

const (
	// https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
	ColorBlack Color = iota + 30
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

const (
	beginBold  = "\x1b[1m"
	beginColor = "\x1b[%dm"
	endFormat  = "\x1b[0m"
)

type Color int

func Colorf(w io.Writer, c Color, bold bool, format string, a ...interface{}) {
	t := fmt.Sprintf(format, a...)
	if es_terminfo.IsOutColorTerminal() {
		if bold {
			_, _ = fmt.Fprintf(w, "\x1b[1m\x1b[%dm%s\x1b[0m", c, t)
		} else {
			_, _ = fmt.Fprintf(w, "\x1b[%dm%s\x1b[0m", c, t)
		}
	} else {
		_, _ = fmt.Fprintf(w, "%s", t)
	}
}

func Colorfln(w io.Writer, c Color, bold bool, format string, a ...interface{}) {
	Colorf(w, c, bold, format, a...)
	_, _ = fmt.Fprintln(w)
}

func Boldf(w io.Writer, format string, a ...interface{}) {
	t := fmt.Sprintf(format, a...)
	if es_terminfo.IsOutColorTerminal() {
		_, _ = fmt.Fprintf(w, "\x1b[1m%s\x1b[0m", t)
	} else {
		_, _ = fmt.Fprintf(w, "%s", t)
	}
}

func Boldfln(w io.Writer, format string, a ...interface{}) {
	Boldf(w, format, a...)
	_, _ = fmt.Fprintln(w)
}
