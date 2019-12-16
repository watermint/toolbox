package ut_doc

import (
	"bufio"
	"bytes"
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_recipe"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/recpie/app_vo_impl"
	"github.com/watermint/toolbox/infra/report/rp_spec"
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

func (z *Commands) authScopes(vo app_vo.ValueObject) (scopes []string, usePersonal, useBusiness bool) {
	l := z.ctl.Log()
	scopes = make([]string, 0)
	sc := make(map[string]bool)

	vc := app_vo_impl.NewValueContainer(vo)
	for _, v := range vc.Values {
		switch v0 := v.(type) {
		case app_conn.ConnBusinessInfo:
			l.Debug("business info", zap.Any("v0", v0))
			sc["business_info"] = true
			useBusiness = true
		case app_conn.ConnBusinessMgmt:
			sc["business_mgmt"] = true
			useBusiness = true
		case app_conn.ConnBusinessFile:
			sc["business_file"] = true
			useBusiness = true
		case app_conn.ConnBusinessAudit:
			sc["business_audit"] = true
			useBusiness = true
		case app_conn.ConnUserFile:
			sc["user_file"] = true
			usePersonal = true
		}
	}
	for s := range sc {
		scopes = append(scopes, s)
	}
	sort.Strings(scopes)

	return scopes, usePersonal, useBusiness
}

func (z *Commands) optionsTable(vo app_vo.ValueObject) string {
	l := z.ctl.Log()
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

	vc := app_vo_impl.NewValueContainer(vo)

	if len(vc.Values) < 1 {
		return ""
	}

	keys := make([]string, 0)
	for k := range vc.Values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		vk := vc.MessageKey(k)
		vd := vc.Values[k]
		switch v := vd.(type) {
		case fd_file.Feed:
			l.Debug("Feed file", zap.Any("v", v))
			vd = ""
		}

		vkd := vk + ".default"
		if mc.Exists(vkd) {
			vd = mc.Text(vkd)
		}
		mt.Row(
			app_msg.M("recipe.dev.doc.options.body.option", app_msg.P{"Option": strcase.ToKebab(k)}),
			app_msg.M(vc.MessageKey(k)),
			app_msg.M("raw", app_msg.P{"Raw": vd}),
		)
	}

	mt.Flush()
	w.Flush()
	return b.String()
}

func (z *Commands) reportTable(rs rp_spec.ReportSpec) string {
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

func (z *Commands) Generate(r app_recipe.SideCarRecipe) error {
	l := z.ctl.Log()
	ui := z.ctl.UI()
	mc := z.ctl.Messages()

	path, name := app_recipe.Path(r)
	path = append(path, name)
	command := strings.Join(path, " ")

	l.Info("Generating command manual", zap.String("command", command))

	msgOrEmpty := func(m app_msg.Message) string {
		if mc.Exists(m.Key()) {
			return mc.Compile(m)
		}
		return ""
	}

	tmplBytes, err := z.ctl.Resource("command.tmpl.md")
	if err != nil {
		l.Error("Template not found", zap.Error(err))
		return err
	}
	tmpl, err := template.New(command).Funcs(msgFuncMap(z.ctl, z.toStdout)).Parse(string(tmplBytes))
	if err != nil {
		l.Error("Unable to compile template", zap.Error(err))
		return err
	}

	authScopes, usePersonal, useBusiness := z.authScopes(r.Requirement())

	params := make(map[string]interface{})
	params["Command"] = command
	params["CommandTitle"] = ui.Text(app_recipe.Title(r).Key())
	params["CommandDesc"] = msgOrEmpty(app_recipe.Desc(r))
	params["CommandArgs"] = msgOrEmpty(app_recipe.RecipeMessage(r, "cli.args"))
	params["CommandNote"] = msgOrEmpty(app_recipe.RecipeMessage(r, "cli.note"))
	params["Options"] = z.optionsTable(r.Requirement())
	params["CommonOptions"] = z.optionsTable(app_opt.NewDefaultCommonOpts())
	params["UseAuth"] = len(authScopes) > 0
	params["UseAuthPersonal"] = usePersonal
	params["UseAuthBusiness"] = useBusiness
	params["AuthScopes"] = authScopes

	reportNames := make([]string, 0)
	reports := make(map[string]string, 0)
	for _, rs := range r.Reports() {
		reportNames = append(reportNames, rs.Name())
		reports[rs.Name()] = z.reportTable(rs)
	}
	sort.Strings(reportNames)
	params["ReportNames"] = reportNames
	params["Reports"] = reports
	params["ReportAvailable"] = len(reportNames) > 0

	out := os.Stdout
	if !z.toStdout {
		outPath := z.path + strings.Join(path, "-") + ".md"
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

	for _, r := range recipes {
		if _, ok := r.(app_recipe.SecretRecipe); ok {
			continue
		}

		switch re := r.(type) {
		case app_recipe.SideCarRecipe:
			if err := z.Generate(re); err != nil {
				return err
			}
		}
	}
	return nil
}
