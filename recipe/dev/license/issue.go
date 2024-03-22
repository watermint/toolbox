package license

import (
	"encoding/base64"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/service/sv_content"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_license"
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
	Scope                   mo_string.OptionalString
	Expiration              mo_time.TimeOptional
	LifecycleExceptionUntil mo_time.TimeOptional
	RecipesAllowed          mo_string.OptionalString
	InfoIssuedLicenseKey    app_msg.Message
	Owner                   string
	Repository              string
	Branch                  string
}

func (z *Issue) Preset() {
	z.AppName = app_definitions.Name
	z.Owner = "watermint"
	z.Repository = "toolbox-supplement"
	z.Branch = "main"
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

	lifecycleExceptionUntil := ""
	if !z.LifecycleExceptionUntil.IsZero() {
		lifecycleExceptionUntil = z.LifecycleExceptionUntil.Time().Format(time.RFC3339)
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
		esl.String("lifecycleExceptionUntil", lifecycleExceptionUntil),
		esl.Strings("recipesAllowed", recipesAllowed),
	)

	lic := app_license.NewLicense(scope).WithExpiration(expiration)
	lic.AppName = z.AppName
	if lifecycleExceptionUntil != "" {
		lic = lic.WithLifecycle(&app_license.LicenseLifecycle{
			ExceptionUntil: lifecycleExceptionUntil,
		})
	}
	if len(recipesAllowed) > 0 {
		lic = lic.WithRecipe(&app_license.LicenseRecipe{
			Allow: recipesAllowed,
		})
	}

	data, key, err := lic.Issue()
	if err != nil {
		l.Debug("Unable to issue a license", esl.Error(err))
		return err
	}

	licenseName := app_license.LicenseName(key)

	c.UI().Info(z.InfoIssuedLicenseKey.With("Key", key))

	sgh := sv_content.New(z.Peer.Client(), z.Owner, z.Repository)

	_, _, err = sgh.Put(
		"licenses/"+licenseName,
		"LICENSE:"+licenseName,
		base64.StdEncoding.EncodeToString(data),
		sv_content.Branch(z.Branch),
	)

	return nil
}

func (z *Issue) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Issue{}, func(r rc_recipe.Recipe) {
		m := r.(*Issue)
		m.Scope = mo_string.NewOptional("scope_test")
		m.Expiration = mo_time.NewOptional(time.Now().Add(24 * time.Hour))
		m.LifecycleExceptionUntil = mo_time.NewOptional(time.Now().Add(24 * time.Hour))
		m.RecipesAllowed = mo_string.NewOptional("recipe1, recipe2")
	})
}
