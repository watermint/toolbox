package member

import (
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_file"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_report"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type RemoveRow struct {
	Email string
}

type RemoveVO struct {
	File     app_file.Data
	Peer     app_conn.ConnBusinessMgmt
	WipeData bool
}

type Remove struct {
}

func (z *Remove) Console() {
}

func (z *Remove) Requirement() app_vo.ValueObject {
	return &RemoveVO{
		WipeData: true,
	}
}

func (z *Remove) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*RemoveVO)

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	svm := sv_member.New(ctx)
	rep, err := k.Report(
		"remove",
		app_report.TransactionHeader(&RemoveRow{}, nil),
	)
	if err != nil {
		return err
	}
	defer rep.Close()

	if err := vo.File.Model(k.Control(), &RemoveRow{}); err != nil {
		return err
	}

	return vo.File.EachRow(func(mod interface{}, rowIndex int) error {
		m := mod.(*RemoveRow)
		mem, err := svm.ResolveByEmail(m.Email)
		if err != nil {
			rep.Failure(api_util.MsgFromError(err), m, nil)
			return nil
		}
		ros := make([]sv_member.RemoveOpt, 0)
		if vo.WipeData {
			ros = append(ros, sv_member.RemoveWipeData())
		}
		err = svm.Remove(mem, ros...)
		if err != nil {
			rep.Failure(api_util.MsgFromError(err), m, nil)
		} else {
			rep.Success(m, nil)
		}
		return nil
	})
}

func (z *Remove) Test(c app_control.Control) error {
	return nil
}
