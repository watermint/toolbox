package module

import (
	"github.com/watermint/toolbox/essentials/go/es_module"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type List struct {
	rc_recipe.RemarkSecret
}

func (z *List) Preset() {
}

func (z *List) Exec(c app_control.Control) error {
	b, err := es_module.ScanBuild()
	if err != nil {
		return err
	}
	l := c.Log()
	l.Info("Go Version", esl.String("version", b.GoVersion()))
	for _, m := range b.Modules() {
		lic, err := es_module.SelectLicenses(m,
			[]es_module.LicenseType{
				es_module.LicenseTypeApache20,
				es_module.LicenseTypeBSD2Clause,
				es_module.LicenseTypeBSD3Clause,
				es_module.LicenseTypeMIT,
			},
			[]es_module.LicenseType{
				es_module.LicenseTypeAGPL,
				es_module.LicenseTypeGPL,
				es_module.LicenseTypeLGPL,
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
	return rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues)
}
