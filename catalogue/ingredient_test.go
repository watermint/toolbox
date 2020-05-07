package catalogue

import (
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"testing"
)

func TestAutoDetectedIngredients(t *testing.T) {
	ad := AutoDetectedIngredients()
	for _, a := range ad {
		s := rc_spec.New(a)
		s.Name()
	}
}
