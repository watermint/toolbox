package rp

import (
	"github.com/google/go-cmp/cmp"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_writer_impl"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestNewSmallCache(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		w := rp_writer_impl.NewMock()
		scw := rp_writer_impl.NewSmallCacheWithThreshold("small_cache_test", w, 3)
		if err := scw.Open(ctl, &rp_writer_impl.MockRecord{}); err != nil {
			t.Error(err)
			return
		}

		if x := len(w.Records()); x != 0 {
			t.Error(x)
		}

		records := []interface{}{
			&rp_writer_impl.MockRecord{SKU: "A001", Quantity: 23},
			&rp_writer_impl.MockRecord{SKU: "A005", Quantity: 48},
			&rp_writer_impl.MockRecord{SKU: "A007", Quantity: 65},
			&rp_writer_impl.MockRecord{SKU: "A013", Quantity: 79},
		}

		// write until threshold
		for i := 0; i < 3; i++ {
			scw.Row(records[i])
			if x := len(w.Records()); x != 0 {
				t.Error(x)
			}
		}

		scw.Row(records[3])

		wroteRecords := w.Records()
		if x := cmp.Diff(records, wroteRecords); x != "" {
			t.Error(x)
		}
	})
}

func TestNewSmallCacheCloseBeforeThreshold(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		w := rp_writer_impl.NewMock()
		scw := rp_writer_impl.NewSmallCacheWithThreshold("small_cache_test", w, 10)
		if err := scw.Open(ctl, &rp_writer_impl.MockRecord{}); err != nil {
			t.Error(err)
			return
		}

		if x := len(w.Records()); x != 0 {
			t.Error(x)
		}

		records := []interface{}{
			&rp_writer_impl.MockRecord{SKU: "A001", Quantity: 23},
			&rp_writer_impl.MockRecord{SKU: "A005", Quantity: 48},
			&rp_writer_impl.MockRecord{SKU: "A007", Quantity: 65},
			&rp_writer_impl.MockRecord{SKU: "A013", Quantity: 79},
		}

		// write until threshold
		for _, record := range records {
			scw.Row(record)
		}
		scw.Close()

		wroteRecords := w.Records()
		if x := cmp.Diff(records, wroteRecords); x != "" {
			t.Error(x)
		}
	})
}
