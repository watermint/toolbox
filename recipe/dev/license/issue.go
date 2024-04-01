package license

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_license"
	"github.com/watermint/toolbox/infra/control/app_license_registry"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"strings"
	"time"
)

type Issue struct {
	rc_recipe.RemarkSecret
	Peer                    gh_conn.ConnGithubRepo
	AppName                 string
	Branch                  string
	Expiration              mo_time.TimeOptional
	InfoIssuedLicenseKey    app_msg.Message
	LicenseeEmail           string
	LicenseeName            string
	LifecycleAvailableAfter int64
	LifecycleWarningAfter   int64
	Owner                   string
	RecipesAllowed          mo_string.OptionalString
	Repository              string
	Scope                   mo_string.OptionalString
}

func (z *Issue) Preset() {
	z.AppName = app_definitions.Name
	z.Owner = "watermint"
	z.Repository = "toolbox-supplement"
	z.Branch = "main"
	z.LifecycleAvailableAfter = int64(app_license.DefaultLifecyclePeriod / time.Second)
	z.LifecycleWarningAfter = int64(app_license.DefaultWarningPeriod(time.Duration(z.LifecycleAvailableAfter)*time.Second) / time.Second)
}

func (z *Issue) Exec(c app_control.Control) error {
	l := c.Log()
	scope := app_license.LicenseScopeBase
	if z.Scope.IsExists() {
		scope = z.Scope.Value()
	}

	// Default expiration is 1 year
	expiration := time.Now().Add(365 * 24 * time.Hour)
	if !z.Expiration.IsZero() {
		expiration = z.Expiration.Time()
	}

	var lifecycleAvailableAfter int64
	if z.LifecycleAvailableAfter > 0 {
		lifecycleAvailableAfter = z.LifecycleAvailableAfter
	}
	var lifecycleWarningAfter int64
	if z.LifecycleWarningAfter > 0 {
		lifecycleWarningAfter = z.LifecycleWarningAfter
	}

	recipesAllowed := make([]string, 0)
	if z.RecipesAllowed.IsExists() {
		recipes := strings.Split(z.RecipesAllowed.Value(), ",")
		for _, recipe := range recipes {
			recipesAllowed = append(recipesAllowed, strings.TrimSpace(recipe))
		}
	}

	l = l.With(
		esl.String("scope", scope),
		esl.Time("expiration", expiration),
		esl.Int64("lifecycleAvailableAfter", lifecycleAvailableAfter),
		esl.Int64("lifecycleWarningAfter", lifecycleWarningAfter),
		esl.Strings("recipesAllowed", recipesAllowed),
	)

	lic := app_license.NewLicense(scope).WithExpiration(expiration)
	lic.AppName = z.AppName
	lic = lic.WithLifecycle(&app_license.LicenseLifecycle{
		AvailableAfter: lifecycleAvailableAfter,
		WarningAfter:   lifecycleWarningAfter,
	})

	if len(recipesAllowed) > 0 {
		lic = lic.WithRecipe(&app_license.LicenseRecipe{
			Allow: recipesAllowed,
		})
	}
	lic = lic.WithLicensee(z.LicenseeName, z.LicenseeEmail)

	registryPath := app_license_registry.DefaultRegistryPath(c.Workspace().Secrets())
	registryDb, err := c.NewOrm(registryPath)
	if err != nil {
		l.Debug("Unable to open the license registry", esl.Error(err))
		return err
	}

	registry := app_license_registry.NewRegistry(
		z.Peer.Client(),
		z.Owner,
		z.Repository,
		z.Branch,
		registryDb,
	)
	key, err := registry.Issue(&lic)

	c.UI().Info(z.InfoIssuedLicenseKey.With("Key", key))

	return nil
}

func (z *Issue) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Issue{}, func(r rc_recipe.Recipe) {
		m := r.(*Issue)
		m.Scope = mo_string.NewOptional("scope_test")
		m.Expiration = mo_time.NewOptional(time.Now().Add(24 * time.Hour))
		m.RecipesAllowed = mo_string.NewOptional("recipe1, recipe2")
		m.LicenseeName = "John Doe"
		m.LicenseeEmail = "john@example.com"
	})
}
