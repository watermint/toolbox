package compare

import (
	"github.com/watermint/toolbox/domain/model/mo_file_diff"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/usecase/uc_compare_local"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type LocalVO struct {
	Peer        app_conn.ConnUserFile
	LocalPath   string
	DropboxPath string
}

const (
	reportLocalDiff = "diff"
	reportLocalSkip = "skip"
)

type Local struct {
}

func (z *Local) Requirement() app_vo.ValueObject {
	return &LocalVO{}
}

func (z *Local) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*LocalVO)
	ui := k.UI()

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	repDiff, err := rp_spec_impl.New(z, k.Control()).Open(reportLocalDiff)
	if err != nil {
		return err
	}
	defer repDiff.Close()
	repSkip, err := rp_spec_impl.New(z, k.Control()).Open(reportLocalSkip)
	if err != nil {
		return err
	}
	defer repSkip.Close()

	diff := func(diff mo_file_diff.Diff) error {
		switch diff.DiffType {
		case mo_file_diff.DiffSkipped:
			repSkip.Row(&diff)
		default:
			repDiff.Row(&diff)
		}
		return nil
	}

	ucl := uc_compare_local.New(ctx, k.UI())
	count, err := ucl.Diff(vo.LocalPath, mo_path.NewPath(vo.DropboxPath), diff)
	if err != nil {
		return err
	}
	ui.Info("recipe.file.compare.local.success", app_msg.P{
		"DiffCount": count,
	})
	return nil
}

func (z *Local) Test(c app_control.Control) error {
	return qt_recipe.ScenarioTest()
}

func (z *Local) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportLocalDiff, &mo_file_diff.Diff{}),
		rp_spec_impl.Spec(reportLocalSkip, &mo_file_diff.Diff{}),
	}
}
