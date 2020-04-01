package teamfolder

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
)

type List struct {
	Peer       rc_conn.ConnBusinessFile
	TeamFolder rp_model.RowReport
}

func (z *List) Preset() {
	z.TeamFolder.SetModel(
		&mo_teamfolder.TeamFolder{},
		rp_model.HiddenColumns(
			"team_folder_id",
		),
	)
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "team_folder", func(cols map[string]string) error {
		if _, ok := cols["name"]; !ok {
			return errors.New("`name` is not found")
		}
		return nil
	})
}

func (z *List) Exec(c app_control.Control) error {
	folders, err := sv_teamfolder.New(z.Peer.Context()).List()
	if err != nil {
		// ApiError will be reported by infra
		return err
	}

	if err := z.TeamFolder.Open(); err != nil {
		return err
	}
	for _, folder := range folders {
		c.Log().Debug("Folder", zap.Any("folder", folder))
		z.TeamFolder.Row(folder)
	}

	return nil
}
