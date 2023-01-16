package conversation

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/slack/api/work_auth"
	"github.com/watermint/toolbox/domain/slack/api/work_conn"
	"github.com/watermint/toolbox/domain/slack/model/mo_message"
	"github.com/watermint/toolbox/domain/slack/service/sv_conversation_history"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"strconv"
	"time"
)

type History struct {
	rc_recipe.RemarkSecret
	Peer     work_conn.ConnSlackApi
	Messages rp_model.RowReport
	Channel  string
	After    mo_time.TimeOptional
}

func (z *History) Preset() {
	z.Peer.SetScopes(work_auth.ScopeChannelsHistory)
	z.Messages.SetModel(&mo_message.Message{})
}

func (z *History) Exec(c app_control.Control) error {
	if err := z.Messages.Open(); err != nil {
		return err
	}
	after := strconv.FormatInt(time.Now().Add(-24*time.Hour).Unix(), 10)
	if !z.After.IsZero() {
		after = strconv.FormatInt(z.After.Time().Unix(), 10)
	}
	inclusive := true

	sv := sv_conversation_history.NewHistory(z.Peer.Client())
	return sv.ListEach(z.Channel, func(m mo_message.Message) {
		z.Messages.Row(&m)
	}, sv_conversation_history.HistoryListOpts{
		After:              after,
		Inclusive:          &inclusive,
		IncludeAllMetadata: &inclusive,
	})
}

func (z *History) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &History{}, func(r rc_recipe.Recipe) {
		m := r.(*History)
		m.Channel = "C1234567890"
	})
}
