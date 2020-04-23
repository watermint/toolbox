package auth

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type Import struct {
	rc_recipe.RemarkSecret
	PeerName string
	EnvName  string
}

func (z *Import) Preset() {
	z.PeerName = app.PeerEndToEndTest
	z.EnvName = app.EnvNameEndToEndToken
}

func (z *Import) Exec(c app_control.Control) error {
	l := c.Log().With(zap.String("peerName", z.PeerName), zap.String("envName", z.EnvName))
	env := ut_runtime.EnvMap()
	e, ok := env[z.EnvName]
	if !ok {
		l.Info("Environment variable not found. Skip import.")
		return nil
	}
	tokens := make(map[string]*oauth2.Token)
	if err := json.Unmarshal([]byte(e), &tokens); err != nil {
		l.Debug("Unable to unmarshal", zap.Error(err))
		return err
	}

	pa := dbx_auth.NewMockWithPreset(z.PeerName, tokens)
	ca := api_auth_impl.NewConsoleCache(c, pa)

	for _, scope := range Scopes {
		if _, err := ca.Auth(scope); err != nil {
			l.Info("Skip loading", zap.String("scope", scope), zap.Error(err))
		} else {
			l.Info("Loaded", zap.String("scope", scope))
		}
	}
	l.Info("Tokens loaded")
	return nil
}

func (z *Import) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Import{}, rc_recipe.NoCustomValues)
}
