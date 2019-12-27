package file

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	namespacefile "github.com/watermint/toolbox/ingredient/team/namespace/file"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type List struct {
	Peer     rc_conn.ConnBusinessFile
	FileList *namespacefile.List
}

func (z *List) Preset() {
}

func (z *List) Exec(k rc_kitchen.Kitchen) error {
	return rc_exec.Exec(k.Control(), z.FileList, func(r rc_recipe.Recipe) {
		rc := r.(*namespacefile.List)
		rc.IncludeTeamFolder = true
		rc.IncludeSharedFolder = false
	})
}

func (z *List) Test(c app_control.Control) error {
	return qt_endtoend.NoTestRequired()
}
