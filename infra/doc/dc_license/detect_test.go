package dc_license

import (
    "testing"

    "github.com/watermint/toolbox/essentials/log/esl"
)

func TestDetect(t *testing.T) {
	inventory, err := Detect()
	if err != nil {
		t.Error(err)
		return
	}

	l := esl.Default()

	for _, info := range inventory {
		l.Info("Library", esl.String("name", info.Package), esl.String("type", info.LicenseType))
	}
}
