package kv_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs"
	"github.com/watermint/toolbox/infra/kvs/kv_storage_impl"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestNew(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		var err error

		db1 := kv_storage_impl.New("coffee1")
		if err := db1.Open(ctl); err != nil {
			t.Error(err)
			return
		}
		defer db1.Close()
		db2 := kv_storage_impl.New("coffee2")
		if err := db2.Open(ctl); err != nil {
			t.Error(err)
			return
		}
		defer db2.Close()
		db3 := kv_storage_impl.New("coffee3")
		if err := db3.Open(ctl); err != nil {
			t.Error(err)
			return
		}
		defer db3.Close()
		db4 := kv_storage_impl.New("coffee4")
		if err := db4.Open(ctl); err != nil {
			t.Error(err)
			return
		}
		defer db4.Close()

		// put/get string
		err = db1.Update(func(coffee kv_kvs.Kvs) error {
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

		// put/get bytes
		err = db2.Update(func(coffee kv_kvs.Kvs) error {
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

		// put/get json
		err = db3.Update(func(coffee kv_kvs.Kvs) error {
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

		// put/get json model
		err = db4.Update(func(coffee kv_kvs.Kvs) error {
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
		err = db4.Update(func(coffee kv_kvs.Kvs) error {
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
