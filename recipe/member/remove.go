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

type DeleteRow struct {
	Email string
}

type DeleteVO struct {
	File     app_file.Data
	Peer     app_conn.ConnBusinessMgmt
	WipeData bool
}

const (
	reportDelete = "delete"
)

type Delete struct {
}

func (z *Delete) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportDelete, rp_model.TransactionHeader(&DeleteRow{}, nil)),
	}
}

func (z *Delete) Console() {
}

func (z *Delete) Requirement() app_vo.ValueObject {
	return &DeleteVO{
		WipeData: true,
	}
}

func (z *Delete) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*DeleteVO)

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	svm := sv_member.New(ctx)
	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportDelete)
	if err != nil {
		return err
	}
	defer rep.Close()

	if err := vo.File.Model(k.Control(), &DeleteRow{}); err != nil {
		return err
	}

	return vo.File.EachRow(func(mod interface{}, rowIndex int) error {
		m := mod.(*DeleteRow)
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

func (z *Delete) Test(c app_control.Control) error {
	return qt_test.HumanInteractionRequired()
}
