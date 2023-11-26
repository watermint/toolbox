package es_module

import (
	"github.com/watermint/toolbox/essentials/http/es_download"
	"github.com/watermint/toolbox/essentials/log/esl"
	"io/fs"
	"io/ioutil"
	"path"
	"regexp"
	"strings"
)

const (
	LicenseTypeUnknown = "unknown"

	LicenseTypeNotFound = "not_found"

	LicenseTypeOther = "other"

	// LicenseTypeApache20 LicenseTypeApache Apache License 2.0
	LicenseTypeApache20 = "apache2_0"

	// LicenseTypeBSD3Clause BSD 3-Clause "New" or "Revised" license
	LicenseTypeBSD3Clause = "bsd3"

	// LicenseTypeBSD2Clause BSD 2-Clause "Simplified" or "FreeBSD" license
	LicenseTypeBSD2Clause = "bsd2"

	// LicenseTypeAGPL GNU Affero General Public License (AGPL)
	LicenseTypeAGPL = "agpl"

	// LicenseTypeGPL GNU General Public License (GPL)
	LicenseTypeGPL = "gpl"

	// LicenseTypeLGPL GNU Library or "Lesser" General Public License (LGPL)
	LicenseTypeLGPL = "lgpl"

	// LicenseTypeMIT MIT license
	LicenseTypeMIT = "mit"

	// LicenseTypeMPL Mozilla Public License
	LicenseTypeMPL = "mpl"

	// LicenseTypeCDDL Common Development and Distribution License
	LicenseTypeCDDL = "cddl"

	// LicenseTypeEclipse Eclipse Public License version
	LicenseTypeEclipse = "eclipse"

	// LicenseTypeFreeType FreeType license
	LicenseTypeFreeType = "freetype"

	// LicenseTypeUnlicense Unlicense
	LicenseTypeUnlicense = "unlicense"
)

type LicenseType string

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
	l := esl.Default()
	licenses = make([]License, 0)
	err = fs.WalkDir(target, ".", func(p string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		ll := l.With(esl.String("path", p))
		lname := strings.ToLower(strings.TrimSuffix(p, path.Ext(p)))

		switch lname {
		case "copying", "licence", "license", "notice", "readme", "unlicence", "unlicense":
			lf, err := target.Open(p)
			if err != nil {
				ll.Debug("Unable to open the license file", esl.Error(err))
				return err
			}
			body, err := ioutil.ReadAll(lf)
			if err != nil {
				ll.Debug("Unable to read the license file", esl.Error(err))
				return err
			}
			_ = lf.Close()

			fileLicenses := ScanLicenseBody(string(body))
			if 0 < len(fileLicenses) {
				licenses = append(licenses, fileLicenses...)
			}
		}

		return nil
	})
	return licenses, err
}

func ScanLicenseBody(body string) (licenses []License) {
	lbody := strings.ToLower(body)
	lbody = strings.ReplaceAll(lbody, "\r\n", " ")
	lbody = strings.ReplaceAll(lbody, "\n", " ")
	lbody = strings.ReplaceAll(lbody, "\t", " ")

	ws := regexp.MustCompile(`\s{2,}`)
	lbody = ws.ReplaceAllLiteralString(lbody, " ")

	scanners := []licenseScanner{
		scanLicenseAGPL,
		scanLicenseApache,
		scanLicenseBSD,
		scanLicenseCDDL,
		scanLicenseEclipse,
		scanLicenseGPL,
		scanLicenseLGPL,
		scanLicenseMIT,
		scanLicenseMPL,
		scanLicenseFreeType,
		scanLicenseUnlicense,
	}

	licenses = make([]License, 0)

	for _, scanner := range scanners {
		if license, match := scanner(lbody, body); match {
			licenses = append(licenses, license)
		}
	}

	return licenses
}

type licenseScanner func(lbody, body string) (license License, match bool)

func scanLicenseUnlicense(lbody, body string) (license License, match bool) {
	if strings.Contains(lbody, "this is free and unencumbered software released into the public domain") {
		return NewLicense(LicenseTypeUnlicense, body), true
	}
	return nil, false
}

func scanLicenseApache(lbody, body string) (license License, match bool) {
	if strings.Contains(lbody, "apache license") &&
		strings.Contains(lbody, "version 2.0") &&
		strings.Contains(lbody, "http://www.apache.org/licenses") {
		return NewLicense(LicenseTypeApache20, body), true
	}
	return nil, false
}

func scanLicenseBSD(lbody, body string) (license License, match bool) {
	if !(strings.Contains(lbody, "redistribution and use in source and binary forms") &&
		strings.Contains(lbody, "with or without modification") &&
		strings.Contains(lbody, "are permitted provided that the following conditions are met")) {
		return nil, false
	}

	if strings.Contains(lbody, "neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission") {
		return NewLicense(LicenseTypeBSD3Clause, body), true
	}
	return NewLicense(LicenseTypeBSD2Clause, body), true
}

func scanLicenseAGPL(lbody, body string) (license License, match bool) {
	if strings.Contains(lbody, "gnu affero general public license") &&
		strings.Contains(lbody, "version 3, 19 november 2007") {
		return NewLicense(LicenseTypeAGPL, body), true
	}
	return nil, false
}

func scanLicenseGPL(lbody, body string) (license License, match bool) {
	if strings.Contains(lbody, "gnu general public license version 2, june 1991") {
		return NewLicense(LicenseTypeGPL, body), true
	}
	if strings.Contains(lbody, "gnu general public license version 3, 29 june 2007") {
		return NewLicense(LicenseTypeGPL, body), true
	}
	return nil, false
}

func scanLicenseLGPL(lbody, body string) (license License, match bool) {
	if strings.Contains(lbody, "gnu lesser general public license version 3, 29 june 2007") {
		return NewLicense(LicenseTypeLGPL, body), true
	}
	return nil, false
}

func scanLicenseMIT(lbody, body string) (license License, match bool) {
	if strings.Contains(lbody, "permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files") {
		return NewLicense(LicenseTypeMIT, body), true
	}
	return nil, false
}

func scanLicenseMPL(lbody, body string) (license License, match bool) {
	if strings.Contains(lbody, "the mozilla public license") {
		return NewLicense(LicenseTypeMPL, body), true
	}
	return nil, false
}

func scanLicenseCDDL(lbody, body string) (license License, match bool) {
	if strings.Contains(lbody, "common development and distribution") {
		return NewLicense(LicenseTypeCDDL, body), true
	}
	return nil, false
}

func scanLicenseEclipse(lbody, body string) (license License, match bool) {
	if strings.Contains(lbody, "eclipse public license") {
		return NewLicense(LicenseTypeEclipse, body), true
	}
	return nil, false
}

func scanLicenseFreeType(lbody, body string) (license License, match bool) {
	if strings.Contains(lbody, "the freetype license, which is similar to the original bsd license with an advertising clause") {
		originalBodyURL := "https://gitlab.freedesktop.org/freetype/freetype/-/raw/master/docs/FTL.TXT"
		origianlBody, err := es_download.DownloadText(esl.Default(), originalBodyURL)
		if err != nil {
			return NewLicense(LicenseTypeFreeType, body), true
		} else {
			return NewLicense(LicenseTypeFreeType, origianlBody), true
		}
	}
	if strings.Contains(lbody, "the freetype project license") &&
		strings.Contains(lbody, "2006-jan-27") {
		return NewLicense(LicenseTypeFreeType, body), true
	}
	return nil, false
}
