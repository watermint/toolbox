package catalogue

import (
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"testing"
)

func TestAutoDetectedRecipes(t *testing.T) {
	ad := AutoDetectedRecipes()
	for _, a := range ad {
		s := rc_spec.New(a)
		s.Name()
	}
}
