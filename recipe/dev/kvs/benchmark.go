package kvs

import (
	"github.com/watermint/toolbox/essentials/kvs/kv_storage_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type Benchmark struct {
	rc_recipe.RemarkSecret
}

func (z *Benchmark) Preset() {
}

func (z *Benchmark) Exec(c app_control.Control) error {
	kvSqlite := kv_storage_impl.NewSqlite("benchmark", c.Log())
	if err := kvSqlite.Open(c.Workspace().KVS()); err != nil {
		return err
	}
	kvSqlite.Close()
	return nil
}

func (z *Benchmark) Test(c app_control.Control) error {
	return nil
}
