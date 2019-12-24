package fd_file

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Rows interface {
	EachRow(exec func(m interface{}, rowIndex int) error) error
}

// Row feed interface for SelfContainedRecipe
type RowFeed interface {
	Rows
	SetModel(m interface{})

	Model() interface{}
	ApplyModel(ctl app_control.Control) error
	Spec() Spec
}

type Spec interface {
	Name() string
	Desc() app_msg.Message
	Columns() []string
	ColumnDesc(col string) app_msg.Message
	ColumnExample(col string) app_msg.Message
}

// File interface for SideCarRecipe
// Deprecated: use RowFeed
type ModelFile interface {
	Rows
	Model(ctl app_control.Control, m interface{}) error
}
