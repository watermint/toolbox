package app_ui

import (
	"bufio"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"go.uber.org/zap"
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

func NewConsole(mc app_msg_container.Container, testMode bool) UI {
	return &console{
		mc:       mc,
		out:      os.Stdout,
		in:       os.Stdin,
		testMode: testMode,
	}
}

type console struct {
	mc       app_msg_container.Container
	out      io.Writer
	in       io.Reader
	testMode bool
}

func (z *console) IsConsole() bool {
	return true
}

func (z *console) IsWeb() bool {
	return false
}

func (z *console) OpenArtifact(path string) {
	z.Info("run.console.open_artifact", app_msg.P("Path", path))
	if z.testMode {
		return
	}

	err := open.Start(path)
	if err != nil {
		z.Error("run.console.open_artifact.error", app_msg.P("Error", err))
	}
}

func (z *console) Text(key string, p ...app_msg.Param) string {
	return z.mc.Compile(app_msg.M(key, p...))
}

func (z *console) Break() {
	fmt.Fprintln(z.out)
}

func (z *console) colorPrint(t string, color int) {
	if runtime.GOOS == "windows" {
		fmt.Fprintf(z.out, "%s\n", t)
	} else {
		fmt.Fprintf(z.out, "\x1b[%dm%s\x1b[0m\n", color, t)
	}
}

func (z *console) boldPrint(t string) {
	if runtime.GOOS == "windows" {
		fmt.Fprintf(z.out, "%s\n", t)
	} else {
		fmt.Fprintf(z.out, "\x1b[1m%s\x1b[0m\n", t)
	}
}

func (z *console) Header(key string, p ...app_msg.Param) {
	m := z.mc.Compile(app_msg.M(key, p...))
	z.boldPrint(m)
}

func (z *console) InfoTable(name string) Table {
	tw := new(tabwriter.Writer)
	tw.Init(z.out, 0, 2, 2, ' ', 0)
	return &consoleTable{
		Container: z.mc,
		Tab:       tw,
	}
}

func (z *console) Info(key string, p ...app_msg.Param) {
	m := z.mc.Compile(app_msg.M(key, p...))
	z.colorPrint(m, ColorWhite)
	app_root.Log().Debug(m)
}

func (z *console) Error(key string, p ...app_msg.Param) {
	m := z.mc.Compile(app_msg.M(key, p...))
	z.colorPrint(m, ColorRed)
	app_root.Log().Debug(m)
}

func (z *console) Success(key string, p ...app_msg.Param) {
	m := z.mc.Compile(app_msg.M(key, p...))
	z.colorPrint(m, ColorGreen)
	app_root.Log().Debug(m)
}

func (z *console) Failure(key string, p ...app_msg.Param) {
	m := z.mc.Compile(app_msg.M(key, p...))
	z.colorPrint(m, ColorRed)
	app_root.Log().Debug(m)
}

func (z *console) AskCont(key string, p ...app_msg.Param) (cont bool, cancel bool) {
	msg := z.mc.Compile(app_msg.M(key, p...))
	app_root.Log().Debug(msg)

	z.colorPrint(msg, ColorCyan)
	br := bufio.NewReader(z.in)
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
	msg := z.mc.Compile(app_msg.M(key, p...))
	z.colorPrint(msg, ColorCyan)
	app_root.Log().Debug(msg)

	br := bufio.NewReader(z.in)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			app_root.Log().Debug("Cancelled")
			return "", true
		}
		text := strings.TrimSpace(string(line))
		if text != "" {
			app_root.Log().Debug("Text entered", zap.String("text", text))
			return text, false
		}

		// ask again
		z.colorPrint(msg, ColorCyan)
	}
}

func (z *console) AskSecure(key string, p ...app_msg.Param) (text string, cancel bool) {
	msg := z.mc.Compile(app_msg.M(key, p...))
	z.colorPrint(msg, ColorCyan)
	app_root.Log().Debug(msg)

	br := bufio.NewReader(z.in)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			app_root.Log().Debug("Cancelled")
			return "", true
		}
		text := strings.TrimSpace(string(line))
		if text != "" {
			app_root.Log().Debug("Secret entered")
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
	var p, q = "", ""

	//if runtime.GOOS != "windows" {
	//	p = "\x1b[1m"
	//	q = "\x1b[0m"
	//}
	fmt.Fprintln(z.Tab, p+strings.Join(h, "\t")+q)
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
