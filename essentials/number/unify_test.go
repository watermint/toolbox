package number

import "testing"

func TestUnify(t *testing.T) {
	// mixed
	{
		g1 := []interface{}{
			1,
			10,
			10.4,
			int64(23),
		}
		f1 := Unify(g1...)
		if f1[0].Int() != 1 {
			t.Error(f1[0].String())
		}
		if f1[1].Int() != 10 {
			t.Error(f1[1].String())
		}
		if f1[2].Int() != 10 {
			t.Error(f1[2].String())
		}
		for _, f := range f1 {
			if !f.IsFloat() {
				t.Error(f.IsFloat())
			}
			if f.IsInt() {
				t.Error(f.IsInt())
			}
		}
	}

	// integer only
	{
		g2 := []interface{}{
			1,
			10,
			20,
			int64(30),
		}
		f2 := Unify(g2...)
		if f2[0].Int() != 1 {
			t.Error(f2[0].String())
		}
		if f2[1].Int() != 10 {
			t.Error(f2[1].String())
		}
		if f2[2].Int() != 20 {
			t.Error(f2[2].String())
		}
		for _, f := range f2 {
			if f.IsFloat() {
				t.Error(f.IsFloat())
			}
			if !f.IsInt() {
				t.Error(f.IsInt())
			}
		}
	}

	// float only
	{
		g3 := []interface{}{
			1.1,
			10.2,
			20.3,
			30.4,
		}
		f3 := Unify(g3...)
		if f3[0].Float64() != 1.1 {
			t.Error(f3[0].String())
		}
		if f3[1].Float64() != 10.2 {
			t.Error(f3[1].String())
		}
		if f3[2].Float64() != 20.3 {
			t.Error(f3[2].String())
		}
		for _, f := range f3 {
			if !f.IsFloat() {
				t.Error(f.IsFloat())
			}
			if f.IsInt() {
				t.Error(f.IsInt())
			}
		}
	}

	// mixed, with invalid
	{
		g4 := []interface{}{
			struct{}{},
			1,
			10,
			10.4,
			int64(23),
		}
		f4 := Unify(g4...)
		if f4[0].Int() != 1 {
			t.Error(f4[0].String())
		}
		if f4[1].Int() != 10 {
			t.Error(f4[1].String())
		}
		if f4[2].Int() != 10 {
			t.Error(f4[2].String())
		}
		for _, f := range f4 {
			if !f.IsFloat() {
				t.Error(f.IsFloat())
			}
			if f.IsInt() {
				t.Error(f.IsInt())
			}
		}
	}
}
