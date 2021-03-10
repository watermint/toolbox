package benchmark

import (
	"github.com/watermint/toolbox/essentials/file/es_filesystem_copier"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_local"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
	"github.com/watermint/toolbox/essentials/file/es_sync"
	"github.com/watermint/toolbox/essentials/model/em_file_random"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Local struct {
	rc_recipe.RemarkSecret
	Path      mo_path.FileSystemPath
	NumFiles  int
	SizeMinKb int
	SizeMaxKb int
}

func (z *Local) Preset() {
	z.NumFiles = 1000
	z.SizeMinKb = 0
	z.SizeMaxKb = 2 * 1024 // 2MiB
}

func (z *Local) Exec(c app_control.Control) error {
	model := em_file_random.NewPoissonTree().Generate(
		em_file_random.NumFiles(z.NumFiles),
		em_file_random.FileSize(int64(z.SizeMinKb*1024), int64(z.SizeMaxKb*1024)),
	)

	copier := es_filesystem_copier.NewModelToLocal(c.Log(), model)
	syncer := es_sync.New(
		c.Log(),
		c.NewQueue(),
		es_filesystem_model.NewFileSystem(model),
		es_filesystem_local.NewFileSystem(),
		copier,
	)

	return syncer.Sync(es_filesystem_model.NewPath("/"),
		es_filesystem_local.NewPath(z.Path.Path()))
}

func (z *Local) Test(c app_control.Control) error {
	workPath, err := qt_file.MakeTestFolder("local", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(workPath)
	}()

	return rc_exec.ExecMock(c, &Local{}, func(r rc_recipe.Recipe) {
		m := r.(*Local)
		m.NumFiles = 10
		m.SizeMaxKb = 10
		m.Path = mo_path.NewFileSystemPath(workPath)
	})
}
