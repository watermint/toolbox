package dbx_filesystem_impl

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"sync"
)

func NewCached(concreteResolver dbx_filesystem.RootNamespaceResolver) dbx_filesystem.RootNamespaceResolver {
	return &cachedRootNamespaceResolver{
		concreteResolver:           concreteResolver,
		teamMemberRootNamespaceIds: make(map[string]string),
	}
}

type cachedRootNamespaceResolver struct {
	individualRootNamespaceId  string
	teamMemberRootNamespaceIds map[string]string
	mutex                      sync.Mutex
	concreteResolver           dbx_filesystem.RootNamespaceResolver
}

func (z *cachedRootNamespaceResolver) ResolveIndividual() (namespaceId string, err error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if z.individualRootNamespaceId != "" {
		return z.individualRootNamespaceId, nil
	}

	namespaceId, err = z.concreteResolver.ResolveIndividual()
	if err != nil {
		return "", err
	}
	z.individualRootNamespaceId = namespaceId
	return namespaceId, nil
}

func (z *cachedRootNamespaceResolver) ResolveTeamMember(teamMemberId string) (namespaceId string, err error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	if z.teamMemberRootNamespaceIds == nil {
		z.teamMemberRootNamespaceIds = make(map[string]string)
	}
	if namespaceId, ok := z.teamMemberRootNamespaceIds[teamMemberId]; ok {
		return namespaceId, nil
	}

	namespaceId, err = z.concreteResolver.ResolveTeamMember(teamMemberId)
	if err != nil {
		return "", err
	}
	z.teamMemberRootNamespaceIds[teamMemberId] = namespaceId
	return namespaceId, nil
}
