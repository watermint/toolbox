package dc_supplemental

import (
	"github.com/watermint/essentials/estring/ecase"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MsgPathVariable struct {
	DocDesc     app_msg.Message
	Title       app_msg.Message
	Overview    app_msg.Message
	VarVariable app_msg.Message
	VarDesc     app_msg.Message
}

var (
	MPathVariable = app_msg.Apply(&MsgPathVariable{}).(*MsgPathVariable)
)

type PathVariable struct {
}

func (z PathVariable) DocDesc() app_msg.Message {
	return MPathVariable.DocDesc
}

func (z PathVariable) DocId() dc_index.DocId {
	return dc_index.DocSupplementalPathVariables
}

func (z PathVariable) Sections() []dc_section.Section {
	return []dc_section.Section{
		&PathVariableDefinitions{},
	}
}

type PathVariableDefinitions struct {
}

func (z PathVariableDefinitions) Title() app_msg.Message {
	return MPathVariable.Title
}

func (z PathVariableDefinitions) Body(ui app_ui.UI) {
	ui.Info(MPathVariable.Overview)
	vt := ui.InfoTable("Path Variables")
	vt.Header(MPathVariable.VarVariable, MPathVariable.VarDesc)

	for _, v := range es_filepath.PathVariables {
		vt.Row(
			app_msg.Raw("{{."+v+"}}"),
			app_msg.ObjMessage(&z, ecase.ToLowerSnakeCase(v)+".desc"),
		)
	}
	vt.Flush()
}
