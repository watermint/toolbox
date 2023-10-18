package kvs

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type DumpResult struct {
	Key       string          `json:"key"`
	Value     string          `json:"value"`
	ValueJson json.RawMessage `json:"value_json"`
}

type Dump struct {
	rc_recipe.RemarkSecret
	Path   mo_path.FileSystemPath
	Result rp_model.RowReport
}

func (z *Dump) Preset() {
	z.Result.SetModel(&DumpResult{})
}

func (z *Dump) Exec(c app_control.Control) error {
	l := c.Log()
	proxy := kv_storage_impl.NewProxy("dump", l)
	if ls, ok := proxy.(kv_storage.Proxy); ok {
		ls.SetEngine(c.Feature().KvsEngine())
	}
	err := proxy.Open(z.Path.Path())
	if err != nil {
		l.Debug("unable to open", esl.Error(err))
		return err
	}

	if err := z.Result.Open(); err != nil {
		return err
	}

	return proxy.View(func(kvs kv_kvs.Kvs) error {
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
