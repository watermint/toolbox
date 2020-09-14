package teamfolder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Add struct {
	rc_recipe.RemarkIrreversible
	Peer                dbx_conn.ConnBusinessFile
	Name                string
	SyncSetting         mo_string.SelectString
	Added               rp_model.RowReport
	ErrorUnableToCreate app_msg.Message
}

func (z *Add) Preset() {
	z.Added.SetModel(&mo_teamfolder.TeamFolder{},
		rp_model.HiddenColumns(
			"team_folder_id",
		),
	)
	z.SyncSetting.SetOptions(
		"default",
		"default", "not_synced",
	)
}

func (z *Add) Exec(c app_control.Control) error {
	ui := c.UI()
	if err := z.Added.Open(); err != nil {
		return err
	}

	opts := make([]sv_teamfolder.CreateOption, 0)
	switch z.SyncSetting.Value() {
	case "default":
		opts = append(opts, sv_teamfolder.SyncDefault())
	case "not_synced":
		opts = append(opts, sv_teamfolder.SyncNoSync())
	}

	folder, err := sv_teamfolder.New(z.Peer.Context()).Create(z.Name, opts...)
	if err != nil {
		ui.Error(z.ErrorUnableToCreate.With("Name", z.Name).With("Error", err))
		return err
	}
	z.Added.Row(folder)
	return nil
}

func (z *Add) Test(c app_control.Control) error {
	err := rc_exec.ExecReplay(c, &Add{}, "replay-teamfolder-add.json.gz", func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.Name = "xxxxx"
	})
	if err != nil {
		return err
	}

	return rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.Name = "Sales"
	})
}
