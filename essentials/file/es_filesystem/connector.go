package es_filesystem

import "github.com/watermint/toolbox/essentials/queue/eq_queue"

type CopyPair struct {
	Source Entry
	Target Path
}

func NewCopyPair(source Entry, target Path) CopyPair {
	return CopyPair{
		Source: source,
		Target: target,
	}
}

type Connector interface {
	// Copy source to target system. Target path must include file name.
	// Connector callbacks onSuccess or onFailure to tell an result.
	// Copy operation may block to wait I/O.
	Copy(source Entry,
		target Path,
		onSuccess func(pair CopyPair, copied Entry),
		onFailure func(pair CopyPair, err FileSystemError))

	// Start up the connector
	Startup(qd eq_queue.Definition) (err FileSystemError)

	// Clean up connector
	Shutdown() (err FileSystemError)
}
