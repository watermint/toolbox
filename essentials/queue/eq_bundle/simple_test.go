package eq_bundle

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestSimpleImpl_BasicBehavior(t *testing.T) {
	factory := eq_pipe.NewTransientSimple(esl.Default())
	bundle := NewSimple(esl.Default(), factory)

	d1 := NewData("", []byte("D00-001"))

	// ensure the queue is empty
	if sizes, total := bundle.Size(); total != 0 {
		t.Error(sizes, total)
	}

	// enqueue
	bundle.Enqueue(d1)

	// ensure queued
	if sizes, total := bundle.Size(); total != 1 {
		t.Error(sizes, total)
	}

	// fetch and compare
	if d, found := bundle.Fetch(); !found {
		t.Error(found)
	} else if !reflect.DeepEqual(d1, d) {
		t.Error(d)
	}

	// no more data
	if d, found := bundle.Fetch(); found {
		t.Error(d, found)
	}

	// complete
	bundle.Complete(d1)

	// close bundle
	bundle.Close()
}

func TestSimpleImpl_Concurrent(t *testing.T) {
	factory := eq_pipe.NewTransientSimple(esl.Default())
	bundle := NewSimple(esl.Default(), factory)

	wgPush := sync.WaitGroup{}
	wgFetch := sync.WaitGroup{}

	waitUnit := 50 * time.Microsecond

	pusher := func(id int) {
		l := esl.Default().With(esl.Int("Pusher", id))
		for j := 0; j < 100; j++ {
			b := fmt.Sprintf("D%02d-%03d", id, j)
			d := NewData(fmt.Sprintf("B%02d", id), []byte(b))
			l.Debug("Push", esl.Any("packet", d), esl.ByteString("data", d.D))
			bundle.Enqueue(d)
			time.Sleep(waitUnit)
		}
		wgPush.Done()
	}
	fetcher := func(id int) {
		l := esl.Default().With(esl.Int("Fetcher", id))
		timeout := time.Now().Add(10 * waitUnit)
		for {
			d, found := bundle.Fetch()
			if !found {
				if timeout.Before(time.Now()) {
					l.Debug("Timeout, return")
					wgFetch.Done()
					return

				} else {
					l.Debug("not found, wait")
					time.Sleep(waitUnit)
				}
			} else {
				l.Info("Process", esl.Any("packet", d), esl.ByteString("data", d.D))
			}
		}
	}

	// Enqueue
	for i := 0; i < 10; i++ {
		wgPush.Add(1)
		go pusher(i)
	}

	// Dequeue
	for i := 0; i < 4; i++ {
		wgFetch.Add(1)
		go fetcher(i)
	}

	wgPush.Wait()
	wgFetch.Wait()

	// close bundle
	bundle.Close()
}
