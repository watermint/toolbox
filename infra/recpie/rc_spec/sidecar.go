package rc_spec

import (
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recpie/rc_conn"
	"github.com/watermint/toolbox/infra/recpie/rc_recipe"
	"github.com/watermint/toolbox/infra/recpie/rc_vo_impl"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"sort"
	"strings"
)

func newSideCar(scr rc_recipe.SideCarRecipe) rc_recipe.Spec {
	path, name := Path(scr)
	cliPath := strings.Join(append(path, name), " ")
	vc := rc_vo_impl.NewValueContainer(scr.Requirement())
	scopes, usePersonal, useBusiness := authScopes(vc)

	return &SpecSideCar{
		scr:             scr,
		scv:             newSideCarValue(vc),
		name:            name,
		cliPath:         cliPath,
		connUsePersonal: usePersonal,
		connUseBusiness: useBusiness,
		connScopes:      scopes,
	}
}

func newSideCarValue(vc *rc_vo_impl.ValueContainer) rc_recipe.SpecValue {
	keys := make([]string, 0)
	for k := range vc.Values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return &SpecSideCarValue{
		vc:        vc,
		valueKeys: keys,
	}
}

func authScopes(vc *rc_vo_impl.ValueContainer) (scopes []string, usePersonal, useBusiness bool) {
	scopes = make([]string, 0)
	sc := make(map[string]bool)

	for _, v := range vc.Values {
		switch v0 := v.(type) {
		case rc_conn.ConnBusinessInfo:
			sc["business_info"] = true
			useBusiness = true
			v0.IsBusinessInfo()
		case rc_conn.ConnBusinessMgmt:
			sc["business_mgmt"] = true
			useBusiness = true
			v0.IsBusinessMgmt()
		case rc_conn.ConnBusinessFile:
			sc["business_file"] = true
			useBusiness = true
			v0.IsBusinessFile()
		case rc_conn.ConnBusinessAudit:
			sc["business_audit"] = true
			useBusiness = true
			v0.IsBusinessAudit()
		case rc_conn.ConnUserFile:
			sc["user_file"] = true
			usePersonal = true
			v0.IsUserFile()
		}
	}
	for s := range sc {
		scopes = append(scopes, s)
	}
	sort.Strings(scopes)

	return scopes, usePersonal, useBusiness
}

type SpecSideCarValue struct {
	vc        *rc_vo_impl.ValueContainer
	valueKeys []string
}

func (z *SpecSideCarValue) ValueNames() []string {
	return z.valueKeys
}

func (z *SpecSideCarValue) ValueDesc(name string) app_msg.Message {
	return app_msg.M(z.vc.MessageKey(name))
}

func (z *SpecSideCarValue) ValueDefault(name string) interface{} {
	switch d := z.vc.Values[name].(type) {
	case fd_file.Feed:
		return ""
	default:
		return d
	}
}

func (z *SpecSideCarValue) ValueCustomDefault(name string) app_msg.MessageOptional {
	return app_msg.M(z.vc.MessageKey(name) + ".default").AsOptional()
}

type SpecSideCar struct {
	scr             rc_recipe.SideCarRecipe
	scv             rc_recipe.SpecValue
	name            string
	cliPath         string
	connUsePersonal bool
	connUseBusiness bool
	connScopes      []string
}

func (z *SpecSideCar) Name() string {
	return z.name
}

func (z *SpecSideCar) Title() app_msg.Message {
	return Title(z.scr)
}

func (z *SpecSideCar) Desc() app_msg.MessageOptional {
	return Desc(z.scr).AsOptional()
}

func (z *SpecSideCar) CliArgs() app_msg.MessageOptional {
	return RecipeMessage(z.scr, "cli.args").AsOptional()
}

func (z *SpecSideCar) CliNote() app_msg.MessageOptional {
	return RecipeMessage(z.scr, "cli.note").AsOptional()
}

func (z *SpecSideCar) Reports() []rp_spec.ReportSpec {
	return z.scr.Reports()
}

func (z *SpecSideCar) CliPath() string {
	return z.cliPath
}

func (z *SpecSideCar) ConnUsePersonal() bool {
	return z.connUsePersonal
}

func (z *SpecSideCar) ConnUseBusiness() bool {
	return z.connUseBusiness
}

func (z *SpecSideCar) ConnScopes() []string {
	return z.connScopes
}

func (z *SpecSideCar) ValueNames() []string {
	return z.scv.ValueNames()
}

func (z *SpecSideCar) ValueDesc(name string) app_msg.Message {
	return z.scv.ValueDesc(name)
}

func (z *SpecSideCar) ValueDefault(name string) interface{} {
	return z.scv.ValueDefault(name)
}

func (z *SpecSideCar) ValueCustomDefault(name string) app_msg.MessageOptional {
	return z.scv.ValueCustomDefault(name)
}
