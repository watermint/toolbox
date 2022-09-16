package file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
)

func DeleteRecursively(client dbx_client.Client, path mo_path.DropboxPath, beforeDelete func(path mo_path.DropboxPath)) error {
	beforeDelete(path)
	_, err := sv_file.NewFiles(client).Remove(path)
	if err == nil {
		return nil
	}
	de := dbx_error.NewErrors(err)
	if de.IsTooManyFiles() {
		entries, err := sv_file.NewFiles(client).List(path)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			if f, ok := entry.File(); ok {
				if err := DeleteRecursively(client, f.Path(), beforeDelete); err != nil {
					return err
				}
			}
			if f, ok := entry.Folder(); ok {
				if err := DeleteRecursively(client, f.Path(), beforeDelete); err != nil {
					return err
				}
			}
		}
		return DeleteRecursively(client, path, beforeDelete)
	} else {
		return err
	}
}
