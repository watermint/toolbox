package file

import (
	"github.com/watermint/toolbox/domain/usecase/uc_file_upload"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"os"
)

type UploadVO struct {
	Peer        rc_conn.OldConnUserFile
	LocalPath   string
	DropboxPath string
	Overwrite   bool
	ChunkSizeKb int
}

type Upload struct {
}

type UploadMO struct {
	Upload *uc_file_upload.UploadMO
}

func (z *Upload) Console() {
}

func (z *Upload) Requirement() rc_vo.ValueObject {
	return &UploadVO{
		ChunkSizeKb: 150 * 1024,
	}
}

func (z *Upload) Exec(k rc_kitchen.Kitchen) error {
	vo := k.Value().(*UploadVO)
	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}
	opts := make([]uc_file_upload.UploadOpt, 0)
	if vo.ChunkSizeKb > 0 {
		opts = append(opts, uc_file_upload.ChunkSizeKb(vo.ChunkSizeKb))
	}
	if vo.Overwrite {
		opts = append(opts, uc_file_upload.Overwrite())
	}

	up := uc_file_upload.New(ctx, rp_spec_impl.New(z, k.Control()), k, opts...)
	_, err = up.Upload(vo.LocalPath, vo.DropboxPath)
	return err
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
		vo := &UploadVO{
			LocalPath:   file,
			DropboxPath: "/" + qt_recipe.TestTeamFolderName,
			Overwrite:   true,
		}
		if !qt_recipe.ApplyTestPeers(c, vo) {
			return qt_endtoend.NotEnoughResource()
		}
		if err := z.Exec(rc_kitchen.NewKitchen(c, vo)); err != nil {
			return err
		}
	}

	// Chunked
	{
		vo := &UploadVO{
			LocalPath:   file,
			DropboxPath: "/" + qt_recipe.TestTeamFolderName,
			Overwrite:   true,
			ChunkSizeKb: 1,
		}
		if !qt_recipe.ApplyTestPeers(c, vo) {
			return qt_endtoend.NotEnoughResource()
		}
		if err := z.Exec(rc_kitchen.NewKitchen(c, vo)); err != nil {
			return err
		}
	}
	return nil
}

func (z *Upload) Reports() []rp_spec.ReportSpec {
	return uc_file_upload.UploadReports()
}
