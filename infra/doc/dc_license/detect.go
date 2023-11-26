package dc_license

import (
	"github.com/watermint/toolbox/essentials/go/es_module"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func Detect() (inventory []LicenseInfo, err error) {
	l := esl.Default()
	inventory = make([]LicenseInfo, 0)
	build, err := es_module.ScanBuild()
	if err != nil {
		l.Debug("Unable to scan build", esl.Error(err))
		return nil, err
	}

	for _, m := range build.Modules() {
		ml, err := es_module.SelectLicenses(m,
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
