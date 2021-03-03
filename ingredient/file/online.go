package file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context_impl"
	"github.com/watermint/toolbox/domain/dropbox/filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_sync"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Online struct {
	Context   dbx_context.Context
	Delete    bool
	Overwrite bool
	SrcPath   mo_path.DropboxPath
	DstPath   mo_path.DropboxPath
	Uploaded  rp_model.TransactionReport
	Skipped   rp_model.TransactionReport
	Deleted   rp_model.RowReport
	Summary   rp_model.RowReport
	Name      mo_filter.Filter
}

func (z *Online) Preset() {
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

func (z *Online) Exec(c app_control.Control) error {
	l := c.Log().With(esl.String("src", z.SrcPath.Path()), esl.String("dest", z.DstPath.Path()))
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

	srcFs := filesystem.NewFileSystem(z.Context)
	tgtFs := filesystem.NewFileSystem(z.Context)
	var conn es_filesystem.Connector
	conn = filesystem.NewDropboxToDropbox(z.Context)

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
			z.Uploaded.Success(mustToDbxEntry(source).Concrete(), mustToDbxEntry(target).Concrete())
			status.copy(source.Size())
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

	syncErr := syncer.Sync(filesystem.NewPath("", z.SrcPath), filesystem.NewPath("", z.DstPath))

	if syncErr != nil {
		l.Debug("Sync finished with an error", esl.Error(syncErr))
	}
	status.finish()

	z.Summary.Row(status.summary)
	return syncErr
}

func (z *Online) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Online{}, func(r rc_recipe.Recipe) {
		m := r.(*Online)
		m.Context = dbx_context_impl.NewMock("mock", c)
		m.SrcPath = qtr_endtoend.NewTestDropboxFolderPath("src")
		m.DstPath = qtr_endtoend.NewTestDropboxFolderPath("dst")
	})
}
