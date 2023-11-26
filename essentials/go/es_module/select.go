package es_module

import "errors"

var (
	ErrorNoLicense = errors.New("no license")
)

func SelectLicenses(module Module, preferred []LicenseType, disallowed []LicenseType) (license License, err error) {
	for _, lic := range module.Licenses() {
		for _, pl := range preferred {
			if lic.Type() == pl {
				return lic, nil
			}
		}
	}
	remaining := make([]License, 0)
	for _, lic := range module.Licenses() {
		da := false
		for _, dl := range disallowed {
			if lic.Type() == dl {
				da = true
				break
			}
		}
		if da {
			continue
		}
		remaining = append(remaining, lic)
	}

	if 0 < len(remaining) {
		return remaining[0], nil
	}
	return nil, ErrorNoLicense
}
