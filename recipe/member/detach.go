package member

import (
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type DetachRow struct {
	Email string `json:"email"`
}

type Detach struct {
	File             fd_file.RowFeed
	Peer             rc_conn.ConnBusinessMgmt
	RevokeTeamShares bool
	OperationLog     rp_model.TransactionReport
}

func (z *Detach) Preset() {
	z.RevokeTeamShares = false
	z.File.SetModel(&DetachRow{})
	z.OperationLog.SetModel(&DetachRow{}, nil)
}

func (z *Detach) Test(c app_control.Control) error {
	return qt_endtoend.HumanInteractionRequired()
}

func (z *Detach) Console() {
}

func (z *Detach) Exec(k rc_kitchen.Kitchen) error {
	ctx := z.Peer.Context()

	svm := sv_member.New(ctx)
	err := z.OperationLog.Open()
	if err != nil {
		return err
	}

	return z.File.EachRow(func(mod interface{}, rowIndex int) error {
		m := mod.(*DetachRow)
		mem, err := svm.ResolveByEmail(m.Email)
		if err != nil {
			z.OperationLog.Failure(err, m)
			return nil
		}
		ros := make([]sv_member.RemoveOpt, 0)
		ros = append(ros, sv_member.Downgrade())
		if !z.RevokeTeamShares {
			ros = append(ros, sv_member.RetainTeamShares())
		}
		err = svm.Remove(mem, ros...)
		if err != nil {
			z.OperationLog.Failure(err, m)
		} else {
			z.OperationLog.Success(m, nil)
		}
		return nil
	})
}
