package dbx_filesystem_impl

import "github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"

func NewEmpty() dbx_filesystem.RootNamespaceResolver {
	return &emptyResolver{}
}

type emptyResolver struct {
}

func (z emptyResolver) ResolveIndividual() (namespaceId string, err error) {
	return "", nil
}

func (z emptyResolver) ResolveTeamMember(teamMemberId string) (namespaceId string, err error) {
	return "", nil
}
