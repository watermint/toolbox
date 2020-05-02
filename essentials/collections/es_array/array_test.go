package es_array

import (
	"github.com/watermint/toolbox/essentials/collections/es_value"
	"testing"
)

func TestEmpty(t *testing.T) {
	e := Empty()
	if !e.IsEmpty() {
		t.Error(e.IsEmpty())
	}
	if e.Size() != 0 {
		t.Error(e.Size())
	}
}

func TestNewByInterface(t *testing.T) {
	a := NewByInterface(1, 2, 3)
	if a.Size() != 3 {
		t.Error(a.Size())
	}
	if a.First().AsNumber().Int() != 1 {
		t.Error(a.First())
	}
	if a.Last().AsNumber().Int() != 3 {
		t.Error(a.Last())
	}

	b := a.Reverse()
	if b.Size() != 3 {
		t.Error(b.Size())
	}
	if b.First().AsNumber().Int() != 3 {
		t.Error(b.First().AsNumber().Int())
	}
	if b.Last().AsNumber().Int() != 1 {
		t.Error(b.Last().AsNumber().Int())
	}
}

func TestArrayImpl_Left(t *testing.T) {
	a := NewByInterface(1, 2, 3)
	b := a.Left(2)

	if b.Size() != 2 {
		t.Error(b.Size())
	}
	if b.First().AsNumber().Int() != 1 {
		t.Error(b.First())
	}
	if b.Last().AsNumber().Int() != 2 {
		t.Error(b.Last())
	}
}

func TestArrayImpl_LeftWhile(t *testing.T) {
	a := NewByInterface(1, 2, 3)
	b := a.LeftWhile(func(v es_value.Value) bool {
		return v.AsNumber().Int() < 3
	})

	if b.Size() != 2 {
		t.Error(b.Size())
	}
	if b.First().AsNumber().Int() != 1 {
		t.Error(b.First())
	}
	if b.Last().AsNumber().Int() != 2 {
		t.Error(b.Last())
	}
}

func TestArrayImpl_Right(t *testing.T) {
	a := NewByInterface(1, 2, 3)
	b := a.Right(2)

	if b.Size() != 2 {
		t.Error(b.Size())
	}
	if b.First().AsNumber().Int() != 2 {
		t.Error(b.First())
	}
	if b.Last().AsNumber().Int() != 3 {
		t.Error(b.Last())
	}
}

func TestArrayImpl_RightWhile(t *testing.T) {
	a := NewByInterface(1, 2, 3)
	b := a.RightWhile(func(v es_value.Value) bool {
		return v.AsNumber().Int() > 1
	})

	if b.Size() != 2 {
		t.Error(b.Size())
	}
	if b.First().AsNumber().Int() != 2 {
		t.Error(b.First())
	}
	if b.Last().AsNumber().Int() != 3 {
		t.Error(b.Last())
	}
}

func TestArrayImpl_Count(t *testing.T) {
	a := NewByInterface(1, 2, 3)
	c := a.Count(func(v es_value.Value) bool {
		return v.AsNumber().Int() > 1
	})
	if c != 2 {
		t.Error(c)
	}
}

func TestArrayImpl_Entries(t *testing.T) {
	a := NewByInterface(1, 2, 3)
	if x := a.Entries()[0].AsNumber().Int(); x != 1 {
		t.Error(x)
	}
	if x := a.Entries()[1].AsNumber().Int(); x != 2 {
		t.Error(x)
	}
	if x := a.Entries()[2].AsNumber().Int(); x != 3 {
		t.Error(x)
	}
}

func TestArrayImpl_Unique(t *testing.T) {
	a := NewByInterface(1, 2, 3, 2, 1)
	b := a.Unique()
	if b.Size() != 3 {
		t.Error(b.Size())
	}
	c := b.Sort()
	if x := c.Entries()[0].AsNumber().Int(); x != 1 {
		t.Error(x)
	}
	if x := c.Entries()[1].AsNumber().Int(); x != 2 {
		t.Error(x)
	}
	if x := c.Entries()[2].AsNumber().Int(); x != 3 {
		t.Error(x)
	}
}

