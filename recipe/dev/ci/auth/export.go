package auth

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type Export struct {
	rc_recipe.RemarkSecret
	Full  dbx_conn.ConnUserFile
	Info  dbx_conn.ConnBusinessInfo
	File  dbx_conn.ConnBusinessFile
	Audit dbx_conn.ConnBusinessAudit
	Mgmt  dbx_conn.ConnBusinessMgmt
}

func (z *Export) Preset() {
	z.Full.SetPeerName(app.PeerEndToEndTest)
	z.Info.SetPeerName(app.PeerEndToEndTest)
	z.File.SetPeerName(app.PeerEndToEndTest)
	z.Audit.SetPeerName(app.PeerEndToEndTest)
	z.Mgmt.SetPeerName(app.PeerEndToEndTest)
}

func (z *Export) Exec(c app_control.Control) error {
	l := c.Log()
	e := make(map[string]*oauth2.Token)
	a := api_auth_impl.NewConsoleCacheOnly(c, app.PeerEndToEndTest)
	for _, s := range Scopes {
		t, err := a.Auth(s)
		if err != nil {
			l.Info("Skip export", zap.Error(err), zap.String("scope", s))
			return nil
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

func (z *Export) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Export{}, rc_recipe.NoCustomValues)
}
