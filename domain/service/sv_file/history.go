package sv_file

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
)

type History interface {
	Revisions(path mo_path.Path) (entries []mo_file.Entry, isDeleted bool, serverDeleted string, err error)

	Restore(path mo_path.Path, revision string) (entry mo_file.Entry, err error)
}
