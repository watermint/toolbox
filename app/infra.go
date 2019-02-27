package app

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/rapid7/go-get-proxied/proxy"
	"github.com/watermint/toolbox/app/app_ui"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

var (
	DropboxFullAppKey           string = ""
	DropboxFullAppSecret        string = ""
	BusinessInfoAppKey          string = ""
	BusinessInfoAppSecret       string = ""
	BusinessFileAppKey          string = ""
	BusinessFileAppSecret       string = ""
	BusinessManagementAppKey    string = ""
	BusinessManagementAppSecret string = ""
	BusinessAuditAppKey         string = ""
	BusinessAuditAppSecret      string = ""
)

var (
	AppName    string = "toolbox"
	AppVersion string = "dev"
	AppHash    string = ""
)

type AppKey struct {
	DropboxFullAppKey           string `json:"DropboxFullAppKey,omitempty"`
	DropboxFullAppSecret        string `json:"DropboxFullAppSecret,omitempty"`
	BusinessInfoAppKey          string `json:"BusinessInfoAppKey,omitempty"`
	BusinessInfoAppSecret       string `json:"BusinessInfoAppSecret,omitempty"`
	BusinessFileAppKey          string `json:"BusinessFileAppKey,omitempty"`
	BusinessFileAppSecret       string `json:"BusinessFileAppSecret,omitempty"`
	BusinessManagementAppKey    string `json:"BusinessManagementAppKey,omitempty"`
	BusinessManagementAppSecret string `json:"BusinessManagementAppSecret,omitempty"`
	BusinessAuditAppKey         string `json:"BusinessAuditAppKey,omitempty"`
	BusinessAuditAppSecret      string `json:"BusinessAuditAppSecret,omitempty"`
}

type Tokens struct {
	DropboxFullToken        string `json:"DropboxFullToken,omitempty"`
	BusinessInfoToken       string `json:"BusinessInfoToken,omitempty"`
	BusinessFileToken       string `json:"BusinessFileToken,omitempty"`
	BusinessManagementToken string `json:"BusinessManagementToken,omitempty"`
	BusinessAuditToken      string `json:"BusinessAuditToken,omitempty"`
}

type ExecContext struct {
	Proxy         string
	WorkPath      string
	TokenFilePath string
	Quiet         bool
	userInterface app_ui.UI
	resources     *rice.Box
	tokens        *Tokens
	logFilePath   string
	logger        *zap.Logger
	messages      *app_ui.UIMessageContainer
}

func NewExecContextForTest() *ExecContext {
	ec := &ExecContext{}
	ec.startup()
	return ec
}

func NewExecContext(bx *rice.Box) *ExecContext {
	ec := &ExecContext{}
	ec.resources = bx
	ec.startup()
	return ec
}

func (z *ExecContext) Msg(key string) app_ui.UIMessage {
	return z.messages.Msg(key)
}

func (z *ExecContext) FileOnWorkPath(name string) string {
	return filepath.Join(z.WorkPath, name)
}

func (z *ExecContext) AuthFile() string {
	return z.FileOnWorkPath(AppName + ".secret")
}

func (z *ExecContext) queueToken(a *dbx_auth.DropboxAuthenticator, business bool) (ac *dbx_api.Context, err error) {
	token, err := a.LoadOrAuth(business)
	if err != nil {
		return nil, err
	}

	ac = dbx_api.NewContext(z.messages, token, z.Log().With(zap.String("token", a.TokenType)))

	return
}

func (z *ExecContext) IsTokensAvailable() bool {
	return z.tokens != nil
}

func (z *ExecContext) LoadOrAuthDropboxFull() (ac *dbx_api.Context, err error) {
	if z.tokens != nil && z.tokens.DropboxFullToken != "" {
		return dbx_api.NewContext(
			z.messages,
			z.tokens.DropboxFullToken,
			z.Log().With(zap.String("token", dbx_auth.DropboxTokenFull)),
		), nil
	}
	a := dbx_auth.NewDropboxAuthenticator(
		z.AuthFile(),
		DropboxFullAppKey,
		DropboxFullAppSecret,
		dbx_auth.DropboxTokenFull,
		z.messages,
		z.Log().With(zap.String("token", dbx_auth.DropboxTokenFull)),
	)
	return z.queueToken(a, false)
}

func (z *ExecContext) LoadOrAuthBusinessInfo() (ac *dbx_api.Context, err error) {
	if z.tokens != nil && z.tokens.BusinessInfoToken != "" {
		return dbx_api.NewContext(
			z.messages,
			z.tokens.BusinessInfoToken,
			z.Log().With(zap.String("token", dbx_auth.DropboxTokenBusinessInfo)),
		), nil
	}
	a := dbx_auth.NewDropboxAuthenticator(
		z.AuthFile(),
		BusinessInfoAppKey,
		BusinessInfoAppSecret,
		dbx_auth.DropboxTokenBusinessInfo,
		z.messages,
		z.Log().With(zap.String("token", dbx_auth.DropboxTokenBusinessInfo)),
	)
	return z.queueToken(a, true)
}

func (z *ExecContext) LoadOrAuthBusinessFile() (ac *dbx_api.Context, err error) {
	if z.tokens != nil && z.tokens.BusinessFileToken != "" {
		return dbx_api.NewContext(
			z.messages,
			z.tokens.BusinessFileToken,
			z.Log().With(zap.String("token", dbx_auth.DropboxTokenBusinessFile)),
		), nil
	}
	a := dbx_auth.NewDropboxAuthenticator(
		z.AuthFile(),
		BusinessFileAppKey,
		BusinessFileAppSecret,
		dbx_auth.DropboxTokenBusinessFile,
		z.messages,
		z.Log().With(zap.String("token", dbx_auth.DropboxTokenBusinessFile)),
	)
	return z.queueToken(a, true)
}

