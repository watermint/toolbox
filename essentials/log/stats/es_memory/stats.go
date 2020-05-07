package es_memory

import (
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_shutdown"
	"runtime"
	"time"
)

const (
	reportInterval = 5 * 1000 * time.Millisecond
)

func reportLoop(t *time.Ticker, l es_log.Logger) {
	for n := range t.C {
		_ = n.Unix()
		DumpMemStats(l)
	}
}

func LaunchReporting(l es_log.Logger) {
	t := time.NewTicker(reportInterval)
	go reportLoop(t, l)
	app_shutdown.AddShutdownHook(func() {
		t.Stop()
	})
}

func DumpMemStats(l es_log.Logger) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	l.Debug("Sys", es_log.Uint64("Sys", mem.Sys), es_log.Uint64("OtherSys", mem.OtherSys))
	l.Debug("Heap stats",
		es_log.Uint64("TotalAlloc", mem.TotalAlloc),
		es_log.Uint64("HeapAlloc", mem.HeapAlloc),
		es_log.Uint64("HeapSys", mem.HeapSys),
		es_log.Uint64("HeapInuse", mem.HeapInuse),
		es_log.Uint64("HeapReleased", mem.HeapReleased),
		es_log.Uint64("Mallocs", mem.Mallocs),
		es_log.Uint64("Free", mem.Frees),
		es_log.Uint64("Live", mem.Mallocs-mem.Frees),
	)
	l.Debug("Stack stats",
		es_log.Uint64("StackInUse", mem.StackInuse),
		es_log.Uint64("StackSys", mem.StackSys),
	)
	l.Debug("GC stats",
		es_log.Uint64("GCSys", mem.GCSys),
		es_log.Uint64("NextGC", mem.NextGC),
		es_log.Uint64("LastGC", mem.LastGC),
		es_log.Uint64("PauseTotalNS", mem.PauseTotalNs),
		es_log.Uint32("NumGC", mem.NumGC),
		es_log.Uint32("NumForcedGC", mem.NumForcedGC),
	)
}
