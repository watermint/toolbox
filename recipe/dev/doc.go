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
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type Doc struct {
}

func (z *Doc) Console() {
}

func (z *Doc) Hidden() {
}

func (z *Doc) Requirement() app_vo.ValueObject {
	return &app_vo.EmptyValueObject{}
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

		book[q] = ui.Text(app_recipe.Desc(r).Key())
	}

	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	app_doc.PrintMarkdown(w, "command", "description", book)
	w.Flush()

	return b.String()
}

func (z *Doc) Exec(k app_kitchen.Kitchen) error {
	commands := z.commands(k)

	readmeBytes, err := ioutil.ReadFile("doc/README.tmpl.md")
	if err != nil {
		return err
	}

	tmpl, err := template.New("README").Parse(string(readmeBytes))
	if err != nil {
		return err
	}
	readmeFile, err := os.Create("README.md")
	if err != nil {
		return err
	}
	defer readmeFile.Close()

	err = tmpl.Execute(readmeFile, map[string]interface{}{
		"Commands": commands,
	})
	if err != nil {
		return err
	}

	return nil
}

func (z *Doc) Test(c app_control.Control) error {
	return z.Exec(app_kitchen.NewKitchen(c, &app_vo.EmptyValueObject{}))
}
