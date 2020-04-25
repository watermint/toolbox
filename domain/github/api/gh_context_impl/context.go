package gh_context_impl

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/api/gh_request"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
	"strings"
)

var (
	ErrorGeneralApiError = errors.New("general api error")
)

func NewNoAuth(ctl app_control.Control) gh_context.Context {
	return &Context{
		scope:     "",
		ac:        nil,
		ctl:       ctl,
		isNoRetry: false,
	}
}

func New(ctl app_control.Control, peerName, scope string, ac api_auth.Context) gh_context.Context {
	return &Context{
		peerName:  peerName,
		scope:     scope,
		ac:        ac,
		ctl:       ctl,
		isNoRetry: false,
	}
}

type Context struct {
	peerName  string
	scope     string
	ac        api_auth.Context
	ctl       app_control.Control
	isNoRetry bool
}

func (z Context) Feature() app_feature.Feature {
	return z.ctl.Feature()
}

func (z Context) ClientHash() string {
	tok := ""
	if z.ac != nil && z.ac.Token() != nil {
		tok = z.ac.Token().AccessToken
	}
	seeds := []string{
		"p", z.peerName,
		"s", z.scope,
		"t", tok,
	}
	return fmt.Sprintf("%x", sha256.Sum224([]byte(strings.Join(seeds, ","))))
}

func (z Context) Log() *zap.Logger {
	return z.ctl.Log()
}

func (z Context) Capture() *zap.Logger {
	return z.ctl.Capture()
}

func (z Context) NoRetryOnError() api_context.Context {
	z.isNoRetry = true
	return &z
}

func (z Context) IsNoRetry() bool {
	return z.isNoRetry
}

func (z Context) Post(endpoint string) api_request.Request {
	return gh_request.NewRpc(&z, z.scope, z.ac.Token(), endpoint, "POST")
}

func (z Context) Get(endpoint string) api_request.Request {
	return gh_request.NewRpc(&z, z.scope, z.ac.Token(), endpoint, "GET")
}

func (z Context) Upload(endpoint string, content ut_io.ReadRewinder) api_request.Request {
	return gh_request.NewUpload(&z, z.scope, z.ac.Token(), endpoint, "POST", content)
}
