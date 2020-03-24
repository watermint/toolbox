package _import

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_url"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type Url struct {
	Peer         rc_conn.ConnUserFile
	Path         mo_path.DropboxPath
	Url          string
	OperationLog rp_model.RowReport
}

func (z *Url) Preset() {
	z.OperationLog.SetModel(&mo_file.ConcreteEntry{})
}

func (z *Url) Exec(c app_control.Control) error {
	ui := c.UI()
	ctx := z.Peer.Context()

	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	path := sv_file_url.PathWithName(z.Path, z.Url)
	ui.InfoK("recipe.file.import.url.progress", app_msg.P{
		"Path": path.Path(),
		"Url":  z.Url,
	})
	entry, err := sv_file_url.New(ctx).Save(path, z.Url)
	if err != nil {
		return err
	}
	z.OperationLog.Row(entry.Concrete())
	return nil
}

func (z *Url) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Url{}, func(r rc_recipe.Recipe) {
		ru := r.(*Url)
		ru.Url = "https://dummyimage.com/10x10/000/fff"
		ru.Path = qt_recipe.NewTestDropboxFolderPath("file-import-url")
	})
}
