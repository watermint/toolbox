package dc_command

import (
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"os"
	"sort"
	"strings"
	"text/template"
)

func NewReadme(
	ctl app_control.Control,
	filename string,
	badge bool,
	markdown bool,
	commandPath string,
) *Readme {
	return &Readme{
		filename:    filename,
		badge:       badge,
		markdown:    markdown,
		ctl:         ctl,
		commandPath: commandPath,
	}
}

type Readme struct {
	filename    string
	badge       bool
	markdown    bool
	commandPath string
	ctl         app_control.Control
}

func (z *Readme) commands() string {
	book := make(map[string]string)
	keys := make([]string, 0)
	cat := app_catalogue.Current()
	recipes := cat.Recipes()

	ui := z.ctl.UI()
	for _, r := range recipes {
		rs := rc_spec.New(r)
		if rs.IsSecret() {
			continue
		}

		p, n := rs.Path()
		p = append(p, n)
		q := strings.Join(p, " ")

		z.ctl.Log().Debug("recipe", es_log.String("recipe", rs.CliPath()))
		book[q] = ui.Text(rs.Title())
		keys = append(keys, q)
	}

	return app_ui.MakeMarkdown(z.ctl.Messages(), func(mui app_ui.UI) {
		mui.WithTable("Commands", func(mt app_ui.Table) {
			mt.Header(
				MCommands.CommandHeaderCommand,
				MCommands.CommandHeaderDesc,
			)
			sort.Strings(keys)
			for _, k := range keys {
				c := k
				if z.markdown {
					c = "[" + k + "](" + z.commandPath + strings.Replace(k, " ", "-", -1) + ".md)"
				}
				mt.RowRaw(c, book[k])
			}
		})
	})
}

func (z *Readme) Generate() error {
	l := z.ctl.Log()
	commands := z.commands()

	l.Info("Generating README", es_log.String("file", z.filename))
	rb := app_resource.Bundle()
	readmeBytes, err := rb.Templates().Bytes("README.tmpl.md")
	if err != nil {
		l.Error("Template not found", es_log.Error(err))
		return err
	}

	tmpl, err := template.New("README").Funcs(msgFuncMap(z.ctl)).Parse(string(readmeBytes))
	if err != nil {
		l.Error("Unable to compile template", es_log.Error(err))
		return err
	}

	bodyUsage := app_ui.MakeConsoleDemo(z.ctl.Messages(), func(cui app_ui.UI) {
		cat := app_catalogue.Current()
		rg := cat.RootGroup()
		rg.PrintUsage(cui, "./tbx", "xx.x.xxx")
	})

	out := es_stdout.NewDefaultOut(z.ctl.Feature().IsTest())
	if !z.ctl.Feature().IsTest() {
		out, err = os.Create(z.filename)
		if err != nil {
			return err
		}
		defer out.Close()
	}

	params := make(map[string]interface{})
	params["Commands"] = commands
	params["BodyUsage"] = bodyUsage

	if z.badge {
		params["Badges"] = app.ProjectStatusBadge
		params["Logo"] = app.ProjectLogo
		params["Release"] = true
	} else {
		params["Badges"] = ""
		params["Logo"] = ""
		params["Release"] = false
	}

	return tmpl.Execute(NewRemoveRedundantLinesWriter(out), params)
}
