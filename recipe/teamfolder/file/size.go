package file

import (
	"github.com/watermint/toolbox/domain/common/model/mo_int"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	namespacefile "github.com/watermint/toolbox/ingredient/team/namespace/file"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"math"
)

type Size struct {
	Peer     dbx_conn.ConnBusinessFile
	FileSize *namespacefile.Size
	Depth    mo_int.RangeInt
}

func (z *Size) Preset() {
	z.Depth.SetRange(1, math.MaxInt32, 1)
}

func (z *Size) Exec(c app_control.Control) error {
	return rc_exec.Exec(c, z.FileSize, func(r rc_recipe.Recipe) {
		rc := r.(*namespacefile.Size)
		rc.IncludeSharedFolder = false
		rc.IncludeTeamFolder = false
		rc.Depth = z.Depth.Value()
		rc.Peer = z.Peer
	})
}

func (z *Size) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}
