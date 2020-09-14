package dc_license

import (
	"bytes"
	"encoding/csv"
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

type Licenses struct {
	Project    LicenseInfo   `json:"project"`
	ThirdParty []LicenseInfo `json:"third_party"`
}

type LicenseInfo struct {
	Package     string `json:"package"`
	Url         string `json:"url"`
	LicenseType string `json:"license_type"`
	LicenseBody string `json:"license_body"`
}

func LoadLicenses(c app_control.Control, path string) (licenses Licenses, err error) {
	l := c.Log()
	lis, err := loadLicenseInfos(c, path)
	if err != nil {
		l.Debug("Unable to load", esl.Error(err))
		return
	}

	thirdPartyPkgs := make([]string, 0)
	thirdPartyLics := make([]LicenseInfo, 0)
	dict := make(map[string]LicenseInfo)

	for _, lic := range lis {
		dict[lic.Package] = lic
		if lic.Package != app.Pkg {
			thirdPartyPkgs = append(thirdPartyPkgs, lic.Package)
		}
	}
	sort.Strings(thirdPartyPkgs)

	if prj, ok := dict[app.Pkg]; !ok {
		l.Debug("Unable to find the project license file")
		return licenses, errors.New("unable to find the project license file")
	} else {
		licenses.Project = prj
	}

	for _, pkg := range thirdPartyPkgs {
		thirdPartyLics = append(thirdPartyLics, dict[pkg])
	}
	licenses.ThirdParty = thirdPartyLics

	return licenses, nil
}

func loadLicenseInfos(c app_control.Control, path string) (licenses []LicenseInfo, err error) {
	l := c.Log()

	l.Debug("Load TOC file")
	tocPath := filepath.Join(path, "licenses.csv")
	toc, err := ioutil.ReadFile(tocPath)
	if err != nil {
		l.Debug("Unable to load TOC", esl.Error(err))
		return nil, err
	}

	licenses = make([]LicenseInfo, 0)

	tocCsv := csv.NewReader(bytes.NewReader(toc))
	for {
		line, err := tocCsv.Read()
		if err == io.EOF {
			l.Debug("EOF")
			return licenses, nil
		}
		if err != nil {
			l.Debug("Unable to load TOC file", esl.Error(err))
			return nil, err
		}

		if len(line) < 3 {
			l.Debug("Unknown data format", esl.Any("line", line))
			return nil, errors.New("unknown data format")
		}

		pkg := line[0]
		url := line[1]
		lic := line[2]

		if url == "Unknown" {
			url = ""
		}
		if lic == "Unknown" {
			lic = ""
		}

		// Looking for a license file
		licPath := filepath.Join(path, "licenses", pkg)
		l.Debug("Looking for a license file", esl.String("path", licPath))

		body, err := loadLicenseFiles(c, licPath)
		if err != nil {
			l.Debug("Unable to load a license file", esl.Error(err))
			return nil, err
		}

		licenses = append(licenses, LicenseInfo{
			Package:     pkg,
			Url:         url,
			LicenseType: lic,
			LicenseBody: body,
		})
	}
}

func loadLicenseFiles(c app_control.Control, path string) (body string, err error) {
	l := c.Log()

	fileEntries, err := ioutil.ReadDir(path)
	if err != nil {
		l.Debug("Unable to load license path, ignore this pkg", esl.Error(err))
		return "", nil
	}

	priorities := []string{"LICENSE", "LICENSE.txt", "LICENSE.md"}

	for _, target := range priorities {
		for _, entry := range fileEntries {
			if entry.Name() == target {
				return loadLicenseFile(c, filepath.Join(path, entry.Name()))
			}
		}
	}

	// otherwise, returns a body first found
	for _, entry := range fileEntries {
		if body, err := loadLicenseFile(c, filepath.Join(path, entry.Name())); err == nil {
			return body, nil
		}
	}

	// Return successfully, if a license file not found
	return "", nil
}

func loadLicenseFile(c app_control.Control, path string) (body string, err error) {
	l := c.Log()
	l.Debug("Load license file", esl.String("path", path))
	bodyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		l.Debug("Unable to load a file", esl.Error(err))
		return "", err
	}
	return string(bodyBytes), nil
}

func MakeTestData(path string) error {
	licenses := []LicenseInfo{
		{
			Package:     "github.com/watermint/toolbox",
			Url:         "Unknown",
			LicenseType: "MIT",
			LicenseBody: "TEST MIT LICENSE toolbox",
		},
		{
			Package:     "github.com/watermint/bwlimit",
			Url:         "https://github.com/watermint/bwlimit/blob/master/LICENSE.md",
			LicenseType: "MIT",
			LicenseBody: "TEST MIT LICENSE bwlimit",
		},
		{
			Package:     "example.com/unknown_mit_no_file",
			Url:         "Unknown",
			LicenseType: "MIT",
			LicenseBody: "",
		},
		{
			Package:     "example.com/unknown_bsd_with_file",
			Url:         "Unknown",
			LicenseType: "BSD-3-Clause",
			LicenseBody: "TEST BSD LICENSE unknown_bsd_with_file",
		},
	}

	// Make TOC
	tocData := &bytes.Buffer{}
	tocCsv := csv.NewWriter(tocData)

	for _, lic := range licenses {
		if err := tocCsv.Write([]string{lic.Package, lic.Url, lic.LicenseType}); err != nil {
			return err
		}
	}
	tocCsv.Flush()

	if err := ioutil.WriteFile(filepath.Join(path, "licenses.csv"), tocData.Bytes(), 0644); err != nil {
		return err
	}

	filenames := map[string]string{
		"github.com/watermint/toolbox":      "README.md",
		"github.com/watermint/bwlimit":      "README",
		"example.com/unknown_bsd_with_file": "NOTICE",
	}

	// Make individual license files
	for _, lic := range licenses {
		if file, ok := filenames[lic.Package]; ok {
			licPath := filepath.Join(path, "licenses", lic.Package)
			if err := os.MkdirAll(licPath, 0755); err != nil {
				return err
			}
			licFile := filepath.Join(licPath, file)
			if err := ioutil.WriteFile(licFile, []byte(lic.LicenseBody), 0644); err != nil {
				return err
			}
		}
	}

	return nil
}
