package pipe_sample

import (
	"github.com/watermint/toolbox/poc/oper/oper_auth"
	"github.com/watermint/toolbox/poc/pipe/pipe_core"
)

type Invite struct {
	Req *InviteRequest
}

func (Invite) Exec(s pipe_core.Session) (res pipe_core.Response, err error) {
	s.Log().Info("start")
	s.Message("invite_start").Tell()

	panic("")
}

type InviteRequest struct {
	Api   *oper_auth.DropboxBusinessManagement
	Email string
}
type InviteResponse struct {
}
