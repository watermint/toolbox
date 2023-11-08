package api_auth

type KeyCredential struct {
	Key string `json:"key"`
}

type KeyEntity struct {
	KeyName     string        `json:"key_name"`
	PeerName    string        `json:"peer_name"`
	Credential  KeyCredential `json:"credential"`
	Description string        `json:"description,omitempty"`
	Timestamp   string        `json:"timestamp,omitempty"`
}

func NewNoAuthKeyEntity() KeyEntity {
	return KeyEntity{}
}

func (z KeyEntity) Entity() Entity {
	return Entity{
		KeyName:     z.KeyName,
		Scope:       "",
		PeerName:    z.PeerName,
		Credential:  z.Credential.Key,
		Description: z.Description,
		Timestamp:   z.Timestamp,
	}
}

func (z KeyEntity) HashSeed() []string {
	return []string{
		"a", z.KeyName,
		"p", z.PeerName,
		"c", z.Credential.Key,
	}
}
