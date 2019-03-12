package sv_file

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
)

type Url interface {
	Save(path mo_path.Path, url string) (entry mo_file.Entry, err error)
}
