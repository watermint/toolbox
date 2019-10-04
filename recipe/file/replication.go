package file

import (
	"github.com/watermint/toolbox/domain/model/mo_file_diff"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/usecase/uc_file_compare"
	"github.com/watermint/toolbox/domain/usecase/uc_file_mirror"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type ReplicationVO struct {
	Src     app_conn.ConnUserFile
	Dst     app_conn.ConnUserFile
	SrcPath string
	DstPath string
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

	ui.Info("recipe.file.replication.conn_src")
	ctxSrc, err := vo.Src.Connect(k.Control())
	if err != nil {
		return err
	}

	ui.Info("recipe.file.replication.conn_dst")
	ctxDst, err := vo.Dst.Connect(k.Control())
	if err != nil {
		return err
	}

	srcPath := mo_path.NewPath(vo.SrcPath)
	dstPath := mo_path.NewPath(vo.DstPath)

	err = uc_file_mirror.New(ctxSrc, ctxDst).Mirror(srcPath, dstPath)
	if err != nil {
		return err
	}
	rep, err := k.Report("replication_diff", &mo_file_diff.Diff{})
	if err != nil {
		return err
	}
	defer rep.Close()

	diff := func(d mo_file_diff.Diff) error {
		rep.Row(&d)
		return nil
	}
	count, err := uc_file_compare.New(ctxSrc, ctxDst).Diff(
		diff,
		uc_file_compare.LeftPath(srcPath),
		uc_file_compare.RightPath(dstPath),
	)
	ui.Info("recipe.file.replication.done", app_msg.P{
		"DiffCount": count,
	})
	if err != nil {
		return err
	}
	return nil
}

func (z *Replication) Test(c app_control.Control) error {
	return nil
}
