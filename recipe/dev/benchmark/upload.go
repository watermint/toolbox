package benchmark

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/filesystem"
	"github.com/watermint/toolbox/domain/dropbox/filesystem/dfs_copier_batch"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
	"github.com/watermint/toolbox/essentials/file/es_sync"
	"github.com/watermint/toolbox/essentials/model/em_file"
	"github.com/watermint/toolbox/essentials/model/em_file_random"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Upload struct {
	rc_recipe.RemarkSecret
	Peer           dbx_conn.ConnScopedIndividual
	Path           mo_path.DropboxPath
	NumFiles       int
	SizeMinKb      int
	SizeMaxKb      int
	Method         mo_string.SelectString
	BlockBlockSize mo_int.RangeInt
	SeqChunkSizeKb mo_int.RangeInt
}

func (z *Upload) Preset() {
	z.Peer.SetScopes(dbx_auth.ScopeFilesContentWrite)
	z.NumFiles = 1000
	z.SizeMinKb = 0
	z.SizeMaxKb = 2 * 1024 // 2MiB
	z.SeqChunkSizeKb.SetRange(1, 150*1024, 64*1024)
	z.BlockBlockSize.SetRange(1, 1000, 50)
	z.Method.SetOptions(
		"block",
		"block",
		"sequential",
	)
}

func (z *Upload) Exec(c app_control.Control) error {
	modelRoot := em_file.NewFolder("data", []em_file.Node{})
	model := em_file_random.NewPoissonTree().Generate(
		em_file_random.NumFiles(z.NumFiles),
		em_file_random.FileSize(int64(z.SizeMinKb*1024), int64(z.SizeMaxKb*1024)),
	)
	model.Rename("Data")
	modelRoot.Add(model)
	var conn es_filesystem.Connector
	switch z.Method.Value() {
	case "block":
		conn = dfs_copier_batch.NewLocalToDropboxBatch(c, z.Peer.Context(), z.BlockBlockSize.Value())

	default:
		conn = filesystem.NewLocalToDropbox(z.Peer.Context(), sv_file_content.ChunkSizeKb(z.SeqChunkSizeKb.Value()))
	}
	copier := filesystem.NewModelToDropbox(c.Log(), modelRoot, conn)
	syncer := es_sync.New(
		c.Log(),
		c.NewQueue(),
		es_filesystem_model.NewFileSystem(modelRoot),
		filesystem.NewFileSystem(z.Peer.Context()),
		copier,
		es_sync.OptimizePreventCreateFolder(!c.Feature().Experiment(app.ExperimentFileSyncDisableReduceCreateFolder)),
	)

	return syncer.Sync(es_filesystem_model.NewPath("/"),
		filesystem.NewPath("", z.Path))
}

func (z *Upload) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Upload{}, func(r rc_recipe.Recipe) {
		m := r.(*Upload)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("benchmark")
	})
}
