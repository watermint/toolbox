package delete

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_team_sharedlink"
	"github.com/watermint/toolbox/essentials/model/mo_string"
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
	BasePath     mo_string.SelectString
}

func (z *Member) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeSharingWrite,
		dbx_auth.ScopeTeamDataMember,
	)
	z.OperationLog.SetModel(
		&uc_team_sharedlink.TargetLinks{},
		&mo_sharedlink.SharedLinkMember{},
		rp_model.HiddenColumns(
			"result.shared_link_id",
			"result.account_id",
			"result.team_member_id",
			"result.status",
		),
	)
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Member) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	member, err := sv_member.New(z.Peer.Client()).ResolveByEmail(z.MemberEmail)
	if err != nil {
		return err
	}

	var onDeleteSuccess uc_team_sharedlink.DeleteOnSuccess = func(t *uc_team_sharedlink.Target) {
		z.OperationLog.Success(&uc_team_sharedlink.TargetLinks{Url: t.Entry.Url}, t.Entry)
	}
	var onDeleteFailure uc_team_sharedlink.DeleteOnFailure = func(t *uc_team_sharedlink.Target, cause error) {
		z.OperationLog.Failure(cause, &uc_team_sharedlink.TargetLinks{Url: t.Entry.Url})
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("delete_link", uc_team_sharedlink.DeleteMemberLink, c, z.Peer.Client(), onDeleteSuccess, onDeleteFailure, dbx_filesystem.AsNamespaceType(z.BasePath.Value()))
		var onSharedLink uc_team_sharedlink.OnSharedLinkMember = func(member *mo_member.Member, entry *mo_sharedlink.SharedLinkMember) {
			q := s.Get("delete_link")
			q.Enqueue(&uc_team_sharedlink.Target{
				Member: member,
				Entry:  entry,
			})
		}
		s.Define("scan_member", uc_team_sharedlink.RetrieveMemberLinks, c, z.Peer.Client(), onSharedLink, dbx_filesystem.AsNamespaceType(z.BasePath.Value()))
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
