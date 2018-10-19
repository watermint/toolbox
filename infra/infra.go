package infra

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/api/auth"
	"github.com/watermint/toolbox/infra/diag"
	"github.com/watermint/toolbox/infra/util"
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
	issuedTokens []string
}

var (
	DropboxFullAppKey           string = ""
	DropboxFullAppSecret        string = ""
	BusinessInfoAppKey          string = "7hvc6534xiq9b1w"
	BusinessInfoAppSecret       string = "08w6ex5q3j76pqv"
	BusinessFileAppKey          string = "40tbehrieaha9w3"
	BusinessFileAppSecret       string = "jh1ylouwv9wbcu4"
	BusinessManagementAppKey    string = "9ihfmr6pu82045o"
	BusinessManagementAppSecret string = "ktvr8lltw5wqyxh"
	BusinessAuditAppKey         string = "6eieiasy6iqlqx1"
	BusinessAuditAppSecret      string = "za0n7k34ca10ows"
)

var (
	AppName    string = "toolbox"
	AppVersion string = "dev"
	AppHash    string = "XXXXXXX"
)

func (opts *ExecContext) FileOnWorkPath(name string) string {
	return filepath.Join(opts.WorkPath, name)
}

func (opts *ExecContext) AuthFile() string {
	return opts.FileOnWorkPath(AppName + ".secret")
}

func (opts *ExecContext) queueToken(a auth.DropboxAuthenticator, business bool) (ac *api.ApiContext, err error) {
	token, err := a.LoadOrAuth(business, !opts.CleanupToken)
	if err != nil {
		return nil, err
	}

	seelog.Debugf("Issued token stored in ExecContext")
	opts.issuedTokens = append(opts.issuedTokens, token)

	ac = api.NewDefaultApiContext(token)

	return
}

func (opts *ExecContext) LoadOrAuthDropboxFull() (ac *api.ApiContext, err error) {
	a := auth.DropboxAuthenticator{
		AuthFile:  opts.AuthFile(),
		AppKey:    DropboxFullAppKey,
		AppSecret: DropboxFullAppSecret,
		TokenType: auth.DropboxTokenFull,
	}
	return opts.queueToken(a, false)
}

func (opts *ExecContext) LoadOrAuthBusinessInfo() (ac *api.ApiContext, err error) {
	a := auth.DropboxAuthenticator{
		AuthFile:  opts.AuthFile(),
		AppKey:    BusinessInfoAppKey,
		AppSecret: BusinessInfoAppSecret,
		TokenType: auth.DropboxTokenBusinessInfo,
	}
	return opts.queueToken(a, true)
}

func (opts *ExecContext) LoadOrAuthBusinessFile() (ac *api.ApiContext, err error) {
	a := auth.DropboxAuthenticator{
		AuthFile:  opts.AuthFile(),
		AppKey:    BusinessFileAppKey,
		AppSecret: BusinessFileAppSecret,
		TokenType: auth.DropboxTokenBusinessFile,
	}
	return opts.queueToken(a, true)
}

func (opts *ExecContext) LoadOrAuthBusinessManagement() (ac *api.ApiContext, err error) {
	a := auth.DropboxAuthenticator{
		AuthFile:  opts.AuthFile(),
		AppKey:    BusinessManagementAppKey,
		AppSecret: BusinessManagementAppSecret,
		TokenType: auth.DropboxTokenBusinessManagement,
	}
	return opts.queueToken(a, true)
}

func (opts *ExecContext) LoadOrAuthBusinessAudit() (ac *api.ApiContext, err error) {
	a := auth.DropboxAuthenticator{
		AuthFile:  opts.AuthFile(),
		AppKey:    BusinessAuditAppKey,
		AppSecret: BusinessAuditAppSecret,
		TokenType: auth.DropboxTokenBusinessAudit,
	}
	return opts.queueToken(a, true)
}

func (opts *ExecContext) Startup() error {
	err := setupWorkPath(opts)
	if err != nil {
		return err
	}

	setupLogger(opts)

	seelog.Infof("[%s] version [%s] hash[%s]", AppName, AppVersion, AppHash)

	if opts.Proxy != "" {
		SetupHttpProxy(opts.Proxy)
	}

	diag.LogDiagnostics()
	diag.LogNetworkDiagnostics()

	err = diag.QuickNetworkDiagnostics()
	if err != nil {
		return errors.New("Unable to reach `www.dropbox.com`. Please check network connection and/or proxy configuration.")
	}

	return nil
}

func (opts *ExecContext) Shutdown() {
	if opts.CleanupToken {
		for _, token := range opts.issuedTokens {
			auth.RevokeToken(token)
		}
	}
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

func (ic *ExecContext) PrepareFlags(flagset *flag.FlagSet) {
	descProxy := "HTTP/HTTPS proxy (hostname:port)"
	flagset.StringVar(&ic.Proxy, "proxy", "", descProxy)

	descWork := fmt.Sprintf("Work directory (default: %s)", DefaultWorkPath())
	flagset.StringVar(&ic.WorkPath, "work", "", descWork)

	descCleanup := "Cleanup token on exit"
	flagset.BoolVar(&ic.CleanupToken, "cleanup-token", false, descCleanup)

	descTrace := "Enable trace level log"
	flagset.BoolVar(&ic.TraceLog, "trace", false, descTrace)
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
