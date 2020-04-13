package file

import (
	"github.com/watermint/toolbox/domain/common/model/mo_int"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/team/namespace/file"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"math"
)

type Size struct {
	IncludeSharedFolder bool
	IncludeTeamFolder   bool
	IncludeMemberFolder bool
	IncludeAppFolder    bool
	Name                string
	Depth               mo_int.RangeInt
	NamespaceSize       *file.Size
	Peer                dbx_conn.ConnBusinessFile
}

func (z *Size) Preset() {
	z.IncludeSharedFolder = true
	z.IncludeTeamFolder = true
	z.Depth.SetRange(1, math.MaxInt32, 1)
}

func (z *Size) Exec(c app_control.Control) error {
	return rc_exec.Exec(c, z.NamespaceSize, func(r rc_recipe.Recipe) {
		rc := r.(*file.Size)
		rc.IncludeSharedFolder = z.IncludeSharedFolder
		rc.IncludeTeamFolder = z.IncludeTeamFolder
		rc.IncludeMemberFolder = z.IncludeMemberFolder
		rc.IncludeAppFolder = z.IncludeAppFolder
		rc.Name = z.Name
		rc.Depth = z.Depth.Value()
	})
}

func (z *Size) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}
