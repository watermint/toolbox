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
	return &UploadVO{}
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

	var entry mo_file.Entry
	if vo.Overwrite {
		entry, err = sv_file_content.NewUpload(ctx).Overwrite(mo_path.NewPath(vo.DropboxPath), vo.LocalPath)
		if err != nil {
			return err
		}
	} else {
		entry, err = sv_file_content.NewUpload(ctx).Add(mo_path.NewPath(vo.DropboxPath), vo.LocalPath)
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
	vo := &UploadVO{
		LocalPath:   file,
		DropboxPath: "/" + app_test.TestTeamFolderName,
		Overwrite:   true,
	}
	if !app_test.ApplyTestPeers(c, vo) {
		return qt_test.NotEnoughResource()
	}
	return z.Exec(app_kitchen.NewKitchen(c, vo))
}

func (z *Upload) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportUpload, &mo_file.ConcreteEntry{}),
	}
}
