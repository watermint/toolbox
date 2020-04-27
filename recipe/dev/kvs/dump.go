package kvs

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs"
	"github.com/watermint/toolbox/infra/kvs/kv_storage_impl"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"go.uber.org/zap"
)

type DumpResult struct {
	Key       string          `json:"key"`
	Value     string          `json:"value"`
	ValueJson json.RawMessage `json:"value_json"`
}

type Dump struct {
	Path   mo_path.FileSystemPath
	Result rp_model.RowReport
}

func (z *Dump) Preset() {
	z.Result.SetModel(&DumpResult{})
}

func (z *Dump) Exec(c app_control.Control) error {
	l := c.Log()
	kv, err := kv_storage_impl.NewWithPath(c, z.Path.Path())
	if err != nil {
		l.Debug("unable to open", zap.Error(err))
		return err
	}

	if err := z.Result.Open(); err != nil {
		return err
	}

	return kv.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEach(func(key string, value []byte) error {
			if gjson.ValidBytes(value) {
				z.Result.Row(&DumpResult{
					Key:       key,
					ValueJson: value,
				})
			} else {
				z.Result.Row(&DumpResult{
					Key:   key,
					Value: string(value),
				})
			}
			return nil
		})
	})
}

func (z *Dump) Test(c app_control.Control) error {
	return qt_errors.ErrorImplementMe
}
