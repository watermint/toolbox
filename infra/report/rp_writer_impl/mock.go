package rp_writer_impl

import (
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"sync"
)

var (
	ErrorMockTheWriterIsNotReady = errors.New("the writer is not ready")
)

type MockRecord struct {
	SKU      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

func NewMock() *Mock {
	return &Mock{
		records:      make([]interface{}, 0),
		recordsMutex: sync.Mutex{},
		isClosed:     false,
		isOpened:     false,
	}
}

type Mock struct {
	name         string
	records      []interface{}
	recordsMutex sync.Mutex
	isClosed     bool
	isOpened     bool
}

func (z *Mock) Records() []interface{} {
	return z.records
}

func (z *Mock) IsOpened() bool {
	return z.isOpened
}

func (z *Mock) IsClosed() bool {
	return z.isClosed
}

func (z *Mock) Name() string {
	return ""
}

func (z *Mock) Row(r interface{}) {
	z.recordsMutex.Lock()
	defer z.recordsMutex.Unlock()

	l := esl.Default()
	if z.isClosed || !z.isOpened {
		l.Warn("The writer is not opened")
		panic(ErrorMockTheWriterIsNotReady)
	}

	l.Debug("Write record", esl.Any("record", r))
	z.records = append(z.records, r)
}

func (z *Mock) Open(ctl app_control.Control, model interface{}, opts ...rp_model.ReportOpt) error {
	z.isOpened = true
	return nil
}

func (z *Mock) Close() {
	z.isClosed = true
}
