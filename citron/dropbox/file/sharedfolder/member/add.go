package member

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_sharedfolder"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Add struct {
	Peer        dbx_conn.ConnScopedIndividual
	Path        mo_path.DropboxPath
	AccessLevel mo_string.SelectString
	Email       string
	Silent      bool
	Message     mo_string.OptionalString
}

func (z *Add) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeSharingWrite,
	)
	z.AccessLevel.SetOptions(
		"editor",
		"editor", "viewer", "viewer_no_comment",
	)
}

func (z *Add) Exec(c app_control.Control) error {
	sfr := uc_sharedfolder.NewResolver(z.Peer.Client())

	sf, err := sfr.Resolve(z.Path)
	if err != nil {
		return err
	}

	opts := make([]sv_sharedfolder_member.AddOption, 0)
	if z.Silent {
		opts = append(opts, sv_sharedfolder_member.AddQuiet())
	}
	if z.Message.IsExists() {
		opts = append(opts, sv_sharedfolder_member.AddCustomMessage(z.Message.Value()))
	}
	err = sv_sharedfolder_member.New(z.Peer.Client(), sf).Add(sv_sharedfolder_member.AddByEmail(z.Email, z.AccessLevel.Value()), opts...)
	if err != nil {
		return err
	}
	return nil
}

func (z *Add) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.Email = "emma@example.com"
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("add")
	})
}
