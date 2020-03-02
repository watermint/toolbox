package batch

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/teamfolder"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Replication struct {
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
		rc.SrcFile.SetPeerName(z.SrcPeerName)
		rc.SrcMgmt.SetPeerName(z.SrcPeerName)
		rc.DstFile.SetPeerName(z.DstPeerName)
		rc.DstMgmt.SetPeerName(z.DstPeerName)
	})
}

func (z *Replication) Test(c app_control.Control) error {
	return qt_errors.ErrorHumanInteractionRequired
}
