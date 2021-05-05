package rc_spec

import (
	"flag"
	"fmt"
	"github.com/watermint/toolbox/essentials/collections/es_array"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_griddata"
	"github.com/watermint/toolbox/infra/data/da_json"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_options"
	"github.com/watermint/toolbox/infra/doc/dc_recipe"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_error_handler"
	"github.com/watermint/toolbox/infra/recipe/rc_group"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_value"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"os"
	"path/filepath"
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

func (z *specValueSelfContained) ErrorHandlers() []rc_error_handler.ErrorHandler {
	handlers := make([]rc_error_handler.ErrorHandler, 0)
	for _, vn := range z.repo.FieldNames() {
		if v, ok := z.repo.FieldValue(vn).(rc_recipe.ValueErrorHandler); ok {
			handlers = append(handlers, v.ErrorHandler())
		}
	}
	return handlers
}

func (z *specValueSelfContained) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.repo.Capture(ctl)
}

func (z *specValueSelfContained) Restore(j es_json.Json, ctl app_control.Control) (rcp rc_recipe.Recipe, err error) {
	l := ctl.Log()

	l.Debug("Restore")
	err = z.repo.Restore(j, ctl)
	if err != nil {
		l.Debug("Unable to restore", esl.Error(err))
		return nil, err
	}

	l.Debug("Apply")
	rcp = z.repo.Apply()

	l.Debug("Spin up")
	rcp, err = z.repo.SpinUp(ctl)
	if err != nil {
		l.Debug("Unable to spin up", esl.Error(err))
		return nil, err
	}
	return rcp, nil
}

func (z *specValueSelfContained) SpecId() string {
	return strings.Join(append(z.path, z.name), "-")
}

func (z *specValueSelfContained) Doc(ui app_ui.UI) *dc_recipe.Recipe {
	// feed ----
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

	// report
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

	// grid data input
	gridDataInputs := make([]*dc_recipe.DocGridDataInput, 0)
	gridDataInputNames := make([]string, 0)
	gridDataInputMap := make(map[string]*dc_recipe.DocGridDataInput)

	for _, g := range z.GridDataInput() {
		gridDataInputMap[g.Name()] = g.Doc(ui)
		gridDataInputNames = append(gridDataInputNames, g.Name())
	}
	sort.Strings(gridDataInputNames)
	for _, g := range gridDataInputNames {
		gridDataInputs = append(gridDataInputs, gridDataInputMap[g])
	}

	// grid data output
	gridDataOutputs := make([]*dc_recipe.DocGridDataOutput, 0)
	gridDataOutputNames := make([]string, 0)
	gridDataOutputMap := make(map[string]*dc_recipe.DocGridDataOutput)

	for _, g := range z.GridDataOutput() {
		gridDataOutputMap[g.Name()] = g.Doc(ui)
		gridDataOutputNames = append(gridDataOutputNames, g.Name())
	}
	sort.Strings(gridDataOutputNames)
	for _, g := range gridDataOutputNames {
		gridDataOutputs = append(gridDataOutputs, gridDataOutputMap[g])
	}

	// text input
	textInputs := make([]*dc_recipe.DocTextInput, 0)
	textInputNames := make([]string, 0)
	textInputMap := make(map[string]*dc_recipe.DocTextInput)

	for _, t := range z.TextInput() {
		textInputMap[t.Name()] = t.Doc(ui)
		textInputNames = append(textInputNames, t.Name())
	}
	sort.Strings(textInputNames)
	for _, t := range textInputNames {
		textInputs = append(textInputs, textInputMap[t])
	}

	// json input
	jsonInputs := make([]*dc_recipe.DocJsonInput, 0)
	jsonInputNames := make([]string, 0)
	jsonInputMap := make(map[string]*dc_recipe.DocJsonInput)

	for _, t := range z.JsonInput() {
		jsonInputMap[t.Name()] = t.Doc(ui)
		jsonInputNames = append(jsonInputNames, t.Name())
	}
	sort.Strings(jsonInputNames)
	for _, t := range jsonInputNames {
		jsonInputs = append(jsonInputs, jsonInputMap[t])
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
		Services:        z.Services(),
		IsSecret:        z.IsSecret(),
		IsConsole:       z.IsConsole(),
		IsExperimental:  z.IsExperimental(),
		IsIrreversible:  z.IsIrreversible(),
		IsTransient:     z.IsTransient(),
		Reports:         reports,
		Feeds:           feeds,
		Values:          values,
		GridDataInput:   gridDataInputs,
		GridDataOutput:  gridDataOutputs,
		TextInput:       textInputs,
		JsonInput:       jsonInputs,
	}
}

func (z *specValueSelfContained) New() rc_recipe.Spec {
	return NewSelfContained(z.scr)
}

func (z *specValueSelfContained) PrintUsage(ui app_ui.UI) {
	rc_group.UsageHeader(ui, z.Title(), app.BuildId)

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

func (z *specValueSelfContained) CliNameRef(media dc_index.MediaType, lg lang.Lang, relPath string) app_msg.Message {
	switch media {
	case dc_index.MediaRepository:
		path := filepath.Join(relPath, z.SpecId()+".md")
		return app_msg.Raw(fmt.Sprintf("[%s](%s)", z.CliPath(), path))
	case dc_index.MediaWeb:
		path := dc_index.WebDocPath(true, dc_index.WebCategoryCommand, z.SpecId(), lg)
		return app_msg.Raw(fmt.Sprintf("[%s](%s)", z.CliPath(), path))
	}
	panic("invalid media type")
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

func (z *specValueSelfContained) GridDataInput() map[string]da_griddata.GridDataInputSpec {
	return z.repo.GridDataInputSpecs()
}

func (z *specValueSelfContained) GridDataOutput() map[string]da_griddata.GridDataOutputSpec {
	return z.repo.GridDataOutputSpecs()
}

func (z *specValueSelfContained) TextInput() map[string]da_text.TextInputSpec {
	return z.repo.TextInputSpecs()
}

func (z *specValueSelfContained) JsonInput() map[string]da_json.JsonInputSpec {
	return z.repo.JsonInputSpecs()
}

func (z *specValueSelfContained) Feeds() map[string]fd_file.Spec {
	return z.repo.FeedSpecs()
}

func (z *specValueSelfContained) Services() []string {
	services := make([]string, 0)
	conns := z.repo.Conns()
	for _, c := range conns {
		services = append(services, c.ServiceName())
	}
	return es_array.NewByString(services...).Unique().Sort().AsStringArray()
}

func (z *specValueSelfContained) ConnUsePersonal() bool {
	use := false
	for _, c := range z.repo.Conns() {
		if c.ServiceName() == api_conn.ServiceDropbox {
			use = true
		}
	}
	return use
}

func (z *specValueSelfContained) ConnUseBusiness() bool {
	use := false
	for _, c := range z.repo.Conns() {
		if c.ServiceName() == api_conn.ServiceDropboxBusiness {
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
