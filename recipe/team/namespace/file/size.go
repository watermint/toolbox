package file

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/team/namespace/file"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type Size struct {
	IncludeSharedFolder bool
	IncludeTeamFolder   bool
	IncludeMemberFolder bool
	IncludeAppFolder    bool
	Name                string
	Depth               int
	NamespaceSize       *file.Size
	Peer                rc_conn.ConnBusinessFile
}

func (z *Size) Preset() {
	z.IncludeSharedFolder = true
	z.IncludeTeamFolder = true
	z.Depth = 1
}

func (z *Size) Exec(k rc_kitchen.Kitchen) error {
	return rc_exec.Exec(k.Control(), z.NamespaceSize, func(r rc_recipe.Recipe) {
		rc := r.(*file.Size)
		rc.IncludeSharedFolder = z.IncludeSharedFolder
		rc.IncludeTeamFolder = z.IncludeTeamFolder
		rc.IncludeMemberFolder = z.IncludeMemberFolder
		rc.IncludeAppFolder = z.IncludeAppFolder
		rc.Name = z.Name
		rc.Depth = z.Depth
	})
}

func (z *Size) Test(c app_control.Control) error {
	return qt_endtoend.NoTestRequired()
}
