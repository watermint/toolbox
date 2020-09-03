package uc_team_content

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
)

type TeamFolderNestedFolderScanWorker struct {
	ctl          app_control.Control
	ctx          dbx_context.Context
	metadata     kv_storage.Storage
	tree         kv_storage.Storage
	teamFolderId string
	descendants  []string
}

func (z *TeamFolderNestedFolderScanWorker) addTree(t *Tree) error {
	return z.tree.Update(func(kvs kv_kvs.Kvs) error {
		return kvs.PutJsonModel(t.NamespaceId, t)
	})
}

func (z *TeamFolderNestedFolderScanWorker) Exec() error {
	// Breadth first search for nested folders
	l := z.ctl.Log().With(esl.String("teamFolderId", z.teamFolderId))
	teamFolderName := ""
	l.Debug("lookup team folder name")
	err := z.metadata.View(func(kvs kv_kvs.Kvs) error {
		tf := &mo_sharedfolder.SharedFolder{}
		if err := kvs.GetJsonModel(z.teamFolderId, tf); err != nil {
			return err
		}
		teamFolderName = tf.Name
		return nil
	})
	if err != nil {
		return err
	}
	l = l.With(esl.String("teamFolderName", teamFolderName))

	err = z.tree.Update(func(kvs kv_kvs.Kvs) error {
		return kvs.PutJsonModel(z.teamFolderId, &Tree{
			NamespaceId:   z.teamFolderId,
			NamespaceName: teamFolderName,
			RelativePath:  mo_path.NewDropboxPath(teamFolderName).Path(),
		})
	})
	if err != nil {
		return err
	}

	l.Debug("search nested folders", esl.Strings("descendants", z.descendants))
	traverse := make(map[string]bool)
	for _, d := range z.descendants {
		traverse[d] = false
	}
	completed := func() bool {
		for _, t := range traverse {
			if !t {
				return false
			}
		}
		return true
	}
	ErrorScanCompleted := errors.New("scan completed")

	ctx := z.ctx.WithPath(dbx_context.Namespace(z.teamFolderId))
	var scan func(path mo_path.DropboxPath) error
	scan = func(path mo_path.DropboxPath) error {
		entries, err := sv_file.NewFiles(ctx).List(path)
		if err != nil {
			l.Debug("Unable to retrieve entries", esl.Error(err), esl.String("path", path.Path()))
			return err
		}

		// Mark nested folders
		for _, entry := range entries {
			if f, ok := entry.Folder(); ok {
				if f.EntrySharedFolderId != "" {
					traverse[f.EntrySharedFolderId] = true
					rp := path.ChildPath(f.EntryName)
					err := z.addTree(&Tree{
						NamespaceId:   f.EntrySharedFolderId,
						NamespaceName: f.EntryName,
						RelativePath:  teamFolderName + rp.Path(),
					})
					if err != nil {
						return err
					}
				}
			}
		}

		// Return if the scan completed
		if completed() {
			return ErrorScanCompleted
		}

		// Dive into descendants
		for _, entry := range entries {
			if f, ok := entry.Folder(); ok {
				if err := scan(path.ChildPath(f.Name())); err != nil {
					return err
				}
			}
		}
		return nil
	}

	if err := scan(mo_path.NewDropboxPath("")); err != nil && err != ErrorScanCompleted {
		l.Debug("The error occurred on scanning team folder", esl.Error(err))
		return err
	}

	l.Debug("Scan completed")

	return nil
}
