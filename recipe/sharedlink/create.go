package sharedlink

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/rc_conn"
	"github.com/watermint/toolbox/infra/recpie/rc_kitchen"
	"github.com/watermint/toolbox/infra/recpie/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_time"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type CreateVO struct {
	Peer     rc_conn.ConnUserFile
	Path     string
	TeamOnly bool
	Password string
	Expires  string
}

const (
	reportCreate = "shared_link"
)

type Create struct {
}

func (z *Create) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportCreate, &mo_sharedlink.Metadata{}),
	}
}

func (z *Create) Console() {
}

func (z *Create) Requirement() rc_vo.ValueObject {
	return &CreateVO{}
}

func (z *Create) Exec(k rc_kitchen.Kitchen) error {
	vo := k.Value().(*CreateVO)
	ui := k.UI()
	opts := make([]sv_sharedlink.LinkOpt, 0)

	if vo.Expires != "" {
		if expires, e := ut_time.ParseTimestamp(vo.Expires); e {
			opts = append(opts, sv_sharedlink.Expires(expires))
		} else {
			ui.Error("recipe.sharedlink.create.err.unsupported_time_format", app_msg.P{
				"Input": vo.Expires,
			})
			return errors.New("invalid time format for expires")
		}
	}
	if vo.TeamOnly {
		opts = append(opts, sv_sharedlink.TeamOnly())
	}
	if vo.Password != "" {
		opts = append(opts, sv_sharedlink.Password(vo.Password))
	}

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	link, err := sv_sharedlink.New(ctx).Create(mo_path.NewPath(vo.Path), opts...)
	if err != nil {
		return err
	}
	ui.Info("recipe.sharedlink.create.success", app_msg.P{
		"Url": link.LinkUrl(),
	})

	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportCreate)
	if err != nil {
		return err
	}
	rep.Row(link.Metadata())
	rep.Close()

	return nil
}

func (z *Create) Test(c app_control.Control) error {
	return qt_recipe.ImplementMe()
}
