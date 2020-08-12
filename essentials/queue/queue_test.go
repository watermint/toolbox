package queue

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe_preserve"
	"github.com/watermint/toolbox/essentials/queue/eq_pump"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"go.uber.org/atomic"
	"golang.org/x/sync/syncmap"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestQueueImpl_Enqueue(t *testing.T) {
	l := esl.Default()
	processor := func(userId string) {
		l.Info("Greet", esl.String("UserId", userId))
	}
	factory := eq_pipe.NewTransientSimple(l)
	queue := New(l, 2, factory, processor)

	queue.Enqueue("U-001")

	b1 := queue.Batch("B01")
	b1.Enqueue("UB-001")
	b1.Enqueue("UB-002")

	queue.Enqueue("U-002")

	queue.Wait()
}

func TestQueueImpl_Suspend_Transient(t *testing.T) {
	waitUnit := 50 * time.Millisecond

	l := esl.Default()
	processor := func(userId string) {
		time.Sleep(waitUnit)
		l.Info("Greet", esl.String("UserId", userId))
	}
	factory := eq_pipe.NewTransientSimple(l)
	queue := New(l, 1, factory, processor)

	queue.Enqueue("U-001")
	queue.Enqueue("U-002")
	queue.Enqueue("U-003")
	queue.Enqueue("U-004")

	session, err := queue.Suspend()
	if err != eq_pipe_preserve.ErrorSessionIsNotAvailable {
		t.Error(session, err)
	}
}

func TestQueueImpl_Suspend(t *testing.T) {
	qt_file.TestWithTestFolder(t, "suspend", false, func(path string) {
		l := esl.Default()
		dataSeq := make([]string, 0)
		for i := 0; i < 10; i++ {
			dataSeq = append(dataSeq, fmt.Sprintf("U-%03d", i))
		}
		proceed := atomic.NewInt32(0)
		proceedData := syncmap.Map{}

		processor := func(userId string) {
			time.Sleep(eq_pump.PollInterval)
			l.Info("Greet", esl.String("UserId", userId))
			proceed.Inc()
			proceedData.Store(userId, true)
		}
		preserver := eq_pipe_preserve.NewFactory(l, path)
		factory := eq_pipe.NewSimple(l, preserver)
		queue := New(l, 1, factory, processor)

		for _, d := range dataSeq {
			queue.Enqueue(d)
		}

		// Wait single interval for process at least one data
		for {
			if x := proceed.Load(); x < 1 {
				l.Debug("Wait for process")
				time.Sleep(eq_pump.PollInterval)
			} else if len(dataSeq) <= int(x) {
				l.Error("Data is not proceed or proceed all data", esl.Int32("x", x))
				t.Error(x)
			} else {
				l.Debug("At least one message proceed", esl.Int32("x", x))
				break
			}
		}
		if x := proceed.Load(); x < 1 || len(dataSeq) <= int(x) {
			l.Error("Data is not proceed or proceed all data", esl.Int32("x", x))
			t.Error(x)
		}

		session, err := queue.Suspend()
		if err != nil {
			t.Error(session, err)
		}
		proceedAtSuspend := proceed.Load()

		// Wait 2 * interval to ensure workers will not consume any data
		time.Sleep(2 * eq_pump.PollInterval)
		if x := proceed.Load(); proceedAtSuspend != x {
			t.Error(x, proceedAtSuspend)
		}

		// Restore
		restored, err := Restore(l, 1, factory, session, processor)
		if err != nil {
			t.Error(err)
		}
		restored.Wait()

		// Verify data
		proceedDataSeq := make([]string, 0)
		proceedData.Range(func(key, value interface{}) bool {
			proceedDataSeq = append(proceedDataSeq, key.(string))
			return true
		})

		sort.Strings(dataSeq)
		sort.Strings(proceedDataSeq)

		if !reflect.DeepEqual(dataSeq, proceedDataSeq) {
			l.Error("Invalid data sequence",
				esl.Strings("dataSeq", dataSeq),
				esl.Strings("processedDataSeq", proceedDataSeq),
			)
			fmt.Println(cmp.Diff(dataSeq, proceedDataSeq))
			t.Error(dataSeq, proceedDataSeq)
		}
	})
}
