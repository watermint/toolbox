package rc_spec

import (
	"flag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_value"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"go.uber.org/zap"
	"sort"
	"strings"
)

func newSelfContained(scr rc_recipe.SelfContainedRecipe) rc_recipe.Spec {
	path, name := rc_recipe.Path(scr)
	cliPath := strings.Join(append(path, name), " ")

	repo := rc_value.NewRepository(scr)

	return &specValueSelfContained{
		name:    name,
		cliPath: cliPath,
		scr:     scr,
		repo:    repo,
	}
}

type specValueSelfContained struct {
	name    string
	cliPath string
	scr     rc_recipe.SelfContainedRecipe
	repo    rc_value.Repository
}

func (z *specValueSelfContained) SpinDown(ctl app_control.Control) error {
	return z.repo.SpinDown(ctl)
}

func (z *specValueSelfContained) ValueNames() []string {
	return z.repo.FieldNames()
}

func (z *specValueSelfContained) ValueDesc(name string) app_msg.Message {
	return z.repo.FieldDesc(name)
}

func (z *specValueSelfContained) ValueDefault(name string) interface{} {
	return z.repo.FieldValueText(name)
}

func (z *specValueSelfContained) ValueCustomDefault(name string) app_msg.MessageOptional {
	return z.repo.FieldCustomDefault(name)
}

func (z *specValueSelfContained) SetFlags(f *flag.FlagSet, ui app_ui.UI) {
	z.repo.ApplyFlags(f, ui)
}

func (z *specValueSelfContained) Name() string {
	return z.name
}

func (z *specValueSelfContained) Title() app_msg.Message {
	return rc_recipe.Title(z.scr)
}

func (z *specValueSelfContained) Desc() app_msg.MessageOptional {
	return rc_recipe.Desc(z.scr).AsOptional()
}

func (z *specValueSelfContained) CliPath() string {
	return z.cliPath
}

func (z *specValueSelfContained) CliArgs() app_msg.MessageOptional {
	return rc_recipe.RecipeMessage(z.scr, "cli.args").AsOptional()
}

func (z *specValueSelfContained) CliNote() app_msg.MessageOptional {
	return rc_recipe.RecipeMessage(z.scr, "cli.note").AsOptional()
}

func (z *specValueSelfContained) Messages() []app_msg.Message {
	panic("implement me")
}

func (z *specValueSelfContained) Reports() []rp_model.Spec {
	reps := make([]rp_model.Spec, 0)
	for _, s := range z.repo.ReportSpecs() {
		reps = append(reps, s)
	}
	return reps
}

func (z *specValueSelfContained) Feeds() map[string]fd_file.Spec {
	return z.repo.FeedSpecs()
}

func (z *specValueSelfContained) ConnUsePersonal() bool {
	use := false
	for _, c := range z.repo.Conns() {
		if c.IsPersonal() {
			use = true
		}
	}
	return use
}

func (z *specValueSelfContained) ConnUseBusiness() bool {
	use := false
	for _, c := range z.repo.Conns() {
		if c.IsBusiness() {
			use = true
		}
	}
	return use
}

func (z *specValueSelfContained) ConnScopes() []string {
	scopes := make([]string, 0)
	scopeLabels := make(map[string]bool)
	for _, c := range z.repo.Conns() {
		scopeLabels[c.ScopeLabel()] = true
	}
	for s := range scopeLabels {
		scopes = append(scopes, s)
	}
	sort.Strings(scopes)
	return scopes
}

func (z *specValueSelfContained) SpinUp(ctl app_control.Control, custom func(r rc_recipe.Recipe)) (rcp rc_recipe.Recipe, k rc_kitchen.Kitchen, err error) {
	l := ctl.Log().With(zap.String("name", z.name))
	rcp = z.repo.Apply()
	custom(rcp)
	_, err = z.repo.SpinUp(ctl)
	if err != nil {
		l.Debug("Unable to spin up", zap.Error(err))
		return nil, nil, err
	}
	return rcp, rc_kitchen.NewKitchen(ctl, rcp), nil
}

func (z *specValueSelfContained) Debug() map[string]interface{} {
	return z.repo.Debug()
}