func TestArrayImpl_Append(t *testing.T) {
	a := NewByInterface(1, 2, 3)
	b := a.Append(NewByInterface(4, 5))
	if b.Size() != 5 {
		t.Error(b.Size())
	}
	if x := b.Entries()[0].AsNumber().Int(); x != 1 {
		t.Error(x)
	}
	if x := b.Entries()[1].AsNumber().Int(); x != 2 {
		t.Error(x)
	}
	if x := b.Entries()[2].AsNumber().Int(); x != 3 {
		t.Error(x)
	}
	if x := b.Entries()[3].AsNumber().Int(); x != 4 {
		t.Error(x)
	}
	if x := b.Entries()[4].AsNumber().Int(); x != 5 {
		t.Error(x)
	}
}

func TestArrayImpl_Intersection(t *testing.T) {
	a := NewByInterface(1, 2, 3)
	b := a.Intersection(NewByInterface(2, 3, 4))
	if b.Size() != 2 {
		t.Error(b.Size())
	}
	c := b.Sort()
	if x := c.Entries()[0].AsNumber().Int(); x != 2 {
		t.Error(x)
	}
	if x := c.Entries()[1].AsNumber().Int(); x != 3 {
		t.Error(x)
	}
}

func TestArrayImpl_Union(t *testing.T) {
	a := NewByInterface(1, 2, 3)
	b := a.Union(NewByInterface(3, 8, 9))
	if b.Size() != 5 {
		t.Error(b.Size())
	}
	c := b.Sort()
	if x := c.Entries()[0].AsNumber().Int(); x != 1 {
		t.Error(x)
	}
	if x := c.Entries()[1].AsNumber().Int(); x != 2 {
		t.Error(x)
	}
	if x := c.Entries()[2].AsNumber().Int(); x != 3 {
		t.Error(x)
	}
	if x := c.Entries()[3].AsNumber().Int(); x != 8 {
		t.Error(x)
	}
	if x := c.Entries()[4].AsNumber().Int(); x != 9 {
		t.Error(x)
	}
}

func TestArrayImpl_Sort(t *testing.T) {
	a := NewByInterface(3, 1, 2)
	b := a.Sort()
	if x := b.Entries()[0].AsNumber().Int(); x != 1 {
		t.Error(x)
	}
	if x := b.Entries()[1].AsNumber().Int(); x != 2 {
		t.Error(x)
	}
	if x := b.Entries()[2].AsNumber().Int(); x != 3 {
		t.Error(x)
	}
}

func TestArrayImpl_AsStringArray(t *testing.T) {
	a := NewByInterface(3, 1, 2)
	b := a.AsStringArray()
	if x := b[0]; x != "3" {
		t.Error(x)
	}
	if x := b[1]; x != "1" {
		t.Error(x)
	}
	if x := b[2]; x != "2" {
		t.Error(x)
	}
}

func TestArrayImpl_AsNumberArray(t *testing.T) {
	a := NewByInterface("3", "1", "2")
	b := a.AsNumberArray()
	if x := b[0]; x.Int() != 3 {
		t.Error(x)
	}
	if x := b[1]; x.Int() != 1 {
		t.Error(x)
	}
	if x := b[2]; x.Int() != 2 {
		t.Error(x)
	}
}

func TestArrayImpl_AsInterfaceArray(t *testing.T) {
	a := NewByInterface("3", "1", "2")
	b := a.AsInterfaceArray()
	if x, ok := b[0].(string); x != "3" || !ok {
		t.Error(x)
	}
	if x, ok := b[1].(string); x != "1" || !ok {
		t.Error(x)
	}
	if x, ok := b[2].(string); x != "2" || !ok {
		t.Error(x)
	}
}

func TestArrayImpl_HashMap(t *testing.T) {
	a := NewByInterface("3", "1", "2")
	b := a.HashMap()
	if x, ok := b["1"]; x.String() != "1" || !ok {
		t.Error(x)
	}
	if x, ok := b["2"]; x.String() != "2" || !ok {
		t.Error(x)
	}
	if x, ok := b["3"]; x.String() != "3" || !ok {
		t.Error(x)
	}
}

func TestArrayImpl_Map(t *testing.T) {
	a := NewByInterface("3", "1", "2")
	b := a.Map(func(v es_value.Value) es_value.Value {
		return es_value.New(v.String() + "E")
	})
	c := b.AsStringArray()
	if x := c[0]; x != "3E" {
		t.Error(x)
	}
	if x := c[1]; x != "1E" {
		t.Error(x)
	}
	if x := c[2]; x != "2E" {
		t.Error(x)
	}
}

func TestArrayImpl_Each(t *testing.T) {
	a := NewByInterface("3", "1", "2")
	a.Each(func(v es_value.Value) {
	})
}
