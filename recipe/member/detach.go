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

type DetachRow struct {
	Email string
}

func (z *DetachRow) Validate() (err error) {
	return nil
}

type DetachVO struct {
	File             app_file.Data
	Peer             app_conn.ConnBusinessMgmt
	RetainTeamShares bool
}

type Detach struct {
}

func (z *Detach) Test(c app_control.Control) error {
	return nil
}

func (z *Detach) Console() {
}

func (z *Detach) Requirement() app_vo.ValueObject {
	return &DetachVO{
		RetainTeamShares: true,
	}
}

func (*Detach) Exec(k app_kitchen.Kitchen) error {
	mvo := k.Value().(*DetachVO)

	ctx, err := mvo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	svm := sv_member.New(ctx)
	rep, err := k.Report(
		"detach",
		app_report.TransactionHeader(&DetachRow{}, nil),
	)
	if err != nil {
		return err
	}
	defer rep.Close()

	if err := mvo.File.Model(k.Control(), &DetachRow{}); err != nil {
		return err
	}

	return mvo.File.EachRow(func(mod interface{}, rowIndex int) error {
		m := mod.(*DetachRow)
		mem, err := svm.ResolveByEmail(m.Email)
		if err != nil {
			rep.Failure(api_util.MsgFromError(err), m, nil)
			return nil
		}
		ros := make([]sv_member.RemoveOpt, 0)
		ros = append(ros, sv_member.Downgrade())
		if mvo.RetainTeamShares {
			ros = append(ros, sv_member.RetainTeamShares())
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
