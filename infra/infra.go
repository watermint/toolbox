package infra

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra/diag"
	"github.com/watermint/toolbox/infra/knowledge"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/integration/auth"
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
    		<filter levels="trace,info,warn,error,critical">
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

type InfraOpts struct {
	Proxy        string
	WorkPath     string
	LogPath      string
	LogMaxSize   uint64
	LogRolls     int
	CleanupToken bool

	issuedTokens []string
}

func (opts *InfraOpts) AuthFile() string {
	return filepath.Join(opts.WorkPath, knowledge.AppName+".secret")
}

func (opts *InfraOpts) queueToken(a auth.DropboxAuthenticator, business bool) (string, error) {
	token, err := a.LoadOrAuth(business, !opts.CleanupToken)
	if err == nil {
		seelog.Debugf("Issued token stored in InfraOpts")
		opts.issuedTokens = append(opts.issuedTokens, token)
	}
	return token, err
}

func (opts *InfraOpts) LoadOrAuthDropboxFull() (string, error) {
	a := auth.DropboxAuthenticator{
		AuthFile:  opts.AuthFile(),
		AppKey:    DropboxFullAppKey,
		AppSecret: DropboxFullAppSecret,
	}
	return opts.queueToken(a, false)
}

func (opts *InfraOpts) LoadOrAuthBusinessInfo() (string, error) {
	a := auth.DropboxAuthenticator{
		AuthFile:  opts.AuthFile(),
		AppKey:    BusinessInfoAppKey,
		AppSecret: BusinessInfoAppSecret,
	}
	return a.LoadOrAuth(true, !opts.CleanupToken)
}

func (opts *InfraOpts) LoadOrAuthBusinessFile() (string, error) {
	a := auth.DropboxAuthenticator{
		AuthFile:  opts.AuthFile(),
		AppKey:    BusinessFileAppKey,
		AppSecret: BusinessFileAppSecret,
	}
	return a.LoadOrAuth(true, !opts.CleanupToken)
}

func (opts *InfraOpts) LoadOrAuthBusinessManagement() (string, error) {
	a := auth.DropboxAuthenticator{
		AuthFile:  opts.AuthFile(),
		AppKey:    BusinessManagementAppKey,
		AppSecret: BusinessManagementAppSecret,
	}
	return a.LoadOrAuth(true, !opts.CleanupToken)
}

func (opts *InfraOpts) Startup() error {
	err := setupWorkPath(opts)
	if err != nil {
		return err
	}

	setupLogger(opts)

	seelog.Infof("[%s] version [%s] hash[%s]", knowledge.AppName, knowledge.AppVersion, knowledge.AppHash)

	if opts.Proxy != "" {
		SetupHttpProxy(opts.Proxy)
	}

	diag.LogDiagnostics()
	diag.LogNetworkDiagnostics()

	err = diag.QuickNetworkDiagnostics()
	if err != nil {
		return errors.New("Unable to reach `www.dropbox.com`. Please check network connection or proxy configuration.")
	}

	return nil
}

func (opts *InfraOpts) Shutdown() {
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
	return filepath.Join(u.HomeDir, "."+knowledge.AppName)
}

func PrepareInfraFlags(flagset *flag.FlagSet) *InfraOpts {
	opts := &InfraOpts{}

	descProxy := "HTTP/HTTPS proxy (hostname:port)"
	flagset.StringVar(&opts.Proxy, "proxy", "", descProxy)

	descWork := fmt.Sprintf("Work directory (default: %s)", DefaultWorkPath())
	flagset.StringVar(&opts.WorkPath, "work", "", descWork)

	descCleanup := "Cleanup token on exit"
	flagset.BoolVar(&opts.CleanupToken, "cleanup-token", false, descCleanup)

	return opts
}

func setupWorkPath(opts *InfraOpts) error {
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
		seelog.Infof("Proxy configuration: HTTP_PROXY[%s]", proxy)
		seelog.Infof("Proxy configuration: HTTPS_PROXY[%s]", proxy)
		os.Setenv("HTTP_PROXY", proxy)
		os.Setenv("HTTPS_PROXY", proxy)
	}
}

func setupLogger(opts *InfraOpts) {
	if opts.LogMaxSize < 1 {
		opts.LogMaxSize = DefaultLogMaxSize
	}
	if opts.LogRolls < 1 {
		opts.LogRolls = DefaultLogRolls
	}
	if opts.LogPath == "" {
		opts.LogPath = filepath.Join(opts.WorkPath, knowledge.AppName+".log")
	}

	logPath = opts.LogPath

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

	seelog.Infof("Logging started: file[%s] maxSize[%d] rolls[%d]", opts.LogPath, opts.LogMaxSize, opts.LogRolls)
}

func ShowUsage(tmpl string, data interface{}) {
	t, err := util.CompileTemplate(tmpl, data)
	if err != nil {
		seelog.Errorf("Unable to create usage template: %v", err)
		panic(err)
	}
	fmt.Fprint(os.Stderr, t)
}
