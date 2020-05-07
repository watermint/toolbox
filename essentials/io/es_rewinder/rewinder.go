package es_rewinder

import (
	"bytes"
	"errors"
	"github.com/watermint/toolbox/essentials/collections/es_number"
	"io"
	"math"
)

type ReadRewinder interface {
	io.Reader
	Rewind() error
	Length() int64
}

func NewReadRewinderOnMemory(c []byte) ReadRewinder {
	sz := int64(len(c))
	return &readerRewinderWithLimit{
		offset: 0,
		limit:  sz,
		length: sz,
		lr:     bytes.NewReader(c),
		r:      bytes.NewReader(c),
	}
}

func NewReadRewinder(r io.ReadSeeker, offset int64) (rr ReadRewinder, err error) {
	return NewReadRewinderWithLimit(r, offset, math.MaxInt64)
}

func NewReadRewinderWithLimit(r io.ReadSeeker, offset, limit int64) (rr ReadRewinder, err error) {
	rrwl := &readerRewinderWithLimit{
		r:      r,
		offset: offset,
		limit:  limit,
	}
	if offset < 0 {
		return nil, errors.New("negative offset")
	}
	if limit < 0 {
		return nil, errors.New("negative limit")
	}
	e, err := r.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, err
	}
	rrwl.length = es_number.Max(es_number.Min(e-offset, limit), 0).Int64()
	if err = rrwl.Rewind(); err != nil {
		return nil, err
	}
	return rrwl, nil
}

type readerRewinderWithLimit struct {
	offset int64
	limit  int64
	length int64
	lr     io.Reader
	r      io.ReadSeeker
}

func (z *readerRewinderWithLimit) Length() int64 {
	return z.length
}

func (z *readerRewinderWithLimit) Read(p []byte) (n int, err error) {
	return z.lr.Read(p)
}

func (z *readerRewinderWithLimit) Rewind() error {
	_, err := z.r.Seek(z.offset, io.SeekStart)
	if err != nil {
		return err
	}
	z.lr = io.LimitReader(z.r, z.limit)
	return nil
}
