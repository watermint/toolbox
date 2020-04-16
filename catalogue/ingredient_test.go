package catalogue

import (
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"testing"
)

func TestAutoDetectedIngredients(t *testing.T) {
	ingredients := AutoDetectedIngredients()
	for _, ig := range ingredients {
		spec := rc_spec.New(ig)
		spec.Name()
	}
}
