package kv

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"math/rand"
	"testing"
)

func BenchmarkMemoryFootprint(b *testing.B) {
	dataSizeMin := 1024
	dataSizeMax := 16384
	mul := 10
	qtr_endtoend.BenchmarkWithControl(b, func(ctl app_control.Control) {
		var err error
		db := kv_storage_impl.InternalNewBadger("benchmark-memory-footprint")
		if err := db.Open(ctl); err != nil {
			b.Error(err)
			return
		}
		defer db.Close()
		l := ctl.Log()
		l.Info("Benchmark", esl.Int("N", b.N))
		for i := 0; i < b.N*mul; i++ {
			err = db.Update(func(bmf kv_kvs.Kvs) error {
				key := fmt.Sprintf("k%x", i)
				dataSize := rand.Intn(dataSizeMax) + dataSizeMin
				data := make([]byte, dataSize)
				if n, err := rand.Read(data); err != nil || n != dataSize {
					b.Error(err, n)
					return err
				}

				if err = bmf.PutBytes(key, data); err != nil {
					b.Error(err)
					return err
				}
				return nil
			})
		}

		for i := 0; i < b.N*mul; i++ {
			err = db.View(func(bmf kv_kvs.Kvs) error {
				key := fmt.Sprintf("k%x", i)
				dataSize := rand.Intn(dataSizeMax) + dataSizeMin
				data := make([]byte, dataSize)
				if n, err := rand.Read(data); err != nil || n != dataSize {
					b.Error(err, n)
					return err
				}

				if _, err = bmf.GetBytes(key); err != nil {
					b.Error(err)
					return err
				}
				return nil
			})
		}

		if err != nil {
			b.Error(err)
		}
	})
}
