package file

import (
	"github.com/watermint/toolbox/domain/usecase/uc_file_upload"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_test"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"os"
)

type UploadVO struct {
	Peer        app_conn.ConnUserFile
	LocalPath   string
	DropboxPath string
	Overwrite   bool
	ChunkSize   int
}

type Upload struct {
}

func (z *Upload) Console() {
}

func (z *Upload) Requirement() app_vo.ValueObject {
	return &UploadVO{
		ChunkSize: 150 * 1024,
	}
}

func (z *Upload) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*UploadVO)
	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}
	opts := make([]uc_file_upload.UploadOpt, 0)
	if vo.ChunkSize > 0 {
		opts = append(opts, uc_file_upload.ChunkSize(vo.ChunkSize*1024))
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
		return qt_test.NotEnoughResource()
	}

	{
		vo := &UploadVO{
			LocalPath:   file,
			DropboxPath: "/" + app_test.TestTeamFolderName,
			Overwrite:   true,
		}
		if !app_test.ApplyTestPeers(c, vo) {
			return qt_test.NotEnoughResource()
		}
		if err := z.Exec(app_kitchen.NewKitchen(c, vo)); err != nil {
			return err
		}
	}

	// Chunked
	{
		vo := &UploadVO{
			LocalPath:   file,
			DropboxPath: "/" + app_test.TestTeamFolderName,
			Overwrite:   true,
			ChunkSize:   1,
		}
		if !app_test.ApplyTestPeers(c, vo) {
			return qt_test.NotEnoughResource()
		}
		if err := z.Exec(app_kitchen.NewKitchen(c, vo)); err != nil {
			return err
		}
	}
	return nil
}

func (z *Upload) Reports() []rp_spec.ReportSpec {
	return uc_file_upload.UploadReports()
}
