package rc_recipe

import (
	"flag"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/doc/dc_recipe"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

const (
	BasePackage = app.Pkg + "/recipe"
)

type Preset interface {
	Preset()
}

type Recipe interface {
	Preset
	Exec(c app_control.Control) error
	Test(c app_control.Control) error
}

type RemarkRecipeSecret interface {
	// True if the recipe is not for general usage.
	IsSecret() bool
}
type RemarkRecipeExperimental interface {
	// True if the recipe is in experimental phase.
	IsExperimental() bool
}
type RemarkRecipeConsole interface {
	// True if the recipe is console mode only.
	IsConsole() bool
}
type RemarkRecipeIrreversible interface {
	// True if the operation is irreversible.
	IsIrreversible() bool
}
type RemarkSecret struct {
}

func (z *RemarkSecret) IsSecret() bool {
	return true
}

type RemarkExperimental struct {
}

func (z *RemarkExperimental) IsExperimental() bool {
	return true
}

type RemarkIrreversible struct {
}

func (z *RemarkIrreversible) IsIrreversible() bool {
	return true
}

type RemarkConsole struct {
}

func (z *RemarkConsole) IsConsole() bool {
	return true
}

type Annotation interface {
	Recipe

	// Returns seed Recipe of this annotation.
	Seed() Recipe

	// True if the recipe is not for general usage.
	IsSecret() bool

	// True if the recipe is not designed for non-console UI.
	IsConsole() bool

	// True if the recipe is in experimental phase.
	IsExperimental() bool

	// True if the operation is irreversible.
	IsIrreversible() bool
}

func NewAnnotated(seed Recipe) Annotation {
	return &AnnotatedRecipe{
		seed: seed,
	}
}

type AnnotatedRecipe struct {
	seed Recipe
}

func (z *AnnotatedRecipe) Preset() {
	z.seed.Preset()
}

func (z *AnnotatedRecipe) Exec(c app_control.Control) error {
	return z.seed.Exec(c)
}

func (z *AnnotatedRecipe) Test(c app_control.Control) error {
	return z.seed.Test(c)
}

func (z *AnnotatedRecipe) Seed() Recipe {
	return z.seed
}

func (z *AnnotatedRecipe) IsSecret() bool {
	if r, ok := z.seed.(RemarkRecipeSecret); ok {
		return r.IsSecret()
	}
	return false
}

func (z *AnnotatedRecipe) IsConsole() bool {
	if r, ok := z.seed.(RemarkRecipeConsole); ok {
		return r.IsConsole()
	}
	return false
}

func (z *AnnotatedRecipe) IsExperimental() bool {
	if r, ok := z.seed.(RemarkRecipeExperimental); ok {
		return r.IsExperimental()
	}
	return false
}

func (z *AnnotatedRecipe) IsIrreversible() bool {
	if r, ok := z.seed.(RemarkRecipeIrreversible); ok {
		return r.IsIrreversible()
	}
	return false
}

type RecipeProperty func(ro *RecipeProps) *RecipeProps
type RecipeProps struct {
	Secret       bool
	Console      bool
	Experimental bool
	Irreversible bool
}

func Secret() RecipeProperty {
	return func(ro *RecipeProps) *RecipeProps {
		ro.Secret = true
		return ro
	}
}
func Console() RecipeProperty {
	return func(ro *RecipeProps) *RecipeProps {
		ro.Console = true
		return ro
	}
}
func Experimental() RecipeProperty {
	return func(ro *RecipeProps) *RecipeProps {
		ro.Experimental = true
		return ro
	}
}
func Irreversible() RecipeProperty {
	return func(ro *RecipeProps) *RecipeProps {
		ro.Irreversible = true
		return ro
	}
}

type annotatedRecipe struct {
	seed Recipe
	opts *RecipeProps
}

func (z *annotatedRecipe) Exec(c app_control.Control) error {
	return z.seed.Exec(c)
}

func (z *annotatedRecipe) Test(c app_control.Control) error {
	return z.seed.Test(c)
}

func (z *annotatedRecipe) Preset() {
	z.seed.Preset()
}

func (z *annotatedRecipe) Seed() Recipe {
	return z.seed
}

func (z *annotatedRecipe) IsSecret() bool {
	return z.opts.Secret
}

func (z *annotatedRecipe) IsConsole() bool {
	return z.opts.Console
}

func (z *annotatedRecipe) IsExperimental() bool {
	return z.opts.Experimental
}

func (z *annotatedRecipe) IsIrreversible() bool {
	return z.opts.Irreversible
}

func Annotate(r Recipe, props ...RecipeProperty) Annotation {
	rp := &RecipeProps{}
	for _, p := range props {
		p(rp)
	}
	return &annotatedRecipe{
		seed: r,
		opts: rp,
	}
}

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

type Spec interface {
	SpecValue

	// Recipe name
	Name() string

	// Recipe title
	Title() app_msg.Message

	// Recipe description
	Desc() app_msg.MessageOptional

	// Recipe remarks
	Remarks() app_msg.MessageOptional

	// Path signature of the recipe
	Path() (path []string, name string)

	// Recipe path on cli
	CliPath() string

	// Recipe argument on cli
	CliArgs() app_msg.MessageOptional

	// Notes for the recipe on cli
	CliNote() app_msg.MessageOptional

	// Spec of reports generated by this recipe
	Reports() []rp_model.Spec

	// Spec of feeds
	Feeds() map[string]fd_file.Spec

	// Messages used by this recipe
	Messages() []app_msg.Message

	// True if this recipe use connection to the Dropbox Personal account
	ConnUsePersonal() bool

	// True if this recipe use connection to the Dropbox Business account
	ConnUseBusiness() bool

	// Returns array of scope of connections to Dropbox account(s)
	ConnScopes() []string

	// Field name and scope label map
	ConnScopeMap() map[string]string

	// Apply values to the new recipe instance
	SpinUp(ctl app_control.Control, custom func(r Recipe)) (rcp Recipe, err error)

	// Serialize values
	Debug() map[string]interface{}

	// SpinDown
	SpinDown(ctl app_control.Control) error

	// True if the recipe is not for general usage.
	IsSecret() bool

	// True if the recipe is not designed for non-console UI.
	IsConsole() bool

	// True if the recipe is in experimental phase.
	IsExperimental() bool

	// True if the operation is irreversible.
	IsIrreversible() bool

	// Print usage
	PrintUsage(ui app_ui.UI)

	// Create new spec
	New() Spec

	// Specification document
	Doc(ui app_ui.UI) *dc_recipe.Recipe
}

func NoCustomValues(r Recipe) {}

type Nop struct {
}

func (z Nop) Preset() {
}

func (z Nop) Exec(c app_control.Control) error {
	return nil
}

func (z Nop) Test(c app_control.Control) error {
	return nil
}
