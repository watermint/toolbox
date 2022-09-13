package member

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_sharedfolder"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Delete struct {
	Peer      dbx_conn.ConnScopedIndividual
	Path      mo_path.DropboxPath
	Email     string
	LeaveCopy bool
}

func (z *Delete) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeSharingWrite,
	)
}

func (z *Delete) Exec(c app_control.Control) error {
	sfr := uc_sharedfolder.NewResolver(z.Peer.Client())

	sf, err := sfr.Resolve(z.Path)
	if err != nil {
		return err
	}

	opts := make([]sv_sharedfolder_member.RemoveOption, 0)
	if z.LeaveCopy {
		opts = append(opts, sv_sharedfolder_member.LeaveACopy())
	}
	err = sv_sharedfolder_member.New(z.Peer.Client(), sf).Remove(sv_sharedfolder_member.RemoveByEmail(z.Email), opts...)
	if err != nil {
		return err
	}
	return nil
}

func (z *Delete) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.Email = "emma@example.com"
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("delete")
	})
}
