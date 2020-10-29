package image

import (
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/essentials/image/ei_exif"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_image"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"path/filepath"
)

type Info struct {
	rc_recipe.RemarkSecret
	Path mo_path.ExistingFileSystemPath
	Exif rp_model.RowReport
}

func (z *Info) Preset() {
	z.Exif.SetModel(&mo_image.Exif{})
}

func (z *Info) Exec(c app_control.Control) error {
	l := c.Log().With(esl.String("path", z.Path.Path()))

	parser := ei_exif.Auto(l)
	exifData, err := parser.Parse(z.Path.Path())
	if err != nil {
		l.Debug("Unable to parse", esl.Error(err))
		return err
	}

	if err = z.Exif.Open(); err != nil {
		return err
	}
	z.Exif.Row(&exifData)
	return nil
}

func (z *Info) Test(c app_control.Control) error {
	rr, err := es_project.DetectRepositoryRoot()
	if err != nil {
		return err
	}
	return rc_exec.Exec(c, &Info{}, func(r rc_recipe.Recipe) {
		m := r.(*Info)
		m.Path = mo_path.NewExistingFileSystemPath(filepath.Join(rr, "test/data/exif_test001.jpg"))
	})
}
