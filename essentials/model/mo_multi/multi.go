package mo_multi

import (
	"flag"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MultiValue interface {
	// Flag name
	Name() string

	// Set flags
	ApplyFlags(f *flag.FlagSet, fieldDesc app_msg.Message, ui app_ui.UI)

	// Nested value fields. (Name + NameSuffix of nested value)
	Fields() []string

	// Description of nested value field. (key: Name + NameSuffix)
	FieldDesc(base app_msg.Message, name string) app_msg.Message
}