func (z *ExecContext) LoadOrAuthBusinessManagement() (ac *dbx_api.Context, err error) {
	if z.tokens != nil && z.tokens.BusinessManagementToken != "" {
		return dbx_api.NewContext(
			z.messages,
			z.tokens.BusinessManagementToken,
			z.Log().With(zap.String("token", dbx_auth.DropboxTokenBusinessManagement)),
		), nil
	}
	a := dbx_auth.NewDropboxAuthenticator(
		z.AuthFile(),
		BusinessManagementAppKey,
		BusinessManagementAppSecret,
		dbx_auth.DropboxTokenBusinessManagement,
		z.messages,
		z.Log().With(zap.String("token", dbx_auth.DropboxTokenBusinessManagement)),
	)
	return z.queueToken(a, true)
}

func (z *ExecContext) LoadOrAuthBusinessAudit() (ac *dbx_api.Context, err error) {
	if z.tokens != nil && z.tokens.BusinessAuditToken != "" {
		return dbx_api.NewContext(
			z.messages,
			z.tokens.BusinessAuditToken,
			z.Log().With(zap.String("token", dbx_auth.DropboxTokenBusinessAudit)),
		), nil
	}
	a := dbx_auth.NewDropboxAuthenticator(
		z.AuthFile(),
		BusinessAuditAppKey,
		BusinessAuditAppSecret,
		dbx_auth.DropboxTokenBusinessAudit,
		z.messages,
		z.Log().With(zap.String("token", dbx_auth.DropboxTokenBusinessAudit)),
	)
	return z.queueToken(a, true)
}

func (z *ExecContext) loadAppKeysFileIfExists() {
	appKeysFile := AppName + ".appkey"
	_, err := os.Stat(appKeysFile)
	if os.IsNotExist(err) {
		return
	}

	ak, err := ioutil.ReadFile(appKeysFile)
	if err != nil {
		z.Log().Debug(
			"Unable to read app keys file",
			zap.String("file", appKeysFile),
			zap.Error(err),
		)
		return
	}
	keys := AppKey{}
	err = json.Unmarshal(ak, &keys)
	if err != nil {
		z.Log().Debug(
			"Unable to unmarshal app keys file",
			zap.String("file", appKeysFile),
			zap.Error(err),
		)
		return
	}

	if keys.DropboxFullAppKey != "" {
		DropboxFullAppKey = keys.DropboxFullAppKey
	}
	if keys.DropboxFullAppSecret != "" {
		DropboxFullAppSecret = keys.DropboxFullAppSecret
	}
	if keys.BusinessInfoAppKey != "" {
		BusinessInfoAppKey = keys.BusinessInfoAppKey
	}
	if keys.BusinessInfoAppSecret != "" {
		BusinessInfoAppSecret = keys.BusinessInfoAppSecret
	}
	if keys.BusinessFileAppKey != "" {
		BusinessFileAppKey = keys.BusinessFileAppKey
	}
	if keys.BusinessFileAppSecret != "" {
		BusinessFileAppSecret = keys.BusinessFileAppSecret
	}
	if keys.BusinessManagementAppKey != "" {
		BusinessManagementAppKey = keys.BusinessManagementAppKey
	}
	if keys.BusinessManagementAppSecret != "" {
		BusinessManagementAppSecret = keys.BusinessManagementAppSecret
	}
	if keys.BusinessAuditAppKey != "" {
		BusinessAuditAppKey = keys.BusinessAuditAppKey
	}
	if keys.BusinessAuditAppSecret != "" {
		BusinessAuditAppSecret = keys.BusinessAuditAppSecret
	}
}

func (z *ExecContext) loadTokensFileIfExists(tokensFilePath string) {
	tokensFile := filepath.Join(tokensFilePath, AppName+".tokens")
	_, err := os.Stat(tokensFile)
	if os.IsNotExist(err) {
		return
	}
	ak, err := ioutil.ReadFile(tokensFile)
	if err != nil {
		z.Log().Debug(
			"Unable to read tokens file",
			zap.String("file", tokensFile),
			zap.Error(err),
		)
		return
	}
	tokens := Tokens{}
	err = json.Unmarshal(ak, &tokens)
	if err != nil {
		z.Log().Debug(
			"Unable to unmarshal tokens file",
			zap.String("file", tokensFile),
			zap.Error(err),
		)
		return
	}

	if tokens.DropboxFullToken != "" &&
		tokens.BusinessInfoToken != "" &&
		tokens.BusinessManagementToken != "" &&
		tokens.BusinessFileToken != "" &&
		tokens.BusinessAuditToken != "" {

		z.Log().Debug(
			"Token file loaded",
			zap.String("file", tokensFile),
		)
		z.tokens = &tokens
	}
}

func (z *ExecContext) startup() error {
	z.setupLoggerConsole()
	z.userInterface = app_ui.NewDefaultCUI()
	z.loadMessages()

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

func (z *ExecContext) applyFlagAppKeys() error {
	z.loadAppKeysFileIfExists()
	z.loadTokensFileIfExists(z.TokenFilePath)
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
	if err := z.applyFlagAppKeys(); err != nil {
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

	descWork := z.Msg("app.common.flag.work").WithArg(z.DefaultWorkPath()).Text()
	f.StringVar(&z.WorkPath, "work", "", descWork)

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
	z.logger.Debug("logger started",
		zap.String("app", AppName),
		zap.String("version", AppVersion),
		zap.String("revision", AppHash),
		zap.String("logfile", logPath),
	)
}

func (z *ExecContext) Log() *zap.Logger {
	if z.logger == nil {
		z.setupLoggerConsole()
	}
	return z.logger
}
