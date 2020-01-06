package sv_file_history

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
)

type History interface {
	Revisions(path mo_path.DropboxPath) (entries []mo_file.Entry, isDeleted bool, serverDeleted string, err error)

	Restore(path mo_path.DropboxPath, revision string) (entry mo_file.Entry, err error)
}
