package member

import (
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type DeleteRow struct {
	Email string `json:"email"`
}

type Delete struct {
	File         fd_file.RowFeed
	Peer         rc_conn.OldConnBusinessMgmt
	WipeData     bool
	OperationLog rp_model.TransactionReport
}

func (z *Delete) Preset() {
	z.WipeData = true
	z.File.SetModel(&DeleteRow{})
	z.OperationLog.Model(&DeleteRow{}, nil)
}

func (z *Delete) Console() {
}

func (z *Delete) Exec(k rc_kitchen.Kitchen) error {
	ctx, err := z.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	svm := sv_member.New(ctx)
	err = z.OperationLog.Open()
	if err != nil {
		return err
	}
	defer z.OperationLog.Close()

	return z.File.EachRow(func(mod interface{}, rowIndex int) error {
		m := mod.(*DeleteRow)
		mem, err := svm.ResolveByEmail(m.Email)
		if err != nil {
			z.OperationLog.Failure(err, m)
			return nil
		}
		ros := make([]sv_member.RemoveOpt, 0)
		if z.WipeData {
			ros = append(ros, sv_member.RemoveWipeData())
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

func (z *Delete) Test(c app_control.Control) error {
	return qt_recipe.HumanInteractionRequired()
}
