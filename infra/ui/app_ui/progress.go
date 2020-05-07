package app_ui

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"sync"
	"time"
)

const (
	progressThreshold       = 1000
	longRunningProgressSpan = 5 * time.Second
)

type MsgProgress struct {
	Progress app_msg.Message
}

var (
	progressCounter sync.Map
	progressMutex   sync.Mutex
	longRunProgress sync.Map
	MProgress       = app_msg.Apply(&MsgProgress{}).(*MsgProgress)
)

func ShowLongRunningProgress(ui UI, seed string, msg app_msg.Message) {
	var last time.Time
	t, ok := longRunProgress.Load(seed)
	if ok {
		last = t.(time.Time)
		if last.Add(longRunningProgressSpan).Before(time.Now()) {
			showLongRunningProgress(ui, seed, msg)
		}
	} else {
		longRunProgress.Store(seed, time.Now())
	}
}

func showLongRunningProgress(ui UI, seed string, msg app_msg.Message) {
	ui.Progress(msg.With("Time", time.Now().Format(time.RFC3339)))
	longRunProgress.Store(seed, time.Now())
}

func ShowProgress(ui UI) {
	ShowProgressWithMessage(ui, MProgress.Progress)
}

func ShowProgressWithMessage(ui UI, msg app_msg.Message) {
	uid := ui.Id()
	progressMutex.Lock()
	v, ok := progressCounter.Load(uid)
	if !ok {
		v = 1
	} else {
		v = v.(int) + 1
	}
	progressCounter.Store(uid, v)
	progressMutex.Unlock()

	vi := v.(int)
	if vi%progressThreshold == 0 {
		ui.Progress(msg.With("Counter", vi))
	}
}
