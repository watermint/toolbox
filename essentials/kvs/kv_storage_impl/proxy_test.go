package kv_storage_impl_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"testing"
)

func kvsTest(t *testing.T, name string, f func(t *testing.T, db kv_storage.Storage)) {
	engines := []kv_storage.KvsEngine{
		kv_storage.KvsEngineBitCask,
		kv_storage.KvsEngineBitcaskTurnstile,
		kv_storage.KvsEngineSqlite,
		kv_storage.KvsEngineSqliteTurnstile,
		kv_storage.KvsEngineBadger,
	}

	for _, engine := range engines {
		name := fmt.Sprintf("kvs_%s_%d", name, engine)
		t.Log("Engine", engine)
		qt_file.TestWithTestFolder(t, name, false, func(path string) {
			db := kv_storage_impl.NewProxy(name, esl.Default())
			if dbp, ok := db.(kv_storage.Proxy); ok {
				dbp.SetEngine(engine)
			} else {
				t.Error("Unable to set engine")
				return
			}
			if err := db.Open(path); err != nil {
				t.Error(err)
				return
			}
			f(t, db)
			db.Close()
		})
	}
}

func TestPutGetString(t *testing.T) {
	kvsTest(t, "coffee_put-get-string", func(t *testing.T, db kv_storage.Storage) {
		var err error

		// put/get string
		err = db.Update(func(coffee kv_kvs.Kvs) error {
			if err = coffee.PutString("A1234", "espresso"); err != nil {
				t.Error(err)
				return err
			}
			if err = coffee.PutString("A5678", "カフェラテ"); err != nil {
				t.Error(err)
				return err
			}
			if v, err := coffee.GetString("A1234"); v != "espresso" || err != nil {
				t.Error(v, err)
				return err
			}
			if v, err := coffee.GetString("A5678"); v != "カフェラテ" || err != nil {
				t.Error(v, err)
				return err
			}
			if err = coffee.Delete("A1234"); err != nil {
				t.Error(err)
				return err
			}
			if v, err := coffee.GetString("A1234"); err == nil || v != "" {
				t.Error(v, err)
				return err
			}
			return nil
		})
		if err != nil {
			t.Error(err)
		}
	})
}

func TestPutGetJSON(t *testing.T) {
	kvsTest(t, "coffee_put-get-json", func(t *testing.T, db kv_storage.Storage) {
		var err error

		// put/get json
		err = db.Update(func(coffee kv_kvs.Kvs) error {
			type SKU struct {
				Name  string `json:"name"`
				Price int    `json:"price"`
			}
			a1234 := &SKU{
				Name:  "Espresso",
				Price: 350,
			}
			a1234raw, err := json.Marshal(a1234)
			if err != nil {
				t.Error(err)
				return err
			}
			a5678 := &SKU{
				Name:  "カフェラテ",
				Price: 500,
			}
			a5678raw, err := json.Marshal(a5678)
			if err != nil {
				t.Error(err)
				return err
			}

			if err != nil {
				t.Error(err)
				return err
			}
			if err = coffee.PutJson("A1234", a1234raw); err != nil {
				t.Error(err)
				return err
			}
			if err = coffee.PutJson("A5678", a5678raw); err != nil {
				t.Error(err)
				return err
			}

			return nil
		})
		if err != nil {
			t.Error(err)
		}
	})
}

func TestPutGetJSONModel(t *testing.T) {
	kvsTest(t, "coffee_put-get-json-model", func(t *testing.T, db kv_storage.Storage) {
		// put/get json model
		upErr := db.Update(func(coffee kv_kvs.Kvs) error {
			type SKU struct {
				Name  string `json:"name"`
				Price int    `json:"price"`
			}
			a1234 := &SKU{
				Name:  "Espresso",
				Price: 350,
			}
			a5678 := &SKU{
				Name:  "カフェラテ",
				Price: 500,
			}

			if puErr := coffee.PutJsonModel("A1234", a1234); puErr != nil {
				t.Error(puErr)
				return puErr
			}
			if puErr := coffee.PutJsonModel("A5678", a5678); puErr != nil {
				t.Error(puErr)
				return puErr
			}
			v := &SKU{}
			if geErr := coffee.GetJsonModel("A1234", v); v.Name != a1234.Name || v.Price != a1234.Price || geErr != nil {
				t.Error(v, geErr)
				return geErr
			}
			v = &SKU{}
			if geErr := coffee.GetJsonModel("A5678", v); v.Name != a5678.Name || v.Price != a5678.Price || geErr != nil {
				t.Error(v, geErr)
				return geErr
			}
			return nil
		})
		if upErr != nil {
			t.Error(upErr)
		}

		// foreach, cursor
		fcErr := db.Update(func(coffee kv_kvs.Kvs) error {
			dat := map[string]string{
				"A1234": "Espresso",
				"A5678": "カフェラテ",
				"A9012": "☕️",
			}
			for k, v := range dat {
				if puErr := coffee.PutString(k, v); puErr != nil {
					t.Error(puErr)
					return puErr
				}
			}

			// foreach
			{
				var foErr error
				found := make(map[string]bool)
				foErr = coffee.ForEach(func(key string, value []byte) error {
					if c, e := dat[key]; !e || c != string(value) {
						t.Error(key, value)
						return errors.New("invalid value")
					}
					found[key] = true
					return nil
				})

				if nil != foErr {
					t.Error(foErr)
					return foErr
				}
				for k := range dat {
					if _, e := found[k]; !e {
						t.Error(k)
						return errors.New("key not found")
					}
				}
			}

			return nil
		})
		if fcErr != nil {
			t.Error(fcErr)
		}
	})
}

func TestPutGetJSONModel2(t *testing.T) {
	kvsTest(t, "coffee_put-get-json-model2", func(t *testing.T, db kv_storage.Storage) {
		var err error

		// foreach model
		err = db.Update(func(coffee kv_kvs.Kvs) error {
			type TestForEachModel struct {
				Name     string `json:"name"`
				Quantity int    `json:"quantity"`
			}

			transactions := make(map[string][]*TestForEachModel)
			transactions["E0123"] = []*TestForEachModel{
				{
					Name:     "Espresso",
					Quantity: 2,
				},
				{
					Name:     "Latte",
					Quantity: 1,
				},
			}
			transactions["E0129"] = []*TestForEachModel{
				{
					Name:     "カフェラテ",
					Quantity: 1,
				},
				{
					Name:     "ブレンド",
					Quantity: 4,
				},
			}

			for k, v := range transactions {
				if err := coffee.PutJsonModel(k, v); err != nil {
					t.Error(err)
				}
			}

			// foreach
			{
				found := make(map[string]bool)
				err = coffee.ForEachModel(&[]*TestForEachModel{}, func(key string, m interface{}) error {
					if c, e := transactions[key]; !e {
						t.Error(key)
						return errors.New("invalid value")
					} else {
						if len(c) != 2 {
							t.Error(c)
						}
					}
					found[key] = true
					return nil
				})
				if err != nil {
					t.Error(err)
					return err
				}
				for k := range transactions {
					if _, e := found[k]; !e {
						t.Error(k)
						return errors.New("key not found")
					}
				}
			}

			return nil
		})
		if err != nil {
			t.Error(err)
		}
	})
}
