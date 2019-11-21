package sync

import (
	"github.com/watermint/toolbox/domain/usecase/uc_file_upload"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type UpVO struct {
	Peer        app_conn.ConnUserFile
	LocalPath   string
	DropboxPath string
}

type Up struct {
}

func (z *Up) Console() {
}

func (z *Up) Requirement() app_vo.ValueObject {
	return &UpVO{}
}

func (z *Up) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*UpVO)
	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}
	up := uc_file_upload.New(ctx, rp_spec_impl.New(z, k.Control()), k,
		uc_file_upload.ChunkSize(150*1024),
		uc_file_upload.CreateFolder(),
		uc_file_upload.Overwrite())

	_, err = up.Upload(vo.LocalPath, vo.DropboxPath)
	return err
}

func (z *Up) Test(c app_control.Control) error {
	return qt_recipe.ScenarioTest()
}

func (z *Up) Reports() []rp_spec.ReportSpec {
	return uc_file_upload.UploadReports()
}
