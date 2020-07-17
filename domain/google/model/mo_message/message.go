package mo_message

import "encoding/json"

type Message struct {
	Raw      json.RawMessage
	Id       string `json:"id" path:"id"`
	ThreadId string `json:"thread_id" path:"threadId"`
	Date     string `json:"date" path:"payload.headers.#(name==\"Date\").value"`
	Subject  string `json:"subject" path:"payload.headers.#(name==\"Subject\").value"`
	To       string `json:"to" path:"payload.headers.#(name==\"To\").value"`
	Cc       string `json:"cc" path:"payload.headers.#(name==\"Cc\").value"`
	From     string `json:"from" path:"payload.headers.#(name==\"From\").value"`
	ReplyTo  string `json:"reply_to" path:"payload.headers.#(name==\"Reply-To\").value"`
}
