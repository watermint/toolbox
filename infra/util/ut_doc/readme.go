package ut_doc

import (
	"bufio"
	"bytes"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"go.uber.org/zap"
	"os"
	"sort"
	"strings"
	"text/template"
)

func NewReadme(
	ctl app_control.Control,
	filename string,
	badge bool,
	toStdout bool,
	markdown bool,
	commandPath string,
) *Readme {
	return &Readme{
		filename:    filename,
		badge:       badge,
		toStdout:    toStdout,
		markdown:    markdown,
		ctl:         ctl,
		commandPath: commandPath,
	}
}

type Readme struct {
	filename    string
	badge       bool
	toStdout    bool
	markdown    bool
	commandPath string
	ctl         app_control.Control
}

func (z *Readme) commands() string {
	book := make(map[string]string)
	keys := make([]string, 0)
	cl := z.ctl.(app_control_launcher.ControlLauncher)
	recipes := cl.Catalogue().Recipes

	ui := z.ctl.UI()
	for _, r := range recipes {
		if _, ok := r.(rc_recipe.SecretRecipe); ok {
			continue
		}

		p, n := rc_recipe.Path(r)
		p = append(p, n)
		q := strings.Join(p, " ")

		book[q] = ui.TextK(rc_recipe.Title(r).Key())
		keys = append(keys, q)
	}
	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	mui := app_ui.NewMarkdown(z.ctl.Messages(), w, false)
	mt := mui.InfoTable("Commands")

	mt.Header(
		app_msg.M("recipe.dev.doc.commands.header.command"),
		app_msg.M("recipe.dev.doc.commands.header.description"),
	)
	sort.Strings(keys)
	for _, k := range keys {
		c := k
		if z.markdown {
			c = "[" + k + "](" + z.commandPath + strings.Replace(k, " ", "-", -1) + ".md)"
		}
		mt.RowRaw(c, book[k])
	}

	mt.Flush()
	w.Flush()

	return b.String()
}

func (z *Readme) Generate() error {
	l := z.ctl.Log()
	commands := z.commands()

	l.Info("Generating README", zap.String("file", z.filename))
	readmeBytes, err := z.ctl.Resource("README.tmpl.md")
	if err != nil {
		l.Error("Template not found", zap.Error(err))
		return err
	}

	tmpl, err := template.New("README").Funcs(msgFuncMap(z.ctl, z.toStdout)).Parse(string(readmeBytes))
	if err != nil {
		l.Error("Unable to compile template", zap.Error(err))
		return err
	}

	out := os.Stdout
	if !z.toStdout {
		out, err = os.Create(z.filename)
		if err != nil {
			return err
		}
		defer out.Close()
	}

	params := make(map[string]interface{})
	params["Commands"] = commands

	if z.badge {
		params["Badges"] = app.ProjectStatusBadge
	} else {
		params["Badges"] = ""
	}

	return tmpl.Execute(NewRemoveRedundantLinesWriter(out), params)
}
