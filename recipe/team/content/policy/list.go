package policy

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_member_folder"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_team_content"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_teamfolder_scanner"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type List struct {
	Peer                           dbx_conn.ConnScopedTeam
	Policy                         rp_model.RowReport
	Folder                         mo_filter.Filter
	ScanTimeout                    mo_string.SelectString
	ErrorUnableToScanMemberFolders app_msg.Message
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeGroupsRead,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
		dbx_auth.ScopeTeamInfoRead,
	)
	z.Policy.SetModel(
		&uc_team_content.FolderPolicy{},
		rp_model.HiddenColumns(
			"owner_team_id",
			"namespace_id",
			"namespace_name",
		),
	)
	z.Folder.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_filter.NewNameSuffixFilter(),
	)
	z.ScanTimeout.SetOptions(string(uc_teamfolder_scanner.ScanTimeoutShort),
		string(uc_teamfolder_scanner.ScanTimeoutShort),
		string(uc_teamfolder_scanner.ScanTimeoutLong),
	)
}

func (z *List) Exec(c app_control.Control) error {
	l := c.Log()

	teamFolderScanner := uc_teamfolder_scanner.New(c, z.Peer.Client(), uc_teamfolder_scanner.ScanTimeoutMode(z.ScanTimeout.Value()))
	teamFolders, err := teamFolderScanner.Scan(z.Folder)
	if err != nil {
		return err
	}
	if err := z.Policy.Open(); err != nil {
		return err
	}

	for _, teamFolder := range teamFolders {
		z.Policy.Row(uc_team_content.NewFolderPolicy(teamFolder.TeamFolder, ""))
		for path, descendant := range teamFolder.NestedFolders {
			z.Policy.Row(uc_team_content.NewFolderPolicy(descendant, path))
		}
	}

	memberFolderScanner := uc_member_folder.New(c, z.Peer.Client())
	memberFolders, err := memberFolderScanner.Scan(z.Folder)
	if err != nil {
		l.Debug("Failed to scan member folders", esl.Error(err))
		c.UI().Error(z.ErrorUnableToScanMemberFolders.With("Error", err))
		memberFolders = make([]*uc_member_folder.MemberNamespace, 0)
	}

	for _, memberFolder := range memberFolders {
		z.Policy.Row(memberFolder.Namespace)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues)
}
