package file

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	namespacefile "github.com/watermint/toolbox/ingredient/ig_dropbox/ig_team/ig_namespace/ig_file"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type List struct {
	Peer                       dbx_conn.ConnScopedTeam
	FileList                   *namespacefile.List
	Folder                     mo_filter.Filter
	ErrorTeamSpaceNotSupported app_msg.Message
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
		dbx_auth.ScopeTeamInfoRead,
	)
	z.Folder.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_filter.NewNameSuffixFilter(),
	)
}

func (z *List) Exec(c app_control.Control) error {
	if ok, _ := sv_team.UnlessTeamFolderApiSupported(z.Peer.Client()); ok {
		c.UI().Error(z.ErrorTeamSpaceNotSupported)
		return errors.New("team space is not supported by this command")
	}

	return rc_exec.Exec(c, z.FileList, func(r rc_recipe.Recipe) {
		rc := r.(*namespacefile.List)
		rc.IncludeTeamFolder = true
		rc.IncludeSharedFolder = false
		rc.Peer = z.Peer
		rc.Folder = z.Folder
	})
}

func (z *List) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}
