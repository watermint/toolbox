package dc_license

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestLoadLicenses(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		qt_file.TestWithTestFolder(t, "license", false, func(path string) {
			if err := MakeTestData(path); err != nil {
				t.Error(err)
				return
			}

			licenses, err := LoadLicenses(ctl, path)
			if err != nil {
				t.Error(err)
				return
			}

			if len(licenses.ThirdParty) != 3 {
				t.Error(es_json.ToJsonString(licenses))
				return
			}
			if licenses.Project.LicenseType != "MIT" {
				t.Error(es_json.ToJsonString(licenses.Project))
			}
		})
	})
}
