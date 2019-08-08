package app

import (
	"errors"
	"flag"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/rapid7/go-get-proxied/proxy"
	"github.com/watermint/toolbox/app/app_ui"
	"github.com/watermint/toolbox/app/app_util"
	"github.com/watermint/toolbox/experimental/app_root"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	AppName       = "toolbox"
	AppVersion    = "`dev`"
	AppHash       = ""
	AppZap        = ""
	AppBuilderKey = ""
	rootContext   *ExecContext
)

const (
	DefaultPeerName = "default"
	MsgNoError      = "app.common.api.err.no_error"
)

const (
	FatalGeneric = iota + 1
	FatalNoAppKey
)

func Root() *ExecContext {
	if rootContext == nil {
		rootContext = NewExecContextForTest()
		return rootContext
	}
	return rootContext
}

type ExecContext struct {
	Proxy           string
	WorkPath        string
	TokenFilePath   string
	Quiet           bool
	isTest          bool
	noCacheToken    bool
	debugMode       bool
	defaultPeerName string
	jobId           string
	userInterface   app_ui.UI
	resources       *rice.Box
	logFilePath     string
	testJobsPath    string
	lang            string
	logger          *zap.Logger
	logWrapper      *app_util.LogWrapper
	messages        *app_ui.UIMessageContainer
	values          map[string]interface{}
}

type TestOpt func(opt *testOpts) *testOpts
type testOpts struct {
	box *rice.Box
}

func WithBox(box *rice.Box) TestOpt {
	return func(opt *testOpts) *testOpts {
		opt.box = box
		return opt
	}
}

func NewExecContextForTest(opts ...TestOpt) *ExecContext {
	to := &testOpts{
		box: nil,
	}
	for _, o := range opts {
		o(to)
	}
	ec := &ExecContext{
		values: make(map[string]interface{}),
	}
	ec.isTest = true
	ec.debugMode = true
	ec.resources = to.box
	ec.startup()
	return ec
}

func NewExecContext(bx *rice.Box) (ec *ExecContext, err error) {
	ec = &ExecContext{
		values: make(map[string]interface{}),
	}
	ec.isTest = false
	ec.resources = bx
	if err := ec.startup(); err != nil {
		return nil, err
	}
	return ec, nil
}

func (z *ExecContext) SetValue(key string, value interface{}) {
	z.values[key] = value
}

func (z *ExecContext) GetValue(key string) (value interface{}, exists bool) {
	value, exists = z.values[key]
	return
}

func (z *ExecContext) NoCacheToken() bool {
	return z.noCacheToken
}

func (z *ExecContext) DefaultPeerName() string {
	return z.defaultPeerName
}

func (z *ExecContext) IsTest() bool {
	return z.isTest
}

func (z *ExecContext) IsDebug() bool {
	return z.debugMode
}

