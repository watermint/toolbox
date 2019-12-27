package teamfolder

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
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
	z.TeamFolder.SetModel(&mo_teamfolder.TeamFolder{})
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "team_folder", func(cols map[string]string) error {
		if _, ok := cols["team_folder_id"]; !ok {
			return errors.New("`team_folder_id` is not found")
		}
		return nil
	})
}

func (z *List) Exec(k rc_kitchen.Kitchen) error {
	folders, err := sv_teamfolder.New(z.Peer.Context()).List()
	if err != nil {
		// ApiError will be reported by infra
		return err
	}

	if err := z.TeamFolder.Open(); err != nil {
		return err
	}
	for _, folder := range folders {
		k.Log().Debug("Folder", zap.Any("folder", folder))
		z.TeamFolder.Row(folder)
	}

	return nil
}
