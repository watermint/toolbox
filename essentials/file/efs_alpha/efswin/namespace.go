package efswin

import (
	"github.com/watermint/toolbox/essentials/file/efs_alpha"
)

type Namespace interface {
	efs_alpha.Namespace

	// Server returns name of the server. Returns empty if no server associated with the namespace.
	Server() string

	// Volume returns drive letter or volume identifier.
	// Returns empty if no volume associated with the namespace.
	// Sample: "C:" or "Volume{b75e2c83-0000-0000-0000-602f00000000}"
	Volume() string
}
