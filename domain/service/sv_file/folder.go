package sv_file

import "github.com/watermint/toolbox/domain/model/mo_file"

type Folder interface {
	// `files/create_folder`
	// options: auto_rename
	Create(path mo_file.Entry) (entry mo_file.Entry, err error)
}
