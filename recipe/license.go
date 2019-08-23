package recipe

import (
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"sort"
)

type License struct {
}

func (*License) Requirement() app_vo.ValueObject {
	return &app_vo.EmptyValueObject{}
}

func (*License) Exec(k app_kitchen.Kitchen) error {
	licenses, order, err := LoadLicense(k.Control())
	if err != nil {
		return err
	}

	for _, pkg := range order {
		k.UI().Header("raw", app_msg.P("Raw", pkg))
		lines := licenses[pkg]
		for _, line := range lines {
			k.UI().Info("raw", app_msg.P("Raw", line))
		}
	}

	return nil
}

const (
	TbxPkg = "github.com/watermint/toolbox"
)

func LoadLicense(ctl app_control.Control) (licenses map[string][]string, order []string, err error) {
	l := ctl.Log()
	lic, err := ctl.Resource("licenses.json")
	if err != nil {
		return nil, nil, err
	}
	licenses = make(map[string][]string)
	if err = json.Unmarshal(lic, &licenses); err != nil {
		l.Error("Invalid license file format", zap.Error(err))
		return nil, nil, err
	}

	if _, ok := licenses[TbxPkg]; !ok {
		l.Error("toolbox license not found")
		return nil, nil, errors.New("toolbox license not found")
	}

	deps := make([]string, 0)
	for pkg := range licenses {
		if pkg != TbxPkg {
			deps = append(deps, pkg)
		}
	}
	sort.Strings(deps)

	order = make([]string, 0)
	order = append(order, TbxPkg)
	order = append(order, deps...)

	return licenses, order, nil
}
