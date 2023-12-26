package api_callback

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/wrapper/lgw_gin"
	"github.com/watermint/toolbox/essentials/runtime/es_open"
	escert "github.com/watermint/toolbox/essentials/security/es_cert"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_template_impl"
	"go.uber.org/atomic"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

const (
	PathPing        = "/ping"
	PathConnect     = "/connect/auth"
	PathSuccess     = "/success"
	PathFailure     = "/failure"
	PathHello       = "/hello"
	DataUriImagePng = "data:image/png;base64,"
)

type MsgCallback struct {
	MsgOpenUrlOnYourBrowser      app_msg.Message
	MsgErrorOpenUrlOnYourBrowser app_msg.Message
	MsgHitEnterToProceed         app_msg.Message
	MsgResultSuccessHeader       app_msg.Message
	MsgResultSuccessBody         app_msg.Message
	MsgResultFailureHeader       app_msg.Message
	MsgResultFailureBody         app_msg.Message
	MsgHelloHeader               app_msg.Message
	MsgHelloBody                 app_msg.Message
}

var (
	shutdownTimeout          = 5 * 1000 * time.Millisecond
	ErrorAnotherServerOnline = errors.New("another server is online")
	MCallback                = app_msg.Apply(&MsgCallback{}).(*MsgCallback)
)

type Callback interface {
	// Auth redirect url
	Url() string

	// Execute OAuth2 flow. This is blocking operation.
	Flow() error

	// Startup the server, this is blocking operation.
	Start() error

	// Shutdown the server
	Shutdown()

	// Wait for the server readiness
	WaitServerReady() bool

	// Handler for web server status check
	Ping(g *gin.Context)

	// Handler for callback
	Connect(g *gin.Context)

	// Handler for Authentication success
	Success(g *gin.Context)

	// Handler for Authentication failure
	Failure(g *gin.Context)
}

type ServerStatus struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Token   string `json:"token"`
}

type Service interface {
	// Compose redirect url that must contain csrf token param `state`.
	Url(redirectUrl string) string

	// Verify with csrf token `state`, and `code` to proceed acquire token.
	Verify(state, code string) bool
}

var (
	instanceId = atomic.Int64{}
)

func New(ctl app_control.Control, s Service, port int, secure bool) Callback {
	var opener es_open.Open
	if ctl.Feature().IsTest() || ctl.Feature().IsQuiet() {
		opener = es_open.NewTestDummy()
	} else {
		opener = es_open.New()
	}
	return &callbackImpl{
		instance: strconv.Itoa(int(instanceId.Add(1))),
		ctl:      ctl,
		service:  s,
		port:     port,
		secure:   secure,
		opener:   opener,
	}
}

func NewWithOpener(ctl app_control.Control, s Service, port int, secure bool, opener es_open.Open) Callback {
	c := New(ctl, s, port, secure)
	c.(*callbackImpl).opener = opener
	return c
}

type callbackImpl struct {
	instance        string
	service         Service
	ctl             app_control.Control
	port            int
	secure          bool
	server          *http.Server
	serverError     error
	serverToken     string
	serverReady     bool
	flowStatus      chan struct{}
	mutex           sync.Mutex
	opener          es_open.Open
	logoImageBase64 template.URL
}

