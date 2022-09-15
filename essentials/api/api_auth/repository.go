package api_auth

type Repository interface {
	// Put to store the Entity.
	Put(entity Entity)

	// Get to retrieve the Entity.
	Get(keyName, scope, peerName string) (entity Entity, found bool)

	// Delete to purge the Entity from the repository.
	Delete(keyName, scope, peerName string)

	// List to retrieve all Entities matches appKeyName/scope combination.
	List(keyName, scope string) (entities []Entity)

	// Close the repository
	Close()
}

type RepositoryTraversable interface {
	All() (entities []Entity)
}
