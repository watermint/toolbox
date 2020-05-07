package device

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_device"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_device"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type List struct {
	Peer   dbx_conn.ConnBusinessFile
	Device rp_model.RowReport
}

func (z *List) Preset() {
	z.Device.SetModel(
		&mo_device.MemberSession{},
		rp_model.HiddenColumns(
			"familiar_name",
			"abbreviated_name",
			"member_folder_id",
			"external_id",
			"account_id",
			"persistent_id",
		),
	)
}

func (z *List) Exec(c app_control.Control) error {
	memberList, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}
	members := mo_member.MapByTeamMemberId(memberList)

	if err := z.Device.Open(); err != nil {
		return err
	}

	sessions, err := sv_device.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	for _, session := range sessions {
		if m, e := members[session.EntryTeamMemberId()]; e {
			ma := mo_device.NewMemberSession(m, session)
			z.Device.Row(ma)
		}
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qtr_endtoend.TestRows(c, "device", func(cols map[string]string) error {
		if _, ok := cols["team_member_id"]; !ok {
			return errors.New("team_member_id is not found")
		}
		return nil
	})
}
