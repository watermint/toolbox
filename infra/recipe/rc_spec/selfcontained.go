package rc_spec

import (
	"flag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_value"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"reflect"
	"sort"
	"strings"
)

func newSelfContained(scr rc_recipe.SelfContainedRecipe) rc_recipe.Spec {
	path, name := Path(scr)
	cliPath := strings.Join(append(path, name), " ")

	vr := rc_value.NewValueRepository()
	if err := vr.Init(scr); err != nil {
		return nil
	}

	scr.Init()
	scopes, usePersonal, useBusiness := authScopesFromVr(vr)

	keys := make([]string, 0)
	for k := range vr.Values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	scv := &specValueSelfContained{
		scr:  scr,
		vr:   vr,
		keys: keys,
	}

	return &specSelfContained{
		scr:             scr,
		scv:             scv,
		vr:              vr,
		name:            name,
		cliPath:         cliPath,
		connUsePersonal: usePersonal,
		connUseBusiness: useBusiness,
		connScopes:      scopes,
	}
}

func authScopesFromVr(vr *rc_value.ValueRepository) (scopes []string, usePersonal, useBusiness bool) {
	scopes = make([]string, 0)
	sc := make(map[string]bool)

	for _, v := range vr.Values {
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

type specValueSelfContained struct {
	scr  rc_recipe.SelfContainedRecipe
	vr   *rc_value.ValueRepository
	keys []string
}

func (z *specValueSelfContained) ValueNames() []string {
	return z.keys
}

func (z *specValueSelfContained) ValueDesc(name string) app_msg.Message {
	return app_msg.M(z.vr.MessageKey(name))
}

func (z *specValueSelfContained) ValueDefault(name string) interface{} {
	return z.vr.Values[name]
}

func (z *specValueSelfContained) ValueCustomDefault(name string) app_msg.MessageOptional {
	return app_msg.M(z.vr.MessageKey(name) + ".default").AsOptional()
}

func (z *specValueSelfContained) SetFlags(f *flag.FlagSet, ui app_ui.UI) {
	z.vr.MakeFlagSet(f, ui)
}

type specSelfContained struct {
	scv             rc_recipe.SpecValue
	scr             rc_recipe.SelfContainedRecipe
	vr              *rc_value.ValueRepository
	name            string
	cliPath         string
	connUsePersonal bool
	connUseBusiness bool
	connScopes      []string
}

func (z *specSelfContained) ValueNames() []string {
	return z.scv.ValueNames()
}

func (z *specSelfContained) ValueDesc(name string) app_msg.Message {
	return z.scv.ValueDesc(name)
}

func (z *specSelfContained) ValueDefault(name string) interface{} {
	switch v := z.scv.ValueDefault(name).(type) {
	case fd_file.RowFeed:
		return ""
	default:
		return v
	}
}

func (z *specSelfContained) ValueCustomDefault(name string) app_msg.MessageOptional {
	return z.scv.ValueCustomDefault(name)
}

func (z *specSelfContained) SetFlags(f *flag.FlagSet, ui app_ui.UI) {
	z.scv.SetFlags(f, ui)
}

func (z *specSelfContained) Name() string {
	return z.name
}

func (z *specSelfContained) Title() app_msg.Message {
	return Title(z.scr)
}

func (z *specSelfContained) Desc() app_msg.MessageOptional {
	return Desc(z.scr).AsOptional()
}

func (z *specSelfContained) CliPath() string {
	return z.cliPath
}

func (z *specSelfContained) CliArgs() app_msg.MessageOptional {
	return RecipeMessage(z.scr, "cli.args").AsOptional()
}

func (z *specSelfContained) CliNote() app_msg.MessageOptional {
	return RecipeMessage(z.scr, "cli.note").AsOptional()
}

func (z *specSelfContained) Reports() []rp_model.Spec {
	rs := make([]rp_model.Spec, 0)
	for _, r := range z.vr.Reports {
		rs = append(rs, r.Spec())
	}
	return rs
}

func (z *specSelfContained) Feeds() map[string]fd_file.Spec {
	return z.vr.FeedSpecs()
}

func (z *specSelfContained) ConnUsePersonal() bool {
	return z.connUsePersonal
}

func (z *specSelfContained) ConnUseBusiness() bool {
	return z.connUseBusiness
}

func (z *specSelfContained) ConnScopes() []string {
	return z.connScopes
}

func (z *specSelfContained) ApplyValues(ctl app_control.Control) (rcp rc_recipe.Recipe, k rc_kitchen.Kitchen, err error) {
	rt := reflect.TypeOf(z.scr).Elem()
	newScr := reflect.New(rt).Interface().(rc_recipe.SelfContainedRecipe)
	newVr := z.vr.Fork(ctl)
	err = newVr.Apply(newScr, ctl)
	if err != nil {
		return nil, nil, err
	}

	app_msg.Apply(newScr)
	return newScr, rc_kitchen.NewKitchen(ctl, newScr), nil
}

func (z *specSelfContained) SerializeValues() map[string]interface{} {
	return z.vr.Serialize()
}
