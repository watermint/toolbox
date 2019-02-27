package app

import (
	"errors"
	"flag"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/rapid7/go-get-proxied/proxy"
	"github.com/watermint/toolbox/app/app_ui"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

var (
	AppName    string = "toolbox"
	AppVersion string = "`dev`"
	AppHash    string = ""
)

type ExecContext struct {
	Proxy         string
	WorkPath      string
	TokenFilePath string
	Quiet         bool
	isTest        bool
	userInterface app_ui.UI
	resources     *rice.Box
	logFilePath   string
	logger        *zap.Logger
	messages      *app_ui.UIMessageContainer
}

func NewExecContextForTest() *ExecContext {
	ec := &ExecContext{}
	ec.isTest = true
	ec.startup()
	return ec
}

func NewExecContext(bx *rice.Box) *ExecContext {
	ec := &ExecContext{}
	ec.isTest = false
	ec.resources = bx
	ec.startup()
	return ec
}

func (z *ExecContext) IsTest() bool {
	return z.isTest
}

func (z *ExecContext) Msg(key string) app_ui.UIMessage {
	return z.messages.Msg(key)
}

func (z *ExecContext) MessageContainer() *app_ui.UIMessageContainer {
	return z.messages
}

func (z *ExecContext) ResourceBytes(path string) ([]byte, error) {
	if z.resources == nil {
		return nil, errors.New("`resources` not found")
	}
	return z.resources.Bytes(path)
}

func (z *ExecContext) FileOnWorkPath(name string) string {
	return filepath.Join(z.WorkPath, name)
}

func (z *ExecContext) AuthFile() string {
	return z.FileOnWorkPath(AppName + ".secret")
}

func (z *ExecContext) startup() error {
	z.setupLoggerConsole()
	z.setupWorkPath()
	z.setupLoggerFile()
	z.logger.Debug("Startup:",
		zap.String("app", AppName),
		zap.String("version", AppVersion),
		zap.String("revision", AppHash),
	)
	z.userInterface = app_ui.NewDefaultCUI()
	z.loadMessages()
	z.logger.Debug("Startup completed")

	return nil
}

func (z *ExecContext) StartupMessage() {
	if !z.Quiet {
		z.Msg("app.common.name").WithData(struct {
			Version string
		}{
			Version: AppVersion,
		}).Tell()
		z.Msg("app.common.license").Tell()
	}
}

func (z *ExecContext) loadMessages() {
	z.messages = app_ui.NewUIMessageContainer(z.resources, z.userInterface, z.logger)
	z.messages.Load()
}

func (z *ExecContext) applyFlagWorkPath() error {
	err := z.setupWorkPath()
	if err != nil {
		return err
	}

	z.setupLoggerFile()
	return nil
}

func (z *ExecContext) applyFlagNetwork() error {
	z.SetupHttpProxy(z.Proxy)
	return nil
}

func (z *ExecContext) ApplyFlags() error {
	if err := z.applyFlagWorkPath(); err != nil {
		return err
	}
	if err := z.applyFlagNetwork(); err != nil {
		return err
	}

	d := Diag{
		ExecContext: z,
	}
	if err := d.Runtime(); err != nil {
		return err
	}
	if err := d.Network(); err != nil {
		return err
	}
	z.StartupMessage()

	return nil
}

func (z *ExecContext) Shutdown() {
	z.Log().Debug("Shutdown")
	z.Log().Sync()
}

func (z *ExecContext) DefaultWorkPath() string {
	u, err := user.Current()
	if err != nil {
		z.Log().Fatal(
			"Unable to determine current user",
			zap.Error(err),
		)
	}
	return filepath.Join(u.HomeDir, "."+AppName)
}

func (z *ExecContext) PrepareFlags(f *flag.FlagSet) {
	descProxy := z.Msg("app.common.flag.proxy").Text()
	f.StringVar(&z.Proxy, "proxy", "", descProxy)

	//descWork := z.Msg("app.common.flag.work").WithArg(z.DefaultWorkPath()).Text()
	//f.StringVar(&z.WorkPath, "work", "", descWork)

	descQuiet := z.Msg("app.common.flag.quiet").Text()
	f.BoolVar(&z.Quiet, "quiet", false, descQuiet)
}

func (z *ExecContext) setupWorkPath() error {
	if z.WorkPath == "" {
		z.WorkPath = z.DefaultWorkPath()
		z.Log().Debug("Setup using default work path",
			zap.String("path", z.WorkPath),
		)
	}

	st, err := os.Stat(z.WorkPath)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(z.WorkPath, 0701)
		if err == nil {
			z.Log().Info(
				"Work directory created",
				zap.String("path", z.WorkPath),
			)
		} else {
			z.Log().Fatal(
				"Unable to create work directory",
				zap.String("path", z.WorkPath),
				zap.Error(err),
			)
			return err
		}
	} else if err != nil {
		z.Log().Fatal(
			"Unable to setup work directory",
			zap.String("path", z.WorkPath),
			zap.Error(err),
		)
	} else if !st.IsDir() {
		z.Log().Fatal(
			"Unable to setup work directory. It's not a directory",
			zap.String("path", z.WorkPath),
		)
		return errors.New("unable to setup work directory")
	} else if st.Mode()&0700 == 0 {
		z.Log().Fatal(
			"Unable to setup work directory. No permission to read/write work directory",
			zap.String("path", z.WorkPath),
		)
		return errors.New("unable to setup work directory")
	}

	return nil
}