func (z *callbackImpl) WaitServerReady() bool {
	for {
		if z.serverReady {
			return true
		}
		if z.serverError != nil {
			return false
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (z *callbackImpl) ping() error {
	l := z.ctl.Log().With(esl.Int("port", z.port), esl.String("instance", z.instance))

	l.Debug("waiting for the server ready")
	hc := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	pingUrl := z.urlForPath(PathPing)
	for {
		time.Sleep(100 * time.Millisecond)
		if z.serverError != nil {
			l.Debug("server startup failure", esl.Error(z.serverError))
			return z.serverError
		}

		l.Debug("ping")
		res, err := hc.Get(pingUrl)
		if err != nil {
			l.Debug("ping failure", esl.Error(err))
			continue
		}
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			l.Debug("unable to read body", esl.Error(err))
			continue
		}
		d := &ServerStatus{}
		if err := json.Unmarshal(b, d); err != nil {
			l.Debug("unable to unmarshal", esl.Error(err))
			continue
		}

		if d.Token != z.serverToken {
			l.Debug("server token unmatched", esl.String("received", d.Token), esl.String("expected", z.serverToken))
			z.Shutdown()
			return ErrorAnotherServerOnline
		}

		l.Debug("quit from waiting loop")
		z.serverReady = true
		return nil
	}
}

func (z *callbackImpl) openUrl(authUrl string) {
	l := z.ctl.Log().With(esl.Int("port", z.port), esl.String("instance", z.instance))
	ui := z.ctl.UI()

	l.Debug("opening auth url", esl.String("url", authUrl))
	ui.Info(MCallback.MsgOpenUrlOnYourBrowser)
	ui.Code(authUrl)
	ui.Break()
	ui.AskProceed(MCallback.MsgHitEnterToProceed)

	if err := z.opener.Open(authUrl); err != nil {
		l.Debug("Unable to open, ask user to open the url")
		ui.Info(MCallback.MsgErrorOpenUrlOnYourBrowser.With("Url", authUrl))
	}
}

func (z *callbackImpl) Flow() error {
	l := z.ctl.Log().With(esl.Int("port", z.port), esl.String("instance", z.instance))

	idle := make(chan struct{})
	url := z.service.Url(z.Url())

	l.Debug("starting server")
	go func() {
		err := z.Start()
		l.Debug("server finished", esl.Error(err))
		close(idle)
	}()

	defer z.Shutdown()

	// waiting for server up
	l.Debug("sending ping to the server")
	if err := z.ping(); err != nil {
		l.Debug("ping failure", esl.Error(err))
		return err
	}

	l.Debug("open url", esl.String("url", url))
	z.openUrl(url)

	// waiting for finish
	l.Debug("waiting for server startup")
	<-idle

	// waiting for the flow finish
	l.Debug("waiting for flow finish")
	<-z.flowStatus

	l.Debug("flow finished")
	return nil
}

func (z *callbackImpl) urlForPath(path string) string {
	if z.secure {
		return fmt.Sprintf("https://localhost:%d%s", z.port, path)
	} else {
		return fmt.Sprintf("http://localhost:%d%s", z.port, path)
	}
}

func (z *callbackImpl) Url() string {
	return z.urlForPath(PathConnect)
}

func (z *callbackImpl) Start() error {
	z.mutex.Lock()
	l := z.ctl.Log().With(esl.Int("port", z.port), esl.String("instance", z.instance))

	// in scope of the lock
	{
		if z.server != nil {
			z.mutex.Unlock()
			l.Debug("The server is already running")
			return nil
		}

		z.flowStatus = make(chan struct{})
		hfs := app_resource.Bundle().Web().HttpFileSystem()
		htp := app_template_impl.NewDev(hfs, z.ctl)
		htr := htp.(render.HTMLRender)
		if !z.ctl.Feature().IsDebug() {
			gin.SetMode(gin.ReleaseMode)
		}
		g := gin.New()
		g.Use(lgw_gin.GinWrapper(l))
		g.Use(lgw_gin.GinRecovery(l))
		g.GET(PathConnect, z.Connect)
		g.GET(PathFailure, z.Failure)
		g.GET(PathSuccess, z.Success)
		g.GET(PathHello, z.Hello)
		g.GET(PathPing, z.Ping)
		if err := htp.Define("result", "layout/simple.html", "pages/auth_result.html"); err != nil {
			z.mutex.Unlock()
			l.Debug("Unable to prepare templates", esl.Error(err))
			return err
		}
		g.StaticFS("/assets", hfs)
		g.HTMLRender = htr

		z.serverToken = sc_random.MustGetSecureRandomString(16)
		z.server = &http.Server{
			Addr:    fmt.Sprintf(":%d", z.port),
			Handler: g,
		}
	}
	z.mutex.Unlock()

	logoImage, err := app_resource.Bundle().Images().Bytes("watermint-toolbox-256x256.png")
	if err != nil {
		l.Debug("unable to load logo image", esl.Error(err))
		return err
	}
	z.logoImageBase64 = template.URL(DataUriImagePng + base64.StdEncoding.EncodeToString(logoImage))

	if z.secure {
		l.Debug("Starting server (https)")
		certPath := filepath.Join(z.ctl.Workspace().Job(), "cert")
		if err := os.MkdirAll(certPath, 0700); err != nil {
			l.Debug("Unable to create cert directory", esl.Error(err))
			return err
		}
		certFile := filepath.Join(certPath, "cert.pem")
		keyFile := filepath.Join(certPath, "key.pem")
		cert, key, err := escert.CreateSelfSigned(28)
		if err != nil {
			l.Debug("Unable to create self signed certificate", esl.Error(err))
			return err
		}
		if err := os.WriteFile(certFile, cert, 0600); err != nil {
			l.Debug("Unable to write cert file", esl.Error(err))
			return err
		}
		if err := os.WriteFile(keyFile, key, 0600); err != nil {
			l.Debug("Unable to write key file", esl.Error(err))
			return err
		}

		if err := z.server.ListenAndServeTLS(certFile, keyFile); err != nil && err != http.ErrServerClosed {
			l.Debug("Server finished with an error", esl.Error(err))
			z.serverError = err
			return err
		}
		l.Debug("Server finished normally")
	} else {
		l.Debug("Starting server (http)")
		if err := z.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Debug("Server finished with an error", esl.Error(err))
			z.serverError = err
			return err
		}
		l.Debug("Server finished normally")
	}

	return nil
}

func (z *callbackImpl) Shutdown() {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	l := z.ctl.Log().With(esl.Int("port", z.port), esl.String("instance", z.instance))
	if z.server == nil {
		l.Debug("Server is not yet started")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := z.server.Shutdown(ctx); err != nil {
		l.Debug("Server finished with an error", esl.Error(err))
	}
	l.Debug("Server stopped")
	z.server = nil
	z.serverToken = ""
	z.serverReady = false
	close(z.flowStatus)
	l.Debug("Flow closed")
}

func (z *callbackImpl) Ping(g *gin.Context) {
	g.JSON(
		http.StatusOK,
		&ServerStatus{
			Name:    app_definitions.Name,
			Version: app_definitions.BuildId,
			Token:   z.serverToken,
		},
	)
}

func (z *callbackImpl) Connect(g *gin.Context) {
	l := z.ctl.Log().With(esl.Int("port", z.port), esl.String("instance", z.instance))
	state := g.Query("state")
	code := g.Query("code")
	l.Debug("Verify state & code")
	if z.service.Verify(state, code) {
		l.Debug("redirect to success")
		g.Redirect(http.StatusTemporaryRedirect, PathSuccess)
	} else {
		l.Debug("redirect to failure")
		g.Redirect(http.StatusTemporaryRedirect, PathFailure)
	}
}

func (z *callbackImpl) Success(g *gin.Context) {
	l := z.ctl.Log().With(esl.Int("port", z.port), esl.String("instance", z.instance))
	ui := z.ctl.UI()
	g.HTML(
		http.StatusOK,
		"result",
		gin.H{
			"Copyright": app_definitions.Copyright,
			"LogoData":  z.logoImageBase64,
			"Header":    ui.Text(MCallback.MsgResultSuccessHeader),
			"Detail":    ui.Text(MCallback.MsgResultSuccessBody),
		},
	)
	z.Shutdown()
	l.Debug("Successfully finished")
}

func (z *callbackImpl) Failure(g *gin.Context) {
	l := z.ctl.Log().With(esl.Int("port", z.port), esl.String("instance", z.instance))
	ui := z.ctl.UI()
	g.HTML(
		http.StatusForbidden,
		"result",
		gin.H{
			"Copyright": app_definitions.Copyright,
			"LogoData":  z.logoImageBase64,
			"Header":    ui.Text(MCallback.MsgResultFailureHeader),
			"Detail":    ui.Text(MCallback.MsgResultFailureBody),
		},
	)
	z.Shutdown()
	l.Debug("Finished with failure")
}

func (z *callbackImpl) Hello(g *gin.Context) {
	l := z.ctl.Log().With(esl.Int("port", z.port), esl.String("instance", z.instance))
	ui := z.ctl.UI()
	g.HTML(
		http.StatusOK,
		"result",
		gin.H{
			"Copyright": app_definitions.Copyright,
			"LogoData":  z.logoImageBase64,
			"Header":    ui.Text(MCallback.MsgHelloHeader),
			"Detail":    ui.Text(MCallback.MsgHelloBody),
		},
	)
	l.Debug("Finished")
}
