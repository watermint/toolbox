package infra

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_auth"
	"github.com/watermint/toolbox/infra/diag"
	"github.com/watermint/toolbox/infra/util"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

const (
	logConfig = `
	<seelog type="adaptive" mininterval="200000000" maxinterval="1000000000" critmsgcount="5">
	<formats>
    		<format id="detail" format="date:%Date(2006-01-02T15:04:05Z07:00)%tloc:%File:%FuncShort:%Line%tlevel:%Level%tmsg:%Msg%n" />
    		<format id="short" format="%Time [%LEVEL] %Msg%n" />
	</formats>
	<outputs formatid="detail">
		{{if .LogPath}}
    		<filter levels="{{.LogLevels}}">
        		<rollingfile formatid="detail" filename="{{.LogPath}}" type="size" maxsize="{{.LogMaxSize}}" maxrolls="{{.LogRolls}}" />
    		</filter>
		{{end}}
		<filter levels="info,warn,error,critical">
        		<console formatid="short" />
    		</filter>
    	</outputs>
	</seelog>
	`
)

const (
	DefaultLogMaxSize = 50 * 1024 * 1024
	DefaultLogRolls   = 7
)

var (
	logPath string
)

type ExecContext struct {
	Proxy        string
	WorkPath     string
	LogPath      string
	LogMaxSize   uint64
	LogRolls     int
	LogLevels    string
	CleanupToken bool
	TraceLog     bool
	tokens       *Tokens
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

func (e *ExecContext) FileOnWorkPath(name string) string {
	return filepath.Join(e.WorkPath, name)
}

func (e *ExecContext) AuthFile() string {
	return e.FileOnWorkPath(AppName + ".secret")
}

func (e *ExecContext) queueToken(a dbx_auth.DropboxAuthenticator, business bool) (ac *dbx_api.Context, err error) {
	token, err := a.LoadOrAuth(business, !e.CleanupToken)
	if err != nil {
		return nil, err
	}

	ac = dbx_api.NewContext(token)

	return
}

func (e *ExecContext) IsTokensAvailable() bool {
	return e.tokens != nil
}

func (e *ExecContext) LoadOrAuthDropboxFull() (ac *dbx_api.Context, err error) {
	if e.tokens != nil && e.tokens.DropboxFullToken != "" {
		return dbx_api.NewContext(e.tokens.DropboxFullToken), nil
	}
	a := dbx_auth.DropboxAuthenticator{
		AuthFile:  e.AuthFile(),
		AppKey:    DropboxFullAppKey,
		AppSecret: DropboxFullAppSecret,
		TokenType: dbx_auth.DropboxTokenFull,
	}
	return e.queueToken(a, false)
}

func (e *ExecContext) LoadOrAuthBusinessInfo() (ac *dbx_api.Context, err error) {
	if e.tokens != nil && e.tokens.BusinessInfoToken != "" {
		return dbx_api.NewContext(e.tokens.BusinessInfoToken), nil
	}
	a := dbx_auth.DropboxAuthenticator{
		AuthFile:  e.AuthFile(),
		AppKey:    BusinessInfoAppKey,
		AppSecret: BusinessInfoAppSecret,
		TokenType: dbx_auth.DropboxTokenBusinessInfo,
	}
	return e.queueToken(a, true)
}

func (e *ExecContext) LoadOrAuthBusinessFile() (ac *dbx_api.Context, err error) {
	if e.tokens != nil && e.tokens.BusinessFileToken != "" {
		return dbx_api.NewContext(e.tokens.BusinessFileToken), nil
	}
	a := dbx_auth.DropboxAuthenticator{
		AuthFile:  e.AuthFile(),
		AppKey:    BusinessFileAppKey,
		AppSecret: BusinessFileAppSecret,
		TokenType: dbx_auth.DropboxTokenBusinessFile,
	}
	return e.queueToken(a, true)
}

func (e *ExecContext) LoadOrAuthBusinessManagement() (ac *dbx_api.Context, err error) {
	if e.tokens != nil && e.tokens.BusinessManagementToken != "" {
		return dbx_api.NewContext(e.tokens.BusinessManagementToken), nil
	}
	a := dbx_auth.DropboxAuthenticator{
		AuthFile:  e.AuthFile(),
		AppKey:    BusinessManagementAppKey,
		AppSecret: BusinessManagementAppSecret,
		TokenType: dbx_auth.DropboxTokenBusinessManagement,
	}
	return e.queueToken(a, true)
}

func (e *ExecContext) LoadOrAuthBusinessAudit() (ac *dbx_api.Context, err error) {
	if e.tokens != nil && e.tokens.BusinessAuditToken != "" {
		return dbx_api.NewContext(e.tokens.BusinessAuditToken), nil
	}
	a := dbx_auth.DropboxAuthenticator{
		AuthFile:  e.AuthFile(),
		AppKey:    BusinessAuditAppKey,
		AppSecret: BusinessAuditAppSecret,
		TokenType: dbx_auth.DropboxTokenBusinessAudit,
	}
	return e.queueToken(a, true)
}

func (e *ExecContext) loadAppKeysFileIfExists() {
	appKeysFile := AppName + ".appkey"
	_, err := os.Stat(appKeysFile)
	if os.IsNotExist(err) {
		return
	}

	ak, err := ioutil.ReadFile(appKeysFile)
	if err != nil {
		seelog.Debugf("Unable to load app keys file: [%s]", appKeysFile)
		return
	}
	keys := AppKey{}
	err = json.Unmarshal(ak, &keys)
	if err != nil {
		seelog.Debugf("Unable to load app keys file: [%s]", appKeysFile)
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

func (e *ExecContext) loadTokensFileIfExists() {
	pwd, _ := os.Getwd()
	seelog.Debugf("Pwd[%s]", pwd)

	tokensFile := AppName + ".tokens"
	_, err := os.Stat(tokensFile)
	if os.IsNotExist(err) {
		return
	}
	ak, err := ioutil.ReadFile(tokensFile)
	if err != nil {
		seelog.Debugf("Unable to load tokens file: [%s]", tokensFile)
		return
	}
	tokens := Tokens{}
	err = json.Unmarshal(ak, &tokens)
	if err != nil {
		seelog.Debugf("Unable to load tokens file: [%s]", tokensFile)
		return
	}

	if tokens.DropboxFullToken != "" &&
		tokens.BusinessInfoToken != "" &&
		tokens.BusinessManagementToken != "" &&
		tokens.BusinessFileToken != "" &&
		tokens.BusinessAuditToken != "" {

		seelog.Debugf("Tokens file [%s] loaded", tokensFile)
		e.tokens = &tokens
	}
}

func (e *ExecContext) StartupForTest() error {
	err := setupWorkPath(e)
	if err != nil {
		return err
	}

	setupLogger(e)
	e.loadAppKeysFileIfExists()
	e.loadTokensFileIfExists()

	return nil
}

func (e *ExecContext) Startup() error {
	err := setupWorkPath(e)
	if err != nil {
		return err
	}

	setupLogger(e)

	seelog.Infof("[%s] version [%s] hash[%s]", AppName, AppVersion, AppHash)

	e.loadAppKeysFileIfExists()
	e.loadTokensFileIfExists()

	if e.Proxy != "" {
		SetupHttpProxy(e.Proxy)
	}

	diag.LogDiagnostics()
	diag.LogNetworkDiagnostics()

	err = diag.QuickNetworkDiagnostics()
	if err != nil {
		return errors.New("unable to reach `www.dropbox.com`. Please check network connection and/or proxy configuration.")
	}

	return nil
}

func (e *ExecContext) Shutdown() {
	seelog.Trace("Shutdown infrastructure")
	seelog.Infof("Log file is at [%s]", logPath)
	seelog.Flush()
}

func DefaultWorkPath() string {
	u, err := user.Current()
	if err != nil {
		log.Fatalf("Unable to determine current user: %v", err)
		panic(err)
	}
	return filepath.Join(u.HomeDir, "."+AppName)
}

func (e *ExecContext) PrepareFlags(flagset *flag.FlagSet) {
	descProxy := "HTTP/HTTPS proxy (hostname:port)"
	flagset.StringVar(&e.Proxy, "proxy", "", descProxy)

	descWork := fmt.Sprintf("Work directory (default: %s)", DefaultWorkPath())
	flagset.StringVar(&e.WorkPath, "work", "", descWork)

	descCleanup := "Cleanup token on exit"
	flagset.BoolVar(&e.CleanupToken, "cleanup-token", false, descCleanup)

	descTrace := "Enable trace level log"
	flagset.BoolVar(&e.TraceLog, "trace", false, descTrace)
}

func setupWorkPath(opts *ExecContext) error {
	if opts.WorkPath == "" {
		opts.WorkPath = DefaultWorkPath()
		log.Printf("Setup using default work path: [%s]", opts.WorkPath)
	}

	st, err := os.Stat(opts.WorkPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(opts.WorkPath, 0701)
			if err != nil {
				log.Fatalf("Unable to create work directory: [%s]", opts.WorkPath)
				return err
			}
		} else {
			return err
		}
	} else {
		if !st.IsDir() {
			return errors.New(fmt.Sprintf("Unable to create work directory, it's not directory: [%s]. ", opts.WorkPath))
		}
		if st.Mode()&0700 == 0 {
			return errors.New(fmt.Sprintf("Unable to read/write work directory: %s", opts.WorkPath))
		}
	}

	return nil
}

func SetupHttpProxy(proxy string) {
	if proxy != "" {
		seelog.Debugf("Proxy configuration: HTTP_PROXY[%s]", proxy)
		seelog.Debugf("Proxy configuration: HTTPS_PROXY[%s]", proxy)
		os.Setenv("HTTP_PROXY", proxy)
		os.Setenv("HTTPS_PROXY", proxy)
	}
}

func setupLogger(opts *ExecContext) {
	if opts.LogMaxSize < 1 {
		opts.LogMaxSize = DefaultLogMaxSize
	}
	if opts.LogRolls < 1 {
		opts.LogRolls = DefaultLogRolls
	}
	if opts.LogPath == "" {
		opts.LogPath = filepath.Join(opts.WorkPath, AppName+".log")
	}

	logPath = opts.LogPath

	if opts.TraceLog {
		opts.LogLevels = "trace,debug,info,warn,error,critical"
	} else {
		opts.LogLevels = "debug,info,warn,error,critical"
	}

	conf, err := util.CompileTemplate(logConfig, opts)
	if err != nil {
		log.Fatalf("Unable to create log config template: %s", err)
		panic(err)
	}
	logger, err := seelog.LoggerFromConfigAsString(conf)
	if err != nil {
		log.Fatalln("Unable to configure seelog", err)
		panic(err)
	} else {
		seelog.ReplaceLogger(logger)
	}

	seelog.Debugf("Logging started: file[%s] maxSize[%d] rolls[%d]", opts.LogPath, opts.LogMaxSize, opts.LogRolls)
}
