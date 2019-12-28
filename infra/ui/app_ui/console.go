package app_ui

import (
	"bufio"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/quality/infra/qt_missingmsg"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
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

const (
	consoleNumRowsThreshold = 500
)

func NewConsole(mc app_msg_container.Container, qm qt_missingmsg.Message, testMode bool) UI {
	return &console{
		mc:       mc,
		out:      os.Stdout,
		in:       os.Stdin,
		testMode: testMode,
		qm:       qm,
	}
}

func CloneConsole(ui UI, mc app_msg_container.Container) UI {
	switch u := ui.(type) {
	case *console:
		return &console{
			mc:       mc,
			out:      u.out,
			in:       u.in,
			testMode: u.testMode,
			qm:       u.qm,
		}

	case *Quiet:
		return &Quiet{
			mc:  mc,
			log: u.log,
		}

	default:
		app_root.Log().Error("Unsupported UI type")
		panic("unsupported UI type")
	}
}

type console struct {
	mc               app_msg_container.Container
	out              io.Writer
	in               io.Reader
	testMode         bool
	qm               qt_missingmsg.Message
	mutex            sync.Mutex
	openArtifactOnce sync.Once
}

func (z *console) Header(m app_msg.Message) {
	z.verifyKey(m.Key())
	t := z.mc.Compile(m)
	z.boldPrint(t)
	z.Break()
}

func (z *console) Text(m app_msg.Message) string {
	z.verifyKey(m.Key())
	return z.mc.Compile(m)
}

func (z *console) TextOrEmpty(m app_msg.Message) string {
	if z.mc.Exists(m.Key()) {
		return z.mc.Compile(m)
	} else {
		return ""
	}
}

func (z *console) Info(m app_msg.Message) {
	z.verifyKey(m.Key())
	t := z.mc.Compile(m)
	z.colorPrint(t, ColorWhite)
	app_root.Log().Debug(t)
}

func (z *console) Error(m app_msg.Message) {
	z.verifyKey(m.Key())
	t := z.mc.Compile(m)
	z.colorPrint(t, ColorRed)
	app_root.Log().Debug(t)
}

func (z *console) IsConsole() bool {
	return true
}

func (z *console) IsWeb() bool {
	return false
}

func (z *console) OpenArtifact(path string) {
	l := app_root.Log()

	if z.testMode {
		return
	}

	z.openArtifactOnce.Do(func() {
		app_root.AddSuccessShutdownHook(func() {
			files, err := ioutil.ReadDir(path)
			if err != nil {
				l.Debug("Unable to read path", zap.Error(err), zap.String("path", path))
				return
			}
			for _, f := range files {
				e := filepath.Ext(f.Name())
				switch strings.ToLower(e) {
				case ".xlsx", ".csv", ".json":
					z.InfoK("run.console.point_artifact", app_msg.P{
						"Path": filepath.Join(path, f.Name()),
					})

				default:
					l.Debug("unsupported extension", zap.String("name", f.Name()))
				}
			}

			z.InfoK("run.console.open_artifact", app_msg.P{"Path": path})
			l.Debug("Register success shutdown hook", zap.String("path", path))
			err = open.Start(path)
			if err != nil {
				z.ErrorK("run.console.open_artifact.error", app_msg.P{"ErrorK": err})
			}
		})
	})
}

func (z *console) verifyKey(key string) {
	if !z.mc.Exists(key) {
		z.qm.NotFound(key)
	}
}

func (z *console) TextK(key string, p ...app_msg.P) string {
	return z.Text(app_msg.M(key, p...))
}

func (z *console) TextOrEmptyK(key string, p ...app_msg.P) string {
	return z.TextOrEmpty(app_msg.M(key, p...))
}

func (z *console) Break() {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	fmt.Fprintln(z.out)
}

func (z *console) colorPrint(t string, color int) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if runtime.GOOS == "windows" {
		fmt.Fprintf(z.out, "%s\n", t)
	} else {
		fmt.Fprintf(z.out, "\x1b[%dm%s\x1b[0m\n", color, t)
	}
}

