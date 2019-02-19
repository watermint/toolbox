package app

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/rapid7/go-get-proxied/proxy"
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

type ExecContext struct {
	Proxy         string
	WorkPath      string
	TokenFilePath string

	tokens      *Tokens
	logFilePath string
	logger      *zap.Logger
}

func NewExecContext() *ExecContext {
	ec := &ExecContext{}
	ec.startup()
	return ec
}

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

func (ec *ExecContext) FileOnWorkPath(name string) string {
	return filepath.Join(ec.WorkPath, name)
}

func (ec *ExecContext) AuthFile() string {
	return ec.FileOnWorkPath(AppName + ".secret")
}

func (ec *ExecContext) queueToken(a dbx_auth.DropboxAuthenticator, business bool) (ac *dbx_api.Context, err error) {
	token, err := a.LoadOrAuth(business)
	if err != nil {
		return nil, err
	}

	ac = dbx_api.NewContext(token, ec.Log().With(zap.String("token", a.TokenType)))

	return
}

func (ec *ExecContext) IsTokensAvailable() bool {
	return ec.tokens != nil
}

func (ec *ExecContext) LoadOrAuthDropboxFull() (ac *dbx_api.Context, err error) {
	if ec.tokens != nil && ec.tokens.DropboxFullToken != "" {
		return dbx_api.NewContext(
			ec.tokens.DropboxFullToken,
			ec.Log().With(zap.String("token", dbx_auth.DropboxTokenFull)),
		), nil
	}
	a := dbx_auth.DropboxAuthenticator{
		AuthFile:  ec.AuthFile(),
		AppKey:    DropboxFullAppKey,
		AppSecret: DropboxFullAppSecret,
		TokenType: dbx_auth.DropboxTokenFull,
		Logger:    ec.Log().With(zap.String("token", dbx_auth.DropboxTokenFull)),
	}
	return ec.queueToken(a, false)
}

func (ec *ExecContext) LoadOrAuthBusinessInfo() (ac *dbx_api.Context, err error) {
	if ec.tokens != nil && ec.tokens.BusinessInfoToken != "" {
		return dbx_api.NewContext(
			ec.tokens.BusinessInfoToken,
			ec.Log().With(zap.String("token", dbx_auth.DropboxTokenBusinessInfo)),
		), nil
	}
	a := dbx_auth.DropboxAuthenticator{
		AuthFile:  ec.AuthFile(),
		AppKey:    BusinessInfoAppKey,
		AppSecret: BusinessInfoAppSecret,
		TokenType: dbx_auth.DropboxTokenBusinessInfo,
		Logger:    ec.Log().With(zap.String("token", dbx_auth.DropboxTokenBusinessInfo)),
	}
	return ec.queueToken(a, true)
}

func (ec *ExecContext) LoadOrAuthBusinessFile() (ac *dbx_api.Context, err error) {
	if ec.tokens != nil && ec.tokens.BusinessFileToken != "" {
		return dbx_api.NewContext(
			ec.tokens.BusinessFileToken,
			ec.Log().With(zap.String("token", dbx_auth.DropboxTokenBusinessFile)),
		), nil
	}
	a := dbx_auth.DropboxAuthenticator{
		AuthFile:  ec.AuthFile(),
		AppKey:    BusinessFileAppKey,
		AppSecret: BusinessFileAppSecret,
		TokenType: dbx_auth.DropboxTokenBusinessFile,
		Logger:    ec.Log().With(zap.String("token", dbx_auth.DropboxTokenBusinessFile)),
	}
	return ec.queueToken(a, true)
}

func (ec *ExecContext) LoadOrAuthBusinessManagement() (ac *dbx_api.Context, err error) {
	if ec.tokens != nil && ec.tokens.BusinessManagementToken != "" {
		return dbx_api.NewContext(
			ec.tokens.BusinessManagementToken,
			ec.Log().With(zap.String("token", dbx_auth.DropboxTokenBusinessManagement)),
		), nil
	}
	a := dbx_auth.DropboxAuthenticator{
		AuthFile:  ec.AuthFile(),
		AppKey:    BusinessManagementAppKey,
		AppSecret: BusinessManagementAppSecret,
		TokenType: dbx_auth.DropboxTokenBusinessManagement,
		Logger:    ec.Log().With(zap.String("token", dbx_auth.DropboxTokenBusinessManagement)),
	}
	return ec.queueToken(a, true)
}

