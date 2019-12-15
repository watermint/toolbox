package kv_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_storage_impl"
	"github.com/watermint/toolbox/infra/kvs/kv_transaction"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestNew(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		db, err := kv_storage_impl.New(ctl, "test")
		if err != nil {
			t.Error(err)
			return
		}

		// put/get string
		err = db.Update(func(tx kv_transaction.Transaction) error {
			coffee, err := tx.Kvs("coffee1")
			if err != nil {
				t.Error(err)
				return err
			}
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
		err = db.Update(func(tx kv_transaction.Transaction) error {
			coffee, err := tx.Kvs("coffee2")
			if err != nil {
				t.Error(err)
				return err
			}
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
		err = db.Update(func(tx kv_transaction.Transaction) error {
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

			coffee, err := tx.Kvs("coffee3")
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
		err = db.Update(func(tx kv_transaction.Transaction) error {
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

			coffee, err := tx.Kvs("coffee4")
			if err != nil {
				t.Error(err)
				return err
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
		err = db.Batch(func(tx kv_transaction.Transaction) error {
			coffee, err := tx.Kvs("coffee4")
			if err != nil {
				t.Error(err)
				return err
			}
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

			// cursor first-next
			{
				found := make(map[string]bool)
				c := coffee.Cursor()

				for k, v, e := c.First(); e; k, v, e = c.Next() {
					if c, e := dat[k]; !e || c != string(v) {
						t.Error(k, v)
						return errors.New("invalid value")
					}
					found[k] = true
				}
				for k := range dat {
					if _, e := found[k]; !e {
						t.Error(k)
						return errors.New("key not found")
					}
				}
			}

			// cursor last-prev
			{
				found := make(map[string]bool)
				c := coffee.Cursor()

				for k, v, e := c.Last(); e; k, v, e = c.Prev() {
					if c, e := dat[k]; !e || c != string(v) {
						t.Error(k, v)
						return errors.New("invalid value")
					}
					found[k] = true
				}
				for k := range dat {
					if _, e := found[k]; !e {
						t.Error(k)
						return errors.New("key not found")
					}
				}
			}

			// cursor seek
			{
				c := coffee.Cursor()
				k, v, e := c.Seek("A5678")
				if !e || k != "A5678" || bytes.Compare(v, []byte("カフェラテ")) != 0 {
					t.Error(k, v, e)
					return errors.New("invalid")
				}
			}
			return nil
		})
		if err != nil {
			t.Error(err)
		}

		// view: foreach, cursor
		err = db.View(func(tx kv_transaction.Transaction) error {
			coffee, err := tx.Kvs("coffee4")
			if err != nil {
				t.Error(err)
				return err
			}
			dat := map[string]string{
				"A1234": "Espresso",
				"A5678": "カフェラテ",
				"A9012": "☕️",
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

			// cursor first-next
			{
				found := make(map[string]bool)
				c := coffee.Cursor()

				for k, v, e := c.First(); e; k, v, e = c.Next() {
					if c, e := dat[k]; !e || c != string(v) {
						t.Error(k, v)
						return errors.New("invalid value")
					}
					found[k] = true
				}
				for k := range dat {
					if _, e := found[k]; !e {
						t.Error(k)
						return errors.New("key not found")
					}
				}
			}

			// cursor last-prev
			{
				found := make(map[string]bool)
				c := coffee.Cursor()

				for k, v, e := c.Last(); e; k, v, e = c.Prev() {
					if c, e := dat[k]; !e || c != string(v) {
						t.Error(k, v)
						return errors.New("invalid value")
					}
					found[k] = true
				}
				for k := range dat {
					if _, e := found[k]; !e {
						t.Error(k)
						return errors.New("key not found")
					}
				}
			}

			// cursor seek
			{
				c := coffee.Cursor()
				k, v, e := c.Seek("A5678")
				if !e || k != "A5678" || bytes.Compare(v, []byte("カフェラテ")) != 0 {
					t.Error(k, v, e)
					return errors.New("invalid")
				}
			}
			return nil
		})
		if err != nil {
			t.Error(err)
		}

		db.Close()
	})
}
