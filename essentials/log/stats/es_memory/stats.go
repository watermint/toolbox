package es_memory

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_shutdown"
	"runtime"
	"time"
)

const (
	reportInterval = 5 * 1000 * time.Millisecond
)

func reportLoop(t *time.Ticker, l esl.Logger) {
	for n := range t.C {
		_ = n.Unix()
		DumpMemStats(l)
	}
}

func LaunchReporting(l esl.Logger) {
	t := time.NewTicker(reportInterval)
	go reportLoop(t, l)
	app_shutdown.AddShutdownHook(func() {
		t.Stop()
	})
}

func DumpMemStats(l esl.Logger) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	l.Debug("Sys", esl.Uint64("Sys", mem.Sys), esl.Uint64("OtherSys", mem.OtherSys))
	l.Debug("Heap stats",
		esl.Uint64("TotalAlloc", mem.TotalAlloc),
		esl.Uint64("HeapAlloc", mem.HeapAlloc),
		esl.Uint64("HeapSys", mem.HeapSys),
		esl.Uint64("HeapInuse", mem.HeapInuse),
		esl.Uint64("HeapReleased", mem.HeapReleased),
		esl.Uint64("Mallocs", mem.Mallocs),
		esl.Uint64("Free", mem.Frees),
		esl.Uint64("Live", mem.Mallocs-mem.Frees),
	)
	l.Debug("Stack stats",
		esl.Uint64("StackInUse", mem.StackInuse),
		esl.Uint64("StackSys", mem.StackSys),
	)
	l.Debug("GC stats",
		esl.Uint64("GCSys", mem.GCSys),
		esl.Uint64("NextGC", mem.NextGC),
		esl.Uint64("LastGC", mem.LastGC),
		esl.Uint64("PauseTotalNS", mem.PauseTotalNs),
		esl.Uint32("NumGC", mem.NumGC),
		esl.Uint32("NumForcedGC", mem.NumForcedGC),
	)
}
