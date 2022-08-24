package file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client_impl"
	"github.com/watermint/toolbox/domain/dropbox/filesystem"
	"github.com/watermint/toolbox/domain/dropbox/filesystem/dfs_copier_batch"
	"github.com/watermint/toolbox/domain/dropbox/filesystem/dfs_local_to_dbx"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_local"
	"github.com/watermint/toolbox/essentials/file/es_sync"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"path/filepath"
)

type MsgUpload struct {
	SkipExists app_msg.Message
	SkipFilter app_msg.Message
	SkipSame   app_msg.Message
	SkipOther  app_msg.Message
}

var (
	MUpload = app_msg.Apply(&MsgUpload{}).(*MsgUpload)
)

type Upload struct {
	Context     dbx_client.Client
	Delete      bool
	Overwrite   bool
	BatchSize   int
	LocalPath   mo_path2.FileSystemPath
	DropboxPath mo_path.DropboxPath
	Uploaded    rp_model.TransactionReport
	Skipped     rp_model.TransactionReport
	Deleted     rp_model.RowReport
	Summary     rp_model.RowReport
	Name        mo_filter.Filter
}

func (z *Upload) Preset() {
	z.Uploaded.SetModel(&es_filesystem.EntryData{}, &mo_file.ConcreteEntry{}, rp_model.HiddenColumns(
		"input.file_system_type",
		"input.name",
		"input.size",
		"input.mod_time",
		"input.is_file",
		"input.is_folder",
		"result.id",
		"result.tag",
		"result.path_lower",
		"result.revision",
		"result.shared_folder_id",
		"result.parent_shared_folder_id",
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

func (z *Upload) Exec(c app_control.Control) error {
	l := c.Log().With(esl.String("src", z.LocalPath.Path()), esl.String("dest", z.DropboxPath.Path()))
	localPath, err := filepath.Abs(z.LocalPath.Path())
	if err != nil {
		l.Debug("Unable to calc abs path", esl.Error(err), esl.String("localPath", z.LocalPath.Path()))
		return err
	}
	localPath = filepath.Clean(localPath)
	if err := z.Uploaded.Open(rp_model.NoConsoleOutput()); err != nil {
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
	l.Debug("Start uploading")

	var srcFs, tgtFs es_filesystem.FileSystem
	var conn es_filesystem.Connector
	var chunkSizeKb int

	srcFs = es_filesystem_local.NewFileSystem()
	if c.Feature().Experiment(app.ExperimentFileSyncNoCacheDropboxFileSystem) {
		tgtFs = filesystem.NewFileSystem(z.Context)
	} else {
		tgtFs, err = filesystem.NewPreScanFileSystem(c, z.Context, z.DropboxPath)
		if err != nil {
			l.Debug("Failed on the pre-scan", esl.Error(err))
			return err
		}
	}

	if c.Feature().Experiment(app.ExperimentFileSyncLegacyLocalToDropboxConnector) {
		chunkSizeKb = 64 * 1024
		conn = dfs_local_to_dbx.NewLocalToDropbox(z.Context,
			sv_file_content.ChunkSizeKb(chunkSizeKb))
	} else {
		chunkSizeKb = 4 * 1024
		conn = dfs_copier_batch.NewLocalToDropboxBatch(c, z.Context, z.BatchSize)
	}

	mustToDbxEntry := func(entry es_filesystem.Entry) mo_file.Entry {
		e, errConvert := filesystem.ToDropboxEntry(entry)
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
			z.Uploaded.Success(source.AsData(), mustToDbxEntry(target).Concrete())
			status.upload(source.Size(), chunkSizeKb*1024)
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
		es_sync.OptimizePreventCreateFolder(!c.Feature().Experiment(app.ExperimentFileSyncDisableReduceCreateFolder)),
	)

	syncErr := syncer.Sync(es_filesystem_local.NewPath(z.LocalPath.Path()), filesystem.NewPath("", z.DropboxPath))

	if syncErr != nil {
		l.Debug("Sync finished with an error", esl.Error(syncErr))
	}
	status.finish()

	z.Summary.Row(status.summary)
	return syncErr
}

func (z *Upload) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Upload{}, func(r rc_recipe.Recipe) {
		m := r.(*Upload)
		m.Context = dbx_client_impl.NewMock("mock", c)
		m.LocalPath = qtr_endtoend.NewTestFileSystemFolderPath(c, "up")
		m.DropboxPath = qtr_endtoend.NewTestDropboxFolderPath("up")
	})
}
