package batch

import (
	"encoding/csv"
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_url"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"os"
	"path/filepath"
)

type MsgUrl struct {
	ProgressImport   app_msg.Message
	ErrorPathMissing app_msg.Message
}

var (
	MUrl = app_msg.Apply(&MsgUrl{}).(*MsgUrl)
)

type UrlRow struct {
	Url  string `json:"url"`
	Path string `json:"path"`
}

type UrlWorker struct {
	row *UrlRow
	ctx dbx_client.Client
	ctl app_control.Control
	rep rp_model.TransactionReport
}

func (z *UrlWorker) Exec() error {
	ui := z.ctl.UI()

	path := sv_file_url.PathWithName(mo_path.NewDropboxPath(z.row.Path), z.row.Url)
	ui.Progress(MUrl.ProgressImport.With("Url", z.row.Url).With("Path", path.Path()))

	entry, err := sv_file_url.New(z.ctx).Save(path, z.row.Url)
	if err != nil {
		z.rep.Failure(err, z.row)
		return err
	}
	z.rep.Success(z.row, entry.Concrete())

	return nil
}

type Url struct {
	rc_recipe.RemarkIrreversible
	Peer            dbx_conn.ConnScopedIndividual
	File            fd_file.RowFeed
	Path            mo_string.OptionalString
	OperationLog    rp_model.TransactionReport
	SkipPathMissing app_msg.Message
	BasePath        mo_string.SelectString
}

func (z *Url) Preset() {
	z.OperationLog.SetModel(
		&UrlRow{},
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
	z.File.SetModel(&UrlRow{})
	z.Peer.SetScopes(dbx_auth.ScopeFilesContentWrite)
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Url) process(row *UrlRow) error {
	path := sv_file_url.PathWithName(mo_path.NewDropboxPath(row.Path), row.Url)

	client := z.Peer.Client().BaseNamespace(dbx_filesystem.AsNamespaceType(z.BasePath.Value()))

	entry, err := sv_file_url.New(client).Save(path, row.Url)
	if err != nil {
		z.OperationLog.Failure(err, row)
		return err
	}
	z.OperationLog.Success(row, entry.Concrete())
	return nil
}

func (z *Url) Exec(c app_control.Control) error {
	ui := c.UI()
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	var lastErr error
	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("url", z.process)
		q := s.Get("url")
		lastErr = z.File.EachRow(func(m interface{}, rowIndex int) error {
			r := m.(*UrlRow)
			var path string
			switch {
			case r.Path != "":
				path = r.Path
			case z.Path.IsExists():
				path = z.Path.Value()
			default:
				z.OperationLog.Skip(z.SkipPathMissing, r)
				ui.Error(MUrl.ErrorPathMissing)
				return errors.New("no path to save")
			}

			q.Enqueue(&UrlRow{
				Url:  r.Url,
				Path: path,
			})
			return nil
		})
	})

	return lastErr
}

func (z *Url) Test(c app_control.Control) error {
	testFilePath := filepath.Join(c.Workspace().Test(), "batch.csv")
	testFile, err := os.Create(testFilePath)
	if err != nil {
		return err
	}
	testCsv := csv.NewWriter(testFile)
	testCsv.Write([]string{"https://dummyimage.com/10x10/000/fff", "/" + qtr_endtoend.TestTeamFolderName + "/file-import-batch-url/fff.png"})
	testCsv.Write([]string{"https://dummyimage.com/10x10/000/eee", "/" + qtr_endtoend.TestTeamFolderName + "/file-import-batch-url/eee.png"})
	testCsv.Flush()

	err = rc_exec.Exec(c, &Url{}, func(r rc_recipe.Recipe) {
		ru := r.(*Url)
		ru.File.SetFilePath(testFilePath)
	})
	if err != nil {
		return err
	}

	return nil
}
