package file

import (
	"github.com/watermint/toolbox/domain/model/mo_file_diff"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/usecase/uc_compare_paths"
	"github.com/watermint/toolbox/domain/usecase/uc_file_mirror"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type ReplicationVO struct {
	Src     rc_conn.OldConnUserFile
	Dst     rc_conn.OldConnUserFile
	SrcPath string
	DstPath string
}

const (
	reportReplication = "replication_diff"
)

type Replication struct {
}

func (z *Replication) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportReplication, &mo_file_diff.Diff{}),
	}
}

func (z *Replication) Console() {
}

func (z *Replication) Requirement() rc_vo.ValueObject {
	return &ReplicationVO{}
}

func (z *Replication) Exec(k rc_kitchen.Kitchen) error {
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

	srcPath := mo_path.NewDropboxPath(vo.SrcPath)
	dstPath := mo_path.NewDropboxPath(vo.DstPath)

	err = uc_file_mirror.New(ctxSrc, ctxDst).Mirror(srcPath, dstPath)
	if err != nil {
		return err
	}
	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportReplication)
	if err != nil {
		return err
	}
	defer rep.Close()

	diff := func(d mo_file_diff.Diff) error {
		rep.Row(&d)
		return nil
	}
	count, err := uc_compare_paths.New(ctxSrc, ctxDst, k.UI()).Diff(srcPath, dstPath, diff)
	ui.Info("recipe.file.replication.done", app_msg.P{
		"DiffCount": count,
	})
	if err != nil {
		return err
	}
	return nil
}

func (z *Replication) Test(c app_control.Control) error {
	return qt_recipe.ImplementMe()
}
