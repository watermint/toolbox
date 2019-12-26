package teamfolder

import (
	"github.com/watermint/toolbox/domain/usecase/uc_teamfolder_mirror"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type ReplicationVO struct {
	SrcFile rc_conn.OldConnBusinessFile
	SrcMgmt rc_conn.OldConnBusinessMgmt
	DstFile rc_conn.OldConnBusinessFile
	DstMgmt rc_conn.OldConnBusinessMgmt
	Name    string
}

type Replication struct {
}

func (z *Replication) Reports() []rp_spec.ReportSpec {
	return uc_teamfolder_mirror.ReportSpec()
}

func (z *Replication) Console() {
}

func (z *Replication) Requirement() rc_vo.ValueObject {
	return &ReplicationVO{}
}

func (z *Replication) Exec(k rc_kitchen.Kitchen) error {
	vo := k.Value().(*ReplicationVO)
	ui := k.UI()

	ui.Info("recipe.teamfolder.replication.conn_src_file")
	ctxFileSrc, err := vo.SrcFile.Connect(k.Control())
	if err != nil {
		return err
	}
	ui.Info("recipe.teamfolder.replication.conn_src_mgmt")
	ctxMgmtSrc, err := vo.SrcMgmt.Connect(k.Control())
	if err != nil {
		return err
	}
	ui.Info("recipe.teamfolder.replication.conn_dst_file")
	ctxFileDst, err := vo.DstFile.Connect(k.Control())
	if err != nil {
		return err
	}
	ui.Info("recipe.teamfolder.replication.conn_dst_mgmt")
	ctxMgmtDst, err := vo.DstMgmt.Connect(k.Control())
	if err != nil {
		return err
	}

	rs := rp_spec_impl.New(z, k.Control())
	ucm := uc_teamfolder_mirror.New(ctxFileSrc, ctxMgmtSrc, ctxFileDst, ctxMgmtDst, k, rs)
	uc, err := ucm.PartialScope([]string{vo.Name})
	if err != nil {
		return err
	}
	err = ucm.Mirror(uc)
	if err != nil {
		ui.Failure("recipe.teamfolder.replication.failure")
		return err
	}
	ui.Success("recipe.teamfolder.replication.success")

	return nil
}

func (z *Replication) Test(c app_control.Control) error {
	return qt_endtoend.HumanInteractionRequired()
}
