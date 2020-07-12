package file

import (
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/team/namespace/file"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type List struct {
	Peer                dbx_conn.ConnBusinessFile
	IncludeMediaInfo    bool
	IncludeDeleted      bool
	IncludeMemberFolder bool
	IncludeSharedFolder bool
	IncludeTeamFolder   bool
	Name                mo_string.OptionalString
	NamespaceFileList   *file.List
}

func (z *List) Preset() {
	z.IncludeTeamFolder = true
	z.IncludeSharedFolder = true
	z.IncludeMemberFolder = false
}

func (z *List) Exec(c app_control.Control) error {
	return rc_exec.Exec(c, z.NamespaceFileList, func(r rc_recipe.Recipe) {
		rc := r.(*file.List)
		rc.IncludeMediaInfo = z.IncludeMediaInfo
		rc.IncludeDeleted = z.IncludeDeleted
		rc.IncludeMemberFolder = z.IncludeMemberFolder
		rc.IncludeSharedFolder = z.IncludeSharedFolder
		rc.IncludeTeamFolder = z.IncludeTeamFolder
		rc.Name = z.Name
		rc.Peer = z.Peer
	})
}

func (z *List) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}
