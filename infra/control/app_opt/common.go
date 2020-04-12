package app_opt

import (
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"runtime"
)

const (
	OutputNone     = "none"
	OutputText     = "text"
	OutputMarkdown = "markdown"
	OutputJson     = "json"
)

type CommonOpts struct {
	// Automatically open the artifact folder, after successful execution
	AutoOpen bool

	// Limit bandwidth to downloading/uploading contents
	BandwidthKb int

	// Set concurrency of worker execution
	Concurrency int

	// Enable debug mode
	Debug bool

	// Enable low memory mode
	LowMemory bool

	// Set output format
	Output string

	// Explicitly set proxy the hostname and the port number
	Proxy mo_string.OptionalString

	// Quiet mode
	Quiet bool

	// Do not store token in the file
	Secure bool

	// Specify workspace path
	Workspace mo_string.OptionalString
}

func (z *CommonOpts) Preset() {
	z.AutoOpen = false
	z.BandwidthKb = 0
	z.Concurrency = runtime.NumCPU()
	z.Debug = false
	z.LowMemory = false
	z.Output = "text"
	z.Proxy = mo_string.NewOptional("")
	z.Quiet = false
	z.Secure = false
	z.Workspace = mo_string.NewOptional("")
}
