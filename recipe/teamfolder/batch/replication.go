package batch

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/teamfolder"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

type Replication struct {
	rc_recipe.RemarkIrreversible
	File        fd_file.RowFeed
	Replication *teamfolder.Replication
	SrcPeerName string
	DstPeerName string
}

func (z *Replication) Preset() {
	z.SrcPeerName = "src"
	z.DstPeerName = "dst"
	z.File.SetModel(&TeamFolderName{})
}

func (z *Replication) Exec(c app_control.Control) error {
	names := make([]string, 0)
	z.File.EachRow(func(m interface{}, rowIndex int) error {
		r := m.(*TeamFolderName)
		names = append(names, r.Name)
		return nil
	})

	return rc_exec.Exec(c, &teamfolder.Replication{}, func(r rc_recipe.Recipe) {
		rc := r.(*teamfolder.Replication)
		rc.TargetNames = names
		rc.Src.SetPeerName(z.SrcPeerName)
		rc.Dst.SetPeerName(z.DstPeerName)
	})
}

func (z *Replication) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Replication{}, func(r rc_recipe.Recipe) {
		f, err := qt_file.MakeTestFile("replication", "Sales\nMarketing\n")
		if err != nil {
			return
		}
		m := r.(*Replication)
		m.File.SetFilePath(f)
	})
	if e, _ := qt_errors.ErrorsForTest(c.Log(), err); e != nil {
		return e
	}
	return qt_errors.ErrorHumanInteractionRequired
}
