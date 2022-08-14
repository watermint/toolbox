package image

import (
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/essentials/image/ei_exif"
	"github.com/watermint/toolbox/essentials/model/mo_image"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"path/filepath"
)

type Exif struct {
	rc_recipe.RemarkTransient
	File     mo_path.ExistingFileSystemPath
	Metadata rp_model.RowReport
}

func (z *Exif) Preset() {
	z.Metadata.SetModel(
		&mo_image.Exif{},
	)
}

func (z *Exif) Exec(c app_control.Control) error {
	if err := z.Metadata.Open(); err != nil {
		return err
	}
	ei := ei_exif.Auto(c.Log())
	m, err := ei.Parse(z.File.Path())
	if err != nil {
		return err
	}
	z.Metadata.Row(m)
	return nil
}

func (z *Exif) Test(c app_control.Control) error {
	wsr, err := es_project.DetectRepositoryRoot()
	if err != nil {
		return err
	}
	return rc_exec.Exec(c, &Exif{}, func(r rc_recipe.Recipe) {
		m := r.(*Exif)
		m.File = mo_path.NewExistingFileSystemPath(filepath.Join(wsr, "recipe/util/image", "exif_test.jpg"))
	})
}
