package dev

import (
	"bufio"
	"bytes"
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
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

const (
	badges = `
[![CircleCI](https://circleci.com/gh/watermint/toolbox.svg?style=svg)](https://circleci.com/gh/watermint/toolbox)
[![Coverage Status](https://coveralls.io/repos/github/watermint/toolbox/badge.svg)](https://coveralls.io/github/watermint/toolbox)
[![Go Report Card](https://goreportcard.com/badge/github.com/watermint/toolbox)](https://goreportcard.com/report/github.com/watermint/toolbox)
`
)

type DocVO struct {
	Test     bool
	Badge    bool
	Filename string
}

type Doc struct {
}

func (z *Doc) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{}
}

func (z *Doc) Console() {
}

func (z *Doc) Hidden() {
}

func (z *Doc) Requirement() app_vo.ValueObject {
	return &DocVO{
		Test:     false,
		Badge:    true,
		Filename: "README.md",
	}
}

func (z *Doc) commands(k app_kitchen.Kitchen) string {
	book := make(map[string]string)
	keys := make([]string, 0)
	cl := k.Control().(app_control_launcher.ControlLauncher)
	recipes := cl.Catalogue()

	ui := k.UI()
	for _, r := range recipes {
		if _, ok := r.(app_recipe.SecretRecipe); ok {
			continue
		}

		p, n := app_recipe.Path(r)
		p = append(p, n)
		q := strings.Join(p, " ")

		book[q] = ui.Text(app_recipe.Title(r).Key())
		keys = append(keys, q)
	}
	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	mui := app_ui.NewMarkdown(k.Control().Messages(), w, false)
	mt := mui.InfoTable("Commands")
	mt.Header(
		app_msg.M("recipe.dev.doc.commands.header.command"),
		app_msg.M("recipe.dev.doc.commands.header.description"),
	)
	sort.Strings(keys)
	for _, k := range keys {
		mt.RowRaw(k, book[k])
	}

	mt.Flush()
	w.Flush()

	return b.String()
}

func (z *Doc) funcMap(k app_kitchen.Kitchen) template.FuncMap {
	vo := k.Value().(*DocVO)
	return template.FuncMap{
		"msg": func(key string) string {
			if vo.Test {
				if !k.Control().Messages().Exists(key) {
					k.UI().Error(key)
				}
			}
			return k.Control().Messages().Text(key)
		},
	}
}

func (z *Doc) readme(k app_kitchen.Kitchen) error {
	vo := k.Value().(*DocVO)
	l := k.Log()
	commands := z.commands(k)

	l.Info("Generating README", zap.String("file", vo.Filename))
	readmeBytes, err := k.Control().Resource("README.tmpl.md")
	if err != nil {
		l.Error("Template not found", zap.Error(err))
		return err
	}

	tmpl, err := template.New("README").Funcs(z.funcMap(k)).Parse(string(readmeBytes))
	if err != nil {
		l.Error("Unable to compile template", zap.Error(err))
		return err
	}

	out := os.Stdout
	if !vo.Test {
		out, err = os.Create(vo.Filename)
		if err != nil {
			return err
		}
		defer out.Close()
	}

	params := make(map[string]interface{})
	params["Commands"] = commands

	if vo.Badge {
		params["Badges"] = badges
	} else {
		params["Badges"] = ""
	}

	return tmpl.Execute(out, params)
}

func (z *Doc) optionsTable(vo app_vo.ValueObject, k app_kitchen.Kitchen) string {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	mc := k.Control().Messages()

	mui := app_ui.NewMarkdown(mc, w, false)
	mt := mui.InfoTable("")

	mt.Header(
		app_msg.M("recipe.dev.doc.options.header.option"),
		app_msg.M("recipe.dev.doc.options.header.description"),
		app_msg.M("recipe.dev.doc.options.header.default"),
	)

	vc := app_vo_impl.NewValueContainer(vo)
	keys := make([]string, 0)
	for k := range vc.Values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		vk := vc.MessageKey(k)
		vd := vc.Values[k]
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

func (z *Doc) reportTable(rs rp_spec.ReportSpec, k app_kitchen.Kitchen) string {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	mc := k.Control().Messages()

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

func (z *Doc) commandManual(r app_recipe.Recipe, k app_kitchen.Kitchen) error {
	vo := k.Value().(*DocVO)
	l := k.Log()
	ui := k.UI()
	mc := k.Control().Messages()

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

	tmplBytes, err := k.Control().Resource("command.tmpl.md")
	if err != nil {
		l.Error("Template not found", zap.Error(err))
		return err
	}
	tmpl, err := template.New(command).Funcs(z.funcMap(k)).Parse(string(tmplBytes))
	if err != nil {
		l.Error("Unable to compile template", zap.Error(err))
		return err
	}

	params := make(map[string]interface{})
	params["Command"] = command
	params["CommandTitle"] = ui.Text(app_recipe.Title(r).Key())
	params["CommandDesc"] = msgOrEmpty(app_recipe.Desc(r))
	params["CommandArgs"] = msgOrEmpty(app_recipe.RecipeMessage(r, "cli.args"))
	params["CommandNote"] = msgOrEmpty(app_recipe.RecipeMessage(r, "cli.note"))
	params["Options"] = z.optionsTable(r.Requirement(), k)
	params["CommonOptions"] = z.optionsTable(app_opt.NewDefaultCommonOpts(), k)

	reportNames := make([]string, 0)
	reports := make(map[string]string, 0)
	for _, rs := range r.Reports() {
		reportNames = append(reportNames, rs.Name())
		reports[rs.Name()] = z.reportTable(rs, k)
	}
	sort.Strings(reportNames)
	params["ReportNames"] = reportNames
	params["Reports"] = reports
	params["ReportAvailable"] = len(reportNames) > 0

	out := os.Stdout
	if !vo.Test {
		outPath := "doc/generated/" + strings.Join(path, "-") + ".md"
		out, err = os.Create(outPath)
		if err != nil {
			l.Error("Unable to create file", zap.Error(err), zap.String("outPath", outPath))
			return err
		}
	}
	return tmpl.Execute(out, params)
}

func (z *Doc) commandManuals(k app_kitchen.Kitchen) error {
	cl := k.Control().(app_control_launcher.ControlLauncher)
	recipes := cl.Catalogue()

	for _, r := range recipes {
		if _, ok := r.(app_recipe.SecretRecipe); ok {
			continue
		}

		if err := z.commandManual(r, k); err != nil {
			return err
		}
	}
	return nil
}

func (z *Doc) Exec(k app_kitchen.Kitchen) error {
	l := k.Log()
	if err := z.readme(k); err != nil {
		l.Error("Failed to generate README", zap.Error(err))
		return err
	}
	if err := z.commandManuals(k); err != nil {
		l.Error("Failed to generate command manuals", zap.Error(err))
		return err
	}

	return nil
}

func (z *Doc) Test(c app_control.Control) error {
	return z.Exec(app_kitchen.NewKitchen(c, &DocVO{
		Test:     true,
		Badge:    false,
		Filename: "",
	}))
}
