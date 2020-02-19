package file

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_url"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink_file"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type List struct {
	Peer     rc_conn.ConnUserFile
	Url      mo_url.Url
	Password string
	FileList rp_model.RowReport
}

func (z *List) Preset() {
	z.FileList.SetModel(&mo_file.ConcreteEntry{})
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.FileList.Open(); err != nil {
		return err
	}

	return sv_sharedlink_file.New(z.Peer.Context()).ListRecursive(z.Url, func(entry mo_file.Entry) {
		z.FileList.Row(entry.Concrete())
	}, sv_sharedlink_file.Password(z.Password))
}

func (z *List) Test(c app_control.Control) error {
	return qt_endtoend.ImplementMe()
}