func (ec *ExecContext) LoadOrAuthBusinessAudit() (ac *dbx_api.Context, err error) {
	if ec.tokens != nil && ec.tokens.BusinessAuditToken != "" {
		return dbx_api.NewContext(
			ec.tokens.BusinessAuditToken,
			ec.Log().With(zap.String("token", dbx_auth.DropboxTokenBusinessAudit)),
		), nil
	}
	a := dbx_auth.DropboxAuthenticator{
		AuthFile:  ec.AuthFile(),
		AppKey:    BusinessAuditAppKey,
		AppSecret: BusinessAuditAppSecret,
		TokenType: dbx_auth.DropboxTokenBusinessAudit,
		Logger:    ec.Log().With(zap.String("token", dbx_auth.DropboxTokenBusinessAudit)),
	}
	return ec.queueToken(a, true)
}

func (ec *ExecContext) loadAppKeysFileIfExists() {
	appKeysFile := AppName + ".appkey"
	_, err := os.Stat(appKeysFile)
	if os.IsNotExist(err) {
		return
	}

	ak, err := ioutil.ReadFile(appKeysFile)
	if err != nil {
		ec.Log().Debug(
			"Unable to read app keys file",
			zap.String("file", appKeysFile),
			zap.Error(err),
		)
		return
	}
	keys := AppKey{}
	err = json.Unmarshal(ak, &keys)
	if err != nil {
		ec.Log().Debug(
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

func (ec *ExecContext) loadTokensFileIfExists(tokensFilePath string) {
	tokensFile := filepath.Join(tokensFilePath, AppName+".tokens")
	_, err := os.Stat(tokensFile)
	if os.IsNotExist(err) {
		return
	}
	ak, err := ioutil.ReadFile(tokensFile)
	if err != nil {
		ec.Log().Debug(
			"Unable to read tokens file",
			zap.String("file", tokensFile),
			zap.Error(err),
		)
		return
	}
	tokens := Tokens{}
	err = json.Unmarshal(ak, &tokens)
	if err != nil {
		ec.Log().Debug(
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

		ec.Log().Debug(
			"Token file loaded",
			zap.String("file", tokensFile),
		)
		ec.tokens = &tokens
	}
}

func (ec *ExecContext) startup() error {
	ec.setupLoggerConsole()
	return nil
}

func (ec *ExecContext) applyFlagWorkPath() error {
	err := ec.setupWorkPath()
	if err != nil {
		return err
	}

	ec.setupLoggerFile()
	return nil
}

func (ec *ExecContext) applyFlagAppKeys() error {
	ec.loadAppKeysFileIfExists()
	ec.loadTokensFileIfExists(ec.TokenFilePath)
	return nil
}

func (ec *ExecContext) applyFlagNetwork() error {
	ec.SetupHttpProxy(ec.Proxy)
	return nil
}

func (ec *ExecContext) ApplyFlags() error {
	if err := ec.applyFlagWorkPath(); err != nil {
		return err
	}
	if err := ec.applyFlagAppKeys(); err != nil {
		return err
	}
	if err := ec.applyFlagNetwork(); err != nil {
		return err
	}

	d := Diag{
		ExecContext: ec,
	}
	if err := d.Runtime(); err != nil {
		return err
	}
	if err := d.Network(); err != nil {
		return err
	}

	return nil
}

func (ec *ExecContext) Shutdown() {
	ec.Log().Debug("Shutdown")
	ec.Log().Sync()
}

func (ec *ExecContext) DefaultWorkPath() string {
	u, err := user.Current()
	if err != nil {
		ec.Log().Fatal(
			"Unable to determine current user",
			zap.Error(err),
		)
	}
	return filepath.Join(u.HomeDir, "."+AppName)
}

func (ec *ExecContext) PrepareFlags(flagset *flag.FlagSet) {
	descProxy := "HTTP/HTTPS proxy (hostname:port)"
	flagset.StringVar(&ec.Proxy, "proxy", "", descProxy)

	descWork := fmt.Sprintf("Work directory (default: %s)", ec.DefaultWorkPath())
	flagset.StringVar(&ec.WorkPath, "work", "", descWork)
}

func (ec *ExecContext) setupWorkPath() error {
	if ec.WorkPath == "" {
		ec.WorkPath = ec.DefaultWorkPath()
		ec.Log().Debug("Setup using default work path",
			zap.String("path", ec.WorkPath),
		)
	}

	st, err := os.Stat(ec.WorkPath)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(ec.WorkPath, 0701)
		if err == nil {
			ec.Log().Info(
				"Work directory created",
				zap.String("path", ec.WorkPath),
			)
		} else {
			ec.Log().Fatal(
				"Unable to create work directory",
				zap.String("path", ec.WorkPath),
				zap.Error(err),
			)
			return err
		}
	} else if err != nil {
		ec.Log().Fatal(
			"Unable to setup work directory",
			zap.String("path", ec.WorkPath),
			zap.Error(err),
		)
	} else if !st.IsDir() {
		ec.Log().Fatal(
			"Unable to setup work directory. It's not a directory",
			zap.String("path", ec.WorkPath),
		)
		return errors.New("unable to setup work directory")
	} else if st.Mode()&0700 == 0 {
		ec.Log().Fatal(
			"Unable to setup work directory. No permission to read/write work directory",
			zap.String("path", ec.WorkPath),
		)
		return errors.New("unable to setup work directory")
	}

	return nil
}

func (ec *ExecContext) SetupHttpProxy(p string) {
	if p != "" {
		os.Setenv("HTTP_PROXY", p)
		os.Setenv("HTTPS_PROXY", p)
		ec.Log().Debug("Proxy configuration",
			zap.String("HTTP_PROXY", p),
			zap.String("HTTPS_PROXY", p),
		)
		return
	}

	detect := proxy.NewProvider("").GetHTTPSProxy("https://api.dropboxapi.com")
	if detect == nil {
		ec.Log().Debug("Proxy configuration",
			zap.String("HTTP_PROXY", ""),
			zap.String("HTTPS_PROXY", ""),
		)
		return
	}

	usr, usrSpecified := detect.Username()
	ec.Log().Debug("Proxy configuration detected",
		zap.String("host", detect.Host()),
		zap.Uint16("port", detect.Port()),
		zap.Bool("user_auth", usrSpecified),
		zap.String("username", usr),
	)
	if usrSpecified {
		ec.Log().Debug("Skip proxy auto detect config because Basic Auth Proxy config not supported")
		ec.Log().Debug("Proxy configuration",
			zap.String("HTTP_PROXY", ""),
			zap.String("HTTPS_PROXY", ""),
		)
		return
	}

	ap := fmt.Sprintf("%s:%d", detect.Host(), detect.Port())
	os.Setenv("HTTP_PROXY", ap)
	os.Setenv("HTTPS_PROXY", ap)
	ec.Log().Debug("Proxy configuration",
		zap.String("HTTP_PROXY", ap),
		zap.String("HTTPS_PROXY", ap),
	)
}

func (ec *ExecContext) consoleLoggerCore() zapcore.Core {
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

func (ec *ExecContext) setupLoggerConsole() *zap.Logger {
	if ec.logger == nil {
		ec.logger = zap.New(ec.consoleLoggerCore())
	}
	return ec.logger
}

func (ec *ExecContext) setupLoggerFile() {
	logPath := filepath.Join(ec.WorkPath, AppName+".log")
	if ec.logFilePath == logPath {
		ec.Log().Debug("Skip setup logger file (path unchanged)",
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
		zapcore.NewTee(zc, ec.consoleLoggerCore()),
	).WithOptions(zap.AddCaller())

	logger.Info("logger started",
		zap.String("app", AppName),
		zap.String("version", AppVersion),
		zap.String("revision", AppHash),
		zap.String("logfile", logPath),
	)

	ec.logger = logger
	ec.logFilePath = logPath

}

func (ec *ExecContext) Log() *zap.Logger {
	if ec.logger == nil {
		ec.setupLoggerConsole()
	}
	return ec.logger
}
