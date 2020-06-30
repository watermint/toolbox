package es_close

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

func New(w io.WriteCloser) io.WriteCloser {
	return &WriteCloser{
		w: w,
	}
}

type WriteCloser struct {
	w io.WriteCloser
}

func (z *WriteCloser) Write(p []byte) (n int, err error) {
	// Catch exceptions in case the writer closed in race condition
	defer func() {
		r := recover()
		if r != nil {
			switch e := r.(type) {
			case error:
				err = e
			default:
				err = errors.New(fmt.Sprintf("%v", e))
			}
		}
	}()

	if z.w != nil {
		n, err = z.w.Write(p)
	} else {
		// report data all written to caller to conform interface definition.
		n = len(p)
	}
	return
}

func (z *WriteCloser) Close() (err error) {
	defer func() {
		r := recover()
		if r != nil {
			switch e := r.(type) {
			case error:
				err = e
			default:
				err = errors.New(fmt.Sprintf("%v", e))
			}
		}
	}()

	if z.w != nil {
		err = z.w.Close()
		z.w = nil
	}
	return
}

func NewNopWriteCloser(w io.Writer) io.WriteCloser {
	return &NopWriteCloser{
		w: w,
	}
}

type NopWriteCloser struct {
	w io.Writer
}

func (z NopWriteCloser) Write(p []byte) (n int, err error) {
	return z.w.Write(p)
}

func (z NopWriteCloser) Close() error {
	return nil
}

func NewNopCloseBuffer() NopCloseBuffer {
	return &nopCloseBufferImpl{buf: bytes.Buffer{}}
}

type NopCloseBuffer interface {
	io.WriteCloser
	Bytes() []byte
	Len() int
	String() string
}

type nopCloseBufferImpl struct {
	buf bytes.Buffer
}

func (z *nopCloseBufferImpl) Write(p []byte) (n int, err error) {
	return z.buf.Write(p)
}

func (z nopCloseBufferImpl) Close() error {
	return nil
}

func (z nopCloseBufferImpl) Bytes() []byte {
	return z.buf.Bytes()
}

func (z nopCloseBufferImpl) Len() int {
	return z.buf.Len()
}

func (z nopCloseBufferImpl) String() string {
	return z.buf.String()
}
