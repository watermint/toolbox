package teamfolder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
)

func IsTeamFolderSupported(ctx dbx_client.Client) (bool, error) {
	t, err := IsTeamSpaceSupported(ctx)
	if err != nil {
		return false, err
	}
	return !t, nil
}

func IsTeamSpaceSupported(ctx dbx_client.Client) (bool, error) {
	info, err := sv_team.New(ctx).Feature()
	if err != nil {
		return false, err
	}

	return info.HasTeamSharedDropbox, nil
}
