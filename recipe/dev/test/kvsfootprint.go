package test

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Kvsfootprint struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkConsole
	Count        int
	Peer         dbx_conn.ConnUserFile
	Entries      kv_storage.Storage
	ProgressLoop app_msg.Message
}

func (z *Kvsfootprint) Preset() {
	z.Count = 1
}

func (z *Kvsfootprint) Exec(c app_control.Control) error {
	l := c.Log()
	ui := c.UI()
	for i := 0; i < z.Count; i++ {
		ui.Progress(z.ProgressLoop.With("Index", i+1))
		sk := func(entry mo_file.Entry) {
			err := z.Entries.Update(func(kvs kv_kvs.Kvs) error {
				return kvs.PutJson(entry.PathDisplay(), entry.Concrete().Raw)
			})
			if err != nil {
				l.Debug("Unable to store", esl.Error(err))
			}
		}
		err := sv_file.NewFiles(z.Peer.Context()).ListChunked(
			mo_path.NewDropboxPath("/"), sk, sv_file.Recursive())
		if err != nil {
			return err
		}
	}
	return nil
}

func (z *Kvsfootprint) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Kvsfootprint{}, rc_recipe.NoCustomValues)
}
