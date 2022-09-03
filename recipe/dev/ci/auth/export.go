package auth

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_oauth"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
)

type Export struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkTransient
}

func (z *Export) Preset() {
}

func (z *Export) Exec(c app_control.Control) error {
	sd := api_auth.OAuthSessionData{
		AppData:  dbx_auth.DropboxIndividual,
		PeerName: app.PeerDeploy,
		Scopes: []string{
			dbx_auth.ScopeFilesContentRead,
			dbx_auth.ScopeFilesContentWrite,
		},
	}
	session := api_auth_oauth.NewSessionCodeAuth(c)
	entity, err := session.Start(sd)
	if err != nil {
		return err
	}

	serialized, err := json.Marshal(&entity)
	if err != nil {
		return err
	}

	ui_out.TextOut(c, string(serialized))
	return nil
}

func (z *Export) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Export{}, rc_recipe.NoCustomValues)
}
