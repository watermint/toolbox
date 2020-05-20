package dc_command

import (
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"sort"
	"strings"
)

func NewReport(spec rc_recipe.Spec) dc_section.Section {
	return &Report{
		spec: spec,
	}
}

type Report struct {
	spec               rc_recipe.Spec
	Header             app_msg.Message
	FileLocation       app_msg.Message
	TableHeaderOs      app_msg.Message
	TableHeaderPath    app_msg.Message
	TableHeaderExample app_msg.Message
	HeaderReport       app_msg.Message
	ReportHeaderName   app_msg.Message
	ReportHeaderDesc   app_msg.Message
	FormatDesc         app_msg.Message
	RemarkMemoryBudget app_msg.Message
	RemarkXlsx         app_msg.Message
}

func (z Report) Title() app_msg.Message {
	return z.Header
}

func (z Report) Body(ui app_ui.UI) {
	ui.Info(z.FileLocation)
	ui.WithTable("Location", func(t app_ui.Table) {
		t.Header(z.TableHeaderOs, z.TableHeaderPath, z.TableHeaderExample)
		t.RowRaw(
			"Windows",
			"`%HOMEPATH%\\.toolbox\\jobs\\[job-id]\\reports`",
			"C:\\Users\\bob\\.toolbox\\jobs\\20190909-115959.597\\reports",
		)
		t.RowRaw(
			"macOS",
			"`$HOME/.toolbox/jobs/[job-id]/reports`",
			"/Users/bob/.toolbox/jobs/20190909-115959.597/reports",
		)
		t.RowRaw(
			"Linux",
			"`$HOME/.toolbox/jobs/[job-id]/reports`",
			"/home/bob/.toolbox/jobs/20190909-115959.597/reports",
		)
	})

	reports := z.spec.Reports()
	sort.Slice(reports, func(i, j int) bool {
		return strings.Compare(reports[i].Name(), reports[j].Name()) < 0
	})
	for _, rs := range reports {
		z.bodyReport(ui, rs)
	}
}

func (z Report) bodyReport(ui app_ui.UI, rs rp_model.Spec) {
	ui.SubHeader(z.HeaderReport.With("Report", rs.Name()))
	ui.Info(rs.Desc())
	ui.Info(z.FormatDesc.With("Name", rs.Name()))

	ui.WithTable(rs.Name(), func(t app_ui.Table) {
		t.Header(z.ReportHeaderName, z.ReportHeaderDesc)
		cols := rs.Columns()
		for _, col := range cols {
			t.Row(app_msg.Raw(col), rs.ColumnDesc(col))
		}
	})
	ui.Break()
	ui.Info(z.RemarkMemoryBudget)
	ui.Break()
	ui.Info(z.RemarkXlsx.With("Name", rs.Name()))
}
