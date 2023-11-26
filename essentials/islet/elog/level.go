package elog

type Level string

const (
	Trace Level = "trace"
	Stat  Level = "stat"
	Debug Level = "debug"
	Info  Level = "info"
	Warn  Level = "warn"
	Error Level = "error"
)
