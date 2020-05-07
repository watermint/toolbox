package member

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

type DetachRow struct {
	Email string `json:"email"`
}

type Detach struct {
	rc_recipe.RemarkIrreversible
	File             fd_file.RowFeed
	Peer             dbx_conn.ConnBusinessMgmt
	RevokeTeamShares bool
	OperationLog     rp_model.TransactionReport
}

func (z *Detach) Preset() {
	z.RevokeTeamShares = false
	z.File.SetModel(&DetachRow{})
	z.OperationLog.SetModel(&DetachRow{}, nil)
}

func (z *Detach) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Detach{}, func(r rc_recipe.Recipe) {
		f, err := qt_file.MakeTestFile("member-detach", "john@example.com\nsmith@example.net\n")
		if err != nil {
			return
		}
		m := r.(*Detach)
		m.File.SetFilePath(f)
		m.RevokeTeamShares = false
	})
	if e, _ := qt_errors.ErrorsForTest(c.Log(), err); e != nil {
		return e
	}
	return qt_errors.ErrorHumanInteractionRequired
}

func (z *Detach) Exec(c app_control.Control) error {
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
