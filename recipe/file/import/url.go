package _import

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file_url"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type UrlVO struct {
	Peer rc_conn.OldConnUserFile
	Path string
	Url  string
}

const (
	urlReport = "import_url"
)

type Url struct {
}

func (z *Url) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(urlReport, &mo_file.ConcreteEntry{}),
	}
}

func (z *Url) Console() {
}

func (z *Url) Requirement() rc_vo.ValueObject {
	return &UrlVO{}
}

func (z *Url) Exec(k rc_kitchen.Kitchen) error {
	vo := k.Value().(*UrlVO)
	ui := k.UI()

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}
	rep, err := rp_spec_impl.New(z, k.Control()).Open(urlReport)
	if err != nil {
		return err
	}
	defer rep.Close()

	path := sv_file_url.PathWithName(mo_path.NewDropboxPath(vo.Path), vo.Url)
	ui.Info("recipe.file.import.url.progress", app_msg.P{
		"Path": path.Path(),
		"Url":  vo.Url,
	})
	entry, err := sv_file_url.New(ctx).Save(path, vo.Url)
	if err != nil {
		return err
	}
	rep.Row(entry.Concrete())
	return nil
}

func (z *Url) Test(c app_control.Control) error {
	vo := &UrlVO{
		Path: "/" + qt_recipe.TestTeamFolderName + "/file-import-url",
		Url:  "https://dummyimage.com/10x10/000/fff",
	}
	if !qt_recipe.ApplyTestPeers(c, vo) {
		return qt_endtoend.NotEnoughResource()
	}
	if err := z.Exec(rc_kitchen.NewKitchen(c, vo)); err != nil {
		return err
	}
	return nil
}
