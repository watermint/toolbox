package file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_url"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink_file"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer     dbx_conn.ConnUserFile
	Url      mo_url.Url
	Password mo_string.OptionalString
	FileList rp_model.RowReport
}

func (z *List) Preset() {
	z.FileList.SetModel(
		&mo_file.ConcreteEntry{},
		rp_model.HiddenColumns(
			"id",
			"path_lower",
			"revision",
			"content_hash",
			"shared_folder_id",
			"parent_shared_folder_id",
		),
		rp_model.HiddenColumns(
			"id",
			"path_lower",
			"revision",
			"content_hash",
			"shared_folder_id",
			"parent_shared_folder_id",
		),
	)
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.FileList.Open(); err != nil {
		return err
	}

	return sv_sharedlink_file.New(z.Peer.Context()).ListRecursive(z.Url, func(entry mo_file.Entry) {
		z.FileList.Row(entry.Concrete())
	}, sv_sharedlink_file.Password(z.Password.Value()))
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, func(r rc_recipe.Recipe) {
		m := r.(*List)
		m.Url, _ = mo_url.NewUrl("https://www.dropbox.com/s/2sn712vy1ovegw8/Prime_Numbers.txt?dl=0")
	})
}
