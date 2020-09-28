package kv_storage_impl

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"math/rand"
	"os"
	"testing"
)

func benchmarkLoadTest(b *testing.B, db kv_storage.Storage) {
	dataSizeMin := 2 << 10
	dataSizeMax := 2 << 16
	mul := 10
	var err error

	l := esl.Default()
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

func BenchmarkMemoryFootprintBitcask(b *testing.B) {
	path, err := qt_file.MakeTestFolder("benchmark-memory-bitcask", false)
	if err != nil {
		b.Error(err)
		return
	}
	defer func() {
		_ = os.RemoveAll(path)
	}()

	db := InternalNewBitcask("benchmark-memory-bitcask", esl.Default())
	if err := db.Open(path); err != nil {
		b.Error(err)
		return
	}
	benchmarkLoadTest(b, db)
	db.Close()
}
