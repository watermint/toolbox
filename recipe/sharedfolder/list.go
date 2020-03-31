package sharedfolder

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type List struct {
	Peer         rc_conn.ConnUserFile
	SharedFolder rp_model.RowReport
}

func (z *List) Preset() {
	z.SharedFolder.SetModel(&mo_sharedfolder.SharedFolder{})
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "shared_folder", func(cols map[string]string) error {
		if _, ok := cols["shared_folder_id"]; !ok {
			return errors.New("shared_folder_id is not found")
		}
		return nil
	})
}

func (z *List) Exec(c app_control.Control) error {
	c.Log().Debug("Scanning folders")
	folders, err := sv_sharedfolder.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	if err := z.SharedFolder.Open(); err != nil {
		return err
	}

	for _, folder := range folders {
		z.SharedFolder.Row(folder)
	}
	return nil
}
