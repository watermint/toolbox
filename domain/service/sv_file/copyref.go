package sv_file

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
)

type CopyRef interface {
	// `files/copy_reference/get`
	Resolve(path mo_path.Path) (entry mo_file.Entry, ref, expires string, err error)

	// `files/copy_reference/save`
	Save(path mo_path.Path, ref string) (entry mo_file.Entry, err error)
}
