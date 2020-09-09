package es_sync

import "github.com/watermint/toolbox/essentials/file/es_filesystem"

type Syncer interface {
	// Sync source to target.
	Sync(source es_filesystem.Path, target es_filesystem.Path) error
}
