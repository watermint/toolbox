package dc_supplemental

import (
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MsgReportingOptions struct {
	DocDesc app_msg.Message
	Title   app_msg.Message

	OverviewTitle app_msg.Message
	OverviewDesc  app_msg.Message

	HiddenColumnTitle             app_msg.Message
	HiddenColumnDesc              app_msg.Message
	HiddenColumnCodeWithoutOption app_msg.Message
	HiddenColumnDescWithOption    app_msg.Message
	HiddenColumnCodeWithOption    app_msg.Message
	HiddenColumnConclusion        app_msg.Message

	OutputFilterTitle    app_msg.Message
	OutputFilterDesc     app_msg.Message
	OutputFilterNote     app_msg.Message
	OutputFilterSegment1 app_msg.Message
	OutputFilterSegment2 app_msg.Message
	OutputFilterCode1    app_msg.Message
	OutputFilterSegment3 app_msg.Message
	OutputFilterCode2    app_msg.Message
	OutputFilterSegment4 app_msg.Message
	OutputFilterCode3    app_msg.Message
	OutputFilterSegment5 app_msg.Message
	OutputFilterCode4    app_msg.Message
}

var (
	MReportingOptions = app_msg.Apply(&MsgReportingOptions{}).(*MsgReportingOptions)
)

type ReportingOptions struct {
}

func (z ReportingOptions) DocId() dc_index.DocId {
	return dc_index.DoCSupplementalReportingOptions
}

func (z ReportingOptions) DocDesc() app_msg.Message {
	return MReportingOptions.DocDesc
}

func (z ReportingOptions) Sections() []dc_section.Section {
	return []dc_section.Section{
		&ReportingOptionsOverview{},
		&ReportingOptionsHiddenColumns{},
		&ReportingOptionsOutputFilter{},
	}
}

type ReportingOptionsOverview struct {
}

func (z ReportingOptionsOverview) Title() app_msg.Message {
	return MReportingOptions.OverviewTitle
}

func (z ReportingOptionsOverview) Body(ui app_ui.UI) {
	ui.Info(MReportingOptions.OverviewDesc)
}

type ReportingOptionsHiddenColumns struct {
}

func (r ReportingOptionsHiddenColumns) Title() app_msg.Message {
	return MReportingOptions.HiddenColumnTitle
}

func (r ReportingOptionsHiddenColumns) Body(ui app_ui.UI) {
	ui.Info(MReportingOptions.HiddenColumnDesc)
	ui.Break()
	ui.Code(ui.Text(MReportingOptions.HiddenColumnCodeWithoutOption))
	ui.Break()
	ui.Info(MReportingOptions.HiddenColumnDescWithOption)
	ui.Break()
	ui.Code(ui.Text(MReportingOptions.HiddenColumnCodeWithOption))
	ui.Break()
	ui.Info(MReportingOptions.HiddenColumnConclusion)
}

type ReportingOptionsOutputFilter struct {
}

func (z ReportingOptionsOutputFilter) Title() app_msg.Message {
	return MReportingOptions.OutputFilterTitle
}

func (z ReportingOptionsOutputFilter) Body(ui app_ui.UI) {
	ui.Info(MReportingOptions.OutputFilterDesc)
	ui.Break()
	ui.Quote(MReportingOptions.OutputFilterNote)
	ui.Break()
	ui.Info(MReportingOptions.OutputFilterSegment1)
	ui.Break()
	ui.Info(MReportingOptions.OutputFilterSegment2)
	ui.Break()
	ui.Code(ui.Text(MReportingOptions.OutputFilterCode1))
	ui.Break()
	ui.Info(MReportingOptions.OutputFilterSegment3)
	ui.Break()
	ui.Code(ui.Text(MReportingOptions.OutputFilterCode2))
	ui.Break()
	ui.Info(MReportingOptions.OutputFilterSegment4)

}
