package sharedlink

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/model/mo_time"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Create struct {
	Peer     rc_conn.ConnUserFile
	Path     mo_path.DropboxPath
	TeamOnly bool
	Password string
	Expires  mo_time.TimeOptional
	Created  rp_model.RowReport
}

func (z *Create) Preset() {
	z.Created.SetModel(&mo_sharedlink.Metadata{})
}

func (z *Create) Exec(c app_control.Control) error {
	ui := c.UI()
	opts := make([]sv_sharedlink.LinkOpt, 0)

	if z.Expires.Ok() {
		opts = append(opts, sv_sharedlink.Expires(z.Expires.Time()))
	}
	if z.TeamOnly {
		opts = append(opts, sv_sharedlink.TeamOnly())
	}
	if z.Password != "" {
		opts = append(opts, sv_sharedlink.Password(z.Password))
	}

	if err := z.Created.Open(); err != nil {
		return err
	}

	link, err := sv_sharedlink.New(z.Peer.Context()).Create(z.Path, opts...)
	if err != nil {
		return err
	}
	ui.InfoK("recipe.sharedlink.create.success", app_msg.P{
		"Url": link.LinkUrl(),
	})

	z.Created.Row(link.Metadata())
	return nil
}

func (z *Create) Test(c app_control.Control) error {
	return qt_errors.ErrorImplementMe
}
