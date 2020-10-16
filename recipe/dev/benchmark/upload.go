package benchmark

import (
	"fmt"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
	"github.com/watermint/toolbox/essentials/file/es_sync"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/em_file"
	"github.com/watermint/toolbox/essentials/model/em_file_random"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Upload struct {
	rc_recipe.RemarkSecret
	Peer        dbx_conn.ConnUserFile
	Path        mo_path.DropboxPath
	NumFiles    int
	SizeMinKb   int
	SizeMaxKb   int
	Shard       int
	ChunkSizeKb mo_int.RangeInt
}

func (z *Upload) Preset() {
	z.NumFiles = 1000
	z.SizeMinKb = 0
	z.SizeMaxKb = 2 * 1024 // 2MiB
	z.Shard = 1
	z.ChunkSizeKb.SetRange(1, 150*1024, 64*1024)
}

func (z *Upload) Exec(c app_control.Control) error {
	l := c.Log()
	modelRoot := em_file.NewFolder("data", []em_file.Node{})
	if 1 < z.Shard {
		for i := 0; i < z.Shard; i++ {
			sharedName := fmt.Sprintf("Shard%d", i)
			model := em_file_random.NewPoissonTree().Generate(
				em_file_random.NumFiles(z.NumFiles/z.Shard),
				em_file_random.FileSize(int64(z.SizeMinKb*1024), int64(z.SizeMaxKb*1024)),
			)
			model.Rename(sharedName)
			modelRoot.Add(model)

			sf, err := sv_sharedfolder.New(z.Peer.Context()).Create(z.Path.ChildPath(sharedName))
			if err != nil {
				l.Warn("Unable to create shard folder", esl.Error(err))
				return err
			}
			l.Info("Create shard shared folder created",
				esl.Int("index", i),
				esl.String("shardName", sharedName),
				esl.Any("sharedFolder", sf),
			)
		}
	} else {
		model := em_file_random.NewPoissonTree().Generate(
			em_file_random.NumFiles(z.NumFiles),
			em_file_random.FileSize(int64(z.SizeMinKb*1024), int64(z.SizeMaxKb*1024)),
		)
		model.Rename("Data")
		modelRoot.Add(model)
	}
	copier := filesystem.NewModelToDropbox(modelRoot, z.Peer.Context(), sv_file_content.ChunkSizeKb(z.ChunkSizeKb.Value()))
	syncer := es_sync.New(
		c.Log(),
		c.Sequence(),
		es_filesystem_model.NewFileSystem(modelRoot),
		filesystem.NewFileSystem(z.Peer.Context()),
		copier,
		es_sync.OptimizePreventCreateFolder(c.Feature().Experiment(app.ExperimentFileSyncReduceCreateFolder)),
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
