package es_json

import (
	"strconv"
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

func TestToJsonString(t *testing.T) {
	{
		v := struct {
			SKU   string `json:"sku"`
			Price int    `json:"price"`
		}{
			SKU:   "A123",
			Price: 123,
		}
		if s := ToJsonString(v); s != `{"sku":"A123","price":123}` {
			t.Error(s)
		}
	}

	{
		v := map[string]string{
			"_sku": "A123",
			"name": "Grapefruit",
		}
		if s := ToJsonString(v); s != `{"_sku":"A123","name":"Grapefruit"}` {
			t.Error(s)
		}
	}

	{
		if s := ToJsonString(nil); s != "null" {
			t.Error(s)
		}
	}

	// non serializable value
	{
		if s := ToJsonString(func() {}); s != "null" {
			t.Error(s)
		}
	}

	{
		v := MustParseString(`{"sku":"A123","price":123}`)
		if s := ToJsonString(v); s != `{"sku":"A123","price":123}` {
			t.Error(s)
		}
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
		if s, e := j.Number(); s != "" || e {
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
		if s, e := j.Number(); s != "" || e {
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
		if s, e := j.Number(); s != "" || e {
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
		if s, e := j.Number(); s != "" || e {
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
		if s, e := j.Number(); !e {
			t.Error(s, e)
		} else {
			if i, err := strconv.ParseInt(s, 10, 64); err != nil || i != 1234 {
				t.Error(s, e, i, err)
			}
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
		if s, e := j.Number(); !e {
			t.Error(s, e)
		} else {
			if f, err := strconv.ParseFloat(s, 64); err != nil || f != 1234.56 {
				t.Error(s, e, f, err)
			}
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
		if s, e := j.Number(); s != "" || e {
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

func TestJsonModelType(t *testing.T) {
	type Order struct {
		Code    string `path:"code"`
		Valid   string `path:"valid"` // string form
		IsValid bool   `path:"valid"` // bool form
	}

	// `valid` exists
	{
		m := `{"code":"A1234","valid":true}`
		m1 := &Order{}
		if j, err := ParseString(m); err != nil {
			t.Error(err)
		} else if err := j.Model(m1); err != nil {
			t.Error(err)
		} else {
			if x := m1.IsValid; !x {
				t.Error(x)
			}
			if x := m1.Valid; x != "true" {
				t.Error(x)
			}
		}
	}

	// `valid` is not exists
	{
		m := `{"code":"A1234"}`
		m1 := &Order{}
		if j, err := ParseString(m); err != nil {
			t.Error(err)
		} else if err := j.Model(m1); err != nil {
			t.Error(err)
		} else {
			if x := m1.IsValid; x {
				t.Error(x)
			}
			if x := m1.Valid; x != "" {
				t.Error(x)
			}
		}
	}
}
