package teamfolder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
)

func IsTeamFolderSupported(ctx dbx_context.Context) (bool, error) {
	t, err := IsTeamSpaceSupported(ctx)
	if err != nil {
		return false, err
	}
	return !t, nil
}

func IsTeamSpaceSupported(ctx dbx_context.Context) (bool, error) {
	info, err := sv_team.New(ctx).Feature()
	if err != nil {
		return false, err
	}

	return info.HasTeamSharedDropbox, nil
}
