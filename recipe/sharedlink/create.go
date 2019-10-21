package sharedlink

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_time"
)

type CreateVO struct {
	Peer     app_conn.ConnUserFile
	Path     string
	TeamOnly bool
	Password string
	Expires  string
}

type Create struct {
}

func (z *Create) Console() {
}

func (z *Create) Requirement() app_vo.ValueObject {
	return &CreateVO{}
}

func (z *Create) Exec(k app_kitchen.Kitchen) error {
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

	if file, ok := link.File(); ok {
		rep, err := k.Report("shared_link", &mo_sharedlink.File{})
		if err != nil {
			return err
		}
		rep.Row(file)
		rep.Close()
	} else {
		rep, err := k.Report("shared_link", &mo_sharedlink.Folder{})
		if err != nil {
			return err
		}
		rep.Row(file)
		rep.Close()
	}
	return nil
}

func (z *Create) Test(c app_control.Control) error {
	return nil
}
