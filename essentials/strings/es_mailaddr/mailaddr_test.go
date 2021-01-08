package es_mailaddr

import "testing"

func TestEscapeSpecial(t *testing.T) {
	// ALPHA/DIGIT only
	{
		if es := EscapeSpecial("test@example.com", "_"); es != "test@example.com" {
			t.Error(es)
		}
		if es := EscapeSpecial("test.test@example.com", "_"); es != "test.test@example.com" {
			t.Error(es)
		}
		if es := EscapeSpecial("0123.test@example.com", "_"); es != "0123.test@example.com" {
			t.Error(es)
		}
	}

	// SPECIAL COMBINED
	{
		if es := EscapeSpecial("#.test@example.com", "_"); es != "_.test@example.com" {
			t.Error(es)
		}
		if es := EscapeSpecial("#.*@example.com", "_"); es != "_._@example.com" {
			t.Error(es)
		}
	}
}

func TestIsEmailAddr(t *testing.T) {
	validEmails := []string{
		"test@example.com",
		"test.test@example.net",
	}
	invalidEmails := []string{
		"test",
	}

	for _, e := range validEmails {
		if x := IsEmailAddr(e); !x {
			t.Error(e, x)
		}
	}
	for _, e := range invalidEmails {
		if x := IsEmailAddr(e); x {
			t.Error(e, x)
		}
	}
}
