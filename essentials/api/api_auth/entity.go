package api_auth

type Entity struct {
	// App key name to retrieve client_id/client_secret
	KeyName string `json:"key_name"`

	// Serialized scope
	Scope string `json:"scope,omitempty"`

	// Peer name
	PeerName string `json:"peer_name"`

	// Serialized credential
	Credential string `json:"credential,omitempty"`

	// Supplemental information (e.g. email address of the authenticated account)
	Description string `json:"description,omitempty"`

	// Timestamp of the entity created/updated (RFC3339 format)
	Timestamp string `json:"timestamp,omitempty"`
}

func (z Entity) NoCredential() EntityNoCredential {
	return EntityNoCredential{
		KeyName:     z.KeyName,
		Scope:       z.Scope,
		PeerName:    z.PeerName,
		Description: z.Description,
		Timestamp:   z.Timestamp,
	}
}

type EntityNoCredential struct {
	// App key name to retrieve client_id/client_secret
	KeyName string `json:"key_name"`

	// Serialized scope
	Scope string `json:"scope"`

	// Peer name
	PeerName string `json:"peer_name"`

	// Supplemental information (e.g. email address of the authenticated account)
	Description string `json:"description"`

	// Timestamp of the entity created/updated (RFC3339 format)
	Timestamp string `json:"timestamp"`
}
