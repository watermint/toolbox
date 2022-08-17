package dbx_conn_impl

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth_attr"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
)

const (
	DefaultPeerName = "default"
)

func connect(scopes []string, peerName string, ctl app_control.Control, app api_auth.OAuthApp) (ctx dbx_context.Context, err error) {
	l := ctl.Log().With(esl.Strings("scopes", scopes), esl.String("peerName", peerName))
	ui := ctl.UI()

	if ctl.Feature().IsTestWithMock() {
		l.Debug("Test with mock")
		return dbx_context_impl.NewMock(peerName, ctl), nil
	}
	if replay, enabled := ctl.Feature().IsTestWithReplay(); enabled {
		l.Debug("Test with replay")
		return dbx_context_impl.NewReplayMock(peerName, ctl, replay), nil
	}
	if replay, enabled := ctl.Feature().IsTestWithSeqReplay(); enabled {
		l.Debug("Test with replay")
		return dbx_context_impl.NewSeqReplayMock(peerName, ctl, replay), nil
	}

	switch {
	case ctl.Feature().IsTest():
		l.Debug("Skip end to end test")
		return dbx_context_impl.NewMock(peerName, ctl), nil

	case ui.IsConsole():
		l.Debug("Connect through console UI")
		c := dbx_auth_attr.NewConsole(ctl, peerName, app)
		ctx, err := c.Start(scopes)
		if err != nil {
			return nil, err
		}
		return dbx_context_impl.New(peerName, ctl, ctx), nil

	}

	l.Debug("Unsupported UI type")
	return nil, errors.New("unsupported UI type")
}
