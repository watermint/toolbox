package api_auth

type BasicRepository interface {
	// Put to store the Entity.
	Put(entity BasicEntity)

	// Get to retrieve the Entity.
	Get(keyName string, peerName string) (entity BasicEntity, found bool)

	// Delete to purge the Entity from the repository.
	Delete(keyName string, peerName string)

	// List to retrieve all Entities matches appKeyName.
	List(keyName string) (entities []BasicEntity)

	// Close the repository
	Close()
}
