package file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_diff"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_compare_paths"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_file_mirror"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Replication struct {
	rc_recipe.RemarkIrreversible
	Src             dbx_conn.ConnUserFile
	Dst             dbx_conn.ConnUserFile
	SrcPath         mo_path.DropboxPath
	DstPath         mo_path.DropboxPath
	ReplicationDiff rp_model.RowReport
	ProgressDone    app_msg.Message
}

func (z *Replication) Preset() {
	z.ReplicationDiff.SetModel(&mo_file_diff.Diff{})
	z.Src.SetPeerName("src")
	z.Dst.SetPeerName("dst")
}

func (z *Replication) Exec(c app_control.Control) error {
	ui := c.UI()

	ctxSrc := z.Src.Context()
	ctxDst := z.Dst.Context()

	err := uc_file_mirror.New(ctxSrc, ctxDst).Mirror(z.SrcPath, z.DstPath)
	if err != nil {
		return err
	}
	if err := z.ReplicationDiff.Open(); err != nil {
		return err
	}
	diff := func(d mo_file_diff.Diff) error {
		z.ReplicationDiff.Row(&d)
		return nil
	}
	count, err := uc_compare_paths.New(ctxSrc, ctxDst, c.UI()).Diff(z.SrcPath, z.DstPath, diff)
	ui.Info(z.ProgressDone.With("DiffCount", count))
	if err != nil {
		return err
	}
	return nil
}

func (z *Replication) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Replication{}, func(r rc_recipe.Recipe) {
		m := r.(*Replication)
		m.SrcPath = qtr_endtoend.NewTestDropboxFolderPath("src")
		m.DstPath = qtr_endtoend.NewTestDropboxFolderPath("dst")
	})
}