func (z *ExecContext) SetupHttpProxy(p string) {
	if p != "" {
		os.Setenv("HTTP_PROXY", p)
		os.Setenv("HTTPS_PROXY", p)
		z.Log().Debug("Proxy configuration",
			zap.String("HTTP_PROXY", p),
			zap.String("HTTPS_PROXY", p),
		)
		return
	}

	detect := proxy.NewProvider("").GetHTTPSProxy("https://api.dropboxapi.com")
	if detect == nil {
		z.Log().Debug("No proxy detected. Use direct connection")
		return
	}

	usr, usrSpecified := detect.Username()
	z.Log().Debug("Proxy configuration detected",
		zap.String("host", detect.Host()),
		zap.Uint16("port", detect.Port()),
		zap.Bool("user_auth", usrSpecified),
		zap.String("username", usr),
	)
	if usrSpecified {
		z.Log().Debug("Skip proxy auto detect config because Basic Auth Proxy config not supported")
		return
	}

	ap := fmt.Sprintf("%s:%d", detect.Host(), detect.Port())
	os.Setenv("HTTP_PROXY", ap)
	os.Setenv("HTTPS_PROXY", ap)
	z.Log().Debug("Proxy configuration (auto detect)",
		zap.String("HTTP_PROXY", ap),
		zap.String("HTTPS_PROXY", ap),
	)
}

func (z *ExecContext) consoleLoggerCore() zapcore.Core {
	en := zapcore.EncoderConfig{
		LevelKey:       "level",
		MessageKey:     "msg",
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	if runtime.GOOS == "windows" {
		en.EncodeLevel = zapcore.CapitalLevelEncoder
	} else {
		en.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	zo := zapcore.AddSync(os.Stdout)
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(en),
		zo,
		zap.InfoLevel,
	)
}

func (z *ExecContext) setupLoggerConsole() *zap.Logger {
	if z.logger == nil {
		z.logger = zap.New(z.consoleLoggerCore())
	}
	return z.logger
}

func (z *ExecContext) setupLoggerFile() {
	logPath := filepath.Join(z.WorkPath, AppName+".log")
	if z.logFilePath == logPath {
		z.Log().Debug("Skip setup logger file (path unchanged)",
			zap.String("path", logPath),
		)
		return
	}

	cfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "trace",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	zo := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    50, // megabytes
		MaxBackups: 10,
		MaxAge:     28, // days
	})

	// route default `log` package output into the file
	log.SetOutput(zo)
	zc := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		zo,
		zap.DebugLevel,
	)

	logger := zap.New(
		zapcore.NewTee(zc, z.consoleLoggerCore()),
	).WithOptions(zap.AddCaller())

	z.logger = logger
	z.logFilePath = logPath
}

func (z *ExecContext) Log() *zap.Logger {
	if z.logger == nil {
		z.setupLoggerConsole()
	}
	return z.logger
}
