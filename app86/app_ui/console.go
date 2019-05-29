package app_ui

import (
	"bufio"
	"fmt"
	"github.com/watermint/toolbox/app86/app_msg"
	"github.com/watermint/toolbox/app86/app_msg_container"
	"github.com/watermint/toolbox/app86/app_root"
	"io"
	"os"
	"runtime"
	"strings"
	"text/tabwriter"
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

func NewConsole(mc app_msg_container.Container) UI {
	return &console{
		Container: mc,
		Out:       os.Stdout,
		In:        os.Stdin,
	}
}

type console struct {
	Container app_msg_container.Container
	Out       io.Writer
	In        io.Reader
}

func (z *console) Break() {
	fmt.Fprintln(z.Out)
}

func (z *console) colorPrint(t string, color int) {
	if runtime.GOOS == "windows" {
		fmt.Fprintf(z.Out, "%s\n", t)
	} else {
		fmt.Fprintf(z.Out, "\x1b[%dm%s\x1b[0m\n", color, t)
	}
}

func (z *console) boldPrint(t string) {
	if runtime.GOOS == "windows" {
		fmt.Fprintf(z.Out, "%s\n", t)
	} else {
		fmt.Fprintf(z.Out, "\x1b[1m%s\x1b[0m\n", t)
	}
}

func (z *console) Header(key string, p ...app_msg.Param) {
	m := z.Container.Compile(app_msg.M(key, p...))
	z.boldPrint(m)
}

func (z *console) InfoTable(border bool) Table {
	tw := new(tabwriter.Writer)
	tw.Init(z.Out, 0, 2, 2, ' ', 0)
	return &consoleTable{
		Container: z.Container,
		Tab:       tw,
	}
}

func (z *console) Info(key string, p ...app_msg.Param) {
	m := z.Container.Compile(app_msg.M(key, p...))
	z.colorPrint(m, ColorWhite)
	app_root.Log().Debug(m)
}

func (z *console) Error(key string, p ...app_msg.Param) {
	m := z.Container.Compile(app_msg.M(key, p...))
	z.colorPrint(m, ColorRed)
	app_root.Log().Debug(m)
}

func (z *console) AskCont(key string, p ...app_msg.Param) (cont bool, cancel bool) {
	msg := z.Container.Compile(app_msg.M(key, p...))
	app_root.Log().Debug(msg)

	z.colorPrint(msg, ColorCyan)
	br := bufio.NewReader(z.In)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			app_root.Log().Debug("Cancelled")
			return false, true
		}
		ans := strings.ToLower(strings.TrimSpace(string(line)))
		switch ans {
		case "y":
			app_root.Log().Debug("Continue")
			return true, false
		case "yes":
			app_root.Log().Debug("Continue")
			return true, false
		case "n":
			app_root.Log().Debug("Do not continue")
			return false, false
		case "no":
			app_root.Log().Debug("Do not continue")
			return false, false
		}

		// ask again
		z.colorPrint(msg, ColorCyan)
	}
}

func (z *console) AskText(key string, p ...app_msg.Param) (text string, cancel bool) {
	msg := z.Container.Compile(app_msg.M(key, p...))
	z.colorPrint(msg, ColorCyan)
	app_root.Log().Debug(msg)

	br := bufio.NewReader(z.In)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			app_root.Log().Debug("Cancelled")
			return "", true
		}
		text := strings.TrimSpace(string(line))
		if text != "" {
			app_root.Log().Debug("Text entered")
			return text, false
		}

		// ask again
		z.colorPrint(msg, ColorCyan)
	}
}

type consoleTable struct {
	Container app_msg_container.Container
	Tab       *tabwriter.Writer
}

func (z *consoleTable) HeaderRaw(h ...string) {
	r := make([]string, 0)
	for _, v := range h {
		if runtime.GOOS == "windows" {
			r = append(r, fmt.Sprintf("%s", v))
		} else {
			r = append(r, fmt.Sprintf("\x1b[1m%s\x1b[0m", v))
		}
	}
	fmt.Fprintln(z.Tab, strings.Join(r, "\t"))
}

func (z *consoleTable) RowRaw(m ...string) {
	fmt.Fprintln(z.Tab, strings.Join(m, "\t"))
}

func (z *consoleTable) Header(h ...app_msg.Message) {
	headers := make([]string, 0)
	for _, hdr := range h {
		headers = append(headers, z.Container.Compile(hdr))
	}
	z.HeaderRaw(headers...)
}

func (z *consoleTable) Row(m ...app_msg.Message) {
	msgs := make([]string, 0)
	for _, msg := range m {
		msgs = append(msgs, z.Container.Compile(msg))
	}
	z.RowRaw(msgs...)
}

func (z *consoleTable) Flush() {
	z.Tab.Flush()
}
