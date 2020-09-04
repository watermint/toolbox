package policy

import (
	"github.com/watermint/toolbox/domain/common/model/mo_filter"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_member_folder"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_team_content"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_teamfolder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type List struct {
	Peer                           dbx_conn.ConnBusinessFile
	Policy                         rp_model.RowReport
	Folder                         mo_filter.Filter
	ErrorUnableToScanMemberFolders app_msg.Message
}

func (z *List) Preset() {
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
}

func (z *List) Exec(c app_control.Control) error {
	l := c.Log()

	teamFolderScanner := uc_teamfolder.New(c, z.Peer.Context())
	teamFolders, err := teamFolderScanner.Scan(z.Folder)
	if err != nil {
		return err
	}

	for _, teamFolder := range teamFolders {
		z.Policy.Row(uc_team_content.NewFolderPolicy(teamFolder.TeamFolder, ""))
		for path, descendant := range teamFolder.NestedFolders {
			z.Policy.Row(uc_team_content.NewFolderPolicy(descendant, path))
		}
	}

	memberFolderScanner := uc_member_folder.New(c, z.Peer.Context())
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
