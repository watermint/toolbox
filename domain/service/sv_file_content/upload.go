package sv_file_content

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"io"
	"os"
	"time"
)

type Upload interface {
	// options: auto_rename, mute, client_modified, strict_conflict, mode
	Upload(path mo_path.Path, in io.Reader, clientModified time.Time) (entry mo_file.Entry, err error)
	UploadFile(path mo_path.Path, file *os.File) (entry mo_file.Entry, err error)
	// options: mode, auto_rename, mute, strict_conflict, duration
	Url(path mo_path.Path) (url string, err error)
}
