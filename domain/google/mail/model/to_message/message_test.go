package to_message

import (
	"encoding/json"
	"testing"
)

func TestMessagePart_With(t *testing.T) {
	m := MessagePart{}.WithTo(Address("Test To", "test-to@example.com")).
		WithCc(Address("Test Cc", "test-cc@example.com"), Address("Test Cc 2", "test-cc2@example.com")).
		WithBcc(Address("", "test-bcc@example.com")).
		WithFrom(Address("Test", "test-from@example.com")).
		WithSubject("Test mail")

	p, err := json.Marshal(m)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(p))
}
