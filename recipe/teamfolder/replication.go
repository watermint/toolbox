package teamfolder

import (
	"github.com/watermint/toolbox/domain/usecase/uc_teamfolder_mirror"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_test"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type ReplicationVO struct {
	SrcFile app_conn.ConnBusinessFile
	SrcMgmt app_conn.ConnBusinessMgmt
	DstFile app_conn.ConnBusinessFile
	DstMgmt app_conn.ConnBusinessMgmt
	Name    string
}

type Replication struct {
}

func (z *Replication) Console() {
}

func (z *Replication) Requirement() app_vo.ValueObject {
	return &ReplicationVO{}
}

func (z *Replication) Exec(k app_kitchen.Kitchen) error {
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

	ucm := uc_teamfolder_mirror.New(ctxFileSrc, ctxMgmtSrc, ctxFileDst, ctxMgmtDst, k)
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
	return qt_test.HumanInteractionRequired()
}
