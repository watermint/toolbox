package mo_message

type Message struct {
	Type string `json:"type" path:"type"`
	User string `json:"user" path:"user"`
	Text string `json:"text" path:"text"`
	Ts   string `json:"ts" path:"ts"`
}
