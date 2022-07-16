package module

import (
	"github.com/watermint/toolbox/essentials/go/go_module"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type List struct {
	rc_recipe.RemarkSecret
}

func (z *List) Preset() {
}

func (z *List) Exec(c app_control.Control) error {
	b, err := go_module.ScanBuild()
	if err != nil {
		return err
	}
	l := c.Log()
	l.Info("Go Version", esl.String("version", b.GoVersion()))
	for _, m := range b.Modules() {
		lic, err := go_module.SelectLicenses(m,
			[]go_module.LicenseType{
				go_module.LicenseTypeApache20,
				go_module.LicenseTypeBSD2Clause,
				go_module.LicenseTypeBSD3Clause,
				go_module.LicenseTypeMIT,
			},
			[]go_module.LicenseType{
				go_module.LicenseTypeAGPL,
				go_module.LicenseTypeGPL,
				go_module.LicenseTypeLGPL,
			},
		)
		if err != nil {
			l.Error("No license", esl.String("Path", m.Path()), esl.String("Version", m.Version()))
			continue
		}
		l.Info("Module", esl.String("Path", m.Path()), esl.String("Version", m.Version()), esl.String("License", string(lic.Type())))
	}

	return nil
}

func (z *List) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}
