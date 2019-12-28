package rc_doc

import (
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func ReportSpec(ui app_ui.UI, r rc_recipe.Recipe) {
	rcpSpec := rc_spec.New(r)
	specs := rcpSpec.Reports()

	if len(specs) < 1 {
		return
	}

	ui.HeaderK("report.recipe.head")

	for _, spec := range specs {
		ui.Break()
		ui.HeaderK("report.recipe.head_report", app_msg.P{"Name": spec.Name()})

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

func RecipeFeedSpec(ui app_ui.UI, r rc_recipe.Recipe) {
	rcpSpec := rc_spec.New(r)
	specs := rcpSpec.Feeds()

	if len(specs) < 1 {
		return
	}

	ui.HeaderK("feed.recipe.head")

	for _, spec := range specs {
		ui.Break()
		ui.HeaderK("feed.recipe.head_report", app_msg.P{"Name": spec.Name()})

		cols := spec.Columns()
		t := ui.InfoTable(spec.Name())

		t.Header(
			app_msg.M("feed.recipe.col_head.name"),
			app_msg.M("feed.recipe.col_head.desc"),
			app_msg.M("feed.recipe.col_head.example"),
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

func FeedSpec(spec fd_file.Spec, ui app_ui.UI) {
	ui.Break()
	ui.HeaderK("feed.recipe.head_report", app_msg.P{"Name": spec.Name()})

	cols := spec.Columns()
	t := ui.InfoTable(spec.Name())

	t.Header(
		app_msg.M("feed.recipe.col_head.name"),
		app_msg.M("feed.recipe.col_head.desc"),
		app_msg.M("feed.recipe.col_head.example"),
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
