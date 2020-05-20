package nw_proxy

import (
	"fmt"
	"github.com/rapid7/go-get-proxied/proxy"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"os"
)

func SetHttpProxy(proxyHostPort string, ctl app_control.Control) {
	l := ctl.Log().With(esl.String("proxy", proxyHostPort))

	if proxyHostPort != "" {
		os.Setenv("HTTP_PROXY", proxyHostPort)
		os.Setenv("HTTPS_PROXY", proxyHostPort)
		l.Debug("Proxy configuration",
			esl.String("HTTP_PROXY", proxyHostPort),
			esl.String("HTTPS_PROXY", proxyHostPort),
		)
		return
	}

	detect := proxy.NewProvider("").GetHTTPSProxy("https://api.dropboxapi.com")
	if detect == nil {
		l.Debug("No proxy detected. Use direct connection")
		return
	}

	usr, usrSpecified := detect.Username()
	ctl.Log().Debug("Proxy configuration detected",
		esl.String("host", detect.Host()),
		esl.Uint16("port", detect.Port()),
		esl.Bool("user_auth", usrSpecified),
		esl.String("username", usr),
	)
	if usrSpecified {
		l.Debug("Skip proxy auto detect config because Basic Auth Proxy config not supported")
		return
	}

	ap := fmt.Sprintf("%s:%d", detect.Host(), detect.Port())
	os.Setenv("HTTP_PROXY", ap)
	os.Setenv("HTTPS_PROXY", ap)
	l.Debug("Proxy configuration (auto detect)",
		esl.String("HTTP_PROXY", ap),
		esl.String("HTTPS_PROXY", ap),
	)
}
