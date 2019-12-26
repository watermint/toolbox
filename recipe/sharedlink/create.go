package sharedlink

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_time"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type Create struct {
	Peer     rc_conn.ConnUserFile
	Path     mo_path.DropboxPath
	TeamOnly bool
	Password string
	Expires  string
	Created  rp_model.RowReport
}

func (z *Create) Preset() {
	z.Created.SetModel(&mo_sharedlink.Metadata{})
}

func (z *Create) Console() {
}

func (z *Create) Exec(k rc_kitchen.Kitchen) error {
	ui := k.UI()
	opts := make([]sv_sharedlink.LinkOpt, 0)

	if z.Expires != "" {
		if expires, e := ut_time.ParseTimestamp(z.Expires); e {
			opts = append(opts, sv_sharedlink.Expires(expires))
		} else {
			ui.Error("recipe.sharedlink.create.err.unsupported_time_format", app_msg.P{
				"Input": z.Expires,
			})
			return errors.New("invalid time format for expires")
		}
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
	ui.Info("recipe.sharedlink.create.success", app_msg.P{
		"Url": link.LinkUrl(),
	})

	z.Created.Row(link.Metadata())
	return nil
}

func (z *Create) Test(c app_control.Control) error {
	return qt_endtoend.ImplementMe()
}
