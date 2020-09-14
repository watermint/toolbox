package kv

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"math/rand"
	"testing"
)

func benchmarkLoadTest(b *testing.B, ctl app_control.Control, db kv_storage.Storage) {
	dataSizeMin := 2 << 10
	dataSizeMax := 2 << 16
	mul := 10
	var err error

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
}

func BenchmarkMemoryFootprintBadger(b *testing.B) {
	qtr_endtoend.BenchmarkWithControl(b, func(ctl app_control.Control) {
		db := kv_storage_impl.InternalNewBadger("benchmark-memory-badger")
		if err := db.Open(ctl); err != nil {
			b.Error(err)
			return
		}
		benchmarkLoadTest(b, ctl, db)
		db.Close()
	})
}

func BenchmarkMemoryFootprintBitcask(b *testing.B) {
	qtr_endtoend.BenchmarkWithControl(b, func(ctl app_control.Control) {
		db := kv_storage_impl.InternalNewBitcask("benchmark-memory-bitcask")
		if err := db.Open(ctl); err != nil {
			b.Error(err)
			return
		}
		benchmarkLoadTest(b, ctl, db)
		db.Close()
	})
}
