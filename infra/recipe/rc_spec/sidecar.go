package rc_spec

import (
	"flag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_vo_impl"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"sort"
	"strings"
)

func newSideCar(scr rc_recipe.SideCarRecipe) rc_recipe.Spec {
	path, name := Path(scr)
	cliPath := strings.Join(append(path, name), " ")
	vc := rc_vo_impl.NewValueContainer(scr.Requirement())
	scopes, usePersonal, useBusiness := authScopesFromVc(vc)

	return &SpecSideCar{
		scr:             scr,
		scv:             newSideCarValue(vc),
		vc:              vc,
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

func authScopesFromVc(vc *rc_vo_impl.ValueContainer) (scopes []string, usePersonal, useBusiness bool) {
	scopes = make([]string, 0)
	sc := make(map[string]bool)

	for _, v := range vc.Values {
		switch v0 := v.(type) {
		case rc_conn.OldConnBusinessInfo:
			sc["business_info"] = true
			useBusiness = true
			v0.IsBusinessInfo()
		case rc_conn.OldConnBusinessMgmt:
			sc["business_mgmt"] = true
			useBusiness = true
			v0.IsBusinessMgmt()
		case rc_conn.OldConnBusinessFile:
			sc["business_file"] = true
			useBusiness = true
			v0.IsBusinessFile()
		case rc_conn.OldConnBusinessAudit:
			sc["business_audit"] = true
			useBusiness = true
			v0.IsBusinessAudit()
		case rc_conn.OldConnUserFile:
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
	scr       rc_recipe.SideCarRecipe
	vc        *rc_vo_impl.ValueContainer
	valueKeys []string
}

func (z *SpecSideCarValue) SetFlags(f *flag.FlagSet, ui app_ui.UI) {
	z.vc.MakeFlagSet(f, ui)
}

func (z *SpecSideCarValue) ValueNames() []string {
	return z.valueKeys
}

func (z *SpecSideCarValue) ValueDesc(name string) app_msg.Message {
	return app_msg.M(z.vc.MessageKey(name))
}

func (z *SpecSideCarValue) ValueDefault(name string) interface{} {
	switch d := z.vc.Values[name].(type) {
	case fd_file.ModelFile:
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
	vc              *rc_vo_impl.ValueContainer
	name            string
	cliPath         string
	connUsePersonal bool
	connUseBusiness bool
	connScopes      []string
}

func (z *SpecSideCar) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *SpecSideCar) Feeds() map[string]fd_file.Spec {
	return map[string]fd_file.Spec{}
}

func (z *SpecSideCar) Debug() map[string]interface{} {
	return z.vc.Serialize()
}

func (z *SpecSideCar) ApplyValues(ctl app_control.Control, custom func(r rc_recipe.Recipe)) (r rc_recipe.Recipe, k rc_kitchen.Kitchen, err error) {
	vo := z.scr.Requirement()
	z.vc.Apply(vo)
	return z.scr, rc_kitchen.NewKitchen(ctl, vo), nil
}

func (z *SpecSideCar) SetFlags(f *flag.FlagSet, ui app_ui.UI) {
	z.scv.SetFlags(f, ui)
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

func (z *SpecSideCar) Reports() []rp_model.Spec {
	specs := make([]rp_model.Spec, 0)
	for _, s := range z.scr.Reports() {
		specs = append(specs, s)
	}
	return specs
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
