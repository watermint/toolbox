package rc_recipe

import (
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/rc_kitchen"
	"github.com/watermint/toolbox/infra/recpie/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

const (
	BasePackage = app.Pkg + "/recipe"
)

type Recipe interface {
	Exec(k rc_kitchen.Kitchen) error
	Test(c app_control.Control) error
}

type SideCarRecipe interface {
	Recipe
	Requirement() rc_vo.ValueObject
	Reports() []rp_spec.ReportSpec
}

type SelfContainedRecipe interface {
	Recipe
	Init()
}

// SecretRecipe will not be listed in available commands.
type SecretRecipe interface {
	Hidden()
}

// Console only recipe will not be listed in web console.
type ConsoleRecipe interface {
	Console()
}

type SpecValue interface {
	// Array of value names
	ValueNames() []string

	// Value description for the name
	ValueDesc(name string) app_msg.Message

	// Value default for the name
	ValueDefault(name string) interface{}

	// Customized value default for the name
	ValueCustomDefault(name string) app_msg.MessageOptional
}

type Spec interface {
	SpecValue

	// Recipe name
	Name() string

	// Recipe title
	Title() app_msg.Message

	// Recipe description
	Desc() app_msg.MessageOptional

	// Recipe path on cli
	CliPath() string

	// Recipe argument on cli
	CliArgs() app_msg.MessageOptional

	// Notes for the recipe on cli
	CliNote() app_msg.MessageOptional

	// Spec of reports generated by this recipe
	Reports() []rp_spec.ReportSpec

	// True if this recipe use connection to the Dropbox Personal account
	ConnUsePersonal() bool

	// True if this recipe use connection to the Dropbox Business account
	ConnUseBusiness() bool

	// Returns array of scope of connections to Dropbox account(s)
	ConnScopes() []string
}