func (z *console) boldPrint(t string) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if runtime.GOOS == "windows" {
		fmt.Fprintf(z.out, "%s\n", t)
	} else {
		fmt.Fprintf(z.out, "\x1b[1m%s\x1b[0m\n", t)
	}
}

func (z *console) HeaderK(key string, p ...app_msg.P) {
	z.Header(app_msg.M(key, p...))
}

func (z *console) InfoTable(name string) Table {
	return newMarkdownTable(z.mc, z.out, false)
	//
	//tw := new(tabwriter.Writer)
	//tw.Init(z.out, 0, 2, 2, ' ', 0)
	//return &consoleTable{
	//	mc:   z.mc,
	//	tab:  tw,
	//	qm:   z.qm,
	//	name: name,
	//	ui:   z,
	//}
}

func (z *console) InfoK(key string, p ...app_msg.P) {
	z.Info(app_msg.M(key, p...))
}

func (z *console) ErrorK(key string, p ...app_msg.P) {
	z.Error(app_msg.M(key, p...))
}

func (z *console) Success(key string, p ...app_msg.P) {
	z.verifyKey(key)
	m := z.mc.Compile(app_msg.M(key, p...))
	z.colorPrint(m, ColorGreen)
	app_root.Log().Debug(m)
}

func (z *console) Failure(key string, p ...app_msg.P) {
	z.verifyKey(key)
	m := z.mc.Compile(app_msg.M(key, p...))
	z.colorPrint(m, ColorRed)
	app_root.Log().Debug(m)
}

func (z *console) AskCont(key string, p ...app_msg.P) (cont bool, cancel bool) {
	z.verifyKey(key)
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

func (z *console) AskText(key string, p ...app_msg.P) (text string, cancel bool) {
	z.verifyKey(key)
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

func (z *console) AskSecure(key string, p ...app_msg.P) (text string, cancel bool) {
	z.verifyKey(key)
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
	mc      app_msg_container.Container
	tab     *tabwriter.Writer
	qm      qt_missingmsg.Message
	mutex   sync.Mutex
	numRows int
	name    string
	ui      UI
}

func (z *consoleTable) HeaderRaw(h ...string) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	var p, q = "", ""

	//if runtime.GOOS != "windows" {
	//	p = "\x1b[1m"
	//	q = "\x1b[0m"
	//}
	fmt.Fprintln(z.tab, p+strings.Join(h, "\t")+q)
}

func (z *consoleTable) RowRaw(m ...string) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	z.numRows++
	if z.numRows <= consoleNumRowsThreshold {
		fmt.Fprintln(z.tab, strings.Join(m, "\t"))
	}
	if z.numRows%consoleNumRowsThreshold == 0 {
		z.ui.InfoK("run.console.progress", app_msg.P{
			"Label":    z.name,
			"Progress": z.numRows,
		})
	}
}

func (z *consoleTable) Header(h ...app_msg.Message) {
	headers := make([]string, 0)
	for _, hdr := range h {
		headers = append(headers, z.mc.Compile(hdr))
	}
	z.HeaderRaw(headers...)
}

func (z *consoleTable) validateMessage(m app_msg.Message) {
	if !z.mc.Exists(m.Key()) {
		z.qm.NotFound(m.Key())
	}
}

func (z *consoleTable) Row(m ...app_msg.Message) {
	msgs := make([]string, 0)
	for _, msg := range m {
		z.validateMessage(msg)
		msgs = append(msgs, z.mc.Compile(msg))
	}
	z.RowRaw(msgs...)
}

func (z *consoleTable) Flush() {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	z.tab.Flush()
	if z.numRows >= consoleNumRowsThreshold {
		z.ui.InfoK("run.console.large_report", app_msg.P{
			"Num": z.numRows,
		})
	}
}
