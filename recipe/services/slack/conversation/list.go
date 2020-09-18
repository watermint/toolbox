package conversation

import (
	"github.com/watermint/toolbox/domain/slack/api/work_auth"
	"github.com/watermint/toolbox/domain/slack/api/work_conn"
	"github.com/watermint/toolbox/domain/slack/model/mo_conversation"
	"github.com/watermint/toolbox/domain/slack/service/sv_conversation"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer         work_conn.ConnSlackApi
	Conversation rp_model.RowReport
}

func (z *List) Preset() {
	z.Peer.SetScopes(work_auth.ScopeChannelsRead)
	z.Conversation.SetModel(&mo_conversation.Conversation{})
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Conversation.Open(); err != nil {
		return err
	}

	return sv_conversation.New(z.Peer.Context()).ListEach(func(c *mo_conversation.Conversation) {
		z.Conversation.Row(c)
	})
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
