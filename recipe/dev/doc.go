package dev

import (
	"errors"
	rice "github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recpie/app_doc"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_recipe"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/legacy/app"
	"github.com/watermint/toolbox/legacy/cmd/cmd_root"
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

func (z *Doc) Exec(k app_kitchen.Kitchen) error {
	book := make(map[string]string)

	// Loading legacy commands
	bx := rice.MustFindBox("../../legacy/resources")
	ec, err := app.NewExecContext(bx)
	if err != nil {
		return errors.New("unable to load legacy resources")
	}

	legacyRoot := cmd_root.NewCommands()
	legacy := app_doc.LegacyCommands(legacyRoot.RootCommand(), ec)

	for k, v := range legacy {
		book[k] = v
	}

	// Loading modern commands
	cl := k.Control().(app_control_launcher.ControlLauncher)
	recpies := cl.Catalogue()

	ui := k.UI()
	for _, r := range recpies {
		if _, ok := r.(app_recipe.SecretRecipe); ok {
			continue
		}

		p, n := app_recipe.Path(r)
		p = append(p, n)
		q := strings.Join(p, " ")

		book[q] = ui.Text(app_recipe.Desc(r).Key())
	}

	app_doc.PrintMarkdown(os.Stdout, "command", "description", book)

	return nil
}

func (z *Doc) Test(c app_control.Control) error {
	return z.Exec(app_kitchen.NewKitchen(c, &app_vo.EmptyValueObject{}))
}
