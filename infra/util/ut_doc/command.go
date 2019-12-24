package ut_doc

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"errors"
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"go.uber.org/zap"
	"os"
	"sort"
	"strings"
	"text/template"
)

func NewCommand(ctl app_control.Control, path string, toStdout bool) *Commands {
	return &Commands{
		ctl:      ctl,
		toStdout: toStdout,
		path:     path,
	}
}

type Commands struct {
	ctl      app_control.Control
	toStdout bool
	path     string
}

func (z *Commands) optionsTable(spec rc_recipe.SpecValue) string {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	mc := z.ctl.Messages()

	mui := app_ui.NewMarkdown(mc, w, false)
	mt := mui.InfoTable("")

	mt.Header(
		app_msg.M("recipe.dev.doc.options.header.option"),
		app_msg.M("recipe.dev.doc.options.header.description"),
		app_msg.M("recipe.dev.doc.options.header.default"),
	)

	if len(spec.ValueNames()) < 1 {
		return ""
	}

	for _, k := range spec.ValueNames() {
		vd := spec.ValueDefault(k)
		vkd := spec.ValueCustomDefault(k)
		if mc.Exists(vkd.Key()) {
			vd = mc.Text(vkd.Key())
		}

		mt.Row(
			app_msg.M("recipe.dev.doc.options.body.option", app_msg.P{"Option": strcase.ToKebab(k)}),
			spec.ValueDesc(k),
			app_msg.M("raw", app_msg.P{"Raw": vd}),
		)
	}

	mt.Flush()
	w.Flush()
	return b.String()
}

func (z *Commands) reportTable(rs rp_model.Spec) string {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	mc := z.ctl.Messages()

	mui := app_ui.NewMarkdown(mc, w, false)
	mt := mui.InfoTable("")

	mt.Header(
		app_msg.M("recipe.dev.doc.report.header.name"),
		app_msg.M("recipe.dev.doc.report.header.description"),
	)

	cols := rs.Columns()
	for _, col := range cols {
		mt.Row(
			app_msg.Raw(col),
			rs.ColumnDesc(col),
		)
	}

	mt.Flush()
	w.Flush()
	return b.String()
}

func (z *Commands) feedTable(spec fd_file.Spec) string {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	mc := z.ctl.Messages()

	mui := app_ui.NewMarkdown(mc, w, false)
	mt := mui.InfoTable(spec.Name())

	mt.Header(
		app_msg.M("recipe.dev.doc.feed.header.name"),
		app_msg.M("recipe.dev.doc.feed.header.description"),
		app_msg.M("recipe.dev.doc.feed.header.example"),
	)

	cols := spec.Columns()
	for _, col := range cols {
		mt.Row(
			app_msg.Raw(col),
			spec.ColumnDesc(col),
			spec.ColumnExample(col),
		)
	}

	mt.Flush()
	w.Flush()
	return b.String()
}

func (z *Commands) feedSample(spec fd_file.Spec) string {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	cw := csv.NewWriter(w)
	ui := z.ctl.UI()

	cols := spec.Columns()
	cw.Write(spec.Columns())

	exRow := make([]string, 0)
	for _, col := range cols {
		exRow = append(exRow, ui.Text(spec.ColumnExample(col).Key()))
	}
	cw.Write(exRow)
	cw.Flush()
	w.Flush()
	return b.String()
}

func (z *Commands) Generate(r rc_recipe.Recipe) error {
	spec := rc_spec.New(r)
	if spec == nil {
		return errors.New("no spec defined for the recipe")
	}

	l := z.ctl.Log()
	ui := z.ctl.UI()

	l.Info("Generating command manual", zap.String("command", spec.CliPath()))

	tmplBytes, err := z.ctl.Resource("command.tmpl.md")
	if err != nil {
		l.Error("Template not found", zap.Error(err))
		return err
	}
	tmpl, err := template.New(spec.CliPath()).Funcs(msgFuncMap(z.ctl, z.toStdout)).Parse(string(tmplBytes))
	if err != nil {
		l.Error("Unable to compile template", zap.Error(err))
		return err
	}
	commonSpec, _, _ := rc_spec.NewCommonValue()

	params := make(map[string]interface{})
	params["Command"] = spec.CliPath()
	params["CommandTitle"] = ui.Text(spec.Title().Key())
	params["CommandDesc"] = ui.TextOrEmpty(spec.Desc().Key())
	params["CommandArgs"] = ui.TextOrEmpty(spec.CliArgs().Key())
	params["CommandNote"] = ui.TextOrEmpty(spec.CliNote().Key())
	params["Options"] = z.optionsTable(spec)
	params["CommonOptions"] = z.optionsTable(commonSpec)
	params["UseAuth"] = len(spec.ConnScopes()) > 0
	params["UseAuthPersonal"] = spec.ConnUsePersonal()
	params["UseAuthBusiness"] = spec.ConnUseBusiness()
	params["AuthScopes"] = spec.ConnScopes()

	feedNames := make([]string, 0)
	feeds := make(map[string]string, 0)
	feedSamples := make(map[string]string, 0)
	for _, fd := range spec.Feeds() {
		feedNames = append(feedNames, fd.Name())
		feeds[fd.Name()] = z.feedTable(fd)
		feedSamples[fd.Name()] = z.feedSample(fd)
	}
	sort.Strings(feedNames)
	params["FeedNames"] = feedNames
	params["Feeds"] = feeds
	params["FeedSamples"] = feedSamples
	params["FeedAvailable"] = len(feedNames) > 0

	reportNames := make([]string, 0)
	reports := make(map[string]string, 0)
	for _, rs := range spec.Reports() {
		reportNames = append(reportNames, rs.Name())
		reports[rs.Name()] = z.reportTable(rs)
	}
	sort.Strings(reportNames)
	params["ReportNames"] = reportNames
	params["Reports"] = reports
	params["ReportAvailable"] = len(reportNames) > 0

	out := os.Stdout
	if !z.toStdout {
		outPath := z.path + strings.ReplaceAll(spec.CliPath(), " ", "-") + ".md"
		out, err = os.Create(outPath)
		if err != nil {
			l.Error("Unable to create file", zap.Error(err), zap.String("outPath", outPath))
			return err
		}
	}
	return tmpl.Execute(NewRemoveRedundantLinesWriter(out), params)
}

func (z *Commands) GenerateAll() error {
	cl := z.ctl.(app_control_launcher.ControlLauncher)
	recipes := cl.Catalogue()
	l := z.ctl.Log()

	numSecret := 0
	numSelfContained := 0
	numSideCar := 0

	for _, r := range recipes {
		if _, ok := r.(rc_recipe.SecretRecipe); ok {
			numSecret++
			continue
		}
		if _, ok := r.(rc_recipe.SelfContainedRecipe); ok {
			numSelfContained++
		}
		if _, ok := r.(rc_recipe.SideCarRecipe); ok {
			numSideCar++
		}
		if err := z.Generate(r); err != nil {
			return err
		}
	}
	l.Info("Recipes", zap.Int("SecretRecipes", numSecret), zap.Int("SideCarRecipe", numSideCar), zap.Int("SelfContainedRecipe", numSelfContained))
	return nil
}
