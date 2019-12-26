package rc_conn_impl

import (
	"errors"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_context_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"go.uber.org/zap"
)

const (
	DefaultPeerName = "default"
)

func connect(tokenType, peerName string, ctl app_control.Control) (ctx api_context.Context, err error) {
	l := ctl.Log().With(zap.String("tokenType", tokenType), zap.String("peerName", peerName))
	ui := ctl.UI()
	switch {
	case ctl.IsTest():
		l.Debug("Connect for testing")
		c := api_auth_impl.NewCached(ctl, api_auth_impl.PeerName(peerName))
		ctx, err = c.Auth(tokenType)
		return

	case ui.IsConsole():
		l.Debug("Connect through console UI")
		c := api_auth_impl.New(ctl, api_auth_impl.PeerName(peerName))
		ctx, err = c.Auth(tokenType)
		return

	case ui.IsWeb():
		l.Debug("Connect through web UI")
		a := api_auth_impl.NewWeb(ctl)
		tokens, err := a.List(tokenType)
		if err != nil {
			return nil, err
		}
		for _, t := range tokens {
			if t.PeerName == peerName {
				c := api_context_impl.New(ctl, t)
				return c, nil
			}
		}
		l.Debug("No peer found in existing connection")
		return nil, errors.New("no peer found")
	}

	l.Debug("Unsupported UI type")
	return nil, errors.New("unsupported UI type")
}
