package catalogue

import (
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"testing"
)

func TestAutoDetectedRecipes(t *testing.T) {
	recipes := AutoDetectedRecipes()
	for _, rc := range recipes {
		spec := rc_spec.New(rc)
		spec.Name()
	}
}
