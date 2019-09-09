package app_network

import (
	"fmt"
	"github.com/rapid7/go-get-proxied/proxy"
	"github.com/watermint/toolbox/infra/control/app_control"
	"go.uber.org/zap"
	"os"
)

func SetHttpProxy(proxyHostPort string, ctl app_control.Control) {
	l := ctl.Log().With(zap.String("proxy", proxyHostPort))

	if proxyHostPort != "" {
		os.Setenv("HTTP_PROXY", proxyHostPort)
		os.Setenv("HTTPS_PROXY", proxyHostPort)
		l.Debug("Proxy configuration",
			zap.String("HTTP_PROXY", proxyHostPort),
			zap.String("HTTPS_PROXY", proxyHostPort),
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
		zap.String("host", detect.Host()),
		zap.Uint16("port", detect.Port()),
		zap.Bool("user_auth", usrSpecified),
		zap.String("username", usr),
	)
	if usrSpecified {
		l.Debug("Skip proxy auto detect config because Basic Auth Proxy config not supported")
		return
	}

	ap := fmt.Sprintf("%s:%d", detect.Host(), detect.Port())
	os.Setenv("HTTP_PROXY", ap)
	os.Setenv("HTTPS_PROXY", ap)
	l.Debug("Proxy configuration (auto detect)",
		zap.String("HTTP_PROXY", ap),
		zap.String("HTTPS_PROXY", ap),
	)
}
