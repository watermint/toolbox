package policy

import (
	"github.com/watermint/toolbox/domain/common/model/mo_filter"
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_team_content"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_teamfolder"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer        dbx_conn.ConnBusinessFile
	Policy      rp_model.RowReport
	Folder      mo_filter.Filter
	ScanTimeout mo_string.SelectString
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
	z.ScanTimeout.SetOptions(string(uc_teamfolder.ScanTimeoutShort),
		string(uc_teamfolder.ScanTimeoutShort),
		string(uc_teamfolder.ScanTimeoutLong),
	)
}

func (z *List) Exec(c app_control.Control) error {
	teamFolderScanner := uc_teamfolder.New(c, z.Peer.Context(), uc_teamfolder.ScanTimeoutMode(z.ScanTimeout.Value()))
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

	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues)
}
