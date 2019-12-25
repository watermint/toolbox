package rc_conn_impl

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
)

// Deprecated:
func NewOldConnBusinessMgmt() rc_conn.OldConnBusinessMgmt {
	return &ConnBusinessMgmt{
		PeerName: DefaultPeerName,
	}
}

// Deprecated:
func NewOldConnBusinessInfo() rc_conn.OldConnBusinessInfo {
	return &ConnBusinessInfo{
		PeerName: DefaultPeerName,
	}
}

// Deprecated:
func NewOldConnBusinessAudit() rc_conn.OldConnBusinessAudit {
	return &ConnBusinessAudit{
		PeerName: DefaultPeerName,
	}
}

// Deprecated:
func NewOldConnBusinessFile() rc_conn.OldConnBusinessFile {
	return &ConnBusinessFile{
		PeerName: DefaultPeerName,
	}
}

// Deprecated:
func NewOldConnUserFile() *ConnUserFile {
	return &ConnUserFile{
		PeerName: DefaultPeerName,
	}
}

// Deprecated:
type ConnBusinessMgmt struct {
	PeerName string
}

func (z *ConnBusinessMgmt) Name() string {
	return z.PeerName
}

func (z *ConnBusinessMgmt) Connect(control app_control.Control) (ctx api_context.Context, err error) {
	return connect(api_auth.DropboxTokenBusinessManagement, z.PeerName, control)
}

func (*ConnBusinessMgmt) IsBusinessMgmt() {
}

// Deprecated:
type ConnBusinessInfo struct {
	PeerName string
}

func (z *ConnBusinessInfo) Name() string {
	return z.PeerName
}

func (z *ConnBusinessInfo) IsBusinessInfo() {
}

func (z *ConnBusinessInfo) Connect(control app_control.Control) (ctx api_context.Context, err error) {
	return connect(api_auth.DropboxTokenBusinessInfo, z.PeerName, control)
}

// Deprecated:
type ConnBusinessFile struct {
	PeerName string
}

func (z *ConnBusinessFile) Name() string {
	return z.PeerName
}

func (z *ConnBusinessFile) IsBusinessFile() {
}

func (z *ConnBusinessFile) Connect(control app_control.Control) (ctx api_context.Context, err error) {
	return connect(api_auth.DropboxTokenBusinessFile, z.PeerName, control)
}

// Deprecated:
type ConnBusinessAudit struct {
	PeerName string
}

func (z *ConnBusinessAudit) Name() string {
	return z.PeerName
}

func (z *ConnBusinessAudit) IsBusinessAudit() {
}

func (z *ConnBusinessAudit) Connect(control app_control.Control) (ctx api_context.Context, err error) {
	return connect(api_auth.DropboxTokenBusinessAudit, z.PeerName, control)
}

// Deprecated:
type ConnUserFile struct {
	PeerName string
}

func (z *ConnUserFile) Name() string {
	return z.Name()
}

func (z *ConnUserFile) IsUserFile() {
}

func (z *ConnUserFile) Connect(control app_control.Control) (ctx api_context.Context, err error) {
	return connect(api_auth.DropboxTokenFull, z.PeerName, control)
}
