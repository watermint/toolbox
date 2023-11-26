package kvs

import (
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/essentials/strings/es_uuid"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Concurrency struct {
	rc_recipe.RemarkSecret
	Count int64
	Data  kv_storage.Storage
}

func (z *Concurrency) Preset() {
	z.Count = 10000
}

func (z *Concurrency) process(key string) error {
	value := es_uuid.NewV4().String()
	return z.Data.Update(func(kvs kv_kvs.Kvs) error {
		return kvs.PutString(key, value)
	})
}

func (z *Concurrency) Exec(c app_control.Control) error {
	queueId := "process"
	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define(queueId, z.process)
		q := s.Get(queueId)
		var i int64
		for i = 0; i < z.Count; i++ {
			q.Enqueue(es_uuid.NewV4().String())
		}
	})
	return nil
}

func (z *Concurrency) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}
