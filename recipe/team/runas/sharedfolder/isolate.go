package sharedfolder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Isolate struct {
	rc_recipe.RemarkIrreversible
	Peer        dbx_conn.ConnScopedTeam
	MemberEmail string
	Isolated    rp_model.TransactionReport
	KeepCopy    bool
}

func (z *Isolate) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeSharingWrite,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
	)
	z.Isolated.SetModel(
		&mo_sharedfolder.SharedFolder{},
		nil,
		rp_model.HiddenColumns(
			"parent_shared_folder_id",
			"team_member_id",
			"namespace_id",
			"owner_team_id",
		),
	)
}

func (z *Isolate) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.Isolated.Open(); err != nil {
		return err
	}

	teamInfo, err := sv_team.New(z.Peer.Context()).Info()
	if err != nil {
		return err
	}

	member, err := sv_member.New(z.Peer.Context()).ResolveByEmail(z.MemberEmail)
	if err != nil {
		return err
	}

	l.Debug("Member found", esl.Any("member", member))

	ctx := z.Peer.Context().AsMemberId(member.TeamMemberId)
	folders, err := sv_sharedfolder.New(ctx).List()
	if err != nil {
		return err
	}

	var lastErr error
	for _, folder := range folders {
		if folder.AccessType == "owner" {
			// unshare
			err = sv_sharedfolder.New(ctx).Remove(folder, sv_sharedfolder.LeaveACopy(z.KeepCopy))
			if err != nil {
				z.Isolated.Failure(err, folder)
				lastErr = err
			} else {
				z.Isolated.Success(folder, nil)
			}
		} else if folder.OwnerTeamId != teamInfo.TeamId {
			// leave
			err = sv_sharedfolder.New(ctx).Leave(folder, sv_sharedfolder.LeaveACopy(z.KeepCopy))
			if err != nil {
				z.Isolated.Failure(err, folder)
				lastErr = err
			} else {
				z.Isolated.Success(folder, nil)
			}
		}
	}
	return lastErr
}

func (z *Isolate) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Isolate{}, func(r rc_recipe.Recipe) {
		m := r.(*Isolate)
		m.MemberEmail = "john@example.com"
	})
}
