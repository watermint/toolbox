package delete

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_filerequest"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_filerequest"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Closed struct {
	rc_recipe.RemarkIrreversible
	Peer     dbx_conn.ConnScopedIndividual
	Deleted  rp_model.RowReport
	BasePath mo_string.SelectString
}

func (z *Closed) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFileRequestsWrite,
	)
	z.Deleted.SetModel(&mo_filerequest.FileRequest{})
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Closed) Exec(c app_control.Control) error {
	if err := z.Deleted.Open(); err != nil {
		return err
	}
	client := z.Peer.Client().BaseNamespace(dbx_filesystem.AsNamespaceType(z.BasePath.Value()))
	frs, err := sv_filerequest.New(client).DeleteAllClosed()
	if err != nil {
		return err
	}
	for _, fr := range frs {
		z.Deleted.Row(fr)
	}
	return nil
}

func (z *Closed) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Closed{}, rc_recipe.NoCustomValues)
}
