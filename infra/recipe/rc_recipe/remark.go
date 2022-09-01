package rc_recipe

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
type RemarkRecipeTransient interface {
	// True if the operation is transient. Logs will not be managed as like regular commands.
	IsTransient() bool
}
type RemarkRecipeDeprecated interface {
	// IsDeprecated returns true if the operation is no longer supported
	IsDeprecated() bool
}

type RemarkSecret struct {
}

func (z RemarkSecret) IsSecret() bool {
	return true
}

type RemarkExperimental struct {
}

func (z RemarkExperimental) IsExperimental() bool {
	return true
}

type RemarkIrreversible struct {
}

func (z RemarkIrreversible) IsIrreversible() bool {
	return true
}

type RemarkConsole struct {
}

func (z RemarkConsole) IsConsole() bool {
	return true
}

type RemarkTransient struct {
}

func (z RemarkTransient) IsTransient() bool {
	return true
}

type RemarkDeprecated struct {
}

func (z RemarkDeprecated) IsDeprecated() bool {
	return true
}
