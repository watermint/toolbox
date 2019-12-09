package file

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file_content"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

type DownloadVO struct {
	Peer        app_conn.ConnUserFile
	DropboxPath string
	LocalPath   string
}

const (
	reportDownload = "download"
)

type Download struct {
}

func (z *Download) Console() {
}

func (z *Download) Requirement() app_vo.ValueObject {
	return &DownloadVO{}
}

func (z *Download) Exec(k app_kitchen.Kitchen) error {
	l := k.Log()
	vo := k.Value().(*DownloadVO)
	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportDownload)
	if err != nil {
		return err
	}
	defer rep.Close()

	entry, f, err := sv_file_content.NewDownload(ctx).Download(mo_path.NewPath(vo.DropboxPath))
	if err != nil {
		return err
	}
	if err := os.Rename(f, filepath.Join(vo.LocalPath, entry.Name())); err != nil {
		l.Debug("Unable to move file to specified path",
			zap.Error(err),
			zap.String("downloaded", f),
			zap.String("destination", vo.LocalPath),
		)
		return err
	}

	rep.Row(entry.Concrete())
	return nil
}

func (z *Download) Test(c app_control.Control) error {
	return qt_recipe.ImplementMe()
}

func (z *Download) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(
			reportDownload,
			&mo_file.ConcreteEntry{},
		),
	}
}
