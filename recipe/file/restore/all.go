package restore

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_folder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_restore"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_revision"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"strings"
)

type TargetPath struct {
	Path string `json:"path"`
}

type All struct {
	rc_recipe.RemarkExperimental
	rc_recipe.RemarkIrreversible
	Peer                  dbx_conn.ConnScopedIndividual
	Path                  mo_path.DropboxPath
	OperationLog          rp_model.TransactionReport
	SkipAlreadyRestored   app_msg.Message
	SkipNotExistOrDeleted app_msg.Message
	SkipIsNotDeleted      app_msg.Message
	ErrorPathNotFound     app_msg.Message
}

func (z *All) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesContentWrite,
	)
	z.OperationLog.SetModel(
		&TargetPath{},
		&mo_file.ConcreteEntry{},
		rp_model.HiddenColumns(
			"result.id",
			"result.path_lower",
			"result.revision",
			"result.content_hash",
			"result.shared_folder_id",
			"result.parent_shared_folder_id",
		),
	)
}

func (z *All) restore(entry *mo_file.Deleted, ctx dbx_client.Client, ctl app_control.Control) error {
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
			z.OperationLog.Success(target, folder.Concrete())
			return nil

		case dbxErr.Path().IsConflict():
			// the path already created
			z.OperationLog.Skip(z.SkipAlreadyRestored, target)
			return nil

		default:
			l.Debug("Unable to create the folder", esl.Error(err))
			z.OperationLog.Failure(err, target)
			return err
		}

	default:
		l.Debug("Unable to retrieve revisions", esl.Error(err))
		z.OperationLog.Failure(err, target)
		return err
	}
	if !revs.IsDeleted {
		l.Debug("The file is not deleted")
		z.OperationLog.Skip(z.SkipIsNotDeleted, target)
		return nil
	}
	if len(revs.Entries) < 1 {
		l.Debug("No revision found")
		err = errors.New("no revisions found for the file")
		z.OperationLog.Failure(err, target)
		return err
	}
	targetRev := revs.Entries[0].Revision
	l.Debug("Restoring to most recent state", esl.String("targetRev", targetRev))

	e, err := sv_file_restore.New(ctx).Restore(entryPath, targetRev)
	dbxErr = dbx_error.NewErrors(err)
	switch {
	case dbxErr == nil:
		z.OperationLog.Success(target, e.Concrete())
		return nil

	case dbxErr.IsInvalidRevision():
		z.OperationLog.Skip(z.SkipNotExistOrDeleted, target)
		return nil

	default:
		z.OperationLog.Failure(err, target)
		return err
	}
}

func (z *All) Exec(c app_control.Control) error {
	l := c.Log()
	ctx := z.Peer.Client()
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	searchBasePath := z.Path
	svf := sv_file.NewFiles(z.Peer.Client())
	for {
		if searchBasePath.IsRoot() {
			break
		}

		m, err := svf.Resolve(searchBasePath)
		dbxErr := dbx_error.NewErrors(err)
		if err == nil {
			searchBasePath = mo_path.NewDropboxPath(m.PathDisplay())
			l.Debug("Restore search from the path", esl.Any("meta", m))
			break
		}

		switch {
		case dbxErr.Path().IsNotFound():
			searchBasePath = searchBasePath.Parent()
			l.Debug("Try with ascendant", esl.String("path", searchBasePath.Path()))
		default:
			l.Debug("Other error, fail", esl.Error(err))
			return err
		}
	}

	targetPathLower := strings.ToLower(z.Path.Path())
	isTargetPath := func(p string) bool {
		if z.Path.IsRoot() {
			return true
		}
		return strings.HasPrefix(strings.ToLower(p), targetPathLower)
	}

	var lastErr error
	proceed := false
	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("restore", z.restore, ctx, c)
		q := s.Get("restore")

		lastErr = sv_file.NewFiles(ctx).ListEach(
			searchBasePath,
			func(entry mo_file.Entry) {
				if !isTargetPath(entry.PathLower()) {
					l.Debug("Skip non target path", esl.String("entryPath", entry.PathDisplay()))
					return
				}
				if d, e := entry.Deleted(); e {
					proceed = true
					q.Enqueue(d)
				}
			},
			sv_file.IncludeDeleted(true),
			sv_file.Recursive(true),
		)
	})

	if !proceed {
		c.UI().Error(z.ErrorPathNotFound.With("Path", z.Path.Path()))
		return errors.New("no deleted file or folder found in the path")
	}

	return lastErr
}

func (z *All) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &All{}, func(r rc_recipe.Recipe) {
		m := r.(*All)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("file-restore")
	})
}
