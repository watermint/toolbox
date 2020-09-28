package nw_proxy

import (
	"fmt"
	"github.com/rapid7/go-get-proxied/proxy"
	"github.com/watermint/toolbox/essentials/log/esl"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	Direct  = "DIRECT"
	noProxy = ""
)

type Configurator interface {
	Get(targetUrl string) (hostPort string, found bool)
}

func Setup(targetUrl string, hostPort string, logger esl.Logger) bool {
	configurators := []Configurator{
		newValueConfigurator(hostPort, logger),
		newAutoConfigurator(logger),
		newDirectConfigurator(),
	}
	l := logger.With(esl.String("targetUrl", targetUrl), esl.String("hostPort", hostPort))
	for _, cfg := range configurators {
		hostPort, found := cfg.Get(targetUrl)
		if !found {
			continue
		}

		var client *http.Client
		if hostPort != noProxy {
			client = &http.Client{
				Transport: &http.Transport{
					Proxy: func(req *http.Request) (proxy *url.URL, err error) {
						return url.Parse(req.URL.Scheme + "://" + hostPort)
					},
				},
			}
		} else {
			client = http.DefaultClient
		}

		l.Debug("Try with a proxy setting")
		res, err := client.Head(targetUrl)
		if err != nil {
			l.Debug("Unable to retrieve, fallback to next configurator", esl.Error(err))
			continue
		}

		l.Debug("Proxy configuration looks good", esl.Int("resCode", res.StatusCode))
		if err := os.Setenv("HTTP_PROXY", hostPort); err != nil {
			l.Debug("Unable to set HTTP_PROXY", esl.Error(err))
			continue
		}
		if err := os.Setenv("HTTPS_PROXY", hostPort); err != nil {
			l.Debug("Unable to set HTTPS_PROXY", esl.Error(err))
			continue
		}

		l.Debug("Proxy configuration",
			esl.String("HTTP_PROXY", hostPort),
			esl.String("HTTPS_PROXY", hostPort),
		)
		return true
	}
	return false
}

func newValueConfigurator(val string, logger esl.Logger) Configurator {
	return &valueConfigurator{
		logger: logger,
		val:    val,
	}
}

type valueConfigurator struct {
	logger esl.Logger
	val    string
}

func (z valueConfigurator) Get(targetUrl string) (hostPort string, found bool) {
	l := z.logger.With(esl.String("value", z.val))
	if strings.ToLower(Direct) == strings.ToLower(z.val) {
		l.Debug("Prefer direct connection")
		return noProxy, true
	}

	if z.val != "" {
		l.Debug("Configuration found")
		return z.val, true
	}

	return noProxy, false
}

func newAutoConfigurator(logger esl.Logger) Configurator {
	return &autoConfigurator{
		logger: logger,
	}
}

type autoConfigurator struct {
	logger esl.Logger
}

func (z autoConfigurator) Get(targetUrl string) (hostPort string, found bool) {
	l := z.logger
	detect := proxy.NewProvider("").GetHTTPSProxy(targetUrl)
	if detect == nil {
		l.Debug("No proxy detected. Use direct connection")
		return noProxy, true
	}

	usr, usrSpecified := detect.Username()
	l.Debug("Proxy configuration detected",
		esl.String("host", detect.Host()),
		esl.Uint16("port", detect.Port()),
		esl.Bool("user_auth", usrSpecified),
		esl.String("username", usr),
	)
	if usrSpecified {
		l.Debug("Skip proxy auto detect config because Basic Auth Proxy config not supported")
		return noProxy, true
	}
	hostPort = fmt.Sprintf("%s:%d", detect.Host(), detect.Port())
	l.Debug("Host and port determined", esl.String("hostPort", hostPort))
	return hostPort, true
}

func newDirectConfigurator() Configurator {
	return &directConfigurator{}
}

type directConfigurator struct {
}

func (z directConfigurator) Get(targetUrl string) (hostPort string, found bool) {
	return noProxy, true
}
