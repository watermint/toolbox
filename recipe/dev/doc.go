package dev

import (
	"fmt"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_recipe"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
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
	cl := k.Control().(app_control_launcher.ControlLauncher)
	recpies := cl.Catalogue()
	book := make(map[string]string)
	names := make([]string, 0)

	for _, r := range recpies {
		if _, ok := r.(app_recipe.SecretRecipe); ok {
			continue
		}

		p, n := app_recipe.Path(r)
		p = append(p, n)
		q := strings.Join(p, " ")

		book[q] = k.UI().Text(app_recipe.Desc(r).Key())
		names = append(names, q)
	}
	sort.Strings(names)

	tw := new(tabwriter.Writer)
	tw.Init(os.Stdout, 0, 2, 1, ' ', tabwriter.Debug)

	row := func(col ...string) {
		cols := make([]string, 0)
		cols = append(cols, col...)
		cols = append(cols, " ")
		fmt.Fprintln(tw, "|"+strings.Join(cols, "\t"))
	}

	row("Recipe", "Description")
	for _, name := range names {
		row("`"+name+"`", book[name])
	}
	tw.Flush()

	return nil
}

func (z *Doc) Test(c app_control.Control) error {
	return z.Exec(app_kitchen.NewKitchen(c, &app_vo.EmptyValueObject{}))
}
