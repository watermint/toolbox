package license

import (
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_license"
	"github.com/watermint/toolbox/infra/control/app_license_key"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"strings"
	"time"
)

type KeySummary struct {
	Key             string `json:"key"`
	ExpirationDate  string `json:"expiration_date"`
	Valid           bool   `json:"valid"`
	LicenseeName    string `json:"licensee_name"`
	LicenseeEmail   string `json:"licensee_email"`
	LicensedRecipes string `json:"licensed_recipes"`
}

type List struct {
	Keys rp_model.RowReport
}

func (z *List) Preset() {
	z.Keys.SetModel(&KeySummary{})
}

func (z *List) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.Keys.Open(); err != nil {
		return err
	}
	keys := app_license_key.AvailableKeys(c.Workspace())
	for _, key := range keys {
		lic, err := app_license.LoadAndCacheLicense(key, app_definitions.SupplementRepositoryLicenseUrl, c.Workspace().Secrets())
		if errors.Is(err, app_license.ErrorUnknownLicenseType) {
			l.Debug("Unknown license type, potentially the key is for different tool", esl.String("key", key))
			continue
		}
		if err != nil {
			l.Warn("Unable to load license", esl.String("key", key), esl.Error(err))
			continue
		}
		recipes := make([]string, 0)
		if lic.Recipe != nil {
			copy(recipes, lic.Recipe.Allow)
		}
		z.Keys.Row(&KeySummary{
			Key:             key,
			ExpirationDate:  lic.LifecycleLimit().Format(time.RFC3339),
			Valid:           lic.IsValid(),
			LicenseeName:    lic.LicenseeName,
			LicenseeEmail:   lic.LicenseeEmail,
			LicensedRecipes: strings.Join(recipes, ","),
		})
	}

	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
