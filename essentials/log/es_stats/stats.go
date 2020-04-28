package es_stats

import (
	"go.uber.org/zap"
	"runtime"
	"time"
)

const (
	reportInterval = 5 * 1000 * time.Millisecond
)

func LaunchReporting(l *zap.Logger) {
	go func() {
		for {
			time.Sleep(reportInterval)
			DumpMemStats(l)
		}
	}()
}

func DumpMemStats(l *zap.Logger) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	l.Debug("Sys", zap.Uint64("Sys", mem.Sys), zap.Uint64("OtherSys", mem.OtherSys))
	l.Debug("Heap stats",
		zap.Uint64("TotalAlloc", mem.TotalAlloc),
		zap.Uint64("HeapAlloc", mem.HeapAlloc),
		zap.Uint64("HeapSys", mem.HeapSys),
		zap.Uint64("HeapInuse", mem.HeapInuse),
		zap.Uint64("HeapReleased", mem.HeapReleased),
		zap.Uint64("Mallocs", mem.Mallocs),
		zap.Uint64("Free", mem.Frees),
		zap.Uint64("Live", mem.Mallocs-mem.Frees),
	)
	l.Debug("Stack stats",
		zap.Uint64("StackInUse", mem.StackInuse),
		zap.Uint64("StackSys", mem.StackSys),
	)
	l.Debug("GC stats",
		zap.Uint64("GCSys", mem.GCSys),
		zap.Uint64("NextGC", mem.NextGC),
		zap.Uint64("LastGC", mem.LastGC),
		zap.Uint64("PauseTotalNS", mem.PauseTotalNs),
		zap.Uint32("NumGC", mem.NumGC),
		zap.Uint32("NumForcedGC", mem.NumForcedGC),
	)
}
