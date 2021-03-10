package es_block

import (
	"bytes"
	"io"
	"testing"
)

func NewSmallReader(r io.Reader, size int) io.Reader {
	return &smallReader{
		r:    r,
		size: size,
	}
}

type smallReader struct {
	r    io.Reader
	size int
}

func (z smallReader) Read(p []byte) (n int, err error) {
	buf := make([]byte, z.size)
	n, err = z.r.Read(buf)
	copy(p, buf)
	return
}

func TestReaderImpl_ReadBlock(t *testing.T) {
	testData := []byte("0123456789")

	{
		r := NewReader(bytes.NewReader(testData), 5)
		{
			b, f, e := r.ReadBlock()
			if string(b) != "01234" || f || e != nil {
				t.Error(string(b), f, e)
			}
		}

		{
			b, f, e := r.ReadBlock()
			if string(b) != "56789" || f || e != nil {
				t.Error(string(b), f, e)
			}
		}

		{
			b, f, e := r.ReadBlock()
			if string(b) != "" || !f || e != nil {
				t.Error(string(b), f, e)
			}
		}
	}

	{
		r := NewReader(bytes.NewReader(testData), 15)
		{
			b, f, e := r.ReadBlock()
			if string(b) != "0123456789" || !f || e != nil {
				t.Error(string(b), f, e)
			}
		}
	}

	{
		r := NewReader(NewSmallReader(bytes.NewReader(testData), 4), 8)
		{
			b, f, e := r.ReadBlock()
			if string(b) != "01234567" || f || e != nil {
				t.Error(string(b), f, e)
			}
		}

		{
			b, f, e := r.ReadBlock()
			if string(b) != "89" || !f || e != nil {
				t.Error(string(b), f, e)
			}
		}
	}

}
