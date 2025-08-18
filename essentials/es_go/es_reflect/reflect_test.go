package es_reflect

import "testing"

func TestNewInstance(t *testing.T) {
	// struct
	{
		type S struct {
			Name  string
			Value string
		}
		s := NewInstance(&S{Name: "s", Value: "v"})
		switch s0 := s.(type) {
		case *S:
			t.Log("ok" + s0.Name + s0.Value)
			if s0.Name != "" || s0.Value != "" {
				t.Error(s0.Name, s0.Value)
			}
		default:
			t.Error("invalid type")
		}
	}
}
