package recipe

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"go.uber.org/zap"
	"sort"
	"strings"
)

type License struct {
}

func (license *License) Test(c app_control.Control) error {
	return nil
}

func (*License) Requirement() app_vo.ValueObject {
	return &app_vo.EmptyValueObject{}
}

func (*License) Exec(k app_kitchen.Kitchen) error {
	tbxLicense, otherLicenses, order, err := LoadLicense(k.Control())
	if err != nil {
		return err
	}

	for _, line := range tbxLicense {
		fmt.Println(line)
	}
	fmt.Printf("\n\n")
	fmt.Println(k.UI().Text("recipe.license.third_party_notice.head"))
	fmt.Printf("\n")
	fmt.Println(k.UI().Text("recipe.license.third_party_notice.body"))
	fmt.Printf("\n")

	for _, pkg := range order {
		pp := pkg
		if strings.HasPrefix(pp, "vendor/") {
			pp = pp[len("vendor/"):]
		}
		fmt.Println(pp + ":")
		fmt.Println(strings.Repeat("-", len(pp)+1))
		fmt.Printf("\n")
		lines := otherLicenses[pkg]
		for _, line := range lines {
			fmt.Println(line)
		}
		fmt.Printf("\n\n")
	}

	return nil
}

const (
	TbxPkg = "github.com/watermint/toolbox"
)

func LoadLicense(ctl app_control.Control) (tbxLicense []string, otherLicenses map[string][]string, order []string, err error) {
	l := ctl.Log()
	lic, err := ctl.Resource("licenses.json")
	if err != nil {
		return nil, nil, nil, err
	}
	otherLicenses = make(map[string][]string)
	licenses := make(map[string][]string)
	if err = json.Unmarshal(lic, &licenses); err != nil {
		l.Error("Invalid license file format", zap.Error(err))
		return nil, nil, nil, err
	}

	if _, ok := licenses[TbxPkg]; !ok {
		l.Error("toolbox license not found")
		return nil, nil, nil, errors.New("toolbox license not found")
	}

	for pkg, ll := range licenses {
		if pkg == TbxPkg {
			tbxLicense = ll
		} else {
			otherLicenses[pkg] = ll
		}
	}

	deps := make([]string, 0)
	for pkg := range otherLicenses {
		if pkg != TbxPkg {
			deps = append(deps, pkg)
		}
	}
	sort.Strings(deps)

	order = make([]string, 0)
	order = append(order, deps...)

	return tbxLicense, licenses, order, nil
}
