package ut_doc

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_group"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_doc"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"go.uber.org/zap"
	"os"
	"sort"
	"strings"
	"text/template"
)

func NewCommand() *Commands {
	return &Commands{}
}

func NewCommandWithPath(path string, toStdout bool) *Commands {
	return &Commands{
		toStdout: toStdout,
		path:     path,
	}
}

type Commands struct {
	toStdout bool
	path     string
}

func (z *Commands) optionsTable(ctl app_control.Control, spec rc_recipe.SpecValue) string {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	mc := ctl.Messages()

	mui := app_ui.NewMarkdown(mc, w, false)
	app_doc.PrintOptionsTable(mui, spec)
	w.Flush()
	return b.String()
}

func (z *Commands) reportTable(ctl app_control.Control, rs rp_model.Spec) string {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	mc := ctl.Messages()

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

func (z *Commands) feedTable(ctl app_control.Control, spec fd_file.Spec) string {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	mc := ctl.Messages()

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

func (z *Commands) feedSample(ctl app_control.Control, spec fd_file.Spec) string {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	cw := csv.NewWriter(w)
	ui := ctl.UI()

	cols := spec.Columns()
	cw.Write(spec.Columns())

	exRow := make([]string, 0)
	for _, col := range cols {
		exRow = append(exRow, ui.Text(spec.ColumnExample(col)))
	}
	cw.Write(exRow)
	cw.Flush()
	w.Flush()
	return b.String()
}

func (z *Commands) Generate(ctl app_control.Control, r rc_recipe.Recipe) error {
	spec := rc_spec.New(r)

	l := ctl.Log()
	ui := ctl.UI()

	l.Info("Generating command manual", zap.String("command", spec.CliPath()))

	tmplBytes, err := ctl.Resource("command.tmpl.md")
	if err != nil {
		l.Error("Template not found", zap.Error(err))
		return err
	}
	tmpl, err := template.New(spec.CliPath()).Funcs(msgFuncMap(ctl, z.toStdout)).Parse(string(tmplBytes))
	if err != nil {
		l.Error("Unable to compile template", zap.Error(err))
		return err
	}
	commonSpec := rc_spec.NewCommonValue()

	authExample := ""
	{
		var b bytes.Buffer
		w := bufio.NewWriter(&b)
		cui := app_ui.NewBufferConsole(ctl.Messages(), w)
		rc_group.AppHeader(cui, "xx.x.xxx")
		cui.Info(api_auth_impl.MCcAuth.OauthSeq1.With("Url", "https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx"))
		cui.Info(api_auth_impl.MCcAuth.OauthSeq2)
		w.Flush()
		authExample = "```\n" + b.String() + "\n```"
	}

	params := make(map[string]interface{})
	params["Command"] = spec.CliPath()
	params["CommandTitle"] = ui.Text(spec.Title())
	params["CommandRemarks"] = ui.TextOrEmpty(spec.Remarks())
	params["CommandDesc"] = ui.TextOrEmpty(spec.Desc())
	params["CommandArgs"] = ui.TextOrEmpty(spec.CliArgs())
	params["CommandNote"] = ui.TextOrEmpty(spec.CliNote())
	params["Options"] = z.optionsTable(ctl, spec)
	params["CommonOptions"] = z.optionsTable(ctl, commonSpec)
	params["UseAuth"] = len(spec.ConnScopes()) > 0
	params["UseAuthPersonal"] = spec.ConnUsePersonal()
	params["UseAuthBusiness"] = spec.ConnUseBusiness()
	params["AuthScopes"] = spec.ConnScopes()
	params["AuthExample"] = authExample

	feedNames := make([]string, 0)
	feeds := make(map[string]string, 0)
	feedSamples := make(map[string]string, 0)
	for _, fd := range spec.Feeds() {
		feedNames = append(feedNames, fd.Name())
		feeds[fd.Name()] = z.feedTable(ctl, fd)
		feedSamples[fd.Name()] = z.feedSample(ctl, fd)
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
		reports[rs.Name()] = z.reportTable(ctl, rs)
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

func (z *Commands) GenerateAll(ctl app_control.Control) error {
	cl := ctl.(app_control_launcher.ControlLauncher)
	recipes := cl.Catalogue().Recipes()
	l := ctl.Log()

	numSecret := 0

	for _, r := range recipes {
		rs := rc_spec.New(r)
		if rs.IsSecret() {
			numSecret++
			continue
		}
		if err := z.Generate(ctl, r); err != nil {
			return err
		}
	}
	l.Info("Recipes", zap.Int("SecretRecipes", numSecret), zap.Int("Recipes", len(recipes)))
	return nil
}
