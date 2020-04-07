package app_ui

import (
	"bufio"
	"fmt"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"github.com/watermint/toolbox/infra/util/ut_open"
	"github.com/watermint/toolbox/infra/util/ut_string"
	"github.com/watermint/toolbox/infra/util/ut_terminal"
	"github.com/watermint/toolbox/quality/infra/qt_missingmsg"
	"github.com/watermint/toolbox/quality/infra/qt_missingmsg_impl"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
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

type MsgConsole struct {
	LargeReport       app_msg.Message
	OpenArtifactError app_msg.Message
	OpenArtifact      app_msg.Message
	PointArtifact     app_msg.Message
	Progress          app_msg.Message
}

var (
	MConsole = app_msg.Apply(&MsgConsole{}).(*MsgConsole)
)

const (
	consoleNumRowsThreshold = 500
)

func NewConsole(mc app_msg_container.Container, qm qt_missingmsg.Message, testMode bool) UI {
	return &console{
		id:       newId(),
		mc:       mc,
		out:      ut_io.NewDefaultOut(testMode),
		in:       os.Stdin,
		testMode: testMode,
		qm:       qm,
		useColor: true,
	}
}

func NewNullConsole(mc app_msg_container.Container, qm qt_missingmsg.Message) UI {
	return &console{
		id:       newId(),
		mc:       mc,
		out:      ioutil.Discard,
		in:       os.Stdin,
		testMode: true,
		qm:       qm,
		useColor: true,
	}
}

func NewBufferConsole(mc app_msg_container.Container, buf io.Writer) UI {
	return &console{
		id:       newId(),
		mc:       mc,
		out:      buf,
		in:       os.Stdin,
		qm:       qt_missingmsg_impl.NewMessageMemory(),
		useColor: false,
	}
}

func CloneConsole(ui UI, mc app_msg_container.Container) UI {
	switch u := ui.(type) {
	case *console:
		return &console{
			id:       newId(),
			mc:       mc,
			out:      u.out,
			in:       u.in,
			testMode: u.testMode,
			qm:       u.qm,
		}

	case *Quiet:
		return NewQuiet(mc)

	default:
		app_root.Log().Error("Unsupported UI type")
		panic("unsupported UI type")
	}
}

type console struct {
	id               string
	mc               app_msg_container.Container
	out              io.Writer
	in               io.Reader
	l                *zap.Logger
	testMode         bool
	useColor         bool
	qm               qt_missingmsg.Message
	mutex            sync.Mutex
	openArtifactOnce sync.Once
}

func (z *console) Id() string {
	return z.id
}

func (z *console) Progress(m app_msg.Message) {
	z.verifyKey(m.Key())
	t := z.mc.Compile(m)
	z.colorPrint(t, ColorCyan)
	z.currentLogger().Debug(t)
}

func (z *console) SubHeader(m app_msg.Message) {
	z.verifyKey(m.Key())
	t := z.mc.Compile(m)
	tl := ut_string.Width(t)
	z.Break()
	z.boldPrint(t)
	z.boldPrint(strings.Repeat("-", tl))
	z.Break()
}

func (z *console) Code(code string) {
	z.Break()
	z.colorPrint(code, ColorBlue)
	z.Break()
}

func (z *console) Exists(m app_msg.Message) bool {
	return z.mc.Exists(m.Key())
}

func (z *console) SetLogger(l *zap.Logger) {
	z.l = l
}

func (z *console) currentLogger() *zap.Logger {
	if z.l == nil {
		return app_root.Log()
	} else {
		return z.l
	}
}

func (z *console) AskCont(m app_msg.Message) (cont bool, cancel bool) {
	z.verifyKey(m.Key())
	msg := z.mc.Compile(m)
	z.currentLogger().Debug(msg)

	z.colorPrint(msg, ColorCyan)
	br := bufio.NewReader(z.in)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			z.currentLogger().Debug("Cancelled")
			return false, true
		}
		ans := strings.ToLower(strings.TrimSpace(string(line)))
		switch ans {
		case "y":
			z.currentLogger().Debug("Continue")
			return true, false
		case "yes":
			z.currentLogger().Debug("Continue")
			return true, false
		case "n":
			z.currentLogger().Debug("Do not continue")
			return false, false
		case "no":
			z.currentLogger().Debug("Do not continue")
			return false, false
		}

		// ask again
		z.colorPrint(msg, ColorCyan)
	}
}

func (z *console) AskText(m app_msg.Message) (text string, cancel bool) {
	z.verifyKey(m.Key())
	msg := z.mc.Compile(m)
	z.colorPrint(msg, ColorCyan)
	z.currentLogger().Debug(msg)

	br := bufio.NewReader(z.in)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			z.currentLogger().Debug("Cancelled")
			return "", true
		}
		text := strings.TrimSpace(string(line))
		if text != "" {
			z.currentLogger().Debug("Text entered", zap.String("text", text))
			return text, false
		}

		// ask again
		z.colorPrint(msg, ColorCyan)
	}
}

