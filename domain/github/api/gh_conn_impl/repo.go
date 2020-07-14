package gh_conn_impl

import (
	"github.com/watermint/toolbox/domain/github/api/gh_auth"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/api/gh_context_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

func NewConnGithubRepo(name string) gh_conn.ConnGithubRepo {
	return &ConnGithubRepo{
		name: name,
	}
}

type ConnGithubRepo struct {
	name string
	ctx  gh_context.Context
}

func (z *ConnGithubRepo) ServiceName() string {
	return api_conn.ServiceGithub
}

func (z *ConnGithubRepo) ScopeLabel() string {
	return gh_auth.ScopeLabelRepo
}

func (z *ConnGithubRepo) Connect(ctl app_control.Control) (err error) {
	l := ctl.Log()
	ui := ctl.UI()
	scope := gh_auth.ScopeRepo

	if ctl.Feature().IsTestWithMock() {
		l.Debug("Test with mock")
		z.ctx = gh_context_impl.NewMock(ctl)
		return nil
	}
	if ctl.Feature().IsTest() && qt_endtoend.IsSkipEndToEndTest() {
		l.Debug("Skip end to end test")
		z.ctx = gh_context_impl.NewMock(ctl)
		return nil
	}
	if !ui.IsConsole() {
		l.Debug("non console UI is not supported")
		return qt_errors.ErrorUnsupportedUI
	}
	a := api_auth_impl.NewConsoleRedirect(ctl, z.name, gh_auth.NewApp(ctl))
	if !ctl.Feature().IsSecure() {
		l.Debug("Enable cache")
		a = api_auth_impl.NewConsoleCache(ctl, a)
	}
	l.Debug("Start auth sequence", esl.String("scope", scope))
	ac, err := a.Auth([]string{scope})
	if err != nil {
		return err
	}
	z.ctx = gh_context_impl.New(ctl, ac)
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
