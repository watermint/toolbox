package file

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/usecase/uc_file_merge"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type Merge struct {
	Peer                rc_conn.ConnUserFile
	From                mo_path.DropboxPath
	To                  mo_path.DropboxPath
	DryRun              bool
	KeepEmptyFolder     bool
	WithinSameNamespace bool
}

func (z *Merge) Preset() {
	z.DryRun = true
}

func (z *Merge) Console() {
}

func (z *Merge) Exec(k rc_kitchen.Kitchen) error {
	ctx := z.Peer.Context()

	ufm := uc_file_merge.New(ctx, k)
	opts := make([]uc_file_merge.MergeOpt, 0)
	if z.DryRun {
		opts = append(opts, uc_file_merge.DryRun())
	}
	if !z.KeepEmptyFolder {
		opts = append(opts, uc_file_merge.ClearEmptyFolder())
	}
	if z.WithinSameNamespace {
		opts = append(opts, uc_file_merge.WithinSameNamespace())
	}

	return ufm.Merge(z.From, z.To, opts...)
}

func (z *Merge) Test(c app_control.Control) error {
	return qt_endtoend.ScenarioTest()
}
