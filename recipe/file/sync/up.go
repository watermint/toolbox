package sync

import (
	"github.com/watermint/toolbox/domain/usecase/uc_file_upload"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/rc_conn"
	"github.com/watermint/toolbox/infra/recpie/rc_kitchen"
	"github.com/watermint/toolbox/infra/recpie/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type UpVO struct {
	Peer              rc_conn.ConnUserFile
	LocalPath         string
	DropboxPath       string
	ExcludeFolderName string
	ChunkSizeKb       int
}

type Up struct {
	//Peer              app_conn.ConnUserFile
	//LocalPath         string
	//DropboxPath       string
	//ExcludeFolderName string
	//ChunkSizeKb       int
	//Excludes          fd_file.Feed
	//UploadMO          *uc_file_upload.UploadMO
	//Upload            rp_spec.ReportSpec
}

//func (z *Up) Init() {
//	z.Upload = rp_spec_impl.Spec(
//		"",
//		rp_model.TransactionHeader(&uc_file_upload.UploadRow{}, &mo_file.ConcreteEntry{}),
//		rp_model.HiddenColumns(
//			"result.id",
//			"result.tag",
//		),
//	)
//	z.ChunkSizeKb = 150 * 1024
//}

func (z *Up) Console() {
}

func (z *Up) Requirement() rc_vo.ValueObject {
	return &UpVO{
		ChunkSizeKb: 150 * 1024,
	}
}

func (z *Up) Exec(k rc_kitchen.Kitchen) error {
	vo := k.Value().(*UpVO)
	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}
	up := uc_file_upload.New(ctx, rp_spec_impl.New(z, k.Control()), k,
		uc_file_upload.ChunkSizeKb(vo.ChunkSizeKb),
		uc_file_upload.CreateFolder(),
		uc_file_upload.ExcludeFolderName(vo.ExcludeFolderName),
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
