package restore

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"strings"
)

type Ext struct {
	rc_recipe.RemarkExperimental
	rc_recipe.RemarkIrreversible
	Peer         dbx_conn.ConnScopedIndividual
	Path         mo_path.DropboxPath
	Ext          string
	OperationLog rp_model.TransactionReport
	BasePath     mo_string.SelectString
}

func (z *Ext) Preset() {
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
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Ext) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	searchBasePath := z.Path
	client := z.Peer.Client().BaseNamespace(dbx_filesystem.AsNamespaceType(z.BasePath.Value()))
	svf := sv_file.NewFiles(client)
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
	targetExtLower := strings.ToLower(z.Ext)
	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("restore", restoreEntry, client, c, z.OperationLog)
		q := s.Get("restore")

		lastErr = sv_file.NewFiles(client).ListEach(
			searchBasePath,
			func(entry mo_file.Entry) {
				if !isTargetPath(entry.PathLower()) {
					l.Debug("Skip non target path", esl.String("entryPath", entry.PathDisplay()))
					return
				}
				if !strings.HasSuffix(entry.PathLower(), targetExtLower) {
					l.Debug("Skip non target extension", esl.String("entryPath", entry.PathDisplay()))
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
		c.UI().Error(MRestore.ErrorPathNotFound.With("Path", z.Path.Path()))
		return errors.New("no deleted file or folder found in the path")
	}

	return lastErr
}

func (z *Ext) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Ext{}, func(r rc_recipe.Recipe) {
		m := r.(*Ext)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("file-restore")
		m.Ext = "txt"
	})
}
