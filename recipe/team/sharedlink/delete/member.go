package delete

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_team_sharedlink"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Member struct {
	rc_recipe.RemarkIrreversible
	Peer         dbx_conn.ConnScopedTeam
	MemberEmail  string
	OperationLog rp_model.TransactionReport
}

func (z *Member) Preset() {
	z.OperationLog.SetModel(
		&TargetLinks{},
		&mo_sharedlink.SharedLinkMember{},
		rp_model.HiddenColumns(
			"result.shared_link_id",
			"result.account_id",
			"result.team_member_id",
			"result.status",
		),
	)
}

func (z *Member) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	member, err := sv_member.New(z.Peer.Context()).ResolveByEmail(z.MemberEmail)
	if err != nil {
		return err
	}

	var onDeleteSuccess uc_team_sharedlink.DeleteOnSuccess = func(t *uc_team_sharedlink.DeleteTarget) {
		z.OperationLog.Success(&TargetLinks{Url: t.Entry.Url}, t.Entry)
	}
	var onDeleteFailure uc_team_sharedlink.DeleteOnFailure = func(t *uc_team_sharedlink.DeleteTarget, cause error) {
		z.OperationLog.Failure(cause, &TargetLinks{Url: t.Entry.Url})
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("delete_link", uc_team_sharedlink.DeleteMemberLink, c, z.Peer.Context(), onDeleteSuccess, onDeleteFailure)
		var onSharedLink uc_team_sharedlink.OnSharedLinkMember = func(member *mo_member.Member, entry *mo_sharedlink.SharedLinkMember) {
			q := s.Get("delete_link")
			q.Enqueue(&uc_team_sharedlink.DeleteTarget{
				Member: member,
				Entry:  entry,
			})
		}
		s.Define("scan_member", uc_team_sharedlink.RetrieveMemberLinks, c, z.Peer.Context(), onSharedLink)
		q := s.Get("scan_member")
		q.Enqueue(member)
	})

	return nil
}

func (z *Member) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Member{}, func(r rc_recipe.Recipe) {
		m := r.(*Member)
		m.MemberEmail = "test@example.com"
	})
}
