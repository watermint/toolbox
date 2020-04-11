package gh_conn_impl

import (
	"errors"
	"github.com/watermint/toolbox/domain/github/api/gh_auth"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/api/gh_context_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"go.uber.org/zap"
)

func NewConnGithubRepo(name string) gh_conn.ConnGithubRepo {
	return &ConnGithubRepo{
		name: name,
	}
}

var (
	ErrorUnsupportedUI      = errors.New("unsupported UI for this auth scope")
	ErrorUnsupportedContext = errors.New("unsupported context type found")
)

type ConnGithubRepo struct {
	name string
	ctx  gh_context.Context
}

func (z *ConnGithubRepo) ScopeLabel() string {
	return gh_auth.ScopeRepo
}

func (z *ConnGithubRepo) Connect(ctl app_control.Control) (err error) {
	l := ctl.Log()
	ui := ctl.UI()
	scope := z.ScopeLabel()

	if c, ok := ctl.(app_control.ControlTestExtension); ok {
		if c.TestValue(qt_endtoend.CtlTestExtUseMock) == true {
			l.Debug("Test with mock")
			z.ctx = gh_context_impl.NewMock()
			return nil
		}
	}
	if ctl.Feature().IsTest() && qt_endtoend.IsSkipEndToEndTest() {
		l.Debug("Skip end to end test")
		z.ctx = gh_context_impl.NewMock()
		return nil
	}
	if !ui.IsConsole() {
		l.Debug("non console UI is not supported")
		return ErrorUnsupportedUI
	}
	a := gh_auth.New(ctl, z.name)
	l.Debug("Start auth sequence", zap.String("scope", scope))
	ac, err := a.Auth(scope)
	if err != nil {
		return err
	}
	z.ctx = gh_context_impl.New(ctl, z.PeerName(), scope, ac)
	return nil
}

func (z *ConnGithubRepo) PeerName() string {
	return z.name
}

func (z *ConnGithubRepo) SetPeerName(name string) {
	z.name = name
}

func (z *ConnGithubRepo) Context() gh_context.Context {
	return z.ctx
}

func (z *ConnGithubRepo) IsRepo() bool {
	return true
}
