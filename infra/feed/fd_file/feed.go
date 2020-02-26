package fd_file

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_doc"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type Rows interface {
	EachRow(exec func(m interface{}, rowIndex int) error) error
}

// Row feed interface for SelfContainedRecipe
type RowFeed interface {
	Rows
	SetModel(m interface{})
	SetFilePath(filePath string)
	FilePath() string

	Model() interface{}
	Open(ctl app_control.Control) error
	Spec() Spec
}

type Spec interface {
	Name() string
	Desc() app_msg.Message
	Columns() []string
	ColumnDesc(col string) app_msg.Message
	ColumnExample(col string) app_msg.Message

	// Generate spec doc of the feed
	Doc(ui app_ui.UI) *rc_doc.Feed
}
