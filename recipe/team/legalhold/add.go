package legalhold

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_legalhold"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_legalhold"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type MemberEmail struct {
	Email string `json:"email"`
}

type Add struct {
	Peer           dbx_conn.ConnScopedTeam
	Name           string
	Member         fd_file.RowFeed
	Policy         rp_model.RowReport
	Description    mo_string.OptionalString
	StartDate      mo_time.TimeOptional
	EndDate        mo_time.TimeOptional
	MemberNotFound app_msg.Message
}

func (z *Add) Preset() {
	//z.Peer.SetScopes(
	//dbx_auth.ScopeMembersRead,
	//dbx_auth.ScopeTeamDataGovernanceWrite,
	//)
	z.Member.SetModel(&MemberEmail{})
	z.Policy.SetModel(&mo_legalhold.Policy{})
}

func (z *Add) Exec(c app_control.Control) error {
	if err := z.Policy.Open(); err != nil {
		return err
	}
	svm := sv_member.NewCached(z.Peer.Client())

	members := make([]*mo_member.Member, 0)
	err := z.Member.EachRow(func(r interface{}, rowIndex int) error {
		m := r.(*MemberEmail)
		if m.Email == "" {
			return nil
		}
		member, err := svm.ResolveByEmail(m.Email)
		if err != nil {
			c.UI().Error(z.MemberNotFound.With("Email", m.Email))
			return err
		}
		members = append(members, member)
		return nil
	})
	if err != nil {
		return err
	}

	policy, err := sv_legalhold.New(z.Peer.Client()).Create(
		z.Name,
		z.Description.Value(),
		z.StartDate.Time(),
		z.EndDate.Time(),
		members,
	)
	if err != nil {
		return err
	}

	z.Policy.Row(policy)
	return nil
}

func (z *Add) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("add", "alex@example.com\nemma@example.com\n")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.Name = "test"
		m.Member.SetFilePath(f)
	})
}
