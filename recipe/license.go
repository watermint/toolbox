package recipe

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/doc/dc_license"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type License struct {
	rc_recipe.RemarkTransient
	ErrorLicenseInfoNotFound app_msg.Message
	ToolboxHeader            app_msg.Message
	ThirdPartyHeader         app_msg.Message
	ThirdPartyNotice         app_msg.Message
	ThirdPartyLink           app_msg.Message
	ThirdPartyPackage        app_msg.Message
	ThirdPartyNoBody         app_msg.Message
}

func (z *License) Preset() {
}

func (z *License) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}

func (z *License) Exec(c app_control.Control) error {
	l := c.Log()
	ui := c.UI()
	licenseData, err := app_resource.Bundle().Data().Bytes("licenses.json")
	if err != nil {
		ui.Error(z.ErrorLicenseInfoNotFound)
		if c.Feature().IsTest() {
			return nil
		}
		return err
	}
	licenses := &dc_license.Licenses{}
	err = json.Unmarshal(licenseData, licenses)
	if err != nil {
		l.Warn("Unable to parse", esl.Error(err), esl.ByteString("data", licenseData))
		ui.Error(z.ErrorLicenseInfoNotFound)
		return err
	}

	ui.Header(z.ToolboxHeader)
	ui.Code(licenses.Project.LicenseBody)

	ui.Header(z.ThirdPartyHeader)
	ui.Info(z.ThirdPartyNotice)

	for _, pkg := range licenses.ThirdParty {
		ui.SubHeader(z.ThirdPartyPackage.With("Package", pkg.Package))
		if pkg.Url != "" {
			ui.Info(z.ThirdPartyLink.With("Url", pkg.Url))
			ui.Break()
		}
		if pkg.LicenseBody != "" {
			ui.Code(pkg.LicenseBody)
		} else {
			ui.Info(z.ThirdPartyNoBody.With("Pkg", pkg.Package))
		}
		ui.Break()
	}

	return nil
}
