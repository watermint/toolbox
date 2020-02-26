package app_ui

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"sync"
)

const (
	progressThreshold = 1000
)

type MsgProgress struct {
	Progress app_msg.Message
}

var (
	progressCounter sync.Map
	progressMutex   sync.Mutex
	MProgress       = app_msg.Apply(&MsgProgress{}).(*MsgProgress)
)

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
