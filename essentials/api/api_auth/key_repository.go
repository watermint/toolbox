package api_auth

type KeyRepository interface {
	// Put to store the Entity.
	Put(entity KeyEntity)

	// Get to retrieve the Entity.
	Get(keyName, peerName string) (entity KeyEntity, found bool)

	// Delete to purge the Entity from the repository.
	Delete(keyName, peerName string)

	// List to retrieve all Entities matches appKeyName.
	List(keyName string) (entities []KeyEntity)

	// Close the repository
	Close()
}
