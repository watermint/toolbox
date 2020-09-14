package eq_stat

import "testing"

func TestStatImpl(t *testing.T) {
	st := New()

	if x, y := st.StatTask("U"); x != 0 || y != 0 {
		t.Error(x, y)
	}

	st.IncrEnqueue("U", "B001")

	if x, y := st.StatTask("U"); x != 0 || y != 1 {
		t.Error(x, y)
	}
	if x, y := st.StatBatch("U"); x != 0 || y != 1 {
		t.Error(x, y)
	}

	st.IncrEnqueue("U", "B001")

	if x, y := st.StatTask("U"); x != 0 || y != 2 {
		t.Error(x, y)
	}
	if x, y := st.StatBatch("U"); x != 0 || y != 1 {
		t.Error(x, y)
	}

	st.IncrEnqueue("U", "B002")

	if x, y := st.StatTask("U"); x != 0 || y != 3 {
		t.Error(x, y)
	}
	if x, y := st.StatBatch("U"); x != 0 || y != 2 {
		t.Error(x, y)
	}

	st.IncrEnqueue("X", "C001")

	if x, y := st.StatTask("U"); x != 0 || y != 3 {
		t.Error(x, y)
	}
	if x, y := st.StatBatch("U"); x != 0 || y != 2 {
		t.Error(x, y)
	}

	st.IncrComplete("U", "B001")

	if x, y := st.StatTask("U"); x != 1 || y != 3 {
		t.Error(x, y)
	}
	if x, y := st.StatBatch("U"); x != 0 || y != 2 {
		t.Error(x, y)
	}

	st.IncrComplete("U", "B001")

	if x, y := st.StatTask("U"); x != 2 || y != 3 {
		t.Error(x, y)
	}
	if x, y := st.StatBatch("U"); x != 1 || y != 2 {
		t.Error(x, y)
	}

	st.IncrComplete("U", "B002")

	if x, y := st.StatTask("U"); x != 3 || y != 3 {
		t.Error(x, y)
	}
	if x, y := st.StatBatch("U"); x != 2 || y != 2 {
		t.Error(x, y)
	}
}
