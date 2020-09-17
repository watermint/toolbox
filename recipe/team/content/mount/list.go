package mount

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_mount"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer   dbx_conn.ConnBusinessFile
	Mount  rp_model.RowReport
	Member mo_filter.Filter
}

func (z *List) Preset() {
	z.Mount.SetModel(&mo_sharedfolder.MemberMount{}, rp_model.HiddenColumns(
		"team_member_id",
		"owner_team_id",
	))
	z.Member.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_filter.NewNameSuffixFilter(),
		mo_filter.NewEmailFilter(),
	)
}

func (z *List) scanMember(member *mo_member.Member, c app_control.Control) error {
	l := c.Log().With(esl.Any("member", member))
	ctx := z.Peer.Context().AsMemberId(member.TeamMemberId)
	l.Debug("Scan member")

	mounts, err := sv_sharedfolder_mount.New(ctx).List()
	if err != nil {
		l.Debug("Unable to scan member mounts", esl.Error(err))
		return err
	}

	for _, mount := range mounts {
		z.Mount.Row(mo_sharedfolder.NewMemberMount(member, mount))
	}
	l.Debug("scan finished")
	return nil
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Mount.Open(); err != nil {
		return err
	}

	members, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	var lastErr error
	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("scan_member", z.scanMember, c)
		q := s.Get("scan_member")

		for _, member := range members {
			q.Enqueue(member)
		}
	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		if err != nil {
			lastErr = err
		}
	}))
	return lastErr
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
