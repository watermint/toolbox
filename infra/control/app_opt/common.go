package app_opt

import "runtime"

type CommonOpts struct {
	Workspace   string
	Debug       bool
	Proxy       string
	Quiet       bool
	Secure      bool
	Concurrency int
}

func NewDefaultCommonOpts() *CommonOpts {
	return &CommonOpts{
		Workspace:   "",
		Debug:       false,
		Proxy:       "",
		Quiet:       false,
		Secure:      false,
		Concurrency: runtime.NumCPU(),
	}
}
