package api_auth

type OAuthRepository interface {
	// Put to store the Entity.
	Put(entity OAuthEntity)

	// Get to retrieve the Entity.
	Get(keyName string, scopes []string, peerName string) (entity OAuthEntity, found bool)

	// Delete to purge the Entity from the repository.
	Delete(keyName string, scopes []string, peerName string)

	// List to retrieve all Entities matches appKeyName/scope combination.
	List(keyName string, scopes []string) (entities []OAuthEntity)

	// Close the repository
	Close()
}
