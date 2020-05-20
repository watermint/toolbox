package recipe

import (
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"sort"
	"strings"
)

type License struct {
	rc_recipe.RemarkTransient
	ErrorLicenseInfoNotFound app_msg.Message
	ToolboxHeader            app_msg.Message
	ThirdPartyHeader         app_msg.Message
	ThirdPartyNotice         app_msg.Message
	ThirdPartyPackage        app_msg.Message
}

func (z *License) Preset() {
}

func (z *License) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}

func (z *License) Exec(c app_control.Control) error {
	ui := c.UI()
	tbxLicense, otherLicenses, order, err := LoadLicense(c)
	if err != nil {
		ui.Error(z.ErrorLicenseInfoNotFound)
		return nil
	}

	ui.Header(z.ToolboxHeader)
	ui.Code(strings.Join(tbxLicense, "\n"))

	ui.Header(z.ThirdPartyHeader)
	ui.Info(z.ThirdPartyNotice)

	for _, pkg := range order {
		pp := pkg
		if strings.HasPrefix(pp, "vendor/") {
			pp = pp[len("vendor/"):]
		}
		ui.SubHeader(z.ThirdPartyPackage.With("Package", pp))
		ui.Code(strings.Join(otherLicenses[pkg], "\n"))
	}

	return nil
}

func LoadLicense(ctl app_control.Control) (tbxLicense []string, otherLicenses map[string][]string, order []string, err error) {
	l := ctl.Log()
	lic, err := app_resource.Bundle().Data().Bytes("licenses.json")
	if err != nil {
		return nil, nil, nil, err
	}
	otherLicenses = make(map[string][]string)
	licenses := make(map[string][]string)
	if err = json.Unmarshal(lic, &licenses); err != nil {
		l.Error("Invalid license file format", esl.Error(err))
		return nil, nil, nil, err
	}

	if _, ok := licenses[app.Pkg]; !ok {
		l.Error("toolbox license not found")
		return nil, nil, nil, errors.New("toolbox license not found")
	}

	for pkg, ll := range licenses {
		if pkg == app.Pkg {
			tbxLicense = ll
		} else {
			otherLicenses[pkg] = ll
		}
	}

	deps := make([]string, 0)
	for pkg := range otherLicenses {
		if pkg != app.Pkg {
			deps = append(deps, pkg)
		}
	}
	sort.Strings(deps)

	order = make([]string, 0)
	order = append(order, deps...)

	return tbxLicense, licenses, order, nil
}
