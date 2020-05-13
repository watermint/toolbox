package rc_spec

import (
	"flag"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/doc/dc_options"
	"github.com/watermint/toolbox/infra/doc/dc_recipe"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_group"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_value"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"os"
	"sort"
	"strings"
)

type MsgSelfContained struct {
	IsExperimental                app_msg.Message
	IsIrreversible                app_msg.Message
	IsExperimentalAndIrreversible app_msg.Message
	RecipeHeaderUsage             app_msg.Message
	RecipeUsage                   app_msg.Message
	RecipeAvailableFlags          app_msg.Message
	RecipeCommonFlags             app_msg.Message
}

var (
	MSelfContained = app_msg.Apply(&MsgSelfContained{}).(*MsgSelfContained)
)

func NewSelfContained(scr rc_recipe.Recipe) rc_recipe.Spec {
	var ann rc_recipe.Annotation
	var repo rc_recipe.Repository

	switch rr := scr.(type) {
	case rc_recipe.Annotation:
		repo = rc_value.NewRepository(rr.Seed())
		ann = rr
		scr = rr.Seed()

	default:
		ann = rc_recipe.NewAnnotated(scr)
		repo = rc_value.NewRepository(scr)
	}

	path, name := es_reflect.Path(rc_recipe.BasePackage, scr)
	cliPath := strings.Join(append(path, name), " ")

	return &specValueSelfContained{
		path:       path,
		name:       name,
		cliPath:    cliPath,
		scr:        scr,
		repo:       repo,
		annotation: ann,
	}
}

type specValueSelfContained struct {
	path       []string
	name       string
	cliPath    string
	annotation rc_recipe.Annotation
	scr        rc_recipe.Recipe
	repo       rc_recipe.Repository
}

func (z *specValueSelfContained) Doc(ui app_ui.UI) *dc_recipe.Recipe {
	feeds := make([]*dc_recipe.Feed, 0)
	feedNames := make([]string, 0)
	feedMaps := make(map[string]*dc_recipe.Feed)

	for _, f := range z.Feeds() {
		feedMaps[f.Name()] = f.Doc(ui)
		feedNames = append(feedNames, f.Name())
	}
	sort.Strings(feedNames)
	for _, f := range feedNames {
		feeds = append(feeds, feedMaps[f])
	}

	reports := make([]*dc_recipe.Report, 0)
	reportNames := make([]string, 0)
	reportMap := make(map[string]*dc_recipe.Report)

	for _, r := range z.Reports() {
		reportMap[r.Name()] = r.Doc(ui)
		reportNames = append(reportNames, r.Name())
	}
	sort.Strings(reportNames)
	for _, r := range reportNames {
		reports = append(reports, reportMap[r])
	}

	values := make([]*dc_recipe.Value, 0)
	valueNames := make([]string, len(z.ValueNames()))
	copy(valueNames, z.ValueNames())
	sort.Strings(valueNames)
	for _, vn := range valueNames {
		var dv app_msg.Message
		cdv := z.ValueCustomDefault(vn)
		if ui.Exists(cdv) {
			dv = cdv
		} else {
			dv = app_msg.Raw(z.ValueDefault(vn))
		}
		v := z.Value(vn)
		if v == nil {
			values = append(values,
				&dc_recipe.Value{
					Name:    vn,
					Default: ui.Text(dv),
					Desc:    ui.Text(z.ValueDesc(vn)),
				},
			)
		} else {
			tn, ta := v.Spec()
			values = append(values,
				&dc_recipe.Value{
					Name:     vn,
					Default:  ui.Text(dv),
					Desc:     ui.Text(z.ValueDesc(vn)),
					TypeName: tn,
					TypeAttr: ta,
				},
			)
		}
	}

	return &dc_recipe.Recipe{
		Name:            z.Name(),
		Title:           ui.Text(z.Title()),
		Desc:            ui.TextOrEmpty(z.Desc()),
		Remarks:         ui.TextOrEmpty(z.Remarks()),
		Path:            z.CliPath(),
		CliArgs:         ui.TextOrEmpty(z.CliArgs()),
		CliNote:         ui.TextOrEmpty(z.CliNote()),
		ConnUsePersonal: z.ConnUsePersonal(),
		ConnUseBusiness: z.ConnUseBusiness(),
		ConnScopes:      z.ConnScopeMap(),
		IsSecret:        z.IsSecret(),
		IsConsole:       z.IsConsole(),
		IsExperimental:  z.IsExperimental(),
		IsIrreversible:  z.IsIrreversible(),
		Feeds:           feeds,
		Reports:         reports,
		Values:          values,
	}
}

