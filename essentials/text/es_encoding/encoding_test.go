package es_encoding

import "testing"

func TestSelectEncoding(t *testing.T) {
	for _, n := range EncodingNames {
		if e := SelectEncoding(n); e == nil {
			t.Error(n)
		}
	}

	if e := SelectEncoding("invalid"); e != nil {
		t.Error(e)
	}
}
