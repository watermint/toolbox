package app_recipe

import (
	"github.com/watermint/toolbox/experimental/app_kitchen"
	"github.com/watermint/toolbox/experimental/app_vo"
)

type Recipe interface {
	Requirement() app_vo.ValueObject
	Exec(k app_kitchen.Kitchen) error
}

// SecretRecipe will not be listed in available commands.
type SecretRecipe interface {
	Hidden()
}

// Console only recipe will not be listed in web console.
type ConsoleRecipe interface {
	Console()
}
