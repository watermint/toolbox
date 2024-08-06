package restore

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_folder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_restore"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_revision"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type MsgRestore struct {
	SkipAlreadyRestored   app_msg.Message
	SkipNotExistOrDeleted app_msg.Message
	SkipIsNotDeleted      app_msg.Message
	ErrorPathNotFound     app_msg.Message
}

var (
	MRestore = app_msg.Apply(&MsgRestore{}).(*MsgRestore)
)

func restoreEntry(entry *mo_file.Deleted, ctx dbx_client.Client, ctl app_control.Control, opLog rp_model.TransactionReport) error {
	l := ctl.Log().With(esl.String("path", entry.EntryPathDisplay))
	target := &TargetPath{
		Path: entry.EntryPathDisplay,
	}
	entryPath := mo_path.NewDropboxPath(entry.EntryPathDisplay)

	revs, err := sv_file_revision.New(ctx).List(entryPath)
	dbxErr := dbx_error.NewErrors(err)
	switch {
	case dbxErr == nil:
		l.Debug("Fall through")

	case dbxErr.Path().IsNotFile():
		l.Debug("Not a file, create the folder")
		folder, err := sv_file_folder.New(ctx).Create(mo_path.NewDropboxPath(entry.EntryPathDisplay))
		dbxErr = dbx_error.NewErrors(err)
		switch {
		case dbxErr == nil:
			opLog.Success(target, folder.Concrete())
			return nil

		case dbxErr.Path().IsConflict():
			// the path already created
			opLog.Skip(MRestore.SkipAlreadyRestored, target)
			return nil

		default:
			l.Debug("Unable to create the folder", esl.Error(err))
			opLog.Failure(err, target)
			return err
		}

	default:
		l.Debug("Unable to retrieve revisions", esl.Error(err))
		opLog.Failure(err, target)
		return err
	}
	if !revs.IsDeleted {
		l.Debug("The file is not deleted")
		opLog.Skip(MRestore.SkipIsNotDeleted, target)
		return nil
	}
	if len(revs.Entries) < 1 {
		l.Debug("No revision found")
		err = errors.New("no revisions found for the file")
		opLog.Failure(err, target)
		return err
	}
	targetRev := revs.Entries[0].Revision
	l.Debug("Restoring to most recent state", esl.String("targetRev", targetRev))

	e, err := sv_file_restore.New(ctx).Restore(entryPath, targetRev)
	dbxErr = dbx_error.NewErrors(err)
	switch {
	case dbxErr == nil:
		opLog.Success(target, e.Concrete())
		return nil

	case dbxErr.IsInvalidRevision():
		opLog.Skip(MRestore.SkipNotExistOrDeleted, target)
		return nil

	default:
		opLog.Failure(err, target)
		return err
	}
}
