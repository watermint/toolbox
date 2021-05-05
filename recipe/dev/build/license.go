package build

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/model/mo_path"
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
	DestPath mo_path.FileSystemPath
}

func (z *License) Preset() {
}

func (z *License) Exec(c app_control.Control) error {
	l := c.Log()
	l.Info("Please ignore logs starts from ERROR:")
	inventory, err := dc_license.Detect(c)
	if err != nil {
		return err
	}

	prjBase, err := es_project.DetectRepositoryRoot()
	if err != nil {
		return err
	}
	licenseBody, err := ioutil.ReadFile(filepath.Join(prjBase, "LICENSE.md"))
	if err != nil {
		return err
	}

	licenses := dc_license.Licenses{
		Project:    dc_license.LicenseInfo{
			Package:     "github.com/watermint/toolbox",
			Url:         "https://github.com/watermint/toolbox",
			LicenseType: "MIT",
			LicenseBody: string(licenseBody),
		},
		ThirdParty: inventory,
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
		m.DestPath = mo_path.NewFileSystemPath(filepath.Join(path, "licenses.json"))
	})
}
