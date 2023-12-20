package file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client_impl"
	"github.com/watermint/toolbox/domain/dropbox/filesystem/dbx_fs"
	"github.com/watermint/toolbox/domain/dropbox/filesystem/dbx_fs_dbx_to_local"
	"github.com/watermint/toolbox/domain/dropbox/filesystem/dbx_fs_dbx_to_local_block"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_local"
	"github.com/watermint/toolbox/essentials/file/es_sync"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"path/filepath"
)

type Download struct {
	Context     dbx_client.Client
	Delete      bool
	Overwrite   bool
	LocalPath   mo_path2.FileSystemPath
	DropboxPath mo_path.DropboxPath
	Downloaded  rp_model.TransactionReport
	Skipped     rp_model.TransactionReport
	Deleted     rp_model.RowReport
	Summary     rp_model.RowReport
	Name        mo_filter.Filter
}

func (z *Download) Preset() {
	z.Downloaded.SetModel(&mo_file.ConcreteEntry{}, &es_filesystem.EntryData{}, rp_model.HiddenColumns(
		"result.file_system_type",
		"result.name",
		"result.size",
		"result.mod_time",
		"result.is_file",
		"result.is_folder",
		"input.id",
		"input.tag",
		"input.path_lower",
		"input.revision",
		"input.shared_folder_id",
		"input.parent_shared_folder_id",
	))
	z.Skipped.SetModel(&es_filesystem.PathData{}, nil, rp_model.HiddenColumns(
		"input.file_system_type",
		"input.attributes",
		"input.entry_namespace.file_system_type",
		"input.entry_namespace.namespace_id",
		"input.entry_namespace.attributes",
	))
	z.Deleted.SetModel(&es_filesystem.PathData{}, rp_model.HiddenColumns(
		"file_system_type",
		"attributes",
		"entry_namespace.file_system_type",
		"entry_namespace.namespace_id",
		"entry_namespace.attributes",
	))
	z.Summary.SetModel(&Summary{})
}

func (z *Download) Exec(c app_control.Control) error {
	l := c.Log().With(esl.String("src", z.LocalPath.Path()), esl.String("dest", z.DropboxPath.Path()))
	localPath, err := filepath.Abs(z.LocalPath.Path())
	if err != nil {
		l.Debug("Unable to calc abs path", esl.Error(err), esl.String("localPath", z.LocalPath.Path()))
		return err
	}
	localPath = filepath.Clean(localPath)
	if err := z.Downloaded.Open(rp_model.NoConsoleOutput()); err != nil {
		return err
	}
	if err := z.Skipped.Open(rp_model.NoConsoleOutput()); err != nil {
		return err
	}
	if err := z.Deleted.Open(rp_model.NoConsoleOutput()); err != nil {
		return err
	}
	if err := z.Summary.Open(); err != nil {
		return err
	}
	l.Debug("Start downloading")

	srcFs := dbx_fs.NewFileSystem(z.Context)
	tgtFs := es_filesystem_local.NewFileSystem()
	var conn es_filesystem.Connector
	if c.Feature().Experiment(app_definitions.ExperimentDbxDownloadBlock) {
		l.Debug("Use block copier")
		conn = dbx_fs_dbx_to_local_block.NewDropboxToLocal(z.Context)
	} else {
		l.Debug("Use standard copier")
		conn = dbx_fs_dbx_to_local.NewDropboxToLocal(z.Context)
	}

	mustToDbxEntry := func(entry es_filesystem.Entry) mo_file.Entry {
		e, errConvert := dbx_fs.ToDropboxEntry(entry)
		if errConvert != nil {
			l.Debug("Unable ot convert", esl.Error(errConvert))
			panic("internal error")
		}
		return e
	}

	status := &Status{}
	status.start()

	syncer := es_sync.New(
		c.Log(),
		c.NewQueue(),
		srcFs,
		tgtFs,
		conn,
		es_sync.SyncDelete(z.Delete),
		es_sync.SyncOverwrite(z.Overwrite),
		es_sync.OnDeleteSuccess(func(target es_filesystem.Path) {
			status.delete()
			z.Deleted.Row(target.AsData())
		}),
		es_sync.OnDeleteFailure(func(target es_filesystem.Path, err es_filesystem.FileSystemError) {
			status.error()
		}),
		es_sync.OnCreateFolderSuccess(func(target es_filesystem.Path) {
			status.createFolder()
		}),
		es_sync.OnCreateFolderFailure(func(target es_filesystem.Path, err es_filesystem.FileSystemError) {
			status.error()
		}),
		es_sync.OnCopySuccess(func(source es_filesystem.Entry, target es_filesystem.Entry) {
			z.Downloaded.Success(mustToDbxEntry(source).Concrete(), target.AsData())
			status.download(source.Size())
		}),
		es_sync.OnCopyFailure(func(source es_filesystem.Path, err es_filesystem.FileSystemError) {
			status.error()
		}),
		es_sync.OnSkip(func(reason es_sync.SkipReason, source es_filesystem.Entry, target es_filesystem.Path) {
			var reasonMsg app_msg.Message
			switch reason {
			case es_sync.SkipExists:
				reasonMsg = MUpload.SkipExists
			case es_sync.SkipFilter:
				reasonMsg = MUpload.SkipFilter
			case es_sync.SkipSame:
				reasonMsg = MUpload.SkipSame
			default:
				reasonMsg = MUpload.SkipOther.With("Reason", reason)
			}
			z.Skipped.Skip(reasonMsg, source.Path().AsData())
			status.skip()
		}),
		es_sync.WithNameFilter(z.Name),
	)

	syncErr := syncer.Sync(dbx_fs.NewPath("", z.DropboxPath), es_filesystem_local.NewPath(z.LocalPath.Path()))
	if syncErr != nil {
		l.Debug("Sync finished with an error", esl.Error(syncErr))
	}
	status.finish()
	z.Summary.Row(status.summary)
	return syncErr
}

func (z *Download) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Download{}, func(r rc_recipe.Recipe) {
		m := r.(*Download)
		m.Context = dbx_client_impl.NewMock("mock", c)
		m.LocalPath = qtr_endtoend.NewTestFileSystemFolderPath(c, "down")
		m.DropboxPath = qtr_endtoend.NewTestDropboxFolderPath("down")
	})
}
