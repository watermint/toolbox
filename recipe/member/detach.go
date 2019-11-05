package member

import (
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_test"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_file"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
)

type DetachRow struct {
	Email string `json:"email"`
}

func (z *DetachRow) Validate() (err error) {
	return nil
}

type DetachVO struct {
	File             app_file.Data
	Peer             app_conn.ConnBusinessMgmt
	RevokeTeamShares bool
}

const (
	reportDetach = "detach"
)

type Detach struct {
}

func (z *Detach) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportDetach, rp_model.TransactionHeader(&DetachRow{}, nil)),
	}
}

func (z *Detach) Test(c app_control.Control) error {
	return qt_test.HumanInteractionRequired()
}

func (z *Detach) Console() {
}

func (z *Detach) Requirement() app_vo.ValueObject {
	return &DetachVO{
		RevokeTeamShares: false,
	}
}

func (z *Detach) Exec(k app_kitchen.Kitchen) error {
	mvo := k.Value().(*DetachVO)

	ctx, err := mvo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	svm := sv_member.New(ctx)
	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportDetach)
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
		if !mvo.RevokeTeamShares {
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
