package app_opt

import "runtime"

type CommonOpts struct {
	Workspace   string
	Debug       bool
	Proxy       string
	Quiet       bool
	Secure      bool
	Concurrency int
	Bandwidth   int
}

func NewDefaultCommonOpts() *CommonOpts {
	return &CommonOpts{
		Workspace:   "",
		Debug:       false,
		Proxy:       "",
		Quiet:       false,
		Secure:      false,
		Bandwidth:   0,
		Concurrency: runtime.NumCPU(),
	}
}
