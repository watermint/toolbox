package member

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type List struct {
	Peer   dbx_conn.ConnScopedIndividual
	Member rp_model.RowReport
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeSharingRead,
	)
	z.Member.SetModel(
		&mo_sharedfolder_member.SharedFolderMember{},
		rp_model.HiddenColumns(
			"shared_folder_id",
			"parent_shared_folder_id",
			"account_id",
			"group_id",
		),
	)
}

func (z *List) listMembers(sf *mo_sharedfolder.SharedFolder, c app_control.Control) error {
	l := c.Log().With(esl.Any("sf", sf))
	l.Debug("Scanning")
	members, err := sv_sharedfolder_member.New(z.Peer.Client(), sf).List()
	if err != nil {
		return err
	}

	for _, member := range members {
		z.Member.Row(mo_sharedfolder_member.NewSharedFolderMember(sf, member))
	}
	return nil
}

func (z *List) Exec(c app_control.Control) error {
	folders, err := sv_sharedfolder.New(z.Peer.Client()).List()
	if err != nil {
		return err
	}

	if err := z.Member.Open(); err != nil {
		return err
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("scan_folder", z.listMembers, c)
		q := s.Get("scan_folder")

		for _, folder := range folders {
			q.Enqueue(folder)
		}
	})

	return nil
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qtr_endtoend.TestRows(c, "member", func(cols map[string]string) error {
		if _, ok := cols["email"]; !ok {
			return errors.New("email is not found")
		}
		return nil
	})
}
