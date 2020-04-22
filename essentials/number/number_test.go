package number

import "testing"

func TestNew(t *testing.T) {
	// int
	{
		i := New(123)
		if !i.IsInt() || !i.IsValid() || i.IsNaN() || i.Int() != 123 || i.String() != "123" {
			t.Error("invalid")
		}
	}
	// int8
	{
		i := New(int8(123))
		if !i.IsInt() || !i.IsValid() || i.IsNaN() || i.Int() != 123 || i.String() != "123" {
			t.Error("invalid")
		}
	}
	// int16
	{
		i := New(int16(123))
		if !i.IsInt() || !i.IsValid() || i.IsNaN() || i.Int() != 123 || i.String() != "123" {
			t.Error("invalid")
		}
	}
	// int32
	{
		i := New(int32(123))
		if !i.IsInt() || !i.IsValid() || i.IsNaN() || i.Int() != 123 || i.String() != "123" {
			t.Error("invalid")
		}
	}
	// int64
	{
		i := New(int64(123))
		if !i.IsInt() || !i.IsValid() || i.IsNaN() || i.Int() != 123 || i.String() != "123" {
			t.Error("invalid")
		}
	}

	// uint8
	{
		i := New(uint8(123))
		if !i.IsInt() || !i.IsValid() || i.IsNaN() || i.Int() != 123 || i.String() != "123" {
			t.Error("invalid")
		}
	}
	// uint16
	{
		i := New(uint16(123))
		if !i.IsInt() || !i.IsValid() || i.IsNaN() || i.Int() != 123 || i.String() != "123" {
			t.Error("invalid")
		}
	}
	// uint32
	{
		i := New(uint32(123))
		if !i.IsInt() || !i.IsValid() || i.IsNaN() || i.Int() != 123 || i.String() != "123" {
			t.Error("invalid")
		}
	}
	// uint64: casted into float64
	{
		i := New(uint64(123))
		if i.IsInt() || !i.IsValid() || i.IsNaN() || i.Int() != 123 || i.String() != "123" {
			t.Error("invalid")
		}
	}

	// string, look like a number
	{
		i := New("123")
		if !i.IsInt() || !i.IsValid() || i.IsNaN() || i.Int() != 123 || i.String() != "123" {
			t.Error("invalid")
		}
	}
	// string, look like a number, with space
	{
		i := New("  123  ")
		if !i.IsInt() || !i.IsValid() || i.IsNaN() || i.Int() != 123 || i.String() != "123" {
			t.Error("invalid")
		}
	}
	// string, look like a float number, with space
	{
		i := New("  123.45  ")
		if i.IsInt() || !i.IsValid() || i.IsNaN() || i.Int() != 123 || i.Float64() != 123.45 || i.String() != "123.45" {
			t.Error("invalid")
		}
	}
	// string, not a number
	{
		i := New("one hundred")
		if i.IsInt() || i.IsValid() || i.IsNaN() || i.Int() != 0 || i.String() != "invalid" {
			t.Error("invalid")
		}
	}
	// number
	{
		i := New("123")
		j := New(i)
		if !j.IsInt() || !j.IsValid() || j.IsNaN() || j.Int() != 123 || j.String() != "123" {
			t.Error("invalid")
		}
	}
}