func (z *specValueSelfContained) New() rc_recipe.Spec {
	return NewSelfContained(z.scr)
}

func (z *specValueSelfContained) PrintUsage(ui app_ui.UI) {
	rc_group.UsageHeader(ui, z.Title(), app.Version)

	ui.Header(MSelfContained.RecipeHeaderUsage)
	ui.Info(MSelfContained.RecipeUsage.
		With("Exec", os.Args[0]).
		With("Recipe", z.CliPath()).
		With("Args", ui.TextOrEmpty(z.CliArgs())))

	ui.Break()
	ui.Header(MSelfContained.RecipeCommonFlags)
	com := NewCommonValue()
	dc_options.PrintOptionsTable(ui, com)

	ui.Header(MSelfContained.RecipeAvailableFlags)
	dc_options.PrintOptionsTable(ui, z)

	ui.Break()
}

func (z *specValueSelfContained) Path() (path []string, name string) {
	return z.path, z.name
}

func (z *specValueSelfContained) IsSecret() bool {
	if z.annotation != nil {
		return z.annotation.IsSecret()
	}
	return false
}

func (z *specValueSelfContained) IsConsole() bool {
	if z.annotation != nil {
		return z.annotation.IsConsole()
	}
	return false
}

func (z *specValueSelfContained) IsExperimental() bool {
	if z.annotation != nil {
		return z.annotation.IsExperimental()
	}
	return false
}

func (z *specValueSelfContained) IsIrreversible() bool {
	if z.annotation != nil {
		return z.annotation.IsIrreversible()
	}
	return false
}

func (z *specValueSelfContained) IsTransient() bool {
	if z.annotation != nil {
		return z.annotation.IsTransient()
	}
	return false
}

func (z *specValueSelfContained) Value(name string) rc_recipe.Value {
	return z.repo.FieldValue(name)
}

func (z *specValueSelfContained) ConnScopeMap() map[string]string {
	scopes := make(map[string]string)
	for k, v := range z.repo.Conns() {
		scopes[k] = v.ScopeLabel()
	}
	return scopes
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
	return app_msg.ObjMessage(z.scr, "title")
}

func (z *specValueSelfContained) Desc() app_msg.MessageOptional {
	return app_msg.ObjMessage(z.scr, "desc").AsOptional()
}

func (z *specValueSelfContained) Remarks() app_msg.MessageOptional {
	switch {
	case z.IsExperimental() && z.IsIrreversible():
		return MSelfContained.IsExperimentalAndIrreversible.AsOptional()
	case z.IsIrreversible():
		return MSelfContained.IsIrreversible.AsOptional()
	case z.IsExperimental():
		return MSelfContained.IsExperimental.AsOptional()
	default:
		return app_msg.Raw("").AsOptional()
	}
}

func (z *specValueSelfContained) CliPath() string {
	return z.cliPath
}

func (z *specValueSelfContained) CliArgs() app_msg.MessageOptional {
	return app_msg.ObjMessage(z.scr, "cli.args").AsOptional()
}

func (z *specValueSelfContained) CliNote() app_msg.MessageOptional {
	return app_msg.ObjMessage(z.scr, "cli.note").AsOptional()
}

func (z *specValueSelfContained) Messages() []app_msg.Message {
	return z.repo.Messages()
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

func (z *specValueSelfContained) SpinUp(ctl app_control.Control, custom func(r rc_recipe.Recipe)) (rcp rc_recipe.Recipe, err error) {
	l := ctl.Log().With(esl.String("name", z.name))
	rcp = z.repo.Apply()
	custom(rcp)
	z.repo.ApplyCustom()
	_, err = z.repo.SpinUp(ctl)
	if err != nil {
		l.Debug("Unable to spin up", esl.Error(err))
		return nil, err
	}
	return rcp, nil
}

func (z *specValueSelfContained) Debug() map[string]interface{} {
	return z.repo.Debug()
}
