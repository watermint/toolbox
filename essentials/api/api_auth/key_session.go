package api_auth

type KeySessionData struct {
	AppData  KeyAppData `json:"app_data"`
	PeerName string     `json:"peer_name"`
}

type KeySession interface {
	Start(session KeySessionData) (entity KeyEntity, err error)
}
