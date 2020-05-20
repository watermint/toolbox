package app_msg

import "testing"

type TestMO struct {
	Hello Message
	Greet Message
}

type TestNestedMO struct {
	Hey  Message
	Test *TestMO
}

func TestApply(t *testing.T) {
	{
		mo1 := &TestMO{}

		Apply(mo1)

		if mo1.Hello.Key() != "infra.ui.app_msg.test_mo.hello" {
			t.Error(mo1.Hello.Key())
		}
		if mo1.Greet.Key() != "infra.ui.app_msg.test_mo.greet" {
			t.Error(mo1.Greet.Key())
		}
	}

	{
		mo2 := &TestNestedMO{}

		Apply(mo2)

		if mo2.Hey.Key() != "infra.ui.app_msg.test_nested_mo.hey" {
			t.Error(mo2.Hey.Key())
		}
		//if mo2.Test.Hello.Key() != "infra.ui.app_msg.test_mo.hello" {
		//	t.Error(mo2.Test.Hello.Key())
		//}
		//if mo2.Test.Greet.Key() != "infra.ui.app_msg.test_mo.greet" {
		//	t.Error(mo2.Test.Greet.Key())
		//}
	}

}
