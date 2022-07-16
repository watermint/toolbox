package dc_license

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"sort"
	"strings"
)

type MsgLicense struct {
	ErrorLicenseInfoNotFound app_msg.Message
	ToolboxHeader            app_msg.Message
	ThirdPartyHeader         app_msg.Message
	ThirdPartyNotice         app_msg.Message
	ThirdPartyLink           app_msg.Message
	ThirdPartyVersion        app_msg.Message
	ThirdPartyPackage        app_msg.Message
	ThirdPartyNoBody         app_msg.Message
}

var (
	MLicense = app_msg.Apply(&MsgLicense{}).(*MsgLicense)
)

func Generate(c app_control.Control, ui app_ui.UI) error {
	l := c.Log()
	licenseData, err := app_resource.Bundle().Data().Bytes("licenses.json")
	if err != nil {
		ui.Error(MLicense.ErrorLicenseInfoNotFound)
		if c.Feature().IsTest() {
			return nil
		}
		return err
	}
	licenses := &Licenses{}
	err = json.Unmarshal(licenseData, licenses)
	if err != nil {
		l.Warn("Unable to parse", esl.Error(err), esl.ByteString("data", licenseData))
		ui.Error(MLicense.ErrorLicenseInfoNotFound)
		return err
	}

	ui.Header(MLicense.ToolboxHeader)
	ui.Code(licenses.Project.LicenseBody)

	ui.Header(MLicense.ThirdPartyHeader)
	ui.Info(MLicense.ThirdPartyNotice)

	sort.Slice(licenses.ThirdParty, func(i, j int) bool {
		return strings.Compare(licenses.ThirdParty[i].Package, licenses.ThirdParty[j].Package) < 0
	})

	for _, pkg := range licenses.ThirdParty {
		ui.SubHeader(MLicense.ThirdPartyPackage.With("Package", pkg.Package))
		if pkg.Version != "" {
			ui.Info(MLicense.ThirdPartyVersion.With("Version", pkg.Version))
			ui.Break()
		}
		if pkg.Url != "" {
			ui.Info(MLicense.ThirdPartyLink.With("Url", pkg.Url))
			ui.Break()
		}
		if pkg.LicenseBody != "" {
			ui.Code(pkg.LicenseBody)
		} else {
			ui.Info(MLicense.ThirdPartyNoBody.With("Pkg", pkg.Package))
		}
		ui.Break()
	}
	return nil
}
