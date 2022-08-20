package api_auth

type Entity struct {
	// App key name to retrieve client_id/client_secret
	KeyName string

	// Serialized scope
	Scope string

	// Peer name
	PeerName string

	// Serialized credential
	Credential string

	// Supplemental information (e.g. email address of the authenticated account)
	Description string
}

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
