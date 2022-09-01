package rc_recipe

import "github.com/watermint/toolbox/infra/control/app_control"

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

	// True if the operation is transient.
	IsTransient() bool

	// IsDeprecated returns true if the operation is deprecated.
	IsDeprecated() bool
}

func NewAnnotated(seed Recipe) Annotation {
	return &AnnotatedRecipe{
		seed: seed,
	}
}

type AnnotatedRecipe struct {
	seed Recipe
}

func (z AnnotatedRecipe) Preset() {
	z.seed.Preset()
}

func (z AnnotatedRecipe) Exec(c app_control.Control) error {
	return z.seed.Exec(c)
}

func (z AnnotatedRecipe) Test(c app_control.Control) error {
	return z.seed.Test(c)
}

func (z AnnotatedRecipe) Seed() Recipe {
	return z.seed
}

func (z AnnotatedRecipe) IsDeprecated() bool {
	if r, ok := z.seed.(RemarkRecipeDeprecated); ok {
		return r.IsDeprecated()
	}
	return false
}

func (z AnnotatedRecipe) IsSecret() bool {
	if r, ok := z.seed.(RemarkRecipeSecret); ok {
		return r.IsSecret()
	}
	return false
}

func (z AnnotatedRecipe) IsConsole() bool {
	if r, ok := z.seed.(RemarkRecipeConsole); ok {
		return r.IsConsole()
	}
	return false
}

func (z AnnotatedRecipe) IsExperimental() bool {
	if r, ok := z.seed.(RemarkRecipeExperimental); ok {
		return r.IsExperimental()
	}
	return false
}

func (z AnnotatedRecipe) IsIrreversible() bool {
	if r, ok := z.seed.(RemarkRecipeIrreversible); ok {
		return r.IsIrreversible()
	}
	return false
}

func (z AnnotatedRecipe) IsTransient() bool {
	if r, ok := z.seed.(RemarkRecipeTransient); ok {
		return r.IsTransient()
	}
	return false
}
