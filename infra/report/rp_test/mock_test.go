package rp

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_writer_impl"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestMock(t *testing.T) {
	m := rp_writer_impl.NewMock()
	if m.IsOpened() || m.IsClosed() {
		t.Error(m.IsOpened(), m.IsClosed())
	}

	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		if err := m.Open(ctl, &rp_writer_impl.MockRecord{}); err != nil {
			t.Error(err)
			return
		}
		if x := len(m.Records()); x != 0 {
			t.Error(x)
		}
		m.Row(&rp_writer_impl.MockRecord{SKU: "A123", Quantity: 91})
		if x := len(m.Records()); x != 1 {
			t.Error(x)
		}
		m.Close()
		if !m.IsClosed() {
			t.Error(m.IsClosed())
		}
	})
}

func TestMock_RowWithOutOpen(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			if err != rp_writer_impl.ErrorMockTheWriterIsNotReady {
				t.Error(err)
			}
		} else {
			t.Error("No panic reported")
		}
	}()

	m := rp_writer_impl.NewMock()
	m.Row(&rp_writer_impl.MockRecord{SKU: "A111", Quantity: 38})
}
