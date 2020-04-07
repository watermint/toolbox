package auth

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type Export struct {
	Full  rc_conn.ConnUserFile
	Info  rc_conn.ConnBusinessInfo
	File  rc_conn.ConnBusinessFile
	Audit rc_conn.ConnBusinessAudit
	Mgmt  rc_conn.ConnBusinessMgmt
}

func (z *Export) Preset() {
	z.Full.SetPeerName(qt_endtoend.EndToEndPeer)
	z.Info.SetPeerName(qt_endtoend.EndToEndPeer)
	z.File.SetPeerName(qt_endtoend.EndToEndPeer)
	z.Audit.SetPeerName(qt_endtoend.EndToEndPeer)
	z.Mgmt.SetPeerName(qt_endtoend.EndToEndPeer)
}

func (z *Export) Exec(c app_control.Control) error {
	l := c.Log()
	e := make(map[string]*oauth2.Token)
	a := dbx_auth.NewConsoleCacheOnly(c, qt_endtoend.EndToEndPeer)
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
	o := ut_io.NewDefaultOut(c.IsTest())
	o.Write(b)
	o.Write([]byte("\n"))
	o.Close()

	return nil
}

func (z *Export) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Export{}, rc_recipe.NoCustomValues)
}
