package sv_file_content

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"io"
)

type Download interface {
	Download(path mo_path.Path, out io.Writer) (entry mo_file.Entry, err error)
	Url(path mo_path.Path) (entry mo_file.Entry, url string, err error)
}
