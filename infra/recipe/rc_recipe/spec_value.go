package rc_recipe

import (
	"flag"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type SpecValue interface {
	// Array of value names
	ValueNames() []string

	// Value description for the name
	ValueDesc(name string) app_msg.Message

	// Value default for the name
	ValueDefault(name string) interface{}

	// Value for the name
	Value(name string) Value

	// Customized value default for the name
	ValueCustomDefault(name string) app_msg.MessageOptional

	// Configure CLI flags
	SetFlags(f *flag.FlagSet, ui app_ui.UI)
}
