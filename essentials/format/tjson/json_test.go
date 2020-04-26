package tjson

import (
	"testing"
)

func TestPath(t *testing.T) {
	if j, err := ParseString(`["orange", "apple", "banana"]`); err != nil {
		t.Error(err)
	} else if k, found := j.Find(PathArrayFirst); !found {
		t.Error(found)
	} else if s, found := k.String(); !found {
		t.Error(found)
	} else {
		t.Log(s)
	}
}

func TestParse(t *testing.T) {
	// object
	if j, err := Parse([]byte("{}")); err == nil {
		if j.IsNull() {
			t.Error(j.IsNull())
		}
		if s, e := j.String(); s != "" || e {
			t.Error(s, e)
		}
		if s, e := j.Array(); s != nil || e {
			t.Error(s, e)
		}
		if s, e := j.Number(); s != nil || e {
			t.Error(s, e)
		}
		if s, e := j.Bool(); s || e {
			t.Error(s, e)
		}
		if s, e := j.Object(); len(s) != 0 || !e {
			t.Error(s, e)
		}
	} else {
		t.Error(err)
	}

	// object 1
	if j, err := Parse([]byte(`{"name":"David"}`)); err == nil {
		if j.IsNull() {
			t.Error(j.IsNull())
		}
		if s, e := j.String(); s != "" || e {
			t.Error(s, e)
		}
		if s, e := j.Array(); s != nil || e {
			t.Error(s, e)
		}
		if s, e := j.Number(); s != nil || e {
			t.Error(s, e)
		}
		if s, e := j.Bool(); s || e {
			t.Error(s, e)
		}
		if s, e := j.Object(); len(s) != 1 || !e {
			t.Error(s, e)
		} else if s1, e1 := s["name"].String(); s1 != "David" || !e1 {
			t.Error(s1, e1)
		}
	} else {
		t.Error(err)
	}

	// true
	if j, err := Parse([]byte(`true`)); err == nil {
		if j.IsNull() {
			t.Error(j.IsNull())
		}
		if s, e := j.String(); s != "" || e {
			t.Error(s, e)
		}
		if s, e := j.Array(); s != nil || e {
			t.Error(s, e)
		}
		if s, e := j.Number(); s != nil || e {
			t.Error(s, e)
		}
		if s, e := j.Bool(); !s || !e {
			t.Error(s, e)
		}
		if s, e := j.Object(); len(s) != 0 || e {
			t.Error(s, e)
		}
	} else {
		t.Error(err)
	}

	// false
	if j, err := Parse([]byte(`false`)); err == nil {
		if j.IsNull() {
			t.Error(j.IsNull())
		}
		if s, e := j.String(); s != "" || e {
			t.Error(s, e)
		}
		if s, e := j.Array(); s != nil || e {
			t.Error(s, e)
		}
		if s, e := j.Number(); s != nil || e {
			t.Error(s, e)
		}
		if s, e := j.Bool(); s || !e {
			t.Error(s, e)
		}
		if s, e := j.Object(); len(s) != 0 || e {
			t.Error(s, e)
		}
	} else {
		t.Error(err)
	}

	// number: int
	if j, err := Parse([]byte(`1234`)); err == nil {
		if j.IsNull() {
			t.Error(j.IsNull())
		}
		if s, e := j.String(); s != "" || e {
			t.Error(s, e)
		}
		if s, e := j.Array(); s != nil || e {
			t.Error(s, e)
		}
		if s, e := j.Number(); s.Int() != 1234 || !e {
			t.Error(s, e)
		}
		if s, e := j.Bool(); s || e {
			t.Error(s, e)
		}
		if s, e := j.Object(); len(s) != 0 || e {
			t.Error(s, e)
		}
	} else {
		t.Error(err)
	}

	// number: float
	if j, err := Parse([]byte(`1234.56`)); err == nil {
		if j.IsNull() {
			t.Error(j.IsNull())
		}
		if s, e := j.String(); s != "" || e {
			t.Error(s, e)
		}
		if s, e := j.Array(); s != nil || e {
			t.Error(s, e)
		}
		if s, e := j.Number(); s.Float64() != 1234.56 || !e {
			t.Error(s, e)
		}
		if s, e := j.Bool(); s || e {
			t.Error(s, e)
		}
		if s, e := j.Object(); len(s) != 0 || e {
			t.Error(s, e)
		}
	} else {
		t.Error(err)
	}

	// string
	if j, err := Parse([]byte(`"tree"`)); err == nil {
		if j.IsNull() {
			t.Error(j.IsNull())
		}
		if s, e := j.String(); s != "tree" || !e {
			t.Error(s, e)
		}
		if s, e := j.Array(); s != nil || e {
			t.Error(s, e)
		}
		if s, e := j.Number(); s != nil || e {
			t.Error(s, e)
		}
		if s, e := j.Bool(); s || e {
			t.Error(s, e)
		}
		if s, e := j.Object(); len(s) != 0 || e {
			t.Error(s, e)
		}
	} else {
		t.Error(err)
	}
}

func TestJsonModel(t *testing.T) {
	type Order struct {
		Code     string `path:"sku.code"`
		Name     string `path:"sku.name"`
		Quantity int    `path:"quantity"`
	}

	m := `{"sku":{"name":"Notebook", "code":"1234"}, "quantity": 48}`
	m1 := &Order{}
	if j, err := ParseString(m); err != nil {
		t.Error(err)
	} else if err := j.Model(m1); err != nil {
		t.Error(err)
	} else {
		if m1.Name != "Notebook" || m1.Code != "1234" || m1.Quantity != 48 {
			t.Error(m1)
		}
	}
}

func TestJsonFind(t *testing.T) {
	m := `{"sku":{"name":"Notebook", "code":"1234"}, "quantity": 48}`
	if j, err := ParseString(m); err != nil {
		t.Error(err)
	} else {
		if k, found := j.Find("sku.name"); !found {
			t.Error(found)
		} else if n, found := k.String(); !found || n != "Notebook" {
			t.Error(n, found)
		}
	}
}

func TestJsonFindModel(t *testing.T) {
	type Order struct {
		Code     string `path:"sku.code"`
		Name     string `path:"sku.name"`
		Quantity int    `path:"quantity"`
	}

	m := `{"transactions":{"sku":{"name":"Notebook", "code":"1234"}, "quantity": 48}}`
	m1 := &Order{}
	if j, err := ParseString(m); err != nil {
		t.Error(err)
	} else if err := j.FindModel("transactions", m1); err != nil {
		t.Error(err)
	} else {
		if m1.Name != "Notebook" || m1.Code != "1234" || m1.Quantity != 48 {
			t.Error(m1)
		}
	}
}
