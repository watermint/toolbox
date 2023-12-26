package catalogue

import (
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"testing"
)

func TestAutoDetectedRecipesCitron(t *testing.T) {
	ad := AutoDetectedRecipesCitron()
	for _, a := range ad {
		s := rc_spec.New(a)
		s.Name()
	}
}
