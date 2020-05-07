package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs"
	"github.com/watermint/toolbox/infra/kvs/kv_storage"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

type ValueKvStorageStorageRecipe struct {
	Data kv_storage.Storage
}

func (z ValueKvStorageStorageRecipe) Preset() {
}

func (z ValueKvStorageStorageRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z ValueKvStorageStorageRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueKvStorageStorage(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueKvStorageStorageRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		if err := flg.Parse([]string{}); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		_ = repo.Apply()

		// Spin up
		rcp2, err := repo.SpinUp(c)
		if err != nil {
			t.Error(err)
			return err
		}
		mod2 := rcp2.(*ValueKvStorageStorageRecipe)
		err = mod2.Data.Update(func(kvs kv_kvs.Kvs) error {
			return kvs.PutString("ping", "pong")
		})
		if err != nil {
			t.Error(err)
		}

		if err := repo.SpinDown(c); err != nil {
			t.Error(err)
			return err
		}

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
