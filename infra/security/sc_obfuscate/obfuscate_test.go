package sc_obfuscate

import (
	"testing"

	"github.com/watermint/toolbox/essentials/log/esl"
)

func TestObfuscateDeobfuscate(t *testing.T) {
	logger := esl.Default()
	key := []byte("test-key-1234567890")
	plain := []byte("hello, obfuscate!")

	obfuscated, err := Obfuscate(logger, key, plain)
	if err != nil {
		t.Fatalf("Obfuscate failed: %v", err)
	}

	deobfuscated, err := Deobfuscate(logger, key, obfuscated)
	if err != nil {
		t.Fatalf("Deobfuscate failed: %v", err)
	}

	if string(deobfuscated) != string(plain) {
		t.Errorf("Deobfuscated data does not match original. Got: %s, Want: %s", deobfuscated, plain)
	}
}

func TestDeobfuscateWithWrongKey(t *testing.T) {
	logger := esl.Default()
	key := []byte("test-key-1234567890")
	wrongKey := []byte("wrong-key-0987654321")
	plain := []byte("hello, obfuscate!")

	obfuscated, err := Obfuscate(logger, key, plain)
	if err != nil {
		t.Fatalf("Obfuscate failed: %v", err)
	}

	_, err = Deobfuscate(logger, wrongKey, obfuscated)
	if err == nil {
		t.Error("Deobfuscate should fail with wrong key, but got no error")
	}
}