func (z *ExecContext) UI() app_ui.UI {
	return z.userInterface
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
func (z *ExecContext) FileOnSecretsPath(name string) string {
	return filepath.Join(z.SecretsPath(), name)
}

func (z *ExecContext) startup() error {
	if z.isTest {
		z.jobId = "test-" + fmt.Sprintf(time.Now().Format("20060102-150405.000"))
		z.testJobsPath = os.TempDir()
	} else {
		z.jobId = fmt.Sprintf(time.Now().Format("20060102-150405.000"))
	}
	z.defaultPeerName = DefaultPeerName
	if runtime.GOOS == "windows" {
		z.userInterface = app_ui.NewDefaultCUI()
	} else {
		z.userInterface = app_ui.NewColorCUI()
	}
	z.setupLoggerConsole()
	z.loadMessages()
	if err := z.setupWorkPath(); err != nil {
		return err
	}
	z.setupLoggerFile()
	z.logger.Debug("Up:",
		zap.String("app", AppName),
		zap.String("version", AppVersion),
		zap.String("revision", AppHash),
	)
	z.logger.Debug("Up completed")
	if z.isTest {
		z.logger.Info("Jobs path", zap.String("path", z.JobsPath()))
	}

	// replace root exec context
	rootContext = z

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

func (z *ExecContext) applyFlagLogger() error {
	z.setupLoggerFile()
	return nil
}

func (z *ExecContext) applyFlagNetwork() error {
	z.SetupHttpProxy(z.Proxy)
	return nil
}

func (z *ExecContext) applyLang() error {
	if z.lang != "" {
		z.messages.UpdateLang(z.lang)
	}
	return nil
}

func (z *ExecContext) applyDebug() error {
	z.userInterface.DebugMode(z.debugMode)
	return nil
}

func (z *ExecContext) ApplyFlags() error {
	if err := z.applyDebug(); err != nil {
		return err
	}
	if err := z.applyFlagWorkPath(); err != nil {
		return err
	}
	if err := z.applyFlagLogger(); err != nil {
		return err
	}
	if err := z.applyLang(); err != nil {
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

func (z *ExecContext) shutdownCleanup() {

}

func (z *ExecContext) Fatal(code int) {
	if z.logWrapper != nil {
		z.logWrapper.Flush()
	}
	z.Log().Debug("Down (Abort)", zap.Int("code", code))
	z.Log().Sync()
	z.shutdownCleanup()
	os.Exit(code)
}

func (z *ExecContext) Shutdown() {
	if z.logWrapper != nil {
		z.logWrapper.Flush()
	}
	z.Log().Debug("Down")
	z.Log().Sync()
	z.shutdownCleanup()
}

func (z *ExecContext) LogsPath() string {
	return filepath.Join(z.JobsPath(), "logs")
}
func (z *ExecContext) JobsPath() string {
	if z.testJobsPath != "" {
		return filepath.Join(z.testJobsPath, z.jobId)
	} else {
		return z.FileOnWorkPath(filepath.Join("jobs", z.jobId))
	}
}
func (z *ExecContext) SecretsPath() string {
	return z.FileOnWorkPath("secrets")
}

func (z *ExecContext) DefaultWorkPath() string {
	for _, e := range os.Environ() {
		v := strings.Split(e, "=")
		if v[0] == "TOOLBOX_HOME" && len(v) > 1 {
			z.Log().Debug("Set work path from $TOOLBOX_HOME", zap.String("home", v[1]))
			return v[1]
		}
	}

	u, err := user.Current()
	if err != nil {
		z.Log().Fatal("Unable to determine current user", zap.Error(err))
	}
	return filepath.Join(u.HomeDir, "."+AppName)
}

func (z *ExecContext) PrepareFlags(f *flag.FlagSet) {
	//descWork := z.Msg("app.common.flag.work").WithArg(z.DefaultWorkPath()).T()
	//f.StringVar(&z.WorkPath, "work", "", descWork)

	descDebug := z.Msg("app.common.flag.debug").T()
	f.BoolVar(&z.debugMode, "debug", false, descDebug)

	descProxy := z.Msg("app.common.flag.proxy").T()
	f.StringVar(&z.Proxy, "proxy", "", descProxy)

	descQuiet := z.Msg("app.common.flag.quiet").T()
	f.BoolVar(&z.Quiet, "quiet", false, descQuiet)

	descAlias := z.Msg("app.common.flag.alias").T()
	f.StringVar(&z.defaultPeerName, "alias", DefaultPeerName, descAlias)

	descSecure := z.Msg("app.common.flag.secure").T()
	f.BoolVar(&z.noCacheToken, "secure", false, descSecure)

	descLang := z.Msg("app.common.flag.lang").T()
	f.StringVar(&z.lang, "lang", "", descLang)
}

func (z *ExecContext) setupDirectory(path string) error {
	failMsg := z.Msg("app.common.infra.err.failed_setup_work_dir")
	failStruct := struct {
		Path  string
		Error string
	}{
		Path: path,
	}

	st, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(path, 0701)
		if err != nil {
			failStruct.Error = err.Error()
			failMsg.WithData(failStruct).TellError()
			z.Log().Fatal("Unable to create work directory", zap.String("path", path), zap.Error(err))
			return err
		}
	} else if err != nil {
		failStruct.Error = err.Error()
		failMsg.WithData(failStruct).TellError()
		z.Log().Fatal("Unable to setup work directory", zap.String("path", path), zap.Error(err))
		return err
	} else if !st.IsDir() {
		failStruct.Error = z.Msg("app.common.infra.err.workdir_is_not_a_dir").T()
		failMsg.WithData(failStruct).TellError()
		z.Log().Fatal("Unable to setup work directory. It's not a directory", zap.String("path", path))
		return errors.New("unable to setup work directory")
	} else if st.Mode()&0700 == 0 {
		failStruct.Error = z.Msg("app.common.infra.err.workdir_no_permission").T()
		failMsg.WithData(failStruct).TellError()
		z.Log().Fatal("Unable to setup work directory. No permission to read/write work directory", zap.String("path", path))
		return errors.New("unable to setup work directory")
	}

	return nil
}

func (z *ExecContext) setupWorkPath() error {
	if z.WorkPath == "" {
		z.WorkPath = z.DefaultWorkPath()
		z.Log().Debug("Setup using default work path",
			zap.String("path", z.WorkPath),
		)
	}

	if err := z.setupDirectory(z.WorkPath); err != nil {
		return err
	}
	if err := z.setupDirectory(z.JobsPath()); err != nil {
		return err
	}
	if err := z.setupDirectory(z.LogsPath()); err != nil {
		return err
	}
	if err := z.setupDirectory(z.SecretsPath()); err != nil {
		return err
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

func (z *ExecContext) setLogger(logger *zap.Logger) {
	if z.logWrapper == nil {
		z.logWrapper = app_util.NewLogWrapper(logger)

		// route default `log` package output into the file
		log.SetOutput(z.logWrapper)
	}
	app_root.SetCompatibleLogger(logger)
	z.logger = logger
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

	level := zap.InfoLevel
	if z.debugMode {
		level = zap.DebugLevel
	}

	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(en),
		zo,
		level,
	)
}

func (z *ExecContext) setupLoggerConsole() *zap.Logger {
	if z.logger == nil {
		z.setLogger(zap.New(z.consoleLoggerCore()))
	}
	return z.logger
}

func (z *ExecContext) setupLoggerFile() {
	logPath := filepath.Join(z.LogsPath(), AppName+".log")

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

	var zo zapcore.WriteSyncer
	f, err := os.Create(logPath)
	if err != nil {
		zo = zapcore.AddSync(os.Stderr)
	} else {
		zo = zapcore.AddSync(f)
	}

	zc := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		zo,
		zap.DebugLevel,
	)

	logger := zap.New(
		zapcore.NewTee(zc, z.consoleLoggerCore()),
	).WithOptions(zap.AddCaller())

	z.setLogger(logger)
	z.logFilePath = logPath
}

func (z *ExecContext) Log() *zap.Logger {
	if z.logger == nil {
		z.setupLoggerConsole()
	}
	return z.logger
}
