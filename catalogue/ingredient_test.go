package catalogue

import (
	"github.com/google/go-cmp/cmp"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"sort"
	"testing"
)

func TestAutoDetectedIngredients(t *testing.T) {
	manualCommands := make([]string, 0)
	manual := Ingredients()
	autoDetectCommands := make([]string, 0)
	autoDetect := AutoDetectedIngredients()

	for _, r := range manual {
		spec := rc_spec.New(r)
		manualCommands = append(manualCommands, spec.CliPath())
	}
	for _, ig := range autoDetect {
		spec := rc_spec.New(ig)
		autoDetectCommands = append(autoDetectCommands, spec.CliPath())
	}
	sort.Strings(manualCommands)
	sort.Strings(autoDetectCommands)
	d := cmp.Diff(manualCommands, autoDetectCommands)
	if d != "" {
		t.Error(d)
	}
}
