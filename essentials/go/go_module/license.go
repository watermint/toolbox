package go_module

import "io/fs"

const (
	LicenseTypeUnknown = iota

	LicenseTypeNotFound

	LicenseTypeOther

	// LicenseTypeApache20 LicenseTypeApache Apache License 2.0
	LicenseTypeApache20

	// LicenseTypeBSD3Clause BSD 3-Clause "New" or "Revised" license
	LicenseTypeBSD3Clause

	// LicenseTypeBSD2Clause BSD 2-Clause "Simplified" or "FreeBSD" license
	LicenseTypeBSD2Clause

	// LicenseTypeGPL GNU General Public License (GPL)
	LicenseTypeGPL

	// LicenseTypeLGPL GNU Library or "Lesser" General Public License (LGPL)
	LicenseTypeLGPL

	// LicenseTypeMIT MIT license
	LicenseTypeMIT

	// LicenseTypeMPL Mozilla Public License
	LicenseTypeMPL

	// LicenseTypeCDDL Common Development and Distribution License
	LicenseTypeCDDL

	// LicenseTypeEclipse Eclipse Public License version
	LicenseTypeEclipse
)

type LicenseType int

type License interface {
	// Type returns license type
	Type() LicenseType

	// Body returns license body text
	Body() string
}

type licenseImpl struct {
	licenseType LicenseType
	licenseBody string
}

func (z licenseImpl) Type() LicenseType {
	return z.licenseType
}

func (z licenseImpl) Body() string {
	return z.licenseBody
}

func NewLicense(licenseType LicenseType, body string) License {
	return licenseImpl{
		licenseType: licenseType,
		licenseBody: body,
	}
}

func NewLicenseNotFound() License {
	return NewLicense(LicenseTypeNotFound, "")
}

func ScanLicense(target fs.FS) (licenses []License, err error) {
	return nil, err
}
