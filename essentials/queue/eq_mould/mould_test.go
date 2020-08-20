package eq_mould

import (
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"github.com/watermint/toolbox/essentials/queue/eq_progress"
	"testing"
)

type MockControl struct {
	Logger esl.Logger
}

func (z MockControl) Log() esl.Logger {
	return z.Logger
}

type MockConn struct {
	peerName string
}

func (z MockConn) PeerName() string {
	return z.peerName
}

type WorkData struct {
	UserId string `json:"user_id"`
}

func TestQueue_Dequeue(t *testing.T) {
	l := esl.Default()
	ctl := MockControl{
		Logger: l.With(esl.Bool("FromContext", true)),
	}
	conn := MockConn{
		peerName: "default",
	}
	handler := eq_progress.NewBar()
	storage := eq_bundle.NewSimple(esl.Default(), handler, eq_pipe.NewTransientSimple(esl.Default()))

	// struct ptr
	{
		f := func(w *WorkData, ctl MockControl, mockConn MockConn) {
			ctl.Log().Info("UserId", esl.String("userId", w.UserId), esl.String("peerName", mockConn.PeerName()))
		}
		mould := New("alpha", storage, f, ctl, conn)
		mould.Pour(&WorkData{
			UserId: "U001",
		})
		if d, found := storage.Fetch(); found {
			mould.Process(d)
		}
	}

	// struct
	{
		f := func(w WorkData, ctl MockControl) {
			ctl.Log().Info("UserId", esl.String("userId", w.UserId))
		}
		mould := New("alpha", storage, f, ctl)
		mould.Pour(WorkData{
			UserId: "U002",
		})
		mould.Pour(WorkData{
			UserId: "U004",
		})
		mould.Pour(WorkData{
			UserId: "U006",
		})
		for i := 0; i < 4; i++ {
			if d, found := storage.Fetch(); found {
				mould.Process(d)
			}
		}
	}

	// plain string with error return
	{
		f := func(userId string, ctl MockControl) error {
			ctl.Log().Info("UserId", esl.String("userId", userId))
			return errors.New("this is wrong")
		}
		mould := New("alpha", storage, f, ctl)
		mould.Pour("U003")
		if d, found := storage.Fetch(); found {
			mould.Process(d)
		}
	}
}

func TestMouldImpl_Batch(t *testing.T) {
	ctl := MockControl{
		Logger: esl.Default().With(esl.Bool("FromContext", true)),
	}
	storage := eq_bundle.NewSimple(esl.Default(), nil, eq_pipe.NewTransientSimple(esl.Default()))

	// struct
	{
		f := func(userId string, ctl MockControl) {
			ctl.Log().Info("UserId", esl.String("userId", userId))
		}
		mould := New("alpha", storage, f, ctl)
		b01 := mould.Batch("B01")
		b02 := mould.Batch("B02")
		b01.Pour("B01_001")
		b02.Pour("C02_101")
		b02.Pour("C02_102")
		b01.Pour("B01_002")

		for i := 0; i < 4; i++ {
			if d, found := storage.Fetch(); found {
				mould.Process(d)
			}
		}
	}
}
