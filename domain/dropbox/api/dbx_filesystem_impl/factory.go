package dbx_filesystem_impl

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"strings"
	"sync"
)

var (
	resolverMutex = sync.Mutex{}
	resolvers     = make(map[string]dbx_filesystem.RootNamespaceResolver)
)

func GetByEntity(client dbx_client.Client, entity api_auth.OAuthEntity) dbx_filesystem.RootNamespaceResolver {
	hashKey := strings.Join(entity.HashSeed(), "/")
	digestBytes := sha256.Sum256([]byte(hashKey))
	digest := hex.EncodeToString(digestBytes[:])
	resolverMutex.Lock()
	defer resolverMutex.Unlock()

	if resolver, ok := resolvers[digest]; ok {
		return resolver
	}
	concreteResolver := newConcrete(client)
	resolver := NewCached(concreteResolver)
	resolvers[digest] = resolver
	return resolver
}
