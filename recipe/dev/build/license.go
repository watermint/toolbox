package build

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/doc/dc_license"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"io/ioutil"
	"os"
	"path/filepath"
)

type License struct {
	rc_recipe.RemarkSecret
	SourcePath mo_path.ExistingFileSystemPath
	DestPath   mo_path.FileSystemPath
}

func (z *License) Preset() {
}

func (z *License) Exec(c app_control.Control) error {
	licenses, err := dc_license.LoadLicenses(c, z.SourcePath.Path())
	if err != nil {
		return err
	}
	licensesData, err := json.Marshal(licenses)
	if err != nil {
		return err
	}

	if c.Feature().IsTest() {
		out := es_stdout.NewDefaultOut(c.Feature())
		_, err := out.Write(licensesData)
		return err
	} else {
		return ioutil.WriteFile(z.DestPath.Path(), licensesData, 0644)
	}
}

func (z *License) Test(c app_control.Control) error {
	path, err := qt_file.MakeTestFolder("license", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(path)
	}()

	if err := dc_license.MakeTestData(path); err != nil {
		return err
	}

	return rc_exec.Exec(c, &License{}, func(r rc_recipe.Recipe) {
		m := r.(*License)
		m.SourcePath = mo_path.NewExistingFileSystemPath(path)
		m.DestPath = mo_path.NewFileSystemPath(filepath.Join(path, "licenses.json"))
	})
}
