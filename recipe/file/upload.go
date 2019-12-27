package file

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/file"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"os"
)

type Upload struct {
	Peer        rc_conn.ConnUserFile
	LocalPath   mo_path.FileSystemPath
	DropboxPath mo_path.DropboxPath
	Overwrite   bool
	ChunkSizeKb int
	Upload      *file.Upload
}

func (z *Upload) Preset() {
	z.ChunkSizeKb = 150 * 1024
}

func (z *Upload) Console() {
}

func (z *Upload) Exec(c app_control.Control) error {
	return rc_exec.Exec(c, z.Upload, func(r rc_recipe.Recipe) {
		ru := r.(*file.Upload)
		ru.EstimateOnly = false
		ru.LocalPath = z.LocalPath
		ru.DropboxPath = z.DropboxPath
		ru.Overwrite = z.Overwrite
		ru.CreateFolder = false
		ru.Context = z.Peer.Context()
		if z.ChunkSizeKb > 0 {
			ru.ChunkSizeKb = z.ChunkSizeKb
		}
	})
}

func (z *Upload) Test(c app_control.Control) error {
	l := c.Log()
	fileCandidates := []string{"README.md", "upload.go", "upload_test.go"}
	file := ""
	for _, f := range fileCandidates {
		if _, err := os.Lstat(f); err == nil {
			file = f
			break
		}
	}
	if file == "" {
		l.Warn("No file to upload")
		return qt_endtoend.NotEnoughResource()
	}

	{
		err := rc_exec.Exec(c, &Upload{}, func(r rc_recipe.Recipe) {
			ru := r.(*Upload)
			ru.LocalPath = mo_path.NewFileSystemPath(file)
			ru.DropboxPath = mo_path.NewDropboxPath("/" + qt_recipe.TestTeamFolderName)
			ru.Overwrite = true
		})
		if err != nil {
			return err
		}
	}

	// Chunked
	{
		err := rc_exec.Exec(c, &Upload{}, func(r rc_recipe.Recipe) {
			ru := r.(*Upload)
			ru.LocalPath = mo_path.NewFileSystemPath(file)
			ru.DropboxPath = mo_path.NewDropboxPath("/" + qt_recipe.TestTeamFolderName)
			ru.Overwrite = true
			ru.ChunkSizeKb = 1
		})
		if err != nil {
			return err
		}
	}
	return nil
}
