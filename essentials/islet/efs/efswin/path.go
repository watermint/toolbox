package efswin

import "github.com/watermint/toolbox/essentials/islet/efs"

// Path for Windows
// https://docs.microsoft.com/en-us/dotnet/standard/io/file-path-formats
type Path interface {
	efs.Path

	// IsUNC returns true if a path format is UNC format.
	IsUNC() bool
}
