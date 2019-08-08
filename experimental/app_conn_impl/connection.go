package app_conn_impl

import (
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/experimental/app_kitchen"
)

const (
	DefaultPeerName = "default"
)

func NewConnBusinessMgmt() *ConnBusinessMgmt {
	return &ConnBusinessMgmt{
		PeerName: DefaultPeerName,
	}
}

func NewConnBusinessInfo() *ConnBusinessInfo {
	return &ConnBusinessInfo{
		PeerName: DefaultPeerName,
	}
}

func connect(tokenType, peerName string, kitchen app_kitchen.Kitchen) (ctx api_context.Context, err error) {
	c := api_auth_impl.NewKc(kitchen, api_auth_impl.PeerName(peerName))
	ctx, err = c.Auth(tokenType)
	return
}

type ConnBusinessMgmt struct {
	PeerName string
}

func (z *ConnBusinessMgmt) Connect(kitchen app_kitchen.Kitchen) (ctx api_context.Context, err error) {
	return connect(api_auth.DropboxTokenBusinessManagement, z.PeerName, kitchen)
}

func (*ConnBusinessMgmt) IsBusinessMgmt() {
}

type ConnBusinessInfo struct {
	PeerName string
}

func (z *ConnBusinessInfo) IsBusinessInfo() {
}

func (z *ConnBusinessInfo) Connect(kitchen app_kitchen.Kitchen) (ctx api_context.Context, err error) {
	return connect(api_auth.DropboxTokenBusinessInfo, z.PeerName, kitchen)
}

type ConnBusinessFile struct {
	PeerName string
}

func (z *ConnBusinessFile) IsBusinessFile() {
}

func (z *ConnBusinessFile) Connect(kitchen app_kitchen.Kitchen) (ctx api_context.Context, err error) {
	return connect(api_auth.DropboxTokenBusinessFile, z.PeerName, kitchen)
}

type ConnBusinessAudit struct {
	PeerName string
}

func (z *ConnBusinessAudit) IsBusinessAudit() {
}

func (z *ConnBusinessAudit) Connect(kitchen app_kitchen.Kitchen) (ctx api_context.Context, err error) {
	return connect(api_auth.DropboxTokenBusinessAudit, z.PeerName, kitchen)
}

type ConnUserFile struct {
	PeerName string
}

func (z *ConnUserFile) IsUserFile() {
}

func (z *ConnUserFile) Connect(kitchen app_kitchen.Kitchen) (ctx api_context.Context, err error) {
	return connect(api_auth.DropboxTokenFull, z.PeerName, kitchen)
}
