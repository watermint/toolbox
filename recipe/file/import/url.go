package _import

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file_url"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type UrlVO struct {
	Peer app_conn.ConnUserFile
	Path string
	Url  string
}

type Url struct {
}

func (z *Url) Console() {
}

func (z *Url) Requirement() app_vo.ValueObject {
	return &UrlVO{}
}

func (z *Url) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*UrlVO)

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}
	rep, err := k.Report("import_url", &mo_file.ConcreteEntry{})
	if err != nil {
		return err
	}
	defer rep.Close()

	entry, err := sv_file_url.New(ctx).Save(mo_path.NewPath(vo.Path), vo.Url)
	if err != nil {
		return err
	}
	rep.Row(entry.Concrete())
	return nil
}

func (z *Url) Test(c app_control.Control) error {
	return nil
}
