package stage

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

// Staging recipe for Dropbox scoped OAuth
type Scoped struct {
	rc_recipe.RemarkSecret
	Individual dbx_conn.ConnScopedIndividual
	Team       dbx_conn.ConnScopedTeam
	FileList   rp_model.RowReport
	MemberList rp_model.RowReport
}

func (z *Scoped) Preset() {
	z.Individual.SetScopes(dbx_auth.ScopeFilesContentRead)
	z.Team.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeTeamInfoRead,
	)
	z.FileList.SetModel(&mo_file.ConcreteEntry{})
	z.MemberList.SetModel(&mo_member.Member{})
}

func (z *Scoped) Exec(c app_control.Control) error {
	if err := z.FileList.Open(); err != nil {
		return err
	}
	entries, err := sv_file.NewFiles(z.Individual.Client()).List(mo_path.NewDropboxPath("/"))
	if err != nil {
		return err
	}
	for _, entry := range entries {
		z.FileList.Row(entry)
	}

	if err := z.MemberList.Open(); err != nil {
		return err
	}
	members, err := sv_member.New(z.Team.Client()).List()
	if err != nil {
		return err
	}
	for _, member := range members {
		z.MemberList.Row(member)
	}

	admin, err := sv_profile.NewTeam(z.Team.Client()).Admin()
	if err != nil {
		return err
	}
	c.Log().Info("Admin", esl.Any("admin", admin))

	return nil
}

func (z *Scoped) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Scoped{}, rc_recipe.NoCustomValues)
}
