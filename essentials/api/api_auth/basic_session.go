package api_auth

type BasicSessionData struct {
	AppData  BasicAppData `json:"app_data"`
	PeerName string       `json:"peer_name"`
}

// BasicSession Basic authentication session
type BasicSession interface {
	Start(session BasicSessionData) (entity BasicEntity, err error)
}
