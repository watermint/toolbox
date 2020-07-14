package stage

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

// Staging recipe for Dropbox scoped OAuth
type Scoped struct {
	Peer dbx_conn.ConnScopedIndividual
	List rp_model.RowReport
}

func (z *Scoped) Preset() {
	z.Peer.SetScopes(dbx_auth.ScopeFilesContentRead)
	z.List.SetModel(&mo_file.ConcreteEntry{})
}

func (z *Scoped) Exec(c app_control.Control) error {
	if err := z.List.Open(); err != nil {
		return err
	}
	entries, err := sv_file.NewFiles(z.Peer.Context()).List(mo_path.NewDropboxPath("/"))
	if err != nil {
		return err
	}

	for _, entry := range entries {
		z.List.Row(entry)
	}
	return nil
}

func (z *Scoped) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Scoped{}, rc_recipe.NoCustomValues)
}
