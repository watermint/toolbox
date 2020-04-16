package catalogue

import (
	"github.com/google/go-cmp/cmp"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"sort"
	"testing"
)

func TestAutoDetectedRecipes(t *testing.T) {
	recipes := AutoDetectedRecipes()
	for _, rc := range recipes {
		spec := rc_spec.New(rc)
		spec.Name()
	}

	manualCommands := make([]string, 0)
	manual := Recipes()
	autoDetectCommands := make([]string, 0)
	autoDetect := AutoDetectedRecipes()

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
