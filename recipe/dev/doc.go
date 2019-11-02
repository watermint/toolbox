package dev

import (
	"bufio"
	"bytes"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recpie/app_doc"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_recipe"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"go.uber.org/zap"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

const (
	badges = `
[![CircleCI](https://circleci.com/gh/watermint/toolbox.svg?style=svg)](https://circleci.com/gh/watermint/toolbox)
[![Coverage Status](https://coveralls.io/repos/github/watermint/toolbox/badge.svg)](https://coveralls.io/github/watermint/toolbox)
[![Go Report Card](https://goreportcard.com/badge/github.com/watermint/toolbox)](https://goreportcard.com/report/github.com/watermint/toolbox)
`
)

type DocVO struct {
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
		Badge:    true,
		Filename: "README.md",
	}
}

func (z *Doc) commands(k app_kitchen.Kitchen) string {
	book := make(map[string]string)
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
	}

	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	app_doc.PrintMarkdown(w, "command", "description", book)
	w.Flush()

	return b.String()
}

func (z *Doc) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*DocVO)
	l := k.Log()
	commands := z.commands(k)

	l.Info("Generating README", zap.String("file", vo.Filename))
	readmeBytes, err := ioutil.ReadFile("doc/README.tmpl.md")
	if err != nil {
		return err
	}

	tmpl, err := template.New("README").Parse(string(readmeBytes))
	if err != nil {
		return err
	}
	readmeFile, err := os.Create(vo.Filename)
	if err != nil {
		return err
	}
	defer readmeFile.Close()

	params := make(map[string]interface{})
	params["Commands"] = commands

	if vo.Badge {
		params["Badges"] = badges
	} else {
		params["Badges"] = ""
	}

	err = tmpl.Execute(readmeFile, params)
	if err != nil {
		return err
	}

	return nil
}

func (z *Doc) Test(c app_control.Control) error {
	return z.Exec(app_kitchen.NewKitchen(c, &app_vo.EmptyValueObject{}))
}
