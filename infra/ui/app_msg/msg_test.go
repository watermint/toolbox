package app_msg

import "testing"

type TestMessageObject struct {
	Hello Message
	Greet Message
}

func TestApply(t *testing.T) {
	{
		mo1 := &TestMessageObject{}

		Apply(mo1)

		if mo1.Hello.Key() != "infra.ui.app_msg.test_message_object.hello" {
			t.Error(mo1.Hello.Key())
		}
		if mo1.Greet.Key() != "infra.ui.app_msg.test_message_object.greet" {
			t.Error(mo1.Greet.Key())
		}
	}
}
