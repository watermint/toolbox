package ut_io

import "testing"

func TestNewDefaultOut(t *testing.T) {
	{
		w := NewDefaultOut(true)
		w.Write([]byte("hello"))
		w.Close()
	}
	{
		w := NewDefaultOut(false)
		w.Write([]byte("Hello"))
		w.Close()
	}
}
