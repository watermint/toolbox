package queue

import (
	"github.com/watermint/toolbox/essentials/log/esl"
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
	ctl := MockControl{
		Logger: esl.Default().With(esl.Bool("FromContext", true)),
	}
	conn := MockConn{
		peerName: "default",
	}

	// struct ptr
	{
		queue := NewQueue(func(w *WorkData, ctl MockControl, mockConn MockConn) {
			ctl.Log().Info("UserId", esl.String("userId", w.UserId), esl.String("peerName", mockConn.PeerName()))
		}, ctl, conn)
		queue.Enqueue(&WorkData{
			UserId: "U001",
		})
		queue.Dequeue()
	}

	// struct
	{
		queue := NewQueue(func(w WorkData, ctl MockControl) {
			ctl.Log().Info("UserId", esl.String("userId", w.UserId))
		}, ctl)
		queue.Enqueue(WorkData{
			UserId: "U002",
		})
		queue.Dequeue()
	}

	// plain string with error return
	{
		queue := NewQueue(func(userId string, ctl MockControl) error {
			ctl.Log().Info("UserId", esl.String("userId", userId))
			return nil
		}, ctl)
		queue.Enqueue("U003")
		queue.Dequeue()
	}
}
