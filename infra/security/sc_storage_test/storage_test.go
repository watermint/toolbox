package sc_obfuscate_test

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_storage"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"os"
	"path/filepath"
	"testing"
)

func TestStorageImpl_PutGet(t *testing.T) {
	d, err := qt_file.MakeTestFolder("obfuscate", false)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		os.RemoveAll(d)
	}()

	type ObfuscateTest struct {
		Name     string `json:"name"`
		Quantity int    `json:"quantity"`
	}

	p := filepath.Join(d, "obfuscate.dat")
	ot0 := &ObfuscateTest{
		Name:     "obfuscate",
		Quantity: 12,
	}

	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		s := sc_storage.NewStorage(ctl)

		if err := s.Get(p, &ObfuscateTest{}); err != sc_storage.ErrorStorageNotFound {
			t.Error(err)
			return
		}

		if err := s.Put(p, ot0); err != nil {
			t.Error(err)
			return
		}

		ot1 := &ObfuscateTest{}
		if err := s.Get(p, ot1); err != nil {
			t.Error(err)
			return
		}

		if ot0.Name != ot1.Name || ot0.Quantity != ot1.Quantity {
			t.Error("invalid")
			return
		}
	})
}