func (z *console) AskSecure(m app_msg.Message) (secure string, cancel bool) {
	z.verifyKey(m.Key())
	msg := z.mc.Compile(m)
	z.colorPrint(msg, ColorCyan)
	z.currentLogger().Debug(msg)

	br := bufio.NewReader(z.in)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			z.currentLogger().Debug("Cancelled")
			return "", true
		}
		text := strings.TrimSpace(string(line))
		if text != "" {
			z.currentLogger().Debug("Secret entered")
			return text, false
		}

		// ask again
		z.colorPrint(msg, ColorCyan)
	}
}

func (z *console) Header(m app_msg.Message) {
	z.verifyKey(m.Key())
	t := z.mc.Compile(m)
	tl := ut_string.Width(t)
	z.Break()
	z.boldPrint(t)
	z.boldPrint(strings.Repeat("=", tl))
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
	z.currentLogger().Debug(t)
}

func (z *console) Error(m app_msg.Message) {
	z.verifyKey(m.Key())
	t := z.mc.Compile(m)
	z.colorPrint(t, ColorRed)
	z.currentLogger().Debug(t)
}

func (z *console) IsConsole() bool {
	return true
}

func (z *console) IsWeb() bool {
	return false
}

func (z *console) OpenArtifact(path string, autoOpen bool) {
	l := z.currentLogger()

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
					z.Info(MConsole.PointArtifact.With(
						"Path", filepath.Join(path, f.Name()),
					))

				default:
					l.Debug("unsupported extension", zap.String("name", f.Name()))
				}
			}

			l.Debug("Register success shutdown hook", zap.String("path", path))
			if z.testMode || !autoOpen {
				return
			}
			z.Info(MConsole.OpenArtifact.With("Path", path))

			if err := ut_open.New().Open(path, true); err != nil {
				z.Error(MConsole.OpenArtifactError.With("Error", err))
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

	if app.IsWindows() || !z.useColor || !ut_terminal.IsTerminal() {
		fmt.Fprintf(z.out, "%s\n", t)
	} else {
		fmt.Fprintf(z.out, "\x1b[%dm%s\x1b[0m\n", color, t)
	}
}

func (z *console) boldPrint(t string) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if app.IsWindows() || !z.useColor {
		fmt.Fprintf(z.out, "%s\n", t)
	} else {
		fmt.Fprintf(z.out, "\x1b[1m%s\x1b[0m\n", t)
	}
}

func (z *console) HeaderK(key string, p ...app_msg.P) {
	z.Header(app_msg.M(key, p...))
}

func (z *console) InfoTable(name string) Table {
	tw := new(tabwriter.Writer)
	tw.Init(z.out, 0, 2, 2, ' ', 0)
	return &consoleTable{
		mc:   z.mc,
		tab:  tw,
		qm:   z.qm,
		name: name,
		ui:   z,
	}
}

func (z *console) InfoK(key string, p ...app_msg.P) {
	z.Info(app_msg.M(key, p...))
}

func (z *console) ErrorK(key string, p ...app_msg.P) {
	z.Error(app_msg.M(key, p...))
}

func (z *console) Success(m app_msg.Message) {
	z.verifyKey(m.Key())
	t := z.mc.Compile(m)
	z.colorPrint(t, ColorGreen)
	z.currentLogger().Debug(t)
}

func (z *console) Failure(m app_msg.Message) {
	z.verifyKey(m.Key())
	t := z.mc.Compile(m)
	z.colorPrint(t, ColorRed)
	z.currentLogger().Debug(t)
}

func (z *console) SuccessK(key string, p ...app_msg.P) {
	z.Success(app_msg.M(key, p...))
}

func (z *console) FailureK(key string, p ...app_msg.P) {
	z.Failure(app_msg.M(key, p...))
}

func (z *console) AskContK(key string, p ...app_msg.P) (cont bool, cancel bool) {
	return z.AskCont(app_msg.M(key, p...))
}

func (z *console) AskTextK(key string, p ...app_msg.P) (text string, cancel bool) {
	return z.AskText(app_msg.M(key, p...))
}

func (z *console) AskSecureK(key string, p ...app_msg.P) (text string, cancel bool) {
	return z.AskSecure(app_msg.M(key, p...))
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
		z.ui.Info(MConsole.Progress.
			With("Label", z.name).
			With("Progress", z.numRows))
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
		z.ui.Info(MConsole.LargeReport.With("Num", z.numRows))
	}
}
