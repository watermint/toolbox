package kv_storage_impl_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"testing"
)

func kvsTest(t *testing.T, name string, f func(t *testing.T, db kv_storage.Storage)) {
	qt_file.TestWithTestFolder(t, "kvs"+name, false, func(path string) {
		db := kv_storage_impl.InternalNewBitcask(name+"_bitcask", esl.Default())
		if err := db.Open(path); err != nil {
			t.Error(err)
			return
		}
		f(t, db)
		db.Close()
	})
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

func TestPutGetBytes(t *testing.T) {
	kvsTest(t, "coffee_put-get-bytes", func(t *testing.T, db kv_storage.Storage) {
		var err error

		// put/get bytes
		err = db.Update(func(coffee kv_kvs.Kvs) error {
			a1234 := []byte("Espresso")
			if err = coffee.PutBytes("A1234", a1234); err != nil {
				t.Error(err)
				return err
			}
			a5678 := []byte("カフェラテ")
			if err = coffee.PutBytes("A5678", a5678); err != nil {
				t.Error(err)
				return err
			}
			if v, err := coffee.GetBytes("A1234"); bytes.Compare(v, a1234) != 0 || err != nil {
				t.Error(v, err)
				return err
			}
			if v, err := coffee.GetBytes("A5678"); bytes.Compare(v, a5678) != 0 || err != nil {
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
			if v, err := coffee.GetBytes("A1234"); bytes.Compare(v, a1234raw) != 0 || err != nil {
				t.Error(v, err)
				return err
			}
			if v, err := coffee.GetBytes("A5678"); bytes.Compare(v, a5678raw) != 0 || err != nil {
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

func TestPutGetJSONModel(t *testing.T) {
	kvsTest(t, "coffee_put-get-json-model", func(t *testing.T, db kv_storage.Storage) {
		var err error

		// put/get json model
		err = db.Update(func(coffee kv_kvs.Kvs) error {
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

			if err = coffee.PutJsonModel("A1234", a1234); err != nil {
				t.Error(err)
				return err
			}
			if err = coffee.PutJsonModel("A5678", a5678); err != nil {
				t.Error(err)
				return err
			}
			v := &SKU{}
			if err := coffee.GetJsonModel("A1234", v); v.Name != a1234.Name || v.Price != a1234.Price || err != nil {
				t.Error(v, err)
				return err
			}
			v = &SKU{}
			if err := coffee.GetJsonModel("A5678", v); v.Name != a5678.Name || v.Price != a5678.Price || err != nil {
				t.Error(v, err)
				return err
			}
			return nil
		})
		if err != nil {
			t.Error(err)
		}

		// foreach, cursor
		err = db.Update(func(coffee kv_kvs.Kvs) error {
			dat := map[string]string{
				"A1234": "Espresso",
				"A5678": "カフェラテ",
				"A9012": "☕️",
			}
			for k, v := range dat {
				if err = coffee.PutString(k, v); err != nil {
					t.Error(err)
					return err
				}
			}

			// foreach
			{
				found := make(map[string]bool)
				err = coffee.ForEach(func(key string, value []byte) error {
					if c, e := dat[key]; !e || c != string(value) {
						t.Error(key, value)
						return errors.New("invalid value")
					}
					found[key] = true
					return nil
				})
				if err != nil {
					t.Error(err)
					return err
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
		if err != nil {
			t.Error(err)
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
