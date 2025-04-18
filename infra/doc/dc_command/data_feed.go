package dc_command

import (
	"encoding/csv"

	"github.com/watermint/toolbox/essentials/encoding/es_csv"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewFeed(spec rc_recipe.Spec) dc_section.Section {
	return &Feed{
		spec: spec,
	}
}

type Feed struct {
	spec               rc_recipe.Spec
	Header             app_msg.Message
	HeaderFeed         app_msg.Message
	TableHeaderName    app_msg.Message
	TableHeaderDesc    app_msg.Message
	TableHeaderExample app_msg.Message
	AboutFormat        app_msg.Message
}

func (z Feed) Title() app_msg.Message {
	return z.Header
}

func (z Feed) Body(ui app_ui.UI) {
	feeds := z.spec.Feeds()
	for _, fs := range feeds {
		z.bodyFeed(ui, fs)
	}
}

func (z Feed) bodyFeed(ui app_ui.UI, fs fd_file.Spec) {
	ui.SubHeader(z.HeaderFeed.With("Feed", fs.Name()))
	ui.Break()
	ui.Info(fs.Desc())

	ui.WithTable(fs.Name(), func(t app_ui.Table) {
		t.Header(z.TableHeaderName, z.TableHeaderDesc, z.TableHeaderExample)

		cols := fs.Columns()
		for _, col := range cols {
			t.Row(app_msg.Raw(col), fs.ColumnDesc(col), fs.ColumnExample(col))
		}
	})
	ui.Break()

	sample := es_csv.MakeCsv(func(w *csv.Writer) {
		cols := fs.Columns()
		vals := make([]string, len(cols))

		for i, col := range cols {
			vals[i] = ui.Text(fs.ColumnExample(col))
		}

		_ = w.Write(cols)
		_ = w.Write(vals)
	})

	ui.Info(z.AboutFormat)
	ui.Code(sample)
}
