package member

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_member_mirror"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type ReplicationRow struct {
	SrcEmail string `json:"src_email"`
	DstEmail string `json:"dst_email"`
}

type Replication struct {
	rc_recipe.RemarkIrreversible
	Src                 dbx_conn.ConnBusinessFile
	Dst                 dbx_conn.ConnBusinessFile
	File                fd_file.RowFeed
	OperationLog        rp_model.TransactionReport
	ProgressReplication app_msg.Message
}

func (z *Replication) Preset() {
	z.File.SetModel(&ReplicationRow{})
	z.OperationLog.SetModel(&ReplicationRow{}, nil)
	z.Src.SetPeerName("src")
	z.Dst.SetPeerName("dst")
}

func (z *Replication) Exec(c app_control.Control) error {
	ui := c.UI()

	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	return z.File.EachRow(func(m interface{}, rowIndex int) error {
		row := m.(*ReplicationRow)

		ui.Progress(z.ProgressReplication.With("SrcEmail", row.SrcEmail).With("DstEmail", row.DstEmail))
		err := uc_member_mirror.New(z.Src.Context(), z.Dst.Context()).Mirror(row.SrcEmail, row.DstEmail)
		if err != nil {
			z.OperationLog.Failure(err, row)
			return err
		}
		z.OperationLog.Success(row, nil)
		return nil
	})
}

func (z *Replication) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Replication{}, func(r rc_recipe.Recipe) {
		f, err := qt_file.MakeTestFile("member-replication", "john@example.com,smith@example.net\n")
		if err != nil {
			return
		}
		m := r.(*Replication)
		m.File.SetFilePath(f)
	})
	if e, _ := qt_recipe.RecipeError(c.Log(), err); e != nil {
		return e
	}
	return qt_errors.ErrorHumanInteractionRequired
}
