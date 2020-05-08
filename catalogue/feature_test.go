package catalogue

import (
	"github.com/watermint/toolbox/infra/control/app_feature"
	"testing"
)

func TestAutoDetectedFeatures(t *testing.T) {
	fs := AutoDetectedFeatures()
	for _, f := range fs {
		app_feature.OptInName(f)
	}
}
