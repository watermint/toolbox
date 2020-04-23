package artifact

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"github.com/watermint/toolbox/recipe/dev/ci/auth"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type Connect struct {
	rc_recipe.RemarkSecret
	Full dbx_conn.ConnUserFile
}

func (z *Connect) Preset() {
	z.Full.SetPeerName(app.PeerDeploy)
}

func (z *Connect) Exec(c app_control.Control) error {
	l := c.Log()
	e := make(map[string]*oauth2.Token)
	a := api_auth_impl.NewConsoleCacheOnly(c, z.Full.PeerName())
	for _, s := range auth.Scopes {
		t, err := a.Auth(s)
		if err != nil {
			l.Info("Skip export", zap.Error(err), zap.String("scope", s))
			continue
		}
		e[s] = t.Token()
	}
	b, err := json.Marshal(e)
	if err != nil {
		return err
	}
	o := ut_io.NewDefaultOut(c.Feature().IsTest())
	o.Write(b)
	o.Write([]byte("\n"))
	o.Close()
	return nil
}

func (z *Connect) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Connect{}, rc_recipe.NoCustomValues)
}
