package teamspace

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
)

var (
	ErrorRootNamespaceNotFound = errors.New("root namespace not found")
)

func FindRootNamespaceIdAsMember(client dbx_client.Client) (rootNamespaceId string, err error) {
	profile, err := sv_profile.NewProfile(client).Current()
	if err != nil {
		return "", err
	}
	return profile.RootNamespaceId, nil
}

func FindRootNamespaceIdAsAdmin(client dbx_client.Client) (rootNamespaceId string, err error) {
	teamfolders, err := sv_teamfolder.New(client).List()
	if err != nil {
		return "", err
	}

	for _, teamfolder := range teamfolders {
		if teamfolder.IsTeamSharedDropbox {
			return teamfolder.TeamFolderId, nil
		}
	}
	return "", ErrorRootNamespaceNotFound
}

func ClientForRootNamespaceAsMember(client dbx_client.Client) (dbx_client.Client, error) {
	rootNamespaceId, err := FindRootNamespaceIdAsMember(client)
	if rootNamespaceId == "" || err != nil {
		return nil, err
	}
	return client.WithPath(dbx_client.Root(rootNamespaceId)), nil
}

func ClientForRootNamespaceAsAdmin(client dbx_client.Client) (dbx_client.Client, error) {
	admin, err := sv_profile.NewTeam(client).Admin()
	if err != nil {
		return nil, err
	}
	rootNamespaceId, err := FindRootNamespaceIdAsAdmin(client)
	if rootNamespaceId == "" || err != nil {
		return nil, err
	}
	return client.AsAdminId(admin.TeamMemberId).WithPath(dbx_client.Root(rootNamespaceId)), nil
}
