package file

import (
	"context"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file_content"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_test"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_model"
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

const (
	reportUpload = "uploaded"
)

type UploadWorker struct {
	basePath  string
	file      string
	ctx       context.Context
	rep       rp_model.Report
	overwrite bool
}

type Upload struct {
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
	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportUpload)
	if err != nil {
		return err
	}
	defer rep.Close()

	up := sv_file_content.NewUpload(ctx, sv_file_content.ChunkSize(int64(vo.ChunkSize*1024)))

	var entry mo_file.Entry
	if vo.Overwrite {
		entry, err = up.Overwrite(mo_path.NewPath(vo.DropboxPath), vo.LocalPath)
		if err != nil {
			return err
		}
	} else {
		entry, err = up.Add(mo_path.NewPath(vo.DropboxPath), vo.LocalPath)
		if err != nil {
			return err
		}
	}
	rep.Row(entry.Concrete())

	return nil
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
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportUpload, &mo_file.ConcreteEntry{}),
	}
}
