package infra

import (
	"bytes"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra/diag"
	"github.com/watermint/toolbox/infra/knowledge"
	"log"
	"os"
	"text/template"
)

const (
	logConfig = `
	<seelog type="adaptive" mininterval="200000000" maxinterval="1000000000" critmsgcount="5">
	<formats>
    		<format id="detail" format="date:%Date(2006-01-02T15:04:05Z07:00)%tloc:%File:%FuncShort:%Line%tlevel:%Level%tmsg:%Msg%n" />
    		<format id="short" format="%Time [%LEVEL][%File:%FuncShort:%Line] %Msg%n" />
	</formats>
	<outputs formatid="detail">
	<!--
		{{if .LogPath}}
    		<filter levels="trace,info,warn,error,critical">
        		<rollingfile formatid="detail" filename="{{.LogPath}}" type="size" maxsize="{{.LogMaxSize}}" maxrolls="{{.LogRolls}}" />
    		</filter>
		{{end}}
		-->
		<filter levels="trace,info,warn,error,critical">
        		<console formatid="short" />
    		</filter>
    	</outputs>
	</seelog>
	`
)

type InfraOpts struct {
	Proxy      string
	WorkPath   string
	LogPath    string
	LogMaxSize uint64
	LogRolls   int
}

func InfraStartup(opts InfraOpts) error {
	replaceLogger()

	seelog.Infof("[%s] version [%s] hash[%s]", knowledge.AppName, knowledge.AppVersion, knowledge.AppHash)
	diag.LogDiagnostics()

	return nil
}

func InfraShutdown() {
	diag.LogNetworkDiagnostics()
	seelog.Trace("Shutdown infrastructure")
	seelog.Flush()
}

func SetupHttpProxy(proxy string) {
	if proxy != "" {
		seelog.Infof("Proxy configuration: HTTP_PROXY[%s]", proxy)
		seelog.Infof("Proxy configuration: HTTPS_PROXY[%s]", proxy)
		os.Setenv("HTTP_PROXY", proxy)
		os.Setenv("HTTPS_PROXY", proxy)

	}
}

func replaceLogger() {
	logger, err := seelog.LoggerFromConfigAsString(logConfig)
	if err != nil {
		log.Fatalln("Unable to configure seelog", err)
	} else {
		seelog.ReplaceLogger(logger)
	}
}

func ShowUsage(tmpl string, data interface{}) {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer

	if err := t.Execute(&buf, data); err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stderr, buf.String())
}
