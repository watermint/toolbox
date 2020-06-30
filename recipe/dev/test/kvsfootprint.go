package test

import (
	"fmt"
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
	Entries1     kv_storage.Storage
	Entries2     kv_storage.Storage
	Entries3     kv_storage.Storage
	Entries4     kv_storage.Storage
	Entries5     kv_storage.Storage
	ProgressLoop app_msg.Message
}

func (z *Kvsfootprint) Preset() {
	z.Count = 1
}

func (z *Kvsfootprint) Exec(c app_control.Control) error {
	l := c.Log()
	ui := c.UI()

	storages := []kv_storage.Storage{
		z.Entries1,
		z.Entries2,
		z.Entries3,
		z.Entries4,
		z.Entries5,
	}

	for i := 0; i < z.Count; i++ {
		ui.Progress(z.ProgressLoop.With("Index", i+1))
		for _, storage := range storages {
			sk := func(entry mo_file.Entry) {
				key := fmt.Sprintf("%x:%s", i, entry.PathDisplay())
				err := storage.Update(func(kvs kv_kvs.Kvs) error {
					return kvs.PutJson(key, entry.Concrete().Raw)
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

			_ = storage.View(func(kvs kv_kvs.Kvs) error {
				return kvs.ForEach(func(key string, value []byte) error {
					return nil
				})
			})
		}
	}
	return nil
}

func (z *Kvsfootprint) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Kvsfootprint{}, rc_recipe.NoCustomValues)
}
