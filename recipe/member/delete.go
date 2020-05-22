package member

import (
	"github.com/watermint/toolbox/domain/common/model/mo_string"
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

type DeleteRow struct {
	Email string `json:"email"`
}

type Delete struct {
	rc_recipe.RemarkIrreversible
	File                            fd_file.RowFeed
	Peer                            dbx_conn.ConnBusinessMgmt
	TransferDestMember              mo_string.OptionalString
	TransferNotifyAdminEmailOnError mo_string.OptionalString
	WipeData                        bool
	OperationLog                    rp_model.TransactionReport
}

func (z *Delete) Preset() {
	z.WipeData = true
	z.File.SetModel(&DeleteRow{})
	z.OperationLog.SetModel(&DeleteRow{}, nil)
}

func (z *Delete) Exec(c app_control.Control) error {
	ctx := z.Peer.Context()

	svm := sv_member.New(ctx)
	err := z.OperationLog.Open()
	if err != nil {
		return err
	}

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
		if z.TransferDestMember.IsExists() {
			ros = append(ros, sv_member.TransferDest(z.TransferDestMember.Value()))
		}
		if z.TransferNotifyAdminEmailOnError.IsExists() {
			ros = append(ros, sv_member.TransferNotifyAdminOnError(z.TransferNotifyAdminEmailOnError.Value()))
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
	err := rc_exec.ExecReplay(c, &Delete{}, "recipe-member-delete-transfer_success.json.gz", func(r rc_recipe.Recipe) {
		f, err := qt_file.MakeTestFile("member-delete", "john@example.com\nsmith@example.net\n")
		if err != nil {
			return
		}
		m := r.(*Delete)
		m.File.SetFilePath(f)
	})
	if err != nil {
		return err
	}

	err = rc_exec.ExecReplay(c, &Delete{}, "recipe-member-delete-recipient_not_verified.json.gz", func(r rc_recipe.Recipe) {
		f, err := qt_file.MakeTestFile("member-delete", "john@example.com\nsmith@example.net\n")
		if err != nil {
			return
		}
		m := r.(*Delete)
		m.File.SetFilePath(f)
	})
	if err != nil {
		return err
	}

	err = rc_exec.ExecMock(c, &Delete{}, func(r rc_recipe.Recipe) {
		f, err := qt_file.MakeTestFile("member-delete", "john@example.com\nsmith@example.net\n")
		if err != nil {
			return
		}
		m := r.(*Delete)
		m.File.SetFilePath(f)
	})
	if e, _ := qt_errors.ErrorsForTest(c.Log(), err); e != nil {
		return err
	}
	return qt_errors.ErrorHumanInteractionRequired
}
