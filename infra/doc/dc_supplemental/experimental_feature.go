package dc_supplemental

import (
	"github.com/watermint/toolbox/essentials/strings/es_case"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MsgExperimentalFeature struct {
	DocDesc     app_msg.Message
	Title       app_msg.Message
	Overview    app_msg.Message
	FeatureName app_msg.Message
	FeatureDesc app_msg.Message
}

var (
	MExperimentalFeature = app_msg.Apply(&MsgExperimentalFeature{}).(*MsgExperimentalFeature)
)

type ExperimentalFeature struct {
}

func (z ExperimentalFeature) DocDesc() app_msg.Message {
	return MExperimentalFeature.DocDesc
}

func (z ExperimentalFeature) DocId() dc_index.DocId {
	return dc_index.DocSupplementalExperimentalFeature
}

func (z ExperimentalFeature) Sections() []dc_section.Section {
	return []dc_section.Section{
		&ExperimentalFeatureDefinitions{},
	}
}

type ExperimentalFeatureDefinitions struct {
}

func (z ExperimentalFeatureDefinitions) Title() app_msg.Message {
	return MExperimentalFeature.Title
}

func (z ExperimentalFeatureDefinitions) Body(ui app_ui.UI) {
	ui.Info(MExperimentalFeature.Overview)
	vt := ui.InfoTable("Experimental features")
	vt.Header(MExperimentalFeature.FeatureName, MExperimentalFeature.FeatureDesc)

	for _, v := range app_definitions.ExperimentalFeatures {
		vt.Row(
			app_msg.Raw(v),
			app_msg.ObjMessage(&z, es_case.ToLowerSnakeCase(v)+".desc"),
		)
	}

	vt.Flush()
}
