package catalogue

import "testing"

func TestAutoDetectedFeatures(t *testing.T) {
	fs := AutoDetectedFeatures()
	for _, f := range fs {
		f.OptInName(f)
	}
}
