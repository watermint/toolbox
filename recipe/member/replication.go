package member

import (
	"github.com/watermint/toolbox/domain/usecase/uc_member_mirror"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type ReplicationRow struct {
	SrcEmail string `json:"src_email"`
	DstEmail string `json:"dst_email"`
}

type Replication struct {
	Src          rc_conn.ConnBusinessFile
	Dst          rc_conn.ConnBusinessFile
	File         fd_file.RowFeed
	OperationLog rp_model.TransactionReport
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

		ui.InfoK("recipe.member.replication.progress", app_msg.P{
			"SrcEmail": row.SrcEmail,
			"DstEmail": row.DstEmail,
		})
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
	return qt_endtoend.HumanInteractionRequired()
}
