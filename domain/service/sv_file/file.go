package sv_file

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"io"
	"os"
	"time"
)

type Files interface {
	Resolve(path mo_path.Path) (entry mo_file.Entry, err error)
	List(path mo_path.Path) (entries []mo_file.Entry, err error)

	Delete(path mo_path.Path) (entry mo_file.Entry, err error)
	DeleteWithRevision(path mo_path.Path, revision string) (entry mo_file.Entry, err error)

	PermDelete(path mo_path.Path) (err error)
	PermDeleteWithRevision(path mo_path.Path, revision string) (err error)
}

type Upload interface {
	// options: auto_rename, mute, client_modified, strict_conflict, mode
	Upload(path mo_path.Path, in io.Reader, clientModified time.Time) (entry mo_file.Entry, err error)
	UploadFile(path mo_path.Path, file *os.File) (entry mo_file.Entry, err error)
	// options: mode, auto_rename, mute, strict_conflict, duration
	Url(path mo_path.Path) (url string, err error)
}

type Download interface {
	Download(path mo_path.Path, out io.Writer) (entry mo_file.Entry, err error)
	Url(path mo_path.Path) (entry mo_file.Entry, url string, err error)
}

type History interface {
	Revisions(path mo_path.Path) (entries []mo_file.Entry, isDeleted bool, serverDeleted string, err error)

	Restore(path mo_path.Path, revision string) (entry mo_file.Entry, err error)
}

type Url interface {
	Save(path mo_path.Path, url string) (entry mo_file.Entry, err error)
}

type Folder interface {
	// `files/create_folder`
	// options: auto_rename
	Create(path mo_file.Entry) (entry mo_file.Entry, err error)
}

type Search interface {
	// `files/search`
	// options: max_results
	ByName(query string) (entries []mo_file.Entry, err error)

	// `files/search`
	// options: max_results
	ByContent(query string) (entries []mo_file.Entry, err error)

	// `files/search`
	// options: max_results
	Deleted(query string) (entries []mo_file.Entry, err error)
}

type CopyRef interface {
	// `files/copy_reference/get`
	Resolve(path mo_path.Path) (entry mo_file.Entry, ref, expires string, err error)

	// `files/copy_reference/save`
	Save(path mo_path.Path, ref string) (entry mo_file.Entry, err error)
}
