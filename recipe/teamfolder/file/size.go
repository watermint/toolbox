package file

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	namespacefile "github.com/watermint/toolbox/ingredient/team/namespace/file"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type Size struct {
	Peer     rc_conn.ConnBusinessFile
	FileSize *namespacefile.Size
	Depth    int
}

func (z *Size) Preset() {
	z.Depth = 1
}

func (z *Size) Exec(c app_control.Control) error {
	return rc_exec.Exec(c, z.FileSize, func(r rc_recipe.Recipe) {
		rc := r.(*namespacefile.Size)
		rc.IncludeSharedFolder = false
		rc.IncludeTeamFolder = false
		rc.Depth = z.Depth
	})
}

func (z *Size) Test(c app_control.Control) error {
	return qt_endtoend.NoTestRequired()
}
