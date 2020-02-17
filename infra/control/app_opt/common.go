package app_opt

import (
	"runtime"
)

type CommonOpts struct {
	Workspace   string
	Debug       bool
	Proxy       string
	Quiet       bool
	Secure      bool
	Concurrency int
	BandwidthKb int
	LowMemory   bool
	AutoOpen    bool
}

func (z *CommonOpts) Preset() {
	z.Workspace = ""
	z.Debug = false
	z.Proxy = ""
	z.Quiet = false
	z.Secure = false
	z.BandwidthKb = 0
	z.Concurrency = runtime.NumCPU()
	z.LowMemory = false
	z.AutoOpen = false
}
