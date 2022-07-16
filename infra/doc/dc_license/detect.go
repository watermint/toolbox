package dc_license

import (
	"github.com/watermint/toolbox/essentials/go/go_module"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func Detect() (inventory []LicenseInfo, err error) {
	l := esl.Default()
	inventory = make([]LicenseInfo, 0)
	build, err := go_module.ScanBuild()
	if err != nil {
		l.Debug("Unable to scan build", esl.Error(err))
		return nil, err
	}

	for _, m := range build.Modules() {
		ml, err := go_module.SelectLicenses(m,
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
			l.Debug("Cannot select the license", esl.Error(err), esl.String("module", m.Path()))
			return nil, err
		}
		inventory = append(inventory,
			LicenseInfo{
				Package:     m.Path(),
				Version:     m.Version(),
				LicenseType: string(ml.Type()),
				LicenseBody: ml.Body(),
			},
		)
	}

	return inventory, nil
}
