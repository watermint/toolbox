package app_conn_impl

import (
	"errors"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_context_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"go.uber.org/zap"
)

const (
	DefaultPeerName = "default"
)

func NewConnBusinessMgmt() app_conn.ConnBusinessMgmt {
	return &ConnBusinessMgmt{
		PeerName: DefaultPeerName,
	}
}

func NewConnBusinessInfo() app_conn.ConnBusinessInfo {
	return &ConnBusinessInfo{
		PeerName: DefaultPeerName,
	}
}

func NewConnBusinessAudit() app_conn.ConnBusinessAudit {
	return &ConnBusinessAudit{
		PeerName: DefaultPeerName,
	}
}

func NewConnBusinessFile() app_conn.ConnBusinessFile {
	return &ConnBusinessFile{
		PeerName: DefaultPeerName,
	}
}

func NewConnUserFile() *ConnUserFile {
	return &ConnUserFile{
		PeerName: DefaultPeerName,
	}
}

func connect(tokenType, peerName string, control app_control.Control) (ctx api_context.Context, err error) {
	l := control.Log().With(zap.String("tokenType", tokenType), zap.String("peerName", peerName))
	ui := control.UI(nil)

	switch {
	case control.IsTest():
		l.Debug("Connect for testing")
		c := api_auth_impl.NewCached(control, api_auth_impl.PeerName(peerName))
		ctx, err = c.Auth(tokenType)
		return

	case ui.IsConsole():
		l.Debug("Connect through console UI")
		c := api_auth_impl.New(control, api_auth_impl.PeerName(peerName))
		ctx, err = c.Auth(tokenType)
		return

	case ui.IsWeb():
		l.Debug("Connect through web UI")
		a := api_auth_impl.NewWeb(control)
		tokens, err := a.List(tokenType)
		if err != nil {
			return nil, err
		}
		for _, t := range tokens {
			if t.PeerName == peerName {
				c := api_context_impl.New(control, t)
				return c, nil
			}
		}
		l.Debug("No peer found in existing connection")
		return nil, errors.New("no peer found")
	}

	l.Debug("Unsupported UI type")
	return nil, errors.New("unsupported UI type")
}

type ConnBusinessMgmt struct {
	PeerName string
}

func (z *ConnBusinessMgmt) Connect(control app_control.Control) (ctx api_context.Context, err error) {
	return connect(api_auth.DropboxTokenBusinessManagement, z.PeerName, control)
}

func (*ConnBusinessMgmt) IsBusinessMgmt() {
}

type ConnBusinessInfo struct {
	PeerName string
}

func (z *ConnBusinessInfo) IsBusinessInfo() {
}

func (z *ConnBusinessInfo) Connect(control app_control.Control) (ctx api_context.Context, err error) {
	return connect(api_auth.DropboxTokenBusinessInfo, z.PeerName, control)
}

type ConnBusinessFile struct {
	PeerName string
}

func (z *ConnBusinessFile) IsBusinessFile() {
}

func (z *ConnBusinessFile) Connect(control app_control.Control) (ctx api_context.Context, err error) {
	return connect(api_auth.DropboxTokenBusinessFile, z.PeerName, control)
}

type ConnBusinessAudit struct {
	PeerName string
}

func (z *ConnBusinessAudit) IsBusinessAudit() {
}

func (z *ConnBusinessAudit) Connect(control app_control.Control) (ctx api_context.Context, err error) {
	return connect(api_auth.DropboxTokenBusinessAudit, z.PeerName, control)
}

type ConnUserFile struct {
	PeerName string
}

func (z *ConnUserFile) IsUserFile() {
}

func (z *ConnUserFile) Connect(control app_control.Control) (ctx api_context.Context, err error) {
	return connect(api_auth.DropboxTokenFull, z.PeerName, control)
}
