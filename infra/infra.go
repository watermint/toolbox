package infra

import (
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra/diag"
	"github.com/watermint/toolbox/infra/knowledge"
	"log"
	"os"
)

const (
	logConfig = `
	<seelog type="adaptive" mininterval="200000000" maxinterval="1000000000" critmsgcount="5">
	<formats>
    		<format id="detail" format="date:%Date(2006-01-02T15:04:05Z07:00)%tloc:%File:%FuncShort:%Line%tlevel:%Level%tmsg:%Msg%n" />
    		<format id="short" format="%Time [%LEVEL][%File:%FuncShort:%Line] %Msg%n" />
	</formats>
	<outputs formatid="detail">
		<filter levels="info,warn,error,critical">
        		<console formatid="short" />
    		</filter>
    	</outputs>
	</seelog>
	`
)

func InfraStartup() error {
	replaceLogger()

	seelog.Infof("[%s] version [%s] hash[%s]", knowledge.AppName, knowledge.AppVersion, knowledge.AppHash)
	diag.LogDiagnostics()

	return nil
}

func InfraShutdown() {
	seelog.Trace("Shutdown infrastructure")
	seelog.Flush()
}

func SetupHttpProxy(proxy string) {
	if proxy != "" {
		seelog.Info("Proxy configuration: HTTP_PROXY[%s]", proxy)
		seelog.Info("Proxy configuration: HTTPS_PROXY[%s]", proxy)
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
