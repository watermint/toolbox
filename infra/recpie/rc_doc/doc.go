package rc_doc

import (
	"github.com/watermint/toolbox/infra/recpie/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func ReportSpec(ui app_ui.UI, r rc_recipe.Recipe) {
	switch scr := r.(type) {
	case rc_recipe.SideCarRecipe:
		specs := scr.Reports()
		if len(specs) < 1 {
			return
		}

		ui.Header("report.recipe.head")

		for _, spec := range specs {
			ui.Break()
			ui.Header("report.recipe.head_report", app_msg.P{"Name": spec.Name()})

			cols := spec.Columns()
			t := ui.InfoTable(spec.Name())

			t.Header(
				app_msg.M("report.recipe.col_head.name"),
				app_msg.M("report.recipe.col_head.desc"),
			)
			for _, col := range cols {
				t.Row(
					app_msg.M("raw", app_msg.P{"Raw": col}),
					spec.ColumnDesc(col),
				)
			}
			t.Flush()
			ui.Break()
		}
	}
}
